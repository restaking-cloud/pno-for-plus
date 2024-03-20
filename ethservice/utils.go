package ethservice

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"time"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func (e *EthService) waitTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	logger := e.log.WithField("tx", tx.Hash().Hex())
	for {
		receipt, err := e.client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt, nil
		}

		if errors.Is(err, ethereum.NotFound) {
			logger.Trace("Transaction not yet mined")
		} else {
			logger.Trace("Receipt retrieval failed", "err", err)
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func (e *EthService) transact(context context.Context, tx *types.Transaction, pk *ecdsa.PrivateKey) (signedTx *types.Transaction, err error) {

	if pk == nil {
		return signedTx, errors.New("private key provided for transaction is nil")
	}

	walletAddress := crypto.PubkeyToAddress(*pk.Public().(*ecdsa.PublicKey))
	
	gasPrice, err := e.client.SuggestGasPrice(context)
	if err != nil {
		return signedTx, fmt.Errorf("failed to retrieve current gas price: %w", err)
	}
	if e.cfg.MaxGasPrice != nil && (e.cfg.MaxGasPrice.Sign() > 0) && gasPrice.Cmp(e.cfg.MaxGasPrice) > 0 {
		return signedTx, fmt.Errorf("gas price (%s) is higher than max gas price (%s)", gasPrice.String(), e.cfg.MaxGasPrice.String())
	}
	gasTip, err := e.client.SuggestGasTipCap(context)
	if err != nil {
		return signedTx, fmt.Errorf("failed to suggest gas tip: %w", err)
	}
	gasLimit, err := e.client.EstimateGas(context, ethereum.CallMsg{
		From: walletAddress,
		To:   tx.To(),
		Data: tx.Data(),
		Value: tx.Value(),
	})
	if err != nil {
		return signedTx, fmt.Errorf("failed to estimate gas: %w", err)
	}
	nonce, err := e.client.PendingNonceAt(context, walletAddress)
	if err != nil {
		return signedTx, fmt.Errorf("failed to get nonce: %w", err)
	}

	fullTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   e.cfg.ChainID,
		Nonce:     nonce,
		GasTipCap: gasTip,
		GasFeeCap: gasPrice,
		Gas:       gasLimit,
		Data:      tx.Data(),
		To:        tx.To(),
	})

	// calculate the total transaction cost to check if the wallet has enough balance
	maxTxCost := new(big.Int).Mul(gasPrice, new(big.Int).SetUint64(gasLimit))
	txCost := new(big.Int).Add(maxTxCost, tx.Value())
	balance, err := e.client.BalanceAt(context, walletAddress, nil)
	if err != nil {
		return signedTx, fmt.Errorf("failed to get wallet balance: %w", err)
	}
	if balance.Cmp(txCost) < 0 {
		return signedTx, fmt.Errorf("wallet balance (%s) is lower than the total transaction cost (%s)", new(big.Float).Quo(new(big.Float).SetInt(balance), new(big.Float).SetInt64(1e18)).String(), new(big.Float).Quo(new(big.Float).SetInt(txCost), new(big.Float).SetInt64(1e18)).String())
	}

	signer := types.LatestSignerForChainID(e.cfg.ChainID)

	signedTx, err = types.SignTx(fullTx, signer, pk)
	if err != nil {
		return signedTx, fmt.Errorf("failed to sign tx: %w", err)
	}

	logger := e.log.WithField("tx", signedTx.Hash().Hex())
	var pending bool

	sendErr := e.client.SendTransaction(context, signedTx)
	if sendErr != nil {
		// check if the transaction was already sent using the hash and if in pending ignore the error
		// this is to handle the case where the transaction was sent but the response was lost or returned a bug error
		_, pending, err = e.client.TransactionByHash(context, signedTx.Hash())
		if err != nil {
			return signedTx, fmt.Errorf("failed to send tx: %w", sendErr)
		}
			
	}
	
	logger.WithField("pending", pending).Info("PNO Module EthService: Transaction sent")

	return signedTx, nil

}

func (e *EthService) transactAndWait(context context.Context, tx *types.Transaction, pk *ecdsa.PrivateKey) (executedTx *types.Transaction, err error) {

	executedTx, err = e.transact(context, tx, pk)
	if err != nil {
		return executedTx, err
	}

	logger := e.log.WithField("tx", executedTx.Hash().Hex())

	logger.Info("PNO Module EthService: Waiting for transaction to be mined")

	receipt, err := e.waitTx(context, executedTx)
	if err != nil {
		return executedTx, fmt.Errorf("failed to wait for tx (%s) to be mined: %w", executedTx.Hash().Hex(), err)
	}

	if receipt.Status != types.ReceiptStatusSuccessful {
		return executedTx, fmt.Errorf("tx (%s) failed in execution", executedTx.Hash().Hex())
	}

	logger.Info("PNO Module EthService: Transaction executed successfully")

	return executedTx, nil

}
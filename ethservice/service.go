package ethservice

import (
	"context"
	"fmt"
	"math/big"

	types "github.com/ethereum/go-ethereum/core/types"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"github.com/restaking-cloud/pno-for-plus/ethservice/config"
	"github.com/restaking-cloud/pno-for-plus/ethservice/contracts"
)

type EthService struct {
	client *ethclient.Client
	cfg    config.EthServiceConfig

	log *logrus.Entry
}

func NewEthService() *EthService {
	return &EthService{}
}

func (e *EthService) Configure(cfg config.EthServiceConfig, logger *logrus.Entry) error {
	e.cfg = cfg
	e.log = logger

	err := e.connect(cfg.ExecutionNodeUrl)
	if err != nil {
		return err
	}

	err = e.configureMulticallContract(common.HexToAddress(contracts.MULTICALL3_CONTRACT_ADDRESS))
	if err != nil {
		return err
	}

	err = e.configureK2LendingContract(cfg.K2LendingContractAddress)
	if err != nil {
		return err
	}

	err = e.configureK2NodeOperatorContract(cfg.K2NodeOperatorContractAddress)
	if err != nil {
		return err
	}

	k2LendingProposerRegistryAddress, err := e.FetchProposerRegistryAddressFromK2Lending()
	if err != nil {
		return err
	}
	k2NodeOperatorProposerRegistryAddress, err := e.FetchProposerRegistryAddressFromK2NodeOperator()
	if err != nil {
		return err
	}

	// check that both k2 contracts are pointing to the same Proposer Registry contract
	if k2LendingProposerRegistryAddress != k2NodeOperatorProposerRegistryAddress {
		return fmt.Errorf("k2 Lending and NodeOperator contracts are not compatible, Proposer Registry addresses do not match")
	}

	pnoRegistryAddress, err := e.FetchPNORegistryAddressFromK2NodeOperator()
	if err != nil {
		return err
	}

	// check that the PNO Registry address matches what is configured if not override
	if pnoRegistryAddress != cfg.PNORegistryContractAddress.String() {
		e.log.WithFields(
			logrus.Fields{
				"configuredPNORegistryAddress":   cfg.PNORegistryContractAddress.String(),
				"nodeOperatorPNORegistryAddress": pnoRegistryAddress,
			},
		).Debug("Proposer Registry contract address mismatch, overriding with the one from the NodeOperator contract")
		// debug log and not warn, since it is expected behaviour to retrieve most up to date from the contracts and not set by user
		cfg.PNORegistryContractAddress = common.HexToAddress(pnoRegistryAddress)
	}

	// check there is a Proposer Registry contract address
	if (cfg.PNORegistryContractAddress == common.Address{}) {
		return fmt.Errorf("pno registry contract address not configured")
		// ideally this would not be possible, since bundled with the module,
		// but in case it happened to be overridden by a wrong address from the NodeOperator contract
	}

	err = e.configurePNORegistryContract(cfg.PNORegistryContractAddress)
	if err != nil {
		return err
	}

	return nil
}

func (e *EthService) ConnectedChainId() *big.Int {
	return e.cfg.ChainID
}

func (e *EthService) Status() (*ethereum.SyncProgress, error) {
	return e.client.SyncProgress(context.Background())
}

func (e *EthService) SetMaxGasPrice(maxGasPrice uint64) {

	e.cfg.MaxGasPrice = big.NewInt(int64(maxGasPrice))

	logger := logrus.WithField("moduleExecution", "pno").WithField("maxGasPrice", e.cfg.MaxGasPrice.String())
	currentGasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		logger.WithError(err).Debug("Failed to retrieve current gas price")
	} else {
		// check if max gas price is more than 30% lower than current gas price
		diff := new(big.Float).Sub(new(big.Float).SetInt(e.cfg.MaxGasPrice), new(big.Float).SetInt(currentGasPrice))
		percentage := new(big.Float).Quo(diff, new(big.Float).SetInt(currentGasPrice))
		if percentage.Cmp(big.NewFloat(-0.3)) < 0 {
			logger.WithFields(
				logrus.Fields{
					"currentGasPrice": currentGasPrice.String() + " gwei",
					"maxGasPrice":     e.cfg.MaxGasPrice.String() + " gwei",
				},
			).Warn("Max gas price is more than 30% lower than current gas price, consider increasing it, else registrations might be paused for a long time")
		}
	}
}

func (e *EthService) FetchProposerRegistryAddressFromK2Lending() (string, error) {

	data, err := e.cfg.K2LendingContractABI.Pack("proposerRegistry")
	if err != nil {
		return "", err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2LendingContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	var contractAddress common.Address
	err = e.cfg.K2LendingContractABI.UnpackIntoInterface(&contractAddress, "proposerRegistry", callResult)
	if err != nil {
		return "", err
	}

	return contractAddress.String(), nil
}

func (e *EthService) FetchProposerRegistryAddressFromK2NodeOperator() (string, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("proposerRegistry")
	if err != nil {
		return "", err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	var contractAddress common.Address
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&contractAddress, "proposerRegistry", callResult)
	if err != nil {
		return "", err
	}

	return contractAddress.String(), nil
}

func (e *EthService) FetchPNORegistryAddressFromK2NodeOperator() (string, error) {

	data, err := e.cfg.K2NodeOperatorContractABI.Pack("pnoRegistry")
	if err != nil {
		return "", err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2NodeOperatorContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return "", err
	}

	var contractAddress common.Address
	err = e.cfg.K2NodeOperatorContractABI.UnpackIntoInterface(&contractAddress, "pnoRegistry", callResult)
	if err != nil {
		return "", err
	}

	return contractAddress.String(), nil
}

func (e *EthService) NativeDelegationCountForNodeOperator(nodeOperatorAddress common.Address) (*big.Int, error) {

	data, err := e.cfg.K2LendingContractABI.Pack("nodeOperatorToBlsPublicKeyCount", nodeOperatorAddress)
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2LendingContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.K2LendingContractABI.UnpackIntoInterface(&callResultDecoded, "nodeOperatorToBlsPublicKeyCount", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking nodeOperatorToBlsPublicKeyCount result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) TotalReDelegationsInETH(userAddress common.Address) (*big.Int, error) {

	data, err := e.cfg.PNORegistryContractABI.Pack("totalReDelegationsInETH", userAddress)
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.PNORegistryContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.PNORegistryContractABI.UnpackIntoInterface(&callResultDecoded, "totalReDelegationsInETH", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking totalReDelegationsInETH result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) TotalReDelegationsInETHFromK2LP(userAddress common.Address) (*big.Int, error) {

	data, err := e.cfg.PNORegistryContractABI.Pack("totalReDelegationsInETHFromK2LP", userAddress)
	if err != nil {
		return nil, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.PNORegistryContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return nil, err
	}

	var callResultDecoded *big.Int
	err = e.cfg.PNORegistryContractABI.UnpackIntoInterface(&callResultDecoded, "totalReDelegationsInETHFromK2LP", callResult)
	if err != nil {
		return nil, fmt.Errorf("error unpacking totalReDelegationsInETHFromK2LP result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) ReDelegateToPreferredNodeOperatorFromNativeDelegation(nodeOperatorAddress common.Address, amount *big.Int) (tx *types.Transaction, err error) {

	data, err := e.cfg.PNORegistryContractABI.Pack("reDelegateToPreferredNodeOperatorFromNativeDelegation", nodeOperatorAddress, amount)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.PNORegistryContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending reDelegateToPreferredNodeOperatorFromNativeDelegation transaction: %w", err)
	}

	return executedTx, nil
}

func (e *EthService) IsPreferredNodeOperator(userAddress common.Address) (bool, error) {

	data, err := e.cfg.PNORegistryContractABI.Pack("isPreferredNodeOperator", userAddress)
	if err != nil {
		return false, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.PNORegistryContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return false, err
	}

	var callResultDecoded bool
	err = e.cfg.PNORegistryContractABI.UnpackIntoInterface(&callResultDecoded, "isPreferredNodeOperator", callResult)
	if err != nil {
		return false, fmt.Errorf("error unpacking isPreferredNodeOperator result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) DelegatedRecipient(userAddress common.Address) (common.Address, error) {

	data, err := e.cfg.K2LendingContractABI.Pack("delegatedClaim", userAddress)
	if err != nil {
		return common.Address{}, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.K2LendingContractAddress,
		Data: data,
	}, nil)
	if err != nil {
		return common.Address{}, err
	}

	var callResultDecoded common.Address
	err = e.cfg.K2LendingContractABI.UnpackIntoInterface(&callResultDecoded, "delegatedClaim", callResult)
	if err != nil {
		return common.Address{}, fmt.Errorf("error unpacking delegatedClaim result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) LienContractAddress() (common.Address, error) {

	data, err := e.cfg.PNORegistryContractABI.Pack("lienContract")
	if err != nil {
		return common.Address{}, err
	}

	callResult, err := e.client.CallContract(context.Background(), ethereum.CallMsg{
		From: e.cfg.ValidatorWalletAddress,
		To:   &e.cfg.PNORegistryContractAddress,
		Data: data,
	}, nil)

	if err != nil {
		return common.Address{}, err
	}

	var callResultDecoded common.Address
	err = e.cfg.PNORegistryContractABI.UnpackIntoInterface(&callResultDecoded, "lienContract", callResult)
	if err != nil {
		return common.Address{}, fmt.Errorf("error unpacking lienContract result: %w", err)
	}

	return callResultDecoded, nil
}

func (e *EthService) SetDelegatedRecipient(recipientAddress common.Address) (tx *types.Transaction, err error) {

	data, err := e.cfg.K2LendingContractABI.Pack("setDelegatedRecipient", recipientAddress)
	if err != nil {
		return nil, err
	}

	executedTx, err := e.transactAndWait(context.Background(), types.NewTx(&types.DynamicFeeTx{
		To:   &e.cfg.K2LendingContractAddress,
		Data: data,
	}), e.cfg.ValidatorWalletPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error sending setDelegatedRecipient transaction: %w", err)
	}

	return executedTx, nil
}
package pno

import (
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/sirupsen/logrus"
)

func (pno *PNOService) reDelegateAvailableNativeDelegationToSelf() (err error) {
	// Redelegate natively delegated validator ETH balance to self
	// as a self-preferencing, if native delegation balance
	// by the currently running node operator is unused in PNO registry

	// As a registered node operator in the Proposer Registry, with one or many
	// keys registered under the configured wallet, this makes the available ETH natively
	// delegated from these validators available to boost your own node operator
	// preferencing score/balance, and thus your own node operator's chances of being
	// as a node runner by other protocol users.

	pno.lock.Lock()
	defer pno.lock.Unlock()

	// Only process a re-delegation attempt every epoch
	currenttime := time.Now()
	if currenttime.Sub(pno.lastAutoRedelegatedTimeStamp) < (32 * 12 * time.Second) {
		return nil
	}
	// Set the last attempted auto-redelegation timestamp
	pno.lastAutoRedelegatedTimeStamp = time.Now()

	// Check if the node operator is registered in the PNO registry
	isRegistered, err := pno.eth.IsPreferredNodeOperator(pno.cfg.ValidatorWalletAddress)
	if err != nil {
		return fmt.Errorf("failed to check if node operator is registered in PNO registry: %v", err)
	}
	if !isRegistered {
		return nil
	}

	// Ensure that the user can redelegate
	canRedelegate, err := pno.ensureCanRedelegate()
	if err != nil {
		return fmt.Errorf("failed to ensure can redelegate: %v", err)
	}
	if !canRedelegate {
		return nil
	}

	// Check amount of native delegations for the configured wallet
	nativeDelegationCount, err := pno.eth.NativeDelegationCountForNodeOperator(pno.cfg.ValidatorWalletAddress)
	if err != nil {
		return fmt.Errorf("failed to get native delegation count for node operator: %v", err)
	}

	// If no native delegations, return
	if nativeDelegationCount.Sign() == 0 {
		return nil
	}

	balancePerDelegation, success := big.NewInt(0).SetString("24000000000000000000", 10)
	if !success {
		return fmt.Errorf("failed to set balance per delegation")
	}
	nativelyDelegatedBalance := big.NewInt(0).Mul(nativeDelegationCount, balancePerDelegation)
	// Check how much has been redelegated by the node operator from the native delegation balance
	redelegatedTotal, err := pno.eth.TotalReDelegationsInETH(pno.cfg.ValidatorWalletAddress)
	if err != nil {
		return fmt.Errorf("failed to get total redelegations: %v", err)
	}
	redelegatedFromK2LP := big.NewInt(0)

	if redelegatedTotal.Sign() > 0 {
		// Check how much has been redelegated by the node operator from the native delegation balance to K2LP
		redelegatedFromK2LP, err = pno.eth.TotalReDelegationsInETHFromK2LP(pno.cfg.ValidatorWalletAddress)
		if err != nil {
			return fmt.Errorf("failed to get total redelegations from K2LP: %v", err)
		}
	}

	redelegatedFromNativeDelegation := big.NewInt(0).Sub(redelegatedTotal, redelegatedFromK2LP)

	availableToRedelegateFromNativeDelegation := big.NewInt(0).Sub(nativelyDelegatedBalance, redelegatedFromNativeDelegation)

	// If there is available balance, redelegate to self
	if availableToRedelegateFromNativeDelegation.Sign() > 0 {
		// Redelegate to self

		tx, err := pno.eth.ReDelegateToPreferredNodeOperatorFromNativeDelegation(pno.cfg.ValidatorWalletAddress, availableToRedelegateFromNativeDelegation)
		if err != nil {
			pno.log.WithError(err).Error("Failed to redelegate to self from native delegation")
			return fmt.Errorf("failed to redelegate to self from native delegation: %w", err)
		}
		pno.log.WithFields(logrus.Fields{
			"txHash":                               tx.Hash().String(),
			"redelegatedAmount":                    availableToRedelegateFromNativeDelegation.String(),
			"totalReDelegated":                     big.NewInt(0).Add(redelegatedTotal, availableToRedelegateFromNativeDelegation).String(),
			"totalReDelegatedFromK2LP":             redelegatedFromK2LP.String(),
			"totalReDelegatedFromNativeDelegation": big.NewInt(0).Add(redelegatedFromNativeDelegation, availableToRedelegateFromNativeDelegation).String(),
		}).Info("Re-delegated to self from native delegation")

	}

	return nil

}

func (pno *PNOService) ensureCanRedelegate() (bool, error) {
	// Ensure that the user can redelegate
	// to any node operator in the PNO registry
	// by checking if they have set the lien contract
	// as the delegated recipient

	var delegatedRecipient common.Address = common.Address{}
	var lienContract common.Address = common.Address{}
	var err error
	var lock sync.Mutex

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		delegatedRecipient, err = pno.eth.DelegatedRecipient(pno.cfg.ValidatorWalletAddress)
		if err != nil {
			lock.Lock()
			err = fmt.Errorf("failed to get delegated recipient: %v", err)
			lock.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		lienContract, err = pno.eth.LienContractAddress()
		if err != nil {
			lock.Lock()
			err = fmt.Errorf("failed to get lien contract address: %v", err)
			lock.Unlock()
		}
	}()

	pno.log.Debug("Ensuring can redelegate")

	wg.Wait()

	if err != nil {
		return false, err
	}

	if lienContract == (common.Address{}) {
		return false, fmt.Errorf("lien contract address not available to check delegated recipient")
	}

	if delegatedRecipient == lienContract {
		return true, nil
	}

	pno.log.WithFields(logrus.Fields{
		"delegatedRecipient": delegatedRecipient.String(),
		"lienContract":       lienContract.String(),
	}).Info("Delegated recipient not set to lien contract")

	pno.log.Info("Setting delegated recipient to lien contract")

	// Set the lien contract as the delegated recipient
	tx, err := pno.eth.SetDelegatedRecipient(lienContract)
	if err != nil {
		return false, fmt.Errorf("failed to set delegated recipient: %v", err)
	}

	pno.log.WithFields(logrus.Fields{
		"txHash":       tx.Hash().String(),
		"userAddress":  pno.cfg.ValidatorWalletAddress.String(),
		"lienContract": lienContract.String(),
	}).Info("Set delegated recipient to lien contract")

	return true, nil
}

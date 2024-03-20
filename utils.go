package pno

import (
	"crypto/ecdsa"
	"fmt"
	"strconv"
	"strings"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pon-network/mev-plus/common"
	pnocommon "github.com/restaking-cloud/pno-for-plus/common"
	"github.com/restaking-cloud/pno-for-plus/config"
)

func (pno *PNOService) parseConfig(moduleFlags common.ModuleFlags) (err error) {
	for flagName, flagValue := range moduleFlags {
		switch flagName {
		case config.WalletPrivateKeyFlag.Name:
			pk, err := crypto.HexToECDSA(strings.TrimPrefix(flagValue, "0x"))
			if err != nil {
				return fmt.Errorf("-%s: invalid wallet private private key %q", config.WalletPrivateKeyFlag.Name, flagValue)
			}
			pno.cfg.ValidatorWalletPrivateKey = pk

			publicKey := pk.Public()

			walletAddress := crypto.PubkeyToAddress(*publicKey.(*ecdsa.PublicKey))

			pno.cfg.ValidatorWalletAddress = walletAddress
		case config.BeaconNodeUrlFlag.Name:
			pno.cfg.BeaconNodeUrl, err = pnocommon.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.BeaconNodeUrlFlag.Name, flagValue)
			}
		case config.ExecutionNodeUrlFlag.Name:
			pno.cfg.ExecutionNodeUrl, err = pnocommon.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.ExecutionNodeUrlFlag.Name, flagValue)
			}
		case config.MaxGasPriceFlag.Name:
			setMaxGasPrice, err := strconv.ParseUint(flagValue, 10, 64)
			if err != nil {
				return fmt.Errorf("-%s: invalid max gas price %q", config.MaxGasPriceFlag.Name, flagValue)
			}
			if setMaxGasPrice <= 0 {
				return fmt.Errorf("-%s: max gas price must be greater than zero", config.MaxGasPriceFlag.Name)
			}
			pno.cfg.MaxGasPrice = setMaxGasPrice
		case config.ListenAddressFlag.Name:
			pno.cfg.ListenAddress, err = pnocommon.CreateUrl(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid url %q", config.ListenAddressFlag.Name, flagValue)
			}
		case config.AutoReDelegateFlag.Name:
			pno.cfg.AutoReDelegateFlag, err = strconv.ParseBool(flagValue)
			if err != nil {
				return fmt.Errorf("-%s: invalid auto redelegate flag %q", config.AutoReDelegateFlag.Name, flagValue)
			}
		case config.K2LendingContractAddressFlag.Name:
			pno.cfg.K2LendingContractAddress = ethCommon.HexToAddress(flagValue)
			if pno.cfg.K2LendingContractAddress == (ethCommon.Address{}) {
				return fmt.Errorf("-%s: invalid address %q", config.K2LendingContractAddressFlag.Name, flagValue)
			}
		case config.K2NodeOperatorContractAddressFlag.Name:
			pno.cfg.K2NodeOperatorContractAddress = ethCommon.HexToAddress(flagValue)
			if pno.cfg.K2NodeOperatorContractAddress == (ethCommon.Address{}) {
				return fmt.Errorf("-%s: invalid address %q", config.K2NodeOperatorContractAddressFlag.Name, flagValue)
			}
		default:
			return fmt.Errorf("unknown flag %q", flagName)
		}

	}

	if len(moduleFlags) > 0 {
		pno.lock.Lock()
		pno.configured = true
		pno.lock.Unlock()
	}

	err = pno.checkConfig()
	if err != nil {
		return err
	}

	return nil
}

func (pno *PNOService) checkConfig() error {

	pno.lock.Lock()
	defer pno.lock.Unlock()

	if !pno.configured {
		// If not set to run, return
		return nil
	}

	pno.configured = false

	// check that the execution node url is set
	if pno.cfg.ExecutionNodeUrl == nil {
		return fmt.Errorf("-%s: execution node url is required", config.ExecutionNodeUrlFlag.Name)
	}

	// check that the beacon node url is set
	if pno.cfg.BeaconNodeUrl == nil {
		return fmt.Errorf("-%s: beacon node url is required", config.BeaconNodeUrlFlag.Name)
	}

	// check that the wallet private key is set
	if pno.cfg.ValidatorWalletPrivateKey == nil {
		return fmt.Errorf("-%s: validator wallet private key is required", config.WalletPrivateKeyFlag.Name)
	}

	pno.configured = true

	return nil
}

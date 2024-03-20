package config

import (
	"crypto/ecdsa"
	"math/big"
	"net/url"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type EthServiceConfig struct {
	ExecutionNodeUrl *url.URL
	ChainID          *big.Int

	MaxGasPrice *big.Int

	K2LendingContractAddress        common.Address
	K2NodeOperatorContractAddress   common.Address
	PNORegistryContractAddress      common.Address

	// ABI
	K2LendingContractABI            *abi.ABI
	K2NodeOperatorContractABI        *abi.ABI
	PNORegistryContractABI          *abi.ABI

	// Multicall
	MulticallContractAddress common.Address
	MulticallContractABI     *abi.ABI

	ValidatorWalletPrivateKey *ecdsa.PrivateKey
	ValidatorWalletAddress    common.Address
}
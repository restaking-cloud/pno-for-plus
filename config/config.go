package config

import (
	"crypto/ecdsa"
	"net/url"

	"github.com/ethereum/go-ethereum/common"
)

type PNOConfig struct {
	ValidatorWalletPrivateKey     *ecdsa.PrivateKey
	ValidatorWalletAddress        common.Address
	ExecutionNodeUrl              *url.URL
	BeaconNodeUrl                 *url.URL
	K2LendingContractAddress	  common.Address
	K2NodeOperatorContractAddress common.Address
	PNORegistryContractAddress    common.Address
	ListenAddress                 *url.URL
	MaxGasPrice                   uint64
	AutoReDelegateFlag			bool
}

var PNOConfigDefaults = PNOConfig{
	ValidatorWalletPrivateKey:     nil,
	ValidatorWalletAddress:        common.Address{},
	ExecutionNodeUrl:              nil,
	BeaconNodeUrl:                 nil,
	K2LendingContractAddress:	  common.Address{},
	K2NodeOperatorContractAddress: common.Address{},
	PNORegistryContractAddress:    common.Address{},
	ListenAddress:                 &url.URL{Scheme: "http", Host: "localhost:9000"},
	MaxGasPrice:                   0,
	AutoReDelegateFlag:            false,
}
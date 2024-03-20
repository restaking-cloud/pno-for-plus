package config

import (
	"github.com/ethereum/go-ethereum/common"
)

// A mapping of chain IDs to their respective PNO known configurations
var PNOConfigConstants = map[uint64]PNOConfig{
	5: {
		K2LendingContractAddress:             common.HexToAddress("0x5f3b5DfA2fBf3fDfEe9fA3d3c6e7fCf3f3f3f3f3f"),
		K2NodeOperatorContractAddress:        common.HexToAddress("0x10b37A1A3e3114fe479B2cf962dB8806c941d2Dc"),
		PNORegistryContractAddress:           common.HexToAddress("0x5f3b5DfA2fBf3fDfEe9fA3d3c6e7fCf3f3f3f3f3f"),
	},
	17000: {
		K2LendingContractAddress:             common.HexToAddress("0x4655512B176243Dd161e61a818899324AE4E9323"),
		K2NodeOperatorContractAddress:        common.HexToAddress("0xe7C28eb37802c4015e65a8c55e182A9d5421Cac3"),
		PNORegistryContractAddress:           common.HexToAddress("0xfFa542ef408440A34809cBb754Aecb28dB1F5c5c"),
	},
}
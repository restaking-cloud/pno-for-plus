package config

import (
	"strings"

	cli "github.com/urfave/cli/v2"
)

var (
	WalletPrivateKeyFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "eth1-private-key",
		Usage:    "The private key of the validator wallet",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
		EnvVars:  []string{"ETH1_PRIVATE_KEY"},
	}
	ExecutionNodeUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "execution-node-url",
		Usage:    "The url of the execution node",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	BeaconNodeUrlFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "beacon-node-url",
		Usage:    "The url of the beacon node",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	ListenAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "listen-address",
		Usage:    "The address to listen on for incoming requests",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	MaxGasPriceFlag = &cli.Uint64Flag{
		Name:     ModuleName + "." + "max-gas-price",
		Usage:    "The maximum gas price to use for transactions, in Gwei",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
		Value:    10,
	}
	AutoReDelegateFlag = &cli.BoolFlag{
		Name:     ModuleName + "." + "auto-redelegate",
		Usage:    "Whether to automatically redelegate available netively delegated balance to self",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	K2LendingContractAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "k2-lending-contract-address",
		Usage:    "The address of the K2 lending contract to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
	K2NodeOperatorContractAddressFlag = &cli.StringFlag{
		Name:     ModuleName + "." + "k2-node-operator-contract-address",
		Usage:    "The address of the K2 node operator contract to override the internal configuration",
		Category: strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
	}
)
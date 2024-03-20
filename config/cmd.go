package config

import (
	"strings"

	cli "github.com/urfave/cli/v2"
)

const ModuleName = "pno"

func NewCommand() *cli.Command {
	return &cli.Command{
		Name:      ModuleName,
		Usage:     "Start the K2 preferred node operator module",
		UsageText: "", // TODO: add usage text
		Category:  strings.ReplaceAll(strings.ToUpper(ModuleName), "_", " "),
		Flags:     pnoFlags(),
	}
}

func pnoFlags() []cli.Flag {
	return []cli.Flag{
		WalletPrivateKeyFlag,
		ListenAddressFlag,
		ExecutionNodeUrlFlag,
		BeaconNodeUrlFlag,
		MaxGasPriceFlag,
		AutoReDelegateFlag,
		K2LendingContractAddressFlag,
		K2NodeOperatorContractAddressFlag,
	}
}
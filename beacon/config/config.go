package config

import (
	"net/url"
	"math/big"
)

type BeaconConfig struct {
	BeaconNodeUrl *url.URL
	ChainID *big.Int
}
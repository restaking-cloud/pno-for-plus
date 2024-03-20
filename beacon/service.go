package beacon

import (
	"math/big"
	"net/http"
	"time"
	"sync"

	"github.com/restaking-cloud/pno-for-plus/beacon/config"
)

type BeaconService struct {
	cfg    config.BeaconConfig
	client *http.Client

	mu sync.Mutex

	currentSlot uint64
}

func NewBeaconService() *BeaconService {
	return &BeaconService{
		client: &http.Client{Timeout: 6 * time.Second}, // get a response in half a slot
	}
}

func (b *BeaconService) Configure(cfg config.BeaconConfig) error {
	b.cfg = cfg

	err := b.connect(cfg.BeaconNodeUrl)
	if err != nil {
		return err
	}

	return nil
}

func (b *BeaconService) ConnectedChainId() *big.Int {
	return b.cfg.ChainID
}
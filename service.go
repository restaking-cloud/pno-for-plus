package pno

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/pon-network/mev-plus/common"
	coreCommon "github.com/pon-network/mev-plus/core/common"
	"github.com/restaking-cloud/pno-for-plus/config"

	apiv1 "github.com/attestantio/go-builder-client/api/v1"

	beaconConfig "github.com/restaking-cloud/pno-for-plus/beacon/config"
	ethConfig "github.com/restaking-cloud/pno-for-plus/ethservice/config"

	"github.com/restaking-cloud/pno-for-plus/beacon"
	"github.com/restaking-cloud/pno-for-plus/ethservice"

	"github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
)

type PNOService struct {
	coreClient *coreCommon.Client
	log        *logrus.Entry
	lock       sync.Mutex

	eth    *ethservice.EthService
	beacon *beacon.BeaconService

	server *http.Server

	exit chan struct{}

	configured bool
	cfg        config.PNOConfig

	lastAutoRedelegatedTimeStamp time.Time
}

func NewPNOService() *PNOService {
	return &PNOService{
		log:    logrus.NewEntry(logrus.New()).WithField("moduleExecution", config.ModuleName),
		eth:    ethservice.NewEthService(),
		beacon: beacon.NewBeaconService(),
		exit:   make(chan struct{}),
		cfg:    config.PNOConfigDefaults,
	}
}

func (pno *PNOService) CliCommand() *cli.Command {
	return config.NewCommand()
}

func (pno *PNOService) Name() string {
	return config.ModuleName
}

func (pno *PNOService) Start() error {

	err := pno.checkConfig()
	// if module configuration has been called and completed without error, this should pose no error
	if err != nil {
		return err
	}

	if !pno.configured {
		// module not configured to run
		return nil
	}

	if pno.server != nil {
		return fmt.Errorf("PNO server already running")
	}

	// start the server
	pno.server = &http.Server{
		Addr:    pno.cfg.ListenAddress.Host,
		Handler: pno.getRouter(),
	}

	go pno.startServer()

	// start monitoring and running background tasks
	go pno.startBackgroundTasks()

	pno.log.WithFields(logrus.Fields{
		"nodeOperatorWalletAddress": pno.cfg.ValidatorWalletAddress,
	}).Info("Started PNO module")

	return nil
}

func (pno *PNOService) startBackgroundTasks() error {

	ctx := context.Background()
	ctxWithCancel, cancel := context.WithCancel(ctx)
	var HeadChan chan beacon.HeadEventData = make(chan beacon.HeadEventData, 64)

	go pno.beacon.SubscribeToHeadEvents(ctxWithCancel, HeadChan)

	var lastEpoch uint64 = 0

	for {
		select {
		case <-pno.exit:
			cancel()
			return nil
		case <-ctxWithCancel.Done():
			cancel()
			return nil
		case headData := <-HeadChan:
			slot := headData.Slot

			epoch := slot / 32
			if epoch != lastEpoch {
				lastEpoch = epoch

				// Re-delegate available balance to self every epoch if there is any
				if pno.cfg.AutoReDelegateFlag {
					err := pno.reDelegateAvailableNativeDelegationToSelf()
					if err != nil {
						pno.log.WithError(err).Error("Failed to re-delegate available balance to self")
					}
				}
			}
		}
	}

}

func (pno *PNOService) Stop() error {

	// stop the server
	err := pno.stopServer()
	if err != nil {
		return err
	}

	close(pno.exit)

	pno.log.Info("Stopped PNO module")

	return nil
}

func (pno *PNOService) ConnectCore(coreClient *coreCommon.Client, pingId string) error {

	// this is the first and only time the client is set and doesnt need a mutex
	pno.coreClient = coreClient

	// test a ping to the core server
	err := pno.coreClient.Ping(pingId)
	if err != nil {
		return err
	}

	return nil
}

func (pno *PNOService) Configure(moduleFlags common.ModuleFlags) (err error) {

	err = pno.parseConfig(moduleFlags)
	if err != nil {
		return err
	}

	// connect to the beacon node and get the chain id configured
	err = pno.beacon.Configure(beaconConfig.BeaconConfig{
		BeaconNodeUrl: pno.cfg.BeaconNodeUrl,
	})
	if err != nil {
		return err
	}

	// retrieve the chain id from the beacon node
	chainId := pno.beacon.ConnectedChainId().Uint64()

	// check if chain id is supported
	knownConfig, ok := config.PNOConfigConstants[chainId]
	if !ok {
		// check if the module has provided contract addresses by the user
		if pno.cfg.K2LendingContractAddress != (ethCommon.Address{}) || pno.cfg.K2NodeOperatorContractAddress != (ethCommon.Address{}) {
			// Check for provided K2LendingContractAddress and K2NodeOperatorContractAddress
			// require the need for both contract addresses to be provided
			if pno.cfg.K2LendingContractAddress == (ethCommon.Address{}) || pno.cfg.K2NodeOperatorContractAddress == (ethCommon.Address{}) {
				return fmt.Errorf("provide both K2LendingContractAddress and K2NodeOperatorContractAddress for chain id %v", chainId)
			}
			pno.log.Debug("User provided K2LendingContractAddress and K2NodeOperatorContractAddress")
		} else {
			return fmt.Errorf("chain id %v is not supported", chainId)
		}
	} else {

		// beacon node chain id is supported, set the rest of the config
		if pno.cfg.K2LendingContractAddress == (ethCommon.Address{}) {
			pno.cfg.K2LendingContractAddress = knownConfig.K2LendingContractAddress
		} else {
			pno.log.Debug("User provided K2LendingContractAddress")
		}

		if pno.cfg.K2NodeOperatorContractAddress == (ethCommon.Address{}) {
			pno.cfg.K2NodeOperatorContractAddress = knownConfig.K2NodeOperatorContractAddress
		} else {
			pno.log.Debug("User provided K2NodeOperatorContractAddress")
		}

		pno.cfg.PNORegistryContractAddress = knownConfig.PNORegistryContractAddress
		// If user provided an addresses for nodeOperator different from known config,
		// the pno registry contract would be overridden from the nodeOperator contract
		// if different from the known config
	}

	// connect to the execution node and get the chain id, and contracts configured
	err = pno.eth.Configure(ethConfig.EthServiceConfig{
		ExecutionNodeUrl:              pno.cfg.ExecutionNodeUrl,
		K2LendingContractAddress:      pno.cfg.K2LendingContractAddress,
		K2NodeOperatorContractAddress: pno.cfg.K2NodeOperatorContractAddress,
		PNORegistryContractAddress:    pno.cfg.PNORegistryContractAddress,
		ValidatorWalletPrivateKey:     pno.cfg.ValidatorWalletPrivateKey,
		ValidatorWalletAddress:        pno.cfg.ValidatorWalletAddress,
	}, pno.log)
	if err != nil {
		return err
	}

	// Ensure that the chain id reported by the beacon node matches the chain id reported by the execution node
	ethChainId := pno.eth.ConnectedChainId().Uint64()
	if ethChainId != chainId {
		// wrong chain id configured for the execution node, needs to match the beacon node (validator truth source)
		return fmt.Errorf("chain id mismatch: beacon node reports %v, execution node reports %v", chainId, ethChainId)
	}

	if pno.cfg.MaxGasPrice > 0 {
		// Then the user has set a max gas price, set it on the eth1 service
		pno.eth.SetMaxGasPrice(pno.cfg.MaxGasPrice)
	}

	return nil
}

func (pno *PNOService) Status() error {

	// check execution node is up
	_, err := pno.eth.Status()
	if err != nil {
		return fmt.Errorf("execution node is down: %v", err)
	}

	return nil

}

func (pno *PNOService) RegisterValidator(_ []apiv1.SignedValidatorRegistration) error {

	if !pno.configured {
		// module not configured to run
		return nil
	}

	if pno.cfg.AutoReDelegateFlag {
		// auto re-delegation to self is enabled
		pno.log.Debug("Triggered Re-Delegation to PNO from RegisterValidator call")
		err := pno.reDelegateAvailableNativeDelegationToSelf()
		if err != nil {
			pno.log.WithError(err).Error("Failed to re-delegate available balance to self")
			return err
		}
	}

	return nil
}

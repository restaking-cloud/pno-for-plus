package beacon

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
)

func (b *BeaconService) connect(url *url.URL) error {
	b.cfg.BeaconNodeUrl = url

	ctx := context.Background()

	id, err := b.networkID(ctx)
	if err != nil {
		return err
	}
	b.cfg.ChainID = id

	synced, err := b.syncProgress(ctx)
	if err != nil {
		return err
	}

	if synced == nil || synced.IsSyncing {
		return fmt.Errorf("beacon node not synced")
	}

	if synced.ElOffline {
		return fmt.Errorf("beacon node not synced, execution layer offline")
	}

	return nil
}

func (b *BeaconService) networkID(ctx context.Context) (chainId *big.Int, err error) {
	spec, err := b.getSpec(ctx)
	if err != nil {
		return chainId, err
	}

	id, ok := spec["DEPOSIT_CHAIN_ID"].(string)
	if !ok {
		return chainId, fmt.Errorf("invalid chain id")
	}

	chainId, ok = new(big.Int).SetString(id, 10)
	if !ok {
		return chainId, fmt.Errorf("invalid chain id")
	}

	return chainId, nil
}

func (b *BeaconService) getSpec(ctx context.Context) (res map[string]interface{}, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", b.cfg.BeaconNodeUrl.String()+SpecPath, nil)
	if err != nil {
		return res, err
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("invalid response (%d): %v", resp.StatusCode, resp)
	}

	var response struct {
		Data map[string]interface{} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return res, err
	}

	return response.Data, nil
}

func (b *BeaconService) syncProgress(ctx context.Context) (res *SyncStatusData, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", b.cfg.BeaconNodeUrl.String()+SyncPath, nil)
	if err != nil {
		return res, err
	}

	resp, err := b.client.Do(req)
	if err != nil {
		return res, err
	}

	if resp.StatusCode != 200 {
		return res, fmt.Errorf("invalid response (%d): %v", resp.StatusCode, resp)
	}

	var response GetSyncStatusResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return res, err
	}

	return response.Data, nil
}
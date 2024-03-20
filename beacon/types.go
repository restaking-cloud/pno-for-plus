package beacon

type GetSyncStatusResponse struct {
	Data *SyncStatusData `json:"data"`
}

type SyncStatusData struct {
	HeadSlot uint64 `json:"head_slot,string"`
	SyncDistance uint64 `json:"sync_distance,string"`
	IsSyncing bool `json:"is_syncing"`
	IsOptimistic bool `json:"is_optimistic"`
	ElOffline bool `json:"el_offline"`
}

type HeadEventData struct {
	Slot  uint64 `json:"slot,string"`
	Block string `json:"block"`
	State string `json:"state"`
}
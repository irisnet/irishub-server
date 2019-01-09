package services

import "github.com/irisnet/irishub-server/models/document"

type SyncService struct {

}

func (syncService SyncService) GetCurrentSyncResult() document.SyncResult {
	r, _ := syncResult.GetCurrentSyncResult()
	return r
}
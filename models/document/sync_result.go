package document

import (
	"github.com/irisnet/irishub-server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollectionNameSyncConf = "sync_task"

type SyncResult struct {
	CurrentBlockHeight int64 `json:"current_height" bson:"current_height"`
}

func (syncResult SyncResult) Name() string {
	return CollectionNameSyncConf
}

func (syncResult SyncResult) Query(query bson.M) (result SyncResult, err error) {
	exOp := func(c *mgo.Collection) error {
		return c.Find(query).One(&result)
	}
	return result, models.ExecCollection(syncResult.Name(), exOp)
}

func (syncResult SyncResult) GetCurrentSyncResult() (result SyncResult, err error) {
	query := bson.M{
		"end_height": 0,
		"status":     "underway",
	}
	return syncResult.Query(query)
}

package document

import (
	"github.com/irisnet/iris-api-server/utils/constants"
	"github.com/irisnet/iris-api-server/utils/helper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	
	"github.com/irisnet/iris-api-server/models"
)

const (
	CollectionNmStakeTx = "tx_stake"
)

// StakeTx
type StakeTx struct {
	TxHash string    `json:"tx_hash" bson:"tx_hash"`
	Time   time.Time `json:"time" bson:"time"`
	Height int64     `json:"height" bson:"height"`
	From   string    `json:"from" bson:"from"`
	PubKey string    `json:"pub_key" bson:"pub_key"`
	Type   string    `json:"type" bson:"type"`
	Amount Coin `json:"amount" bson:"amount"`
}

func (m StakeTx) Name() string {
	return CollectionNmStakeTx
}

func (m StakeTx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": m.TxHash}
}

func (m StakeTx) Query(
	query bson.M, fields bson.M,
	skip int, limit int, sorts ...string) (results []StakeTx, err error) {
	
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(m.Name(), exop)
}

// get list of stake tx
func (m StakeTx) GetStakeTxList(
	address string, pubKey string, txType string,
	startTime time.Time, endTime time.Time,
	skip int, limit int, sorts []string) (
		[]StakeTx, error,
) {
	
	query := bson.M{
	}
	
	if address != "" {
		query["from"] = address
	}
	if pubKey != "" {
		query["pub_key"] = pubKey
	}
	if txType != "" {
		query["type"] = txType
	}
	
	if startTime.IsZero() {
		startTime, _ = helper.ParseTime(constants.TIME_LAYOUT_FULL, constants.TIME_START)
	}
	if endTime.IsZero() {
		endTime = time.Now()
	}
	
	query["time"] = bson.M{
		"$gte": startTime,
		"$lte": endTime,
	}
	fields := bson.M{}
	
	return m.Query(query, fields, skip, limit, sorts...)

}

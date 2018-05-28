package document

import (
	"time"
	
	"github.com/irisnet/iris-api-server/models"
	"github.com/irisnet/iris-api-server/modules/logger"
	"github.com/irisnet/iris-api-server/utils/constants"
	"github.com/irisnet/iris-api-server/utils/helper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmCommonTx = "tx_common"
)

type CommonTx struct {
	TxHash string     `json:"tx_hash" bson:"tx_hash"`
	Time   time.Time  `json:"time" bson:"time"`
	Height int64      `json:"height" bson:"height"`
	From   string     `json:"from" bson:"from"`
	To     string     `json:"to" bson:"to"`
	Amount Coins      `json:"amount" bson:"amount"`
	Type   string     `json:"type" bson:"type"`
	
	Candidate Candidate `json:"candidate"`
}

func (d CommonTx) Name() string {
	return CollectionNmCommonTx
}

func (d CommonTx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": d.TxHash}
}

func (d CommonTx) Query(
	query bson.M, fields bson.M,
	skip int, limit int, sorts ...string) (results []CommonTx, err error) {
	
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}

func (d CommonTx) GetList(address string, txType string,
	startTime time.Time, endTime time.Time,
	skip int, limit int, sorts []string) (
	[]CommonTx, error) {
	
	query := bson.M{
	}
	
	if txType == "" {
		query = bson.M{
			"$or": []bson.M{
				bson.M{"from": address},
				bson.M{"to": address},
			},
		}
	} else {
		switch txType {
		case constants.TxTypeCoinReceive:
			query["to"] = address
			break
		case constants.TxTypeCoinSend, constants.TxTypeStakeDelegate, constants.TxTypeStakeUnBond:
			query["from"] = address
			break
		}
		query["type"] = constants.TxTypeFrontMapDb[txType]
	}
	
	if startTime.IsZero() {
		startTime, _ = helper.ParseFullTime(constants.TIME_START)
	}
	if endTime.IsZero() {
		endTime = time.Now()
	}
	query["time"] = bson.M{
		"$gte": startTime,
		"$lte": endTime,
	}
	fields := bson.M{}
	logger.Info.Println(helper.ToJson(query))
	
	return d.Query(query, fields, skip, limit, sorts...)
}

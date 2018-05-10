package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNmStakeTx = "tx_stake"
)

// StakeTx
type StakeTx struct {
	TxHash string    `bson:"tx_hash"`
	Time   time.Time `bson:"time"`
	Height int64     `bson:"height"`
	From   string    `bson:"from"`
	PubKey string    `bson:"pub_key"`
	Type   string    `bson:"type"`
	Amount Coin `bson:"amount"`
}

func (c StakeTx) Name() string {
	return CollectionNmStakeTx
}

func (c StakeTx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": c.TxHash}
}

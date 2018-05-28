package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNmCoinTx = "tx_coin"
)

type CoinTx struct {
	TxHash string     `bson:"tx_hash"`
	Time   time.Time  `bson:"time"`
	Height int64      `bson:"height"`
	From   string     `bson:"from"`
	To     string     `bson:"to"`
	Amount Coins `bson:"amount"`
}

func (c CoinTx) Name() string {
	return CollectionNmCoinTx
}

func (c CoinTx) PkKvPair() map[string]interface{} {
	return bson.M{"tx_hash": c.TxHash}
}


package document

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	CollectionNmAccount = "account"
)

type Account struct {
	Address string     `bson:"address"`
	Amount  Coins `bson:"amount"`
	Time    time.Time  `bson:"time"`
	Height  int64      `bson:"height"`
}

func (a Account) Name() string {
	return CollectionNmAccount
}

func (a Account) PkKvPair() map[string]interface{} {
	return bson.M{"address": a.Address}
}

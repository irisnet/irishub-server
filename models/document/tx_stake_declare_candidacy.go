package document

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeTxDeclareCandidacy = "tx_stake"
)

// Description
type Description struct {
	Moniker  string `json:"moniker" bson:"moniker"`
	Identity string `json:"identity" bson:"identity"`
	Website  string `json:"website" bson:"website"`
	Details  string `json:"details" bson:"details"`
}

type StakeTxDeclareCandidacy struct {
	StakeTx `bson:"stake_tx"`
	Description `bson:"description"`
}

func (s StakeTxDeclareCandidacy) Name() string  {
	return CollectionNmStakeTxDeclareCandidacy
}

func (s StakeTxDeclareCandidacy) PkKvPair() map[string]interface{}  {
	return bson.M{"stake_tx.tx_hash": s.TxHash}
}
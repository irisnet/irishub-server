package document

import (
	"github.com/irisnet/irishub-server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const CollectionNmValidatorHistory = "validator_history"

type ValidatorHistory struct {
	Candidate  `bson:"candidate"`
	UpdateTime time.Time `bson:"update_time"`
}

func (v ValidatorHistory) Name() string {
	return CollectionNmValidatorHistory
}

func (v ValidatorHistory) PkKvPair() map[string]interface{} {
	return bson.M{"candidate.address": v.Address}
}

func (v ValidatorHistory) QueryAll() (vs []ValidatorHistory) {
	queryOp := func(c *mgo.Collection) error {
		return c.Find(nil).All(&vs)
	}
	models.ExecCollection(v.Name(), queryOp)
	return vs
}

package document

import (
	"github.com/irisnet/irishub-server/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollectionName = "validator_up_time"

type ValidatorUpTime struct {
	ValAddress string  `bson:"val_address"`
	UpTime     float64 `bson:"up_time"`
}

func (d ValidatorUpTime) Name() string {
	return CollectionName
}

func (d ValidatorUpTime) PkKvPair() map[string]interface{} {
	return bson.M{"val_address": d.ValAddress}
}

func (d ValidatorUpTime) Query(query bson.M, sorts []string, skip, limit int) (
	results []ValidatorUpTime, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}

func (d ValidatorUpTime) GetUpTime(valAddress []string) ([]ValidatorUpTime, error) {
	q := bson.M{
		"val_address": bson.M{
			"$in": valAddress,
		},
	}
	var sorts []string
	skip := 0
	limit := len(valAddress)

	return d.Query(q, sorts, skip, limit)
}

package document

import (
	"github.com/irisnet/iris-api-server/modules/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/irisnet/iris-api-server/models"
)

const (
	CollectionNmStakeRoleDelegator = "stake_role_delegator"
)

type Delegator struct {
	Address string `bson:"address"`
	PubKey  string `bson:"pub_key"`
	Shares  int64  `bson:"shares"`
}

func (d Delegator) Name() string {
	return CollectionNmStakeRoleDelegator
}

func (d Delegator) PkKvPair() map[string]interface{} {
	return bson.M{"address": d.Address, "pub_key": d.PubKey}
}

func (d Delegator) Query(
	query bson.M, skip int, limit int, sorts ...string,
	) (results []Delegator, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sorts...).
			Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}

func (d Delegator) GetDelegatorListByAddressAndPubKeys(address string, pubKeys []string,
	) ([]Delegator, error) {

	query := bson.M{
		"address": address,
		"pub_key": &bson.M{
			"$in": pubKeys,
		},
	}
	sorts := make([]string, 0)

	delegator, err := d.Query(query, 0, len(pubKeys), sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return delegator, err
}

package document

import (
	"time"

	"github.com/irisnet/irishub-server/modules/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/irisnet/irishub-server/models"
)

const (
	CollectionNmStakeRoleDelegator = "stake_role_delegator"
)

type Delegator struct {
	Address string `json:"address" bson:"address"`
	PubKey  string `json:"pub_key" bson:"pub_key"`
	Shares  int64  `json:"shares" bson:"shares"`
	UpdateTime  time.Time   `json:"update_time" bson:"update_time"`
}

type DelegatorShares struct {
	TotalShares uint64 `json:"total_shares" bson:"total_shares"`
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

func (d Delegator) GetDelegatorListByAddress(address string, skip int,
	limit int, sorts []string) ([]Delegator, error) {

	query := bson.M{
		"address": address,
	}

	delegator, err := d.Query(query, skip, limit, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return delegator, err
}

func (d Delegator) GetTotalSharesByAddress(address string) (DelegatorShares, error)  {
	var value DelegatorShares
	
	q := func(c *mgo.Collection) error {
		m := []bson.M{
			{
				"$project": bson.M{
					"_id": 0,
				},
				
			},
			{
				"$match": bson.M{
					"address": address,
				},
			},
			{
				"$group": bson.M{
					"_id": "$address",
					"total_shares": bson.M{"$sum": "$shares"},
				},
			},
		}
		return c.Pipe(m).One(&value)
	}
	
	err := models.ExecCollection(d.Name(), q)
	
	if err !=  nil {
		return DelegatorShares{}, err
	}
	return value, nil


}
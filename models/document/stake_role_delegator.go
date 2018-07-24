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
	Address       string    `json:"address" bson:"address"`
	ValidatorAddr string    `json:"pub_key" bson:"validator_addr"` // validator ValidatorAddress
	Shares        int64     `json:"shares" bson:"shares"`
	UpdateTime    time.Time `json:"update_time" bson:"update_time"`
}

type DelegatorShares struct {
	ValidatorAddr string `bson:"_id"`
	TotalShares float64 `json:"total_shares" bson:"total_shares"`
}


func (d Delegator) Name() string {
	return CollectionNmStakeRoleDelegator
}

func (d Delegator) PkKvPair() map[string]interface{} {
	return bson.M{"address": d.Address, "validator_addr": d.ValidatorAddr}
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

func (d Delegator) GetDelegatorListByAddressAndValidatorAddrs(address string, valAddrs []string,
	) ([]Delegator, error) {

	query := bson.M{
		"address": address,
		"validator_addr": &bson.M{
			"$in": valAddrs,
		},
		"shares": &bson.M{
			"$gt": 0,
		},
	}
	sorts := make([]string, 0)

	delegator, err := d.Query(query, 0, len(valAddrs), sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return delegator, err
}

func (d Delegator) GetDelegatorListByAddress(address string, skip int,
	limit int, sorts []string) ([]Delegator, error) {

	query := bson.M{
		"address": address,
		"shares": &bson.M{
			"$gt": 0,
		},
	}

	delegator, err := d.Query(query, skip, limit, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return delegator, err
}

func (d Delegator) GetTotalSharesByAddress(address string) ([]DelegatorShares, error)  {
	var value []DelegatorShares
	
	q := func(c *mgo.Collection) error {
		m := []bson.M{
			{
				"$match": bson.M{
					"address": address,
				},
			},
			{
				"$group": bson.M{
					"_id" : "$validator_addr",
					"total_shares": bson.M{"$sum": "$shares"},
				},
			},
		}
		return c.Pipe(m).All(&value)
	}
	
	err := models.ExecCollection(d.Name(), q)
	
	if err !=  nil {
		return nil, err
	}
	return value, nil


}
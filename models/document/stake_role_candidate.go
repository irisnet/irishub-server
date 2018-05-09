package document

import (

	"github.com/irisnet/iris-api-server/models"
	"github.com/irisnet/iris-api-server/modules/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleCandidate = "stake_role_candidate"
)

type Candidate struct {
	Address     string      `json:"address" bson:"address"` // owner
	PubKey      string      `json:"pub_key" bson:"pub_key"`
	Shares      int64       `json:"shares" bson:"shares"`
	VotingPower uint64      `json:"voting_power" bson:"voting_power"` // Voting power if pubKey is a considered a validator
	Description Description `json:"description" bson:"description"`  // Description terms for the candidate

	Delegators []Delegator `json:"delegators"`
}

func (d Candidate) Name() string {
	return CollectionNmStakeRoleCandidate
}

func (d Candidate) PkKvPair() map[string]interface{} {
	return bson.M{"pub_key": d.PubKey}
}

func (d Candidate) Query(
	query bson.M, skip int, limit int, sorts ...string,
	) (results []Candidate, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}


func (d Candidate) GetCandidatesList(sorts []string, skip int, limit int) ([]Candidate, error)  {
	query := bson.M{
		"shares": &bson.M{
			"$ne": 0,
		},
	}
	candidates, err := d.Query(query, skip, limit, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return candidates, err
}

func (d Candidate) GetCandidatesListByPubKeys(pubKeys []string) ([]Candidate, error)  {
	query := bson.M{
		"pub_key": &bson.M{
			"$in": pubKeys,
		},
	}
	sorts := make([]string, 0)

	candidates, err := d.Query(query, 0, len(pubKeys), sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return candidates, err
}


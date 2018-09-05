package document

import (
	"github.com/irisnet/irishub-server/models"
	"github.com/irisnet/irishub-server/modules/logger"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleCandidate = "stake_role_candidate"
)

type Candidate struct {
	Address     string         `json:"address" bson:"address"` // owner
	PubKey      string         `json:"pub_key" bson:"pub_key"`
	PubKeyAddr  string         `bson:"pub_key_addr"` // validator address
	Shares      float64        `json:"shares" bson:"tokens"`
	Revoked     bool           `bson:"revoked"`
	Description ValDescription `json:"description" bson:"description"` // Description terms for the candidate

	VotingPower float64 `json:"voting_power"` // Voting power if pubKey is a considered a validator
	UpTime      float64
	Delegators  []Delegator `json:"delegators"`
}

func (d Candidate) Name() string {
	return CollectionNmStakeRoleCandidate
}

func (d Candidate) PkKvPair() map[string]interface{} {
	return bson.M{"address": d.Address}
}

func (d Candidate) Query(
	query bson.M, skip int, limit int, sorts ...string,
) (results []Candidate, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}

func (d Candidate) GetCandidatesList(q string, sorts []string, skip int, limit int) ([]Candidate, error) {
	query := bson.M{
		"revoked": false,
	}
	if q != "" {
		query["description.moniker"] = &bson.M{
			"$regex":   q,
			"$options": "$i",
		}
	}
	candidates, err := d.Query(query, skip, limit, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return candidates, err
}

func (d Candidate) GetCandidatesListByValidatorAddrs(valAddrs []string) ([]Candidate, error) {
	query := bson.M{
		"address": &bson.M{
			"$in": valAddrs,
		},
		"revoked": false,
	}
	sorts := make([]string, 0)

	candidates, err := d.Query(query, 0, len(valAddrs), sorts...)

	if err != nil {
		logger.Error.Println(err)
	}

	return candidates, err
}

func (d Candidate) GetTotalShares() (float64, error) {
	type result struct {
		Id          string  `bson:"_id"`
		TotalShares float64 `bson:"total_shares"`
	}
	var value result

	q := func(c *mgo.Collection) error {
		m := []bson.M{
			{
				"$match": bson.M{
					"revoked": false,
				},
			},
			{
				"$group": bson.M{
					"_id":          "test",
					"total_shares": bson.M{"$sum": "$tokens"},
				},
			},
		}
		return c.Pipe(m).One(&value)
	}

	err := models.ExecCollection(d.Name(), q)

	if err != nil && err.Error() != mgo.ErrNotFound.Error() {
		return 0, err
	}
	return value.TotalShares, nil
}

func (d Candidate) GetCandidateDetail(valAddr string) (Candidate, error) {
	query := bson.M{
		"address": valAddr,
	}
	sorts := make([]string, 0)

	candidates, err := d.Query(query, 0, 1, sorts...)

	if err != nil {
		logger.Error.Println(err)
	}
	if len(candidates) > 0 {
		return candidates[0], err
	}
	return Candidate{}, err
}

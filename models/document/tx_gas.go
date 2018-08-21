package document

import (
	"github.com/irisnet/irishub-server/models"
	"github.com/irisnet/irishub-server/utils/constants"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const CollectionNmTxGas = "tx_gas"

type TxGas struct {
	TxType   string   `bson:"tx_type"`
	GasUsed  GasUsed  `bson:"gas_used"`
	GasPrice GasPrice `bson:"gas_price"`
}

type GasUsed struct {
	MinGasUsed float64 `bson:"min_gas_used"`
	MaxGasUsed float64 `bson:"max_gas_used"`
	AvgGasUsed float64 `bson:"avg_gas_used"`
}

type GasPrice struct {
	Denom       string  `bson:"denom"`
	MinGasPrice float64 `bson:"min_gas_price"`
	MaxGasPrice float64 `bson:"max_gas_price"`
	AvgGasPrice float64 `bson:"avg_gas_price"`
}

func (d TxGas) Name() string {
	return CollectionNmTxGas
}

func (d TxGas) PkKvPair() map[string]interface{} {
	return bson.M{"tx_type": d.TxType}
}

func (d TxGas) Query(query bson.M, fields bson.M,
	skip int, limit int, sorts ...string) (results []TxGas, err error) {
	exop := func(c *mgo.Collection) error {
		return c.Find(query).Select(fields).Sort(sorts...).Skip(skip).Limit(limit).All(&results)
	}
	return results, models.ExecCollection(d.Name(), exop)
}

func (d TxGas) GetTxGas(txType string) (TxGas, error) {
	var (
		sorts    []string
		dbTxType string
	)

	dbTxType = constants.TxTypeFrontMapDb[txType]

	// not supported tx type
	if dbTxType == "" {
		return TxGas{}, nil
	}
	query := bson.M{
		"tx_type": dbTxType,
	}
	fields := bson.M{}
	skip := 0
	limit := 1

	res, err := d.Query(query, fields, skip, limit, sorts...)
	if err != nil {
		return TxGas{}, err
	}
	// result is empty
	if res == nil {
		return TxGas{}, nil
	}
	return res[0], nil
}

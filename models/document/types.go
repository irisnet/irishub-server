package document

type Coins []Coin

type Coin struct {
	Denom  string `bson:"denom"`
	Amount int64  `bson:"amount"`
}



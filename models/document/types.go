package document

type Coins []Coin

type Coin struct {
	Denom  string `bson:"denom"`
	Amount float64  `bson:"amount"`
}



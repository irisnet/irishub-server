package vo

const (
	CoinTypeIris  = "iris"
	CoinTypeAtto  = "iris-atto"
	CoinTypeFemto = "iris-femto"
	CoinTypePico  = "iris-pico"
	CoinTypeNano  = "iris-nano"
	CoinTypeMicro = "iris-micro"
	CoinTypeMilli = "iris-milli"
)

var coinsMap = make(map[string]float64)

func init() {
	coinsMap[CoinTypeIris] = float64(1)
	coinsMap[CoinTypeMilli] = float64(1000)
	coinsMap[CoinTypeMicro] = float64(1000000)
	coinsMap[CoinTypeNano] = float64(1000000000)
	coinsMap[CoinTypePico] = float64(1000000000000)
	coinsMap[CoinTypeFemto] = float64(1000000000000000)
	coinsMap[CoinTypeAtto] = float64(1000000000000000000)
}

type Fee struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}

type Address struct {
	Chain string `json:"chain"`
	App   string `json:"app"`
	Addr  string `json:"addr"`
}

type Coin struct {
	Denom  string  `json:"denom"`
	Amount float64 `json:"amount"`
}

func (coin Coin) Covert(denom string) Coin {
	srcPreci := coinsMap[coin.Denom]
	dstPreci := coinsMap[denom]

	dstAmt := coin.Amount * (dstPreci / srcPreci)
	return Coin{
		Denom:  denom,
		Amount: dstAmt,
	}
}

type Memo struct {
	Id   int64  `json:"id,omitempty"`
	Text []byte `json:"text,omitempty"`
}

type PubKey struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

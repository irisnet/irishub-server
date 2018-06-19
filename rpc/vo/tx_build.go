package vo

type BuildTxReqVO struct {
	Fees Fee `json:"fees"`
	Multi bool `json:"multi"`
	Sequence int64 `json:"sequence"`
	From Address `json:"from"`
	To Address `json:"to"`
	Amount []Coin `json:"amount"`
	Memo Memo `json:"memo"`
}

package vo

type BuildTxVO struct {
	Fees Fee `json:"fees"`
	Multi bool `json:"multi"`
	Sequence uint64 `json:"sequence"`
	From Address `json:"from"`
	To Address `json:"to"`
	Amount []Coin `json:"amount"`
	Memo Memo `json:"memo"`
}

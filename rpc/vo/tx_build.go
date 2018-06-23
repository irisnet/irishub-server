package vo

type BuildTxReqVO struct {
	Fees Fee `json:"fees"`
	Multi bool `json:"multi"`
	Sequence int64 `json:"sequence"`
	From Address `json:"from"`
	To Address `json:"to"`
	Amount []Coin `json:"amount"`
	Memo Memo `json:"memo"`
	TxType string `json:"txType"`
}

type BuildDelegateTxReqVO struct {
	Fees Fee `json:"fees"`
	Sequence int64 `json:"sequence"`
	From Address `json:"from"`
	PubKey PubKey `json:"pub_key"`
	Amount Coin `json:"amount"`
}

type BuildUnBondTxReqVO struct {
	Fees Fee `json:"fees"`
	Sequence int64 `json:"sequence"`
	From Address `json:"from"`
	PubKey PubKey `json:"pub_key"`
	Amount int64 `json:"amount"`
}

type BuildTxResVO struct {
	Data []byte
	Ext []byte
}




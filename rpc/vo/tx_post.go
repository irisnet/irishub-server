package vo

type PostTxReqVO struct {
	Tx []byte `json:"tx,omitempty"`
}

type PostTxResVO struct {
	TxHash string
}

type SimulateTxReqVO struct {
	Tx []byte `json:"tx,omitempty"`
}

type Record struct {
	ValAddress string `json:"valAddress"`
	Name       string `json:"name"`
	Amount     *Coin  `json:"amount"`
}

type SimulateTxResVO struct {
	Records []Record `json:"records"`
	Gas     int64    `json:"gas"`
}

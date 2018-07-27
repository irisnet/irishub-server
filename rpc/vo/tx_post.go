package vo

type PostTxReqVO struct {
	Tx []byte `json:"tx,omitempty"`
}

type PostTxResVO struct {
	TxHash string
}



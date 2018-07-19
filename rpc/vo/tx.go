package vo

import "github.com/irisnet/irishub-server/models/document"

type TxListReqVO struct {
	Address              string   `json:"address,omitempty"`
	Page                 int64   `json:"page,omitempty"`
	PerPage              int64   `json:"perPage,omitempty"`
	Status               string   `json:"status,omitempty"`
	Type                 string   `json:"type,omitempty"`
	StartTime            string   `json:"startTime,omitempty"`
	EndTime              string   `json:"endTime,omitempty"`
	Sort                 string   `json:"sort,omitempty"`
	Q                    string   `json:"q,omitempty"`
}

type TxListResVO struct {
	Txs []document.CommonTx
}

type TxDetailReqVO struct {
	TxHash string `json:"txHash,omitempty"`
}

type TxDetailResVO struct {
	Tx document.CommonTx
}




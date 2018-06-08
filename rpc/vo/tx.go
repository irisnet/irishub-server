package vo

import "github.com/irisnet/irishub-server/models/document"

type TxListReqVO struct {
	Address              string   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Page                 uint64   `protobuf:"varint,2,opt,name=page" json:"page,omitempty"`
	PerPage              uint64   `protobuf:"varint,3,opt,name=perPage" json:"perPage,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	StartTime            string   `protobuf:"bytes,6,opt,name=startTime" json:"startTime,omitempty"`
	EndTime              string   `protobuf:"bytes,7,opt,name=endTime" json:"endTime,omitempty"`
	Sort                 string   `protobuf:"bytes,8,opt,name=sort" json:"sort,omitempty"`
	Q                    string   `protobuf:"bytes,9,opt,name=q" json:"q,omitempty"`
}

type TxListResVO struct {
	Txs []document.CommonTx
}



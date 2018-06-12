package vo

type ShareReqVO struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

type ShareResVO struct {
	TotalShare uint64
}


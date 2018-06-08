package vo

type BalanceReqVO struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

type BalanceResVO struct {
	Height uint64 `json:"height"`
	Data BalanceResDataVO `json:"data"`
	
}

type BalanceResDataVO struct {
	Coins []*Coin `json:"coins"`
}




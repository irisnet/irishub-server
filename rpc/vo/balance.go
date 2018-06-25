package vo

type BalanceReqVO struct {
	Address string `json:"address,omitempty"`
}

type BalanceResVO struct {
	Height int64 `json:"height"`
	Data BalanceResDataVO `json:"data"`
	
}

type BalanceResDataVO struct {
	Coins []*Coin `json:"coins"`
}




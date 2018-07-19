package vo

type AccountNumReqVO struct {
	Address string
}

type AccountNumResVO struct {
	AccountNum int64 `json:""`
}

type SequenceReqVO struct {
	Address string `json:"address"`
}

type SequenceResVO struct {
	Sequence int64 `json:"data"`
	Ext []byte
}

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



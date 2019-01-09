package vo

import "github.com/irisnet/irishub-server/models/document"

type RewardInfoReqVO struct {
	DelAddr string `json:"delAddr"`
	ValAddr string `json:"valAddr"`
}

type RewardInfoResVo struct {
	DelAddr      string              `json:"delAddr"`
	WithdrawAddr string              `json:"withdrawAddr"`
	Txs          []document.CommonTx `json:"txs"`
}

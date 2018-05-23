package vo

type StakeTxListVO struct {
	BaseVO
	Address string `form:"address" binding:"required"`
	PubKey string `form:"pub_key"`
	TxType string `form:"tx_type"`
	StartTime string `form:"start_time"`
	EndTime string `form:"end_time"`
	Sort string `form:"sort"`
}


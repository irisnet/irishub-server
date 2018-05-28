package vo

type CommonTxListVO struct {
	BaseVO
	Address string `form:"address" binding:"required"`
	TxType string `form:"tx_type"`
	StartTime string `form:"start_time"`
	EndTime string `form:"end_time"`
	Sort string `form:"sort"`
}

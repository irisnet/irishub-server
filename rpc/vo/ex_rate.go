package vo

type ExRateReqVO struct {
	ValidatorAddress string
}

type ExRateResVO struct {
	ExRate string `json:"token_shares_rate"`
}

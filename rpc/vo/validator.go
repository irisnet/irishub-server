package vo

import "github.com/irisnet/irishub-server/models/document"

// ==============================================
// request of get validator list
// ==============================================

type ValidatorListReqVO struct {
	Address string `json:"address,omitempty"`
	Page    int16  `json:"page,omitempty"`
	PerPage int16  `json:"perPage,omitempty"`
	Sort    string `json:"sort,omitempty"`
	Q       string `json:"q,omitempty"`
}

type ValidatorListResVO struct {
	Candidates []document.Candidate
}

// ==============================================
// request of get validator detail
// ==============================================

type ValidatorDetailReqVO struct {
	DelAddr string `json:"address,omitempty"`
	ValAddr string `json:"pubKey,omitempty"`
}

type ValidatorDetailResVO struct {
	Candidate document.Candidate
}

// ==============================================
// request of get validator exchange rate
// ==============================================

type ValidatorExRateReqVO struct {
	ValidatorAddress string
}

type ValidatorExRateResVO struct {
	ExRate float64 `json:"token_shares_rate"`
}

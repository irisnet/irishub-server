package vo

import (
	"github.com/irisnet/irishub-server/models/document"
)

// =============================================
// request of getDelegatorCandidateList
// =============================================

type DelegatorCandidateListReqVO struct {
	Address              string   `json:"address,omitempty"`
	Page                 int16   `json:"page,omitempty"`
	PerPage              int16   `json:"perPage,omitempty"`
	Sort                 string   `json:"sort,omitempty"`
	Q                    string   `json:"q,omitempty"`
}

type DelegatorCandidateListResVO struct {
	Candidates []document.Candidate
}

// =============================================
// request of getDelegatorTotalShares
// =============================================

type TotalShareReqVO struct {
	Address string `json:"address,omitempty"`
}

type TotalShareResVO struct {
	TotalShare float64
}

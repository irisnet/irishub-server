package vo

import "github.com/irisnet/irishub-server/models/document"

type CandidateListReqVO struct {
	Address              string   `json:"address,omitempty"`
	Page                 int16   `json:"page,omitempty"`
	PerPage              int16   `json:"perPage,omitempty"`
	Sort                 string   `json:"sort,omitempty"`
	Q                    string   `json:"q,omitempty"`
}

type CandidateListResVO struct {
	Candidates []document.Candidate
}

type CandidateDetailReqVO struct {
	Address string `json:"address,omitempty"`
	ValAddr string `json:"pubKey,omitempty"`
}

type CandidateDetailResVO struct {
	Candidate document.Candidate
}

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




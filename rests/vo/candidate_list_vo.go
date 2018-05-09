package vo

type CandidateListVo struct {
	BaseVO
	Address string `json:"address"`
	Sort string `json:"sort"`
	Q string `json:"q"`
}


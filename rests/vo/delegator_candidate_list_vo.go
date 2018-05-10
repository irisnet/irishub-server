package vo

type DelegatorCandidateListVo struct {
	BaseVO
	Address string `form:"address"`
	Sort string `form:"sort"`
	Q string `form:"q"`
}


package vo

type SequenceReqVO struct {
	Address string `json:"address"`
}

type SequenceResVO struct {
	Sequence uint64 `json:"data"`
	Height uint64 `json:"height"`
}




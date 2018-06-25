package vo

type SequenceReqVO struct {
	Address string `json:"address"`
}

type SequenceResVO struct {
	Sequence int64 `json:"data"`
	Height int64 `json:"height"`
}




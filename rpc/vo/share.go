package vo

type ShareReqVO struct {
	Address string `json:"address,omitempty"`
}

type ShareResVO struct {
	TotalShare uint64
}


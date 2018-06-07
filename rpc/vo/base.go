package vo

type Fee struct {
	Denom string `json:"denom"`
	Amount int64 `json:"amount"`
}

type Address struct {
	Chain string `json:"chain"`
	App string `json:"app"`
	Addr string `json:"addr"`
}

type Coin struct {
	Denom string `json:"denom"`
	Amount int64 `json:"amount"`
}

type Memo struct {
	Id uint64   `json:"id,omitempty"`
	Text []byte   `json:"text,omitempty"`
}







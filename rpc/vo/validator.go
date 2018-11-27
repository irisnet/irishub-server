package vo

import (
	"github.com/irisnet/irishub-server/models/document"
	"time"
)

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

type Validator struct {
	OperatorAddr       string      `json:"operator_address"`
	ConsPubKey         string      `json:"consensus_pubkey"`
	Jailed             bool        `json:"jailed"`
	Status             byte        `json:"status"`
	Tokens             string      `json:"tokens"`
	DelegatorShares    string      `json:"delegator_shares"`
	Description        Description `json:"description"`
	BondHeight         string      `json:"bond_height"`
	BondIntraTxCounter int16       `json:"bond_intra_tx_counter"`
	UnbondingHeight    string      `json:"unbonding_height"`
	UnbondingMinTime   time.Time   `json:"unbonding_time"`
	Commission         Commission  `json:"commission"`
}

type Description struct {
	Moniker  string `json:"moniker"`  // name
	Identity string `json:"identity"` // optional identity signature (ex. UPort or Keybase)
	Website  string `json:"website"`  // optional website link
	Details  string `json:"details"`  // optional details
}

type Commission struct {
	Rate          string    `json:"rate"`
	MaxRate       string    `json:"max_rate"`
	MaxChangeRate string    `json:"max_change_rate"`
	UpdateTime    time.Time `json:"update_time"`
}

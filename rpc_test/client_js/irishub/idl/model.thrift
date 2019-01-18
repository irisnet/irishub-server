namespace go model

struct DelegatorUnbondingDelegation {
    1: double tokens,
    2: string minTime
}

struct Delegator {
	1: string address,
	2: string valAddress,
	3: string shares,
    4: double bondedTokens,
    5: DelegatorUnbondingDelegation unbondingDelegation
}

struct CandidateDescription {
	1: string details,
	2: string identity,
	3: string moniker,
	4: string website
}

struct Candidate {
	1: string address,
	2: string pubKey,
	3: double shares,
	4: double votingPower,
	5: CandidateDescription description,
	6: list<Delegator> delegators,
	7: double upTime
	8: string type
	9: i8 number
	10: i8 lift
}

/** common exception
 * @param errCode, error code
 * @param errMsg, error message
 */
exception Exception {
  1: i32 errCode,
  2: string errMsg
}

// ==================================
// define method request and response
// ==================================

struct CandidateListRequest {
	1: string address,
	2: i16 page,
	3: i16 perPage,
	4: string sort,
	5: string q
}

struct CandidateDetailRequest {
	1: string delAddress
	2: string valAddress
}

struct DelegatorCandidateListRequest {
	1: string address,
	2: i16 page,
	3: i16 perPage,
	4: string sort,
	5: string q
}

struct TotalShareRequest {
	1: string address
}

struct TotalShareResponse {
	1: double totalShares,
	2: double bondedTokens,
	3: double unbondingTokens
}

struct ValidatorExRateRequest {
    1: string validatorAddress
}

struct ValidatorExRateResponse {
    1: double tokenSharesRate
}

struct withdrawInfo{
    1: string delAddr
    2: string withdrawAddr
}

struct WithdrawAddrRequest {
    1: list<string> delAddrs
}

struct WithdrawAddrResponse {
    1: list<withdrawInfo> withdrawInfo
}

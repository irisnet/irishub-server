package services

import (
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

type CandidateService struct {
}

func (s CandidateService) List(reqVO vo.CandidateListReqVO) (vo.CandidateListResVO, errors.IrisError)  {
	sorts := helper.ParseParamSort(reqVO.Sort)
	
	var (
		resVO vo.CandidateListResVO
	)
	
	skip, limit := helper.ParseParamPage(int(reqVO.Page), int(reqVO.PerPage))
	address := reqVO.Address
	q := reqVO.Q

	// query all candidates
	candidates, err := candidateModel.GetCandidatesList(q, sorts, skip, limit)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	if candidates == nil {
		return resVO, irisErr
	}

	// get total shares
	totalShares, err := s.getTotalShares()
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// query detail of candidate which i have delegated
	var (
		validatorAddrs []string
	)
	for _, candidate := range candidates {
		validatorAddrs = append(validatorAddrs, candidate.Address)
	}
	delegator, err := delegatorModel.GetDelegatorListByAddressAndValidatorAddrs(address, validatorAddrs)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}
	for i, cd := range candidates {
		candidates[i] = s.buildCandidates(cd, delegator, totalShares)
	}
	
	resVO = vo.CandidateListResVO{
		Candidates: candidates,
	}

	return resVO, irisErr
}

func (s CandidateService) Detail(reqVO vo.CandidateDetailReqVO) (
	vo.CandidateDetailResVO, errors.IrisError) {
	
	var (
		resVO vo.CandidateDetailResVO
	)

	// query detail info of candidate
	candidate, err := candidateModel.GetCandidateDetail(reqVO.ValAddr)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// not found
	if candidate.Address == "" {
		return resVO, irisErr
	}

	// get total shares
	totalShares, err := candidateModel.GetTotalShares()
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// query detail of candidate which i have delegated
	var (
		validatorAddrs = []string{candidate.Address}
	)
	delegator, err := delegatorModel.GetDelegatorListByAddressAndValidatorAddrs(reqVO.Address, validatorAddrs)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}
	candidate = s.buildCandidates(candidate, delegator, totalShares)
	
	resVO = vo.CandidateDetailResVO{
		Candidate: candidate,
	}

	return resVO, irisErr
}

func (s CandidateService) DelegatorCandidateList(reqVO vo.DelegatorCandidateListReqVO) (vo.DelegatorCandidateListResVO, errors.IrisError)  {
	
	var (
		resVO vo.DelegatorCandidateListResVO
	)
	
	sorts := helper.ParseParamSort(reqVO.Sort)
	skip, limit := helper.ParseParamPage(int(reqVO.Page), int(reqVO.PerPage))
	
	address := reqVO.Address

	// query delegator list by address
	delegator, err := delegatorModel.GetDelegatorListByAddress(address, skip, limit, sorts)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	if delegator == nil {
		return resVO, irisErr
	}

	// get total shares
	totalShares, err := s.getTotalShares()
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// query all candidate which delegator have delegated
	var (
		valAddrs []string
	)
	for _, de := range delegator {
		valAddrs = append(valAddrs, de.ValidatorAddr)
	}
	candidates, err := candidateModel.GetCandidatesListByValidatorAddrs(valAddrs)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	for i, cd := range candidates {
		candidates[i] = s.buildCandidates(cd, delegator, totalShares)
	}
	
	resVO = vo.DelegatorCandidateListResVO{
		Candidates: candidates,
	}

	return resVO, irisErr
}

// build data
func (s CandidateService) buildCandidates(
	cd document.Candidate,
	delegator []document.Delegator,
	totalShares float64,
) document.Candidate {

	delegators := make([]document.Delegator, 0)
	for _, de := range delegator {
		if cd.Address == de.ValidatorAddr {
			delegators = append(delegators, de)
			cd.Delegators = delegators
			break
		}
	}
	if totalShares != 0 {
		cd.VotingPower = float64(cd.Shares) / totalShares
	}


	return cd
}

// get total shares
func (s CandidateService) getTotalShares() (float64, error) {
	return candidateModel.GetTotalShares()
}

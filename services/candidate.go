package services

import (
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

type CandidateService struct {
}

func (s CandidateService) List(listVo vo.CandidateListReqVO) (vo.CandidateListResVO, errors.IrisError)  {
	sorts := helper.ParseParamSort(listVo.Sort)
	
	var (
		resVO vo.CandidateListResVO
	)
	
	skip, limit := helper.ParseParamPage(int(listVo.Page), int(listVo.PerPage))
	address := listVo.Address

	// query all candidates
	candidates, err := candidateModel.GetCandidatesList(sorts, skip, limit)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// get total shares
	totalShares, err := s.getTotalShares()
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// query detail of candidate which i have delegated
	var (
		pubKeys []string
	)
	for _, candidate := range candidates {
		pubKeys = append(pubKeys, candidate.PubKey)
	}
	delegator, err := delegatorModel.GetDelegatorListByAddressAndPubKeys(address, pubKeys)
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

// func (s CandidateService) DelegatorCandidateList(reqVO vo.CandidateDetailReqVO) ([]document.Candidate, errors.IrisError)  {
// 	sort := listVo.Sort
// 	var sorts []string
// 	if sort != "" {
// 		sorts = strings.Split(sort, ",")
// 	} else {
// 		sorts = []string{"-update_time"}
// 	}
//
// 	skip := (listVo.Page - 1) * listVo.PerPage
// 	limit := listVo.PerPage
// 	address := listVo.Address
//
// 	// query delegator list by address
// 	delegator, err := delegatorModel.GetDelegatorListByAddress(address, skip, limit, sorts)
// 	if err != nil {
// 		return nil, ConvertSysErr(err)
// 	}
//
// 	// get total shares
// 	totalShares, err := s.getTotalShares()
// 	if err != nil {
// 		return nil, ConvertSysErr(err)
// 	}
//
// 	// query all candidate which delegator have delegated
// 	var (
// 		pubKeys []string
// 	)
// 	for _, de := range delegator {
// 		pubKeys = append(pubKeys, de.PubKey)
// 	}
// 	candidates, err := candidateModel.GetCandidatesListByPubKeys(pubKeys)
// 	if err != nil {
// 		return nil, ConvertSysErr(err)
// 	}
//
// 	for i, cd := range candidates {
// 		candidates[i] = s.buildCandidates(cd, delegator, totalShares)
// 	}
//
// 	return candidates, irisErr
// }

func (s CandidateService) Detail(reqVO vo.CandidateDetailReqVO) (
	vo.CandidateDetailResVO, errors.IrisError) {
	
	var (
		resVO vo.CandidateDetailResVO
	)

	// query detail info of candidate
	candidate, err := candidateModel.GetCandidateDetail(reqVO.PubKey)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// get total shares
	totalShares, err := candidateModel.GetTotalShares()
	if err != nil {
		return resVO, ConvertSysErr(err)
	}

	// query detail of candidate which i have delegated
	var (
		pubKeys = []string{candidate.PubKey}
	)
	delegator, err := delegatorModel.GetDelegatorListByAddressAndPubKeys(reqVO.Address, pubKeys)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}
	candidate = s.buildCandidates(candidate, delegator, totalShares)
	
	resVO = vo.CandidateDetailResVO{
		Candidate: candidate,
	}

	return resVO, irisErr
}

// build data
func (s CandidateService) buildCandidates(
	cd document.Candidate,
	delegator []document.Delegator,
	totalShares uint64,
) document.Candidate {

	delegators := make([]document.Delegator, 0)
	for _, de := range delegator {
		if cd.PubKey == de.PubKey {
			delegators = append(delegators, de)
			cd.Delegators = delegators
			break
		}
	}
	cd.VotingPower = float64(cd.Shares) / float64(totalShares)

	return cd
}

// get total shares
func (s CandidateService) getTotalShares() (uint64, error) {
	return candidateModel.GetTotalShares()
}

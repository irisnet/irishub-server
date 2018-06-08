package services

import (
	"strings"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/rests/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

type CandidateService struct {
}

func (s CandidateService) List(listVo vo.CandidateListVo) ([]document.Candidate, errors.IrisError)  {
	sorts := helper.ParseParamSort(listVo.Sort)
	skip, limit := helper.ParseParamPage(listVo.Page, listVo.PerPage)
	address := listVo.Address

	// query all candidates
	candidates, err := candidateModel.GetCandidatesList(sorts, skip, limit)
	if err != nil {
		return nil, ConvertSysErr(err)
	}

	// get total shares
	totalShares, err := s.getTotalShares()
	if err != nil {
		return nil, ConvertSysErr(err)
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
		return nil, ConvertSysErr(err)
	}
	for i, cd := range candidates {
		candidates[i] = s.buildCandidates(cd, delegator, totalShares)
	}

	return candidates, irisErr
}

func (s CandidateService) DelegatorCandidateList(listVo vo.DelegatorCandidateListVo) ([]document.Candidate, errors.IrisError)  {
	sort := listVo.Sort
	var sorts []string
	if sort != "" {
		sorts = strings.Split(sort, ",")
	} else {
		sorts = []string{"-update_time"}
	}

	skip := (listVo.Page - 1) * listVo.PerPage
	limit := listVo.PerPage
	address := listVo.Address

	// query delegator list by address
	delegator, err := delegatorModel.GetDelegatorListByAddress(address, skip, limit, sorts)
	if err != nil {
		return nil, ConvertSysErr(err)
	}

	// get total shares
	totalShares, err := s.getTotalShares()
	if err != nil {
		return nil, ConvertSysErr(err)
	}

	// query all candidate which delegator have delegated
	var (
		pubKeys []string
	)
	for _, de := range delegator {
		pubKeys = append(pubKeys, de.PubKey)
	}
	candidates, err := candidateModel.GetCandidatesListByPubKeys(pubKeys)
	if err != nil {
		return nil, ConvertSysErr(err)
	}

	for i, cd := range candidates {
		candidates[i] = s.buildCandidates(cd, delegator, totalShares)
	}

	return candidates, irisErr
}

func (s CandidateService) Detail(pubKey string, address string) (document.Candidate, errors.IrisError) {

	// query detail info of candidate
	candidate, err := candidateModel.GetCandidateDetail(pubKey)
	if err != nil {
		return candidate, ConvertSysErr(err)
	}

	// get total shares
	totalShares, err := candidateModel.GetTotalShares()
	if err != nil {
		return document.Candidate{}, ConvertSysErr(err)
	}

	// query detail of candidate which i have delegated
	var (
		pubKeys = []string{candidate.PubKey}
	)
	delegator, err := delegatorModel.GetDelegatorListByAddressAndPubKeys(address, pubKeys)
	if err != nil {
		return document.Candidate{}, ConvertSysErr(err)
	}
	candidate = s.buildCandidates(candidate, delegator, totalShares)

	return candidate, irisErr
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
		}
		break
	}
	cd.VotingPower = float64(cd.Shares) / float64(totalShares)

	return cd
}

// get total shares
func (s CandidateService) getTotalShares() (uint64, error) {
	return candidateModel.GetTotalShares()
}

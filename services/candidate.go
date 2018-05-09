package services

import (
	"strings"

	"github.com/irisnet/iris-api-server/models/document"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
)

type CandidateService struct {
}

func (s CandidateService) List(listVo vo.CandidateListVo) ([]document.Candidate, errors.IrisError)  {
	sort := listVo.Sort
	var (
		sorts []string
	)
	if sort != "" {
		sorts = strings.Split(sort, ",")
	}
	skip := (listVo.Page - 1) * listVo.PerPage
	limit := listVo.PerPage
	address := listVo.Address

	// query all candidates
	candidates, err := candidateModel.GetCandidatesList(sorts, skip, limit)
	if err != nil {
		irisErr = irisErr.New(errors.EC50001, err.Error())
		return nil, irisErr
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
		irisErr = irisErr.New(errors.EC50001, err.Error())
		return nil, irisErr
	}
	for i, cd := range candidates {
		delegators := make([]document.Delegator, 0)
		for _, de := range delegator {
			if cd.PubKey == de.PubKey {
				delegators = append(delegators, de)
				cd.Delegators = delegators
			}
			break
		}
		candidates[i] = cd
	}

	return candidates, irisErr
}

func (s CandidateService) DelegatorCandidateList(listVo vo.DelegatorCandidateListVo) ([]document.Candidate, errors.IrisError)  {
	sort := listVo.Sort
	var sorts []string
	if sort != "" {
		sorts = strings.Split(sort, ",")
	}

	skip := (listVo.Page - 1) * listVo.PerPage
	limit := listVo.PerPage
	address := listVo.Address

	// query delegator list by address
	delegator, err := delegatorModel.GetDelegatorListByAddress(address, skip, limit, sorts)
	if err != nil {
		irisErr = irisErr.New(errors.EC50001, err.Error())
		return nil, irisErr
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
		irisErr = irisErr.New(errors.EC50001, err.Error())
		return nil, irisErr
	}

	for i, cd := range candidates {
		delegators := make([]document.Delegator, 0)
		for _, de := range delegator {
			if cd.PubKey == de.PubKey {
				delegators = append(delegators, de)
				cd.Delegators = delegators
			}
			break
		}
		candidates[i] = cd
	}

	return candidates, irisErr
}

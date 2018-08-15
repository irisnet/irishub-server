package services

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
)

type ValidatorService struct {
}

func (s ValidatorService) List(reqVO vo.ValidatorListReqVO) (vo.ValidatorListResVO, errors.IrisError) {
	sorts := helper.ParseParamSort(reqVO.Sort)

	var (
		resVO vo.ValidatorListResVO
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
	totalShares, err := s.GetTotalShares()
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
		candidates[i] = s.buildValidator(cd, delegator, totalShares)
	}

	resVO = vo.ValidatorListResVO{
		Candidates: candidates,
	}

	return resVO, irisErr
}

func (s ValidatorService) GetValidatorExRate(reqVO vo.ValidatorExRateReqVO) (
	vo.ValidatorExRateResVO, errors.IrisError) {

	var (
		resVO vo.ValidatorExRateResVO
	)

	address := reqVO.ValidatorAddress

	uri := fmt.Sprintf(constants.HttpUriGetExRate, address)

	statusCode, resBytes := HttpClientGetData(uri)

	// statusCode != 200
	if !helper.SliceContains(constants.SuccessStatusCodes, statusCode) {
		return resVO, ConvertSysErr(fmt.Errorf(string(resBytes)))
	}

	if err := json.Unmarshal(resBytes, &resVO); err != nil {
		return resVO, ConvertSysErr(err)
	}

	return resVO, irisErr
}

// get total shares
func (s ValidatorService) GetTotalShares() (float64, error) {
	return candidateModel.GetTotalShares()
}

// build data
func (s ValidatorService) buildValidator(cd document.Candidate,
	delegators []document.Delegator, totalShares float64) document.Candidate {

	resDelegators := make([]document.Delegator, 0)
	for _, d := range delegators {
		if cd.Address == d.ValidatorAddr {
			// calculate token by delegator's share
			reqVO := vo.ValidatorExRateReqVO{
				ValidatorAddress: cd.Address,
			}
			res, err := s.GetValidatorExRate(reqVO)
			if err.IsNotNull() {
				logger.Error.Printf("Can't getValidatorExRate, err is %v\n", err)
			}
			d.BondedTokens = float64(d.Shares) * res.ExRate

			resDelegators = append(resDelegators, d)

			cd.Delegators = resDelegators
			break
		}
	}
	if totalShares != 0 {
		cd.VotingPower = float64(cd.Shares) / totalShares
	}

	return cd
}

func (s ValidatorService) Detail(reqVO vo.ValidatorDetailReqVO) (
	vo.ValidatorDetailResVO, errors.IrisError) {

	var (
		resVO vo.ValidatorDetailResVO
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
	delegator, err := delegatorModel.GetDelegatorListByAddressAndValidatorAddrs(reqVO.DelAddr, validatorAddrs)
	if err != nil {
		return resVO, ConvertSysErr(err)
	}
	candidate = s.buildValidator(candidate, delegator, totalShares)

	resVO = vo.ValidatorDetailResVO{
		Candidate: candidate,
	}

	return resVO, irisErr
}

package services

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

type DelegatorService struct {
}

var (
	validatorService ValidatorService
)

func (s DelegatorService) DelegatorCandidateList(reqVO vo.DelegatorCandidateListReqVO) (vo.DelegatorCandidateListResVO, errors.IrisError) {

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
	totalShares, err := validatorService.GetTotalShares()
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
		candidates[i] = validatorService.buildCandidates(cd, delegator, totalShares)
	}

	resVO = vo.DelegatorCandidateListResVO{
		Candidates: candidates,
	}

	return resVO, irisErr
}

func (s DelegatorService) GetDelegatorTotalShare(reqVO vo.DelegatorTotalShareReqVO) (vo.DelegatorTotalShareResVO, errors.IrisError) {

	var (
		resVO                                                vo.DelegatorTotalShareResVO
		totalShares, totalBondedTokens, totalUnbondingTokens float64
	)

	delegatorShares, err := delegatorModel.GetTotalSharesByAddress(reqVO.Address)
	// can't find record by address
	if err != nil {
		return resVO, irisErr
	}

	if len(delegatorShares) > 0 {
		for _, v := range delegatorShares {

			reqVO := vo.ValidatorExRateReqVO{
				ValidatorAddress: v.ValidatorAddr,
			}
			res, err := validatorService.GetValidatorExRate(reqVO)

			if err.IsNotNull() {
				logger.Error.Printf("Can't get validator exRate, valAddr is %v\n", v.ValidatorAddr)
				continue
			}

			totalShares += v.TotalShares
			totalBondedTokens += v.TotalShares * res.ExRate
			totalUnbondingTokens += v.TotalUnbondingTokens
		}
	}

	resVO = vo.DelegatorTotalShareResVO{
		TotalShares:          totalShares,
		ToTalBondedTokens:    totalBondedTokens,
		ToTalUnbondingTokens: totalUnbondingTokens,
	}
	return resVO, irisErr
}

package services

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/utils/helper"
	"github.com/irisnet/irishub-server/modules/logger"
)

type DelegatorService struct {

}

var (
	validatorService ValidatorService
)

func (s DelegatorService) DelegatorCandidateList(reqVO vo.DelegatorCandidateListReqVO) (vo.DelegatorCandidateListResVO, errors.IrisError)  {

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

func (s DelegatorService) GetDelegatorTotalShare(reqVO vo.TotalShareReqVO) (vo.TotalShareResVO, errors.IrisError) {

	var (
		resVO       vo.TotalShareResVO
		totalTokens float64
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

			totalTokens += v.TotalShares * res.ExRate
		}
	}

	resVO = vo.TotalShareResVO{
		// TODO: set value of shares equal tokens,
		// next version will change correct
		TotalShare: totalTokens,
	}
	return resVO, irisErr
}

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
		resVO                vo.DelegatorCandidateListResVO
		tmValAddrs, valAddrs []string
		methodName           = "DelegatorValidatorList"
	)

	sorts := helper.ParseParamSort(reqVO.Sort)
	skip, limit := helper.ParseParamPage(int(reqVO.Page), int(reqVO.PerPage))

	address := reqVO.Address

	// query delegator list by address
	delegator, err := delegatorModel.GetDelegatorListByAddress(address, skip, limit, sorts)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	if delegator == nil {
		return resVO, irisErr
	}

	// get total shares
	totalShares, err := validatorService.GetTotalShares()
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	// query all candidate which delegator have delegated
	for _, de := range delegator {
		valAddrs = append(valAddrs, de.ValidatorAddr)
	}
	candidates, err := candidateModel.GetCandidatesListByValidatorAddrs(valAddrs)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	// get validator up time info
	for _, v := range candidates {
		tmValAddrs = append(tmValAddrs, v.PubKeyAddr)
	}
	valUpTimes, err := valUpTimeModel.GetUpTime(tmValAddrs)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	for i, cd := range candidates {
		candidates[i] = validatorService.buildValidator(cd, delegator, valUpTimes, totalShares)
	}

	resVO = vo.DelegatorCandidateListResVO{
		Candidates: candidates,
	}

	return resVO, irisErr
}

func (s DelegatorService) GetDelegatorTotalShare(reqVO vo.DelegatorTotalShareReqVO) (vo.DelegatorTotalShareResVO, errors.IrisError) {

	var (
		resVO                          vo.DelegatorTotalShareResVO
		totalShares, totalBondedTokens float64
		methodName                     = "GetDelegatorTotalShares"
	)

	// get total shares and bonded tokens
	// note: delegatorShares represent shares which delegator bonded on one validator,
	//       result is grouped by validator address
	delegatorShares, err := delegatorModel.GetTotalSharesByAddress(reqVO.Address)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
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
		}
	}

	// get total unbonding tokens
	totalUnbondingTokens, err := delegatorModel.GetTotalUnbondingTokens(reqVO.Address)
	if err != nil {
		logger.Error.Printf("%v: err is %v\n", methodName, err)
		return resVO, ConvertSysErr(err)
	}

	resVO = vo.DelegatorTotalShareResVO{
		TotalShares:          totalShares,
		ToTalBondedTokens:    totalBondedTokens,
		ToTalUnbondingTokens: totalUnbondingTokens,
	}
	return resVO, irisErr
}

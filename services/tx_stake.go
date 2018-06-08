package services

import (
	"time"
	
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/models/document"
	"github.com/irisnet/irishub-server/rests/vo"
	"github.com/irisnet/irishub-server/utils/helper"
)

type StakeTxService struct {
}

// get list of stake tx
func (s StakeTxService) GetList(vo vo.StakeTxListVO) ([]document.StakeTx, errors.IrisError) {
	var (
		startTime time.Time
		endTime time.Time
		err error
	)
	
	if vo.StartTime != "" {
		startTime, err = helper.ParseFullTime(vo.StartTime)
		if err != nil {
			return nil, irisErr.New(errors.EC40002, errors.EM40002)
		}
	}
	
	if vo.EndTime != "" {
		endTime, err = helper.ParseFullTime(vo.EndTime)
		if err != nil {
			return nil, irisErr.New(errors.EC40002, errors.EM40002)
		}
	}
	
	sorts := helper.ParseParamSort(vo.Sort)
	skip, limit := helper.ParseParamPage(vo.Page, vo.PerPage)
	
	
	stakeTxs, err := stakeTxModel.GetStakeTxList(vo.Address, vo.PubKey, vo.TxType, startTime, endTime, skip, limit, sorts)
	if err != nil {
		return nil, ConvertSysErr(err)
	}
	
	return stakeTxs, irisErr
}



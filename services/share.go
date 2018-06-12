package services

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
)

type ShareService struct {

}

func (s ShareService) GetDelegatorTotalShare(reqVO vo.ShareReqVO) (vo.ShareResVO, errors.IrisError) {
	
	var (
		resVO vo.ShareResVO
	)
	
	delegatorShares, err := delegatorModel.GetTotalSharesByAddress(reqVO.Address)
	// can't find record by address
	if err != nil {
		return resVO, irisErr
	}
	
	resVO = vo.ShareResVO{
		TotalShare: delegatorShares.TotalShares,
	}
	return resVO, irisErr
}

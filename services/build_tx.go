package services

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc/vo"
)

type BuildTxService struct {

}

// Deprecated: no longer use
func (s BuildTxService) BuildTx(reqVO vo.BuildTxReqVO) (vo.BuildTxResVO, errors.IrisError) {
	var (
		resVO vo.BuildTxResVO
	)
	
	return resVO, irisErr
}

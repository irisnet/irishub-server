package irishub

import (
	"github.com/irisnet/irishub-server/services"
	irisProtoc "github.com/irisnet/irisnet-rpc/irishub/codegen/server/model"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

type WithdrawInfoHandler struct {
}

func (c WithdrawInfoHandler) Handler(ctx context.Context, req *irisProtoc.WithdrawAddrRequest) (
	*irisProtoc.WithdrawAddrResponse, error) {

	if len(req.DelAddrs) == 0 {
		return nil, services.ConvertBadRequestErr(errors.New("delAddr is empty"))
	}

	var results = make([]*irisProtoc.WithdrawInfo, len(req.DelAddrs))

	for i, delAddr := range req.DelAddrs {
		addr, err := accountService.QueryWithdrawAddr(delAddr)
		if err.IsNotNull() {
			continue
		}
		results[i] = &irisProtoc.WithdrawInfo{
			DelAddr:      delAddr,
			WithdrawAddr: addr,
		}
	}

	return &irisProtoc.WithdrawAddrResponse{
		WithdrawInfo: results,
	}, nil
}

package blockchain

import (
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/rpc"
	"github.com/irisnet/irishub-server/services"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	"golang.org/x/net/context"
	"reflect"
)

var (
	txGasHandler TxGasHandler

	buildTxHandler BuildTxHandler
	buildTxService services.BuildTxService

	postTxHandler PostTxHandler
	postTxService services.PostTxService

	simulateTxHandler SimulateTxHandler
	simulateTxService services.SimulateTxService

	accountService  services.AccountService
	sequenceHandler SequenceHandler
	balanceHandler  BalanceHandler

	txListHandler TxListHandler
	txService     services.TxService

	txDetailHandler        TxDetailHandler
	queryRewardInfoHandler QueryRewardInfoHandler
)

func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)

	ok, er := rpc.DoFilters(reflect.TypeOf(req).String())
	if !ok {
		return nil, BuildException(er)
	}

	switch req.(type) {
	case *commonProtoc.TxGasRequest:
		res, err = txGasHandler.Handler(ctx, req.(*commonProtoc.TxGasRequest))
		break
	case *commonProtoc.BuildTxRequest:
		res, err = buildTxHandler.Handler(ctx, req.(*commonProtoc.BuildTxRequest))
		break
	case *commonProtoc.PostTxRequest:
		res, err = postTxHandler.Handler(ctx, req.(*commonProtoc.PostTxRequest))
		break
	case *commonProtoc.SimulateTxRequest:
		res, err = simulateTxHandler.Handler(ctx, req.(*commonProtoc.SimulateTxRequest))
		break
	case *commonProtoc.SequenceRequest:
		res, err = sequenceHandler.Handler(ctx, req.(*commonProtoc.SequenceRequest))
		break
	case *commonProtoc.BalanceRequest:
		res, err = balanceHandler.Handler(ctx, req.(*commonProtoc.BalanceRequest))
		break
	case *commonProtoc.TxListRequest:
		res, err = txListHandler.Handler(ctx, req.(*commonProtoc.TxListRequest))
		break
	case *commonProtoc.TxDetailRequest:
		res, err = txDetailHandler.Handler(ctx, req.(*commonProtoc.TxDetailRequest))
		break
	case *commonProtoc.RewardInfoRequest:
		res, err = queryRewardInfoHandler.Handler(ctx, req.(*commonProtoc.RewardInfoRequest))
		break
	}

	return res, err
}

func BuildException(err errors.IrisError) error {
	var (
		exception commonProtoc.Exception
	)
	exception.ErrCode = int32(err.ErrCode)
	exception.ErrMsg = err.ErrMsg
	return &exception
}

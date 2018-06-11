package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/irishub-server/services"
	"golang.org/x/net/context"
)

var (
	buildTxController  BuildTxController
	buildTxService services.BuildTxService
	
	postTxController   PostTxController
	postTxService  services.PostTxService
	
	sequenceController SequenceController
	sequenceService services.SequenceService
	
	balanceController BalanceController
	balanceService services.BalanceService
	
	txListController TxListController
	txService services.TxService
	
	txDetailController TxDetailController
)

func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	
	switch req.(type) {
	case *chainModel.BuildTxRequest:
		res, err = buildTxController.Handler(ctx, req.(*chainModel.BuildTxRequest))
		break
	case *chainModel.PostTxRequest:
		res, err = postTxController.Handler(ctx, req.(*chainModel.PostTxRequest))
		break
	case *chainModel.SequenceRequest:
		res, err = sequenceController.Handler(ctx, req.(*chainModel.SequenceRequest))
		break
	case *chainModel.BalanceRequest:
		res, err = balanceController.Handler(ctx, req.(*chainModel.BalanceRequest))
		break
	case *chainModel.TxListRequest:
		res, err = txListController.Handler(ctx, req.(*chainModel.TxListRequest))
		break
	case *chainModel.TxDetailRequest:
		res, err = txDetailController.Handler(ctx, req.(*chainModel.TxDetailRequest))
		break
	}
	
	return res, err
}

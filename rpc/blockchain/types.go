package blockchain

import (
	chainModel "github.com/irisnet/blockchain-rpc/codegen/server"
	"github.com/irisnet/iris-api-server/services"
	"golang.org/x/net/context"
)

var (
	buildTxController  BuildTxController
	buildTxService services.BuildTxService
	
	postTxController   PostTxController
	postTxService  services.PostTxService
)

func Handler(ctx context.Context, req interface{}) (interface{}, error) {
	var (
		res interface{}
		err error
	)
	
	switch req.(type) {
	case chainModel.BuildTxRequest:
		res, err = buildTxController.Handler(ctx, req.(*chainModel.BuildTxRequest))
		break
	case chainModel.PostTxRequest:
		res, err = postTxController.Handler(ctx, req.(*chainModel.PostTxRequest))
		break
	}
	
	return res, err
}

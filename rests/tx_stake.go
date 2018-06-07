package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
)

type StakeTxRoute struct {
}

func RegisterStakeTxRoute(r *gin.Engine)  {
	stakeTxRoute := StakeTxRoute{}
	
	rg := r.Group("/stake_txs")
	{
		rg.GET("", stakeTxRoute.GetList)
	}
}

func (r StakeTxRoute) GetList(c *gin.Context) {
	var listVo vo.StakeTxListVO
	err := c.ShouldBindQuery(&listVo)
	if err != nil {
		irisErr = irisErr.New(errors.EC40001, errors.EM40001 + err.Error())
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	response, irisErr := stakeTxService.GetList(listVo)
	if irisErr.IsNotNull() {
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	c.JSON(OK, BuildResponse(response))
}

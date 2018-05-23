package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/rests/errors"
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

// @Summary stake tx list
// @Description get list of stake tx
// @Tags txs
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param per_page query int true "per_page"
// @Param address query string true "user address"
// @Param pub_key query string false "pubKet of candidate"
// @Param start_time query string false "tx time"
// @Param end_time query string false "tx time"
// @Param sort query string false "order"
// @Success 200 {array} document.StakeTx "content of data"
// @router /stake_txs [get]
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

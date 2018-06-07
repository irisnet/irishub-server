package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
)

type CommonTxRoute struct {
}

func RegisterCommonTxRoute(r *gin.Engine)  {
	commonTxRoute := CommonTxRoute{}
	
	rg := r.Group("/txs")
	{
		rg.GET("", commonTxRoute.GetList)
	}
}

// @Summary tx list
// @Description get list of tx
// @Tags txs
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param per_page query int true "per_page"
// @Param address query string true "wallet address"
// @Param tx_type query string false "tx type" Enums(send, receive, delegate, unbond)
// @Param start_time query string false "tx time"
// @Param end_time query string false "tx time"
// @Param sort query string false "order" Enums(-time)
// @Success 200 {string} string "generate response data error, please see result"
// @router /txs [get]
func (r CommonTxRoute) GetList(c *gin.Context) {
	var listVo vo.CommonTxListVO
	err := c.ShouldBindQuery(&listVo)
	if err != nil {
		irisErr = irisErr.New(errors.EC40001, errors.EM40001 + err.Error())
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	response, irisErr := commonTxService.GetList(listVo)
	if irisErr.IsNotNull() {
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	c.JSON(OK, BuildResponse(response))
}

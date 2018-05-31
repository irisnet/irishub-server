package rests

import (
	"github.com/gin-gonic/gin"
)

type ShareRoute struct {
}


func RegisterRoutesShare(r *gin.Engine) {
	
	shareRoute := ShareRoute{}
	
	rg := r.Group("/shares")
	{
		rg.GET("/delegator/:address", shareRoute.GetTotalShares)
	}
}

// @Summary shares own by a delegator
// @Description get total shares which own by a delegator
// @Tags stake
// @Accept json
// @Produce json
// @Param address path string true "user address"
// @Success 200 {object} document.DelegatorShares "content of data"
// @router /shares/delegator/{address} [get]
func (r ShareRoute) GetTotalShares(c *gin.Context)  {
	address := c.Param("address")
	response, irisErr := delegatorService.GetTotalShares(address)
	if irisErr.IsNotNull() {
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	c.JSON(OK, response)
}

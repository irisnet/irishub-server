package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
)

type DelegatorRoute struct {
}


func RegisterRoutesDelegator(r *gin.Engine) {

	delegatorRoute := DelegatorRoute{}

	rg := r.Group("/delegators")
	{
		rg.GET("/:address/candidates", delegatorRoute.DelegatorCandidateList)
		rg.GET("/:address/candidates/:pub_key", delegatorRoute.DelegatorCandidateDetail)
	}
}

// @Summary delegator candidate list
// @Description list of candidate which delegated by address
// @Tags stake
// @Accept json
// @Produce json
// @Param address path string true "user address"
// @Param page query int true "page"
// @Param per_page query int true "per_page"
// @Param sort query string false "order"
// @Success 200 {array} document.Candidate "content of data"
// @router /delegators/{address}/candidates [get]
func (r DelegatorRoute) DelegatorCandidateList(c *gin.Context)  {
	address := c.Param("address")
	var (
		listVo vo.DelegatorCandidateListVo
	)
	err := c.ShouldBindQuery(&listVo)
	if err != nil {
		irisErr = irisErr.New(errors.EC40001, errors.EM40001)
		c.JSON(OK, BuildExpResponse(irisErr))
	}
	listVo.Address = address

	response, iriErr := candidateService.DelegatorCandidateList(listVo)
	if iriErr.IsNotNull() {
		c.JSON(OK, BuildExpResponse(irisErr))
	}
	c.JSON(OK, response)

}

// @Summary delegator candidate detail
// @Description detail of candidate which delegated by address
// @Tags stake
// @Accept json
// @Produce json
// @Param address path string true "user address"
// @Param pub_key path string true "public key of candidate"
// @Success 200 {object} document.Candidate "content of data"
// @router /delegators/{address}/candidates/{pub_key} [get]
func (r DelegatorRoute) DelegatorCandidateDetail(c *gin.Context)  {
	pubKey := c.Param("pub_key")
	address := c.Param("address")

	response, iriErr := candidateService.Detail(pubKey, address)
	if iriErr.IsNotNull() {
		c.JSON(OK, BuildExpResponse(irisErr))
	}
	c.JSON(OK, response)

}
package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
	"github.com/irisnet/iris-api-server/services"
)

type CandidateRoute struct {
}

var (
	service services.CandidateService
	irisErr errors.IrisError
)

func RegisterRoutesCandidate(r *gin.Engine) {

	candidateRoute := CandidateRoute{}

	rg := r.Group("/v1/candidates")
	{
		rg.GET("", candidateRoute.List)
	}
}

func (cr CandidateRoute) List(c *gin.Context) {
	var listVo vo.CandidateListVo
	err := c.ShouldBindQuery(&listVo)
	if err != nil {
		irisErr = irisErr.New(errors.EC40001, errors.EM40001)
		c.JSON(HttpStatusOk, BuildExpResponse(irisErr))
	}
	candidates, irisErr := service.List(listVo)
	if irisErr.IsNotNull() {
		c.JSON(HttpStatusOk, BuildExpResponse(irisErr))
	}
	c.JSON(HttpStatusOk, BuildResponse(candidates))
}



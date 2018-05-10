package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/rests/errors"
	"github.com/irisnet/iris-api-server/rests/vo"
)

type CandidateRoute struct {
}

func RegisterRoutesCandidate(r *gin.Engine) {

	candidateRoute := CandidateRoute{}

	rg := r.Group("/candidates")
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
	response, irisErr := candidateService.List(listVo)
	if irisErr.IsNotNull() {
		c.JSON(HttpStatusOk, BuildExpResponse(irisErr))
	}
	c.JSON(HttpStatusOk, BuildResponse(response))
}



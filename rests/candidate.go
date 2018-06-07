package rests

import (
	"github.com/gin-gonic/gin"
	"github.com/irisnet/iris-api-server/errors"
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

// @Summary candidate list
// @Description the list of candidates
// @Tags stake
// @Accept json
// @Produce json
// @Param address query string true "user address"
// @Param page query int true "page"
// @Param per_page query int true "per_page"
// @Param sort query string false "order"
// @Success 200 {array} document.Candidate "content of data"
// @router /candidates [get]
func (cr CandidateRoute) List(c *gin.Context) {
	var listVo vo.CandidateListVo
	err := c.ShouldBindQuery(&listVo)
	if err != nil {
		irisErr = irisErr.New(errors.EC40001, errors.EM40001)
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	response, irisErr := candidateService.List(listVo)
	if irisErr.IsNotNull() {
		c.JSON(OK, BuildExpResponse(irisErr))
		return
	}
	c.JSON(OK, BuildResponse(response))
}



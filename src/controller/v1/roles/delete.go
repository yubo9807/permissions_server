package roles

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	type Params struct {
		Id string `form:"id" binding:"required"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.State.ErrorParams()
		return
	}

	spider.Correlation.DeleteCorrelationRoles(params.Id)
	spider.Common.Delete("roles", params.Id)
	service.State.Success()
}

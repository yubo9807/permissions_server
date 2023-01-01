package menu

import (
	"server/src/service"
	"server/src/spider"

	"github.com/gin-gonic/gin"
)

// 更新基本信息
func Update(ctx *gin.Context) {
	type Params struct {
		Id     string `form:"id" binding:"required"`
		Name   string `form:"name" binding:"required"`
		Hidden bool   `form:"hidden"`
		Title  string `form:"title"`
	}
	var params Params
	if err := ctx.ShouldBind(&params); err != nil {
		service.ErrorParams()
		return
	}

	rows := spider.MenuQuery(params.Name, "")
	if len(rows) > 0 {
		service.ErrorCustom("菜单'" + params.Name + "'已存在")
		return
	}

	spider.MenuModify(params.Id, params.Name, params.Hidden, params.Title)
	service.Success()
}

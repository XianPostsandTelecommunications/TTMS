/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 18:49
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
)

type tag struct{}

func (tag) AddTagForMovie(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.AddTagParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.Tag.AddTagsForMovie(param.MovieID, param.Tag); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (tag) GetTagsFromMovie(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.GetTagsParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	rep, err := logic.Group.Tag.GetTagsFromMovie(param.MovieID)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rep)
}

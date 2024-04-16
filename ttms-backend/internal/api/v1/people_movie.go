/**
 * @Author: lenovo
 * @Description:
 * @File:  people_movie
 * @Version: 1.0.0
 * @Date: 2023/06/01 22:24
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/middleware"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
)

type pm struct{}

func (pm) UserMovieAction(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.UserMovie2Action
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	//等于1的话,直接就是赞成
	context, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(err)
		return
	}

	if param.Choice == 1 {
		if err := logic.Group.Pm.CheckIsRepeat(context.ID, param.MovieID); err != nil {
			rly.Reply(err)
			return
		}
		_, err := logic.Group.Pm.CreateFavorToMovie(ctx, param.MovieID)
		if err != nil {
			rly.Reply(err)
			return
		}
	} else {
		if err := logic.Group.Pm.DelUserFavor(ctx, param.MovieID); err != nil {
			rly.Reply(err)
			return
		}
	}
	rly.Reply(nil)
}

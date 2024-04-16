/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema
 * @Version: 1.0.0
 * @Date: 2023/06/05 8:17
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
	"mognolia/internal/pkg/utils"
)

type cinema struct{}

func (c *cinema) CreateCinema(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.CreateCinema
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.S().Infof("failed to shouldBindJSON: %v", err)
		return
	}
	if err := logic.Group.Cinema.CreateCinema(param); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (c *cinema) DelCinema(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.DeleteCinema
	if err := ctx.ShouldBindJSON(&param); err != nil {
		zap.S().Infof("failed to shouldBindJSON: %v", err)
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.Cinema.DeleteCinema(param.CinemaID); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (c *cinema) GetAllCinemaByPage(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	iStr := ctx.Param("page")
	i := utils.StringToIDMust(iStr)
	result, err := logic.Group.Cinema.GetAllCinemaByPage(int(i))
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.ReplyList(nil, result)
}

func (c *cinema) UpdateCinema(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.UpdateCinemaInfo
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.Cinema.UpdateCinemaAvatarOrName(param); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (c *cinema) GetCinemaDetails(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.GetCinemaDetails
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	cinemaDetail, err := logic.Group.Cinema.GetCinemaDetails(param.CinemaID)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, cinemaDetail)
}

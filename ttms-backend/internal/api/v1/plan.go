/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 0:40
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mognolia/internal/api/base"
	"mognolia/internal/global"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
	"mognolia/internal/pkg/utils"
)

type plan struct{}

func (plan) CreatePlan(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.CreatePlan
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		global.Logger.Info("Error creating plan")
		return
	}
	result, err := logic.Group.Plan.CreatePlan(param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

func (plan) DelPlan(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.DelPlan
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.Plan.DelPlan(req.PlanID); err != nil {
		zap.S().Infof("logic.Group.Plan.DelPlan failed: %v", err)
		return
	}
	rly.Reply(nil)
}

func (plan) GetPlans(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	iStr := ctx.Param("page")
	i := utils.StringToIDMust(iStr)
	result, err := logic.Group.Plan.GetPlans(int(i))
	if err != nil {
		zap.S().Infof("logic.Group.GetPlans error err: %v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

func (plan) GetPlansByMovieIDAndPeriod(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.GetPlansByMovieIDAndPeriod
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := req.Check(); err != nil {
		rly.Reply(err)
		return
	}
	result, err := logic.Group.Plan.GetPlansByMovieIDAndPeriod(req)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

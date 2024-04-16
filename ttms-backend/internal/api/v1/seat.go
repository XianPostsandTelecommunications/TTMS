/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 6:24
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
)

type seat struct{}

func (s *seat) ShowSeats(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	req := request.ShowSeatsParam{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	result, err := logic.Group.Seat.ShowSeats(req.PlanID, req.CinemaID)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

func (s *seat) UpdateSeatStatus(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	param := request.UpdateSeatStatus{}
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.Seat.UpdateSeatStatus(param); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

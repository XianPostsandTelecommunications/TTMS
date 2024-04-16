/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/05 8:20
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/middleware"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"
	"mognolia/internal/pkg/utils"
)

type ticket struct{}

func (ticket) GetTicketsByPlanID(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	iStr := ctx.Param("planID")
	i := utils.StringToIDMust(iStr)
	result, err := logic.Group.Ticket.GetTicketsByPlan(i)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

func (ticket) SoldTicket(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.LockTickets
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(err)
		return
	}
	result, err := logic.Group.Ticket.LockTickets(content.ID, req)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)
}

func (ticket) PayTicket(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.PayTicketReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(err)
		return
	}

	if err := logic.Group.Ticket.PayTicket(content.ID, req); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (ticket) BackTicket(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.BackTicketParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(err)
		return
	}
	if err := logic.Group.Ticket.BackTicket(param, content.ID); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (ticket) LockOneTicket(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.LockOneTicketReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	content, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(err)
		return
	}
	t := logic.NewTicket()
	if err := t.LockOneTicket(req, content.ID); err != nil {
		rly.Reply(err)
		return
	}

}

func (ticket) SearchTicket(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var req request.SearchTicketReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	ok, err := logic.Group.Ticket.SearchTicket(req.OrderID)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, gin.H{
		"exist": ok,
	})
}

func (ticket) GetAllOrders(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	content, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(err)
		return
	}
	result, err := logic.Group.Ticket.GetAllOrders(content.ID)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, result)

}

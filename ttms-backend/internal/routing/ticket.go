/**
 * @Author: lenovo
 * @Description:
 * @File:  ticket
 * @Version: 1.0.0
 * @Date: 2023/06/05 8:14
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type ticket struct {
}

func (ticket) Init(r *gin.RouterGroup) {
	g := r.Group("/tickets", middleware.Auth())
	g.GET("/seats/:planID", v1.Group.Ticket.GetTicketsByPlanID)
	g.POST("/soldTicket", v1.Group.Ticket.SoldTicket)
	g.POST("/pay", v1.Group.Ticket.PayTicket)
	g.POST("/isPay", v1.Group.Ticket.SearchTicket)
	g.POST("/backTicket", v1.Group.Ticket.BackTicket)
	g.GET("/getOrders", v1.Group.Ticket.GetAllOrders)
}

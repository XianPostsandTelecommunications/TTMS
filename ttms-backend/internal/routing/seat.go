/**
 * @Author: lenovo
 * @Description:
 * @File:  seat
 * @Version: 1.0.0
 * @Date: 2023/06/04 22:24
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type seat struct{}

func (seat) Init(r *gin.RouterGroup) {
	g := r.Group("/seat", middleware.Auth(), middleware.AuthManager())
	{
		g.GET("/showSeatsByCinemaID/:cinemaID", v1.Group.Seat.ShowSeats)
		g.PUT("/update", v1.Group.Seat.UpdateSeatStatus)
	}
}

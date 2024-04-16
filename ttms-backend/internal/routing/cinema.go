/**
 * @Author: lenovo
 * @Description:
 * @File:  cinema
 * @Version: 1.0.0
 * @Date: 2023/06/04 1:08
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type cinema struct{}

func (c *cinema) Init(r *gin.RouterGroup) {
	g := r.Group("/cinema", middleware.Auth(), middleware.AuthManager())
	{
		g.POST("/create", v1.Group.Cinema.CreateCinema)
		g.DELETE("/delete", v1.Group.Cinema.DelCinema)
		g.GET("/list/:page", v1.Group.Cinema.GetAllCinemaByPage)
		g.PUT("/update", v1.Group.Cinema.UpdateCinema)
		g.GET("/details", v1.Group.Cinema.GetCinemaDetails)
	}
}

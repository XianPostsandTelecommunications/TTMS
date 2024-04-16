/**
 * @Author: lenovo
 * @Description:
 * @File:  people_movie
 * @Version: 1.0.0
 * @Date: 2023/06/01 22:17
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type pm struct{}

func (pm) Init(r *gin.RouterGroup) {
	g := r.Group("/peopleMovie", middleware.Auth())
	{
		g.POST("", v1.Group.Pm.UserMovieAction)
	}
}

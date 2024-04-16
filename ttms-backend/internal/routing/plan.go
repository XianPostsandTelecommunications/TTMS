/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/05 8:13
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type plan struct{}

func (plan) Init(r *gin.RouterGroup) {
	g := r.Group("/plan", middleware.Auth(), middleware.AuthManager())
	{
		g.POST("/create", v1.Group.Plan.CreatePlan)
		g.DELETE("/delete", v1.Group.Plan.DelPlan)
	}

	p := r.Group("/plan", middleware.Auth())
	{
		p.POST("/byMovieIDAndPeriod", v1.Group.Plan.GetPlansByMovieIDAndPeriod)
		p.GET("/list/:page", v1.Group.Plan.GetPlans)
	}
}

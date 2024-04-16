/**
 * @Author: lenovo
 * @Description:
 * @File:  tag
 * @Version: 1.0.0
 * @Date: 2023/05/31 21:53
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type tag struct{}

func (tag) Init(r *gin.RouterGroup) {
	{
		g := r.Group("/tags", middleware.Auth(), middleware.AuthManager())
		{
			g.POST("/addTag", v1.Group.Tag.AddTagForMovie)
			g.GET("/getTags", v1.Group.Tag.GetTagsFromMovie)
		}

	}
}

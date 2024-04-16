/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/05/31 11:26
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type file struct{}

func (file) Init(r *gin.RouterGroup) {
	f := r.Group("/file", middleware.Auth())
	{
		f.POST("/upload", v1.Group.File.Upload)
	}
}

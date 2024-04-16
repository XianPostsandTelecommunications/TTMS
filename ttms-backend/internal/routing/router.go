/**
 * @Author: lenovo
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:12
 */

package routing

import (
	gs "github.com/swaggo/gin-swagger"
	"mognolia/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	_ "mognolia/internal/docs"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())
	r.LoadHTMLFiles("./templates/index.html") // 加载html
	r.Static("/assets", "./assets")           // 加载静态文件
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})
	root := r.Group("/api/v1")
	{
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
		r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
		Group.User.Init(root)
		Group.Email.Init(root)
		Group.Movie.Init(root)
		Group.File.Init(root)
		Group.Tag.Init(root)
		Group.Pm.Init(root)
		Group.Seat.Init(root)
		Group.Cinema.Init(root)
		Group.Plan.Init(root)
		Group.Ticket.Init(root)
	}
	return r
}

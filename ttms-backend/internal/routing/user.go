/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:21
 */

package routing

import (
	"github.com/gin-gonic/gin"
	v1 "mognolia/internal/api/v1"
	"mognolia/internal/middleware"
)

type user struct{}

func (u *user) Init(r *gin.RouterGroup) {

	r.POST("/register", v1.Group.User.Register)
	r.POST("/login", v1.Group.User.Login)
	r.GET("/isRePeat", v1.Group.User.IsRePeat)
	r.POST("/refreshToken", v1.Group.User.RefreshToken)
	g := r.Group("/user", middleware.Auth())
	{

		g.POST("/findUser", v1.Group.User.FindUser)
		g.PUT("/modifyAvatar", v1.Group.User.ModifyAvatar)
		g.PUT("/modifyPassword", v1.Group.User.ModifyPassword)
		g.PUT("/modifyEmail", v1.Group.User.ModifyEmail)
		g.PUT("/updateInfo", v1.Group.User.UpdateUserInfo)

	}
	manager := r.Group("/manager", middleware.Auth(), middleware.AuthManager())
	{
		manager.GET("/list", v1.Group.User.List)
		manager.POST("/register", v1.Group.User.Register)
		manager.DELETE("/:UserID", v1.Group.User.DeleteUser)
		manager.POST("/createManager", v1.Group.User.CreateManager)
	}
}

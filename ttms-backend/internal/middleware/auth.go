/**
 * @Author: lenovo
 * @Description:
 * @File:  auth
 * @Version: 1.0.0
 * @Date: 2023/05/29 22:20
 */

package middleware

import (
	"mognolia/internal/dao"
	"mognolia/internal/global"
	"mognolia/internal/model"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app"
	"mognolia/internal/pkg/app/errcode"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) (*model.Content, errcode.Err) {
	v, ok := ctx.Get(global.Settings.Token.AuthKey)
	if !ok {
		return nil, errcode.ErrUnauthorizedAuthNotExist
	}
	return v.(*model.Content), nil
}
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		tokenString := ctx.GetHeader(global.Settings.Token.AuthType)
		payLoad, err := global.Maker.VerifyToken(tokenString)
		if err != nil {
			rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
			ctx.Abort()
			return
		}
		t := &model.Content{}
		if err := t.UnMarshal(payLoad.Content); err != nil {
			rly.Reply(errcode.ErrParamsNotValid)
			ctx.Abort()
			return
		}
		if result := dao.Group.DB.Model(&automigrate.User{}).Where("id =?", t.ID).Find(&automigrate.User{}); result.RowsAffected == 0 {
			rly.Reply(myerr.UserNotExists)
			ctx.Abort()
			return
		}
		global.Logger.Info(global.Settings.Token.AuthKey)
		ctx.Set(global.Settings.Token.AuthKey, t)
		ctx.Next()
	}
}

func AuthManager() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rly := app.NewResponse(ctx)
		context, err := GetContext(ctx)
		if err != nil {
			rly.Reply(err)
			ctx.Abort()
			global.Logger.Info("GetContext failed ")
			return
		}
		id, role := context.ID, context.Role
		if role != automigrate.Admin {
			rly.Reply(myerr.AuthNotEnough)
			ctx.Abort()
			return
		}
		u := automigrate.User{}
		if result := dao.Group.DB.Model(&automigrate.User{}).Where("id =?", id).Find(&u); result.RowsAffected == 0 {
			rly.Reply(myerr.UserNotExists)
			ctx.Abort()
			return
		} else {
			if u.Role != automigrate.Admin {
				rly.Reply(myerr.AuthNotEnough)
				ctx.Abort()
				return
			}
		}
		ctx.Next()
	}
}

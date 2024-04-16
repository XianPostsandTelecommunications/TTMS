/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:42
 */

package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mognolia/internal/api/base"
	"mognolia/internal/global"
	"mognolia/internal/logic"
	"mognolia/internal/middleware"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app"
	"mognolia/internal/pkg/app/errcode"
	"mognolia/internal/pkg/utils"
)

type user struct{}

// Register
// @Tags      user
// @Summary   注册
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data           body      request.RegisterParam  true  "注册"
// @Success   200            {object}  common.State{data=reply.RegisterRly}  "1001:参数有误 1003:系统错误 20001:用户已存在 "
// @Router    /api/v1/register [post]
func (u *user) Register(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.RegisterParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		zap.S().Infof("should bind failed err:%v", err)
		return
	}
	rsp, err := logic.Group.User.Register(ctx, param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rsp)
}

// Login
// @Tags      user
// @Summary   登录
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data           body      request.LoginParam  true  "登录"
// @Success   200            {object}  common.State{data=reply.LoginRly}  "1001:参数有误 1003:系统错误 30001:验证码失效或者有误 20003：密码不能为空"
// @Router    /api/v1/login [post]
func (u *user) Login(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.LoginParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	rep, err := logic.Group.User.Login(param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rep)
}

// RefreshToken
// @Tags      user
// @Summary   刷新token
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data           body      request.RefreshTokenParam  true  "刷新"
// @Success   200            {object}  common.State{data=reply.RefreshRly}  "1001:参数有误 1003:系统错误 2001:鉴权失败 30003:refresh token失效"
// @Router    /api/v1/refreshToken [post]
func (u *user) RefreshToken(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.RefreshTokenParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	rep, err := logic.Group.User.Refresh(param)
	if err != nil {
		global.Logger.Info("RefreshToken generated error")
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rep)
}

// FindUser
// @Tags      user
// @Summary   根据用户名查找用户
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           body      request.FindParam  true  "刷新"
// @Success   200            {object}  common.State{data=reply.UserInfo}  "1001:参数有误 1003:系统错误 20002:用户不存在"
// @Router    /api/v1/user/findUser [post]
func (u *user) FindUser(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.FindParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		global.Logger.Error("ctx.ShouldBindJSON failed")
		base.HandleValidatorError(ctx, err)
		return
	}
	rep, err := logic.Group.User.FindUserInfo(param.UserName)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rep)
}

// IsRePeat
// @Tags      user
// @Summary   判断用户名是否存在
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data           query      request.IsRePeat  true  "是否重复"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误 2001:鉴权失败 20001:用户已存在 20002:用户不存在"
// @Router    /api/v1/isRePeat [get]
func (u *user) IsRePeat(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var params request.IsRePeat
	if err := ctx.ShouldBindQuery(&params); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	uid, err := logic.Group.User.CheckIsRePeat(params.Username)
	if err != nil {
		zap.S().Infof("logic.Group.User.FindUserInfo failed,err:%v", err)
		rly.Reply(err)
		return
	}
	rly.Reply(nil, uid)
}

// List
// @Tags      FuncMgr
// @Summary   获取用户列表
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     page           path      string                 true  "页码"
// @Success   200            {object}  common.State{data=reply.UserList}  "1001:参数有误 1003:系统错误 2001:鉴权失败 30004:权限不够 30005:无相关记录"
// @Router    /api/v1/FuncMgr/list [get]
func (u *user) List(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	content, err := middleware.GetContext(ctx)
	if err != nil {
		rly.Reply(errcode.ErrUnauthorizedAuthNotExist)
		return
	}
	if content.Role != automigrate.Admin {
		rly.Reply(myerr.AuthNotEnough)
		return
	}
	//i := ctx.Query("page")
	//istr := utils.StringToIDMust(i)
	rsp, err := logic.Group.User.GetList()
	if err != nil {
		rly.Reply(err)
		zap.S().Infof("logic group user list failed: %v", err)
		return
	}
	rly.Reply(nil, rsp)
}

// ModifyPassword
// @Tags      user
// @Summary   用户修改密码
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query      request.ModifyPassword true  "修改密码"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误  20002:用户不存在 30001:验证码有误"
// @Router    /api/v1/user/modifyPassword [put]
func (u *user) ModifyPassword(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.ModifyPassword
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.User.ModifyPassword(ctx, param.EmailCode, param.NewPassword); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// ModifyAvatar
// @Tags      user
// @Summary   用户修改头像
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query     request.ModifyAvatar  true  "修改头像"
// @Success   200            {object}  common.State{data=request.ModifyAvatar}  "1001:参数有误 1003:系统错误  20002:用户不存在"
// @Router    /api/v1/user/modifyAvatar [put]
func (u *user) ModifyAvatar(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.ModifyAvatar
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		fmt.Println(err)
		return
	}
	if err := logic.Group.User.ModifyAvatar(ctx, param.NewAvatar); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// ModifyEmail
// @Tags      user
// @Summary   用户修改头像
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query     request.ModifyEmail  true  "修改邮箱"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误  20002:用户不存在 30001:验证码有误"
// @Router    /api/v1/user/modifyEmail [put]
func (u *user) ModifyEmail(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.ModifyEmail
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.User.ModifyEmail(ctx, param.EmailCode, param.NewEmail); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

// UpdateUserInfo
// @Tags      user
// @Summary   用户修改性别,个性签名
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query      request.UpdateInfo  true  "修改个性签名,性别"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误  20002:用户不存在"
// @Router    /api/v1/user/updateInfo [put]
func (u *user) UpdateUserInfo(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.UpdateInfo
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.User.UpdateUserInfo(ctx, param); err != nil {
		rly.Reply(err)
	}
	rly.Reply(nil)
}

// DeleteUser
// @Tags      FuncMgr
// @Summary   管理员删除用户
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Param     data           query      request.DelUser  true  "删除用户"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误  20002:用户不存在"
// @Router    /api/v1/user/update [put]
func (u *user) DeleteUser(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	iStr := ctx.Param("UserID")
	i := utils.StringToIDMust(iStr)
	if err := logic.Group.User.DeleteUser(uint(i)); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

func (u *user) CreateManager(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.CreateManager
	if err := ctx.ShouldBindJSON(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if err := logic.Group.User.CreateManager(param.Username); err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

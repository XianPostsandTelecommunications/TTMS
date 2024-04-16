/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/05/29 12:11
 */

package v1

import (
	"fmt"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/pkg/app"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type email struct{}

// SendEmail
// @Tags      email
// @Summary   发邮件
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     data           query     request.SendEmailParams  true  "是否重复"
// @Success   200            {object}  common.State{}  "1001:参数有误 1003:系统错误 2001:鉴权失败 20001:用户已存在 30002:发送次数过多"
// @Router    /api/v1/email/send [post]
func (e *email) SendEmail(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.SendEmailParams
	if err := ctx.ShouldBindJSON(&param); err != nil {
		fmt.Println(param.Email)
		base.HandleValidatorError(ctx, err)
		zap.S().Info("ctx.ShouldBindJSON failed", zap.Any("err", err))
		return
	}
	fmt.Println(param.Email)
	if err := logic.Group.Email.SendEmail(param.Email); err != nil {
		zap.S().Info("logic.Group.Email.SendEmail ")
		rly.Reply(err)
		return
	}
	rly.Reply(nil)
}

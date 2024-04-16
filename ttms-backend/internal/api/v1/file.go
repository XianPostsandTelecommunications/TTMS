/**
 * @Author: lenovo
 * @Description:
 * @File:  file
 * @Version: 1.0.0
 * @Date: 2023/05/31 11:04
 */

package v1

import (
	"github.com/gin-gonic/gin"
	"mognolia/internal/api/base"
	"mognolia/internal/logic"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app"
)

type file struct{}

// Upload
// @Tags      file
// @Summary   上传文件
// @Security  BasicAuth
// @accept    application/json
// @Produce   application/json
// @Param     x_token  header    string                 true  "x_token 用户令牌"
// @Success   200            {object}  common.State{reply.FileRly}  "1001:参数有误 1003:系统错误 2001:鉴权失败 "
// @Router    /api/v1/file/send [post]
func (file) Upload(ctx *gin.Context) {
	rly := app.NewResponse(ctx)
	var param request.FileParam
	if err := ctx.ShouldBind(&param); err != nil {
		base.HandleValidatorError(ctx, err)
		return
	}
	if param.File == nil {
		rly.Reply(myerr.FileIsEmpty)
		return
	}
	rep, err := logic.Group.File.UploadFile(param)
	if err != nil {
		rly.Reply(err)
		return
	}
	rly.Reply(nil, rep)
}

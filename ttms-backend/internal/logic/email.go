/**
 * @Author: lenovo
 * @Description:
 * @File:  email
 * @Version: 1.0.0
 * @Date: 2023/05/29 12:19
 */

package logic

import (
	"mognolia/internal/global"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
	email2 "mognolia/internal/work/email"

	"go.uber.org/zap"
)

type email struct{}

func (e *email) SendEmail(email string) errcode.Err {
	send := email2.NewSendCodeTask(email)
	if ok := email2.CheckEmailBeMask(email); ok {
		return myerr.SendTooMany
	}
	global.Worker.SendTask(send.SendTask()) //异步发送消息
	go func() {
		result := send.GetChanResult()
		if result.Error != nil {
			switch result.Error {
			case email2.ErrSendTooMany:
				zap.S().Infof("发送过于频繁")
			default:
				zap.S().Infof("failed,error: %v", result.Error)
				zap.S().Info("其他错误")
			}

		}
	}()
	return nil
}

/**
 * @Author: lenovo
 * @Description:
 * @File:  myerr
 * @Version: 1.0.0
 * @Date: 2023/05/29 11:02
 */

package myerr

import "mognolia/internal/pkg/app/errcode"

var (
	UserAlreadyExists       = errcode.NewErr(20001, "用户已经存在了")
	UserNotExists           = errcode.NewErr(20002, "用户不存在")
	PasswordCanNotBeEmpty   = errcode.NewErr(20003, "密码不能为空")
	InvalidPassword         = errcode.NewErr(20004, "密码错误")
	CodeInvalid             = errcode.NewErr(30001, "验证码失效或者是不正确")
	SendTooMany             = errcode.NewErr(30002, "发送次数过多")
	RefreshTokenInvalid     = errcode.NewErr(30003, "refreshToken 失效")
	AuthNotEnough           = errcode.NewErr(30004, "非管理员，权限不够")
	NoRecords               = errcode.NewErr(30005, "无相关记录")
	FileIsEmpty             = errcode.NewErr(30006, "文件为空")
	ErrMovieAlreadyExists   = errcode.NewErr(30007, "电影已经存在了")
	ErrHasFavoriteThisMovie = errcode.NewErr(30008, "用户已经关注了这部电影")
	ErrLockTickets          = errcode.NewErr(30009, "票已经被锁定了")
	ErrCreateOrder          = errcode.NewErr(30010, "生成订单失败")
	NotLockTicket           = errcode.NewErr(30011, "用户没有锁票不能购买")
)

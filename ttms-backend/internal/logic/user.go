/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:46
 */

package logic

import (
	"errors"
	"fmt"
	"github.com/0RAJA/Rutils/pkg/password"
	"github.com/gin-gonic/gin"
	"mognolia/internal/dao/mysql/query"
	"mognolia/internal/global"
	"mognolia/internal/middleware"
	"mognolia/internal/model"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/reply"
	"mognolia/internal/model/request"
	"mognolia/internal/myerr"
	"mognolia/internal/pkg/app/errcode"
	password2 "mognolia/internal/pkg/password"
	"mognolia/internal/pkg/token"
	workemail "mognolia/internal/work/email"
	"time"

	"gorm.io/gorm"
)

type user struct{}
type Result struct {
	Token   string
	Payload *token.Payload
	Error   error
}

func CreateToken(resultChan chan<- Result, userID uint, role automigrate.Roler, expireTime time.Duration) func() {
	return func() {
		defer close(resultChan)
		content, _ := model.NewContent(userID, role).Marshal()
		tokenString, payLoad, err := global.Maker.CreateToken(content, expireTime)
		resultChan <- Result{Token: tokenString, Payload: payLoad, Error: err}
	}
}
func (u *user) Register(ctx *gin.Context, params request.RegisterParam) (*reply.RegisterRly, errcode.Err) {
	//先判断一下是否注册过了
	content, err := middleware.GetContext(ctx)
	var role automigrate.Roler
	if errors.Is(err, errcode.ErrUnauthorizedAuthNotExist) {
		role = automigrate.Vistor
	} else {
		role = content.Role
	}

	q := query.NewUser()
	_, err1 := q.FindUserByEmail(params.Email)
	if err1 != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			//执行注册操作
			if ok := workemail.CheckEmailAndCodeValid(params.Email, params.EmailCode); !ok {
				return nil, myerr.CodeInvalid
			}
			hashPassword, err := password.HashPassword(params.Password)
			if err != nil {
				return nil, errcode.ErrServer
			}

			if params.Avatar == "" {
				params.Avatar = global.Settings.Rule.DefaultAccountAvatar
			}
			if params.Gender == "" {
				params.Gender = string(automigrate.UnKnown)
			}
			userInfo, err := q.Register(params, hashPassword, role)
			if err != nil {
				return nil, errcode.ErrServer
			}
			return &reply.RegisterRly{UserID: userInfo.ID}, nil
		}
		return nil, errcode.ErrServer
	}
	return nil, myerr.UserAlreadyExists
}

func (u *user) Login(params request.LoginParam) (*reply.LoginRly, errcode.Err) {
	q := query.NewUser()
	userInfo, err := q.FindUserByEmail(params.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.UserNotExists
		}
		return nil, errcode.ErrServer
	}
	switch params.LoginType {
	case 2:
		if len(params.Password) == 0 {
			return nil, myerr.PasswordCanNotBeEmpty
		}
		if err := password.CheckPassword(params.Password, userInfo.Password); err != nil {
			return nil, myerr.InvalidPassword
		}
	case 1:
		if ok := workemail.CheckEmailAndCodeValid(params.Email, params.EmailCode); !ok {
			return nil, myerr.CodeInvalid
		}
	default:
		return nil, errcode.ErrParamsNotValid
	}
	//可以生成AccessToken了
	accessToken := make(chan Result, 1)
	refreshToken := make(chan Result, 1)
	fmt.Println(global.Settings.Token.AccessTokenTime)
	global.Worker.SendTask(CreateToken(accessToken, userInfo.ID, userInfo.Role, global.Settings.Token.AccessTokenTime))
	accessResult := <-accessToken
	if accessResult.Error != nil {
		return nil, errcode.ErrServer
	}
	global.Worker.SendTask(CreateToken(refreshToken, userInfo.ID, userInfo.Role, global.Settings.Token.RefreshTokenTime))

	refreshResult := <-refreshToken
	if refreshResult.Error != nil {
		return nil, errcode.ErrServer
	}
	return &reply.LoginRly{
		UserID:       userInfo.ID,
		AccessToken:  accessResult.Token,
		RefreshToken: refreshResult.Token,
		PayLoad:      accessResult.Payload,
		UserInfo: reply.UserInfoReply{
			UserID:    userInfo.ID,
			AvatarURL: userInfo.Avatar,
			Role:      string(userInfo.Role),
			UserName:  userInfo.UserName,
			Email:     userInfo.Email,
			Signature: userInfo.Signature,
		},
	}, nil
}

func (u *user) Refresh(param request.RefreshTokenParam) (*reply.RefreshRly, errcode.Err) {
	accPayLoad, err := global.Maker.VerifyToken(param.AccessToken)
	if err != nil {
		return nil, errcode.ErrUnauthorizedAuthNotExist
	}
	t := model.Content{}
	if err := t.UnMarshal(accPayLoad.Content); err != nil {
		return nil, errcode.ErrServer
	}
	RefpayLoad, err := global.Maker.VerifyToken(param.RefreshToken)
	if err != nil {
		return nil, errcode.ErrUnauthorizedAuthNotExist
	}
	if RefpayLoad.ExpiredAt.Before(time.Now()) {
		return nil, myerr.RefreshTokenInvalid
	}
	accessResult := make(chan Result, 1)
	global.Worker.SendTask(CreateToken(accessResult, t.ID, t.Role, global.Settings.Token.AccessTokenTime))
	r := <-accessResult
	if r.Error != nil {
		return nil, errcode.ErrServer
	}
	return &reply.RefreshRly{AccessToken: r.Token}, nil
}

func (u *user) FindUserInfo(username string) (*reply.UserInfo, errcode.Err) {
	q := query.NewUser()
	userInfo, err := q.FindUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.UserNotExists
		}
		return nil, errcode.ErrServer
	}
	return &reply.UserInfo{
		Username:  userInfo.UserName,
		Signature: userInfo.Signature,
		Gender:    string(userInfo.Gender),
		Avatar:    userInfo.Avatar,
		Role:      string(userInfo.Role),
	}, nil
}

func (u *user) CheckIsRePeat(username string) (uint, errcode.Err) {
	q := query.NewUser()
	uInfo, err := q.FindUserByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, myerr.UserNotExists
		} else {
			return 0, errcode.ErrServer
		}
	}
	return uInfo.ID, myerr.UserAlreadyExists
}

func (u *user) GetList() (*reply.UserList, errcode.Err) {
	var result reply.UserList
	q := query.NewUser()
	allInfos, err := q.GetAllUser()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, myerr.NoRecords
		} else {
			return nil, errcode.ErrServer
		}
	}
	result.Total = len(allInfos)
	for _, userInfo := range allInfos {
		result.UserInfos = append(result.UserInfos, &reply.UserInfo{
			UserID:    userInfo.ID,
			Username:  userInfo.UserName,
			Signature: userInfo.Signature,
			Gender:    string(userInfo.Gender),
			Avatar:    userInfo.Avatar,
			Role:      string(userInfo.Role),
		})
	}
	return &result, nil
}

func (u *user) ModifyPassword(ctx *gin.Context, emailCode, password string) errcode.Err {
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return err
	}
	uid := content.ID
	q := query.NewUser()
	uInfo, err1 := q.GetUserInfoByID(uid)
	if err != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			return myerr.UserNotExists
		}
		return errcode.ErrServer
	}
	if ok := workemail.CheckEmailAndCodeValid(uInfo.Email, emailCode); !ok {
		return myerr.CodeInvalid
	}
	hashPassord, _ := password2.HashPassword(password)
	if err := q.ModifyPassword(uid, hashPassord); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (u *user) ModifyAvatar(ctx *gin.Context, avatar string) errcode.Err {
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return err
	}
	uid := content.ID
	q := query.NewUser()
	_, err1 := q.GetUserInfoByID(uid)
	if err != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			return myerr.UserNotExists
		}
		return errcode.ErrServer
	}
	if err := q.ModifyAvatar(uid, avatar); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (u *user) ModifyEmail(ctx *gin.Context, emailCode, email string) errcode.Err {
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return err
	}
	uid := content.ID
	q := query.NewUser()
	uInfo, err1 := q.GetUserInfoByID(uid)
	if err != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			return myerr.UserNotExists
		}
		return errcode.ErrServer
	}
	if ok := workemail.CheckEmailAndCodeValid(uInfo.Email, emailCode); !ok {
		return myerr.CodeInvalid
	}
	if err := q.ModifyEmail(uid, email); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (u *user) UpdateUserInfo(ctx *gin.Context, param request.UpdateInfo) errcode.Err {
	content, err := middleware.GetContext(ctx)
	if err != nil {
		return err
	}
	uid := content.ID
	q := query.NewUser()
	_, err1 := q.GetUserInfoByID(uid)
	if err != nil {
		if errors.Is(err1, gorm.ErrRecordNotFound) {
			return myerr.UserNotExists
		}
		return errcode.ErrServer
	}
	if err := q.UpdateUserInfo(uid, param.Gender, param.Signature); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (u *user) DeleteUser(uid uint) errcode.Err {
	q := query.NewUser()
	if err := q.DelUser(uid); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

func (user) CreateManager(username string) errcode.Err {
	q := query.NewUser()
	userInfo, err := q.FindUserByUsername(username)
	if err != nil || userInfo.ID == 0 {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return myerr.NoRecords
		}
		return errcode.ErrServer
	}
	if err := q.UpdateUserRole(username); err != nil {
		return errcode.ErrServer.WithDetails(err.Error())
	}
	return nil
}

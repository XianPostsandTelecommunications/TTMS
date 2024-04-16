/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 9:31
 */

package request

type RegisterParam struct {
	Email     string `json:"Email" binding:"required"`                     //邮箱
	UserName  string `json:"UserName" binding:"required"`                  //用户名
	Password  string `json:"Password" binding:"required,gte=3,lte=12"`     //密码长度介于3和12之间
	EmailCode string `json:"EmailCode" binding:"required,gte=6,lte=6"`     //邮箱验证码，长度为6
	Signature string `json:"Signature" `                                   //个性签名
	Gender    string `json:"Gender" binding:"omitempty,oneof=male female"` //性别
	Avatar    string `json:"Avatar"`                                       //头像
}

type LoginParam struct {
	Email     string `json:"Email" binding:"required" `     //
	Password  string `json:"Password"`                      //密码
	EmailCode string `json:"EmailCode"`                     //
	LoginType int    `json:"LoginType" binding:"oneof=1 2"` //1表示验证码登录，2表示密码登录
}

type RefreshTokenParam struct {
	AccessToken  string `json:"AccessToken" binding:"required"`
	RefreshToken string `json:"RefreshToken" binding:"required"`
}

type FindParam struct {
	UserName string `json:"UserName" binding:"required"`
}

type IsRePeat struct {
	Username string `json:"Username" binding:"required"`
}

type ModifyPassword struct {
	EmailCode   string `json:"EmailCode" binding:"required,gte=6,lte=6"`
	NewPassword string `json:"NewPassword" binding:"required"`
}

type ModifyEmail struct {
	EmailCode string `json:"EmailCode" binding:"required,gte=6,lte=6"`
	NewEmail  string `json:"NewEmail" binding:"required"`
}

type ModifyAvatar struct {
	NewAvatar string `json:"newAvatar" binding:"required"`
}

type UpdateInfo struct {
	Signature string `json:"signature"`
	Gender    string `json:"gender"`
}

type DelUser struct {
	UserID uint `json:"userID" binding:"required"`
}

type GetList struct {
	Page int `json:"page"`
}

type CreateManager struct {
	Username string `json:"Username" binding:"required"`
}

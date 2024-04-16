/**
 * @Author: lenovo
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2023/05/29 11:04
 */

package query

import (
	"mognolia/internal/dao"
	"mognolia/internal/model/automigrate"
	"mognolia/internal/model/request"

	"gorm.io/gorm"
)

type user struct{}

func NewUser() *user {
	return &user{}
}
func (u *user) FindUserByEmail(email string) (*automigrate.User, error) {
	var user automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("email =?", email).Find(&user); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (u *user) FindUserByUsername(username string) (*automigrate.User, error) {
	var user automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("user_name =?", username).Find(&user); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &user, nil
}

func (u *user) Register(us request.RegisterParam, hashPassword string, role automigrate.Roler) (*automigrate.User, error) {
	user := &automigrate.User{
		UserName:  us.UserName,
		Email:     us.Email,
		Password:  hashPassword,
		Signature: us.Signature,
		Avatar:    us.Avatar,
		Gender:    automigrate.Gend(us.Gender),
		Role:      role,
	}
	if result := dao.Group.DB.Model(&automigrate.User{}).Create(user); result.RowsAffected == 0 {
		return nil, result.Error
	}
	return user, nil
}

func (u *user) GetList() ([]automigrate.User, error) {
	var UserInfos []automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Find(&UserInfos); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return UserInfos, nil
}

func (u *user) GetAllUser() ([]automigrate.User, error) {
	var UserInfos []automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Find(&UserInfos); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return UserInfos, nil
}

func (u *user) GetUserInfoByID(uid uint) (*automigrate.User, error) {
	var us automigrate.User
	if result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", uid).Find(&us); result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &us, nil
}

func (u *user) ModifyPassword(uid uint, hashPassword string) error {
	result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", uid).Update("password", hashPassword)
	return result.Error
}
func (u *user) ModifyEmail(uid uint, newEmail string) error {
	result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", uid).Update("email", newEmail)
	return result.Error
}

func (u *user) ModifyAvatar(uid uint, avatar string) error {
	result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", uid).Update("avatar", avatar)
	return result.Error
}

func (u *user) UpdateUserInfo(uid uint, Gender, Signature string) error {
	result := dao.Group.DB.Model(&automigrate.User{}).Where("id = ?", uid).Updates(automigrate.User{Gender: automigrate.Gend(Gender), Signature: Signature})
	return result.Error
}

func (u user) DelUser(uid uint) error {
	if result := dao.Group.DB.Model(&automigrate.User{}).Delete(&automigrate.User{}, uid); result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (user) UpdateUserRole(userName string) error {
	result := dao.Group.DB.Model(&automigrate.User{}).Where("user_name =?", userName).Updates(&automigrate.User{Role: automigrate.Admin})
	return result.Error
}

package services

import (
	"studygo/binders"
	"studygo/constant"
	"studygo/models"
	"studygo/tools"
)

// UserService 相关业务
type UserService struct{}

// CreateUser 创建用户
func (obj *UserService) CreateUser(user *binders.User) {
	salt := tools.GetRandomString(6)
	hash := tools.PasswordHash(&user.Password, &salt)

	userModel := models.User{Username: user.Username, Age: user.Age}
	userModel.Password = hash
	userModel.Salt = salt

	models.DB.Create(&userModel)
}

// VerifyUser 验证用户
func (obj *UserService) VerifyUser(userBind *binders.User, check byte) (models.User, bool) {
	var (
		userEntity models.User
	)
	models.DB.Where("username = ?", userBind.Username).First(&userEntity)

	if constant.UserVerifyCheckPwd == check&constant.UserVerifyCheckPwd {
		if !tools.PasswordVerify(&userBind.Password, &userEntity.Salt, &userEntity.Password) {
			return userEntity, false
		}
	}

	if constant.UserVerifyCheckStat == check&constant.UserVerifyCheckStat {
		if userEntity.Status != 1 {
			return userEntity, false
		}
	}
	return userEntity, true
}

package services

import (
	"studygo/binders"
	"studygo/models"
)

// UserService 相关业务
type UserService struct{}

// CreateUser 创建用户
func (obj *UserService) CreateUser(user *binders.User) {
	user.Age = 18
	userModel := models.User{Username: user.Username, Password: user.Password}
	models.DB.Create(&userModel)
}

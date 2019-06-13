package api

import (
	"fmt"
	"studygo/binders"
	"studygo/services"

	"github.com/gin-gonic/gin"
)

// Signup 注册
func Signup(c *gin.Context) {
	var userService *services.UserService
	var userBind binders.User

	if c.ShouldBind(&userBind) == nil {
		fmt.Println(userBind)
		userService.CreateUser(&userBind)
		c.JSON(200, gin.H{
			"status": 200,
			"msg":    "成功",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": 400,
		"msg":    "参数错误",
	})

}

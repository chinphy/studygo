package api

import (
	"fmt"
	"studygo/binders"
	"studygo/constant"
	"studygo/services"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// Signin 登录
func Signin(c *gin.Context) {
	var (
		userService *services.UserService
		userBind    binders.User
		tokenKey    = []byte("This is secret!")
	)

	if c.ShouldBind(&userBind) == nil {
		fmt.Println("userBind:", userBind)
		userEntity, ok := userService.VerifyUser(&userBind, constant.UserVerifyCheckPwd|constant.UserVerifyCheckStat)
		if ok {
			claims := &jwt.StandardClaims{
				NotBefore: int64(time.Now().Unix()),
				ExpiresAt: int64(time.Now().Unix() + 1000),
				Issuer:    userEntity.Username,
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenStr, err := token.SignedString(tokenKey)

			if err == nil {
				c.JSON(200, gin.H{
					"status":       200,
					"access_token": tokenStr,
				})
				return
			}
		}
	}

	c.JSON(200, gin.H{
		"status": 500,
	})
}

package services

import (
	"studygo/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Sign 登录注册签到
type Sign struct{}

var (
	tokenKey = []byte("This is secret!")
)

// Signup 注册
func (t *Sign) Signup(c *gin.Context) {

}

// Signin 登录
func (t *Sign) Signin(c *gin.Context) {
	var user models.User
	models.DB.Where("username = ?", c.PostForm("username")).First(&user)

	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    c.PostForm("username"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(tokenKey)

	if err != nil {
		c.JSON(200, gin.H{
			"status": 500,
		})
	} else {
		c.JSON(200, gin.H{
			"status":       200,
			"access_token": tokenStr,
		})
	}
}

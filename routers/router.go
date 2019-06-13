package routers

import (
	"studygo/api"
	"studygo/services"

	"github.com/gin-gonic/gin"
)

//Router ...
func Router() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		v1.GET("/someJSON", func(c *gin.Context) {
			data := map[string]interface{}{
				"lang": "zh",
			}
			c.JSON(200, data)
		})

		var c *services.Sign
		v1.POST("/signin", c.Signin)
		v1.POST("/signup", api.Signup)
	}

	return r
}

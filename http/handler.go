package http

import (
	firebase "firebase-authenticator/driver"

	"github.com/gin-gonic/gin"
)

//User is the struct to hold the received token value
type User struct {
	Token string `json:"token" binding:"required"`
}

//Run inits and starts all services
func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/token", generateToken)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func generateToken(c *gin.Context) {
	key := c.GetHeader("X-API-Key")
	if key != "key-value" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})

		return
	}

	var u User
	c.BindJSON(&u)
	token, err := firebase.CreateCustomToken(u.Token)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"token":   token,
		"message": "success",
	})
}

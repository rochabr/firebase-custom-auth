package http

import (
	"errors"
	"strings"

	auth "firebase-custom-auth/http/security"

	firebase "github.com/rochabr/firebase-custom-auth/driver"

	"github.com/gin-gonic/gin"
)

//Run inits and starts all services
func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/token", generateToken)
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func generateToken(c *gin.Context) {
	key := c.GetHeader("Authorization")
	token, err := parseBearer(key)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})

		return
	}

	customToken, err := firebase.CreateCustomToken(token)
	if err != nil || token == "" {
		c.JSON(500, gin.H{
			"message": err.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"token":   customToken,
		"message": "success",
	})
}

func parseBearer(key string) (string, error) {
	bearer := strings.Split(key, "Bearer ")

	if len(bearer) != 2 {
		return "", errors.New("Invalid bearer token")
	}

	token, err := auth.Validate(bearer[1])
	return token, err
}

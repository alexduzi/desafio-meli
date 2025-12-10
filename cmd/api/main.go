package main

import (
	"net/http"
	"project/internal/infra/database"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from GO!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	_, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	r := setupRouter()

	r.Run()
}

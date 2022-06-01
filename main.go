package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getPort() string {
	p := os.Getenv("APP_PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		hostname, _ := os.Hostname()
		ctx.JSON(http.StatusOK, "Pod name is "+hostname)
	})

	router.Run(getPort())
}

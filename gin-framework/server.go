package main

import (
	"gin-framework/controller"
	"gin-framework/middlewares"
	"gin-framework/service"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func setupLogOutput() {
	if f, err := os.Create("gin.log"); err == nil {
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	}
}

func main() {
	setupLogOutput()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}

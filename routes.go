package main

import (
	"github.com/gin-gonic/gin"
	"govue.demo/go_web_0/controller"
	"govue.demo/go_web_0/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/reg", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	return r
}

package router

import (
	"github.com/gin-gonic/gin"

	"github.com/forg002-ctrl/managment_system/pkg/hadlers/account"
	"github.com/forg002-ctrl/managment_system/pkg/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default();

	router.Use(middleware.LoggerMiddleware())

	api := router.Group("/api/v1")
	{
		// api.POST("/account")
		// api.GET("/accounts")
		// api.GET("/account/:id")
		api.GET("account", handlers.CreateAccount)
	}

	return router
}
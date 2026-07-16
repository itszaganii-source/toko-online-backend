package routes

import (
	"toko-online-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine) {
	// Order routes
	router.POST("/orders", controllers.CreateOrder)
}

package routes

import (
	"toko-online-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupOrderRoutes(router *gin.Engine) {
	// Order routes
	router.GET("/orders", controllers.GetOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:id/status", controllers.UpdateOrderStatus)
}

package routes

import (
	"toko-online-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(router *gin.Engine) {
	// Product routes
	router.POST("/products", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProductByID)
}

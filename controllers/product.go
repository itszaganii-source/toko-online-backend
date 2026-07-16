package controllers

import (
	"net/http"
	"toko-online-backend/config"
	"toko-online-backend/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON request to product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save product to database
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product",
		})
		return
	}

	// Return success response with created product
	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"data":    product,
	})
}

func GetProducts(c *gin.Context) {
	var products []models.Product

	// Get all products from database
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve products",
		})
		return
	}

	// Return products in JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "Products retrieved successfully",
		"data":    products,
	})
}

func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	// Get product by ID from database
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found",
		})
		return
	}

	// Return product in JSON
	c.JSON(http.StatusOK, gin.H{
		"message": "Product retrieved successfully",
		"data":    product,
	})
}

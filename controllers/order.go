package controllers

import (
	"net/http"
	"toko-online-backend/config"
	"toko-online-backend/models"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	CustomerName string `json:"customer_name" binding:"required"`
	Items        []struct {
		ProductID uint `json:"product_id" binding:"required"`
		Quantity  int  `json:"quantity" binding:"required"`
	} `json:"items" binding:"required"`
}

type UpdateOrderStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Start transaction
	tx := config.DB.Begin()

	// Create order
	order := models.Order{
		CustomerName: req.CustomerName,
		TotalPrice:   0,
		Status:       "Pending",
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create order",
		})
		return
	}

	totalPrice := 0

	// Process each item
	for _, item := range req.Items {
		var product models.Product

		// Get product from database
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Product not found",
			})
			return
		}

		// Check stock availability
		if product.Stock < item.Quantity {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Stok produk " + product.Name + " tidak mencukupi",
			})
			return
		}

		// Reduce product stock
		product.Stock -= item.Quantity
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to update product stock",
			})
			return
		}

		// Calculate item price
		itemPrice := product.Price * item.Quantity
		totalPrice += itemPrice

		// Create order item
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create order item",
			})
			return
		}
	}

	// Update order total price
	order.TotalPrice = totalPrice
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update order total price",
		})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to commit transaction",
		})
		return
	}

	// Fetch complete order with items
	config.DB.Preload("OrderItems.Product").First(&order, order.ID)

	// Return success response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
		"data":    order,
	})
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	// Fetch all orders with their items
	if err := config.DB.Preload("OrderItems.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch orders",
		})
		return
	}

	// Return orders
	c.JSON(http.StatusOK, orders)
}

func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var req UpdateOrderStatusRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find order by ID
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Order not found",
		})
		return
	}

	// Update order status
	order.Status = req.Status
	if err := config.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update order status",
		})
		return
	}

	// Return updated order
	c.JSON(http.StatusOK, order)
}

package routes

import (
	"order-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Order Service Running",
		})
	})

	router.POST("/orders", handlers.CreateOrder)

	router.GET("/orders", handlers.GetOrders)
}

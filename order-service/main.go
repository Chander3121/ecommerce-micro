package main

import (
	"log"
	"order-service/config"
	"order-service/models"
	"order-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.Order{})

	routes.SetupRoutes(router)

	log.Println("Order Service running on port 3003")

	router.Run(":3003")
}

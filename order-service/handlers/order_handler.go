package handlers

import (
  "net/http"
  "order-service/clients"
  "order-service/config"
  "order-service/models"

  "github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
  token := c.GetHeader("Authorization")

  if token == "" {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error": "Missing token",
    })

    return
  }

  validateResponse, err := clients.ValidateToken(token)

  if err != nil || !validateResponse.Valid {
    c.JSON(http.StatusUnauthorized, gin.H{
      "error": "Invalid token",
    })

    return
  }

  var order models.Order

  if err := c.ShouldBindJSON(&order); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": err.Error(),
    })

    return
  }

  product, err := clients.GetProduct(
    int(order.ProductID),
  )

  if err != nil {
    c.JSON(http.StatusBadGateway, gin.H{
      "error": "Failed to fetch product",
    })

    return
  }

  if product.Stock < order.Quantity {
    c.JSON(http.StatusBadRequest, gin.H{
      "error": "Insufficient stock",
    })

    return
  }

  order.UserID = uint(validateResponse.User.ID)

  order.TotalPrice =
    product.Price * float64(order.Quantity)

  order.Status = "pending"

  config.DB.Create(&order)

  c.JSON(http.StatusCreated, gin.H{
    "order": order,
    "product": product,
  })
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	config.DB.Find(&orders)

	c.JSON(http.StatusOK, orders)
}

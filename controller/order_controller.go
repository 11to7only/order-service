package controller

import (
	"fmt"
	"net/http"
	"order-service/model"
	"order-service/service"

	"github.com/gin-gonic/gin"
	"strconv"
)

func Home(c *gin.Context) {
    fmt.Println("Service Home called...")
    c.JSON(http.StatusOK, gin.H{
        "message": "Order Service is running successfully!",
    })
}

func Health(c *gin.Context) {
    fmt.Println("Service Health called...")
    c.JSON(http.StatusOK, gin.H{
        "message": "Application is UP and Running",
    })
}


func GetOrders(c *gin.Context) {
	fmt.Println("Controller GetOrders called...")

	orders := service.GetAllOrders()

	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	fmt.Println("Controller CreateOrder called...")

	var order model.Order

	// Read JSON request body
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	service.AddOrder(order)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Order created successfully",
	})
}

func GetOrderByID(c *gin.Context) {
    fmt.Println("Controller GetOrderByID called...")
    idParam := c.Param("id")

    id, err := strconv.Atoi(idParam)
    	if err != nil {
    		c.JSON(http.StatusBadRequest, gin.H{
    			"error": "Invalid order ID",
    		})
    		return
    	}

    order := service.GetOrderByID(id)

    if order == nil {
        c.JSON(http.StatusNotFound, gin.H{
            "error": "Order not found",
        })
        return
    }
    c.JSON(http.StatusOK, order)
}
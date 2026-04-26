package routes

import (
	"order-service/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/", controller.Home)
	router.GET("/health", controller.Health)

	router.GET("/orders", controller.GetOrders)
	router.POST("/orders", controller.CreateOrder)
	router.GET("/orders/:id", controller.GetOrderByID)
}
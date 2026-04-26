package main

import (
	"fmt"
	"order-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
    routes.SetupRoutes(router)

	port := "8080"

	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
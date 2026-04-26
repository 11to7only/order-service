package service

import (
	"fmt"
	"order-service/model"
	"time"
)

var orders []model.Order
var nextID = 1

func GetAllOrders() []model.Order {
	fmt.Println("Service GetAllOrders called...")
	return orders
}

func AddOrder(order model.Order) {
	fmt.Println("Service AddOrder called...")

	order.ID = nextID
	order.CreatedAt = time.Now()

	nextID++
	orders = append(orders, order)
}

func GetOrderByID(id int) *model.Order {
    fmt.Println("Service GetOrderByID called...")
    for _, order := range orders {
        if order.ID == id {
            return &order
        }
    }
    return nil
}
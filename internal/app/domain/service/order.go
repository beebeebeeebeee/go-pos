package service

import "go-pos/internal/app/domain/entity"

type IOrderService interface {
	CreateNewOrder() entity.Order
	GetLatestOrderID() entity.OrderID
	GetOrderByID(orderID entity.OrderID) entity.Order
	UpdateOrderItem(order entity.Order, item entity.Item, adjustment int) entity.Order
}

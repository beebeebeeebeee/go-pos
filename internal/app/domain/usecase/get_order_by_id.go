package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func GetOrderByID(orderService service.IOrderService, id entity.OrderID) entity.Order {
	return orderService.GetOrderByID(id)
}

package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func CreateNewOrder(orderService service.IOrderService) entity.Order {
	return orderService.CreateNewOrder()
}

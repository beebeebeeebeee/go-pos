package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func GetLatestOrderID(orderService service.IOrderService) entity.OrderID {
	return orderService.GetLatestOrderID()
}

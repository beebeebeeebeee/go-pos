package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func UpdateOrderItem(
	orderService service.IOrderService,
	itemService service.IItemService,
	orderID entity.OrderID,
	itemID entity.ItemID,
	quantity int,
) entity.Order {
	order := GetOrderByID(orderService, orderID)
	item := GetItemByID(itemService, itemID)

	return orderService.UpdateOrderItem(order, item, quantity)
}

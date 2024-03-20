package service

import (
	"github.com/google/uuid"
	"go-pos/internal/app/domain/entity"
)

type OrderService struct {
	orders []entity.Order
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) CreateNewOrder() entity.Order {
	order := entity.Order{
		ID: entity.OrderID(uuid.NewString()),
	}
	o.orders = append(o.orders, order)
	return order
}

func (o *OrderService) GetLatestOrderID() entity.OrderID {
	if len(o.orders) == 0 {
		o.CreateNewOrder()
	}

	return o.orders[len(o.orders)-1].ID
}

func (o *OrderService) GetOrderByID(orderID entity.OrderID) entity.Order {
	for _, order := range o.orders {
		if order.ID == orderID {
			return order
		}
	}

	return entity.Order{}
}

func (o *OrderService) UpdateOrderItem(order entity.Order, item entity.Item, adjustment int) entity.Order {
	var orderItem *entity.OrderItem
	var orderItemIndex int
	for i, it := range order.OrderItems {
		if it.Item.ID == item.ID {
			orderItem = &order.OrderItems[i]
			orderItemIndex = i
			break
		}
	}

	if adjustment == 0 {
		return order
	}

	if orderItem == nil {
		if adjustment < 0 {
			return order
		}

		orderItem = &entity.OrderItem{
			Item:     item,
			Quantity: adjustment,
		}
		order.OrderItems = append(order.OrderItems, *orderItem)
	} else {
		orderItem.Quantity = orderItem.Quantity + adjustment

		if orderItem.Quantity == 0 {
			order.OrderItems = append(order.OrderItems[:orderItemIndex], order.OrderItems[orderItemIndex+1:]...)
			return order
		}
	}

	var orderTotalPrice float64
	for i, it := range order.OrderItems {
		order.OrderItems[i].TotalPrice = it.Item.Price * float64(it.Quantity)
		orderTotalPrice += order.OrderItems[i].TotalPrice
	}
	order.TotalPrice = orderTotalPrice

	for i, it := range o.orders {
		if it.ID == order.ID {
			o.orders[i] = order
			break
		}
	}

	return order
}

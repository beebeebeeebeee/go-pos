package controller

import (
	"context"
	"go-pos/internal/app/domain"
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/usecase"
)

type OrderController struct {
	ctx context.Context
	di  *domain.DI
}

func NewOrderController(di *domain.DI) *OrderController {
	return &OrderController{
		di: di,
	}
}

func (o *OrderController) SetCtx(ctx context.Context) {
	o.ctx = ctx
}

func (o *OrderController) CreateNewOrder() entity.Order {
	return usecase.CreateNewOrder(o.di.OrderService)
}

func (o *OrderController) GetLatestOrder() entity.Order {
	orderID := usecase.GetLatestOrderID(o.di.OrderService)
	return usecase.GetOrderByID(o.di.OrderService, orderID)
}

func (o *OrderController) UpdateOrderItem(order entity.OrderID, item entity.ItemID, quantity int) entity.Order {
	return usecase.UpdateOrderItem(o.di.OrderService, o.di.ItemService, order, item, quantity)
}

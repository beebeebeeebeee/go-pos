package controller

import (
	"context"
	"go-pos/internal/app/domain"
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

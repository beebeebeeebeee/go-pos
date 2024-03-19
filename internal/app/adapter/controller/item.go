package controller

import (
	"context"
	"go-pos/internal/app/domain"
	"go-pos/internal/app/domain/entity"
)

type ItemController struct {
	ctx context.Context
	di  *domain.DI
}

func NewItemController(
	di *domain.DI,
) *ItemController {
	return &ItemController{
		di: di,
	}
}

func (i *ItemController) SetCtx(ctx context.Context) {
	i.ctx = ctx
}

func (i *ItemController) GetItemList() []entity.Item {
	return []entity.Item{
		{
			Name:  "細餅",
			Price: 500.00,
		},
		{
			Name:  "大餅",
			Price: 1000.00,
		},
	}
}

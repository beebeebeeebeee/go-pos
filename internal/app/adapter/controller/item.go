package controller

import (
	"context"
	"go-pos/internal/app/domain"
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/usecase"
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
	return usecase.GetAllItems(i.di.ItemService)
}

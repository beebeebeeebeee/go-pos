package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func GetAllItems(itemService service.IItemService) []entity.Item {
	return itemService.GetAllItems()
}

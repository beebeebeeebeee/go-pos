package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func GetItemByID(itemService service.IItemService, id entity.ItemID) entity.Item {
	return itemService.GetItemByID(id)
}

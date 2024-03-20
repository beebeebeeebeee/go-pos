package service

import (
	"github.com/google/uuid"
	"go-pos/internal/app/domain/entity"
)

var (
	細餅 = entity.NewItem(entity.ItemID(uuid.NewString()), "細餅", 500.00)
	大餅 = entity.NewItem(entity.ItemID(uuid.NewString()), "大餅", 1000.00)
)

var ItemList = []entity.Item{
	細餅,
	大餅,
}

type ItemService struct {
}

func NewItemService() *ItemService {
	return &ItemService{}
}

func (s *ItemService) GetAllItems() []entity.Item {
	return ItemList
}

func (s *ItemService) GetItemByID(itemID entity.ItemID) entity.Item {
	for _, v := range ItemList {
		if v.ID == itemID {
			return v
		}
	}

	return entity.Item{}
}

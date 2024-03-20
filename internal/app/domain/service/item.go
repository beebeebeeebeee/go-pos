package service

import "go-pos/internal/app/domain/entity"

type IItemService interface {
	GetAllItems() []entity.Item
	GetItemByID(itemID entity.ItemID) entity.Item
}

package service

import "go-pos/internal/app/domain/entity"

type IReceiptService interface {
	GenerateReceipt(path *string, order entity.Order) (string, error)
}

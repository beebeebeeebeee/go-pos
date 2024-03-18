package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func GenerateReceipt(receiptService service.IReceiptService, order entity.Order) (string, error) {
	receiptText, err := receiptService.GenerateReceipt(nil, order)
	if err != nil {
		return "", err
	}

	return receiptText, nil
}

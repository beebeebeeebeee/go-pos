package service

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/infrastructure/receipt"
)

type ReceiptService struct {
	client *receipt.Client
}

func NewReceiptService(client *receipt.Client) *ReceiptService {
	return &ReceiptService{
		client: client,
	}
}

func (r *ReceiptService) GenerateReceipt(path *string, order entity.Order) (string, error) {
	receiptText, err := r.client.GetReceipt(path, order)
	if err != nil {
		return "", err
	}

	return receiptText, nil
}

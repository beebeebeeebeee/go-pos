package controller

import (
	"context"
	"go-pos/internal/app/domain"
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/usecase"
	"log"
)

type ReceiptController struct {
	ctx context.Context
	di  *domain.DI
}

func NewReceiptController(di *domain.DI) *ReceiptController {
	return &ReceiptController{
		di: di,
	}
}

func (a *ReceiptController) SetCtx(ctx context.Context) {
	a.ctx = ctx
}

var order = entity.Order{
	ID: "123",
	OrderItems: []entity.OrderItem{
		{
			Item: entity.Item{
				Name:  "大餅",
				Price: 1000.00,
			},
			Quantity:   37,
			TotalPrice: 37000.00,
		},
	},
	TotalPrice: 37000.00,
}

func (a *ReceiptController) GetReceipt() (string, error) {
	receiptText, err := usecase.GenerateReceipt(a.di.ReceiptService, order)
	if err != nil {
		log.Println("error while get receipt:", err)
		return "", err
	}

	return receiptText, nil
}

func (a *ReceiptController) PrintReceipt() error {
	if _, err := usecase.PrintReceipt(a.di.ReceiptService, a.di.PrinterService, order); err != nil {
		log.Println("error while print receipt:", err)
		return err
	}

	return nil
}

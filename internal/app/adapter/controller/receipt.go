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
	OrderID: "123",
	OrderItems: []entity.OrderItem{
		{
			Item: entity.Item{
				Name:  "Iced Americano",
				Price: 40.00,
			},
			Quantity:   1,
			TotalPrice: 40.00,
		},
		{
			Item: entity.Item{
				Name:  "Drip Coffee",
				Price: 80.00,
			},
			Quantity:   1,
			TotalPrice: 80.00,
		},
	},
	TotalPrice: 120.00,
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

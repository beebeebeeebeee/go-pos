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

func (a *ReceiptController) GetReceipt(orderID entity.OrderID) (string, error) {
	order := usecase.GetOrderByID(a.di.OrderService, orderID)

	receiptText, err := usecase.GenerateReceipt(a.di.ReceiptService, order)
	if err != nil {
		log.Println("error while get receipt:", err)
		return "", err
	}

	return receiptText, nil
}

func (a *ReceiptController) PrintReceipt(orderID entity.OrderID) error {
	order := usecase.GetOrderByID(a.di.OrderService, orderID)

	if _, err := usecase.PrintReceipt(a.di.ReceiptService, a.di.PrinterService, order); err != nil {
		log.Println("error while print receipt:", err)
		return err
	}

	return nil
}

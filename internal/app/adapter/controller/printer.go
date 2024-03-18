package controller

import (
	"context"
	"go-pos/internal/app/domain"
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/usecase"
)

type PrinterController struct {
	ctx context.Context
	di  *domain.DI
}

func NewPrinterController(di *domain.DI) *PrinterController {
	return &PrinterController{
		di: di,
	}
}

func (p *PrinterController) SetCtx(ctx context.Context) {
	p.ctx = ctx
}

func (p *PrinterController) GetPrinterList() []entity.Printer {
	return usecase.GetPrinterList(p.di.PrinterService)
}

func (p *PrinterController) SetSelectedPrinter(printer entity.Printer) []entity.Printer {
	usecase.SetSelectedPrinter(p.di.PrinterService, printer)
	return usecase.GetPrinterList(p.di.PrinterService)
}

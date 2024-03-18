package service

import "go-pos/internal/app/domain/entity"

type IPrinterService interface {
	GetPrinterList() []entity.Printer
	SetSelectedPrinter(printer entity.Printer)
	PrintFile(path string) error
}

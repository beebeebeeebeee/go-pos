package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func SetSelectedPrinter(printerService service.IPrinterService, printer entity.Printer) {
	printerService.SetSelectedPrinter(printer)
}

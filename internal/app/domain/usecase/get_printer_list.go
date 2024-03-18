package usecase

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
)

func GetPrinterList(printerService service.IPrinterService) []entity.Printer {
	return printerService.GetPrinterList()
}

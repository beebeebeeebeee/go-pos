package usecase

import "go-pos/internal/app/domain/service"

func PrintFile(printerService service.IPrinterService, path string) error {
	return printerService.PrintFile(path)
}

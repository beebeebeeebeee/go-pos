package domain

import "go-pos/internal/app/domain/service"

type DI struct {
	PrinterService service.IPrinterService
	ReceiptService service.IReceiptService
}

func NewDI(
	printerService service.IPrinterService,
	receiptService service.IReceiptService,
) *DI {
	return &DI{
		PrinterService: printerService,
		ReceiptService: receiptService,
	}
}

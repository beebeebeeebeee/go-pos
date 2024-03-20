package domain

import "go-pos/internal/app/domain/service"

type DI struct {
	ItemService    service.IItemService
	OrderService   service.IOrderService
	PrinterService service.IPrinterService
	ReceiptService service.IReceiptService
}

func NewDI(
	itemService service.IItemService,
	orderService service.IOrderService,
	printerService service.IPrinterService,
	receiptService service.IReceiptService,
) *DI {
	return &DI{
		ItemService:    itemService,
		OrderService:   orderService,
		PrinterService: printerService,
		ReceiptService: receiptService,
	}
}

package main

import (
	"context"
	"embed"
	"go-pos/internal/app/adapter/controller"
	"go-pos/internal/app/adapter/service"
	"go-pos/internal/app/domain"
	"go-pos/internal/app/infrastructure/receipt"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:internal/app/presentation/dist
var assets embed.FS

func main() {
	receiptClient := receipt.NewClient(float64(12), 38, "Company Name")

	printerService := service.NewPrinterService()
	receiptService := service.NewReceiptService(receiptClient)

	di := domain.NewDI(
		printerService,
		receiptService,
	)

	itemController := controller.NewItemController(di)
	orderController := controller.NewOrderController(di)
	receiptController := controller.NewReceiptController(di)
	printerController := controller.NewPrinterController(di)

	err := wails.Run(&options.App{
		Title:  "go-pos",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			itemController.SetCtx(ctx)
			orderController.SetCtx(ctx)
			receiptController.SetCtx(ctx)
			printerController.SetCtx(ctx)
		},
		Bind: []interface{}{
			itemController,
			orderController,
			receiptController,
			printerController,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

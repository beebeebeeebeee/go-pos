package main

import (
	"changeme/src"
	"changeme/src/domain/entity"
	"changeme/src/receipt"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	//a.Printer()
}

func (a *App) Printer() string {
	dir, _ := os.MkdirTemp("", "example")
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)

	tmpPath := path.Join(dir, fmt.Sprintf("%s.pdf", uuid.New().String()))
	//tmpPath := "hello.pdf"

	receiptClient := receipt.NewClient(
		10,
		24,
		"Fung Coffee",
	)

	err := receiptClient.GetReceipt(tmpPath, entity.Order{
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
	})
	if err != nil {
		return err.Error()
	}

	printer := src.NewDefaultConnection().Dests[0]
	if printer == nil {
		return "No printer found"
	}

	jobID := printer.TestPrint(tmpPath)
	log.Println(jobID)

	return printer.Name
	//return "Hello"
}

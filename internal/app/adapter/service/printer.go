package service

import (
	"errors"
	"go-pos/internal/app/adapter/util/converter"
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/infrastructure/cups"
)

type PrinterService struct {
	SelectedPrinter *entity.Printer
}

func NewPrinterService() *PrinterService {
	return &PrinterService{}
}

func (p *PrinterService) GetPrinterList() []entity.Printer {
	dests := cups.NewDefaultConnection().Dests
	var printers []entity.Printer
	for _, dest := range dests {
		printer := converter.PrinterCupsToDomain(*dest)
		if p.SelectedPrinter != nil && printer.Name == p.SelectedPrinter.Name {
			printer.Selected = true
		}
		printers = append(printers, printer)
	}

	return printers
}

func (p *PrinterService) SetSelectedPrinter(printer entity.Printer) {
	p.SelectedPrinter = &printer
}

func (p *PrinterService) PrintFile(path string) error {
	if p.SelectedPrinter == nil {
		return errors.New("no printer selected")
	}

	var printer *cups.Dest
	for _, dest := range cups.NewDefaultConnection().Dests {
		if dest.Name == p.SelectedPrinter.Name {
			printer = dest
			break
		}
	}
	if printer == nil {
		return errors.New("selected printer not found")
	}

	jobID := printer.PrintFile(path)
	if jobID == 0 {
		return errors.New("failed to print file")
	}

	return nil
}

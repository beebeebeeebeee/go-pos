package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/domain/service"
	"os"
	"path"
)

var mode = "test"

func PrintReceipt(receiptService service.IReceiptService, printerService service.IPrinterService, order entity.Order) (string, error) {
	dir, _ := os.MkdirTemp("", "example")
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)

	tmpPath := path.Join(dir, fmt.Sprintf("%s.pdf", uuid.New().String()))

	receiptText, err := receiptService.GenerateReceipt(&tmpPath, order)
	if err != nil {
		return "", err
	}

	if mode == "test" {
		source, err := os.ReadFile(tmpPath)
		if err != nil {
			return "", err
		}

		err = os.WriteFile("preview.pdf", source, 0644)

		return receiptText, nil
	}

	if err = PrintFile(printerService, tmpPath); err != nil {
		return "", err
	}

	return receiptText, nil
}

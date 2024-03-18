package converter

import (
	"go-pos/internal/app/domain/entity"
	"go-pos/internal/app/infrastructure/cups"
)

func PrinterCupsToDomain(p cups.Dest) entity.Printer {
	return entity.Printer{
		Name:      p.Name,
		IsDefault: p.IsDefault,
		Options:   p.Options,
	}
}

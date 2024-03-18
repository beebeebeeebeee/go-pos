package receipt

import (
	"go-pos/internal/app/domain/entity"
)

type Client struct {
	FontSize    float64
	LineSpacing int
	CompanyName string
}

func NewClient(
	fontSize float64,
	lineSpacing int,
	companyName string,
) *Client {
	return &Client{
		FontSize:    fontSize,
		LineSpacing: lineSpacing,
		CompanyName: companyName,
	}
}

func (c *Client) GetReceipt(
	path *string,
	orders entity.Order,
) (string, error) {
	printer := NewPrinter(c.FontSize, c.LineSpacing)

	printer.PrintLine(printer.GetCenterText(c.CompanyName))
	printer.PrintLine("Tel: 1234567890")
	printer.PrintLine(printer.GetDivider())

	for _, orderItem := range orders.OrderItems {
		printer.PrintLine(printer.GetStringWithSpacing(orderItem.Item.Name, printer.FormatPrice(orderItem.TotalPrice)))
	}

	printer.PrintLine(printer.GetDivider())
	printer.PrintLine(printer.GetStringWithSpacing("Total", printer.FormatPrice(orders.TotalPrice)))

	if path != nil {
		err := printer.pdf.OutputFileAndClose(*path)
		if err != nil {
			return "", err
		}
	}

	return printer.text, nil
}

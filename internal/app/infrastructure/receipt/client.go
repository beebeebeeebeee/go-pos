package receipt

import (
	"fmt"
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
	printer.PrintLine(printer.GetDivider())

	for _, orderItem := range orders.OrderItems {
		printer.PrintLine(orderItem.Item.Name)
		printer.PrintLine(printer.GetStringWithSpacing(fmt.Sprintf("%dx %s", orderItem.Quantity, printer.FormatPrice(orderItem.Item.Price)), printer.FormatPrice(orderItem.TotalPrice)))
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

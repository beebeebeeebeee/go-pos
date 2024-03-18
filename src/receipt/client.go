package receipt

import (
	"changeme/src/domain/entity"
	"fmt"
	"github.com/go-pdf/fpdf"
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
	path string,
	orders entity.Order,
) error {
	line := 0

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("SpaceMono", "", "SpaceMono-Regular.ttf")
	pdf.AddPage()
	pdf.SetFont("SpaceMono", "", c.FontSize)
	pdf.SetX(10)

	c.PrintLine(pdf, &line, c.GetCenterText(c.CompanyName))
	c.PrintLine(pdf, &line, c.getDivider())

	for _, orderItem := range orders.OrderItems {
		c.PrintLine(pdf, &line, c.GetStringWithSpacing(orderItem.Item.Name, formatPrice(orderItem.TotalPrice)))
	}

	c.PrintLine(pdf, &line, c.getDivider())
	c.PrintLine(pdf, &line, c.GetStringWithSpacing("Total", formatPrice(orders.TotalPrice)))

	err := pdf.OutputFileAndClose(path)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) PrintLine(pdf *fpdf.Fpdf, line *int, text string) {
	pdf.SetY(10 + c.FontSize/2*float64(*line))
	pdf.Cell(0, 0, text)
	*line++
}

func (c *Client) getDivider() string {
	var result string
	for i := 0; i < c.LineSpacing; i++ {
		result += "-"
	}
	return result
}

func (c *Client) GetCenterText(text string) string {
	return getSpacing(c.LineSpacing/2-len(text)/2) + text
}

func (c *Client) GetStringWithSpacing(text1 string, text2 string) string {
	return text1 + getSpacing(c.LineSpacing-len(text1)-len(text2)) + text2
}

func getSpacing(spaces int) string {
	var result string
	for i := 0; i < spaces; i++ {
		result += " "
	}
	return result
}

func formatPrice(price float64) string {
	return "$" + fmt.Sprintf("%.2f", price)
}

package receipt

import (
	"fmt"
	"github.com/go-pdf/fpdf"
)

type Printer struct {
	FontSize    float64
	LineSpacing int

	pdf  *fpdf.Fpdf
	text string
	line int
}

func NewPrinter(
	fontSize float64,
	lineSpacing int,
) *Printer {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddUTF8Font("SpaceMono", "", "fonts/SpaceMono-Regular.ttf")
	pdf.AddPage()
	pdf.SetFont("SpaceMono", "", fontSize)
	pdf.SetX(10)

	return &Printer{
		FontSize:    fontSize,
		LineSpacing: lineSpacing,

		line: 0,
		pdf:  pdf,
	}
}

func (p *Printer) PrintLine(text string) {
	p.text = p.text + text + "\n"

	p.pdf.SetY(10 + p.FontSize/2*float64(p.line))
	p.pdf.Cell(0, 0, text)
	p.line++
}

func (p *Printer) GetCenterText(text string) string {
	return p.GetSpacing(p.LineSpacing/2-len(text)/2) + text
}

func (p *Printer) GetStringWithSpacing(text1 string, text2 string) string {
	return text1 + p.GetSpacing(p.LineSpacing-len(text1)-len(text2)) + text2
}

func (p *Printer) GetDivider() string {
	var result string
	for i := 0; i < p.LineSpacing; i++ {
		result += "-"
	}
	return result
}

func (p *Printer) GetSpacing(spaces int) string {
	var result string
	for i := 0; i < spaces; i++ {
		result += " "
	}
	return result
}

func (p *Printer) FormatPrice(price float64) string {
	return "$" + fmt.Sprintf("%.2f", price)
}

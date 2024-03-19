package receipt

import (
	"github.com/go-pdf/fpdf"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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
	pdf.AddUTF8Font("AssetsMono", "", "fonts/MPLUS1Code-VariableFont_wght.ttf")
	pdf.AddPage()
	pdf.SetFont("AssetsMono", "", fontSize)
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
	return p.GetSpacing(p.LineSpacing/2-p.GetTextLength(text)/2) + text
}

func (p *Printer) GetStringWithSpacing(text1 string, text2 string) string {
	return text1 + p.GetSpacing(p.LineSpacing-p.GetTextLength(text1)-p.GetTextLength(text2)) + text2
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
	return message.NewPrinter(language.English).Sprintf("$%.2f", price)
}

func (p *Printer) GetTextLength(text string) int {
	// non english characters are count as 2
	var length int
	for _, char := range text {
		if char > 127 {
			length += 2
		} else {
			length++
		}
	}

	return length
}

package pdfType

import (
	"fmt"

	"github.com/andrewarrow/BuisnessPDF/generator"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func getAddressLine(address FullPersonInfo) string {
	var addressSenderSmallText = ""

	addressSenderSmallText += address.CompanyName
	if address.CompanyName != "" && (address.FullForename != "" || address.FullSurname != "") {
		addressSenderSmallText += ", "
	}

	addressSenderSmallText += address.FullForename
	if address.FullSurname != "" {
		addressSenderSmallText += " "
	}
	addressSenderSmallText += address.FullSurname

	addressSenderSmallText += fmt.Sprintf(" - %s %s",
		address.Address.Road,
		address.Address.HouseNumber,
	)

	if address.Address.StreetSupplement != "" {
		addressSenderSmallText += ", "
		addressSenderSmallText += address.Address.StreetSupplement
	}

	addressSenderSmallText += fmt.Sprintf(", %s %s %s",
		address.Address.CountryCode,
		address.Address.ZipCode,
		address.Address.CityName,
	)

	return addressSenderSmallText
}

func germanNumber[T float64 | int](n T) string {
	p := message.NewPrinter(language.English)

	switch fmt.Sprintf("%T", *new(T)) {
	case "float64":
		return p.Sprintf("%.2f", n)
	case "int":
		return p.Sprintf("%d", n)
	default:
		return "GERMAN NUMBER FAILED"
	}
}

func getCellWith(pdfGen *generator.PDFGenerator, percent float64) float64 {
	maxSavePrintingWidth, _ := pdfGen.GetPdf().GetPageSize()
	maxSavePrintingWidth = maxSavePrintingWidth - pdfGen.GetMarginLeft() - pdfGen.GetMarginRight()

	return (percent * maxSavePrintingWidth) / 100.0
}

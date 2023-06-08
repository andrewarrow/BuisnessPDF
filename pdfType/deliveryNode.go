package pdfType

import (
	"encoding/json"
	errorsWithStack "github.com/go-errors/errors"
	"github.com/jung-kurt/gofpdf"
	"github.com/rs/zerolog"
	"io"
	"net/http"
)

type DeliveryNode struct {
	data          deliveryNodeRequestData
	meta          pdfMeta
	logger        *zerolog.Logger
	printErrStack bool
}

type deliveryNodeRequestData struct {
}

func NewDeliveryNode(logger *zerolog.Logger) *DeliveryNode {
	return &DeliveryNode{
		data: deliveryNodeRequestData{},
		meta: pdfMeta{
			margin: pdfMargin{
				left:   25,
				right:  20,
				top:    45,
				bottom: 0,
			},
			font: pdfFont{
				fontName:    "openSans",
				sizeDefault: 10,
				sizeSmall:   8,
				SizeLarge:   15,
			},
		},
		logger:        logger,
		printErrStack: logger.GetLevel() <= zerolog.DebugLevel,
	}
}

func (d *DeliveryNode) SetDataFromRequest(request *http.Request) (err error) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			d.LogError(err)
		}
	}(request.Body)

	err = json.NewDecoder(request.Body).Decode(&d.data)
	if err != nil {
		return err
	}

	err = d.validateData()
	if err != nil {
		d.data = deliveryNodeRequestData{}
		return err
	}

	return nil
}

func (d *DeliveryNode) GeneratePDF() (*gofpdf.Fpdf, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DeliveryNode) LogError(err error) {
	var errStr string

	if _, ok := err.(*errorsWithStack.Error); ok && d.printErrStack {
		errStr = err.(*errorsWithStack.Error).ErrorStack()
	} else {
		errStr = err.Error()
	}

	d.logger.Error().Msgf(errStr)
}

func (d *DeliveryNode) validateData() (err error) {
	return err
}

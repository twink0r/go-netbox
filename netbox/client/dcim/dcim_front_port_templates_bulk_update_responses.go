// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// DcimFrontPortTemplatesBulkUpdateReader is a Reader for the DcimFrontPortTemplatesBulkUpdate structure.
type DcimFrontPortTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesBulkUpdateOK creates a DcimFrontPortTemplatesBulkUpdateOK with default headers values
func NewDcimFrontPortTemplatesBulkUpdateOK() *DcimFrontPortTemplatesBulkUpdateOK {
	return &DcimFrontPortTemplatesBulkUpdateOK{}
}

/* DcimFrontPortTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesBulkUpdateOK dcim front port templates bulk update o k
*/
type DcimFrontPortTemplatesBulkUpdateOK struct {
	Payload *models.FrontPortTemplate
}

func (o *DcimFrontPortTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortTemplatesBulkUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesBulkUpdateDefault creates a DcimFrontPortTemplatesBulkUpdateDefault with default headers values
func NewDcimFrontPortTemplatesBulkUpdateDefault(code int) *DcimFrontPortTemplatesBulkUpdateDefault {
	return &DcimFrontPortTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortTemplatesBulkUpdateDefault dcim front port templates bulk update default
*/
type DcimFrontPortTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates bulk update default response
func (o *DcimFrontPortTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/][%d] dcim_front-port-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

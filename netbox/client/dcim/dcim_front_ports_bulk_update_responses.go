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

// DcimFrontPortsBulkUpdateReader is a Reader for the DcimFrontPortsBulkUpdate structure.
type DcimFrontPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsBulkUpdateOK creates a DcimFrontPortsBulkUpdateOK with default headers values
func NewDcimFrontPortsBulkUpdateOK() *DcimFrontPortsBulkUpdateOK {
	return &DcimFrontPortsBulkUpdateOK{}
}

/* DcimFrontPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsBulkUpdateOK dcim front ports bulk update o k
*/
type DcimFrontPortsBulkUpdateOK struct {
	Payload *models.FrontPort
}

func (o *DcimFrontPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcimFrontPortsBulkUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimFrontPortsBulkUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsBulkUpdateDefault creates a DcimFrontPortsBulkUpdateDefault with default headers values
func NewDcimFrontPortsBulkUpdateDefault(code int) *DcimFrontPortsBulkUpdateDefault {
	return &DcimFrontPortsBulkUpdateDefault{
		_statusCode: code,
	}
}

/* DcimFrontPortsBulkUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortsBulkUpdateDefault dcim front ports bulk update default
*/
type DcimFrontPortsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports bulk update default response
func (o *DcimFrontPortsBulkUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimFrontPortsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcim_front-ports_bulk_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimFrontPortsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

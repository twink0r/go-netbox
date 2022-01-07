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

// DcimPowerFeedsBulkPartialUpdateReader is a Reader for the DcimPowerFeedsBulkPartialUpdate structure.
type DcimPowerFeedsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerFeedsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsBulkPartialUpdateOK creates a DcimPowerFeedsBulkPartialUpdateOK with default headers values
func NewDcimPowerFeedsBulkPartialUpdateOK() *DcimPowerFeedsBulkPartialUpdateOK {
	return &DcimPowerFeedsBulkPartialUpdateOK{}
}

/* DcimPowerFeedsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsBulkPartialUpdateOK dcim power feeds bulk partial update o k
*/
type DcimPowerFeedsBulkPartialUpdateOK struct {
	Payload *models.PowerFeed
}

func (o *DcimPowerFeedsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/][%d] dcimPowerFeedsBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimPowerFeedsBulkPartialUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsBulkPartialUpdateDefault creates a DcimPowerFeedsBulkPartialUpdateDefault with default headers values
func NewDcimPowerFeedsBulkPartialUpdateDefault(code int) *DcimPowerFeedsBulkPartialUpdateDefault {
	return &DcimPowerFeedsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimPowerFeedsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerFeedsBulkPartialUpdateDefault dcim power feeds bulk partial update default
*/
type DcimPowerFeedsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds bulk partial update default response
func (o *DcimPowerFeedsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerFeedsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-feeds/][%d] dcim_power-feeds_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimPowerFeedsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

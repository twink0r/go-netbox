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

// DcimRegionsUpdateReader is a Reader for the DcimRegionsUpdate structure.
type DcimRegionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRegionsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsUpdateOK creates a DcimRegionsUpdateOK with default headers values
func NewDcimRegionsUpdateOK() *DcimRegionsUpdateOK {
	return &DcimRegionsUpdateOK{}
}

/* DcimRegionsUpdateOK describes a response with status code 200, with default header values.

DcimRegionsUpdateOK dcim regions update o k
*/
type DcimRegionsUpdateOK struct {
	Payload *models.Region
}

func (o *DcimRegionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimRegionsUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsUpdateDefault creates a DcimRegionsUpdateDefault with default headers values
func NewDcimRegionsUpdateDefault(code int) *DcimRegionsUpdateDefault {
	return &DcimRegionsUpdateDefault{
		_statusCode: code,
	}
}

/* DcimRegionsUpdateDefault describes a response with status code -1, with default header values.

DcimRegionsUpdateDefault dcim regions update default
*/
type DcimRegionsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions update default response
func (o *DcimRegionsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimRegionsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcim_regions_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimRegionsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// CircuitsCircuitsUpdateReader is a Reader for the CircuitsCircuitsUpdate structure.
type CircuitsCircuitsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCircuitsCircuitsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitsUpdateOK creates a CircuitsCircuitsUpdateOK with default headers values
func NewCircuitsCircuitsUpdateOK() *CircuitsCircuitsUpdateOK {
	return &CircuitsCircuitsUpdateOK{}
}

/* CircuitsCircuitsUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitsUpdateOK circuits circuits update o k
*/
type CircuitsCircuitsUpdateOK struct {
	Payload *models.Circuit
}

func (o *CircuitsCircuitsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuits/{id}/][%d] circuitsCircuitsUpdateOK  %+v", 200, o.Payload)
}
func (o *CircuitsCircuitsUpdateOK) GetPayload() *models.Circuit {
	return o.Payload
}

func (o *CircuitsCircuitsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Circuit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitsUpdateDefault creates a CircuitsCircuitsUpdateDefault with default headers values
func NewCircuitsCircuitsUpdateDefault(code int) *CircuitsCircuitsUpdateDefault {
	return &CircuitsCircuitsUpdateDefault{
		_statusCode: code,
	}
}

/* CircuitsCircuitsUpdateDefault describes a response with status code -1, with default header values.

CircuitsCircuitsUpdateDefault circuits circuits update default
*/
type CircuitsCircuitsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuits update default response
func (o *CircuitsCircuitsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *CircuitsCircuitsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuits/{id}/][%d] circuits_circuits_update default  %+v", o._statusCode, o.Payload)
}
func (o *CircuitsCircuitsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

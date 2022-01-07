// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// VirtualizationInterfacesBulkPartialUpdateReader is a Reader for the VirtualizationInterfacesBulkPartialUpdate structure.
type VirtualizationInterfacesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationInterfacesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesBulkPartialUpdateOK creates a VirtualizationInterfacesBulkPartialUpdateOK with default headers values
func NewVirtualizationInterfacesBulkPartialUpdateOK() *VirtualizationInterfacesBulkPartialUpdateOK {
	return &VirtualizationInterfacesBulkPartialUpdateOK{}
}

/* VirtualizationInterfacesBulkPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationInterfacesBulkPartialUpdateOK virtualization interfaces bulk partial update o k
*/
type VirtualizationInterfacesBulkPartialUpdateOK struct {
	Payload *models.VMInterface
}

func (o *VirtualizationInterfacesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/interfaces/][%d] virtualizationInterfacesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *VirtualizationInterfacesBulkPartialUpdateOK) GetPayload() *models.VMInterface {
	return o.Payload
}

func (o *VirtualizationInterfacesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VMInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationInterfacesBulkPartialUpdateDefault creates a VirtualizationInterfacesBulkPartialUpdateDefault with default headers values
func NewVirtualizationInterfacesBulkPartialUpdateDefault(code int) *VirtualizationInterfacesBulkPartialUpdateDefault {
	return &VirtualizationInterfacesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* VirtualizationInterfacesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesBulkPartialUpdateDefault virtualization interfaces bulk partial update default
*/
type VirtualizationInterfacesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces bulk partial update default response
func (o *VirtualizationInterfacesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *VirtualizationInterfacesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/interfaces/][%d] virtualization_interfaces_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *VirtualizationInterfacesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

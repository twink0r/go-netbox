// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamVrfsBulkDeleteReader is a Reader for the IpamVrfsBulkDelete structure.
type IpamVrfsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamVrfsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamVrfsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVrfsBulkDeleteNoContent creates a IpamVrfsBulkDeleteNoContent with default headers values
func NewIpamVrfsBulkDeleteNoContent() *IpamVrfsBulkDeleteNoContent {
	return &IpamVrfsBulkDeleteNoContent{}
}

/* IpamVrfsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamVrfsBulkDeleteNoContent ipam vrfs bulk delete no content
*/
type IpamVrfsBulkDeleteNoContent struct {
}

func (o *IpamVrfsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipamVrfsBulkDeleteNoContent ", 204)
}

func (o *IpamVrfsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamVrfsBulkDeleteDefault creates a IpamVrfsBulkDeleteDefault with default headers values
func NewIpamVrfsBulkDeleteDefault(code int) *IpamVrfsBulkDeleteDefault {
	return &IpamVrfsBulkDeleteDefault{
		_statusCode: code,
	}
}

/* IpamVrfsBulkDeleteDefault describes a response with status code -1, with default header values.

IpamVrfsBulkDeleteDefault ipam vrfs bulk delete default
*/
type IpamVrfsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vrfs bulk delete default response
func (o *IpamVrfsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *IpamVrfsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/vrfs/][%d] ipam_vrfs_bulk_delete default  %+v", o._statusCode, o.Payload)
}
func (o *IpamVrfsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/fbreckle/go-netbox/netbox/models"
)

// IpamL2vpnsListReader is a Reader for the IpamL2vpnsList structure.
type IpamL2vpnsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamL2vpnsListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnsListOK creates a IpamL2vpnsListOK with default headers values
func NewIpamL2vpnsListOK() *IpamL2vpnsListOK {
	return &IpamL2vpnsListOK{}
}

/* IpamL2vpnsListOK describes a response with status code 200, with default header values.

IpamL2vpnsListOK ipam l2vpns list o k
*/
type IpamL2vpnsListOK struct {
	Payload *IpamL2vpnsListOKBody
}

func (o *IpamL2vpnsListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/l2vpns/][%d] ipamL2vpnsListOK  %+v", 200, o.Payload)
}
func (o *IpamL2vpnsListOK) GetPayload() *IpamL2vpnsListOKBody {
	return o.Payload
}

func (o *IpamL2vpnsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IpamL2vpnsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnsListDefault creates a IpamL2vpnsListDefault with default headers values
func NewIpamL2vpnsListDefault(code int) *IpamL2vpnsListDefault {
	return &IpamL2vpnsListDefault{
		_statusCode: code,
	}
}

/* IpamL2vpnsListDefault describes a response with status code -1, with default header values.

IpamL2vpnsListDefault ipam l2vpns list default
*/
type IpamL2vpnsListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpns list default response
func (o *IpamL2vpnsListDefault) Code() int {
	return o._statusCode
}

func (o *IpamL2vpnsListDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/l2vpns/][%d] ipam_l2vpns_list default  %+v", o._statusCode, o.Payload)
}
func (o *IpamL2vpnsListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IpamL2vpnsListOKBody ipam l2vpns list o k body
swagger:model IpamL2vpnsListOKBody
*/
type IpamL2vpnsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.L2VPN `json:"results"`
}

// Validate validates this ipam l2vpns list o k body
func (o *IpamL2vpnsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamL2vpnsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamL2vpnsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IpamL2vpnsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamL2vpnsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamL2vpnsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamL2vpnsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IpamL2vpnsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamL2vpnsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamL2vpnsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamL2vpnsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ipam l2vpns list o k body based on the context it is used
func (o *IpamL2vpnsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IpamL2vpnsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamL2vpnsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipamL2vpnsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IpamL2vpnsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IpamL2vpnsListOKBody) UnmarshalBinary(b []byte) error {
	var res IpamL2vpnsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

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

// DcimSiteGroupsPartialUpdateReader is a Reader for the DcimSiteGroupsPartialUpdate structure.
type DcimSiteGroupsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSiteGroupsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSiteGroupsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimSiteGroupsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSiteGroupsPartialUpdateOK creates a DcimSiteGroupsPartialUpdateOK with default headers values
func NewDcimSiteGroupsPartialUpdateOK() *DcimSiteGroupsPartialUpdateOK {
	return &DcimSiteGroupsPartialUpdateOK{}
}

/* DcimSiteGroupsPartialUpdateOK describes a response with status code 200, with default header values.

DcimSiteGroupsPartialUpdateOK dcim site groups partial update o k
*/
type DcimSiteGroupsPartialUpdateOK struct {
	Payload *models.SiteGroup
}

func (o *DcimSiteGroupsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/site-groups/{id}/][%d] dcimSiteGroupsPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimSiteGroupsPartialUpdateOK) GetPayload() *models.SiteGroup {
	return o.Payload
}

func (o *DcimSiteGroupsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SiteGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSiteGroupsPartialUpdateDefault creates a DcimSiteGroupsPartialUpdateDefault with default headers values
func NewDcimSiteGroupsPartialUpdateDefault(code int) *DcimSiteGroupsPartialUpdateDefault {
	return &DcimSiteGroupsPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimSiteGroupsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimSiteGroupsPartialUpdateDefault dcim site groups partial update default
*/
type DcimSiteGroupsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim site groups partial update default response
func (o *DcimSiteGroupsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimSiteGroupsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/site-groups/{id}/][%d] dcim_site-groups_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimSiteGroupsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSiteGroupsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

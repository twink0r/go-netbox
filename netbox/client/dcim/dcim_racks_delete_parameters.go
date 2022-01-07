// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimRacksDeleteParams creates a new DcimRacksDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksDeleteParams() *DcimRacksDeleteParams {
	return &DcimRacksDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksDeleteParamsWithTimeout creates a new DcimRacksDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRacksDeleteParamsWithTimeout(timeout time.Duration) *DcimRacksDeleteParams {
	return &DcimRacksDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRacksDeleteParamsWithContext creates a new DcimRacksDeleteParams object
// with the ability to set a context for a request.
func NewDcimRacksDeleteParamsWithContext(ctx context.Context) *DcimRacksDeleteParams {
	return &DcimRacksDeleteParams{
		Context: ctx,
	}
}

// NewDcimRacksDeleteParamsWithHTTPClient creates a new DcimRacksDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksDeleteParamsWithHTTPClient(client *http.Client) *DcimRacksDeleteParams {
	return &DcimRacksDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRacksDeleteParams contains all the parameters to send to the API endpoint
   for the dcim racks delete operation.

   Typically these are written to a http.Request.
*/
type DcimRacksDeleteParams struct {

	/* ID.

	   A unique integer value identifying this rack.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksDeleteParams) WithDefaults() *DcimRacksDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks delete params
func (o *DcimRacksDeleteParams) WithTimeout(timeout time.Duration) *DcimRacksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks delete params
func (o *DcimRacksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks delete params
func (o *DcimRacksDeleteParams) WithContext(ctx context.Context) *DcimRacksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks delete params
func (o *DcimRacksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks delete params
func (o *DcimRacksDeleteParams) WithHTTPClient(client *http.Client) *DcimRacksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks delete params
func (o *DcimRacksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim racks delete params
func (o *DcimRacksDeleteParams) WithID(id int64) *DcimRacksDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks delete params
func (o *DcimRacksDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

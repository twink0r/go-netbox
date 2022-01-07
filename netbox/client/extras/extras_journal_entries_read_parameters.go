// Code generated by go-swagger; DO NOT EDIT.

package extras

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

// NewExtrasJournalEntriesReadParams creates a new ExtrasJournalEntriesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJournalEntriesReadParams() *ExtrasJournalEntriesReadParams {
	return &ExtrasJournalEntriesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJournalEntriesReadParamsWithTimeout creates a new ExtrasJournalEntriesReadParams object
// with the ability to set a timeout on a request.
func NewExtrasJournalEntriesReadParamsWithTimeout(timeout time.Duration) *ExtrasJournalEntriesReadParams {
	return &ExtrasJournalEntriesReadParams{
		timeout: timeout,
	}
}

// NewExtrasJournalEntriesReadParamsWithContext creates a new ExtrasJournalEntriesReadParams object
// with the ability to set a context for a request.
func NewExtrasJournalEntriesReadParamsWithContext(ctx context.Context) *ExtrasJournalEntriesReadParams {
	return &ExtrasJournalEntriesReadParams{
		Context: ctx,
	}
}

// NewExtrasJournalEntriesReadParamsWithHTTPClient creates a new ExtrasJournalEntriesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJournalEntriesReadParamsWithHTTPClient(client *http.Client) *ExtrasJournalEntriesReadParams {
	return &ExtrasJournalEntriesReadParams{
		HTTPClient: client,
	}
}

/* ExtrasJournalEntriesReadParams contains all the parameters to send to the API endpoint
   for the extras journal entries read operation.

   Typically these are written to a http.Request.
*/
type ExtrasJournalEntriesReadParams struct {

	/* ID.

	   A unique integer value identifying this journal entry.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras journal entries read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJournalEntriesReadParams) WithDefaults() *ExtrasJournalEntriesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras journal entries read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJournalEntriesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) WithTimeout(timeout time.Duration) *ExtrasJournalEntriesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) WithContext(ctx context.Context) *ExtrasJournalEntriesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) WithHTTPClient(client *http.Client) *ExtrasJournalEntriesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) WithID(id int64) *ExtrasJournalEntriesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras journal entries read params
func (o *ExtrasJournalEntriesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJournalEntriesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

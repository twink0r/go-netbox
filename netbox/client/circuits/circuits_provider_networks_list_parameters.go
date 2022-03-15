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

package circuits

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

// NewCircuitsProviderNetworksListParams creates a new CircuitsProviderNetworksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProviderNetworksListParams() *CircuitsProviderNetworksListParams {
	return &CircuitsProviderNetworksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProviderNetworksListParamsWithTimeout creates a new CircuitsProviderNetworksListParams object
// with the ability to set a timeout on a request.
func NewCircuitsProviderNetworksListParamsWithTimeout(timeout time.Duration) *CircuitsProviderNetworksListParams {
	return &CircuitsProviderNetworksListParams{
		timeout: timeout,
	}
}

// NewCircuitsProviderNetworksListParamsWithContext creates a new CircuitsProviderNetworksListParams object
// with the ability to set a context for a request.
func NewCircuitsProviderNetworksListParamsWithContext(ctx context.Context) *CircuitsProviderNetworksListParams {
	return &CircuitsProviderNetworksListParams{
		Context: ctx,
	}
}

// NewCircuitsProviderNetworksListParamsWithHTTPClient creates a new CircuitsProviderNetworksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProviderNetworksListParamsWithHTTPClient(client *http.Client) *CircuitsProviderNetworksListParams {
	return &CircuitsProviderNetworksListParams{
		HTTPClient: client,
	}
}

/* CircuitsProviderNetworksListParams contains all the parameters to send to the API endpoint
   for the circuits provider networks list operation.

   Typically these are written to a http.Request.
*/
type CircuitsProviderNetworksListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionEmpty.
	DescriptionEmpty *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Provider.
	Provider *string

	// Providern.
	Providern *string

	// ProviderID.
	ProviderID *string

	// ProviderIDn.
	ProviderIDn *string

	// Q.
	Q *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits provider networks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksListParams) WithDefaults() *CircuitsProviderNetworksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits provider networks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProviderNetworksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithTimeout(timeout time.Duration) *CircuitsProviderNetworksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithContext(ctx context.Context) *CircuitsProviderNetworksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithHTTPClient(client *http.Client) *CircuitsProviderNetworksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithCreated(created *string) *CircuitsProviderNetworksListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithCreatedGte(createdGte *string) *CircuitsProviderNetworksListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithCreatedLte(createdLte *string) *CircuitsProviderNetworksListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescription(description *string) *CircuitsProviderNetworksListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionEmpty adds the descriptionEmpty to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionEmpty(descriptionEmpty *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionEmpty(descriptionEmpty)
	return o
}

// SetDescriptionEmpty adds the descriptionEmpty to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionEmpty(descriptionEmpty *string) {
	o.DescriptionEmpty = descriptionEmpty
}

// WithDescriptionIc adds the descriptionIc to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionIc(descriptionIc *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionIe(descriptionIe *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionIew(descriptionIew *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionIsw(descriptionIsw *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionn(descriptionn *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionNic(descriptionNic *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionNie(descriptionNie *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionNiew(descriptionNiew *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithDescriptionNisw(descriptionNisw *string) *CircuitsProviderNetworksListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithID(id *string) *CircuitsProviderNetworksListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithIDGt(iDGt *string) *CircuitsProviderNetworksListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithIDGte(iDGte *string) *CircuitsProviderNetworksListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithIDLt(iDLt *string) *CircuitsProviderNetworksListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithIDLte(iDLte *string) *CircuitsProviderNetworksListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithIDn(iDn *string) *CircuitsProviderNetworksListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithLastUpdated(lastUpdated *string) *CircuitsProviderNetworksListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithLastUpdatedGte(lastUpdatedGte *string) *CircuitsProviderNetworksListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithLastUpdatedLte(lastUpdatedLte *string) *CircuitsProviderNetworksListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithLimit(limit *int64) *CircuitsProviderNetworksListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithName(name *string) *CircuitsProviderNetworksListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameEmpty(nameEmpty *string) *CircuitsProviderNetworksListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameIc(nameIc *string) *CircuitsProviderNetworksListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameIe(nameIe *string) *CircuitsProviderNetworksListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameIew(nameIew *string) *CircuitsProviderNetworksListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameIsw(nameIsw *string) *CircuitsProviderNetworksListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNamen(namen *string) *CircuitsProviderNetworksListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameNic(nameNic *string) *CircuitsProviderNetworksListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameNie(nameNie *string) *CircuitsProviderNetworksListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameNiew(nameNiew *string) *CircuitsProviderNetworksListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithNameNisw(nameNisw *string) *CircuitsProviderNetworksListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithOffset(offset *int64) *CircuitsProviderNetworksListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithProvider adds the provider to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithProvider(provider *string) *CircuitsProviderNetworksListParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithProvidern adds the providern to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithProvidern(providern *string) *CircuitsProviderNetworksListParams {
	o.SetProvidern(providern)
	return o
}

// SetProvidern adds the providerN to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetProvidern(providern *string) {
	o.Providern = providern
}

// WithProviderID adds the providerID to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithProviderID(providerID *string) *CircuitsProviderNetworksListParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WithProviderIDn adds the providerIDn to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithProviderIDn(providerIDn *string) *CircuitsProviderNetworksListParams {
	o.SetProviderIDn(providerIDn)
	return o
}

// SetProviderIDn adds the providerIdN to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetProviderIDn(providerIDn *string) {
	o.ProviderIDn = providerIDn
}

// WithQ adds the q to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithQ(q *string) *CircuitsProviderNetworksListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithTag(tag *string) *CircuitsProviderNetworksListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) WithTagn(tagn *string) *CircuitsProviderNetworksListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the circuits provider networks list params
func (o *CircuitsProviderNetworksListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProviderNetworksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionEmpty != nil {

		// query param description__empty
		var qrDescriptionEmpty string

		if o.DescriptionEmpty != nil {
			qrDescriptionEmpty = *o.DescriptionEmpty
		}
		qDescriptionEmpty := qrDescriptionEmpty
		if qDescriptionEmpty != "" {

			if err := r.SetQueryParam("description__empty", qDescriptionEmpty); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if o.Providern != nil {

		// query param provider__n
		var qrProvidern string

		if o.Providern != nil {
			qrProvidern = *o.Providern
		}
		qProvidern := qrProvidern
		if qProvidern != "" {

			if err := r.SetQueryParam("provider__n", qProvidern); err != nil {
				return err
			}
		}
	}

	if o.ProviderID != nil {

		// query param provider_id
		var qrProviderID string

		if o.ProviderID != nil {
			qrProviderID = *o.ProviderID
		}
		qProviderID := qrProviderID
		if qProviderID != "" {

			if err := r.SetQueryParam("provider_id", qProviderID); err != nil {
				return err
			}
		}
	}

	if o.ProviderIDn != nil {

		// query param provider_id__n
		var qrProviderIDn string

		if o.ProviderIDn != nil {
			qrProviderIDn = *o.ProviderIDn
		}
		qProviderIDn := qrProviderIDn
		if qProviderIDn != "" {

			if err := r.SetQueryParam("provider_id__n", qProviderIDn); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

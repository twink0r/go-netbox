// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIpamRouteTargetsListParams creates a new IpamRouteTargetsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRouteTargetsListParams() *IpamRouteTargetsListParams {
	return &IpamRouteTargetsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsListParamsWithTimeout creates a new IpamRouteTargetsListParams object
// with the ability to set a timeout on a request.
func NewIpamRouteTargetsListParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsListParams {
	return &IpamRouteTargetsListParams{
		timeout: timeout,
	}
}

// NewIpamRouteTargetsListParamsWithContext creates a new IpamRouteTargetsListParams object
// with the ability to set a context for a request.
func NewIpamRouteTargetsListParamsWithContext(ctx context.Context) *IpamRouteTargetsListParams {
	return &IpamRouteTargetsListParams{
		Context: ctx,
	}
}

// NewIpamRouteTargetsListParamsWithHTTPClient creates a new IpamRouteTargetsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRouteTargetsListParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsListParams {
	return &IpamRouteTargetsListParams{
		HTTPClient: client,
	}
}

/* IpamRouteTargetsListParams contains all the parameters to send to the API endpoint
   for the ipam route targets list operation.

   Typically these are written to a http.Request.
*/
type IpamRouteTargetsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// ExportingVrf.
	ExportingVrf *string

	// ExportingVrfn.
	ExportingVrfn *string

	// ExportingVrfID.
	ExportingVrfID *string

	// ExportingVrfIDn.
	ExportingVrfIDn *string

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

	// ImportingVrf.
	ImportingVrf *string

	// ImportingVrfn.
	ImportingVrfn *string

	// ImportingVrfID.
	ImportingVrfID *string

	// ImportingVrfIDn.
	ImportingVrfIDn *string

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

	// Q.
	Q *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// Tenantn.
	Tenantn *string

	// TenantGroup.
	TenantGroup *string

	// TenantGroupn.
	TenantGroupn *string

	// TenantGroupID.
	TenantGroupID *string

	// TenantGroupIDn.
	TenantGroupIDn *string

	// TenantID.
	TenantID *string

	// TenantIDn.
	TenantIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam route targets list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsListParams) WithDefaults() *IpamRouteTargetsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam route targets list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithContext(ctx context.Context) *IpamRouteTargetsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithCreated(created *string) *IpamRouteTargetsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithCreatedGte(createdGte *string) *IpamRouteTargetsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithCreatedLte(createdLte *string) *IpamRouteTargetsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithExportingVrf adds the exportingVrf to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithExportingVrf(exportingVrf *string) *IpamRouteTargetsListParams {
	o.SetExportingVrf(exportingVrf)
	return o
}

// SetExportingVrf adds the exportingVrf to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetExportingVrf(exportingVrf *string) {
	o.ExportingVrf = exportingVrf
}

// WithExportingVrfn adds the exportingVrfn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithExportingVrfn(exportingVrfn *string) *IpamRouteTargetsListParams {
	o.SetExportingVrfn(exportingVrfn)
	return o
}

// SetExportingVrfn adds the exportingVrfN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetExportingVrfn(exportingVrfn *string) {
	o.ExportingVrfn = exportingVrfn
}

// WithExportingVrfID adds the exportingVrfID to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithExportingVrfID(exportingVrfID *string) *IpamRouteTargetsListParams {
	o.SetExportingVrfID(exportingVrfID)
	return o
}

// SetExportingVrfID adds the exportingVrfId to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetExportingVrfID(exportingVrfID *string) {
	o.ExportingVrfID = exportingVrfID
}

// WithExportingVrfIDn adds the exportingVrfIDn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithExportingVrfIDn(exportingVrfIDn *string) *IpamRouteTargetsListParams {
	o.SetExportingVrfIDn(exportingVrfIDn)
	return o
}

// SetExportingVrfIDn adds the exportingVrfIdN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetExportingVrfIDn(exportingVrfIDn *string) {
	o.ExportingVrfIDn = exportingVrfIDn
}

// WithID adds the id to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithID(id *string) *IpamRouteTargetsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithIDGt(iDGt *string) *IpamRouteTargetsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithIDGte(iDGte *string) *IpamRouteTargetsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithIDLt(iDLt *string) *IpamRouteTargetsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithIDLte(iDLte *string) *IpamRouteTargetsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithIDn(iDn *string) *IpamRouteTargetsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithImportingVrf adds the importingVrf to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithImportingVrf(importingVrf *string) *IpamRouteTargetsListParams {
	o.SetImportingVrf(importingVrf)
	return o
}

// SetImportingVrf adds the importingVrf to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetImportingVrf(importingVrf *string) {
	o.ImportingVrf = importingVrf
}

// WithImportingVrfn adds the importingVrfn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithImportingVrfn(importingVrfn *string) *IpamRouteTargetsListParams {
	o.SetImportingVrfn(importingVrfn)
	return o
}

// SetImportingVrfn adds the importingVrfN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetImportingVrfn(importingVrfn *string) {
	o.ImportingVrfn = importingVrfn
}

// WithImportingVrfID adds the importingVrfID to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithImportingVrfID(importingVrfID *string) *IpamRouteTargetsListParams {
	o.SetImportingVrfID(importingVrfID)
	return o
}

// SetImportingVrfID adds the importingVrfId to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetImportingVrfID(importingVrfID *string) {
	o.ImportingVrfID = importingVrfID
}

// WithImportingVrfIDn adds the importingVrfIDn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithImportingVrfIDn(importingVrfIDn *string) *IpamRouteTargetsListParams {
	o.SetImportingVrfIDn(importingVrfIDn)
	return o
}

// SetImportingVrfIDn adds the importingVrfIdN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetImportingVrfIDn(importingVrfIDn *string) {
	o.ImportingVrfIDn = importingVrfIDn
}

// WithLastUpdated adds the lastUpdated to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithLastUpdated(lastUpdated *string) *IpamRouteTargetsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamRouteTargetsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamRouteTargetsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithLimit(limit *int64) *IpamRouteTargetsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithName(name *string) *IpamRouteTargetsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameEmpty(nameEmpty *string) *IpamRouteTargetsListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameIc(nameIc *string) *IpamRouteTargetsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameIe(nameIe *string) *IpamRouteTargetsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameIew(nameIew *string) *IpamRouteTargetsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameIsw(nameIsw *string) *IpamRouteTargetsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNamen(namen *string) *IpamRouteTargetsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameNic(nameNic *string) *IpamRouteTargetsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameNie(nameNie *string) *IpamRouteTargetsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameNiew(nameNiew *string) *IpamRouteTargetsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithNameNisw(nameNisw *string) *IpamRouteTargetsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithOffset(offset *int64) *IpamRouteTargetsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithQ(q *string) *IpamRouteTargetsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTag(tag *string) *IpamRouteTargetsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTagn(tagn *string) *IpamRouteTargetsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenant(tenant *string) *IpamRouteTargetsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantn(tenantn *string) *IpamRouteTargetsListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantGroup(tenantGroup *string) *IpamRouteTargetsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantGroupn(tenantGroupn *string) *IpamRouteTargetsListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantGroupID(tenantGroupID *string) *IpamRouteTargetsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamRouteTargetsListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantID(tenantID *string) *IpamRouteTargetsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam route targets list params
func (o *IpamRouteTargetsListParams) WithTenantIDn(tenantIDn *string) *IpamRouteTargetsListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam route targets list params
func (o *IpamRouteTargetsListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ExportingVrf != nil {

		// query param exporting_vrf
		var qrExportingVrf string

		if o.ExportingVrf != nil {
			qrExportingVrf = *o.ExportingVrf
		}
		qExportingVrf := qrExportingVrf
		if qExportingVrf != "" {

			if err := r.SetQueryParam("exporting_vrf", qExportingVrf); err != nil {
				return err
			}
		}
	}

	if o.ExportingVrfn != nil {

		// query param exporting_vrf__n
		var qrExportingVrfn string

		if o.ExportingVrfn != nil {
			qrExportingVrfn = *o.ExportingVrfn
		}
		qExportingVrfn := qrExportingVrfn
		if qExportingVrfn != "" {

			if err := r.SetQueryParam("exporting_vrf__n", qExportingVrfn); err != nil {
				return err
			}
		}
	}

	if o.ExportingVrfID != nil {

		// query param exporting_vrf_id
		var qrExportingVrfID string

		if o.ExportingVrfID != nil {
			qrExportingVrfID = *o.ExportingVrfID
		}
		qExportingVrfID := qrExportingVrfID
		if qExportingVrfID != "" {

			if err := r.SetQueryParam("exporting_vrf_id", qExportingVrfID); err != nil {
				return err
			}
		}
	}

	if o.ExportingVrfIDn != nil {

		// query param exporting_vrf_id__n
		var qrExportingVrfIDn string

		if o.ExportingVrfIDn != nil {
			qrExportingVrfIDn = *o.ExportingVrfIDn
		}
		qExportingVrfIDn := qrExportingVrfIDn
		if qExportingVrfIDn != "" {

			if err := r.SetQueryParam("exporting_vrf_id__n", qExportingVrfIDn); err != nil {
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

	if o.ImportingVrf != nil {

		// query param importing_vrf
		var qrImportingVrf string

		if o.ImportingVrf != nil {
			qrImportingVrf = *o.ImportingVrf
		}
		qImportingVrf := qrImportingVrf
		if qImportingVrf != "" {

			if err := r.SetQueryParam("importing_vrf", qImportingVrf); err != nil {
				return err
			}
		}
	}

	if o.ImportingVrfn != nil {

		// query param importing_vrf__n
		var qrImportingVrfn string

		if o.ImportingVrfn != nil {
			qrImportingVrfn = *o.ImportingVrfn
		}
		qImportingVrfn := qrImportingVrfn
		if qImportingVrfn != "" {

			if err := r.SetQueryParam("importing_vrf__n", qImportingVrfn); err != nil {
				return err
			}
		}
	}

	if o.ImportingVrfID != nil {

		// query param importing_vrf_id
		var qrImportingVrfID string

		if o.ImportingVrfID != nil {
			qrImportingVrfID = *o.ImportingVrfID
		}
		qImportingVrfID := qrImportingVrfID
		if qImportingVrfID != "" {

			if err := r.SetQueryParam("importing_vrf_id", qImportingVrfID); err != nil {
				return err
			}
		}
	}

	if o.ImportingVrfIDn != nil {

		// query param importing_vrf_id__n
		var qrImportingVrfIDn string

		if o.ImportingVrfIDn != nil {
			qrImportingVrfIDn = *o.ImportingVrfIDn
		}
		qImportingVrfIDn := qrImportingVrfIDn
		if qImportingVrfIDn != "" {

			if err := r.SetQueryParam("importing_vrf_id__n", qImportingVrfIDn); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string

		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {

			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string

		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {

			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupn != nil {

		// query param tenant_group__n
		var qrTenantGroupn string

		if o.TenantGroupn != nil {
			qrTenantGroupn = *o.TenantGroupn
		}
		qTenantGroupn := qrTenantGroupn
		if qTenantGroupn != "" {

			if err := r.SetQueryParam("tenant_group__n", qTenantGroupn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string

		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {

			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupIDn != nil {

		// query param tenant_group_id__n
		var qrTenantGroupIDn string

		if o.TenantGroupIDn != nil {
			qrTenantGroupIDn = *o.TenantGroupIDn
		}
		qTenantGroupIDn := qrTenantGroupIDn
		if qTenantGroupIDn != "" {

			if err := r.SetQueryParam("tenant_group_id__n", qTenantGroupIDn); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
				return err
			}
		}
	}

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string

		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {

			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

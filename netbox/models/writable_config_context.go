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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableConfigContext writable config context
//
// swagger:model WritableConfigContext
type WritableConfigContext struct {

	// cluster groups
	// Unique: true
	ClusterGroups []int64 `json:"cluster_groups"`

	// cluster types
	// Unique: true
	ClusterTypes []int64 `json:"cluster_types"`

	// clusters
	// Unique: true
	Clusters []int64 `json:"clusters"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Data
	// Required: true
	Data interface{} `json:"data"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device types
	// Unique: true
	DeviceTypes []int64 `json:"device_types"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is active
	IsActive bool `json:"is_active,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// locations
	// Unique: true
	Locations []int64 `json:"locations"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// platforms
	// Unique: true
	Platforms []int64 `json:"platforms"`

	// regions
	// Unique: true
	Regions []int64 `json:"regions"`

	// roles
	// Unique: true
	Roles []int64 `json:"roles"`

	// site groups
	// Unique: true
	SiteGroups []int64 `json:"site_groups"`

	// sites
	// Unique: true
	Sites []int64 `json:"sites"`

	// tags
	// Unique: true
	Tags []string `json:"tags"`

	// tenant groups
	// Unique: true
	TenantGroups []int64 `json:"tenant_groups"`

	// tenants
	// Unique: true
	Tenants []int64 `json:"tenants"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Weight
	// Maximum: 32767
	// Minimum: 0
	Weight *int64 `json:"weight,omitempty"`
}

// Validate validates this writable config context
func (m *WritableConfigContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSiteGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSites(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConfigContext) validateClusterGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("cluster_groups", "body", m.ClusterGroups); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateClusterTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("cluster_types", "body", m.ClusterTypes); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateClusters(formats strfmt.Registry) error {
	if swag.IsZero(m.Clusters) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusters", "body", m.Clusters); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateData(formats strfmt.Registry) error {

	if m.Data == nil {
		return errors.Required("data", "body", nil)
	}

	return nil
}

func (m *WritableConfigContext) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateDeviceTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceTypes) { // not required
		return nil
	}

	if err := validate.UniqueItems("device_types", "body", m.DeviceTypes); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateLocations(formats strfmt.Registry) error {
	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	if err := validate.UniqueItems("locations", "body", m.Locations); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validatePlatforms(formats strfmt.Registry) error {
	if swag.IsZero(m.Platforms) { // not required
		return nil
	}

	if err := validate.UniqueItems("platforms", "body", m.Platforms); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateRegions(formats strfmt.Registry) error {
	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateRoles(formats strfmt.Registry) error {
	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if err := validate.UniqueItems("roles", "body", m.Roles); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateSiteGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.SiteGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("site_groups", "body", m.SiteGroups); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateSites(formats strfmt.Registry) error {
	if swag.IsZero(m.Sites) { // not required
		return nil
	}

	if err := validate.UniqueItems("sites", "body", m.Sites); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if err := validate.UniqueItems("tags", "body", m.Tags); err != nil {
		return err
	}

	for i := 0; i < len(m.Tags); i++ {

		if err := validate.Pattern("tags"+"."+strconv.Itoa(i), "body", m.Tags[i], `^[-\w]+$`); err != nil {
			return err
		}

	}

	return nil
}

func (m *WritableConfigContext) validateTenantGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.TenantGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenant_groups", "body", m.TenantGroups); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateTenants(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenants) { // not required
		return nil
	}

	if err := validate.UniqueItems("tenants", "body", m.Tenants); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) validateWeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Weight) { // not required
		return nil
	}

	if err := validate.MinimumInt("weight", "body", *m.Weight, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("weight", "body", *m.Weight, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable config context based on the context it is used
func (m *WritableConfigContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableConfigContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *WritableConfigContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableConfigContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableConfigContext) UnmarshalBinary(b []byte) error {
	var res WritableConfigContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

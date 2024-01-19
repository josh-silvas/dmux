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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritablePowerFeed writable power feed
//
// swagger:model WritablePowerFeed
type WritablePowerFeed struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// Amperage
	// Maximum: 32767
	// Minimum: 1
	Amperage int64 `json:"amperage,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable end
	// Read Only: true
	// Min Length: 1
	CableEnd string `json:"cable_end,omitempty"`

	// Comments
	Comments string `json:"comments,omitempty"`

	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoints []*string `json:"connected_endpoints"`

	// Connected endpoints reachable
	// Read Only: true
	ConnectedEndpointsReachable *bool `json:"connected_endpoints_reachable,omitempty"`

	// Connected endpoints type
	// Read Only: true
	ConnectedEndpointsType string `json:"connected_endpoints_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	//
	// Return the appropriate serializer for the link termination model.
	//
	// Read Only: true
	LinkPeers []*string `json:"link_peers"`

	// Link peers type
	// Read Only: true
	LinkPeersType string `json:"link_peers_type,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Max utilization
	//
	// Maximum permissible draw (percentage)
	// Maximum: 100
	// Minimum: 1
	MaxUtilization int64 `json:"max_utilization,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Phase
	// Enum: [single-phase three-phase]
	Phase string `json:"phase,omitempty"`

	// Power panel
	// Required: true
	PowerPanel *int64 `json:"power_panel"`

	// Rack
	Rack *int64 `json:"rack,omitempty"`

	// Status
	// Enum: [offline active planned failed]
	Status string `json:"status,omitempty"`

	// Supply
	// Enum: [ac dc]
	Supply string `json:"supply,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// Type
	// Enum: [primary redundant]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Voltage
	// Maximum: 32767
	// Minimum: -32768
	Voltage *int64 `json:"voltage,omitempty"`
}

// Validate validates this writable power feed
func (m *WritablePowerFeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmperage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCableEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxUtilization(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePowerPanel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoltage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePowerFeed) validateAmperage(formats strfmt.Registry) error {
	if swag.IsZero(m.Amperage) { // not required
		return nil
	}

	if err := validate.MinimumInt("amperage", "body", m.Amperage, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("amperage", "body", m.Amperage, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritablePowerFeed) validateCableEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.CableEnd) { // not required
		return nil
	}

	if err := validate.MinLength("cable_end", "body", m.CableEnd, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateMaxUtilization(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxUtilization) { // not required
		return nil
	}

	if err := validate.MinimumInt("max_utilization", "body", m.MaxUtilization, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max_utilization", "body", m.MaxUtilization, 100, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateName(formats strfmt.Registry) error {

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

var writablePowerFeedTypePhasePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["single-phase","three-phase"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypePhasePropEnum = append(writablePowerFeedTypePhasePropEnum, v)
	}
}

const (

	// WritablePowerFeedPhaseSingleDashPhase captures enum value "single-phase"
	WritablePowerFeedPhaseSingleDashPhase string = "single-phase"

	// WritablePowerFeedPhaseThreeDashPhase captures enum value "three-phase"
	WritablePowerFeedPhaseThreeDashPhase string = "three-phase"
)

// prop value enum
func (m *WritablePowerFeed) validatePhaseEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypePhasePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validatePhase(formats strfmt.Registry) error {
	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhaseEnum("phase", "body", m.Phase); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validatePowerPanel(formats strfmt.Registry) error {

	if err := validate.Required("power_panel", "body", m.PowerPanel); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["offline","active","planned","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeStatusPropEnum = append(writablePowerFeedTypeStatusPropEnum, v)
	}
}

const (

	// WritablePowerFeedStatusOffline captures enum value "offline"
	WritablePowerFeedStatusOffline string = "offline"

	// WritablePowerFeedStatusActive captures enum value "active"
	WritablePowerFeedStatusActive string = "active"

	// WritablePowerFeedStatusPlanned captures enum value "planned"
	WritablePowerFeedStatusPlanned string = "planned"

	// WritablePowerFeedStatusFailed captures enum value "failed"
	WritablePowerFeedStatusFailed string = "failed"
)

// prop value enum
func (m *WritablePowerFeed) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var writablePowerFeedTypeSupplyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ac","dc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeSupplyPropEnum = append(writablePowerFeedTypeSupplyPropEnum, v)
	}
}

const (

	// WritablePowerFeedSupplyAc captures enum value "ac"
	WritablePowerFeedSupplyAc string = "ac"

	// WritablePowerFeedSupplyDc captures enum value "dc"
	WritablePowerFeedSupplyDc string = "dc"
)

// prop value enum
func (m *WritablePowerFeed) validateSupplyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypeSupplyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateSupply(formats strfmt.Registry) error {
	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	// value enum
	if err := m.validateSupplyEnum("supply", "body", m.Supply); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var writablePowerFeedTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","redundant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerFeedTypeTypePropEnum = append(writablePowerFeedTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerFeedTypePrimary captures enum value "primary"
	WritablePowerFeedTypePrimary string = "primary"

	// WritablePowerFeedTypeRedundant captures enum value "redundant"
	WritablePowerFeedTypeRedundant string = "redundant"
)

// prop value enum
func (m *WritablePowerFeed) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerFeedTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerFeed) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) validateVoltage(formats strfmt.Registry) error {
	if swag.IsZero(m.Voltage) { // not required
		return nil
	}

	if err := validate.MinimumInt("voltage", "body", *m.Voltage, -32768, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("voltage", "body", *m.Voltage, 32767, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power feed based on the context it is used
func (m *WritablePowerFeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCableEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointsType(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateLinkPeers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLinkPeersType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
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

func (m *WritablePowerFeed) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateCableEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_end", "body", string(m.CableEnd)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateConnectedEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints", "body", []*string(m.ConnectedEndpoints)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateConnectedEndpointsReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_reachable", "body", m.ConnectedEndpointsReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateConnectedEndpointsType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_type", "body", string(m.ConnectedEndpointsType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateLinkPeers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers", "body", []*string(m.LinkPeers)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateLinkPeersType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers_type", "body", string(m.LinkPeersType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerFeed) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritablePowerFeed) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerFeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerFeed) UnmarshalBinary(b []byte) error {
	var res WritablePowerFeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

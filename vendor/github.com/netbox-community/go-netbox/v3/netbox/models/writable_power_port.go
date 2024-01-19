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

// WritablePowerPort writable power port
//
// swagger:model WritablePowerPort
type WritablePowerPort struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// Allocated draw
	//
	// Allocated power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	AllocatedDraw *int64 `json:"allocated_draw,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable end
	// Read Only: true
	// Min Length: 1
	CableEnd string `json:"cable_end,omitempty"`

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

	// Device
	// Required: true
	Device *int64 `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

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

	// Maximum draw
	//
	// Maximum power draw (watts)
	// Maximum: 32767
	// Minimum: 1
	MaximumDraw *int64 `json:"maximum_draw,omitempty"`

	// Module
	Module *int64 `json:"module,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// Type
	//
	// Physical port type
	// Enum: [iec-60320-c6 iec-60320-c8 iec-60320-c14 iec-60320-c16 iec-60320-c20 iec-60320-c22 iec-60309-p-n-e-4h iec-60309-p-n-e-6h iec-60309-p-n-e-9h iec-60309-2p-e-4h iec-60309-2p-e-6h iec-60309-2p-e-9h iec-60309-3p-e-4h iec-60309-3p-e-6h iec-60309-3p-e-9h iec-60309-3p-n-e-4h iec-60309-3p-n-e-6h iec-60309-3p-n-e-9h nema-1-15p nema-5-15p nema-5-20p nema-5-30p nema-5-50p nema-6-15p nema-6-20p nema-6-30p nema-6-50p nema-10-30p nema-10-50p nema-14-20p nema-14-30p nema-14-50p nema-14-60p nema-15-15p nema-15-20p nema-15-30p nema-15-50p nema-15-60p nema-l1-15p nema-l5-15p nema-l5-20p nema-l5-30p nema-l5-50p nema-l6-15p nema-l6-20p nema-l6-30p nema-l6-50p nema-l10-30p nema-l14-20p nema-l14-30p nema-l14-50p nema-l14-60p nema-l15-20p nema-l15-30p nema-l15-50p nema-l15-60p nema-l21-20p nema-l21-30p nema-l22-30p cs6361c cs6365c cs8165c cs8265c cs8365c cs8465c ita-c ita-e ita-f ita-ef ita-g ita-h ita-i ita-j ita-k ita-l ita-m ita-n ita-o usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b usb-micro-ab usb-3-b usb-3-micro-b dc-terminal saf-d-grid neutrik-powercon-20 neutrik-powercon-32 neutrik-powercon-true1 neutrik-powercon-true1-top ubiquiti-smartpower hardwired other]
	Type string `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable power port
func (m *WritablePowerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocatedDraw(formats); err != nil {
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

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaximumDraw(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritablePowerPort) validateAllocatedDraw(formats strfmt.Registry) error {
	if swag.IsZero(m.AllocatedDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("allocated_draw", "body", *m.AllocatedDraw, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("allocated_draw", "body", *m.AllocatedDraw, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateCable(formats strfmt.Registry) error {
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

func (m *WritablePowerPort) validateCableEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.CableEnd) { // not required
		return nil
	}

	if err := validate.MinLength("cable_end", "body", m.CableEnd, 1); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateMaximumDraw(formats strfmt.Registry) error {
	if swag.IsZero(m.MaximumDraw) { // not required
		return nil
	}

	if err := validate.MinimumInt("maximum_draw", "body", *m.MaximumDraw, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("maximum_draw", "body", *m.MaximumDraw, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateTags(formats strfmt.Registry) error {
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

var writablePowerPortTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["iec-60320-c6","iec-60320-c8","iec-60320-c14","iec-60320-c16","iec-60320-c20","iec-60320-c22","iec-60309-p-n-e-4h","iec-60309-p-n-e-6h","iec-60309-p-n-e-9h","iec-60309-2p-e-4h","iec-60309-2p-e-6h","iec-60309-2p-e-9h","iec-60309-3p-e-4h","iec-60309-3p-e-6h","iec-60309-3p-e-9h","iec-60309-3p-n-e-4h","iec-60309-3p-n-e-6h","iec-60309-3p-n-e-9h","nema-1-15p","nema-5-15p","nema-5-20p","nema-5-30p","nema-5-50p","nema-6-15p","nema-6-20p","nema-6-30p","nema-6-50p","nema-10-30p","nema-10-50p","nema-14-20p","nema-14-30p","nema-14-50p","nema-14-60p","nema-15-15p","nema-15-20p","nema-15-30p","nema-15-50p","nema-15-60p","nema-l1-15p","nema-l5-15p","nema-l5-20p","nema-l5-30p","nema-l5-50p","nema-l6-15p","nema-l6-20p","nema-l6-30p","nema-l6-50p","nema-l10-30p","nema-l14-20p","nema-l14-30p","nema-l14-50p","nema-l14-60p","nema-l15-20p","nema-l15-30p","nema-l15-50p","nema-l15-60p","nema-l21-20p","nema-l21-30p","nema-l22-30p","cs6361c","cs6365c","cs8165c","cs8265c","cs8365c","cs8465c","ita-c","ita-e","ita-f","ita-ef","ita-g","ita-h","ita-i","ita-j","ita-k","ita-l","ita-m","ita-n","ita-o","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","usb-micro-ab","usb-3-b","usb-3-micro-b","dc-terminal","saf-d-grid","neutrik-powercon-20","neutrik-powercon-32","neutrik-powercon-true1","neutrik-powercon-true1-top","ubiquiti-smartpower","hardwired","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writablePowerPortTypeTypePropEnum = append(writablePowerPortTypeTypePropEnum, v)
	}
}

const (

	// WritablePowerPortTypeIecDash60320DashC6 captures enum value "iec-60320-c6"
	WritablePowerPortTypeIecDash60320DashC6 string = "iec-60320-c6"

	// WritablePowerPortTypeIecDash60320DashC8 captures enum value "iec-60320-c8"
	WritablePowerPortTypeIecDash60320DashC8 string = "iec-60320-c8"

	// WritablePowerPortTypeIecDash60320DashC14 captures enum value "iec-60320-c14"
	WritablePowerPortTypeIecDash60320DashC14 string = "iec-60320-c14"

	// WritablePowerPortTypeIecDash60320DashC16 captures enum value "iec-60320-c16"
	WritablePowerPortTypeIecDash60320DashC16 string = "iec-60320-c16"

	// WritablePowerPortTypeIecDash60320DashC20 captures enum value "iec-60320-c20"
	WritablePowerPortTypeIecDash60320DashC20 string = "iec-60320-c20"

	// WritablePowerPortTypeIecDash60320DashC22 captures enum value "iec-60320-c22"
	WritablePowerPortTypeIecDash60320DashC22 string = "iec-60320-c22"

	// WritablePowerPortTypeIecDash60309DashpDashnDasheDash4h captures enum value "iec-60309-p-n-e-4h"
	WritablePowerPortTypeIecDash60309DashpDashnDasheDash4h string = "iec-60309-p-n-e-4h"

	// WritablePowerPortTypeIecDash60309DashpDashnDasheDash6h captures enum value "iec-60309-p-n-e-6h"
	WritablePowerPortTypeIecDash60309DashpDashnDasheDash6h string = "iec-60309-p-n-e-6h"

	// WritablePowerPortTypeIecDash60309DashpDashnDasheDash9h captures enum value "iec-60309-p-n-e-9h"
	WritablePowerPortTypeIecDash60309DashpDashnDasheDash9h string = "iec-60309-p-n-e-9h"

	// WritablePowerPortTypeIecDash60309Dash2pDasheDash4h captures enum value "iec-60309-2p-e-4h"
	WritablePowerPortTypeIecDash60309Dash2pDasheDash4h string = "iec-60309-2p-e-4h"

	// WritablePowerPortTypeIecDash60309Dash2pDasheDash6h captures enum value "iec-60309-2p-e-6h"
	WritablePowerPortTypeIecDash60309Dash2pDasheDash6h string = "iec-60309-2p-e-6h"

	// WritablePowerPortTypeIecDash60309Dash2pDasheDash9h captures enum value "iec-60309-2p-e-9h"
	WritablePowerPortTypeIecDash60309Dash2pDasheDash9h string = "iec-60309-2p-e-9h"

	// WritablePowerPortTypeIecDash60309Dash3pDasheDash4h captures enum value "iec-60309-3p-e-4h"
	WritablePowerPortTypeIecDash60309Dash3pDasheDash4h string = "iec-60309-3p-e-4h"

	// WritablePowerPortTypeIecDash60309Dash3pDasheDash6h captures enum value "iec-60309-3p-e-6h"
	WritablePowerPortTypeIecDash60309Dash3pDasheDash6h string = "iec-60309-3p-e-6h"

	// WritablePowerPortTypeIecDash60309Dash3pDasheDash9h captures enum value "iec-60309-3p-e-9h"
	WritablePowerPortTypeIecDash60309Dash3pDasheDash9h string = "iec-60309-3p-e-9h"

	// WritablePowerPortTypeIecDash60309Dash3pDashnDasheDash4h captures enum value "iec-60309-3p-n-e-4h"
	WritablePowerPortTypeIecDash60309Dash3pDashnDasheDash4h string = "iec-60309-3p-n-e-4h"

	// WritablePowerPortTypeIecDash60309Dash3pDashnDasheDash6h captures enum value "iec-60309-3p-n-e-6h"
	WritablePowerPortTypeIecDash60309Dash3pDashnDasheDash6h string = "iec-60309-3p-n-e-6h"

	// WritablePowerPortTypeIecDash60309Dash3pDashnDasheDash9h captures enum value "iec-60309-3p-n-e-9h"
	WritablePowerPortTypeIecDash60309Dash3pDashnDasheDash9h string = "iec-60309-3p-n-e-9h"

	// WritablePowerPortTypeNemaDash1Dash15p captures enum value "nema-1-15p"
	WritablePowerPortTypeNemaDash1Dash15p string = "nema-1-15p"

	// WritablePowerPortTypeNemaDash5Dash15p captures enum value "nema-5-15p"
	WritablePowerPortTypeNemaDash5Dash15p string = "nema-5-15p"

	// WritablePowerPortTypeNemaDash5Dash20p captures enum value "nema-5-20p"
	WritablePowerPortTypeNemaDash5Dash20p string = "nema-5-20p"

	// WritablePowerPortTypeNemaDash5Dash30p captures enum value "nema-5-30p"
	WritablePowerPortTypeNemaDash5Dash30p string = "nema-5-30p"

	// WritablePowerPortTypeNemaDash5Dash50p captures enum value "nema-5-50p"
	WritablePowerPortTypeNemaDash5Dash50p string = "nema-5-50p"

	// WritablePowerPortTypeNemaDash6Dash15p captures enum value "nema-6-15p"
	WritablePowerPortTypeNemaDash6Dash15p string = "nema-6-15p"

	// WritablePowerPortTypeNemaDash6Dash20p captures enum value "nema-6-20p"
	WritablePowerPortTypeNemaDash6Dash20p string = "nema-6-20p"

	// WritablePowerPortTypeNemaDash6Dash30p captures enum value "nema-6-30p"
	WritablePowerPortTypeNemaDash6Dash30p string = "nema-6-30p"

	// WritablePowerPortTypeNemaDash6Dash50p captures enum value "nema-6-50p"
	WritablePowerPortTypeNemaDash6Dash50p string = "nema-6-50p"

	// WritablePowerPortTypeNemaDash10Dash30p captures enum value "nema-10-30p"
	WritablePowerPortTypeNemaDash10Dash30p string = "nema-10-30p"

	// WritablePowerPortTypeNemaDash10Dash50p captures enum value "nema-10-50p"
	WritablePowerPortTypeNemaDash10Dash50p string = "nema-10-50p"

	// WritablePowerPortTypeNemaDash14Dash20p captures enum value "nema-14-20p"
	WritablePowerPortTypeNemaDash14Dash20p string = "nema-14-20p"

	// WritablePowerPortTypeNemaDash14Dash30p captures enum value "nema-14-30p"
	WritablePowerPortTypeNemaDash14Dash30p string = "nema-14-30p"

	// WritablePowerPortTypeNemaDash14Dash50p captures enum value "nema-14-50p"
	WritablePowerPortTypeNemaDash14Dash50p string = "nema-14-50p"

	// WritablePowerPortTypeNemaDash14Dash60p captures enum value "nema-14-60p"
	WritablePowerPortTypeNemaDash14Dash60p string = "nema-14-60p"

	// WritablePowerPortTypeNemaDash15Dash15p captures enum value "nema-15-15p"
	WritablePowerPortTypeNemaDash15Dash15p string = "nema-15-15p"

	// WritablePowerPortTypeNemaDash15Dash20p captures enum value "nema-15-20p"
	WritablePowerPortTypeNemaDash15Dash20p string = "nema-15-20p"

	// WritablePowerPortTypeNemaDash15Dash30p captures enum value "nema-15-30p"
	WritablePowerPortTypeNemaDash15Dash30p string = "nema-15-30p"

	// WritablePowerPortTypeNemaDash15Dash50p captures enum value "nema-15-50p"
	WritablePowerPortTypeNemaDash15Dash50p string = "nema-15-50p"

	// WritablePowerPortTypeNemaDash15Dash60p captures enum value "nema-15-60p"
	WritablePowerPortTypeNemaDash15Dash60p string = "nema-15-60p"

	// WritablePowerPortTypeNemaDashL1Dash15p captures enum value "nema-l1-15p"
	WritablePowerPortTypeNemaDashL1Dash15p string = "nema-l1-15p"

	// WritablePowerPortTypeNemaDashL5Dash15p captures enum value "nema-l5-15p"
	WritablePowerPortTypeNemaDashL5Dash15p string = "nema-l5-15p"

	// WritablePowerPortTypeNemaDashL5Dash20p captures enum value "nema-l5-20p"
	WritablePowerPortTypeNemaDashL5Dash20p string = "nema-l5-20p"

	// WritablePowerPortTypeNemaDashL5Dash30p captures enum value "nema-l5-30p"
	WritablePowerPortTypeNemaDashL5Dash30p string = "nema-l5-30p"

	// WritablePowerPortTypeNemaDashL5Dash50p captures enum value "nema-l5-50p"
	WritablePowerPortTypeNemaDashL5Dash50p string = "nema-l5-50p"

	// WritablePowerPortTypeNemaDashL6Dash15p captures enum value "nema-l6-15p"
	WritablePowerPortTypeNemaDashL6Dash15p string = "nema-l6-15p"

	// WritablePowerPortTypeNemaDashL6Dash20p captures enum value "nema-l6-20p"
	WritablePowerPortTypeNemaDashL6Dash20p string = "nema-l6-20p"

	// WritablePowerPortTypeNemaDashL6Dash30p captures enum value "nema-l6-30p"
	WritablePowerPortTypeNemaDashL6Dash30p string = "nema-l6-30p"

	// WritablePowerPortTypeNemaDashL6Dash50p captures enum value "nema-l6-50p"
	WritablePowerPortTypeNemaDashL6Dash50p string = "nema-l6-50p"

	// WritablePowerPortTypeNemaDashL10Dash30p captures enum value "nema-l10-30p"
	WritablePowerPortTypeNemaDashL10Dash30p string = "nema-l10-30p"

	// WritablePowerPortTypeNemaDashL14Dash20p captures enum value "nema-l14-20p"
	WritablePowerPortTypeNemaDashL14Dash20p string = "nema-l14-20p"

	// WritablePowerPortTypeNemaDashL14Dash30p captures enum value "nema-l14-30p"
	WritablePowerPortTypeNemaDashL14Dash30p string = "nema-l14-30p"

	// WritablePowerPortTypeNemaDashL14Dash50p captures enum value "nema-l14-50p"
	WritablePowerPortTypeNemaDashL14Dash50p string = "nema-l14-50p"

	// WritablePowerPortTypeNemaDashL14Dash60p captures enum value "nema-l14-60p"
	WritablePowerPortTypeNemaDashL14Dash60p string = "nema-l14-60p"

	// WritablePowerPortTypeNemaDashL15Dash20p captures enum value "nema-l15-20p"
	WritablePowerPortTypeNemaDashL15Dash20p string = "nema-l15-20p"

	// WritablePowerPortTypeNemaDashL15Dash30p captures enum value "nema-l15-30p"
	WritablePowerPortTypeNemaDashL15Dash30p string = "nema-l15-30p"

	// WritablePowerPortTypeNemaDashL15Dash50p captures enum value "nema-l15-50p"
	WritablePowerPortTypeNemaDashL15Dash50p string = "nema-l15-50p"

	// WritablePowerPortTypeNemaDashL15Dash60p captures enum value "nema-l15-60p"
	WritablePowerPortTypeNemaDashL15Dash60p string = "nema-l15-60p"

	// WritablePowerPortTypeNemaDashL21Dash20p captures enum value "nema-l21-20p"
	WritablePowerPortTypeNemaDashL21Dash20p string = "nema-l21-20p"

	// WritablePowerPortTypeNemaDashL21Dash30p captures enum value "nema-l21-30p"
	WritablePowerPortTypeNemaDashL21Dash30p string = "nema-l21-30p"

	// WritablePowerPortTypeNemaDashL22Dash30p captures enum value "nema-l22-30p"
	WritablePowerPortTypeNemaDashL22Dash30p string = "nema-l22-30p"

	// WritablePowerPortTypeCs6361c captures enum value "cs6361c"
	WritablePowerPortTypeCs6361c string = "cs6361c"

	// WritablePowerPortTypeCs6365c captures enum value "cs6365c"
	WritablePowerPortTypeCs6365c string = "cs6365c"

	// WritablePowerPortTypeCs8165c captures enum value "cs8165c"
	WritablePowerPortTypeCs8165c string = "cs8165c"

	// WritablePowerPortTypeCs8265c captures enum value "cs8265c"
	WritablePowerPortTypeCs8265c string = "cs8265c"

	// WritablePowerPortTypeCs8365c captures enum value "cs8365c"
	WritablePowerPortTypeCs8365c string = "cs8365c"

	// WritablePowerPortTypeCs8465c captures enum value "cs8465c"
	WritablePowerPortTypeCs8465c string = "cs8465c"

	// WritablePowerPortTypeItaDashc captures enum value "ita-c"
	WritablePowerPortTypeItaDashc string = "ita-c"

	// WritablePowerPortTypeItaDashe captures enum value "ita-e"
	WritablePowerPortTypeItaDashe string = "ita-e"

	// WritablePowerPortTypeItaDashf captures enum value "ita-f"
	WritablePowerPortTypeItaDashf string = "ita-f"

	// WritablePowerPortTypeItaDashEf captures enum value "ita-ef"
	WritablePowerPortTypeItaDashEf string = "ita-ef"

	// WritablePowerPortTypeItaDashg captures enum value "ita-g"
	WritablePowerPortTypeItaDashg string = "ita-g"

	// WritablePowerPortTypeItaDashh captures enum value "ita-h"
	WritablePowerPortTypeItaDashh string = "ita-h"

	// WritablePowerPortTypeItaDashi captures enum value "ita-i"
	WritablePowerPortTypeItaDashi string = "ita-i"

	// WritablePowerPortTypeItaDashj captures enum value "ita-j"
	WritablePowerPortTypeItaDashj string = "ita-j"

	// WritablePowerPortTypeItaDashk captures enum value "ita-k"
	WritablePowerPortTypeItaDashk string = "ita-k"

	// WritablePowerPortTypeItaDashl captures enum value "ita-l"
	WritablePowerPortTypeItaDashl string = "ita-l"

	// WritablePowerPortTypeItaDashm captures enum value "ita-m"
	WritablePowerPortTypeItaDashm string = "ita-m"

	// WritablePowerPortTypeItaDashn captures enum value "ita-n"
	WritablePowerPortTypeItaDashn string = "ita-n"

	// WritablePowerPortTypeItaDasho captures enum value "ita-o"
	WritablePowerPortTypeItaDasho string = "ita-o"

	// WritablePowerPortTypeUsbDasha captures enum value "usb-a"
	WritablePowerPortTypeUsbDasha string = "usb-a"

	// WritablePowerPortTypeUsbDashb captures enum value "usb-b"
	WritablePowerPortTypeUsbDashb string = "usb-b"

	// WritablePowerPortTypeUsbDashc captures enum value "usb-c"
	WritablePowerPortTypeUsbDashc string = "usb-c"

	// WritablePowerPortTypeUsbDashMiniDasha captures enum value "usb-mini-a"
	WritablePowerPortTypeUsbDashMiniDasha string = "usb-mini-a"

	// WritablePowerPortTypeUsbDashMiniDashb captures enum value "usb-mini-b"
	WritablePowerPortTypeUsbDashMiniDashb string = "usb-mini-b"

	// WritablePowerPortTypeUsbDashMicroDasha captures enum value "usb-micro-a"
	WritablePowerPortTypeUsbDashMicroDasha string = "usb-micro-a"

	// WritablePowerPortTypeUsbDashMicroDashb captures enum value "usb-micro-b"
	WritablePowerPortTypeUsbDashMicroDashb string = "usb-micro-b"

	// WritablePowerPortTypeUsbDashMicroDashAb captures enum value "usb-micro-ab"
	WritablePowerPortTypeUsbDashMicroDashAb string = "usb-micro-ab"

	// WritablePowerPortTypeUsbDash3Dashb captures enum value "usb-3-b"
	WritablePowerPortTypeUsbDash3Dashb string = "usb-3-b"

	// WritablePowerPortTypeUsbDash3DashMicroDashb captures enum value "usb-3-micro-b"
	WritablePowerPortTypeUsbDash3DashMicroDashb string = "usb-3-micro-b"

	// WritablePowerPortTypeDcDashTerminal captures enum value "dc-terminal"
	WritablePowerPortTypeDcDashTerminal string = "dc-terminal"

	// WritablePowerPortTypeSafDashdDashGrid captures enum value "saf-d-grid"
	WritablePowerPortTypeSafDashdDashGrid string = "saf-d-grid"

	// WritablePowerPortTypeNeutrikDashPowerconDash20 captures enum value "neutrik-powercon-20"
	WritablePowerPortTypeNeutrikDashPowerconDash20 string = "neutrik-powercon-20"

	// WritablePowerPortTypeNeutrikDashPowerconDash32 captures enum value "neutrik-powercon-32"
	WritablePowerPortTypeNeutrikDashPowerconDash32 string = "neutrik-powercon-32"

	// WritablePowerPortTypeNeutrikDashPowerconDashTrue1 captures enum value "neutrik-powercon-true1"
	WritablePowerPortTypeNeutrikDashPowerconDashTrue1 string = "neutrik-powercon-true1"

	// WritablePowerPortTypeNeutrikDashPowerconDashTrue1DashTop captures enum value "neutrik-powercon-true1-top"
	WritablePowerPortTypeNeutrikDashPowerconDashTrue1DashTop string = "neutrik-powercon-true1-top"

	// WritablePowerPortTypeUbiquitiDashSmartpower captures enum value "ubiquiti-smartpower"
	WritablePowerPortTypeUbiquitiDashSmartpower string = "ubiquiti-smartpower"

	// WritablePowerPortTypeHardwired captures enum value "hardwired"
	WritablePowerPortTypeHardwired string = "hardwired"

	// WritablePowerPortTypeOther captures enum value "other"
	WritablePowerPortTypeOther string = "other"
)

// prop value enum
func (m *WritablePowerPort) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writablePowerPortTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritablePowerPort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable power port based on the context it is used
func (m *WritablePowerPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *WritablePowerPort) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritablePowerPort) contextValidateCableEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_end", "body", string(m.CableEnd)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateConnectedEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints", "body", []*string(m.ConnectedEndpoints)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateConnectedEndpointsReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_reachable", "body", m.ConnectedEndpointsReachable); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateConnectedEndpointsType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoints_type", "body", string(m.ConnectedEndpointsType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateLinkPeers(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers", "body", []*string(m.LinkPeers)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateLinkPeersType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "link_peers_type", "body", string(m.LinkPeersType)); err != nil {
		return err
	}

	return nil
}

func (m *WritablePowerPort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *WritablePowerPort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritablePowerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritablePowerPort) UnmarshalBinary(b []byte) error {
	var res WritablePowerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

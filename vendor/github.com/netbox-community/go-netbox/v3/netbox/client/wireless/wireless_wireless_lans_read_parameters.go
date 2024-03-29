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

package wireless

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

// NewWirelessWirelessLansReadParams creates a new WirelessWirelessLansReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLansReadParams() *WirelessWirelessLansReadParams {
	return &WirelessWirelessLansReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLansReadParamsWithTimeout creates a new WirelessWirelessLansReadParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLansReadParamsWithTimeout(timeout time.Duration) *WirelessWirelessLansReadParams {
	return &WirelessWirelessLansReadParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLansReadParamsWithContext creates a new WirelessWirelessLansReadParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLansReadParamsWithContext(ctx context.Context) *WirelessWirelessLansReadParams {
	return &WirelessWirelessLansReadParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLansReadParamsWithHTTPClient creates a new WirelessWirelessLansReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLansReadParamsWithHTTPClient(client *http.Client) *WirelessWirelessLansReadParams {
	return &WirelessWirelessLansReadParams{
		HTTPClient: client,
	}
}

/*
WirelessWirelessLansReadParams contains all the parameters to send to the API endpoint

	for the wireless wireless lans read operation.

	Typically these are written to a http.Request.
*/
type WirelessWirelessLansReadParams struct {

	/* ID.

	   A unique integer value identifying this Wireless LAN.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless lans read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLansReadParams) WithDefaults() *WirelessWirelessLansReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless lans read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLansReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) WithTimeout(timeout time.Duration) *WirelessWirelessLansReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) WithContext(ctx context.Context) *WirelessWirelessLansReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) WithHTTPClient(client *http.Client) *WirelessWirelessLansReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) WithID(id int64) *WirelessWirelessLansReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the wireless wireless lans read params
func (o *WirelessWirelessLansReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLansReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

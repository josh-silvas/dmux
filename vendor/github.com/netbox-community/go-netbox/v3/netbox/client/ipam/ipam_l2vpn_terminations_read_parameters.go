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
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIpamL2vpnTerminationsReadParams creates a new IpamL2vpnTerminationsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnTerminationsReadParams() *IpamL2vpnTerminationsReadParams {
	return &IpamL2vpnTerminationsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnTerminationsReadParamsWithTimeout creates a new IpamL2vpnTerminationsReadParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnTerminationsReadParamsWithTimeout(timeout time.Duration) *IpamL2vpnTerminationsReadParams {
	return &IpamL2vpnTerminationsReadParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnTerminationsReadParamsWithContext creates a new IpamL2vpnTerminationsReadParams object
// with the ability to set a context for a request.
func NewIpamL2vpnTerminationsReadParamsWithContext(ctx context.Context) *IpamL2vpnTerminationsReadParams {
	return &IpamL2vpnTerminationsReadParams{
		Context: ctx,
	}
}

// NewIpamL2vpnTerminationsReadParamsWithHTTPClient creates a new IpamL2vpnTerminationsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnTerminationsReadParamsWithHTTPClient(client *http.Client) *IpamL2vpnTerminationsReadParams {
	return &IpamL2vpnTerminationsReadParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnTerminationsReadParams contains all the parameters to send to the API endpoint

	for the ipam l2vpn terminations read operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnTerminationsReadParams struct {

	/* ID.

	   A unique integer value identifying this L2VPN termination.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpn terminations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsReadParams) WithDefaults() *IpamL2vpnTerminationsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpn terminations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) WithTimeout(timeout time.Duration) *IpamL2vpnTerminationsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) WithContext(ctx context.Context) *IpamL2vpnTerminationsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) WithHTTPClient(client *http.Client) *IpamL2vpnTerminationsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) WithID(id int64) *IpamL2vpnTerminationsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam l2vpn terminations read params
func (o *IpamL2vpnTerminationsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnTerminationsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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

// NewDcimVirtualDeviceContextsReadParams creates a new DcimVirtualDeviceContextsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualDeviceContextsReadParams() *DcimVirtualDeviceContextsReadParams {
	return &DcimVirtualDeviceContextsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualDeviceContextsReadParamsWithTimeout creates a new DcimVirtualDeviceContextsReadParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualDeviceContextsReadParamsWithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsReadParams {
	return &DcimVirtualDeviceContextsReadParams{
		timeout: timeout,
	}
}

// NewDcimVirtualDeviceContextsReadParamsWithContext creates a new DcimVirtualDeviceContextsReadParams object
// with the ability to set a context for a request.
func NewDcimVirtualDeviceContextsReadParamsWithContext(ctx context.Context) *DcimVirtualDeviceContextsReadParams {
	return &DcimVirtualDeviceContextsReadParams{
		Context: ctx,
	}
}

// NewDcimVirtualDeviceContextsReadParamsWithHTTPClient creates a new DcimVirtualDeviceContextsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualDeviceContextsReadParamsWithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsReadParams {
	return &DcimVirtualDeviceContextsReadParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualDeviceContextsReadParams contains all the parameters to send to the API endpoint

	for the dcim virtual device contexts read operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualDeviceContextsReadParams struct {

	/* ID.

	   A unique integer value identifying this virtual device context.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual device contexts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsReadParams) WithDefaults() *DcimVirtualDeviceContextsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual device contexts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) WithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) WithContext(ctx context.Context) *DcimVirtualDeviceContextsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) WithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) WithID(id int64) *DcimVirtualDeviceContextsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual device contexts read params
func (o *DcimVirtualDeviceContextsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualDeviceContextsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

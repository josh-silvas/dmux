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

package tenancy

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
)

// NewTenancyContactGroupsBulkDeleteParams creates a new TenancyContactGroupsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyContactGroupsBulkDeleteParams() *TenancyContactGroupsBulkDeleteParams {
	return &TenancyContactGroupsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyContactGroupsBulkDeleteParamsWithTimeout creates a new TenancyContactGroupsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewTenancyContactGroupsBulkDeleteParamsWithTimeout(timeout time.Duration) *TenancyContactGroupsBulkDeleteParams {
	return &TenancyContactGroupsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewTenancyContactGroupsBulkDeleteParamsWithContext creates a new TenancyContactGroupsBulkDeleteParams object
// with the ability to set a context for a request.
func NewTenancyContactGroupsBulkDeleteParamsWithContext(ctx context.Context) *TenancyContactGroupsBulkDeleteParams {
	return &TenancyContactGroupsBulkDeleteParams{
		Context: ctx,
	}
}

// NewTenancyContactGroupsBulkDeleteParamsWithHTTPClient creates a new TenancyContactGroupsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyContactGroupsBulkDeleteParamsWithHTTPClient(client *http.Client) *TenancyContactGroupsBulkDeleteParams {
	return &TenancyContactGroupsBulkDeleteParams{
		HTTPClient: client,
	}
}

/*
TenancyContactGroupsBulkDeleteParams contains all the parameters to send to the API endpoint

	for the tenancy contact groups bulk delete operation.

	Typically these are written to a http.Request.
*/
type TenancyContactGroupsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy contact groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactGroupsBulkDeleteParams) WithDefaults() *TenancyContactGroupsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy contact groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyContactGroupsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy contact groups bulk delete params
func (o *TenancyContactGroupsBulkDeleteParams) WithTimeout(timeout time.Duration) *TenancyContactGroupsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy contact groups bulk delete params
func (o *TenancyContactGroupsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy contact groups bulk delete params
func (o *TenancyContactGroupsBulkDeleteParams) WithContext(ctx context.Context) *TenancyContactGroupsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy contact groups bulk delete params
func (o *TenancyContactGroupsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy contact groups bulk delete params
func (o *TenancyContactGroupsBulkDeleteParams) WithHTTPClient(client *http.Client) *TenancyContactGroupsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy contact groups bulk delete params
func (o *TenancyContactGroupsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyContactGroupsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

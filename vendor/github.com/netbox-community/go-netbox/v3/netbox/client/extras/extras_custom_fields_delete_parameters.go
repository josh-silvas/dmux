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

// NewExtrasCustomFieldsDeleteParams creates a new ExtrasCustomFieldsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsDeleteParams() *ExtrasCustomFieldsDeleteParams {
	return &ExtrasCustomFieldsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsDeleteParamsWithTimeout creates a new ExtrasCustomFieldsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsDeleteParams {
	return &ExtrasCustomFieldsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsDeleteParamsWithContext creates a new ExtrasCustomFieldsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsDeleteParamsWithContext(ctx context.Context) *ExtrasCustomFieldsDeleteParams {
	return &ExtrasCustomFieldsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsDeleteParamsWithHTTPClient creates a new ExtrasCustomFieldsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsDeleteParams {
	return &ExtrasCustomFieldsDeleteParams{
		HTTPClient: client,
	}
}

/*
ExtrasCustomFieldsDeleteParams contains all the parameters to send to the API endpoint

	for the extras custom fields delete operation.

	Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this custom field.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsDeleteParams) WithDefaults() *ExtrasCustomFieldsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) WithContext(ctx context.Context) *ExtrasCustomFieldsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) WithID(id int64) *ExtrasCustomFieldsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields delete params
func (o *ExtrasCustomFieldsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

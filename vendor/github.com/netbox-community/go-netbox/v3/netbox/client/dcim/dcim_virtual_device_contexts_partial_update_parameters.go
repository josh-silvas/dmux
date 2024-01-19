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

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// NewDcimVirtualDeviceContextsPartialUpdateParams creates a new DcimVirtualDeviceContextsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualDeviceContextsPartialUpdateParams() *DcimVirtualDeviceContextsPartialUpdateParams {
	return &DcimVirtualDeviceContextsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualDeviceContextsPartialUpdateParamsWithTimeout creates a new DcimVirtualDeviceContextsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualDeviceContextsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsPartialUpdateParams {
	return &DcimVirtualDeviceContextsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimVirtualDeviceContextsPartialUpdateParamsWithContext creates a new DcimVirtualDeviceContextsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimVirtualDeviceContextsPartialUpdateParamsWithContext(ctx context.Context) *DcimVirtualDeviceContextsPartialUpdateParams {
	return &DcimVirtualDeviceContextsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimVirtualDeviceContextsPartialUpdateParamsWithHTTPClient creates a new DcimVirtualDeviceContextsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualDeviceContextsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsPartialUpdateParams {
	return &DcimVirtualDeviceContextsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimVirtualDeviceContextsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the dcim virtual device contexts partial update operation.

	Typically these are written to a http.Request.
*/
type DcimVirtualDeviceContextsPartialUpdateParams struct {

	// Data.
	Data *models.WritableVirtualDeviceContext

	/* ID.

	   A unique integer value identifying this virtual device context.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual device contexts partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WithDefaults() *DcimVirtualDeviceContextsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual device contexts partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualDeviceContextsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimVirtualDeviceContextsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WithContext(ctx context.Context) *DcimVirtualDeviceContextsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimVirtualDeviceContextsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WithData(data *models.WritableVirtualDeviceContext) *DcimVirtualDeviceContextsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) SetData(data *models.WritableVirtualDeviceContext) {
	o.Data = data
}

// WithID adds the id to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WithID(id int64) *DcimVirtualDeviceContextsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual device contexts partial update params
func (o *DcimVirtualDeviceContextsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualDeviceContextsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimLocationsBulkDeleteReader is a Reader for the DcimLocationsBulkDelete structure.
type DcimLocationsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimLocationsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimLocationsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsBulkDeleteNoContent creates a DcimLocationsBulkDeleteNoContent with default headers values
func NewDcimLocationsBulkDeleteNoContent() *DcimLocationsBulkDeleteNoContent {
	return &DcimLocationsBulkDeleteNoContent{}
}

/*
DcimLocationsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimLocationsBulkDeleteNoContent dcim locations bulk delete no content
*/
type DcimLocationsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim locations bulk delete no content response has a 2xx status code
func (o *DcimLocationsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations bulk delete no content response has a 3xx status code
func (o *DcimLocationsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations bulk delete no content response has a 4xx status code
func (o *DcimLocationsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations bulk delete no content response has a 5xx status code
func (o *DcimLocationsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations bulk delete no content response a status code equal to that given
func (o *DcimLocationsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim locations bulk delete no content response
func (o *DcimLocationsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimLocationsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/][%d] dcimLocationsBulkDeleteNoContent ", 204)
}

func (o *DcimLocationsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/locations/][%d] dcimLocationsBulkDeleteNoContent ", 204)
}

func (o *DcimLocationsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimLocationsBulkDeleteDefault creates a DcimLocationsBulkDeleteDefault with default headers values
func NewDcimLocationsBulkDeleteDefault(code int) *DcimLocationsBulkDeleteDefault {
	return &DcimLocationsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimLocationsBulkDeleteDefault dcim locations bulk delete default
*/
type DcimLocationsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim locations bulk delete default response has a 2xx status code
func (o *DcimLocationsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations bulk delete default response has a 3xx status code
func (o *DcimLocationsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations bulk delete default response has a 4xx status code
func (o *DcimLocationsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations bulk delete default response has a 5xx status code
func (o *DcimLocationsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations bulk delete default response a status code equal to that given
func (o *DcimLocationsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim locations bulk delete default response
func (o *DcimLocationsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimLocationsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/][%d] dcim_locations_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/locations/][%d] dcim_locations_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

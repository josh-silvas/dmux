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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// TenancyContactsCreateReader is a Reader for the TenancyContactsCreate structure.
type TenancyContactsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyContactsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactsCreateCreated creates a TenancyContactsCreateCreated with default headers values
func NewTenancyContactsCreateCreated() *TenancyContactsCreateCreated {
	return &TenancyContactsCreateCreated{}
}

/*
TenancyContactsCreateCreated describes a response with status code 201, with default header values.

TenancyContactsCreateCreated tenancy contacts create created
*/
type TenancyContactsCreateCreated struct {
	Payload *models.Contact
}

// IsSuccess returns true when this tenancy contacts create created response has a 2xx status code
func (o *TenancyContactsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts create created response has a 3xx status code
func (o *TenancyContactsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts create created response has a 4xx status code
func (o *TenancyContactsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts create created response has a 5xx status code
func (o *TenancyContactsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts create created response a status code equal to that given
func (o *TenancyContactsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the tenancy contacts create created response
func (o *TenancyContactsCreateCreated) Code() int {
	return 201
}

func (o *TenancyContactsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/contacts/][%d] tenancyContactsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyContactsCreateCreated) String() string {
	return fmt.Sprintf("[POST /tenancy/contacts/][%d] tenancyContactsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyContactsCreateCreated) GetPayload() *models.Contact {
	return o.Payload
}

func (o *TenancyContactsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Contact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactsCreateDefault creates a TenancyContactsCreateDefault with default headers values
func NewTenancyContactsCreateDefault(code int) *TenancyContactsCreateDefault {
	return &TenancyContactsCreateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactsCreateDefault describes a response with status code -1, with default header values.

TenancyContactsCreateDefault tenancy contacts create default
*/
type TenancyContactsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this tenancy contacts create default response has a 2xx status code
func (o *TenancyContactsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contacts create default response has a 3xx status code
func (o *TenancyContactsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contacts create default response has a 4xx status code
func (o *TenancyContactsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contacts create default response has a 5xx status code
func (o *TenancyContactsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contacts create default response a status code equal to that given
func (o *TenancyContactsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the tenancy contacts create default response
func (o *TenancyContactsCreateDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /tenancy/contacts/][%d] tenancy_contacts_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsCreateDefault) String() string {
	return fmt.Sprintf("[POST /tenancy/contacts/][%d] tenancy_contacts_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// IpamRirsPartialUpdateReader is a Reader for the IpamRirsPartialUpdate structure.
type IpamRirsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRirsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamRirsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsPartialUpdateOK creates a IpamRirsPartialUpdateOK with default headers values
func NewIpamRirsPartialUpdateOK() *IpamRirsPartialUpdateOK {
	return &IpamRirsPartialUpdateOK{}
}

/*
IpamRirsPartialUpdateOK describes a response with status code 200, with default header values.

IpamRirsPartialUpdateOK ipam rirs partial update o k
*/
type IpamRirsPartialUpdateOK struct {
	Payload *models.RIR
}

// IsSuccess returns true when this ipam rirs partial update o k response has a 2xx status code
func (o *IpamRirsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs partial update o k response has a 3xx status code
func (o *IpamRirsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs partial update o k response has a 4xx status code
func (o *IpamRirsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs partial update o k response has a 5xx status code
func (o *IpamRirsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs partial update o k response a status code equal to that given
func (o *IpamRirsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ipam rirs partial update o k response
func (o *IpamRirsPartialUpdateOK) Code() int {
	return 200
}

func (o *IpamRirsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/rirs/{id}/][%d] ipamRirsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRirsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/rirs/{id}/][%d] ipamRirsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRirsPartialUpdateOK) GetPayload() *models.RIR {
	return o.Payload
}

func (o *IpamRirsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RIR)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRirsPartialUpdateDefault creates a IpamRirsPartialUpdateDefault with default headers values
func NewIpamRirsPartialUpdateDefault(code int) *IpamRirsPartialUpdateDefault {
	return &IpamRirsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRirsPartialUpdateDefault describes a response with status code -1, with default header values.

IpamRirsPartialUpdateDefault ipam rirs partial update default
*/
type IpamRirsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this ipam rirs partial update default response has a 2xx status code
func (o *IpamRirsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs partial update default response has a 3xx status code
func (o *IpamRirsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs partial update default response has a 4xx status code
func (o *IpamRirsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs partial update default response has a 5xx status code
func (o *IpamRirsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs partial update default response a status code equal to that given
func (o *IpamRirsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the ipam rirs partial update default response
func (o *IpamRirsPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *IpamRirsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/rirs/{id}/][%d] ipam_rirs_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/rirs/{id}/][%d] ipam_rirs_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
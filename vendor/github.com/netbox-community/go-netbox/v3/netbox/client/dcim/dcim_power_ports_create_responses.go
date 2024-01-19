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

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// DcimPowerPortsCreateReader is a Reader for the DcimPowerPortsCreate structure.
type DcimPowerPortsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPortsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerPortsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortsCreateCreated creates a DcimPowerPortsCreateCreated with default headers values
func NewDcimPowerPortsCreateCreated() *DcimPowerPortsCreateCreated {
	return &DcimPowerPortsCreateCreated{}
}

/*
DcimPowerPortsCreateCreated describes a response with status code 201, with default header values.

DcimPowerPortsCreateCreated dcim power ports create created
*/
type DcimPowerPortsCreateCreated struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports create created response has a 2xx status code
func (o *DcimPowerPortsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports create created response has a 3xx status code
func (o *DcimPowerPortsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports create created response has a 4xx status code
func (o *DcimPowerPortsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports create created response has a 5xx status code
func (o *DcimPowerPortsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports create created response a status code equal to that given
func (o *DcimPowerPortsCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the dcim power ports create created response
func (o *DcimPowerPortsCreateCreated) Code() int {
	return 201
}

func (o *DcimPowerPortsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcimPowerPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPortsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcimPowerPortsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPortsCreateCreated) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsCreateDefault creates a DcimPowerPortsCreateDefault with default headers values
func NewDcimPowerPortsCreateDefault(code int) *DcimPowerPortsCreateDefault {
	return &DcimPowerPortsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortsCreateDefault describes a response with status code -1, with default header values.

DcimPowerPortsCreateDefault dcim power ports create default
*/
type DcimPowerPortsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power ports create default response has a 2xx status code
func (o *DcimPowerPortsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power ports create default response has a 3xx status code
func (o *DcimPowerPortsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power ports create default response has a 4xx status code
func (o *DcimPowerPortsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power ports create default response has a 5xx status code
func (o *DcimPowerPortsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power ports create default response a status code equal to that given
func (o *DcimPowerPortsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power ports create default response
func (o *DcimPowerPortsCreateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerPortsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcim_power-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/power-ports/][%d] dcim_power-ports_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

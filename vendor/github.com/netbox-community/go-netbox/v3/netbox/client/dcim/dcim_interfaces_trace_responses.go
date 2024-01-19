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

// DcimInterfacesTraceReader is a Reader for the DcimInterfacesTrace structure.
type DcimInterfacesTraceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesTraceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesTraceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInterfacesTraceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesTraceOK creates a DcimInterfacesTraceOK with default headers values
func NewDcimInterfacesTraceOK() *DcimInterfacesTraceOK {
	return &DcimInterfacesTraceOK{}
}

/*
DcimInterfacesTraceOK describes a response with status code 200, with default header values.

DcimInterfacesTraceOK dcim interfaces trace o k
*/
type DcimInterfacesTraceOK struct {
	Payload *models.Interface
}

// IsSuccess returns true when this dcim interfaces trace o k response has a 2xx status code
func (o *DcimInterfacesTraceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces trace o k response has a 3xx status code
func (o *DcimInterfacesTraceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces trace o k response has a 4xx status code
func (o *DcimInterfacesTraceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces trace o k response has a 5xx status code
func (o *DcimInterfacesTraceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces trace o k response a status code equal to that given
func (o *DcimInterfacesTraceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim interfaces trace o k response
func (o *DcimInterfacesTraceOK) Code() int {
	return 200
}

func (o *DcimInterfacesTraceOK) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/trace/][%d] dcimInterfacesTraceOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesTraceOK) String() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/trace/][%d] dcimInterfacesTraceOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesTraceOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesTraceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesTraceDefault creates a DcimInterfacesTraceDefault with default headers values
func NewDcimInterfacesTraceDefault(code int) *DcimInterfacesTraceDefault {
	return &DcimInterfacesTraceDefault{
		_statusCode: code,
	}
}

/*
DcimInterfacesTraceDefault describes a response with status code -1, with default header values.

DcimInterfacesTraceDefault dcim interfaces trace default
*/
type DcimInterfacesTraceDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim interfaces trace default response has a 2xx status code
func (o *DcimInterfacesTraceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interfaces trace default response has a 3xx status code
func (o *DcimInterfacesTraceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interfaces trace default response has a 4xx status code
func (o *DcimInterfacesTraceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interfaces trace default response has a 5xx status code
func (o *DcimInterfacesTraceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interfaces trace default response a status code equal to that given
func (o *DcimInterfacesTraceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim interfaces trace default response
func (o *DcimInterfacesTraceDefault) Code() int {
	return o._statusCode
}

func (o *DcimInterfacesTraceDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/trace/][%d] dcim_interfaces_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesTraceDefault) String() string {
	return fmt.Sprintf("[GET /dcim/interfaces/{id}/trace/][%d] dcim_interfaces_trace default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesTraceDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesTraceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// DcimConsoleServerPortsUpdateReader is a Reader for the DcimConsoleServerPortsUpdate structure.
type DcimConsoleServerPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsUpdateOK creates a DcimConsoleServerPortsUpdateOK with default headers values
func NewDcimConsoleServerPortsUpdateOK() *DcimConsoleServerPortsUpdateOK {
	return &DcimConsoleServerPortsUpdateOK{}
}

/*
DcimConsoleServerPortsUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortsUpdateOK dcim console server ports update o k
*/
type DcimConsoleServerPortsUpdateOK struct {
	Payload *models.ConsoleServerPort
}

// IsSuccess returns true when this dcim console server ports update o k response has a 2xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports update o k response has a 3xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports update o k response has a 4xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports update o k response has a 5xx status code
func (o *DcimConsoleServerPortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports update o k response a status code equal to that given
func (o *DcimConsoleServerPortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim console server ports update o k response
func (o *DcimConsoleServerPortsUpdateOK) Code() int {
	return 200
}

func (o *DcimConsoleServerPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcimConsoleServerPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortsUpdateOK) GetPayload() *models.ConsoleServerPort {
	return o.Payload
}

func (o *DcimConsoleServerPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsUpdateDefault creates a DcimConsoleServerPortsUpdateDefault with default headers values
func NewDcimConsoleServerPortsUpdateDefault(code int) *DcimConsoleServerPortsUpdateDefault {
	return &DcimConsoleServerPortsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortsUpdateDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortsUpdateDefault dcim console server ports update default
*/
type DcimConsoleServerPortsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim console server ports update default response has a 2xx status code
func (o *DcimConsoleServerPortsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server ports update default response has a 3xx status code
func (o *DcimConsoleServerPortsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server ports update default response has a 4xx status code
func (o *DcimConsoleServerPortsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server ports update default response has a 5xx status code
func (o *DcimConsoleServerPortsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server ports update default response a status code equal to that given
func (o *DcimConsoleServerPortsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim console server ports update default response
func (o *DcimConsoleServerPortsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimConsoleServerPortsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-ports/{id}/][%d] dcim_console-server-ports_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
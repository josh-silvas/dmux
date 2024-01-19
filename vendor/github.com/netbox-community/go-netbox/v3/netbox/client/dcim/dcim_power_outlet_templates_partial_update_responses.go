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

// DcimPowerOutletTemplatesPartialUpdateReader is a Reader for the DcimPowerOutletTemplatesPartialUpdate structure.
type DcimPowerOutletTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimPowerOutletTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletTemplatesPartialUpdateOK creates a DcimPowerOutletTemplatesPartialUpdateOK with default headers values
func NewDcimPowerOutletTemplatesPartialUpdateOK() *DcimPowerOutletTemplatesPartialUpdateOK {
	return &DcimPowerOutletTemplatesPartialUpdateOK{}
}

/*
DcimPowerOutletTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletTemplatesPartialUpdateOK dcim power outlet templates partial update o k
*/
type DcimPowerOutletTemplatesPartialUpdateOK struct {
	Payload *models.PowerOutletTemplate
}

// IsSuccess returns true when this dcim power outlet templates partial update o k response has a 2xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlet templates partial update o k response has a 3xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlet templates partial update o k response has a 4xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlet templates partial update o k response has a 5xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlet templates partial update o k response a status code equal to that given
func (o *DcimPowerOutletTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim power outlet templates partial update o k response
func (o *DcimPowerOutletTemplatesPartialUpdateOK) Code() int {
	return 200
}

func (o *DcimPowerOutletTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/{id}/][%d] dcimPowerOutletTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletTemplatesPartialUpdateOK) GetPayload() *models.PowerOutletTemplate {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutletTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletTemplatesPartialUpdateDefault creates a DcimPowerOutletTemplatesPartialUpdateDefault with default headers values
func NewDcimPowerOutletTemplatesPartialUpdateDefault(code int) *DcimPowerOutletTemplatesPartialUpdateDefault {
	return &DcimPowerOutletTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

DcimPowerOutletTemplatesPartialUpdateDefault dcim power outlet templates partial update default
*/
type DcimPowerOutletTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim power outlet templates partial update default response has a 2xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlet templates partial update default response has a 3xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlet templates partial update default response has a 4xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlet templates partial update default response has a 5xx status code
func (o *DcimPowerOutletTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlet templates partial update default response a status code equal to that given
func (o *DcimPowerOutletTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim power outlet templates partial update default response
func (o *DcimPowerOutletTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimPowerOutletTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-outlet-templates/{id}/][%d] dcim_power-outlet-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

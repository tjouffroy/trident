// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnmpUsersGetReader is a Reader for the SnmpUsersGet structure.
type SnmpUsersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpUsersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpUsersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpUsersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpUsersGetOK creates a SnmpUsersGetOK with default headers values
func NewSnmpUsersGetOK() *SnmpUsersGetOK {
	return &SnmpUsersGetOK{}
}

/*
SnmpUsersGetOK describes a response with status code 200, with default header values.

OK
*/
type SnmpUsersGetOK struct {
	Payload *models.SnmpUser
}

// IsSuccess returns true when this snmp users get o k response has a 2xx status code
func (o *SnmpUsersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snmp users get o k response has a 3xx status code
func (o *SnmpUsersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snmp users get o k response has a 4xx status code
func (o *SnmpUsersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snmp users get o k response has a 5xx status code
func (o *SnmpUsersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snmp users get o k response a status code equal to that given
func (o *SnmpUsersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the snmp users get o k response
func (o *SnmpUsersGetOK) Code() int {
	return 200
}

func (o *SnmpUsersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users/{engine_id}/{name}][%d] snmpUsersGetOK %s", 200, payload)
}

func (o *SnmpUsersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users/{engine_id}/{name}][%d] snmpUsersGetOK %s", 200, payload)
}

func (o *SnmpUsersGetOK) GetPayload() *models.SnmpUser {
	return o.Payload
}

func (o *SnmpUsersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnmpUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnmpUsersGetDefault creates a SnmpUsersGetDefault with default headers values
func NewSnmpUsersGetDefault(code int) *SnmpUsersGetDefault {
	return &SnmpUsersGetDefault{
		_statusCode: code,
	}
}

/*
SnmpUsersGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnmpUsersGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snmp users get default response has a 2xx status code
func (o *SnmpUsersGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snmp users get default response has a 3xx status code
func (o *SnmpUsersGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snmp users get default response has a 4xx status code
func (o *SnmpUsersGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snmp users get default response has a 5xx status code
func (o *SnmpUsersGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snmp users get default response a status code equal to that given
func (o *SnmpUsersGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snmp users get default response
func (o *SnmpUsersGetDefault) Code() int {
	return o._statusCode
}

func (o *SnmpUsersGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users/{engine_id}/{name}][%d] snmp_users_get default %s", o._statusCode, payload)
}

func (o *SnmpUsersGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /support/snmp/users/{engine_id}/{name}][%d] snmp_users_get default %s", o._statusCode, payload)
}

func (o *SnmpUsersGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpUsersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

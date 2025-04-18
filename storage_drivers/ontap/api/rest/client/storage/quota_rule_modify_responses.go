// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// QuotaRuleModifyReader is a Reader for the QuotaRuleModify structure.
type QuotaRuleModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QuotaRuleModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQuotaRuleModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewQuotaRuleModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQuotaRuleModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQuotaRuleModifyOK creates a QuotaRuleModifyOK with default headers values
func NewQuotaRuleModifyOK() *QuotaRuleModifyOK {
	return &QuotaRuleModifyOK{}
}

/*
QuotaRuleModifyOK describes a response with status code 200, with default header values.

OK
*/
type QuotaRuleModifyOK struct {
	Payload *models.QuotaRuleJobLinkResponse
}

// IsSuccess returns true when this quota rule modify o k response has a 2xx status code
func (o *QuotaRuleModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota rule modify o k response has a 3xx status code
func (o *QuotaRuleModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota rule modify o k response has a 4xx status code
func (o *QuotaRuleModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota rule modify o k response has a 5xx status code
func (o *QuotaRuleModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this quota rule modify o k response a status code equal to that given
func (o *QuotaRuleModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the quota rule modify o k response
func (o *QuotaRuleModifyOK) Code() int {
	return 200
}

func (o *QuotaRuleModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/quota/rules/{uuid}][%d] quotaRuleModifyOK %s", 200, payload)
}

func (o *QuotaRuleModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/quota/rules/{uuid}][%d] quotaRuleModifyOK %s", 200, payload)
}

func (o *QuotaRuleModifyOK) GetPayload() *models.QuotaRuleJobLinkResponse {
	return o.Payload
}

func (o *QuotaRuleModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaRuleJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaRuleModifyAccepted creates a QuotaRuleModifyAccepted with default headers values
func NewQuotaRuleModifyAccepted() *QuotaRuleModifyAccepted {
	return &QuotaRuleModifyAccepted{}
}

/*
QuotaRuleModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QuotaRuleModifyAccepted struct {
	Payload *models.QuotaRuleJobLinkResponse
}

// IsSuccess returns true when this quota rule modify accepted response has a 2xx status code
func (o *QuotaRuleModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this quota rule modify accepted response has a 3xx status code
func (o *QuotaRuleModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this quota rule modify accepted response has a 4xx status code
func (o *QuotaRuleModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this quota rule modify accepted response has a 5xx status code
func (o *QuotaRuleModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this quota rule modify accepted response a status code equal to that given
func (o *QuotaRuleModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the quota rule modify accepted response
func (o *QuotaRuleModifyAccepted) Code() int {
	return 202
}

func (o *QuotaRuleModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/quota/rules/{uuid}][%d] quotaRuleModifyAccepted %s", 202, payload)
}

func (o *QuotaRuleModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/quota/rules/{uuid}][%d] quotaRuleModifyAccepted %s", 202, payload)
}

func (o *QuotaRuleModifyAccepted) GetPayload() *models.QuotaRuleJobLinkResponse {
	return o.Payload
}

func (o *QuotaRuleModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QuotaRuleJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQuotaRuleModifyDefault creates a QuotaRuleModifyDefault with default headers values
func NewQuotaRuleModifyDefault(code int) *QuotaRuleModifyDefault {
	return &QuotaRuleModifyDefault{
		_statusCode: code,
	}
}

/*
	QuotaRuleModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5308478 | User mapping can only be turned on for UNIX usernames or Windows account names. |
| 5308501 | Mapping from Windows user to UNIX user for user rule was unsuccessful. |
| 5308502 | Mapping from UNIX user to Windows user for user rule was unsuccessful. |
| 5308545 | The specified quota rule UUID is invalid. |
| 5308561 | Failed to obtain volume quota state or invalid quota state obtained for volume. |
| 5308567 | Quota policy rule modify operation succeeded, but quota resize failed due to internal error. |
| 5308573 | Input value is greater than limit for field. |
| 5308574 | Input value is out of range for field. |
| 5308575 | Input value is incorrectly larger than listed field. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type QuotaRuleModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this quota rule modify default response has a 2xx status code
func (o *QuotaRuleModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this quota rule modify default response has a 3xx status code
func (o *QuotaRuleModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this quota rule modify default response has a 4xx status code
func (o *QuotaRuleModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this quota rule modify default response has a 5xx status code
func (o *QuotaRuleModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this quota rule modify default response a status code equal to that given
func (o *QuotaRuleModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the quota rule modify default response
func (o *QuotaRuleModifyDefault) Code() int {
	return o._statusCode
}

func (o *QuotaRuleModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/quota/rules/{uuid}][%d] quota_rule_modify default %s", o._statusCode, payload)
}

func (o *QuotaRuleModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/quota/rules/{uuid}][%d] quota_rule_modify default %s", o._statusCode, payload)
}

func (o *QuotaRuleModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QuotaRuleModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

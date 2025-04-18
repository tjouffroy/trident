// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// SnaplockLogCreateReader is a Reader for the SnaplockLogCreate structure.
type SnaplockLogCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLogCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockLogCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSnaplockLogCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLogCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLogCreateCreated creates a SnaplockLogCreateCreated with default headers values
func NewSnaplockLogCreateCreated() *SnaplockLogCreateCreated {
	return &SnaplockLogCreateCreated{}
}

/*
SnaplockLogCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockLogCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SnaplockLogJobLinkResponse
}

// IsSuccess returns true when this snaplock log create created response has a 2xx status code
func (o *SnaplockLogCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log create created response has a 3xx status code
func (o *SnaplockLogCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log create created response has a 4xx status code
func (o *SnaplockLogCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log create created response has a 5xx status code
func (o *SnaplockLogCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log create created response a status code equal to that given
func (o *SnaplockLogCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the snaplock log create created response
func (o *SnaplockLogCreateCreated) Code() int {
	return 201
}

func (o *SnaplockLogCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplockLogCreateCreated %s", 201, payload)
}

func (o *SnaplockLogCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplockLogCreateCreated %s", 201, payload)
}

func (o *SnaplockLogCreateCreated) GetPayload() *models.SnaplockLogJobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SnaplockLogJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogCreateAccepted creates a SnaplockLogCreateAccepted with default headers values
func NewSnaplockLogCreateAccepted() *SnaplockLogCreateAccepted {
	return &SnaplockLogCreateAccepted{}
}

/*
SnaplockLogCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnaplockLogCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.SnaplockLogJobLinkResponse
}

// IsSuccess returns true when this snaplock log create accepted response has a 2xx status code
func (o *SnaplockLogCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock log create accepted response has a 3xx status code
func (o *SnaplockLogCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock log create accepted response has a 4xx status code
func (o *SnaplockLogCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock log create accepted response has a 5xx status code
func (o *SnaplockLogCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock log create accepted response a status code equal to that given
func (o *SnaplockLogCreateAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the snaplock log create accepted response
func (o *SnaplockLogCreateAccepted) Code() int {
	return 202
}

func (o *SnaplockLogCreateAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplockLogCreateAccepted %s", 202, payload)
}

func (o *SnaplockLogCreateAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplockLogCreateAccepted %s", 202, payload)
}

func (o *SnaplockLogCreateAccepted) GetPayload() *models.SnaplockLogJobLinkResponse {
	return o.Payload
}

func (o *SnaplockLogCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.SnaplockLogJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLogCreateDefault creates a SnaplockLogCreateDefault with default headers values
func NewSnaplockLogCreateDefault(code int) *SnaplockLogCreateDefault {
	return &SnaplockLogCreateDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLogCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
|   918236    | The specified volume name and UUID refer to different volumes  |
|   918253    | Incorrect format for the retention period, duration must be in the ISO-8601 format |
| 13763161    | Audit logging is already configured for the SVM  |
| 14090340    | {field} is a required field  |
| 14090343    | Invalid Field  |
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
*/
type SnaplockLogCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock log create default response has a 2xx status code
func (o *SnaplockLogCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock log create default response has a 3xx status code
func (o *SnaplockLogCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock log create default response has a 4xx status code
func (o *SnaplockLogCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock log create default response has a 5xx status code
func (o *SnaplockLogCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock log create default response a status code equal to that given
func (o *SnaplockLogCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock log create default response
func (o *SnaplockLogCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockLogCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplock_log_create default %s", o._statusCode, payload)
}

func (o *SnaplockLogCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/audit-logs][%d] snaplock_log_create default %s", o._statusCode, payload)
}

func (o *SnaplockLogCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLogCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

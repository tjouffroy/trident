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

// StorageAvailabilityZoneModifyReader is a Reader for the StorageAvailabilityZoneModify structure.
type StorageAvailabilityZoneModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StorageAvailabilityZoneModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStorageAvailabilityZoneModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewStorageAvailabilityZoneModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStorageAvailabilityZoneModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStorageAvailabilityZoneModifyOK creates a StorageAvailabilityZoneModifyOK with default headers values
func NewStorageAvailabilityZoneModifyOK() *StorageAvailabilityZoneModifyOK {
	return &StorageAvailabilityZoneModifyOK{}
}

/*
StorageAvailabilityZoneModifyOK describes a response with status code 200, with default header values.

OK
*/
type StorageAvailabilityZoneModifyOK struct {
	Payload *models.StorageAvailabilityZoneJobLinkResponse
}

// IsSuccess returns true when this storage availability zone modify o k response has a 2xx status code
func (o *StorageAvailabilityZoneModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage availability zone modify o k response has a 3xx status code
func (o *StorageAvailabilityZoneModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage availability zone modify o k response has a 4xx status code
func (o *StorageAvailabilityZoneModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage availability zone modify o k response has a 5xx status code
func (o *StorageAvailabilityZoneModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this storage availability zone modify o k response a status code equal to that given
func (o *StorageAvailabilityZoneModifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the storage availability zone modify o k response
func (o *StorageAvailabilityZoneModifyOK) Code() int {
	return 200
}

func (o *StorageAvailabilityZoneModifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/availability-zones/{uuid}][%d] storageAvailabilityZoneModifyOK %s", 200, payload)
}

func (o *StorageAvailabilityZoneModifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/availability-zones/{uuid}][%d] storageAvailabilityZoneModifyOK %s", 200, payload)
}

func (o *StorageAvailabilityZoneModifyOK) GetPayload() *models.StorageAvailabilityZoneJobLinkResponse {
	return o.Payload
}

func (o *StorageAvailabilityZoneModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageAvailabilityZoneJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageAvailabilityZoneModifyAccepted creates a StorageAvailabilityZoneModifyAccepted with default headers values
func NewStorageAvailabilityZoneModifyAccepted() *StorageAvailabilityZoneModifyAccepted {
	return &StorageAvailabilityZoneModifyAccepted{}
}

/*
StorageAvailabilityZoneModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type StorageAvailabilityZoneModifyAccepted struct {
	Payload *models.StorageAvailabilityZoneJobLinkResponse
}

// IsSuccess returns true when this storage availability zone modify accepted response has a 2xx status code
func (o *StorageAvailabilityZoneModifyAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this storage availability zone modify accepted response has a 3xx status code
func (o *StorageAvailabilityZoneModifyAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this storage availability zone modify accepted response has a 4xx status code
func (o *StorageAvailabilityZoneModifyAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this storage availability zone modify accepted response has a 5xx status code
func (o *StorageAvailabilityZoneModifyAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this storage availability zone modify accepted response a status code equal to that given
func (o *StorageAvailabilityZoneModifyAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the storage availability zone modify accepted response
func (o *StorageAvailabilityZoneModifyAccepted) Code() int {
	return 202
}

func (o *StorageAvailabilityZoneModifyAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/availability-zones/{uuid}][%d] storageAvailabilityZoneModifyAccepted %s", 202, payload)
}

func (o *StorageAvailabilityZoneModifyAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/availability-zones/{uuid}][%d] storageAvailabilityZoneModifyAccepted %s", 202, payload)
}

func (o *StorageAvailabilityZoneModifyAccepted) GetPayload() *models.StorageAvailabilityZoneJobLinkResponse {
	return o.Payload
}

func (o *StorageAvailabilityZoneModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageAvailabilityZoneJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStorageAvailabilityZoneModifyDefault creates a StorageAvailabilityZoneModifyDefault with default headers values
func NewStorageAvailabilityZoneModifyDefault(code int) *StorageAvailabilityZoneModifyDefault {
	return &StorageAvailabilityZoneModifyDefault{
		_statusCode: code,
	}
}

/*
	StorageAvailabilityZoneModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9240632 | The provided nearly full threshold percent cannot be greater than full threshold percent. |
| 9240633 | The provided full threshold percent cannot be smaller than nearly full threshold percent. |
| 9240634 | The provided storage availability zone UUID is not valid. |
| 262144015 | Failed to retrieve storage availability zone information for provided storage availability zone UUID. |
| 262144016 | Failed to retrieve the node information for provided storage availability zone UUID. |
| 262144017 | This operation is not supported on this platform or has been disabled. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type StorageAvailabilityZoneModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this storage availability zone modify default response has a 2xx status code
func (o *StorageAvailabilityZoneModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this storage availability zone modify default response has a 3xx status code
func (o *StorageAvailabilityZoneModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this storage availability zone modify default response has a 4xx status code
func (o *StorageAvailabilityZoneModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this storage availability zone modify default response has a 5xx status code
func (o *StorageAvailabilityZoneModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this storage availability zone modify default response a status code equal to that given
func (o *StorageAvailabilityZoneModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the storage availability zone modify default response
func (o *StorageAvailabilityZoneModifyDefault) Code() int {
	return o._statusCode
}

func (o *StorageAvailabilityZoneModifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/availability-zones/{uuid}][%d] storage_availability_zone_modify default %s", o._statusCode, payload)
}

func (o *StorageAvailabilityZoneModifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PATCH /storage/availability-zones/{uuid}][%d] storage_availability_zone_modify default %s", o._statusCode, payload)
}

func (o *StorageAvailabilityZoneModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *StorageAvailabilityZoneModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// LocalCifsGroupCollectionGetReader is a Reader for the LocalCifsGroupCollectionGet structure.
type LocalCifsGroupCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LocalCifsGroupCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLocalCifsGroupCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLocalCifsGroupCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLocalCifsGroupCollectionGetOK creates a LocalCifsGroupCollectionGetOK with default headers values
func NewLocalCifsGroupCollectionGetOK() *LocalCifsGroupCollectionGetOK {
	return &LocalCifsGroupCollectionGetOK{}
}

/*
LocalCifsGroupCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type LocalCifsGroupCollectionGetOK struct {
	Payload *models.LocalCifsGroupResponse
}

// IsSuccess returns true when this local cifs group collection get o k response has a 2xx status code
func (o *LocalCifsGroupCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this local cifs group collection get o k response has a 3xx status code
func (o *LocalCifsGroupCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this local cifs group collection get o k response has a 4xx status code
func (o *LocalCifsGroupCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this local cifs group collection get o k response has a 5xx status code
func (o *LocalCifsGroupCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this local cifs group collection get o k response a status code equal to that given
func (o *LocalCifsGroupCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the local cifs group collection get o k response
func (o *LocalCifsGroupCollectionGetOK) Code() int {
	return 200
}

func (o *LocalCifsGroupCollectionGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups][%d] localCifsGroupCollectionGetOK %s", 200, payload)
}

func (o *LocalCifsGroupCollectionGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups][%d] localCifsGroupCollectionGetOK %s", 200, payload)
}

func (o *LocalCifsGroupCollectionGetOK) GetPayload() *models.LocalCifsGroupResponse {
	return o.Payload
}

func (o *LocalCifsGroupCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalCifsGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLocalCifsGroupCollectionGetDefault creates a LocalCifsGroupCollectionGetDefault with default headers values
func NewLocalCifsGroupCollectionGetDefault(code int) *LocalCifsGroupCollectionGetDefault {
	return &LocalCifsGroupCollectionGetDefault{
		_statusCode: code,
	}
}

/*
LocalCifsGroupCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type LocalCifsGroupCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this local cifs group collection get default response has a 2xx status code
func (o *LocalCifsGroupCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this local cifs group collection get default response has a 3xx status code
func (o *LocalCifsGroupCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this local cifs group collection get default response has a 4xx status code
func (o *LocalCifsGroupCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this local cifs group collection get default response has a 5xx status code
func (o *LocalCifsGroupCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this local cifs group collection get default response a status code equal to that given
func (o *LocalCifsGroupCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the local cifs group collection get default response
func (o *LocalCifsGroupCollectionGetDefault) Code() int {
	return o._statusCode
}

func (o *LocalCifsGroupCollectionGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups][%d] local_cifs_group_collection_get default %s", o._statusCode, payload)
}

func (o *LocalCifsGroupCollectionGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /protocols/cifs/local-groups][%d] local_cifs_group_collection_get default %s", o._statusCode, payload)
}

func (o *LocalCifsGroupCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LocalCifsGroupCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

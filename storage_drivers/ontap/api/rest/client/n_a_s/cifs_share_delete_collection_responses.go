// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CifsShareDeleteCollectionReader is a Reader for the CifsShareDeleteCollection structure.
type CifsShareDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsShareDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsShareDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsShareDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsShareDeleteCollectionOK creates a CifsShareDeleteCollectionOK with default headers values
func NewCifsShareDeleteCollectionOK() *CifsShareDeleteCollectionOK {
	return &CifsShareDeleteCollectionOK{}
}

/*
CifsShareDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type CifsShareDeleteCollectionOK struct {
}

// IsSuccess returns true when this cifs share delete collection o k response has a 2xx status code
func (o *CifsShareDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs share delete collection o k response has a 3xx status code
func (o *CifsShareDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs share delete collection o k response has a 4xx status code
func (o *CifsShareDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs share delete collection o k response has a 5xx status code
func (o *CifsShareDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs share delete collection o k response a status code equal to that given
func (o *CifsShareDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cifs share delete collection o k response
func (o *CifsShareDeleteCollectionOK) Code() int {
	return 200
}

func (o *CifsShareDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares][%d] cifsShareDeleteCollectionOK", 200)
}

func (o *CifsShareDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/cifs/shares][%d] cifsShareDeleteCollectionOK", 200)
}

func (o *CifsShareDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCifsShareDeleteCollectionDefault creates a CifsShareDeleteCollectionDefault with default headers values
func NewCifsShareDeleteCollectionDefault(code int) *CifsShareDeleteCollectionDefault {
	return &CifsShareDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	CifsShareDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 655393     | Standard admin shares cannot be removed |
*/
type CifsShareDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this cifs share delete collection default response has a 2xx status code
func (o *CifsShareDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs share delete collection default response has a 3xx status code
func (o *CifsShareDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs share delete collection default response has a 4xx status code
func (o *CifsShareDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs share delete collection default response has a 5xx status code
func (o *CifsShareDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs share delete collection default response a status code equal to that given
func (o *CifsShareDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cifs share delete collection default response
func (o *CifsShareDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *CifsShareDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/shares][%d] cifs_share_delete_collection default %s", o._statusCode, payload)
}

func (o *CifsShareDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/cifs/shares][%d] cifs_share_delete_collection default %s", o._statusCode, payload)
}

func (o *CifsShareDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsShareDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CifsShareDeleteCollectionBody cifs share delete collection body
swagger:model CifsShareDeleteCollectionBody
*/
type CifsShareDeleteCollectionBody struct {

	// cifs share response inline records
	CifsShareResponseInlineRecords []*models.CifsShare `json:"records,omitempty"`
}

// Validate validates this cifs share delete collection body
func (o *CifsShareDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCifsShareResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CifsShareDeleteCollectionBody) validateCifsShareResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.CifsShareResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.CifsShareResponseInlineRecords); i++ {
		if swag.IsZero(o.CifsShareResponseInlineRecords[i]) { // not required
			continue
		}

		if o.CifsShareResponseInlineRecords[i] != nil {
			if err := o.CifsShareResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cifs share delete collection body based on the context it is used
func (o *CifsShareDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCifsShareResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CifsShareDeleteCollectionBody) contextValidateCifsShareResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.CifsShareResponseInlineRecords); i++ {

		if o.CifsShareResponseInlineRecords[i] != nil {
			if err := o.CifsShareResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CifsShareDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CifsShareDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res CifsShareDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

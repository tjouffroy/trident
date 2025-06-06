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

// VscanConfigDeleteCollectionReader is a Reader for the VscanConfigDeleteCollection structure.
type VscanConfigDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VscanConfigDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVscanConfigDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVscanConfigDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVscanConfigDeleteCollectionOK creates a VscanConfigDeleteCollectionOK with default headers values
func NewVscanConfigDeleteCollectionOK() *VscanConfigDeleteCollectionOK {
	return &VscanConfigDeleteCollectionOK{}
}

/*
VscanConfigDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type VscanConfigDeleteCollectionOK struct {
}

// IsSuccess returns true when this vscan config delete collection o k response has a 2xx status code
func (o *VscanConfigDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this vscan config delete collection o k response has a 3xx status code
func (o *VscanConfigDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this vscan config delete collection o k response has a 4xx status code
func (o *VscanConfigDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this vscan config delete collection o k response has a 5xx status code
func (o *VscanConfigDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this vscan config delete collection o k response a status code equal to that given
func (o *VscanConfigDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the vscan config delete collection o k response
func (o *VscanConfigDeleteCollectionOK) Code() int {
	return 200
}

func (o *VscanConfigDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/vscan][%d] vscanConfigDeleteCollectionOK", 200)
}

func (o *VscanConfigDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /protocols/vscan][%d] vscanConfigDeleteCollectionOK", 200)
}

func (o *VscanConfigDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVscanConfigDeleteCollectionDefault creates a VscanConfigDeleteCollectionDefault with default headers values
func NewVscanConfigDeleteCollectionDefault(code int) *VscanConfigDeleteCollectionDefault {
	return &VscanConfigDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
	VscanConfigDeleteCollectionDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 10027259   | A scanner-pool, an On-Access policy, or an On-Demand policy might fail to get deleted due to either a systematic error or some hardware failure. The error code returned details the failure along with the reason for the failure. For example, \"Failed to delete On-Access policy \"sp1\". Reason: \"Failed to delete policy. Reason: policy must be disabled before being deleted.\". Retry the operation.\"
*/
type VscanConfigDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this vscan config delete collection default response has a 2xx status code
func (o *VscanConfigDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this vscan config delete collection default response has a 3xx status code
func (o *VscanConfigDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this vscan config delete collection default response has a 4xx status code
func (o *VscanConfigDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this vscan config delete collection default response has a 5xx status code
func (o *VscanConfigDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this vscan config delete collection default response a status code equal to that given
func (o *VscanConfigDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the vscan config delete collection default response
func (o *VscanConfigDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *VscanConfigDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan][%d] vscan_config_delete_collection default %s", o._statusCode, payload)
}

func (o *VscanConfigDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /protocols/vscan][%d] vscan_config_delete_collection default %s", o._statusCode, payload)
}

func (o *VscanConfigDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VscanConfigDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
VscanConfigDeleteCollectionBody vscan config delete collection body
swagger:model VscanConfigDeleteCollectionBody
*/
type VscanConfigDeleteCollectionBody struct {

	// vscan response inline records
	VscanResponseInlineRecords []*models.Vscan `json:"records,omitempty"`
}

// Validate validates this vscan config delete collection body
func (o *VscanConfigDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateVscanResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanConfigDeleteCollectionBody) validateVscanResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.VscanResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.VscanResponseInlineRecords); i++ {
		if swag.IsZero(o.VscanResponseInlineRecords[i]) { // not required
			continue
		}

		if o.VscanResponseInlineRecords[i] != nil {
			if err := o.VscanResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this vscan config delete collection body based on the context it is used
func (o *VscanConfigDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateVscanResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *VscanConfigDeleteCollectionBody) contextValidateVscanResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.VscanResponseInlineRecords); i++ {

		if o.VscanResponseInlineRecords[i] != nil {
			if err := o.VscanResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *VscanConfigDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *VscanConfigDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res VscanConfigDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

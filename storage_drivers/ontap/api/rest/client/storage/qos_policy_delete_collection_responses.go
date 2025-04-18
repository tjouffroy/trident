// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// QosPolicyDeleteCollectionReader is a Reader for the QosPolicyDeleteCollection structure.
type QosPolicyDeleteCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyDeleteCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQosPolicyDeleteCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyDeleteCollectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyDeleteCollectionOK creates a QosPolicyDeleteCollectionOK with default headers values
func NewQosPolicyDeleteCollectionOK() *QosPolicyDeleteCollectionOK {
	return &QosPolicyDeleteCollectionOK{}
}

/*
QosPolicyDeleteCollectionOK describes a response with status code 200, with default header values.

OK
*/
type QosPolicyDeleteCollectionOK struct {
}

// IsSuccess returns true when this qos policy delete collection o k response has a 2xx status code
func (o *QosPolicyDeleteCollectionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy delete collection o k response has a 3xx status code
func (o *QosPolicyDeleteCollectionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy delete collection o k response has a 4xx status code
func (o *QosPolicyDeleteCollectionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy delete collection o k response has a 5xx status code
func (o *QosPolicyDeleteCollectionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy delete collection o k response a status code equal to that given
func (o *QosPolicyDeleteCollectionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qos policy delete collection o k response
func (o *QosPolicyDeleteCollectionOK) Code() int {
	return 200
}

func (o *QosPolicyDeleteCollectionOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies][%d] qosPolicyDeleteCollectionOK", 200)
}

func (o *QosPolicyDeleteCollectionOK) String() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies][%d] qosPolicyDeleteCollectionOK", 200)
}

func (o *QosPolicyDeleteCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewQosPolicyDeleteCollectionDefault creates a QosPolicyDeleteCollectionDefault with default headers values
func NewQosPolicyDeleteCollectionDefault(code int) *QosPolicyDeleteCollectionDefault {
	return &QosPolicyDeleteCollectionDefault{
		_statusCode: code,
	}
}

/*
QosPolicyDeleteCollectionDefault describes a response with status code -1, with default header values.

Error
*/
type QosPolicyDeleteCollectionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this qos policy delete collection default response has a 2xx status code
func (o *QosPolicyDeleteCollectionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos policy delete collection default response has a 3xx status code
func (o *QosPolicyDeleteCollectionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos policy delete collection default response has a 4xx status code
func (o *QosPolicyDeleteCollectionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos policy delete collection default response has a 5xx status code
func (o *QosPolicyDeleteCollectionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos policy delete collection default response a status code equal to that given
func (o *QosPolicyDeleteCollectionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qos policy delete collection default response
func (o *QosPolicyDeleteCollectionDefault) Code() int {
	return o._statusCode
}

func (o *QosPolicyDeleteCollectionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/qos/policies][%d] qos_policy_delete_collection default %s", o._statusCode, payload)
}

func (o *QosPolicyDeleteCollectionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage/qos/policies][%d] qos_policy_delete_collection default %s", o._statusCode, payload)
}

func (o *QosPolicyDeleteCollectionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyDeleteCollectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
QosPolicyDeleteCollectionBody qos policy delete collection body
swagger:model QosPolicyDeleteCollectionBody
*/
type QosPolicyDeleteCollectionBody struct {

	// error
	Error *models.Error `json:"error,omitempty"`

	// qos policy response inline records
	QosPolicyResponseInlineRecords []*models.QosPolicy `json:"records,omitempty"`
}

// Validate validates this qos policy delete collection body
func (o *QosPolicyDeleteCollectionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQosPolicyResponseInlineRecords(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QosPolicyDeleteCollectionBody) validateError(formats strfmt.Registry) error {
	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if o.Error != nil {
		if err := o.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *QosPolicyDeleteCollectionBody) validateQosPolicyResponseInlineRecords(formats strfmt.Registry) error {
	if swag.IsZero(o.QosPolicyResponseInlineRecords) { // not required
		return nil
	}

	for i := 0; i < len(o.QosPolicyResponseInlineRecords); i++ {
		if swag.IsZero(o.QosPolicyResponseInlineRecords[i]) { // not required
			continue
		}

		if o.QosPolicyResponseInlineRecords[i] != nil {
			if err := o.QosPolicyResponseInlineRecords[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + "records" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this qos policy delete collection body based on the context it is used
func (o *QosPolicyDeleteCollectionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateQosPolicyResponseInlineRecords(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *QosPolicyDeleteCollectionBody) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if o.Error != nil {
		if err := o.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info" + "." + "error")
			}
			return err
		}
	}

	return nil
}

func (o *QosPolicyDeleteCollectionBody) contextValidateQosPolicyResponseInlineRecords(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.QosPolicyResponseInlineRecords); i++ {

		if o.QosPolicyResponseInlineRecords[i] != nil {
			if err := o.QosPolicyResponseInlineRecords[i].ContextValidate(ctx, formats); err != nil {
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
func (o *QosPolicyDeleteCollectionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *QosPolicyDeleteCollectionBody) UnmarshalBinary(b []byte) error {
	var res QosPolicyDeleteCollectionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

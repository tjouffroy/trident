// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// WwpnAliasCreateReader is a Reader for the WwpnAliasCreate structure.
type WwpnAliasCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WwpnAliasCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewWwpnAliasCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWwpnAliasCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWwpnAliasCreateCreated creates a WwpnAliasCreateCreated with default headers values
func NewWwpnAliasCreateCreated() *WwpnAliasCreateCreated {
	return &WwpnAliasCreateCreated{}
}

/*
WwpnAliasCreateCreated describes a response with status code 201, with default header values.

Created
*/
type WwpnAliasCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.WwpnAliasResponse
}

// IsSuccess returns true when this wwpn alias create created response has a 2xx status code
func (o *WwpnAliasCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wwpn alias create created response has a 3xx status code
func (o *WwpnAliasCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wwpn alias create created response has a 4xx status code
func (o *WwpnAliasCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this wwpn alias create created response has a 5xx status code
func (o *WwpnAliasCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this wwpn alias create created response a status code equal to that given
func (o *WwpnAliasCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the wwpn alias create created response
func (o *WwpnAliasCreateCreated) Code() int {
	return 201
}

func (o *WwpnAliasCreateCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/fc/wwpn-aliases][%d] wwpnAliasCreateCreated %s", 201, payload)
}

func (o *WwpnAliasCreateCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/fc/wwpn-aliases][%d] wwpnAliasCreateCreated %s", 201, payload)
}

func (o *WwpnAliasCreateCreated) GetPayload() *models.WwpnAliasResponse {
	return o.Payload
}

func (o *WwpnAliasCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.WwpnAliasResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWwpnAliasCreateDefault creates a WwpnAliasCreateDefault with default headers values
func NewWwpnAliasCreateDefault(code int) *WwpnAliasCreateDefault {
	return &WwpnAliasCreateDefault{
		_statusCode: code,
	}
}

/*
	WwpnAliasCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | The specified SVM does not exist or cannot be accessed by this user. |
| 2621706 | The specified `svm.uuid` and `svm.name` do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either `svm.name` or `svm.uuid` must be supplied. |
| 5373982 | An invalid WWN was specified. The length is incorrect. |
| 5373983 | An invalid WWN was specified. The format is incorrect. |
| 5374869 | The alias already exists. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type WwpnAliasCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this wwpn alias create default response has a 2xx status code
func (o *WwpnAliasCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wwpn alias create default response has a 3xx status code
func (o *WwpnAliasCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wwpn alias create default response has a 4xx status code
func (o *WwpnAliasCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wwpn alias create default response has a 5xx status code
func (o *WwpnAliasCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wwpn alias create default response a status code equal to that given
func (o *WwpnAliasCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the wwpn alias create default response
func (o *WwpnAliasCreateDefault) Code() int {
	return o._statusCode
}

func (o *WwpnAliasCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/fc/wwpn-aliases][%d] wwpn_alias_create default %s", o._statusCode, payload)
}

func (o *WwpnAliasCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/fc/wwpn-aliases][%d] wwpn_alias_create default %s", o._statusCode, payload)
}

func (o *WwpnAliasCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *WwpnAliasCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

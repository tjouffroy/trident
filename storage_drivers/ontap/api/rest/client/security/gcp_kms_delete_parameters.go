// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGcpKmsDeleteParams creates a new GcpKmsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGcpKmsDeleteParams() *GcpKmsDeleteParams {
	return &GcpKmsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGcpKmsDeleteParamsWithTimeout creates a new GcpKmsDeleteParams object
// with the ability to set a timeout on a request.
func NewGcpKmsDeleteParamsWithTimeout(timeout time.Duration) *GcpKmsDeleteParams {
	return &GcpKmsDeleteParams{
		timeout: timeout,
	}
}

// NewGcpKmsDeleteParamsWithContext creates a new GcpKmsDeleteParams object
// with the ability to set a context for a request.
func NewGcpKmsDeleteParamsWithContext(ctx context.Context) *GcpKmsDeleteParams {
	return &GcpKmsDeleteParams{
		Context: ctx,
	}
}

// NewGcpKmsDeleteParamsWithHTTPClient creates a new GcpKmsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGcpKmsDeleteParamsWithHTTPClient(client *http.Client) *GcpKmsDeleteParams {
	return &GcpKmsDeleteParams{
		HTTPClient: client,
	}
}

/*
GcpKmsDeleteParams contains all the parameters to send to the API endpoint

	for the gcp kms delete operation.

	Typically these are written to a http.Request.
*/
type GcpKmsDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* UUID.

	   Google Cloud KMS UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gcp kms delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsDeleteParams) WithDefaults() *GcpKmsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gcp kms delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GcpKmsDeleteParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := GcpKmsDeleteParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the gcp kms delete params
func (o *GcpKmsDeleteParams) WithTimeout(timeout time.Duration) *GcpKmsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gcp kms delete params
func (o *GcpKmsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gcp kms delete params
func (o *GcpKmsDeleteParams) WithContext(ctx context.Context) *GcpKmsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gcp kms delete params
func (o *GcpKmsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gcp kms delete params
func (o *GcpKmsDeleteParams) WithHTTPClient(client *http.Client) *GcpKmsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gcp kms delete params
func (o *GcpKmsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeout adds the returnTimeout to the gcp kms delete params
func (o *GcpKmsDeleteParams) WithReturnTimeout(returnTimeout *int64) *GcpKmsDeleteParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the gcp kms delete params
func (o *GcpKmsDeleteParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithUUID adds the uuid to the gcp kms delete params
func (o *GcpKmsDeleteParams) WithUUID(uuid string) *GcpKmsDeleteParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the gcp kms delete params
func (o *GcpKmsDeleteParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *GcpKmsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

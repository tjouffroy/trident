// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewFpolicyPersistentStoreCreateParams creates a new FpolicyPersistentStoreCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyPersistentStoreCreateParams() *FpolicyPersistentStoreCreateParams {
	return &FpolicyPersistentStoreCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyPersistentStoreCreateParamsWithTimeout creates a new FpolicyPersistentStoreCreateParams object
// with the ability to set a timeout on a request.
func NewFpolicyPersistentStoreCreateParamsWithTimeout(timeout time.Duration) *FpolicyPersistentStoreCreateParams {
	return &FpolicyPersistentStoreCreateParams{
		timeout: timeout,
	}
}

// NewFpolicyPersistentStoreCreateParamsWithContext creates a new FpolicyPersistentStoreCreateParams object
// with the ability to set a context for a request.
func NewFpolicyPersistentStoreCreateParamsWithContext(ctx context.Context) *FpolicyPersistentStoreCreateParams {
	return &FpolicyPersistentStoreCreateParams{
		Context: ctx,
	}
}

// NewFpolicyPersistentStoreCreateParamsWithHTTPClient creates a new FpolicyPersistentStoreCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyPersistentStoreCreateParamsWithHTTPClient(client *http.Client) *FpolicyPersistentStoreCreateParams {
	return &FpolicyPersistentStoreCreateParams{
		HTTPClient: client,
	}
}

/*
FpolicyPersistentStoreCreateParams contains all the parameters to send to the API endpoint

	for the fpolicy persistent store create operation.

	Typically these are written to a http.Request.
*/
type FpolicyPersistentStoreCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.FpolicyPersistentStore

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecords *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy persistent store create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyPersistentStoreCreateParams) WithDefaults() *FpolicyPersistentStoreCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy persistent store create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyPersistentStoreCreateParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(false)
	)

	val := FpolicyPersistentStoreCreateParams{
		ReturnRecords: &returnRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) WithTimeout(timeout time.Duration) *FpolicyPersistentStoreCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) WithContext(ctx context.Context) *FpolicyPersistentStoreCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) WithHTTPClient(client *http.Client) *FpolicyPersistentStoreCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) WithInfo(info *models.FpolicyPersistentStore) *FpolicyPersistentStoreCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) SetInfo(info *models.FpolicyPersistentStore) {
	o.Info = info
}

// WithReturnRecords adds the returnRecords to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) WithReturnRecords(returnRecords *bool) *FpolicyPersistentStoreCreateParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithSvmUUID adds the svmUUID to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) WithSvmUUID(svmUUID string) *FpolicyPersistentStoreCreateParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the fpolicy persistent store create params
func (o *FpolicyPersistentStoreCreateParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyPersistentStoreCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

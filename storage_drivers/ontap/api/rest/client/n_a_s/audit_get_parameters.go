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
)

// NewAuditGetParams creates a new AuditGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuditGetParams() *AuditGetParams {
	return &AuditGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuditGetParamsWithTimeout creates a new AuditGetParams object
// with the ability to set a timeout on a request.
func NewAuditGetParamsWithTimeout(timeout time.Duration) *AuditGetParams {
	return &AuditGetParams{
		timeout: timeout,
	}
}

// NewAuditGetParamsWithContext creates a new AuditGetParams object
// with the ability to set a context for a request.
func NewAuditGetParamsWithContext(ctx context.Context) *AuditGetParams {
	return &AuditGetParams{
		Context: ctx,
	}
}

// NewAuditGetParamsWithHTTPClient creates a new AuditGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuditGetParamsWithHTTPClient(client *http.Client) *AuditGetParams {
	return &AuditGetParams{
		HTTPClient: client,
	}
}

/*
AuditGetParams contains all the parameters to send to the API endpoint

	for the audit get operation.

	Typically these are written to a http.Request.
*/
type AuditGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditGetParams) WithDefaults() *AuditGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuditGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the audit get params
func (o *AuditGetParams) WithTimeout(timeout time.Duration) *AuditGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the audit get params
func (o *AuditGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the audit get params
func (o *AuditGetParams) WithContext(ctx context.Context) *AuditGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the audit get params
func (o *AuditGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the audit get params
func (o *AuditGetParams) WithHTTPClient(client *http.Client) *AuditGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the audit get params
func (o *AuditGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the audit get params
func (o *AuditGetParams) WithFields(fields []string) *AuditGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the audit get params
func (o *AuditGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithSvmUUID adds the svmUUID to the audit get params
func (o *AuditGetParams) WithSvmUUID(svmUUID string) *AuditGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the audit get params
func (o *AuditGetParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *AuditGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
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

// bindParamAuditGet binds the parameter fields
func (o *AuditGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

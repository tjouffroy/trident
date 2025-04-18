// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewStorageUnitSnapshotModifyParams creates a new StorageUnitSnapshotModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStorageUnitSnapshotModifyParams() *StorageUnitSnapshotModifyParams {
	return &StorageUnitSnapshotModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStorageUnitSnapshotModifyParamsWithTimeout creates a new StorageUnitSnapshotModifyParams object
// with the ability to set a timeout on a request.
func NewStorageUnitSnapshotModifyParamsWithTimeout(timeout time.Duration) *StorageUnitSnapshotModifyParams {
	return &StorageUnitSnapshotModifyParams{
		timeout: timeout,
	}
}

// NewStorageUnitSnapshotModifyParamsWithContext creates a new StorageUnitSnapshotModifyParams object
// with the ability to set a context for a request.
func NewStorageUnitSnapshotModifyParamsWithContext(ctx context.Context) *StorageUnitSnapshotModifyParams {
	return &StorageUnitSnapshotModifyParams{
		Context: ctx,
	}
}

// NewStorageUnitSnapshotModifyParamsWithHTTPClient creates a new StorageUnitSnapshotModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewStorageUnitSnapshotModifyParamsWithHTTPClient(client *http.Client) *StorageUnitSnapshotModifyParams {
	return &StorageUnitSnapshotModifyParams{
		HTTPClient: client,
	}
}

/*
StorageUnitSnapshotModifyParams contains all the parameters to send to the API endpoint

	for the storage unit snapshot modify operation.

	Typically these are written to a http.Request.
*/
type StorageUnitSnapshotModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.StorageUnitSnapshot

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeout *int64

	/* StorageUnitUUID.

	   Storage Unit UUID
	*/
	StorageUnitUUID string

	/* UUID.

	   Snapshot UUID
	*/
	UUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage unit snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageUnitSnapshotModifyParams) WithDefaults() *StorageUnitSnapshotModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage unit snapshot modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StorageUnitSnapshotModifyParams) SetDefaults() {
	var (
		returnTimeoutDefault = int64(0)
	)

	val := StorageUnitSnapshotModifyParams{
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithTimeout(timeout time.Duration) *StorageUnitSnapshotModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithContext(ctx context.Context) *StorageUnitSnapshotModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithHTTPClient(client *http.Client) *StorageUnitSnapshotModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithInfo(info *models.StorageUnitSnapshot) *StorageUnitSnapshotModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetInfo(info *models.StorageUnitSnapshot) {
	o.Info = info
}

// WithReturnTimeout adds the returnTimeout to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithReturnTimeout(returnTimeout *int64) *StorageUnitSnapshotModifyParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithStorageUnitUUID adds the storageUnitUUID to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithStorageUnitUUID(storageUnitUUID string) *StorageUnitSnapshotModifyParams {
	o.SetStorageUnitUUID(storageUnitUUID)
	return o
}

// SetStorageUnitUUID adds the storageUnitUuid to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetStorageUnitUUID(storageUnitUUID string) {
	o.StorageUnitUUID = storageUnitUUID
}

// WithUUID adds the uuid to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) WithUUID(uuid string) *StorageUnitSnapshotModifyParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the storage unit snapshot modify params
func (o *StorageUnitSnapshotModifyParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StorageUnitSnapshotModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

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

	// path param storage_unit.uuid
	if err := r.SetPathParam("storage_unit.uuid", o.StorageUnitUUID); err != nil {
		return err
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

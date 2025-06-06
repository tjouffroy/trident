// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewCloudStoreCollectionGetParams creates a new CloudStoreCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudStoreCollectionGetParams() *CloudStoreCollectionGetParams {
	return &CloudStoreCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudStoreCollectionGetParamsWithTimeout creates a new CloudStoreCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCloudStoreCollectionGetParamsWithTimeout(timeout time.Duration) *CloudStoreCollectionGetParams {
	return &CloudStoreCollectionGetParams{
		timeout: timeout,
	}
}

// NewCloudStoreCollectionGetParamsWithContext creates a new CloudStoreCollectionGetParams object
// with the ability to set a context for a request.
func NewCloudStoreCollectionGetParamsWithContext(ctx context.Context) *CloudStoreCollectionGetParams {
	return &CloudStoreCollectionGetParams{
		Context: ctx,
	}
}

// NewCloudStoreCollectionGetParamsWithHTTPClient creates a new CloudStoreCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudStoreCollectionGetParamsWithHTTPClient(client *http.Client) *CloudStoreCollectionGetParams {
	return &CloudStoreCollectionGetParams{
		HTTPClient: client,
	}
}

/*
CloudStoreCollectionGetParams contains all the parameters to send to the API endpoint

	for the cloud store collection get operation.

	Typically these are written to a http.Request.
*/
type CloudStoreCollectionGetParams struct {

	/* AggregateName.

	   Filter by aggregate.name
	*/
	AggregateName *string

	/* AggregateUUID.

	   Aggregate UUID
	*/
	AggregateUUID string

	/* Availability.

	   Filter by availability
	*/
	Availability *string

	/* AvailabilityAtPartner.

	   Filter by availability_at_partner
	*/
	AvailabilityAtPartner *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* MirrorDegraded.

	   Filter by mirror_degraded
	*/
	MirrorDegraded *bool

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Primary.

	   Filter by primary
	*/
	Primary *bool

	/* ResyncProgress.

	   Filter by resync-progress
	*/
	ResyncProgress *int64

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* TargetName.

	   Filter by target.name
	*/
	TargetName *string

	/* TargetUUID.

	   Filter by target.uuid
	*/
	TargetUUID *string

	/* UnavailableReasonMessage.

	   Filter by unavailable_reason.message
	*/
	UnavailableReasonMessage *string

	/* UnreclaimedSpaceThreshold.

	   Filter by unreclaimed_space_threshold
	*/
	UnreclaimedSpaceThreshold *int64

	/* Used.

	   Filter by used
	*/
	Used *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud store collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudStoreCollectionGetParams) WithDefaults() *CloudStoreCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud store collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudStoreCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := CloudStoreCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithTimeout(timeout time.Duration) *CloudStoreCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithContext(ctx context.Context) *CloudStoreCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithHTTPClient(client *http.Client) *CloudStoreCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregateName adds the aggregateName to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithAggregateName(aggregateName *string) *CloudStoreCollectionGetParams {
	o.SetAggregateName(aggregateName)
	return o
}

// SetAggregateName adds the aggregateName to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetAggregateName(aggregateName *string) {
	o.AggregateName = aggregateName
}

// WithAggregateUUID adds the aggregateUUID to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithAggregateUUID(aggregateUUID string) *CloudStoreCollectionGetParams {
	o.SetAggregateUUID(aggregateUUID)
	return o
}

// SetAggregateUUID adds the aggregateUuid to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetAggregateUUID(aggregateUUID string) {
	o.AggregateUUID = aggregateUUID
}

// WithAvailability adds the availability to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithAvailability(availability *string) *CloudStoreCollectionGetParams {
	o.SetAvailability(availability)
	return o
}

// SetAvailability adds the availability to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetAvailability(availability *string) {
	o.Availability = availability
}

// WithAvailabilityAtPartner adds the availabilityAtPartner to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithAvailabilityAtPartner(availabilityAtPartner *string) *CloudStoreCollectionGetParams {
	o.SetAvailabilityAtPartner(availabilityAtPartner)
	return o
}

// SetAvailabilityAtPartner adds the availabilityAtPartner to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetAvailabilityAtPartner(availabilityAtPartner *string) {
	o.AvailabilityAtPartner = availabilityAtPartner
}

// WithFields adds the fields to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithFields(fields []string) *CloudStoreCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithMaxRecords adds the maxRecords to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithMaxRecords(maxRecords *int64) *CloudStoreCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithMirrorDegraded adds the mirrorDegraded to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithMirrorDegraded(mirrorDegraded *bool) *CloudStoreCollectionGetParams {
	o.SetMirrorDegraded(mirrorDegraded)
	return o
}

// SetMirrorDegraded adds the mirrorDegraded to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetMirrorDegraded(mirrorDegraded *bool) {
	o.MirrorDegraded = mirrorDegraded
}

// WithOrderBy adds the orderBy to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithOrderBy(orderBy []string) *CloudStoreCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPrimary adds the primary to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithPrimary(primary *bool) *CloudStoreCollectionGetParams {
	o.SetPrimary(primary)
	return o
}

// SetPrimary adds the primary to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetPrimary(primary *bool) {
	o.Primary = primary
}

// WithResyncProgress adds the resyncProgress to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithResyncProgress(resyncProgress *int64) *CloudStoreCollectionGetParams {
	o.SetResyncProgress(resyncProgress)
	return o
}

// SetResyncProgress adds the resyncProgress to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetResyncProgress(resyncProgress *int64) {
	o.ResyncProgress = resyncProgress
}

// WithReturnRecords adds the returnRecords to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithReturnRecords(returnRecords *bool) *CloudStoreCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *CloudStoreCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithTargetName adds the targetName to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithTargetName(targetName *string) *CloudStoreCollectionGetParams {
	o.SetTargetName(targetName)
	return o
}

// SetTargetName adds the targetName to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetTargetName(targetName *string) {
	o.TargetName = targetName
}

// WithTargetUUID adds the targetUUID to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithTargetUUID(targetUUID *string) *CloudStoreCollectionGetParams {
	o.SetTargetUUID(targetUUID)
	return o
}

// SetTargetUUID adds the targetUuid to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetTargetUUID(targetUUID *string) {
	o.TargetUUID = targetUUID
}

// WithUnavailableReasonMessage adds the unavailableReasonMessage to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithUnavailableReasonMessage(unavailableReasonMessage *string) *CloudStoreCollectionGetParams {
	o.SetUnavailableReasonMessage(unavailableReasonMessage)
	return o
}

// SetUnavailableReasonMessage adds the unavailableReasonMessage to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetUnavailableReasonMessage(unavailableReasonMessage *string) {
	o.UnavailableReasonMessage = unavailableReasonMessage
}

// WithUnreclaimedSpaceThreshold adds the unreclaimedSpaceThreshold to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithUnreclaimedSpaceThreshold(unreclaimedSpaceThreshold *int64) *CloudStoreCollectionGetParams {
	o.SetUnreclaimedSpaceThreshold(unreclaimedSpaceThreshold)
	return o
}

// SetUnreclaimedSpaceThreshold adds the unreclaimedSpaceThreshold to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetUnreclaimedSpaceThreshold(unreclaimedSpaceThreshold *int64) {
	o.UnreclaimedSpaceThreshold = unreclaimedSpaceThreshold
}

// WithUsed adds the used to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) WithUsed(used *int64) *CloudStoreCollectionGetParams {
	o.SetUsed(used)
	return o
}

// SetUsed adds the used to the cloud store collection get params
func (o *CloudStoreCollectionGetParams) SetUsed(used *int64) {
	o.Used = used
}

// WriteToRequest writes these params to a swagger request
func (o *CloudStoreCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AggregateName != nil {

		// query param aggregate.name
		var qrAggregateName string

		if o.AggregateName != nil {
			qrAggregateName = *o.AggregateName
		}
		qAggregateName := qrAggregateName
		if qAggregateName != "" {

			if err := r.SetQueryParam("aggregate.name", qAggregateName); err != nil {
				return err
			}
		}
	}

	// path param aggregate.uuid
	if err := r.SetPathParam("aggregate.uuid", o.AggregateUUID); err != nil {
		return err
	}

	if o.Availability != nil {

		// query param availability
		var qrAvailability string

		if o.Availability != nil {
			qrAvailability = *o.Availability
		}
		qAvailability := qrAvailability
		if qAvailability != "" {

			if err := r.SetQueryParam("availability", qAvailability); err != nil {
				return err
			}
		}
	}

	if o.AvailabilityAtPartner != nil {

		// query param availability_at_partner
		var qrAvailabilityAtPartner string

		if o.AvailabilityAtPartner != nil {
			qrAvailabilityAtPartner = *o.AvailabilityAtPartner
		}
		qAvailabilityAtPartner := qrAvailabilityAtPartner
		if qAvailabilityAtPartner != "" {

			if err := r.SetQueryParam("availability_at_partner", qAvailabilityAtPartner); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.MirrorDegraded != nil {

		// query param mirror_degraded
		var qrMirrorDegraded bool

		if o.MirrorDegraded != nil {
			qrMirrorDegraded = *o.MirrorDegraded
		}
		qMirrorDegraded := swag.FormatBool(qrMirrorDegraded)
		if qMirrorDegraded != "" {

			if err := r.SetQueryParam("mirror_degraded", qMirrorDegraded); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.Primary != nil {

		// query param primary
		var qrPrimary bool

		if o.Primary != nil {
			qrPrimary = *o.Primary
		}
		qPrimary := swag.FormatBool(qrPrimary)
		if qPrimary != "" {

			if err := r.SetQueryParam("primary", qPrimary); err != nil {
				return err
			}
		}
	}

	if o.ResyncProgress != nil {

		// query param resync-progress
		var qrResyncProgress int64

		if o.ResyncProgress != nil {
			qrResyncProgress = *o.ResyncProgress
		}
		qResyncProgress := swag.FormatInt64(qrResyncProgress)
		if qResyncProgress != "" {

			if err := r.SetQueryParam("resync-progress", qResyncProgress); err != nil {
				return err
			}
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

	if o.TargetName != nil {

		// query param target.name
		var qrTargetName string

		if o.TargetName != nil {
			qrTargetName = *o.TargetName
		}
		qTargetName := qrTargetName
		if qTargetName != "" {

			if err := r.SetQueryParam("target.name", qTargetName); err != nil {
				return err
			}
		}
	}

	if o.TargetUUID != nil {

		// query param target.uuid
		var qrTargetUUID string

		if o.TargetUUID != nil {
			qrTargetUUID = *o.TargetUUID
		}
		qTargetUUID := qrTargetUUID
		if qTargetUUID != "" {

			if err := r.SetQueryParam("target.uuid", qTargetUUID); err != nil {
				return err
			}
		}
	}

	if o.UnavailableReasonMessage != nil {

		// query param unavailable_reason.message
		var qrUnavailableReasonMessage string

		if o.UnavailableReasonMessage != nil {
			qrUnavailableReasonMessage = *o.UnavailableReasonMessage
		}
		qUnavailableReasonMessage := qrUnavailableReasonMessage
		if qUnavailableReasonMessage != "" {

			if err := r.SetQueryParam("unavailable_reason.message", qUnavailableReasonMessage); err != nil {
				return err
			}
		}
	}

	if o.UnreclaimedSpaceThreshold != nil {

		// query param unreclaimed_space_threshold
		var qrUnreclaimedSpaceThreshold int64

		if o.UnreclaimedSpaceThreshold != nil {
			qrUnreclaimedSpaceThreshold = *o.UnreclaimedSpaceThreshold
		}
		qUnreclaimedSpaceThreshold := swag.FormatInt64(qrUnreclaimedSpaceThreshold)
		if qUnreclaimedSpaceThreshold != "" {

			if err := r.SetQueryParam("unreclaimed_space_threshold", qUnreclaimedSpaceThreshold); err != nil {
				return err
			}
		}
	}

	if o.Used != nil {

		// query param used
		var qrUsed int64

		if o.Used != nil {
			qrUsed = *o.Used
		}
		qUsed := swag.FormatInt64(qrUsed)
		if qUsed != "" {

			if err := r.SetQueryParam("used", qUsed); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCloudStoreCollectionGet binds the parameter fields
func (o *CloudStoreCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCloudStoreCollectionGet binds the parameter order_by
func (o *CloudStoreCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}

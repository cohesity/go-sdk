// Code generated by go-swagger; DO NOT EDIT.

package object

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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewConstructMetaInfoParams creates a new ConstructMetaInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConstructMetaInfoParams() *ConstructMetaInfoParams {
	return &ConstructMetaInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConstructMetaInfoParamsWithTimeout creates a new ConstructMetaInfoParams object
// with the ability to set a timeout on a request.
func NewConstructMetaInfoParamsWithTimeout(timeout time.Duration) *ConstructMetaInfoParams {
	return &ConstructMetaInfoParams{
		timeout: timeout,
	}
}

// NewConstructMetaInfoParamsWithContext creates a new ConstructMetaInfoParams object
// with the ability to set a context for a request.
func NewConstructMetaInfoParamsWithContext(ctx context.Context) *ConstructMetaInfoParams {
	return &ConstructMetaInfoParams{
		Context: ctx,
	}
}

// NewConstructMetaInfoParamsWithHTTPClient creates a new ConstructMetaInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewConstructMetaInfoParamsWithHTTPClient(client *http.Client) *ConstructMetaInfoParams {
	return &ConstructMetaInfoParams{
		HTTPClient: client,
	}
}

/*
ConstructMetaInfoParams contains all the parameters to send to the API endpoint

	for the construct meta info operation.

	Typically these are written to a http.Request.
*/
type ConstructMetaInfoParams struct {

	/* Body.

	   Specifies the parameters to construct meta info for desired workflow.
	*/
	Body *models.ConstructMetaInfoRequest

	/* SnapshotID.

	   Specifies the snapshot id.
	*/
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the construct meta info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstructMetaInfoParams) WithDefaults() *ConstructMetaInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the construct meta info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConstructMetaInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the construct meta info params
func (o *ConstructMetaInfoParams) WithTimeout(timeout time.Duration) *ConstructMetaInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the construct meta info params
func (o *ConstructMetaInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the construct meta info params
func (o *ConstructMetaInfoParams) WithContext(ctx context.Context) *ConstructMetaInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the construct meta info params
func (o *ConstructMetaInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the construct meta info params
func (o *ConstructMetaInfoParams) WithHTTPClient(client *http.Client) *ConstructMetaInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the construct meta info params
func (o *ConstructMetaInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the construct meta info params
func (o *ConstructMetaInfoParams) WithBody(body *models.ConstructMetaInfoRequest) *ConstructMetaInfoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the construct meta info params
func (o *ConstructMetaInfoParams) SetBody(body *models.ConstructMetaInfoRequest) {
	o.Body = body
}

// WithSnapshotID adds the snapshotID to the construct meta info params
func (o *ConstructMetaInfoParams) WithSnapshotID(snapshotID string) *ConstructMetaInfoParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the construct meta info params
func (o *ConstructMetaInfoParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *ConstructMetaInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param snapshotId
	if err := r.SetPathParam("snapshotId", o.SnapshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

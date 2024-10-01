// Code generated by go-swagger; DO NOT EDIT.

package disks

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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// NewUpdateDiskParams creates a new UpdateDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDiskParams() *UpdateDiskParams {
	return &UpdateDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDiskParamsWithTimeout creates a new UpdateDiskParams object
// with the ability to set a timeout on a request.
func NewUpdateDiskParamsWithTimeout(timeout time.Duration) *UpdateDiskParams {
	return &UpdateDiskParams{
		timeout: timeout,
	}
}

// NewUpdateDiskParamsWithContext creates a new UpdateDiskParams object
// with the ability to set a context for a request.
func NewUpdateDiskParamsWithContext(ctx context.Context) *UpdateDiskParams {
	return &UpdateDiskParams{
		Context: ctx,
	}
}

// NewUpdateDiskParamsWithHTTPClient creates a new UpdateDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDiskParamsWithHTTPClient(client *http.Client) *UpdateDiskParams {
	return &UpdateDiskParams{
		HTTPClient: client,
	}
}

/*
UpdateDiskParams contains all the parameters to send to the API endpoint

	for the update disk operation.

	Typically these are written to a http.Request.
*/
type UpdateDiskParams struct {

	// Body.
	Body *models.DiskInternal

	/* ID.

	   Id of the Disk

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDiskParams) WithDefaults() *UpdateDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDiskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update disk params
func (o *UpdateDiskParams) WithTimeout(timeout time.Duration) *UpdateDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update disk params
func (o *UpdateDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update disk params
func (o *UpdateDiskParams) WithContext(ctx context.Context) *UpdateDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update disk params
func (o *UpdateDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update disk params
func (o *UpdateDiskParams) WithHTTPClient(client *http.Client) *UpdateDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update disk params
func (o *UpdateDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update disk params
func (o *UpdateDiskParams) WithBody(body *models.DiskInternal) *UpdateDiskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update disk params
func (o *UpdateDiskParams) SetBody(body *models.DiskInternal) {
	o.Body = body
}

// WithID adds the id to the update disk params
func (o *UpdateDiskParams) WithID(id int64) *UpdateDiskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update disk params
func (o *UpdateDiskParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

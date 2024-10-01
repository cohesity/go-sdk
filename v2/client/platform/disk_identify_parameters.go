// Code generated by go-swagger; DO NOT EDIT.

package platform

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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// NewDiskIdentifyParams creates a new DiskIdentifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDiskIdentifyParams() *DiskIdentifyParams {
	return &DiskIdentifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDiskIdentifyParamsWithTimeout creates a new DiskIdentifyParams object
// with the ability to set a timeout on a request.
func NewDiskIdentifyParamsWithTimeout(timeout time.Duration) *DiskIdentifyParams {
	return &DiskIdentifyParams{
		timeout: timeout,
	}
}

// NewDiskIdentifyParamsWithContext creates a new DiskIdentifyParams object
// with the ability to set a context for a request.
func NewDiskIdentifyParamsWithContext(ctx context.Context) *DiskIdentifyParams {
	return &DiskIdentifyParams{
		Context: ctx,
	}
}

// NewDiskIdentifyParamsWithHTTPClient creates a new DiskIdentifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDiskIdentifyParamsWithHTTPClient(client *http.Client) *DiskIdentifyParams {
	return &DiskIdentifyParams{
		HTTPClient: client,
	}
}

/*
DiskIdentifyParams contains all the parameters to send to the API endpoint

	for the disk identify operation.

	Typically these are written to a http.Request.
*/
type DiskIdentifyParams struct {

	/* Body.

	   Specifies the parameter to identify disk.
	*/
	Body *models.DiskIdentify

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the disk identify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiskIdentifyParams) WithDefaults() *DiskIdentifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the disk identify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DiskIdentifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the disk identify params
func (o *DiskIdentifyParams) WithTimeout(timeout time.Duration) *DiskIdentifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disk identify params
func (o *DiskIdentifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disk identify params
func (o *DiskIdentifyParams) WithContext(ctx context.Context) *DiskIdentifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disk identify params
func (o *DiskIdentifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disk identify params
func (o *DiskIdentifyParams) WithHTTPClient(client *http.Client) *DiskIdentifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disk identify params
func (o *DiskIdentifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the disk identify params
func (o *DiskIdentifyParams) WithBody(body *models.DiskIdentify) *DiskIdentifyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the disk identify params
func (o *DiskIdentifyParams) SetBody(body *models.DiskIdentify) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DiskIdentifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

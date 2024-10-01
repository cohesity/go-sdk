// Code generated by go-swagger; DO NOT EDIT.

package ip

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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// NewUnconfigureIPParams creates a new UnconfigureIPParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnconfigureIPParams() *UnconfigureIPParams {
	return &UnconfigureIPParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnconfigureIPParamsWithTimeout creates a new UnconfigureIPParams object
// with the ability to set a timeout on a request.
func NewUnconfigureIPParamsWithTimeout(timeout time.Duration) *UnconfigureIPParams {
	return &UnconfigureIPParams{
		timeout: timeout,
	}
}

// NewUnconfigureIPParamsWithContext creates a new UnconfigureIPParams object
// with the ability to set a context for a request.
func NewUnconfigureIPParamsWithContext(ctx context.Context) *UnconfigureIPParams {
	return &UnconfigureIPParams{
		Context: ctx,
	}
}

// NewUnconfigureIPParamsWithHTTPClient creates a new UnconfigureIPParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnconfigureIPParamsWithHTTPClient(client *http.Client) *UnconfigureIPParams {
	return &UnconfigureIPParams{
		HTTPClient: client,
	}
}

/*
UnconfigureIPParams contains all the parameters to send to the API endpoint

	for the unconfigure Ip operation.

	Typically these are written to a http.Request.
*/
type UnconfigureIPParams struct {

	// Body.
	Body *models.IPUnconfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unconfigure Ip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnconfigureIPParams) WithDefaults() *UnconfigureIPParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unconfigure Ip params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnconfigureIPParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unconfigure Ip params
func (o *UnconfigureIPParams) WithTimeout(timeout time.Duration) *UnconfigureIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unconfigure Ip params
func (o *UnconfigureIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unconfigure Ip params
func (o *UnconfigureIPParams) WithContext(ctx context.Context) *UnconfigureIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unconfigure Ip params
func (o *UnconfigureIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unconfigure Ip params
func (o *UnconfigureIPParams) WithHTTPClient(client *http.Client) *UnconfigureIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unconfigure Ip params
func (o *UnconfigureIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unconfigure Ip params
func (o *UnconfigureIPParams) WithBody(body *models.IPUnconfig) *UnconfigureIPParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unconfigure Ip params
func (o *UnconfigureIPParams) SetBody(body *models.IPUnconfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UnconfigureIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

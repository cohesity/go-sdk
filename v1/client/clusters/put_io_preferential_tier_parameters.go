// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewPutIoPreferentialTierParams creates a new PutIoPreferentialTierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutIoPreferentialTierParams() *PutIoPreferentialTierParams {
	return &PutIoPreferentialTierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutIoPreferentialTierParamsWithTimeout creates a new PutIoPreferentialTierParams object
// with the ability to set a timeout on a request.
func NewPutIoPreferentialTierParamsWithTimeout(timeout time.Duration) *PutIoPreferentialTierParams {
	return &PutIoPreferentialTierParams{
		timeout: timeout,
	}
}

// NewPutIoPreferentialTierParamsWithContext creates a new PutIoPreferentialTierParams object
// with the ability to set a context for a request.
func NewPutIoPreferentialTierParamsWithContext(ctx context.Context) *PutIoPreferentialTierParams {
	return &PutIoPreferentialTierParams{
		Context: ctx,
	}
}

// NewPutIoPreferentialTierParamsWithHTTPClient creates a new PutIoPreferentialTierParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutIoPreferentialTierParamsWithHTTPClient(client *http.Client) *PutIoPreferentialTierParams {
	return &PutIoPreferentialTierParams{
		HTTPClient: client,
	}
}

/*
PutIoPreferentialTierParams contains all the parameters to send to the API endpoint

	for the put io preferential tier operation.

	Typically these are written to a http.Request.
*/
type PutIoPreferentialTierParams struct {

	// Body.
	Body *models.IoPreferentialTier

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put io preferential tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIoPreferentialTierParams) WithDefaults() *PutIoPreferentialTierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put io preferential tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutIoPreferentialTierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put io preferential tier params
func (o *PutIoPreferentialTierParams) WithTimeout(timeout time.Duration) *PutIoPreferentialTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put io preferential tier params
func (o *PutIoPreferentialTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put io preferential tier params
func (o *PutIoPreferentialTierParams) WithContext(ctx context.Context) *PutIoPreferentialTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put io preferential tier params
func (o *PutIoPreferentialTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put io preferential tier params
func (o *PutIoPreferentialTierParams) WithHTTPClient(client *http.Client) *PutIoPreferentialTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put io preferential tier params
func (o *PutIoPreferentialTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put io preferential tier params
func (o *PutIoPreferentialTierParams) WithBody(body *models.IoPreferentialTier) *PutIoPreferentialTierParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put io preferential tier params
func (o *PutIoPreferentialTierParams) SetBody(body *models.IoPreferentialTier) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PutIoPreferentialTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

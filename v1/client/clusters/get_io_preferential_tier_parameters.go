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
)

// NewGetIoPreferentialTierParams creates a new GetIoPreferentialTierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIoPreferentialTierParams() *GetIoPreferentialTierParams {
	return &GetIoPreferentialTierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIoPreferentialTierParamsWithTimeout creates a new GetIoPreferentialTierParams object
// with the ability to set a timeout on a request.
func NewGetIoPreferentialTierParamsWithTimeout(timeout time.Duration) *GetIoPreferentialTierParams {
	return &GetIoPreferentialTierParams{
		timeout: timeout,
	}
}

// NewGetIoPreferentialTierParamsWithContext creates a new GetIoPreferentialTierParams object
// with the ability to set a context for a request.
func NewGetIoPreferentialTierParamsWithContext(ctx context.Context) *GetIoPreferentialTierParams {
	return &GetIoPreferentialTierParams{
		Context: ctx,
	}
}

// NewGetIoPreferentialTierParamsWithHTTPClient creates a new GetIoPreferentialTierParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIoPreferentialTierParamsWithHTTPClient(client *http.Client) *GetIoPreferentialTierParams {
	return &GetIoPreferentialTierParams{
		HTTPClient: client,
	}
}

/*
GetIoPreferentialTierParams contains all the parameters to send to the API endpoint

	for the get io preferential tier operation.

	Typically these are written to a http.Request.
*/
type GetIoPreferentialTierParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get io preferential tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIoPreferentialTierParams) WithDefaults() *GetIoPreferentialTierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get io preferential tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIoPreferentialTierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get io preferential tier params
func (o *GetIoPreferentialTierParams) WithTimeout(timeout time.Duration) *GetIoPreferentialTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get io preferential tier params
func (o *GetIoPreferentialTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get io preferential tier params
func (o *GetIoPreferentialTierParams) WithContext(ctx context.Context) *GetIoPreferentialTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get io preferential tier params
func (o *GetIoPreferentialTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get io preferential tier params
func (o *GetIoPreferentialTierParams) WithHTTPClient(client *http.Client) *GetIoPreferentialTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get io preferential tier params
func (o *GetIoPreferentialTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIoPreferentialTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

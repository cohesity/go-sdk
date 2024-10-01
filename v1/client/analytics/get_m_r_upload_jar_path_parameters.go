// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// NewGetMRUploadJarPathParams creates a new GetMRUploadJarPathParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMRUploadJarPathParams() *GetMRUploadJarPathParams {
	return &GetMRUploadJarPathParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMRUploadJarPathParamsWithTimeout creates a new GetMRUploadJarPathParams object
// with the ability to set a timeout on a request.
func NewGetMRUploadJarPathParamsWithTimeout(timeout time.Duration) *GetMRUploadJarPathParams {
	return &GetMRUploadJarPathParams{
		timeout: timeout,
	}
}

// NewGetMRUploadJarPathParamsWithContext creates a new GetMRUploadJarPathParams object
// with the ability to set a context for a request.
func NewGetMRUploadJarPathParamsWithContext(ctx context.Context) *GetMRUploadJarPathParams {
	return &GetMRUploadJarPathParams{
		Context: ctx,
	}
}

// NewGetMRUploadJarPathParamsWithHTTPClient creates a new GetMRUploadJarPathParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMRUploadJarPathParamsWithHTTPClient(client *http.Client) *GetMRUploadJarPathParams {
	return &GetMRUploadJarPathParams{
		HTTPClient: client,
	}
}

/*
GetMRUploadJarPathParams contains all the parameters to send to the API endpoint

	for the get m r upload jar path operation.

	Typically these are written to a http.Request.
*/
type GetMRUploadJarPathParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get m r upload jar path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMRUploadJarPathParams) WithDefaults() *GetMRUploadJarPathParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get m r upload jar path params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMRUploadJarPathParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get m r upload jar path params
func (o *GetMRUploadJarPathParams) WithTimeout(timeout time.Duration) *GetMRUploadJarPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get m r upload jar path params
func (o *GetMRUploadJarPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get m r upload jar path params
func (o *GetMRUploadJarPathParams) WithContext(ctx context.Context) *GetMRUploadJarPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get m r upload jar path params
func (o *GetMRUploadJarPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get m r upload jar path params
func (o *GetMRUploadJarPathParams) WithHTTPClient(client *http.Client) *GetMRUploadJarPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get m r upload jar path params
func (o *GetMRUploadJarPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetMRUploadJarPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

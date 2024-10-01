// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewUnregisterTrustedCaParams creates a new UnregisterTrustedCaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnregisterTrustedCaParams() *UnregisterTrustedCaParams {
	return &UnregisterTrustedCaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnregisterTrustedCaParamsWithTimeout creates a new UnregisterTrustedCaParams object
// with the ability to set a timeout on a request.
func NewUnregisterTrustedCaParamsWithTimeout(timeout time.Duration) *UnregisterTrustedCaParams {
	return &UnregisterTrustedCaParams{
		timeout: timeout,
	}
}

// NewUnregisterTrustedCaParamsWithContext creates a new UnregisterTrustedCaParams object
// with the ability to set a context for a request.
func NewUnregisterTrustedCaParamsWithContext(ctx context.Context) *UnregisterTrustedCaParams {
	return &UnregisterTrustedCaParams{
		Context: ctx,
	}
}

// NewUnregisterTrustedCaParamsWithHTTPClient creates a new UnregisterTrustedCaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnregisterTrustedCaParamsWithHTTPClient(client *http.Client) *UnregisterTrustedCaParams {
	return &UnregisterTrustedCaParams{
		HTTPClient: client,
	}
}

/*
UnregisterTrustedCaParams contains all the parameters to send to the API endpoint

	for the unregister trusted ca operation.

	Typically these are written to a http.Request.
*/
type UnregisterTrustedCaParams struct {

	/* ID.

	   Specifies the id of the certificate to be unregistered.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unregister trusted ca params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnregisterTrustedCaParams) WithDefaults() *UnregisterTrustedCaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unregister trusted ca params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnregisterTrustedCaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) WithTimeout(timeout time.Duration) *UnregisterTrustedCaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) WithContext(ctx context.Context) *UnregisterTrustedCaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) WithHTTPClient(client *http.Client) *UnregisterTrustedCaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) WithID(id string) *UnregisterTrustedCaParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unregister trusted ca params
func (o *UnregisterTrustedCaParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnregisterTrustedCaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

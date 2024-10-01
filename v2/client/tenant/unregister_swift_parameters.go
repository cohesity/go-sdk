// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewUnregisterSwiftParams creates a new UnregisterSwiftParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnregisterSwiftParams() *UnregisterSwiftParams {
	return &UnregisterSwiftParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnregisterSwiftParamsWithTimeout creates a new UnregisterSwiftParams object
// with the ability to set a timeout on a request.
func NewUnregisterSwiftParamsWithTimeout(timeout time.Duration) *UnregisterSwiftParams {
	return &UnregisterSwiftParams{
		timeout: timeout,
	}
}

// NewUnregisterSwiftParamsWithContext creates a new UnregisterSwiftParams object
// with the ability to set a context for a request.
func NewUnregisterSwiftParamsWithContext(ctx context.Context) *UnregisterSwiftParams {
	return &UnregisterSwiftParams{
		Context: ctx,
	}
}

// NewUnregisterSwiftParamsWithHTTPClient creates a new UnregisterSwiftParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnregisterSwiftParamsWithHTTPClient(client *http.Client) *UnregisterSwiftParams {
	return &UnregisterSwiftParams{
		HTTPClient: client,
	}
}

/*
UnregisterSwiftParams contains all the parameters to send to the API endpoint

	for the unregister swift operation.

	Typically these are written to a http.Request.
*/
type UnregisterSwiftParams struct {

	/* Body.

	   Specifies the parameters to unregister a Swift service from Keystone server.
	*/
	Body *models.UnregisterSwiftParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unregister swift params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnregisterSwiftParams) WithDefaults() *UnregisterSwiftParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unregister swift params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnregisterSwiftParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unregister swift params
func (o *UnregisterSwiftParams) WithTimeout(timeout time.Duration) *UnregisterSwiftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unregister swift params
func (o *UnregisterSwiftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unregister swift params
func (o *UnregisterSwiftParams) WithContext(ctx context.Context) *UnregisterSwiftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unregister swift params
func (o *UnregisterSwiftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unregister swift params
func (o *UnregisterSwiftParams) WithHTTPClient(client *http.Client) *UnregisterSwiftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unregister swift params
func (o *UnregisterSwiftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the unregister swift params
func (o *UnregisterSwiftParams) WithBody(body *models.UnregisterSwiftParams) *UnregisterSwiftParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the unregister swift params
func (o *UnregisterSwiftParams) SetBody(body *models.UnregisterSwiftParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UnregisterSwiftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

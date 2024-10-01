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

// NewUpdateTenantSwiftParams creates a new UpdateTenantSwiftParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTenantSwiftParams() *UpdateTenantSwiftParams {
	return &UpdateTenantSwiftParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantSwiftParamsWithTimeout creates a new UpdateTenantSwiftParams object
// with the ability to set a timeout on a request.
func NewUpdateTenantSwiftParamsWithTimeout(timeout time.Duration) *UpdateTenantSwiftParams {
	return &UpdateTenantSwiftParams{
		timeout: timeout,
	}
}

// NewUpdateTenantSwiftParamsWithContext creates a new UpdateTenantSwiftParams object
// with the ability to set a context for a request.
func NewUpdateTenantSwiftParamsWithContext(ctx context.Context) *UpdateTenantSwiftParams {
	return &UpdateTenantSwiftParams{
		Context: ctx,
	}
}

// NewUpdateTenantSwiftParamsWithHTTPClient creates a new UpdateTenantSwiftParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTenantSwiftParamsWithHTTPClient(client *http.Client) *UpdateTenantSwiftParams {
	return &UpdateTenantSwiftParams{
		HTTPClient: client,
	}
}

/*
UpdateTenantSwiftParams contains all the parameters to send to the API endpoint

	for the update tenant swift operation.

	Typically these are written to a http.Request.
*/
type UpdateTenantSwiftParams struct {

	/* Body.

	   Specifies the parameters to update a Swift configuration.
	*/
	Body *models.SwiftParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tenant swift params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantSwiftParams) WithDefaults() *UpdateTenantSwiftParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tenant swift params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantSwiftParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tenant swift params
func (o *UpdateTenantSwiftParams) WithTimeout(timeout time.Duration) *UpdateTenantSwiftParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant swift params
func (o *UpdateTenantSwiftParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant swift params
func (o *UpdateTenantSwiftParams) WithContext(ctx context.Context) *UpdateTenantSwiftParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant swift params
func (o *UpdateTenantSwiftParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant swift params
func (o *UpdateTenantSwiftParams) WithHTTPClient(client *http.Client) *UpdateTenantSwiftParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant swift params
func (o *UpdateTenantSwiftParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tenant swift params
func (o *UpdateTenantSwiftParams) WithBody(body *models.SwiftParams) *UpdateTenantSwiftParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tenant swift params
func (o *UpdateTenantSwiftParams) SetBody(body *models.SwiftParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantSwiftParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

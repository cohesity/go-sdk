// Code generated by go-swagger; DO NOT EDIT.

package principals

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

// NewUpdateLinuxCredentialsParams creates a new UpdateLinuxCredentialsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateLinuxCredentialsParams() *UpdateLinuxCredentialsParams {
	return &UpdateLinuxCredentialsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLinuxCredentialsParamsWithTimeout creates a new UpdateLinuxCredentialsParams object
// with the ability to set a timeout on a request.
func NewUpdateLinuxCredentialsParamsWithTimeout(timeout time.Duration) *UpdateLinuxCredentialsParams {
	return &UpdateLinuxCredentialsParams{
		timeout: timeout,
	}
}

// NewUpdateLinuxCredentialsParamsWithContext creates a new UpdateLinuxCredentialsParams object
// with the ability to set a context for a request.
func NewUpdateLinuxCredentialsParamsWithContext(ctx context.Context) *UpdateLinuxCredentialsParams {
	return &UpdateLinuxCredentialsParams{
		Context: ctx,
	}
}

// NewUpdateLinuxCredentialsParamsWithHTTPClient creates a new UpdateLinuxCredentialsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateLinuxCredentialsParamsWithHTTPClient(client *http.Client) *UpdateLinuxCredentialsParams {
	return &UpdateLinuxCredentialsParams{
		HTTPClient: client,
	}
}

/*
UpdateLinuxCredentialsParams contains all the parameters to send to the API endpoint

	for the update linux credentials operation.

	Typically these are written to a http.Request.
*/
type UpdateLinuxCredentialsParams struct {

	// Body.
	Body *models.UpdateLinuxPasswordReqParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update linux credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLinuxCredentialsParams) WithDefaults() *UpdateLinuxCredentialsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update linux credentials params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateLinuxCredentialsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) WithTimeout(timeout time.Duration) *UpdateLinuxCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) WithContext(ctx context.Context) *UpdateLinuxCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) WithHTTPClient(client *http.Client) *UpdateLinuxCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) WithBody(body *models.UpdateLinuxPasswordReqParams) *UpdateLinuxCredentialsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update linux credentials params
func (o *UpdateLinuxCredentialsParams) SetBody(body *models.UpdateLinuxPasswordReqParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLinuxCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

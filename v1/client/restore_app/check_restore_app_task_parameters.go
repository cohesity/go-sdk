// Code generated by go-swagger; DO NOT EDIT.

package restore_app

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

// NewCheckRestoreAppTaskParams creates a new CheckRestoreAppTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCheckRestoreAppTaskParams() *CheckRestoreAppTaskParams {
	return &CheckRestoreAppTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCheckRestoreAppTaskParamsWithTimeout creates a new CheckRestoreAppTaskParams object
// with the ability to set a timeout on a request.
func NewCheckRestoreAppTaskParamsWithTimeout(timeout time.Duration) *CheckRestoreAppTaskParams {
	return &CheckRestoreAppTaskParams{
		timeout: timeout,
	}
}

// NewCheckRestoreAppTaskParamsWithContext creates a new CheckRestoreAppTaskParams object
// with the ability to set a context for a request.
func NewCheckRestoreAppTaskParamsWithContext(ctx context.Context) *CheckRestoreAppTaskParams {
	return &CheckRestoreAppTaskParams{
		Context: ctx,
	}
}

// NewCheckRestoreAppTaskParamsWithHTTPClient creates a new CheckRestoreAppTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewCheckRestoreAppTaskParamsWithHTTPClient(client *http.Client) *CheckRestoreAppTaskParams {
	return &CheckRestoreAppTaskParams{
		HTTPClient: client,
	}
}

/*
CheckRestoreAppTaskParams contains all the parameters to send to the API endpoint

	for the check restore app task operation.

	Typically these are written to a http.Request.
*/
type CheckRestoreAppTaskParams struct {

	/* Body.

	   Check Restore Application Task argument
	*/
	Body *models.CheckRestoreAppTaskArg

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the check restore app task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckRestoreAppTaskParams) WithDefaults() *CheckRestoreAppTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the check restore app task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CheckRestoreAppTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the check restore app task params
func (o *CheckRestoreAppTaskParams) WithTimeout(timeout time.Duration) *CheckRestoreAppTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check restore app task params
func (o *CheckRestoreAppTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check restore app task params
func (o *CheckRestoreAppTaskParams) WithContext(ctx context.Context) *CheckRestoreAppTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check restore app task params
func (o *CheckRestoreAppTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check restore app task params
func (o *CheckRestoreAppTaskParams) WithHTTPClient(client *http.Client) *CheckRestoreAppTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check restore app task params
func (o *CheckRestoreAppTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the check restore app task params
func (o *CheckRestoreAppTaskParams) WithBody(body *models.CheckRestoreAppTaskArg) *CheckRestoreAppTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check restore app task params
func (o *CheckRestoreAppTaskParams) SetBody(body *models.CheckRestoreAppTaskArg) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CheckRestoreAppTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

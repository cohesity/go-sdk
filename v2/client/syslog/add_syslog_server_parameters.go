// Code generated by go-swagger; DO NOT EDIT.

package syslog

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

// NewAddSyslogServerParams creates a new AddSyslogServerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddSyslogServerParams() *AddSyslogServerParams {
	return &AddSyslogServerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddSyslogServerParamsWithTimeout creates a new AddSyslogServerParams object
// with the ability to set a timeout on a request.
func NewAddSyslogServerParamsWithTimeout(timeout time.Duration) *AddSyslogServerParams {
	return &AddSyslogServerParams{
		timeout: timeout,
	}
}

// NewAddSyslogServerParamsWithContext creates a new AddSyslogServerParams object
// with the ability to set a context for a request.
func NewAddSyslogServerParamsWithContext(ctx context.Context) *AddSyslogServerParams {
	return &AddSyslogServerParams{
		Context: ctx,
	}
}

// NewAddSyslogServerParamsWithHTTPClient creates a new AddSyslogServerParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddSyslogServerParamsWithHTTPClient(client *http.Client) *AddSyslogServerParams {
	return &AddSyslogServerParams{
		HTTPClient: client,
	}
}

/*
AddSyslogServerParams contains all the parameters to send to the API endpoint

	for the add syslog server operation.

	Typically these are written to a http.Request.
*/
type AddSyslogServerParams struct {

	/* Body.

	   Specifies parameters to add syslog server.
	*/
	Body *models.SyslogServer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add syslog server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSyslogServerParams) WithDefaults() *AddSyslogServerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add syslog server params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddSyslogServerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add syslog server params
func (o *AddSyslogServerParams) WithTimeout(timeout time.Duration) *AddSyslogServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add syslog server params
func (o *AddSyslogServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add syslog server params
func (o *AddSyslogServerParams) WithContext(ctx context.Context) *AddSyslogServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add syslog server params
func (o *AddSyslogServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add syslog server params
func (o *AddSyslogServerParams) WithHTTPClient(client *http.Client) *AddSyslogServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add syslog server params
func (o *AddSyslogServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add syslog server params
func (o *AddSyslogServerParams) WithBody(body *models.SyslogServer) *AddSyslogServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add syslog server params
func (o *AddSyslogServerParams) SetBody(body *models.SyslogServer) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddSyslogServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

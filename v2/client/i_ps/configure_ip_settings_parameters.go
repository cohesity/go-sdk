// Code generated by go-swagger; DO NOT EDIT.

package i_ps

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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewConfigureIPSettingsParams creates a new ConfigureIPSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConfigureIPSettingsParams() *ConfigureIPSettingsParams {
	return &ConfigureIPSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConfigureIPSettingsParamsWithTimeout creates a new ConfigureIPSettingsParams object
// with the ability to set a timeout on a request.
func NewConfigureIPSettingsParamsWithTimeout(timeout time.Duration) *ConfigureIPSettingsParams {
	return &ConfigureIPSettingsParams{
		timeout: timeout,
	}
}

// NewConfigureIPSettingsParamsWithContext creates a new ConfigureIPSettingsParams object
// with the ability to set a context for a request.
func NewConfigureIPSettingsParamsWithContext(ctx context.Context) *ConfigureIPSettingsParams {
	return &ConfigureIPSettingsParams{
		Context: ctx,
	}
}

// NewConfigureIPSettingsParamsWithHTTPClient creates a new ConfigureIPSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewConfigureIPSettingsParamsWithHTTPClient(client *http.Client) *ConfigureIPSettingsParams {
	return &ConfigureIPSettingsParams{
		HTTPClient: client,
	}
}

/*
ConfigureIPSettingsParams contains all the parameters to send to the API endpoint

	for the configure IP settings operation.

	Typically these are written to a http.Request.
*/
type ConfigureIPSettingsParams struct {

	/* Body.

	   Specifies the parameters to configure a IP settings on a cluster.
	*/
	Body *models.IPConfigParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the configure IP settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigureIPSettingsParams) WithDefaults() *ConfigureIPSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the configure IP settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConfigureIPSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the configure IP settings params
func (o *ConfigureIPSettingsParams) WithTimeout(timeout time.Duration) *ConfigureIPSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the configure IP settings params
func (o *ConfigureIPSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the configure IP settings params
func (o *ConfigureIPSettingsParams) WithContext(ctx context.Context) *ConfigureIPSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the configure IP settings params
func (o *ConfigureIPSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the configure IP settings params
func (o *ConfigureIPSettingsParams) WithHTTPClient(client *http.Client) *ConfigureIPSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the configure IP settings params
func (o *ConfigureIPSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the configure IP settings params
func (o *ConfigureIPSettingsParams) WithBody(body *models.IPConfigParams) *ConfigureIPSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the configure IP settings params
func (o *ConfigureIPSettingsParams) SetBody(body *models.IPConfigParams) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ConfigureIPSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

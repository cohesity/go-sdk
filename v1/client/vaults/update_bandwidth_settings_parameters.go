// Code generated by go-swagger; DO NOT EDIT.

package vaults

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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewUpdateBandwidthSettingsParams creates a new UpdateBandwidthSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateBandwidthSettingsParams() *UpdateBandwidthSettingsParams {
	return &UpdateBandwidthSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBandwidthSettingsParamsWithTimeout creates a new UpdateBandwidthSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateBandwidthSettingsParamsWithTimeout(timeout time.Duration) *UpdateBandwidthSettingsParams {
	return &UpdateBandwidthSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateBandwidthSettingsParamsWithContext creates a new UpdateBandwidthSettingsParams object
// with the ability to set a context for a request.
func NewUpdateBandwidthSettingsParamsWithContext(ctx context.Context) *UpdateBandwidthSettingsParams {
	return &UpdateBandwidthSettingsParams{
		Context: ctx,
	}
}

// NewUpdateBandwidthSettingsParamsWithHTTPClient creates a new UpdateBandwidthSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateBandwidthSettingsParamsWithHTTPClient(client *http.Client) *UpdateBandwidthSettingsParams {
	return &UpdateBandwidthSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateBandwidthSettingsParams contains all the parameters to send to the API endpoint

	for the update bandwidth settings operation.

	Typically these are written to a http.Request.
*/
type UpdateBandwidthSettingsParams struct {

	/* Body.

	   Request to update global bandwidth limits settings.
	*/
	Body *models.VaultBandwidthLimits

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update bandwidth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBandwidthSettingsParams) WithDefaults() *UpdateBandwidthSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update bandwidth settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateBandwidthSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) WithTimeout(timeout time.Duration) *UpdateBandwidthSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) WithContext(ctx context.Context) *UpdateBandwidthSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) WithHTTPClient(client *http.Client) *UpdateBandwidthSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) WithBody(body *models.VaultBandwidthLimits) *UpdateBandwidthSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update bandwidth settings params
func (o *UpdateBandwidthSettingsParams) SetBody(body *models.VaultBandwidthLimits) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBandwidthSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

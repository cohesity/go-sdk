// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// NewUpdateExternalTargetSettingsParams creates a new UpdateExternalTargetSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateExternalTargetSettingsParams() *UpdateExternalTargetSettingsParams {
	return &UpdateExternalTargetSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateExternalTargetSettingsParamsWithTimeout creates a new UpdateExternalTargetSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateExternalTargetSettingsParamsWithTimeout(timeout time.Duration) *UpdateExternalTargetSettingsParams {
	return &UpdateExternalTargetSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateExternalTargetSettingsParamsWithContext creates a new UpdateExternalTargetSettingsParams object
// with the ability to set a context for a request.
func NewUpdateExternalTargetSettingsParamsWithContext(ctx context.Context) *UpdateExternalTargetSettingsParams {
	return &UpdateExternalTargetSettingsParams{
		Context: ctx,
	}
}

// NewUpdateExternalTargetSettingsParamsWithHTTPClient creates a new UpdateExternalTargetSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateExternalTargetSettingsParamsWithHTTPClient(client *http.Client) *UpdateExternalTargetSettingsParams {
	return &UpdateExternalTargetSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateExternalTargetSettingsParams contains all the parameters to send to the API endpoint

	for the update external target settings operation.

	Typically these are written to a http.Request.
*/
type UpdateExternalTargetSettingsParams struct {

	/* Body.

	   Specifies the parameters to update a External Target Settings.
	*/
	Body *models.GlobalBandwidthSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update external target settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateExternalTargetSettingsParams) WithDefaults() *UpdateExternalTargetSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update external target settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateExternalTargetSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) WithTimeout(timeout time.Duration) *UpdateExternalTargetSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) WithContext(ctx context.Context) *UpdateExternalTargetSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) WithHTTPClient(client *http.Client) *UpdateExternalTargetSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) WithBody(body *models.GlobalBandwidthSettings) *UpdateExternalTargetSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update external target settings params
func (o *UpdateExternalTargetSettingsParams) SetBody(body *models.GlobalBandwidthSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateExternalTargetSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package app_instance

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
	"github.com/go-openapi/swag"

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// NewUpdateAppInstanceSettingsParams creates a new UpdateAppInstanceSettingsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAppInstanceSettingsParams() *UpdateAppInstanceSettingsParams {
	return &UpdateAppInstanceSettingsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppInstanceSettingsParamsWithTimeout creates a new UpdateAppInstanceSettingsParams object
// with the ability to set a timeout on a request.
func NewUpdateAppInstanceSettingsParamsWithTimeout(timeout time.Duration) *UpdateAppInstanceSettingsParams {
	return &UpdateAppInstanceSettingsParams{
		timeout: timeout,
	}
}

// NewUpdateAppInstanceSettingsParamsWithContext creates a new UpdateAppInstanceSettingsParams object
// with the ability to set a context for a request.
func NewUpdateAppInstanceSettingsParamsWithContext(ctx context.Context) *UpdateAppInstanceSettingsParams {
	return &UpdateAppInstanceSettingsParams{
		Context: ctx,
	}
}

// NewUpdateAppInstanceSettingsParamsWithHTTPClient creates a new UpdateAppInstanceSettingsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAppInstanceSettingsParamsWithHTTPClient(client *http.Client) *UpdateAppInstanceSettingsParams {
	return &UpdateAppInstanceSettingsParams{
		HTTPClient: client,
	}
}

/*
UpdateAppInstanceSettingsParams contains all the parameters to send to the API endpoint

	for the update app instance settings operation.

	Typically these are written to a http.Request.
*/
type UpdateAppInstanceSettingsParams struct {

	/* AppInstanceID.

	   Specifies the app instance Id.

	   Format: int64
	*/
	AppInstanceID int64

	/* Body.

	   Request to update app instance settings.
	*/
	Body *models.UpdateAppInstanceSettingsParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update app instance settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppInstanceSettingsParams) WithDefaults() *UpdateAppInstanceSettingsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update app instance settings params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppInstanceSettingsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) WithTimeout(timeout time.Duration) *UpdateAppInstanceSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) WithContext(ctx context.Context) *UpdateAppInstanceSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) WithHTTPClient(client *http.Client) *UpdateAppInstanceSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppInstanceID adds the appInstanceID to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) WithAppInstanceID(appInstanceID int64) *UpdateAppInstanceSettingsParams {
	o.SetAppInstanceID(appInstanceID)
	return o
}

// SetAppInstanceID adds the appInstanceId to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) SetAppInstanceID(appInstanceID int64) {
	o.AppInstanceID = appInstanceID
}

// WithBody adds the body to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) WithBody(body *models.UpdateAppInstanceSettingsParameters) *UpdateAppInstanceSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update app instance settings params
func (o *UpdateAppInstanceSettingsParams) SetBody(body *models.UpdateAppInstanceSettingsParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppInstanceSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param appInstanceId
	if err := r.SetPathParam("appInstanceId", swag.FormatInt64(o.AppInstanceID)); err != nil {
		return err
	}
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

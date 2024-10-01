// Code generated by go-swagger; DO NOT EDIT.

package m_f_a

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

// NewUpdateSupportMFAConfigParams creates a new UpdateSupportMFAConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSupportMFAConfigParams() *UpdateSupportMFAConfigParams {
	return &UpdateSupportMFAConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSupportMFAConfigParamsWithTimeout creates a new UpdateSupportMFAConfigParams object
// with the ability to set a timeout on a request.
func NewUpdateSupportMFAConfigParamsWithTimeout(timeout time.Duration) *UpdateSupportMFAConfigParams {
	return &UpdateSupportMFAConfigParams{
		timeout: timeout,
	}
}

// NewUpdateSupportMFAConfigParamsWithContext creates a new UpdateSupportMFAConfigParams object
// with the ability to set a context for a request.
func NewUpdateSupportMFAConfigParamsWithContext(ctx context.Context) *UpdateSupportMFAConfigParams {
	return &UpdateSupportMFAConfigParams{
		Context: ctx,
	}
}

// NewUpdateSupportMFAConfigParamsWithHTTPClient creates a new UpdateSupportMFAConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSupportMFAConfigParamsWithHTTPClient(client *http.Client) *UpdateSupportMFAConfigParams {
	return &UpdateSupportMFAConfigParams{
		HTTPClient: client,
	}
}

/*
UpdateSupportMFAConfigParams contains all the parameters to send to the API endpoint

	for the update support m f a config operation.

	Typically these are written to a http.Request.
*/
type UpdateSupportMFAConfigParams struct {

	/* Body.

	   The update request for the MFA Settings
	*/
	Body *models.SupportMfaConfigInfo

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update support m f a config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSupportMFAConfigParams) WithDefaults() *UpdateSupportMFAConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update support m f a config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSupportMFAConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) WithTimeout(timeout time.Duration) *UpdateSupportMFAConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) WithContext(ctx context.Context) *UpdateSupportMFAConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) WithHTTPClient(client *http.Client) *UpdateSupportMFAConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) WithBody(body *models.SupportMfaConfigInfo) *UpdateSupportMFAConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update support m f a config params
func (o *UpdateSupportMFAConfigParams) SetBody(body *models.SupportMfaConfigInfo) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSupportMFAConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package views

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

// NewUpdateViewUserQuotaParams creates a new UpdateViewUserQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateViewUserQuotaParams() *UpdateViewUserQuotaParams {
	return &UpdateViewUserQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateViewUserQuotaParamsWithTimeout creates a new UpdateViewUserQuotaParams object
// with the ability to set a timeout on a request.
func NewUpdateViewUserQuotaParamsWithTimeout(timeout time.Duration) *UpdateViewUserQuotaParams {
	return &UpdateViewUserQuotaParams{
		timeout: timeout,
	}
}

// NewUpdateViewUserQuotaParamsWithContext creates a new UpdateViewUserQuotaParams object
// with the ability to set a context for a request.
func NewUpdateViewUserQuotaParamsWithContext(ctx context.Context) *UpdateViewUserQuotaParams {
	return &UpdateViewUserQuotaParams{
		Context: ctx,
	}
}

// NewUpdateViewUserQuotaParamsWithHTTPClient creates a new UpdateViewUserQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateViewUserQuotaParamsWithHTTPClient(client *http.Client) *UpdateViewUserQuotaParams {
	return &UpdateViewUserQuotaParams{
		HTTPClient: client,
	}
}

/*
UpdateViewUserQuotaParams contains all the parameters to send to the API endpoint

	for the update view user quota operation.

	Typically these are written to a http.Request.
*/
type UpdateViewUserQuotaParams struct {

	/* Body.

	   update user quota params.
	*/
	Body *models.ViewUserQuotaParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update view user quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewUserQuotaParams) WithDefaults() *UpdateViewUserQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update view user quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewUserQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update view user quota params
func (o *UpdateViewUserQuotaParams) WithTimeout(timeout time.Duration) *UpdateViewUserQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update view user quota params
func (o *UpdateViewUserQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update view user quota params
func (o *UpdateViewUserQuotaParams) WithContext(ctx context.Context) *UpdateViewUserQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update view user quota params
func (o *UpdateViewUserQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update view user quota params
func (o *UpdateViewUserQuotaParams) WithHTTPClient(client *http.Client) *UpdateViewUserQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update view user quota params
func (o *UpdateViewUserQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update view user quota params
func (o *UpdateViewUserQuotaParams) WithBody(body *models.ViewUserQuotaParameters) *UpdateViewUserQuotaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update view user quota params
func (o *UpdateViewUserQuotaParams) SetBody(body *models.ViewUserQuotaParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateViewUserQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

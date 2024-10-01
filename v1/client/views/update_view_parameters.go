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

// NewUpdateViewParams creates a new UpdateViewParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateViewParams() *UpdateViewParams {
	return &UpdateViewParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateViewParamsWithTimeout creates a new UpdateViewParams object
// with the ability to set a timeout on a request.
func NewUpdateViewParamsWithTimeout(timeout time.Duration) *UpdateViewParams {
	return &UpdateViewParams{
		timeout: timeout,
	}
}

// NewUpdateViewParamsWithContext creates a new UpdateViewParams object
// with the ability to set a context for a request.
func NewUpdateViewParamsWithContext(ctx context.Context) *UpdateViewParams {
	return &UpdateViewParams{
		Context: ctx,
	}
}

// NewUpdateViewParamsWithHTTPClient creates a new UpdateViewParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateViewParamsWithHTTPClient(client *http.Client) *UpdateViewParams {
	return &UpdateViewParams{
		HTTPClient: client,
	}
}

/*
UpdateViewParams contains all the parameters to send to the API endpoint

	for the update view operation.

	Typically these are written to a http.Request.
*/
type UpdateViewParams struct {

	/* Body.

	   Request to update a view.
	*/
	Body *models.UpdateViewParam

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewParams) WithDefaults() *UpdateViewParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update view params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateViewParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update view params
func (o *UpdateViewParams) WithTimeout(timeout time.Duration) *UpdateViewParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update view params
func (o *UpdateViewParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update view params
func (o *UpdateViewParams) WithContext(ctx context.Context) *UpdateViewParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update view params
func (o *UpdateViewParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update view params
func (o *UpdateViewParams) WithHTTPClient(client *http.Client) *UpdateViewParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update view params
func (o *UpdateViewParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update view params
func (o *UpdateViewParams) WithBody(body *models.UpdateViewParam) *UpdateViewParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update view params
func (o *UpdateViewParams) SetBody(body *models.UpdateViewParam) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateViewParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

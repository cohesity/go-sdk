// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewUpdateTenantEntityParams creates a new UpdateTenantEntityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTenantEntityParams() *UpdateTenantEntityParams {
	return &UpdateTenantEntityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantEntityParamsWithTimeout creates a new UpdateTenantEntityParams object
// with the ability to set a timeout on a request.
func NewUpdateTenantEntityParamsWithTimeout(timeout time.Duration) *UpdateTenantEntityParams {
	return &UpdateTenantEntityParams{
		timeout: timeout,
	}
}

// NewUpdateTenantEntityParamsWithContext creates a new UpdateTenantEntityParams object
// with the ability to set a context for a request.
func NewUpdateTenantEntityParamsWithContext(ctx context.Context) *UpdateTenantEntityParams {
	return &UpdateTenantEntityParams{
		Context: ctx,
	}
}

// NewUpdateTenantEntityParamsWithHTTPClient creates a new UpdateTenantEntityParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTenantEntityParamsWithHTTPClient(client *http.Client) *UpdateTenantEntityParams {
	return &UpdateTenantEntityParams{
		HTTPClient: client,
	}
}

/*
UpdateTenantEntityParams contains all the parameters to send to the API endpoint

	for the update tenant entity operation.

	Typically these are written to a http.Request.
*/
type UpdateTenantEntityParams struct {

	/* Body.

	   Request to associate entity for tenant.
	*/
	Body *models.TenantEntityUpdateParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tenant entity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantEntityParams) WithDefaults() *UpdateTenantEntityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tenant entity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantEntityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tenant entity params
func (o *UpdateTenantEntityParams) WithTimeout(timeout time.Duration) *UpdateTenantEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant entity params
func (o *UpdateTenantEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant entity params
func (o *UpdateTenantEntityParams) WithContext(ctx context.Context) *UpdateTenantEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant entity params
func (o *UpdateTenantEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant entity params
func (o *UpdateTenantEntityParams) WithHTTPClient(client *http.Client) *UpdateTenantEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant entity params
func (o *UpdateTenantEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tenant entity params
func (o *UpdateTenantEntityParams) WithBody(body *models.TenantEntityUpdateParameters) *UpdateTenantEntityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tenant entity params
func (o *UpdateTenantEntityParams) SetBody(body *models.TenantEntityUpdateParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

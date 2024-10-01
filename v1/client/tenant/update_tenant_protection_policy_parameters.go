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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewUpdateTenantProtectionPolicyParams creates a new UpdateTenantProtectionPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTenantProtectionPolicyParams() *UpdateTenantProtectionPolicyParams {
	return &UpdateTenantProtectionPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantProtectionPolicyParamsWithTimeout creates a new UpdateTenantProtectionPolicyParams object
// with the ability to set a timeout on a request.
func NewUpdateTenantProtectionPolicyParamsWithTimeout(timeout time.Duration) *UpdateTenantProtectionPolicyParams {
	return &UpdateTenantProtectionPolicyParams{
		timeout: timeout,
	}
}

// NewUpdateTenantProtectionPolicyParamsWithContext creates a new UpdateTenantProtectionPolicyParams object
// with the ability to set a context for a request.
func NewUpdateTenantProtectionPolicyParamsWithContext(ctx context.Context) *UpdateTenantProtectionPolicyParams {
	return &UpdateTenantProtectionPolicyParams{
		Context: ctx,
	}
}

// NewUpdateTenantProtectionPolicyParamsWithHTTPClient creates a new UpdateTenantProtectionPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTenantProtectionPolicyParamsWithHTTPClient(client *http.Client) *UpdateTenantProtectionPolicyParams {
	return &UpdateTenantProtectionPolicyParams{
		HTTPClient: client,
	}
}

/*
UpdateTenantProtectionPolicyParams contains all the parameters to send to the API endpoint

	for the update tenant protection policy operation.

	Typically these are written to a http.Request.
*/
type UpdateTenantProtectionPolicyParams struct {

	/* Body.

	   Request to update existing protection policies.
	*/
	Body *models.TenantProtectionPolicyUpdateParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tenant protection policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantProtectionPolicyParams) WithDefaults() *UpdateTenantProtectionPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tenant protection policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantProtectionPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) WithTimeout(timeout time.Duration) *UpdateTenantProtectionPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) WithContext(ctx context.Context) *UpdateTenantProtectionPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) WithHTTPClient(client *http.Client) *UpdateTenantProtectionPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) WithBody(body *models.TenantProtectionPolicyUpdateParameters) *UpdateTenantProtectionPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tenant protection policy params
func (o *UpdateTenantProtectionPolicyParams) SetBody(body *models.TenantProtectionPolicyUpdateParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantProtectionPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

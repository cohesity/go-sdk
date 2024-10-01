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

// NewUpdateTenantLdapProviderParams creates a new UpdateTenantLdapProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateTenantLdapProviderParams() *UpdateTenantLdapProviderParams {
	return &UpdateTenantLdapProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateTenantLdapProviderParamsWithTimeout creates a new UpdateTenantLdapProviderParams object
// with the ability to set a timeout on a request.
func NewUpdateTenantLdapProviderParamsWithTimeout(timeout time.Duration) *UpdateTenantLdapProviderParams {
	return &UpdateTenantLdapProviderParams{
		timeout: timeout,
	}
}

// NewUpdateTenantLdapProviderParamsWithContext creates a new UpdateTenantLdapProviderParams object
// with the ability to set a context for a request.
func NewUpdateTenantLdapProviderParamsWithContext(ctx context.Context) *UpdateTenantLdapProviderParams {
	return &UpdateTenantLdapProviderParams{
		Context: ctx,
	}
}

// NewUpdateTenantLdapProviderParamsWithHTTPClient creates a new UpdateTenantLdapProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateTenantLdapProviderParamsWithHTTPClient(client *http.Client) *UpdateTenantLdapProviderParams {
	return &UpdateTenantLdapProviderParams{
		HTTPClient: client,
	}
}

/*
UpdateTenantLdapProviderParams contains all the parameters to send to the API endpoint

	for the update tenant ldap provider operation.

	Typically these are written to a http.Request.
*/
type UpdateTenantLdapProviderParams struct {

	/* Body.

	   Request to update existing ldap providers.
	*/
	Body *models.TenantLdapProviderUpdateParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update tenant ldap provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantLdapProviderParams) WithDefaults() *UpdateTenantLdapProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update tenant ldap provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateTenantLdapProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) WithTimeout(timeout time.Duration) *UpdateTenantLdapProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) WithContext(ctx context.Context) *UpdateTenantLdapProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) WithHTTPClient(client *http.Client) *UpdateTenantLdapProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) WithBody(body *models.TenantLdapProviderUpdateParameters) *UpdateTenantLdapProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update tenant ldap provider params
func (o *UpdateTenantLdapProviderParams) SetBody(body *models.TenantLdapProviderUpdateParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateTenantLdapProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

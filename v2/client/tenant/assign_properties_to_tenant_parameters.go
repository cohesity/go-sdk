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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// NewAssignPropertiesToTenantParams creates a new AssignPropertiesToTenantParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssignPropertiesToTenantParams() *AssignPropertiesToTenantParams {
	return &AssignPropertiesToTenantParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssignPropertiesToTenantParamsWithTimeout creates a new AssignPropertiesToTenantParams object
// with the ability to set a timeout on a request.
func NewAssignPropertiesToTenantParamsWithTimeout(timeout time.Duration) *AssignPropertiesToTenantParams {
	return &AssignPropertiesToTenantParams{
		timeout: timeout,
	}
}

// NewAssignPropertiesToTenantParamsWithContext creates a new AssignPropertiesToTenantParams object
// with the ability to set a context for a request.
func NewAssignPropertiesToTenantParamsWithContext(ctx context.Context) *AssignPropertiesToTenantParams {
	return &AssignPropertiesToTenantParams{
		Context: ctx,
	}
}

// NewAssignPropertiesToTenantParamsWithHTTPClient creates a new AssignPropertiesToTenantParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssignPropertiesToTenantParamsWithHTTPClient(client *http.Client) *AssignPropertiesToTenantParams {
	return &AssignPropertiesToTenantParams{
		HTTPClient: client,
	}
}

/*
AssignPropertiesToTenantParams contains all the parameters to send to the API endpoint

	for the assign properties to tenant operation.

	Typically these are written to a http.Request.
*/
type AssignPropertiesToTenantParams struct {

	// Body.
	Body *models.TenantAssignmentsParams

	/* ID.

	   The Tenant id.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the assign properties to tenant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignPropertiesToTenantParams) WithDefaults() *AssignPropertiesToTenantParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the assign properties to tenant params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssignPropertiesToTenantParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) WithTimeout(timeout time.Duration) *AssignPropertiesToTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) WithContext(ctx context.Context) *AssignPropertiesToTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) WithHTTPClient(client *http.Client) *AssignPropertiesToTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) WithBody(body *models.TenantAssignmentsParams) *AssignPropertiesToTenantParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) SetBody(body *models.TenantAssignmentsParams) {
	o.Body = body
}

// WithID adds the id to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) WithID(id string) *AssignPropertiesToTenantParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the assign properties to tenant params
func (o *AssignPropertiesToTenantParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AssignPropertiesToTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

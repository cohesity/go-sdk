// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// NewUpdatePreferredDomainControllersParams creates a new UpdatePreferredDomainControllersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdatePreferredDomainControllersParams() *UpdatePreferredDomainControllersParams {
	return &UpdatePreferredDomainControllersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePreferredDomainControllersParamsWithTimeout creates a new UpdatePreferredDomainControllersParams object
// with the ability to set a timeout on a request.
func NewUpdatePreferredDomainControllersParamsWithTimeout(timeout time.Duration) *UpdatePreferredDomainControllersParams {
	return &UpdatePreferredDomainControllersParams{
		timeout: timeout,
	}
}

// NewUpdatePreferredDomainControllersParamsWithContext creates a new UpdatePreferredDomainControllersParams object
// with the ability to set a context for a request.
func NewUpdatePreferredDomainControllersParamsWithContext(ctx context.Context) *UpdatePreferredDomainControllersParams {
	return &UpdatePreferredDomainControllersParams{
		Context: ctx,
	}
}

// NewUpdatePreferredDomainControllersParamsWithHTTPClient creates a new UpdatePreferredDomainControllersParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdatePreferredDomainControllersParamsWithHTTPClient(client *http.Client) *UpdatePreferredDomainControllersParams {
	return &UpdatePreferredDomainControllersParams{
		HTTPClient: client,
	}
}

/*
UpdatePreferredDomainControllersParams contains all the parameters to send to the API endpoint

	for the update preferred domain controllers operation.

	Typically these are written to a http.Request.
*/
type UpdatePreferredDomainControllersParams struct {

	/* Body.

	   Request to update preferred domain controllers of an Active Directory.
	*/
	Body []*models.PreferredDomainController

	/* Name.

	   Specifies the Active Directory Domain Name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update preferred domain controllers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePreferredDomainControllersParams) WithDefaults() *UpdatePreferredDomainControllersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update preferred domain controllers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdatePreferredDomainControllersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) WithTimeout(timeout time.Duration) *UpdatePreferredDomainControllersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) WithContext(ctx context.Context) *UpdatePreferredDomainControllersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) WithHTTPClient(client *http.Client) *UpdatePreferredDomainControllersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) WithBody(body []*models.PreferredDomainController) *UpdatePreferredDomainControllersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) SetBody(body []*models.PreferredDomainController) {
	o.Body = body
}

// WithName adds the name to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) WithName(name string) *UpdatePreferredDomainControllersParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update preferred domain controllers params
func (o *UpdatePreferredDomainControllersParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePreferredDomainControllersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

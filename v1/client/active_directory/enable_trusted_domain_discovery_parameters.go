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

// NewEnableTrustedDomainDiscoveryParams creates a new EnableTrustedDomainDiscoveryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnableTrustedDomainDiscoveryParams() *EnableTrustedDomainDiscoveryParams {
	return &EnableTrustedDomainDiscoveryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnableTrustedDomainDiscoveryParamsWithTimeout creates a new EnableTrustedDomainDiscoveryParams object
// with the ability to set a timeout on a request.
func NewEnableTrustedDomainDiscoveryParamsWithTimeout(timeout time.Duration) *EnableTrustedDomainDiscoveryParams {
	return &EnableTrustedDomainDiscoveryParams{
		timeout: timeout,
	}
}

// NewEnableTrustedDomainDiscoveryParamsWithContext creates a new EnableTrustedDomainDiscoveryParams object
// with the ability to set a context for a request.
func NewEnableTrustedDomainDiscoveryParamsWithContext(ctx context.Context) *EnableTrustedDomainDiscoveryParams {
	return &EnableTrustedDomainDiscoveryParams{
		Context: ctx,
	}
}

// NewEnableTrustedDomainDiscoveryParamsWithHTTPClient creates a new EnableTrustedDomainDiscoveryParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnableTrustedDomainDiscoveryParamsWithHTTPClient(client *http.Client) *EnableTrustedDomainDiscoveryParams {
	return &EnableTrustedDomainDiscoveryParams{
		HTTPClient: client,
	}
}

/*
EnableTrustedDomainDiscoveryParams contains all the parameters to send to the API endpoint

	for the enable trusted domain discovery operation.

	Typically these are written to a http.Request.
*/
type EnableTrustedDomainDiscoveryParams struct {

	/* Body.

	   Request to update enable trusted domains state of an Active Directory.
	*/
	Body *models.UpdateTrustedDomainEnableParams

	/* Name.

	   Specifies the Active Directory Domain Name.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enable trusted domain discovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableTrustedDomainDiscoveryParams) WithDefaults() *EnableTrustedDomainDiscoveryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enable trusted domain discovery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnableTrustedDomainDiscoveryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) WithTimeout(timeout time.Duration) *EnableTrustedDomainDiscoveryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) WithContext(ctx context.Context) *EnableTrustedDomainDiscoveryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) WithHTTPClient(client *http.Client) *EnableTrustedDomainDiscoveryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) WithBody(body *models.UpdateTrustedDomainEnableParams) *EnableTrustedDomainDiscoveryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) SetBody(body *models.UpdateTrustedDomainEnableParams) {
	o.Body = body
}

// WithName adds the name to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) WithName(name string) *EnableTrustedDomainDiscoveryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the enable trusted domain discovery params
func (o *EnableTrustedDomainDiscoveryParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *EnableTrustedDomainDiscoveryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

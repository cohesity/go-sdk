// Code generated by go-swagger; DO NOT EDIT.

package remote_clusters

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

// NewValidateRemoteClusterParams creates a new ValidateRemoteClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidateRemoteClusterParams() *ValidateRemoteClusterParams {
	return &ValidateRemoteClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidateRemoteClusterParamsWithTimeout creates a new ValidateRemoteClusterParams object
// with the ability to set a timeout on a request.
func NewValidateRemoteClusterParamsWithTimeout(timeout time.Duration) *ValidateRemoteClusterParams {
	return &ValidateRemoteClusterParams{
		timeout: timeout,
	}
}

// NewValidateRemoteClusterParamsWithContext creates a new ValidateRemoteClusterParams object
// with the ability to set a context for a request.
func NewValidateRemoteClusterParamsWithContext(ctx context.Context) *ValidateRemoteClusterParams {
	return &ValidateRemoteClusterParams{
		Context: ctx,
	}
}

// NewValidateRemoteClusterParamsWithHTTPClient creates a new ValidateRemoteClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidateRemoteClusterParamsWithHTTPClient(client *http.Client) *ValidateRemoteClusterParams {
	return &ValidateRemoteClusterParams{
		HTTPClient: client,
	}
}

/*
ValidateRemoteClusterParams contains all the parameters to send to the API endpoint

	for the validate remote cluster operation.

	Typically these are written to a http.Request.
*/
type ValidateRemoteClusterParams struct {

	/* Body.

	   Specifies the request to validate Remote Cluster.
	*/
	Body *models.ValidateRemoteClusterConnectionParam

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validate remote cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateRemoteClusterParams) WithDefaults() *ValidateRemoteClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validate remote cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidateRemoteClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validate remote cluster params
func (o *ValidateRemoteClusterParams) WithTimeout(timeout time.Duration) *ValidateRemoteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate remote cluster params
func (o *ValidateRemoteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate remote cluster params
func (o *ValidateRemoteClusterParams) WithContext(ctx context.Context) *ValidateRemoteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate remote cluster params
func (o *ValidateRemoteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate remote cluster params
func (o *ValidateRemoteClusterParams) WithHTTPClient(client *http.Client) *ValidateRemoteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate remote cluster params
func (o *ValidateRemoteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the validate remote cluster params
func (o *ValidateRemoteClusterParams) WithBody(body *models.ValidateRemoteClusterConnectionParam) *ValidateRemoteClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the validate remote cluster params
func (o *ValidateRemoteClusterParams) SetBody(body *models.ValidateRemoteClusterConnectionParam) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateRemoteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package remote_cluster

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
	"github.com/go-openapi/swag"

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// NewUpdateRemoteClusterParams creates a new UpdateRemoteClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateRemoteClusterParams() *UpdateRemoteClusterParams {
	return &UpdateRemoteClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRemoteClusterParamsWithTimeout creates a new UpdateRemoteClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateRemoteClusterParamsWithTimeout(timeout time.Duration) *UpdateRemoteClusterParams {
	return &UpdateRemoteClusterParams{
		timeout: timeout,
	}
}

// NewUpdateRemoteClusterParamsWithContext creates a new UpdateRemoteClusterParams object
// with the ability to set a context for a request.
func NewUpdateRemoteClusterParamsWithContext(ctx context.Context) *UpdateRemoteClusterParams {
	return &UpdateRemoteClusterParams{
		Context: ctx,
	}
}

// NewUpdateRemoteClusterParamsWithHTTPClient creates a new UpdateRemoteClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateRemoteClusterParamsWithHTTPClient(client *http.Client) *UpdateRemoteClusterParams {
	return &UpdateRemoteClusterParams{
		HTTPClient: client,
	}
}

/*
UpdateRemoteClusterParams contains all the parameters to send to the API endpoint

	for the update remote cluster operation.

	Typically these are written to a http.Request.
*/
type UpdateRemoteClusterParams struct {

	/* Body.

	   Request to update a remote Cluster.
	*/
	Body *models.RegisterRemoteCluster

	/* ID.

	   id of the remote Cluster

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update remote cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRemoteClusterParams) WithDefaults() *UpdateRemoteClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update remote cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateRemoteClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update remote cluster params
func (o *UpdateRemoteClusterParams) WithTimeout(timeout time.Duration) *UpdateRemoteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update remote cluster params
func (o *UpdateRemoteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update remote cluster params
func (o *UpdateRemoteClusterParams) WithContext(ctx context.Context) *UpdateRemoteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update remote cluster params
func (o *UpdateRemoteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update remote cluster params
func (o *UpdateRemoteClusterParams) WithHTTPClient(client *http.Client) *UpdateRemoteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update remote cluster params
func (o *UpdateRemoteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update remote cluster params
func (o *UpdateRemoteClusterParams) WithBody(body *models.RegisterRemoteCluster) *UpdateRemoteClusterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update remote cluster params
func (o *UpdateRemoteClusterParams) SetBody(body *models.RegisterRemoteCluster) {
	o.Body = body
}

// WithID adds the id to the update remote cluster params
func (o *UpdateRemoteClusterParams) WithID(id int64) *UpdateRemoteClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update remote cluster params
func (o *UpdateRemoteClusterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRemoteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

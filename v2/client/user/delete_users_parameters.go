// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewDeleteUsersParams creates a new DeleteUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsersParams() *DeleteUsersParams {
	return &DeleteUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsersParamsWithTimeout creates a new DeleteUsersParams object
// with the ability to set a timeout on a request.
func NewDeleteUsersParamsWithTimeout(timeout time.Duration) *DeleteUsersParams {
	return &DeleteUsersParams{
		timeout: timeout,
	}
}

// NewDeleteUsersParamsWithContext creates a new DeleteUsersParams object
// with the ability to set a context for a request.
func NewDeleteUsersParamsWithContext(ctx context.Context) *DeleteUsersParams {
	return &DeleteUsersParams{
		Context: ctx,
	}
}

// NewDeleteUsersParamsWithHTTPClient creates a new DeleteUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsersParamsWithHTTPClient(client *http.Client) *DeleteUsersParams {
	return &DeleteUsersParams{
		HTTPClient: client,
	}
}

/*
DeleteUsersParams contains all the parameters to send to the API endpoint

	for the delete users operation.

	Typically these are written to a http.Request.
*/
type DeleteUsersParams struct {

	/* Body.

	   If the Cohesity user was created against an Active Directory/IdP user, the referenced principal user on the Active Directory/IdP domain is NOT deleted. Only the user on the Cohesity Cluster is deleted.
	*/
	Body *models.DeleteUsersRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsersParams) WithDefaults() *DeleteUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete users params
func (o *DeleteUsersParams) WithTimeout(timeout time.Duration) *DeleteUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete users params
func (o *DeleteUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete users params
func (o *DeleteUsersParams) WithContext(ctx context.Context) *DeleteUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete users params
func (o *DeleteUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete users params
func (o *DeleteUsersParams) WithHTTPClient(client *http.Client) *DeleteUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete users params
func (o *DeleteUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete users params
func (o *DeleteUsersParams) WithBody(body *models.DeleteUsersRequest) *DeleteUsersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete users params
func (o *DeleteUsersParams) SetBody(body *models.DeleteUsersRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

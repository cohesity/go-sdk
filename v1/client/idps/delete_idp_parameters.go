// Code generated by go-swagger; DO NOT EDIT.

package idps

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
)

// NewDeleteIdpParams creates a new DeleteIdpParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteIdpParams() *DeleteIdpParams {
	return &DeleteIdpParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteIdpParamsWithTimeout creates a new DeleteIdpParams object
// with the ability to set a timeout on a request.
func NewDeleteIdpParamsWithTimeout(timeout time.Duration) *DeleteIdpParams {
	return &DeleteIdpParams{
		timeout: timeout,
	}
}

// NewDeleteIdpParamsWithContext creates a new DeleteIdpParams object
// with the ability to set a context for a request.
func NewDeleteIdpParamsWithContext(ctx context.Context) *DeleteIdpParams {
	return &DeleteIdpParams{
		Context: ctx,
	}
}

// NewDeleteIdpParamsWithHTTPClient creates a new DeleteIdpParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteIdpParamsWithHTTPClient(client *http.Client) *DeleteIdpParams {
	return &DeleteIdpParams{
		HTTPClient: client,
	}
}

/*
DeleteIdpParams contains all the parameters to send to the API endpoint

	for the delete idp operation.

	Typically these are written to a http.Request.
*/
type DeleteIdpParams struct {

	/* ID.

	   Specifies the Id assigned for the IdP Service by the Cluster.

	   Format: int64
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete idp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdpParams) WithDefaults() *DeleteIdpParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete idp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteIdpParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete idp params
func (o *DeleteIdpParams) WithTimeout(timeout time.Duration) *DeleteIdpParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete idp params
func (o *DeleteIdpParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete idp params
func (o *DeleteIdpParams) WithContext(ctx context.Context) *DeleteIdpParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete idp params
func (o *DeleteIdpParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete idp params
func (o *DeleteIdpParams) WithHTTPClient(client *http.Client) *DeleteIdpParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete idp params
func (o *DeleteIdpParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete idp params
func (o *DeleteIdpParams) WithID(id int64) *DeleteIdpParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete idp params
func (o *DeleteIdpParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteIdpParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

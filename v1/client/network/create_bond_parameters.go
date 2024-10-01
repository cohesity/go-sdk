// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewCreateBondParams creates a new CreateBondParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateBondParams() *CreateBondParams {
	return &CreateBondParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateBondParamsWithTimeout creates a new CreateBondParams object
// with the ability to set a timeout on a request.
func NewCreateBondParamsWithTimeout(timeout time.Duration) *CreateBondParams {
	return &CreateBondParams{
		timeout: timeout,
	}
}

// NewCreateBondParamsWithContext creates a new CreateBondParams object
// with the ability to set a context for a request.
func NewCreateBondParamsWithContext(ctx context.Context) *CreateBondParams {
	return &CreateBondParams{
		Context: ctx,
	}
}

// NewCreateBondParamsWithHTTPClient creates a new CreateBondParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateBondParamsWithHTTPClient(client *http.Client) *CreateBondParams {
	return &CreateBondParams{
		HTTPClient: client,
	}
}

/*
CreateBondParams contains all the parameters to send to the API endpoint

	for the create bond operation.

	Typically these are written to a http.Request.
*/
type CreateBondParams struct {

	// Body.
	Body *models.CreateBondParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create bond params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBondParams) WithDefaults() *CreateBondParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create bond params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateBondParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create bond params
func (o *CreateBondParams) WithTimeout(timeout time.Duration) *CreateBondParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create bond params
func (o *CreateBondParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create bond params
func (o *CreateBondParams) WithContext(ctx context.Context) *CreateBondParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create bond params
func (o *CreateBondParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create bond params
func (o *CreateBondParams) WithHTTPClient(client *http.Client) *CreateBondParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create bond params
func (o *CreateBondParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create bond params
func (o *CreateBondParams) WithBody(body *models.CreateBondParameters) *CreateBondParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create bond params
func (o *CreateBondParams) SetBody(body *models.CreateBondParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateBondParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// NewCreateExternalTargetParams creates a new CreateExternalTargetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateExternalTargetParams() *CreateExternalTargetParams {
	return &CreateExternalTargetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExternalTargetParamsWithTimeout creates a new CreateExternalTargetParams object
// with the ability to set a timeout on a request.
func NewCreateExternalTargetParamsWithTimeout(timeout time.Duration) *CreateExternalTargetParams {
	return &CreateExternalTargetParams{
		timeout: timeout,
	}
}

// NewCreateExternalTargetParamsWithContext creates a new CreateExternalTargetParams object
// with the ability to set a context for a request.
func NewCreateExternalTargetParamsWithContext(ctx context.Context) *CreateExternalTargetParams {
	return &CreateExternalTargetParams{
		Context: ctx,
	}
}

// NewCreateExternalTargetParamsWithHTTPClient creates a new CreateExternalTargetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateExternalTargetParamsWithHTTPClient(client *http.Client) *CreateExternalTargetParams {
	return &CreateExternalTargetParams{
		HTTPClient: client,
	}
}

/*
CreateExternalTargetParams contains all the parameters to send to the API endpoint

	for the create external target operation.

	Typically these are written to a http.Request.
*/
type CreateExternalTargetParams struct {

	/* Body.

	   Specifies the parameters to create a External Target.
	*/
	Body *models.ExternalTarget

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create external target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExternalTargetParams) WithDefaults() *CreateExternalTargetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create external target params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateExternalTargetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create external target params
func (o *CreateExternalTargetParams) WithTimeout(timeout time.Duration) *CreateExternalTargetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create external target params
func (o *CreateExternalTargetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create external target params
func (o *CreateExternalTargetParams) WithContext(ctx context.Context) *CreateExternalTargetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create external target params
func (o *CreateExternalTargetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create external target params
func (o *CreateExternalTargetParams) WithHTTPClient(client *http.Client) *CreateExternalTargetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create external target params
func (o *CreateExternalTargetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create external target params
func (o *CreateExternalTargetParams) WithBody(body *models.ExternalTarget) *CreateExternalTargetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create external target params
func (o *CreateExternalTargetParams) SetBody(body *models.ExternalTarget) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExternalTargetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package patch_management

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

// NewApplyPatchesParams creates a new ApplyPatchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplyPatchesParams() *ApplyPatchesParams {
	return &ApplyPatchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplyPatchesParamsWithTimeout creates a new ApplyPatchesParams object
// with the ability to set a timeout on a request.
func NewApplyPatchesParamsWithTimeout(timeout time.Duration) *ApplyPatchesParams {
	return &ApplyPatchesParams{
		timeout: timeout,
	}
}

// NewApplyPatchesParamsWithContext creates a new ApplyPatchesParams object
// with the ability to set a context for a request.
func NewApplyPatchesParamsWithContext(ctx context.Context) *ApplyPatchesParams {
	return &ApplyPatchesParams{
		Context: ctx,
	}
}

// NewApplyPatchesParamsWithHTTPClient creates a new ApplyPatchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplyPatchesParamsWithHTTPClient(client *http.Client) *ApplyPatchesParams {
	return &ApplyPatchesParams{
		HTTPClient: client,
	}
}

/*
ApplyPatchesParams contains all the parameters to send to the API endpoint

	for the apply patches operation.

	Typically these are written to a http.Request.
*/
type ApplyPatchesParams struct {

	/* Body.

	   Request to apply patches.
	*/
	Body *models.ApplyPatchesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the apply patches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplyPatchesParams) WithDefaults() *ApplyPatchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the apply patches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplyPatchesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the apply patches params
func (o *ApplyPatchesParams) WithTimeout(timeout time.Duration) *ApplyPatchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the apply patches params
func (o *ApplyPatchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the apply patches params
func (o *ApplyPatchesParams) WithContext(ctx context.Context) *ApplyPatchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the apply patches params
func (o *ApplyPatchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the apply patches params
func (o *ApplyPatchesParams) WithHTTPClient(client *http.Client) *ApplyPatchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the apply patches params
func (o *ApplyPatchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the apply patches params
func (o *ApplyPatchesParams) WithBody(body *models.ApplyPatchesRequest) *ApplyPatchesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the apply patches params
func (o *ApplyPatchesParams) SetBody(body *models.ApplyPatchesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ApplyPatchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

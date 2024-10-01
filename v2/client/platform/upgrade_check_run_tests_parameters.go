// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// NewUpgradeCheckRunTestsParams creates a new UpgradeCheckRunTestsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeCheckRunTestsParams() *UpgradeCheckRunTestsParams {
	return &UpgradeCheckRunTestsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeCheckRunTestsParamsWithTimeout creates a new UpgradeCheckRunTestsParams object
// with the ability to set a timeout on a request.
func NewUpgradeCheckRunTestsParamsWithTimeout(timeout time.Duration) *UpgradeCheckRunTestsParams {
	return &UpgradeCheckRunTestsParams{
		timeout: timeout,
	}
}

// NewUpgradeCheckRunTestsParamsWithContext creates a new UpgradeCheckRunTestsParams object
// with the ability to set a context for a request.
func NewUpgradeCheckRunTestsParamsWithContext(ctx context.Context) *UpgradeCheckRunTestsParams {
	return &UpgradeCheckRunTestsParams{
		Context: ctx,
	}
}

// NewUpgradeCheckRunTestsParamsWithHTTPClient creates a new UpgradeCheckRunTestsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeCheckRunTestsParamsWithHTTPClient(client *http.Client) *UpgradeCheckRunTestsParams {
	return &UpgradeCheckRunTestsParams{
		HTTPClient: client,
	}
}

/*
UpgradeCheckRunTestsParams contains all the parameters to send to the API endpoint

	for the upgrade check run tests operation.

	Typically these are written to a http.Request.
*/
type UpgradeCheckRunTestsParams struct {

	/* Body.

	   Run upgrade checks on cluster.
	*/
	Body *models.UpgradeCheckRunTestsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade check run tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeCheckRunTestsParams) WithDefaults() *UpgradeCheckRunTestsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade check run tests params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeCheckRunTestsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) WithTimeout(timeout time.Duration) *UpgradeCheckRunTestsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) WithContext(ctx context.Context) *UpgradeCheckRunTestsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) WithHTTPClient(client *http.Client) *UpgradeCheckRunTestsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) WithBody(body *models.UpgradeCheckRunTestsRequest) *UpgradeCheckRunTestsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade check run tests params
func (o *UpgradeCheckRunTestsParams) SetBody(body *models.UpgradeCheckRunTestsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeCheckRunTestsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

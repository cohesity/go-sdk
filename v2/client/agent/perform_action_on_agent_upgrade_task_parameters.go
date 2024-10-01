// Code generated by go-swagger; DO NOT EDIT.

package agent

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

	"github.com/cohesity/go-sdk/v2/models"
)

// NewPerformActionOnAgentUpgradeTaskParams creates a new PerformActionOnAgentUpgradeTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformActionOnAgentUpgradeTaskParams() *PerformActionOnAgentUpgradeTaskParams {
	return &PerformActionOnAgentUpgradeTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformActionOnAgentUpgradeTaskParamsWithTimeout creates a new PerformActionOnAgentUpgradeTaskParams object
// with the ability to set a timeout on a request.
func NewPerformActionOnAgentUpgradeTaskParamsWithTimeout(timeout time.Duration) *PerformActionOnAgentUpgradeTaskParams {
	return &PerformActionOnAgentUpgradeTaskParams{
		timeout: timeout,
	}
}

// NewPerformActionOnAgentUpgradeTaskParamsWithContext creates a new PerformActionOnAgentUpgradeTaskParams object
// with the ability to set a context for a request.
func NewPerformActionOnAgentUpgradeTaskParamsWithContext(ctx context.Context) *PerformActionOnAgentUpgradeTaskParams {
	return &PerformActionOnAgentUpgradeTaskParams{
		Context: ctx,
	}
}

// NewPerformActionOnAgentUpgradeTaskParamsWithHTTPClient creates a new PerformActionOnAgentUpgradeTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformActionOnAgentUpgradeTaskParamsWithHTTPClient(client *http.Client) *PerformActionOnAgentUpgradeTaskParams {
	return &PerformActionOnAgentUpgradeTaskParams{
		HTTPClient: client,
	}
}

/*
PerformActionOnAgentUpgradeTaskParams contains all the parameters to send to the API endpoint

	for the perform action on agent upgrade task operation.

	Typically these are written to a http.Request.
*/
type PerformActionOnAgentUpgradeTaskParams struct {

	/* Body.

	   Specifies the parameters to perform an action on an agent upgrade task.
	*/
	Body *models.AgentUpgradeTaskActionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the perform action on agent upgrade task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformActionOnAgentUpgradeTaskParams) WithDefaults() *PerformActionOnAgentUpgradeTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the perform action on agent upgrade task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformActionOnAgentUpgradeTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) WithTimeout(timeout time.Duration) *PerformActionOnAgentUpgradeTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) WithContext(ctx context.Context) *PerformActionOnAgentUpgradeTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) WithHTTPClient(client *http.Client) *PerformActionOnAgentUpgradeTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) WithBody(body *models.AgentUpgradeTaskActionRequest) *PerformActionOnAgentUpgradeTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the perform action on agent upgrade task params
func (o *PerformActionOnAgentUpgradeTaskParams) SetBody(body *models.AgentUpgradeTaskActionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PerformActionOnAgentUpgradeTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

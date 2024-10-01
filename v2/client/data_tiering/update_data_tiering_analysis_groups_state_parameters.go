// Code generated by go-swagger; DO NOT EDIT.

package data_tiering

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

// NewUpdateDataTieringAnalysisGroupsStateParams creates a new UpdateDataTieringAnalysisGroupsStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateDataTieringAnalysisGroupsStateParams() *UpdateDataTieringAnalysisGroupsStateParams {
	return &UpdateDataTieringAnalysisGroupsStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDataTieringAnalysisGroupsStateParamsWithTimeout creates a new UpdateDataTieringAnalysisGroupsStateParams object
// with the ability to set a timeout on a request.
func NewUpdateDataTieringAnalysisGroupsStateParamsWithTimeout(timeout time.Duration) *UpdateDataTieringAnalysisGroupsStateParams {
	return &UpdateDataTieringAnalysisGroupsStateParams{
		timeout: timeout,
	}
}

// NewUpdateDataTieringAnalysisGroupsStateParamsWithContext creates a new UpdateDataTieringAnalysisGroupsStateParams object
// with the ability to set a context for a request.
func NewUpdateDataTieringAnalysisGroupsStateParamsWithContext(ctx context.Context) *UpdateDataTieringAnalysisGroupsStateParams {
	return &UpdateDataTieringAnalysisGroupsStateParams{
		Context: ctx,
	}
}

// NewUpdateDataTieringAnalysisGroupsStateParamsWithHTTPClient creates a new UpdateDataTieringAnalysisGroupsStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateDataTieringAnalysisGroupsStateParamsWithHTTPClient(client *http.Client) *UpdateDataTieringAnalysisGroupsStateParams {
	return &UpdateDataTieringAnalysisGroupsStateParams{
		HTTPClient: client,
	}
}

/*
UpdateDataTieringAnalysisGroupsStateParams contains all the parameters to send to the API endpoint

	for the update data tiering analysis groups state operation.

	Typically these are written to a http.Request.
*/
type UpdateDataTieringAnalysisGroupsStateParams struct {

	/* Body.

	     Specifies the parameters to perform an action of list of data tiering
	analysis groups.
	*/
	Body *models.UpdateDataTieringStateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update data tiering analysis groups state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataTieringAnalysisGroupsStateParams) WithDefaults() *UpdateDataTieringAnalysisGroupsStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update data tiering analysis groups state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateDataTieringAnalysisGroupsStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) WithTimeout(timeout time.Duration) *UpdateDataTieringAnalysisGroupsStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) WithContext(ctx context.Context) *UpdateDataTieringAnalysisGroupsStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) WithHTTPClient(client *http.Client) *UpdateDataTieringAnalysisGroupsStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) WithBody(body *models.UpdateDataTieringStateRequest) *UpdateDataTieringAnalysisGroupsStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update data tiering analysis groups state params
func (o *UpdateDataTieringAnalysisGroupsStateParams) SetBody(body *models.UpdateDataTieringStateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDataTieringAnalysisGroupsStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

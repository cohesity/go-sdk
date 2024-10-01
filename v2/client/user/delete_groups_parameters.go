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

// NewDeleteGroupsParams creates a new DeleteGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteGroupsParams() *DeleteGroupsParams {
	return &DeleteGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteGroupsParamsWithTimeout creates a new DeleteGroupsParams object
// with the ability to set a timeout on a request.
func NewDeleteGroupsParamsWithTimeout(timeout time.Duration) *DeleteGroupsParams {
	return &DeleteGroupsParams{
		timeout: timeout,
	}
}

// NewDeleteGroupsParamsWithContext creates a new DeleteGroupsParams object
// with the ability to set a context for a request.
func NewDeleteGroupsParamsWithContext(ctx context.Context) *DeleteGroupsParams {
	return &DeleteGroupsParams{
		Context: ctx,
	}
}

// NewDeleteGroupsParamsWithHTTPClient creates a new DeleteGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteGroupsParamsWithHTTPClient(client *http.Client) *DeleteGroupsParams {
	return &DeleteGroupsParams{
		HTTPClient: client,
	}
}

/*
DeleteGroupsParams contains all the parameters to send to the API endpoint

	for the delete groups operation.

	Typically these are written to a http.Request.
*/
type DeleteGroupsParams struct {

	/* Body.

	   Specifies the delete request payload.
	*/
	Body *models.DeleteGroupsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupsParams) WithDefaults() *DeleteGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteGroupsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete groups params
func (o *DeleteGroupsParams) WithTimeout(timeout time.Duration) *DeleteGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete groups params
func (o *DeleteGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete groups params
func (o *DeleteGroupsParams) WithContext(ctx context.Context) *DeleteGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete groups params
func (o *DeleteGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete groups params
func (o *DeleteGroupsParams) WithHTTPClient(client *http.Client) *DeleteGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete groups params
func (o *DeleteGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete groups params
func (o *DeleteGroupsParams) WithBody(body *models.DeleteGroupsRequest) *DeleteGroupsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete groups params
func (o *DeleteGroupsParams) SetBody(body *models.DeleteGroupsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

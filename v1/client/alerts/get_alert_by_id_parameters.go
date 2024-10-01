// Code generated by go-swagger; DO NOT EDIT.

package alerts

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
)

// NewGetAlertByIDParams creates a new GetAlertByIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAlertByIDParams() *GetAlertByIDParams {
	return &GetAlertByIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlertByIDParamsWithTimeout creates a new GetAlertByIDParams object
// with the ability to set a timeout on a request.
func NewGetAlertByIDParamsWithTimeout(timeout time.Duration) *GetAlertByIDParams {
	return &GetAlertByIDParams{
		timeout: timeout,
	}
}

// NewGetAlertByIDParamsWithContext creates a new GetAlertByIDParams object
// with the ability to set a context for a request.
func NewGetAlertByIDParamsWithContext(ctx context.Context) *GetAlertByIDParams {
	return &GetAlertByIDParams{
		Context: ctx,
	}
}

// NewGetAlertByIDParamsWithHTTPClient creates a new GetAlertByIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAlertByIDParamsWithHTTPClient(client *http.Client) *GetAlertByIDParams {
	return &GetAlertByIDParams{
		HTTPClient: client,
	}
}

/*
GetAlertByIDParams contains all the parameters to send to the API endpoint

	for the get alert by Id operation.

	Typically these are written to a http.Request.
*/
type GetAlertByIDParams struct {

	/* ID.

	   Unique id of the Alert to return.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get alert by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertByIDParams) WithDefaults() *GetAlertByIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get alert by Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAlertByIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get alert by Id params
func (o *GetAlertByIDParams) WithTimeout(timeout time.Duration) *GetAlertByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alert by Id params
func (o *GetAlertByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alert by Id params
func (o *GetAlertByIDParams) WithContext(ctx context.Context) *GetAlertByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alert by Id params
func (o *GetAlertByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alert by Id params
func (o *GetAlertByIDParams) WithHTTPClient(client *http.Client) *GetAlertByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alert by Id params
func (o *GetAlertByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get alert by Id params
func (o *GetAlertByIDParams) WithID(id string) *GetAlertByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get alert by Id params
func (o *GetAlertByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlertByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

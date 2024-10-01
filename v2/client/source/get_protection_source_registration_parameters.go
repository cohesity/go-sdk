// Code generated by go-swagger; DO NOT EDIT.

package source

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

// NewGetProtectionSourceRegistrationParams creates a new GetProtectionSourceRegistrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProtectionSourceRegistrationParams() *GetProtectionSourceRegistrationParams {
	return &GetProtectionSourceRegistrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProtectionSourceRegistrationParamsWithTimeout creates a new GetProtectionSourceRegistrationParams object
// with the ability to set a timeout on a request.
func NewGetProtectionSourceRegistrationParamsWithTimeout(timeout time.Duration) *GetProtectionSourceRegistrationParams {
	return &GetProtectionSourceRegistrationParams{
		timeout: timeout,
	}
}

// NewGetProtectionSourceRegistrationParamsWithContext creates a new GetProtectionSourceRegistrationParams object
// with the ability to set a context for a request.
func NewGetProtectionSourceRegistrationParamsWithContext(ctx context.Context) *GetProtectionSourceRegistrationParams {
	return &GetProtectionSourceRegistrationParams{
		Context: ctx,
	}
}

// NewGetProtectionSourceRegistrationParamsWithHTTPClient creates a new GetProtectionSourceRegistrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProtectionSourceRegistrationParamsWithHTTPClient(client *http.Client) *GetProtectionSourceRegistrationParams {
	return &GetProtectionSourceRegistrationParams{
		HTTPClient: client,
	}
}

/*
GetProtectionSourceRegistrationParams contains all the parameters to send to the API endpoint

	for the get protection source registration operation.

	Typically these are written to a http.Request.
*/
type GetProtectionSourceRegistrationParams struct {

	/* ID.

	   Specifies the id of the Protection Source registration.

	   Format: int64
	*/
	ID int64

	/* RequestInitiatorType.

	   Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
	*/
	RequestInitiatorType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get protection source registration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionSourceRegistrationParams) WithDefaults() *GetProtectionSourceRegistrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get protection source registration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProtectionSourceRegistrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) WithTimeout(timeout time.Duration) *GetProtectionSourceRegistrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) WithContext(ctx context.Context) *GetProtectionSourceRegistrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) WithHTTPClient(client *http.Client) *GetProtectionSourceRegistrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) WithID(id int64) *GetProtectionSourceRegistrationParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) SetID(id int64) {
	o.ID = id
}

// WithRequestInitiatorType adds the requestInitiatorType to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) WithRequestInitiatorType(requestInitiatorType *string) *GetProtectionSourceRegistrationParams {
	o.SetRequestInitiatorType(requestInitiatorType)
	return o
}

// SetRequestInitiatorType adds the requestInitiatorType to the get protection source registration params
func (o *GetProtectionSourceRegistrationParams) SetRequestInitiatorType(requestInitiatorType *string) {
	o.RequestInitiatorType = requestInitiatorType
}

// WriteToRequest writes these params to a swagger request
func (o *GetProtectionSourceRegistrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.RequestInitiatorType != nil {

		// header param requestInitiatorType
		if err := r.SetHeaderParam("requestInitiatorType", *o.RequestInitiatorType); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

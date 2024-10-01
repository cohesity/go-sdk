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
	"github.com/go-openapi/swag"

	"github.com/cohesity/go-sdk/v2/models"
)

// NewUpdateInterfaceGroupParams creates a new UpdateInterfaceGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateInterfaceGroupParams() *UpdateInterfaceGroupParams {
	return &UpdateInterfaceGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateInterfaceGroupParamsWithTimeout creates a new UpdateInterfaceGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateInterfaceGroupParamsWithTimeout(timeout time.Duration) *UpdateInterfaceGroupParams {
	return &UpdateInterfaceGroupParams{
		timeout: timeout,
	}
}

// NewUpdateInterfaceGroupParamsWithContext creates a new UpdateInterfaceGroupParams object
// with the ability to set a context for a request.
func NewUpdateInterfaceGroupParamsWithContext(ctx context.Context) *UpdateInterfaceGroupParams {
	return &UpdateInterfaceGroupParams{
		Context: ctx,
	}
}

// NewUpdateInterfaceGroupParamsWithHTTPClient creates a new UpdateInterfaceGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateInterfaceGroupParamsWithHTTPClient(client *http.Client) *UpdateInterfaceGroupParams {
	return &UpdateInterfaceGroupParams{
		HTTPClient: client,
	}
}

/*
UpdateInterfaceGroupParams contains all the parameters to send to the API endpoint

	for the update interface group operation.

	Typically these are written to a http.Request.
*/
type UpdateInterfaceGroupParams struct {

	/* Body.

	   Parameters to update an interface group on the cluster.
	*/
	Body *models.InterfaceGroupParams

	/* ID.

	   Id of the interface group.

	   Format: int32
	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update interface group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInterfaceGroupParams) WithDefaults() *UpdateInterfaceGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update interface group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateInterfaceGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update interface group params
func (o *UpdateInterfaceGroupParams) WithTimeout(timeout time.Duration) *UpdateInterfaceGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update interface group params
func (o *UpdateInterfaceGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update interface group params
func (o *UpdateInterfaceGroupParams) WithContext(ctx context.Context) *UpdateInterfaceGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update interface group params
func (o *UpdateInterfaceGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update interface group params
func (o *UpdateInterfaceGroupParams) WithHTTPClient(client *http.Client) *UpdateInterfaceGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update interface group params
func (o *UpdateInterfaceGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update interface group params
func (o *UpdateInterfaceGroupParams) WithBody(body *models.InterfaceGroupParams) *UpdateInterfaceGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update interface group params
func (o *UpdateInterfaceGroupParams) SetBody(body *models.InterfaceGroupParams) {
	o.Body = body
}

// WithID adds the id to the update interface group params
func (o *UpdateInterfaceGroupParams) WithID(id int32) *UpdateInterfaceGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update interface group params
func (o *UpdateInterfaceGroupParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateInterfaceGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

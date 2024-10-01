// Code generated by go-swagger; DO NOT EDIT.

package object

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

// NewAssociateEntityMetadataParams creates a new AssociateEntityMetadataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAssociateEntityMetadataParams() *AssociateEntityMetadataParams {
	return &AssociateEntityMetadataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAssociateEntityMetadataParamsWithTimeout creates a new AssociateEntityMetadataParams object
// with the ability to set a timeout on a request.
func NewAssociateEntityMetadataParamsWithTimeout(timeout time.Duration) *AssociateEntityMetadataParams {
	return &AssociateEntityMetadataParams{
		timeout: timeout,
	}
}

// NewAssociateEntityMetadataParamsWithContext creates a new AssociateEntityMetadataParams object
// with the ability to set a context for a request.
func NewAssociateEntityMetadataParamsWithContext(ctx context.Context) *AssociateEntityMetadataParams {
	return &AssociateEntityMetadataParams{
		Context: ctx,
	}
}

// NewAssociateEntityMetadataParamsWithHTTPClient creates a new AssociateEntityMetadataParams object
// with the ability to set a custom HTTPClient for a request.
func NewAssociateEntityMetadataParamsWithHTTPClient(client *http.Client) *AssociateEntityMetadataParams {
	return &AssociateEntityMetadataParams{
		HTTPClient: client,
	}
}

/*
AssociateEntityMetadataParams contains all the parameters to send to the API endpoint

	for the associate entity metadata operation.

	Typically these are written to a http.Request.
*/
type AssociateEntityMetadataParams struct {

	/* Body.

	   Specifies the parameters to associate metadata with entities in the entity hierarchy.
	*/
	Body *models.AssociateEntityMetadataRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the associate entity metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociateEntityMetadataParams) WithDefaults() *AssociateEntityMetadataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the associate entity metadata params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AssociateEntityMetadataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the associate entity metadata params
func (o *AssociateEntityMetadataParams) WithTimeout(timeout time.Duration) *AssociateEntityMetadataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the associate entity metadata params
func (o *AssociateEntityMetadataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the associate entity metadata params
func (o *AssociateEntityMetadataParams) WithContext(ctx context.Context) *AssociateEntityMetadataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the associate entity metadata params
func (o *AssociateEntityMetadataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the associate entity metadata params
func (o *AssociateEntityMetadataParams) WithHTTPClient(client *http.Client) *AssociateEntityMetadataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the associate entity metadata params
func (o *AssociateEntityMetadataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the associate entity metadata params
func (o *AssociateEntityMetadataParams) WithBody(body *models.AssociateEntityMetadataRequest) *AssociateEntityMetadataParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the associate entity metadata params
func (o *AssociateEntityMetadataParams) SetBody(body *models.AssociateEntityMetadataRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AssociateEntityMetadataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

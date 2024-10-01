// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

	"github.com/cohesity/go-sdk/v1/models"
)

// NewDeleteActiveDirectoryEntryParams creates a new DeleteActiveDirectoryEntryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteActiveDirectoryEntryParams() *DeleteActiveDirectoryEntryParams {
	return &DeleteActiveDirectoryEntryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteActiveDirectoryEntryParamsWithTimeout creates a new DeleteActiveDirectoryEntryParams object
// with the ability to set a timeout on a request.
func NewDeleteActiveDirectoryEntryParamsWithTimeout(timeout time.Duration) *DeleteActiveDirectoryEntryParams {
	return &DeleteActiveDirectoryEntryParams{
		timeout: timeout,
	}
}

// NewDeleteActiveDirectoryEntryParamsWithContext creates a new DeleteActiveDirectoryEntryParams object
// with the ability to set a context for a request.
func NewDeleteActiveDirectoryEntryParamsWithContext(ctx context.Context) *DeleteActiveDirectoryEntryParams {
	return &DeleteActiveDirectoryEntryParams{
		Context: ctx,
	}
}

// NewDeleteActiveDirectoryEntryParamsWithHTTPClient creates a new DeleteActiveDirectoryEntryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteActiveDirectoryEntryParamsWithHTTPClient(client *http.Client) *DeleteActiveDirectoryEntryParams {
	return &DeleteActiveDirectoryEntryParams{
		HTTPClient: client,
	}
}

/*
DeleteActiveDirectoryEntryParams contains all the parameters to send to the API endpoint

	for the delete active directory entry operation.

	Typically these are written to a http.Request.
*/
type DeleteActiveDirectoryEntryParams struct {

	/* Body.

	   Request to delete a join with an Active Directory.
	*/
	Body *models.ActiveDirectoryEntry

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete active directory entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteActiveDirectoryEntryParams) WithDefaults() *DeleteActiveDirectoryEntryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete active directory entry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteActiveDirectoryEntryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) WithTimeout(timeout time.Duration) *DeleteActiveDirectoryEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) WithContext(ctx context.Context) *DeleteActiveDirectoryEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) WithHTTPClient(client *http.Client) *DeleteActiveDirectoryEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) WithBody(body *models.ActiveDirectoryEntry) *DeleteActiveDirectoryEntryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete active directory entry params
func (o *DeleteActiveDirectoryEntryParams) SetBody(body *models.ActiveDirectoryEntry) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteActiveDirectoryEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

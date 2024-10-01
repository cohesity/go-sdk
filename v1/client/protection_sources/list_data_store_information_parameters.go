// Code generated by go-swagger; DO NOT EDIT.

package protection_sources

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

// NewListDataStoreInformationParams creates a new ListDataStoreInformationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDataStoreInformationParams() *ListDataStoreInformationParams {
	return &ListDataStoreInformationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDataStoreInformationParamsWithTimeout creates a new ListDataStoreInformationParams object
// with the ability to set a timeout on a request.
func NewListDataStoreInformationParamsWithTimeout(timeout time.Duration) *ListDataStoreInformationParams {
	return &ListDataStoreInformationParams{
		timeout: timeout,
	}
}

// NewListDataStoreInformationParamsWithContext creates a new ListDataStoreInformationParams object
// with the ability to set a context for a request.
func NewListDataStoreInformationParamsWithContext(ctx context.Context) *ListDataStoreInformationParams {
	return &ListDataStoreInformationParams{
		Context: ctx,
	}
}

// NewListDataStoreInformationParamsWithHTTPClient creates a new ListDataStoreInformationParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDataStoreInformationParamsWithHTTPClient(client *http.Client) *ListDataStoreInformationParams {
	return &ListDataStoreInformationParams{
		HTTPClient: client,
	}
}

/*
ListDataStoreInformationParams contains all the parameters to send to the API endpoint

	for the list data store information operation.

	Typically these are written to a http.Request.
*/
type ListDataStoreInformationParams struct {

	/* SourceID.

	   Specifies the id of the virtual machine in vmware environment.

	   Format: int64
	*/
	SourceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list data store information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDataStoreInformationParams) WithDefaults() *ListDataStoreInformationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list data store information params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDataStoreInformationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list data store information params
func (o *ListDataStoreInformationParams) WithTimeout(timeout time.Duration) *ListDataStoreInformationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list data store information params
func (o *ListDataStoreInformationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list data store information params
func (o *ListDataStoreInformationParams) WithContext(ctx context.Context) *ListDataStoreInformationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list data store information params
func (o *ListDataStoreInformationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list data store information params
func (o *ListDataStoreInformationParams) WithHTTPClient(client *http.Client) *ListDataStoreInformationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list data store information params
func (o *ListDataStoreInformationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSourceID adds the sourceID to the list data store information params
func (o *ListDataStoreInformationParams) WithSourceID(sourceID int64) *ListDataStoreInformationParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the list data store information params
func (o *ListDataStoreInformationParams) SetSourceID(sourceID int64) {
	o.SourceID = sourceID
}

// WriteToRequest writes these params to a swagger request
func (o *ListDataStoreInformationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param sourceId
	qrSourceID := o.SourceID
	qSourceID := swag.FormatInt64(qrSourceID)
	if qSourceID != "" {

		if err := r.SetQueryParam("sourceId", qSourceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

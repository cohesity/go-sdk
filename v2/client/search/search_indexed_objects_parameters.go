// Code generated by go-swagger; DO NOT EDIT.

package search

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

// NewSearchIndexedObjectsParams creates a new SearchIndexedObjectsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSearchIndexedObjectsParams() *SearchIndexedObjectsParams {
	return &SearchIndexedObjectsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSearchIndexedObjectsParamsWithTimeout creates a new SearchIndexedObjectsParams object
// with the ability to set a timeout on a request.
func NewSearchIndexedObjectsParamsWithTimeout(timeout time.Duration) *SearchIndexedObjectsParams {
	return &SearchIndexedObjectsParams{
		timeout: timeout,
	}
}

// NewSearchIndexedObjectsParamsWithContext creates a new SearchIndexedObjectsParams object
// with the ability to set a context for a request.
func NewSearchIndexedObjectsParamsWithContext(ctx context.Context) *SearchIndexedObjectsParams {
	return &SearchIndexedObjectsParams{
		Context: ctx,
	}
}

// NewSearchIndexedObjectsParamsWithHTTPClient creates a new SearchIndexedObjectsParams object
// with the ability to set a custom HTTPClient for a request.
func NewSearchIndexedObjectsParamsWithHTTPClient(client *http.Client) *SearchIndexedObjectsParams {
	return &SearchIndexedObjectsParams{
		HTTPClient: client,
	}
}

/*
SearchIndexedObjectsParams contains all the parameters to send to the API endpoint

	for the search indexed objects operation.

	Typically these are written to a http.Request.
*/
type SearchIndexedObjectsParams struct {

	/* Body.

	   Specifies the parameters to search for indexed objects.
	*/
	Body *models.SearchIndexedObjectsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the search indexed objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchIndexedObjectsParams) WithDefaults() *SearchIndexedObjectsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the search indexed objects params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SearchIndexedObjectsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the search indexed objects params
func (o *SearchIndexedObjectsParams) WithTimeout(timeout time.Duration) *SearchIndexedObjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search indexed objects params
func (o *SearchIndexedObjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search indexed objects params
func (o *SearchIndexedObjectsParams) WithContext(ctx context.Context) *SearchIndexedObjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search indexed objects params
func (o *SearchIndexedObjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search indexed objects params
func (o *SearchIndexedObjectsParams) WithHTTPClient(client *http.Client) *SearchIndexedObjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search indexed objects params
func (o *SearchIndexedObjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the search indexed objects params
func (o *SearchIndexedObjectsParams) WithBody(body *models.SearchIndexedObjectsRequest) *SearchIndexedObjectsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the search indexed objects params
func (o *SearchIndexedObjectsParams) SetBody(body *models.SearchIndexedObjectsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SearchIndexedObjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

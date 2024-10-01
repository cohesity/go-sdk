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
)

// NewGetIpmiUsersParams creates a new GetIpmiUsersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIpmiUsersParams() *GetIpmiUsersParams {
	return &GetIpmiUsersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIpmiUsersParamsWithTimeout creates a new GetIpmiUsersParams object
// with the ability to set a timeout on a request.
func NewGetIpmiUsersParamsWithTimeout(timeout time.Duration) *GetIpmiUsersParams {
	return &GetIpmiUsersParams{
		timeout: timeout,
	}
}

// NewGetIpmiUsersParamsWithContext creates a new GetIpmiUsersParams object
// with the ability to set a context for a request.
func NewGetIpmiUsersParamsWithContext(ctx context.Context) *GetIpmiUsersParams {
	return &GetIpmiUsersParams{
		Context: ctx,
	}
}

// NewGetIpmiUsersParamsWithHTTPClient creates a new GetIpmiUsersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIpmiUsersParamsWithHTTPClient(client *http.Client) *GetIpmiUsersParams {
	return &GetIpmiUsersParams{
		HTTPClient: client,
	}
}

/*
GetIpmiUsersParams contains all the parameters to send to the API endpoint

	for the get ipmi users operation.

	Typically these are written to a http.Request.
*/
type GetIpmiUsersParams struct {

	/* Ids.

	   Node ids to filter node IPMI users.
	*/
	Ids []int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ipmi users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpmiUsersParams) WithDefaults() *GetIpmiUsersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ipmi users params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIpmiUsersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ipmi users params
func (o *GetIpmiUsersParams) WithTimeout(timeout time.Duration) *GetIpmiUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ipmi users params
func (o *GetIpmiUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ipmi users params
func (o *GetIpmiUsersParams) WithContext(ctx context.Context) *GetIpmiUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ipmi users params
func (o *GetIpmiUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ipmi users params
func (o *GetIpmiUsersParams) WithHTTPClient(client *http.Client) *GetIpmiUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ipmi users params
func (o *GetIpmiUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get ipmi users params
func (o *GetIpmiUsersParams) WithIds(ids []int64) *GetIpmiUsersParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get ipmi users params
func (o *GetIpmiUsersParams) SetIds(ids []int64) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *GetIpmiUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ids != nil {

		// binding items for ids
		joinedIds := o.bindParamIds(reg)

		// query array param ids
		if err := r.SetQueryParam("ids", joinedIds...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetIpmiUsers binds the parameter ids
func (o *GetIpmiUsersParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []int64

		idsIIV := swag.FormatInt64(idsIIR) // int64 as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: ""
	idsIS := swag.JoinByFormat(idsIC, "")

	return idsIS
}

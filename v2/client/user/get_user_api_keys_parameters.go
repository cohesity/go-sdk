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
	"github.com/go-openapi/swag"
)

// NewGetUserAPIKeysParams creates a new GetUserAPIKeysParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserAPIKeysParams() *GetUserAPIKeysParams {
	return &GetUserAPIKeysParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAPIKeysParamsWithTimeout creates a new GetUserAPIKeysParams object
// with the ability to set a timeout on a request.
func NewGetUserAPIKeysParamsWithTimeout(timeout time.Duration) *GetUserAPIKeysParams {
	return &GetUserAPIKeysParams{
		timeout: timeout,
	}
}

// NewGetUserAPIKeysParamsWithContext creates a new GetUserAPIKeysParams object
// with the ability to set a context for a request.
func NewGetUserAPIKeysParamsWithContext(ctx context.Context) *GetUserAPIKeysParams {
	return &GetUserAPIKeysParams{
		Context: ctx,
	}
}

// NewGetUserAPIKeysParamsWithHTTPClient creates a new GetUserAPIKeysParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserAPIKeysParamsWithHTTPClient(client *http.Client) *GetUserAPIKeysParams {
	return &GetUserAPIKeysParams{
		HTTPClient: client,
	}
}

/*
GetUserAPIKeysParams contains all the parameters to send to the API endpoint

	for the get user API keys operation.

	Typically these are written to a http.Request.
*/
type GetUserAPIKeysParams struct {

	/* Ids.

	   Filter by API Key Ids
	*/
	Ids []string

	/* IsActive.

	   If true, the response will only include API keys which are active. Returns all keys if the query param is not set.
	*/
	IsActive *bool

	/* IsExpired.

	   If true, the response will only include API keys which has been expired. Returns all keys if the query param is not set.
	*/
	IsExpired *bool

	/* UserSid.

	   Specify the SID of the API key owner.
	*/
	UserSid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user API keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAPIKeysParams) WithDefaults() *GetUserAPIKeysParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user API keys params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserAPIKeysParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user API keys params
func (o *GetUserAPIKeysParams) WithTimeout(timeout time.Duration) *GetUserAPIKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user API keys params
func (o *GetUserAPIKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user API keys params
func (o *GetUserAPIKeysParams) WithContext(ctx context.Context) *GetUserAPIKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user API keys params
func (o *GetUserAPIKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user API keys params
func (o *GetUserAPIKeysParams) WithHTTPClient(client *http.Client) *GetUserAPIKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user API keys params
func (o *GetUserAPIKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the get user API keys params
func (o *GetUserAPIKeysParams) WithIds(ids []string) *GetUserAPIKeysParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get user API keys params
func (o *GetUserAPIKeysParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithIsActive adds the isActive to the get user API keys params
func (o *GetUserAPIKeysParams) WithIsActive(isActive *bool) *GetUserAPIKeysParams {
	o.SetIsActive(isActive)
	return o
}

// SetIsActive adds the isActive to the get user API keys params
func (o *GetUserAPIKeysParams) SetIsActive(isActive *bool) {
	o.IsActive = isActive
}

// WithIsExpired adds the isExpired to the get user API keys params
func (o *GetUserAPIKeysParams) WithIsExpired(isExpired *bool) *GetUserAPIKeysParams {
	o.SetIsExpired(isExpired)
	return o
}

// SetIsExpired adds the isExpired to the get user API keys params
func (o *GetUserAPIKeysParams) SetIsExpired(isExpired *bool) {
	o.IsExpired = isExpired
}

// WithUserSid adds the userSid to the get user API keys params
func (o *GetUserAPIKeysParams) WithUserSid(userSid string) *GetUserAPIKeysParams {
	o.SetUserSid(userSid)
	return o
}

// SetUserSid adds the userSid to the get user API keys params
func (o *GetUserAPIKeysParams) SetUserSid(userSid string) {
	o.UserSid = userSid
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAPIKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IsActive != nil {

		// query param isActive
		var qrIsActive bool

		if o.IsActive != nil {
			qrIsActive = *o.IsActive
		}
		qIsActive := swag.FormatBool(qrIsActive)
		if qIsActive != "" {

			if err := r.SetQueryParam("isActive", qIsActive); err != nil {
				return err
			}
		}
	}

	if o.IsExpired != nil {

		// query param isExpired
		var qrIsExpired bool

		if o.IsExpired != nil {
			qrIsExpired = *o.IsExpired
		}
		qIsExpired := swag.FormatBool(qrIsExpired)
		if qIsExpired != "" {

			if err := r.SetQueryParam("isExpired", qIsExpired); err != nil {
				return err
			}
		}
	}

	// path param userSid
	if err := r.SetPathParam("userSid", o.UserSid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetUserAPIKeys binds the parameter ids
func (o *GetUserAPIKeysParams) bindParamIds(formats strfmt.Registry) []string {
	idsIR := o.Ids

	var idsIC []string
	for _, idsIIR := range idsIR { // explode []string

		idsIIV := idsIIR // string as string
		idsIC = append(idsIC, idsIIV)
	}

	// items.CollectionFormat: ""
	idsIS := swag.JoinByFormat(idsIC, "")

	return idsIS
}

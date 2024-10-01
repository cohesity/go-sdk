// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service_group

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

// NewGetIcapConnectionStatusParams creates a new GetIcapConnectionStatusParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetIcapConnectionStatusParams() *GetIcapConnectionStatusParams {
	return &GetIcapConnectionStatusParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetIcapConnectionStatusParamsWithTimeout creates a new GetIcapConnectionStatusParams object
// with the ability to set a timeout on a request.
func NewGetIcapConnectionStatusParamsWithTimeout(timeout time.Duration) *GetIcapConnectionStatusParams {
	return &GetIcapConnectionStatusParams{
		timeout: timeout,
	}
}

// NewGetIcapConnectionStatusParamsWithContext creates a new GetIcapConnectionStatusParams object
// with the ability to set a context for a request.
func NewGetIcapConnectionStatusParamsWithContext(ctx context.Context) *GetIcapConnectionStatusParams {
	return &GetIcapConnectionStatusParams{
		Context: ctx,
	}
}

// NewGetIcapConnectionStatusParamsWithHTTPClient creates a new GetIcapConnectionStatusParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetIcapConnectionStatusParamsWithHTTPClient(client *http.Client) *GetIcapConnectionStatusParams {
	return &GetIcapConnectionStatusParams{
		HTTPClient: client,
	}
}

/*
GetIcapConnectionStatusParams contains all the parameters to send to the API endpoint

	for the get icap connection status operation.

	Typically these are written to a http.Request.
*/
type GetIcapConnectionStatusParams struct {

	/* IcapUris.

	   Specifies the list of icap uri.
	*/
	IcapUris []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get icap connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIcapConnectionStatusParams) WithDefaults() *GetIcapConnectionStatusParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get icap connection status params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetIcapConnectionStatusParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get icap connection status params
func (o *GetIcapConnectionStatusParams) WithTimeout(timeout time.Duration) *GetIcapConnectionStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get icap connection status params
func (o *GetIcapConnectionStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get icap connection status params
func (o *GetIcapConnectionStatusParams) WithContext(ctx context.Context) *GetIcapConnectionStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get icap connection status params
func (o *GetIcapConnectionStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get icap connection status params
func (o *GetIcapConnectionStatusParams) WithHTTPClient(client *http.Client) *GetIcapConnectionStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get icap connection status params
func (o *GetIcapConnectionStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIcapUris adds the icapUris to the get icap connection status params
func (o *GetIcapConnectionStatusParams) WithIcapUris(icapUris []string) *GetIcapConnectionStatusParams {
	o.SetIcapUris(icapUris)
	return o
}

// SetIcapUris adds the icapUris to the get icap connection status params
func (o *GetIcapConnectionStatusParams) SetIcapUris(icapUris []string) {
	o.IcapUris = icapUris
}

// WriteToRequest writes these params to a swagger request
func (o *GetIcapConnectionStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IcapUris != nil {

		// binding items for icapUris
		joinedIcapUris := o.bindParamIcapUris(reg)

		// query array param icapUris
		if err := r.SetQueryParam("icapUris", joinedIcapUris...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetIcapConnectionStatus binds the parameter icapUris
func (o *GetIcapConnectionStatusParams) bindParamIcapUris(formats strfmt.Registry) []string {
	icapUrisIR := o.IcapUris

	var icapUrisIC []string
	for _, icapUrisIIR := range icapUrisIR { // explode []string

		icapUrisIIV := icapUrisIIR // string as string
		icapUrisIC = append(icapUrisIC, icapUrisIIV)
	}

	// items.CollectionFormat: ""
	icapUrisIS := swag.JoinByFormat(icapUrisIC, "")

	return icapUrisIS
}

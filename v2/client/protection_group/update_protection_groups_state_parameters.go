// Code generated by go-swagger; DO NOT EDIT.

package protection_group

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

// NewUpdateProtectionGroupsStateParams creates a new UpdateProtectionGroupsStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateProtectionGroupsStateParams() *UpdateProtectionGroupsStateParams {
	return &UpdateProtectionGroupsStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProtectionGroupsStateParamsWithTimeout creates a new UpdateProtectionGroupsStateParams object
// with the ability to set a timeout on a request.
func NewUpdateProtectionGroupsStateParamsWithTimeout(timeout time.Duration) *UpdateProtectionGroupsStateParams {
	return &UpdateProtectionGroupsStateParams{
		timeout: timeout,
	}
}

// NewUpdateProtectionGroupsStateParamsWithContext creates a new UpdateProtectionGroupsStateParams object
// with the ability to set a context for a request.
func NewUpdateProtectionGroupsStateParamsWithContext(ctx context.Context) *UpdateProtectionGroupsStateParams {
	return &UpdateProtectionGroupsStateParams{
		Context: ctx,
	}
}

// NewUpdateProtectionGroupsStateParamsWithHTTPClient creates a new UpdateProtectionGroupsStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateProtectionGroupsStateParamsWithHTTPClient(client *http.Client) *UpdateProtectionGroupsStateParams {
	return &UpdateProtectionGroupsStateParams{
		HTTPClient: client,
	}
}

/*
UpdateProtectionGroupsStateParams contains all the parameters to send to the API endpoint

	for the update protection groups state operation.

	Typically these are written to a http.Request.
*/
type UpdateProtectionGroupsStateParams struct {

	/* Body.

	   Specifies the parameters to perform an action of list of Protection Groups.
	*/
	Body *models.UpdateProtectionGroupsStateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update protection groups state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProtectionGroupsStateParams) WithDefaults() *UpdateProtectionGroupsStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update protection groups state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateProtectionGroupsStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) WithTimeout(timeout time.Duration) *UpdateProtectionGroupsStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) WithContext(ctx context.Context) *UpdateProtectionGroupsStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) WithHTTPClient(client *http.Client) *UpdateProtectionGroupsStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) WithBody(body *models.UpdateProtectionGroupsStateRequest) *UpdateProtectionGroupsStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update protection groups state params
func (o *UpdateProtectionGroupsStateParams) SetBody(body *models.UpdateProtectionGroupsStateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProtectionGroupsStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// NewCreateApplicationsRecoverTaskParams creates a new CreateApplicationsRecoverTaskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateApplicationsRecoverTaskParams() *CreateApplicationsRecoverTaskParams {
	return &CreateApplicationsRecoverTaskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateApplicationsRecoverTaskParamsWithTimeout creates a new CreateApplicationsRecoverTaskParams object
// with the ability to set a timeout on a request.
func NewCreateApplicationsRecoverTaskParamsWithTimeout(timeout time.Duration) *CreateApplicationsRecoverTaskParams {
	return &CreateApplicationsRecoverTaskParams{
		timeout: timeout,
	}
}

// NewCreateApplicationsRecoverTaskParamsWithContext creates a new CreateApplicationsRecoverTaskParams object
// with the ability to set a context for a request.
func NewCreateApplicationsRecoverTaskParamsWithContext(ctx context.Context) *CreateApplicationsRecoverTaskParams {
	return &CreateApplicationsRecoverTaskParams{
		Context: ctx,
	}
}

// NewCreateApplicationsRecoverTaskParamsWithHTTPClient creates a new CreateApplicationsRecoverTaskParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateApplicationsRecoverTaskParamsWithHTTPClient(client *http.Client) *CreateApplicationsRecoverTaskParams {
	return &CreateApplicationsRecoverTaskParams{
		HTTPClient: client,
	}
}

/*
CreateApplicationsRecoverTaskParams contains all the parameters to send to the API endpoint

	for the create applications recover task operation.

	Typically these are written to a http.Request.
*/
type CreateApplicationsRecoverTaskParams struct {

	/* Body.

	     Request to create a Restore Task for recovering Applications like SQL DB.
	volumes to mount points.
	*/
	Body *models.ApplicationsRestoreTaskRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create applications recover task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateApplicationsRecoverTaskParams) WithDefaults() *CreateApplicationsRecoverTaskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create applications recover task params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateApplicationsRecoverTaskParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) WithTimeout(timeout time.Duration) *CreateApplicationsRecoverTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) WithContext(ctx context.Context) *CreateApplicationsRecoverTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) WithHTTPClient(client *http.Client) *CreateApplicationsRecoverTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) WithBody(body *models.ApplicationsRestoreTaskRequest) *CreateApplicationsRecoverTaskParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create applications recover task params
func (o *CreateApplicationsRecoverTaskParams) SetBody(body *models.ApplicationsRestoreTaskRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateApplicationsRecoverTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

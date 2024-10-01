// Code generated by go-swagger; DO NOT EDIT.

package principals

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

// NewResetS3SecretKeyParams creates a new ResetS3SecretKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResetS3SecretKeyParams() *ResetS3SecretKeyParams {
	return &ResetS3SecretKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResetS3SecretKeyParamsWithTimeout creates a new ResetS3SecretKeyParams object
// with the ability to set a timeout on a request.
func NewResetS3SecretKeyParamsWithTimeout(timeout time.Duration) *ResetS3SecretKeyParams {
	return &ResetS3SecretKeyParams{
		timeout: timeout,
	}
}

// NewResetS3SecretKeyParamsWithContext creates a new ResetS3SecretKeyParams object
// with the ability to set a context for a request.
func NewResetS3SecretKeyParamsWithContext(ctx context.Context) *ResetS3SecretKeyParams {
	return &ResetS3SecretKeyParams{
		Context: ctx,
	}
}

// NewResetS3SecretKeyParamsWithHTTPClient creates a new ResetS3SecretKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewResetS3SecretKeyParamsWithHTTPClient(client *http.Client) *ResetS3SecretKeyParams {
	return &ResetS3SecretKeyParams{
		HTTPClient: client,
	}
}

/*
ResetS3SecretKeyParams contains all the parameters to send to the API endpoint

	for the reset s3 secret key operation.

	Typically these are written to a http.Request.
*/
type ResetS3SecretKeyParams struct {

	/* Body.

	   Request to reset the S3 secret access key for the specified Cohesity user.
	*/
	Body *models.ResetS3SecretKeyParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reset s3 secret key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetS3SecretKeyParams) WithDefaults() *ResetS3SecretKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reset s3 secret key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResetS3SecretKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) WithTimeout(timeout time.Duration) *ResetS3SecretKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) WithContext(ctx context.Context) *ResetS3SecretKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) WithHTTPClient(client *http.Client) *ResetS3SecretKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) WithBody(body *models.ResetS3SecretKeyParameters) *ResetS3SecretKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reset s3 secret key params
func (o *ResetS3SecretKeyParams) SetBody(body *models.ResetS3SecretKeyParameters) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ResetS3SecretKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

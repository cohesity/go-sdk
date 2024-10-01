// Code generated by go-swagger; DO NOT EDIT.

package identity_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// UpdateIdentityProviderReader is a Reader for the UpdateIdentityProvider structure.
type UpdateIdentityProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIdentityProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIdentityProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateIdentityProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIdentityProviderOK creates a UpdateIdentityProviderOK with default headers values
func NewUpdateIdentityProviderOK() *UpdateIdentityProviderOK {
	return &UpdateIdentityProviderOK{}
}

/*
UpdateIdentityProviderOK describes a response with status code 200, with default header values.

Success
*/
type UpdateIdentityProviderOK struct {
	Payload *models.IdentityProviderConfiguration
}

// IsSuccess returns true when this update identity provider o k response has a 2xx status code
func (o *UpdateIdentityProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update identity provider o k response has a 3xx status code
func (o *UpdateIdentityProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update identity provider o k response has a 4xx status code
func (o *UpdateIdentityProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update identity provider o k response has a 5xx status code
func (o *UpdateIdentityProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update identity provider o k response a status code equal to that given
func (o *UpdateIdentityProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update identity provider o k response
func (o *UpdateIdentityProviderOK) Code() int {
	return 200
}

func (o *UpdateIdentityProviderOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /idps/{id}][%d] updateIdentityProviderOK %s", 200, payload)
}

func (o *UpdateIdentityProviderOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /idps/{id}][%d] updateIdentityProviderOK %s", 200, payload)
}

func (o *UpdateIdentityProviderOK) GetPayload() *models.IdentityProviderConfiguration {
	return o.Payload
}

func (o *UpdateIdentityProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IdentityProviderConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIdentityProviderDefault creates a UpdateIdentityProviderDefault with default headers values
func NewUpdateIdentityProviderDefault(code int) *UpdateIdentityProviderDefault {
	return &UpdateIdentityProviderDefault{
		_statusCode: code,
	}
}

/*
UpdateIdentityProviderDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateIdentityProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update identity provider default response has a 2xx status code
func (o *UpdateIdentityProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update identity provider default response has a 3xx status code
func (o *UpdateIdentityProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update identity provider default response has a 4xx status code
func (o *UpdateIdentityProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update identity provider default response has a 5xx status code
func (o *UpdateIdentityProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update identity provider default response a status code equal to that given
func (o *UpdateIdentityProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update identity provider default response
func (o *UpdateIdentityProviderDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIdentityProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /idps/{id}][%d] UpdateIdentityProvider default %s", o._statusCode, payload)
}

func (o *UpdateIdentityProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /idps/{id}][%d] UpdateIdentityProvider default %s", o._statusCode, payload)
}

func (o *UpdateIdentityProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateIdentityProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

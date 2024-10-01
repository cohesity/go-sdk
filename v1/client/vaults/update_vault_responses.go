// Code generated by go-swagger; DO NOT EDIT.

package vaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// UpdateVaultReader is a Reader for the UpdateVault structure.
type UpdateVaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateVaultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateVaultOK creates a UpdateVaultOK with default headers values
func NewUpdateVaultOK() *UpdateVaultOK {
	return &UpdateVaultOK{}
}

/*
UpdateVaultOK describes a response with status code 200, with default header values.

Success
*/
type UpdateVaultOK struct {
	Payload *models.Vault
}

// IsSuccess returns true when this update vault o k response has a 2xx status code
func (o *UpdateVaultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update vault o k response has a 3xx status code
func (o *UpdateVaultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update vault o k response has a 4xx status code
func (o *UpdateVaultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update vault o k response has a 5xx status code
func (o *UpdateVaultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update vault o k response a status code equal to that given
func (o *UpdateVaultOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update vault o k response
func (o *UpdateVaultOK) Code() int {
	return 200
}

func (o *UpdateVaultOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/vaults/{id}][%d] updateVaultOK %s", 200, payload)
}

func (o *UpdateVaultOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/vaults/{id}][%d] updateVaultOK %s", 200, payload)
}

func (o *UpdateVaultOK) GetPayload() *models.Vault {
	return o.Payload
}

func (o *UpdateVaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVaultDefault creates a UpdateVaultDefault with default headers values
func NewUpdateVaultDefault(code int) *UpdateVaultDefault {
	return &UpdateVaultDefault{
		_statusCode: code,
	}
}

/*
UpdateVaultDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateVaultDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update vault default response has a 2xx status code
func (o *UpdateVaultDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update vault default response has a 3xx status code
func (o *UpdateVaultDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update vault default response has a 4xx status code
func (o *UpdateVaultDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update vault default response has a 5xx status code
func (o *UpdateVaultDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update vault default response a status code equal to that given
func (o *UpdateVaultDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update vault default response
func (o *UpdateVaultDefault) Code() int {
	return o._statusCode
}

func (o *UpdateVaultDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/vaults/{id}][%d] UpdateVault default %s", o._statusCode, payload)
}

func (o *UpdateVaultDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/vaults/{id}][%d] UpdateVault default %s", o._statusCode, payload)
}

func (o *UpdateVaultDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateVaultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

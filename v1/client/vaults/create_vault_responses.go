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

// CreateVaultReader is a Reader for the CreateVault structure.
type CreateVaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateVaultCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateVaultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateVaultCreated creates a CreateVaultCreated with default headers values
func NewCreateVaultCreated() *CreateVaultCreated {
	return &CreateVaultCreated{}
}

/*
CreateVaultCreated describes a response with status code 201, with default header values.

Success
*/
type CreateVaultCreated struct {
	Payload *models.Vault
}

// IsSuccess returns true when this create vault created response has a 2xx status code
func (o *CreateVaultCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create vault created response has a 3xx status code
func (o *CreateVaultCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create vault created response has a 4xx status code
func (o *CreateVaultCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create vault created response has a 5xx status code
func (o *CreateVaultCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create vault created response a status code equal to that given
func (o *CreateVaultCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create vault created response
func (o *CreateVaultCreated) Code() int {
	return 201
}

func (o *CreateVaultCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/vaults][%d] createVaultCreated %s", 201, payload)
}

func (o *CreateVaultCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/vaults][%d] createVaultCreated %s", 201, payload)
}

func (o *CreateVaultCreated) GetPayload() *models.Vault {
	return o.Payload
}

func (o *CreateVaultCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Vault)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVaultDefault creates a CreateVaultDefault with default headers values
func NewCreateVaultDefault(code int) *CreateVaultDefault {
	return &CreateVaultDefault{
		_statusCode: code,
	}
}

/*
CreateVaultDefault describes a response with status code -1, with default header values.

Error
*/
type CreateVaultDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create vault default response has a 2xx status code
func (o *CreateVaultDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create vault default response has a 3xx status code
func (o *CreateVaultDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create vault default response has a 4xx status code
func (o *CreateVaultDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create vault default response has a 5xx status code
func (o *CreateVaultDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create vault default response a status code equal to that given
func (o *CreateVaultDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create vault default response
func (o *CreateVaultDefault) Code() int {
	return o._statusCode
}

func (o *CreateVaultDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/vaults][%d] CreateVault default %s", o._statusCode, payload)
}

func (o *CreateVaultDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/vaults][%d] CreateVault default %s", o._statusCode, payload)
}

func (o *CreateVaultDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateVaultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

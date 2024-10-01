// Code generated by go-swagger; DO NOT EDIT.

package user

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

// CreateUserAPIKeyReader is a Reader for the CreateUserAPIKey structure.
type CreateUserAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserAPIKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateUserAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateUserAPIKeyCreated creates a CreateUserAPIKeyCreated with default headers values
func NewCreateUserAPIKeyCreated() *CreateUserAPIKeyCreated {
	return &CreateUserAPIKeyCreated{}
}

/*
CreateUserAPIKeyCreated describes a response with status code 201, with default header values.

Success
*/
type CreateUserAPIKeyCreated struct {
	Payload *models.CreatedUserAPIKey
}

// IsSuccess returns true when this create user Api key created response has a 2xx status code
func (o *CreateUserAPIKeyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user Api key created response has a 3xx status code
func (o *CreateUserAPIKeyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user Api key created response has a 4xx status code
func (o *CreateUserAPIKeyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user Api key created response has a 5xx status code
func (o *CreateUserAPIKeyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user Api key created response a status code equal to that given
func (o *CreateUserAPIKeyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user Api key created response
func (o *CreateUserAPIKeyCreated) Code() int {
	return 201
}

func (o *CreateUserAPIKeyCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/{userSid}/api-keys][%d] createUserApiKeyCreated %s", 201, payload)
}

func (o *CreateUserAPIKeyCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/{userSid}/api-keys][%d] createUserApiKeyCreated %s", 201, payload)
}

func (o *CreateUserAPIKeyCreated) GetPayload() *models.CreatedUserAPIKey {
	return o.Payload
}

func (o *CreateUserAPIKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatedUserAPIKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserAPIKeyDefault creates a CreateUserAPIKeyDefault with default headers values
func NewCreateUserAPIKeyDefault(code int) *CreateUserAPIKeyDefault {
	return &CreateUserAPIKeyDefault{
		_statusCode: code,
	}
}

/*
CreateUserAPIKeyDefault describes a response with status code -1, with default header values.

Error
*/
type CreateUserAPIKeyDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create user API key default response has a 2xx status code
func (o *CreateUserAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create user API key default response has a 3xx status code
func (o *CreateUserAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create user API key default response has a 4xx status code
func (o *CreateUserAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create user API key default response has a 5xx status code
func (o *CreateUserAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create user API key default response a status code equal to that given
func (o *CreateUserAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create user API key default response
func (o *CreateUserAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *CreateUserAPIKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/{userSid}/api-keys][%d] CreateUserAPIKey default %s", o._statusCode, payload)
}

func (o *CreateUserAPIKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /users/{userSid}/api-keys][%d] CreateUserAPIKey default %s", o._statusCode, payload)
}

func (o *CreateUserAPIKeyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

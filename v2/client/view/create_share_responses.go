// Code generated by go-swagger; DO NOT EDIT.

package view

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

// CreateShareReader is a Reader for the CreateShare structure.
type CreateShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateShareCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateShareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateShareCreated creates a CreateShareCreated with default headers values
func NewCreateShareCreated() *CreateShareCreated {
	return &CreateShareCreated{}
}

/*
CreateShareCreated describes a response with status code 201, with default header values.

Success
*/
type CreateShareCreated struct {
	Payload *models.Share
}

// IsSuccess returns true when this create share created response has a 2xx status code
func (o *CreateShareCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create share created response has a 3xx status code
func (o *CreateShareCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create share created response has a 4xx status code
func (o *CreateShareCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create share created response has a 5xx status code
func (o *CreateShareCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create share created response a status code equal to that given
func (o *CreateShareCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create share created response
func (o *CreateShareCreated) Code() int {
	return 201
}

func (o *CreateShareCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file-services/shares][%d] createShareCreated %s", 201, payload)
}

func (o *CreateShareCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file-services/shares][%d] createShareCreated %s", 201, payload)
}

func (o *CreateShareCreated) GetPayload() *models.Share {
	return o.Payload
}

func (o *CreateShareCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Share)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateShareDefault creates a CreateShareDefault with default headers values
func NewCreateShareDefault(code int) *CreateShareDefault {
	return &CreateShareDefault{
		_statusCode: code,
	}
}

/*
CreateShareDefault describes a response with status code -1, with default header values.

Error
*/
type CreateShareDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create share default response has a 2xx status code
func (o *CreateShareDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create share default response has a 3xx status code
func (o *CreateShareDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create share default response has a 4xx status code
func (o *CreateShareDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create share default response has a 5xx status code
func (o *CreateShareDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create share default response a status code equal to that given
func (o *CreateShareDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create share default response
func (o *CreateShareDefault) Code() int {
	return o._statusCode
}

func (o *CreateShareDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file-services/shares][%d] CreateShare default %s", o._statusCode, payload)
}

func (o *CreateShareDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file-services/shares][%d] CreateShare default %s", o._statusCode, payload)
}

func (o *CreateShareDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateShareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

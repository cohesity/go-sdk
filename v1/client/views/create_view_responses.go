// Code generated by go-swagger; DO NOT EDIT.

package views

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v1/models"
)

// CreateViewReader is a Reader for the CreateView structure.
type CreateViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateViewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateViewCreated creates a CreateViewCreated with default headers values
func NewCreateViewCreated() *CreateViewCreated {
	return &CreateViewCreated{}
}

/*
CreateViewCreated describes a response with status code 201, with default header values.

Success
*/
type CreateViewCreated struct {
	Payload *models.View
}

// IsSuccess returns true when this create view created response has a 2xx status code
func (o *CreateViewCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create view created response has a 3xx status code
func (o *CreateViewCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create view created response has a 4xx status code
func (o *CreateViewCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create view created response has a 5xx status code
func (o *CreateViewCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create view created response a status code equal to that given
func (o *CreateViewCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create view created response
func (o *CreateViewCreated) Code() int {
	return 201
}

func (o *CreateViewCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views][%d] createViewCreated %s", 201, payload)
}

func (o *CreateViewCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views][%d] createViewCreated %s", 201, payload)
}

func (o *CreateViewCreated) GetPayload() *models.View {
	return o.Payload
}

func (o *CreateViewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.View)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateViewDefault creates a CreateViewDefault with default headers values
func NewCreateViewDefault(code int) *CreateViewDefault {
	return &CreateViewDefault{
		_statusCode: code,
	}
}

/*
CreateViewDefault describes a response with status code -1, with default header values.

Error
*/
type CreateViewDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create view default response has a 2xx status code
func (o *CreateViewDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create view default response has a 3xx status code
func (o *CreateViewDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create view default response has a 4xx status code
func (o *CreateViewDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create view default response has a 5xx status code
func (o *CreateViewDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create view default response a status code equal to that given
func (o *CreateViewDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create view default response
func (o *CreateViewDefault) Code() int {
	return o._statusCode
}

func (o *CreateViewDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views][%d] CreateView default %s", o._statusCode, payload)
}

func (o *CreateViewDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views][%d] CreateView default %s", o._statusCode, payload)
}

func (o *CreateViewDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package view_boxes

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

// CreateViewBoxReader is a Reader for the CreateViewBox structure.
type CreateViewBoxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateViewBoxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateViewBoxCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateViewBoxDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateViewBoxCreated creates a CreateViewBoxCreated with default headers values
func NewCreateViewBoxCreated() *CreateViewBoxCreated {
	return &CreateViewBoxCreated{}
}

/*
CreateViewBoxCreated describes a response with status code 201, with default header values.

Success
*/
type CreateViewBoxCreated struct {
	Payload *models.ViewBox
}

// IsSuccess returns true when this create view box created response has a 2xx status code
func (o *CreateViewBoxCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create view box created response has a 3xx status code
func (o *CreateViewBoxCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create view box created response has a 4xx status code
func (o *CreateViewBoxCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create view box created response has a 5xx status code
func (o *CreateViewBoxCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create view box created response a status code equal to that given
func (o *CreateViewBoxCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create view box created response
func (o *CreateViewBoxCreated) Code() int {
	return 201
}

func (o *CreateViewBoxCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewBoxes][%d] createViewBoxCreated %s", 201, payload)
}

func (o *CreateViewBoxCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewBoxes][%d] createViewBoxCreated %s", 201, payload)
}

func (o *CreateViewBoxCreated) GetPayload() *models.ViewBox {
	return o.Payload
}

func (o *CreateViewBoxCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ViewBox)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateViewBoxDefault creates a CreateViewBoxDefault with default headers values
func NewCreateViewBoxDefault(code int) *CreateViewBoxDefault {
	return &CreateViewBoxDefault{
		_statusCode: code,
	}
}

/*
CreateViewBoxDefault describes a response with status code -1, with default header values.

Error
*/
type CreateViewBoxDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create view box default response has a 2xx status code
func (o *CreateViewBoxDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create view box default response has a 3xx status code
func (o *CreateViewBoxDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create view box default response has a 4xx status code
func (o *CreateViewBoxDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create view box default response has a 5xx status code
func (o *CreateViewBoxDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create view box default response a status code equal to that given
func (o *CreateViewBoxDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create view box default response
func (o *CreateViewBoxDefault) Code() int {
	return o._statusCode
}

func (o *CreateViewBoxDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewBoxes][%d] CreateViewBox default %s", o._statusCode, payload)
}

func (o *CreateViewBoxDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewBoxes][%d] CreateViewBox default %s", o._statusCode, payload)
}

func (o *CreateViewBoxDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateViewBoxDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

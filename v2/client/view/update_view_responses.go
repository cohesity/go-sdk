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

// UpdateViewReader is a Reader for the UpdateView structure.
type UpdateViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateViewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateViewOK creates a UpdateViewOK with default headers values
func NewUpdateViewOK() *UpdateViewOK {
	return &UpdateViewOK{}
}

/*
UpdateViewOK describes a response with status code 200, with default header values.

Success
*/
type UpdateViewOK struct {
	Payload *models.View
}

// IsSuccess returns true when this update view o k response has a 2xx status code
func (o *UpdateViewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update view o k response has a 3xx status code
func (o *UpdateViewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update view o k response has a 4xx status code
func (o *UpdateViewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update view o k response has a 5xx status code
func (o *UpdateViewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update view o k response a status code equal to that given
func (o *UpdateViewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update view o k response
func (o *UpdateViewOK) Code() int {
	return 200
}

func (o *UpdateViewOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /file-services/views/{id}][%d] updateViewOK %s", 200, payload)
}

func (o *UpdateViewOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /file-services/views/{id}][%d] updateViewOK %s", 200, payload)
}

func (o *UpdateViewOK) GetPayload() *models.View {
	return o.Payload
}

func (o *UpdateViewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.View)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateViewDefault creates a UpdateViewDefault with default headers values
func NewUpdateViewDefault(code int) *UpdateViewDefault {
	return &UpdateViewDefault{
		_statusCode: code,
	}
}

/*
UpdateViewDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateViewDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update view default response has a 2xx status code
func (o *UpdateViewDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update view default response has a 3xx status code
func (o *UpdateViewDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update view default response has a 4xx status code
func (o *UpdateViewDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update view default response has a 5xx status code
func (o *UpdateViewDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update view default response a status code equal to that given
func (o *UpdateViewDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update view default response
func (o *UpdateViewDefault) Code() int {
	return o._statusCode
}

func (o *UpdateViewDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /file-services/views/{id}][%d] UpdateView default %s", o._statusCode, payload)
}

func (o *UpdateViewDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /file-services/views/{id}][%d] UpdateView default %s", o._statusCode, payload)
}

func (o *UpdateViewDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

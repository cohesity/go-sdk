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

// CloneViewReader is a Reader for the CloneView structure.
type CloneViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloneViewCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloneViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloneViewCreated creates a CloneViewCreated with default headers values
func NewCloneViewCreated() *CloneViewCreated {
	return &CloneViewCreated{}
}

/*
CloneViewCreated describes a response with status code 201, with default header values.

Success
*/
type CloneViewCreated struct {
	Payload *models.View
}

// IsSuccess returns true when this clone view created response has a 2xx status code
func (o *CloneViewCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clone view created response has a 3xx status code
func (o *CloneViewCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone view created response has a 4xx status code
func (o *CloneViewCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone view created response has a 5xx status code
func (o *CloneViewCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this clone view created response a status code equal to that given
func (o *CloneViewCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the clone view created response
func (o *CloneViewCreated) Code() int {
	return 201
}

func (o *CloneViewCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/clone][%d] cloneViewCreated %s", 201, payload)
}

func (o *CloneViewCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/clone][%d] cloneViewCreated %s", 201, payload)
}

func (o *CloneViewCreated) GetPayload() *models.View {
	return o.Payload
}

func (o *CloneViewCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.View)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneViewDefault creates a CloneViewDefault with default headers values
func NewCloneViewDefault(code int) *CloneViewDefault {
	return &CloneViewDefault{
		_statusCode: code,
	}
}

/*
CloneViewDefault describes a response with status code -1, with default header values.

Error
*/
type CloneViewDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this clone view default response has a 2xx status code
func (o *CloneViewDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this clone view default response has a 3xx status code
func (o *CloneViewDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this clone view default response has a 4xx status code
func (o *CloneViewDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this clone view default response has a 5xx status code
func (o *CloneViewDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this clone view default response a status code equal to that given
func (o *CloneViewDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the clone view default response
func (o *CloneViewDefault) Code() int {
	return o._statusCode
}

func (o *CloneViewDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/clone][%d] CloneView default %s", o._statusCode, payload)
}

func (o *CloneViewDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/clone][%d] CloneView default %s", o._statusCode, payload)
}

func (o *CloneViewDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CloneViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

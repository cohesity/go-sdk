// Code generated by go-swagger; DO NOT EDIT.

package destroy_clone

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

// DestroyCloneTaskReader is a Reader for the DestroyCloneTask structure.
type DestroyCloneTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyCloneTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDestroyCloneTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDestroyCloneTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDestroyCloneTaskOK creates a DestroyCloneTaskOK with default headers values
func NewDestroyCloneTaskOK() *DestroyCloneTaskOK {
	return &DestroyCloneTaskOK{}
}

/*
DestroyCloneTaskOK describes a response with status code 200, with default header values.

Success
*/
type DestroyCloneTaskOK struct {
	Payload *models.RestoreTaskWrapper
}

// IsSuccess returns true when this destroy clone task o k response has a 2xx status code
func (o *DestroyCloneTaskOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this destroy clone task o k response has a 3xx status code
func (o *DestroyCloneTaskOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this destroy clone task o k response has a 4xx status code
func (o *DestroyCloneTaskOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this destroy clone task o k response has a 5xx status code
func (o *DestroyCloneTaskOK) IsServerError() bool {
	return false
}

// IsCode returns true when this destroy clone task o k response a status code equal to that given
func (o *DestroyCloneTaskOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the destroy clone task o k response
func (o *DestroyCloneTaskOK) Code() int {
	return 200
}

func (o *DestroyCloneTaskOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /destroyclone/{id}][%d] destroyCloneTaskOK %s", 200, payload)
}

func (o *DestroyCloneTaskOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /destroyclone/{id}][%d] destroyCloneTaskOK %s", 200, payload)
}

func (o *DestroyCloneTaskOK) GetPayload() *models.RestoreTaskWrapper {
	return o.Payload
}

func (o *DestroyCloneTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestoreTaskWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDestroyCloneTaskDefault creates a DestroyCloneTaskDefault with default headers values
func NewDestroyCloneTaskDefault(code int) *DestroyCloneTaskDefault {
	return &DestroyCloneTaskDefault{
		_statusCode: code,
	}
}

/*
DestroyCloneTaskDefault describes a response with status code -1, with default header values.

(empty)
*/
type DestroyCloneTaskDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this destroy clone task default response has a 2xx status code
func (o *DestroyCloneTaskDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this destroy clone task default response has a 3xx status code
func (o *DestroyCloneTaskDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this destroy clone task default response has a 4xx status code
func (o *DestroyCloneTaskDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this destroy clone task default response has a 5xx status code
func (o *DestroyCloneTaskDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this destroy clone task default response a status code equal to that given
func (o *DestroyCloneTaskDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the destroy clone task default response
func (o *DestroyCloneTaskDefault) Code() int {
	return o._statusCode
}

func (o *DestroyCloneTaskDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /destroyclone/{id}][%d] DestroyCloneTask default %s", o._statusCode, payload)
}

func (o *DestroyCloneTaskDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /destroyclone/{id}][%d] DestroyCloneTask default %s", o._statusCode, payload)
}

func (o *DestroyCloneTaskDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *DestroyCloneTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

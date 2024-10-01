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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v1/models"
)

// CloneDirectoryReader is a Reader for the CloneDirectory structure.
type CloneDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloneDirectoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCloneDirectoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCloneDirectoryCreated creates a CloneDirectoryCreated with default headers values
func NewCloneDirectoryCreated() *CloneDirectoryCreated {
	return &CloneDirectoryCreated{}
}

/*
CloneDirectoryCreated describes a response with status code 201, with default header values.

Success
*/
type CloneDirectoryCreated struct {
}

// IsSuccess returns true when this clone directory created response has a 2xx status code
func (o *CloneDirectoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clone directory created response has a 3xx status code
func (o *CloneDirectoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone directory created response has a 4xx status code
func (o *CloneDirectoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone directory created response has a 5xx status code
func (o *CloneDirectoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this clone directory created response a status code equal to that given
func (o *CloneDirectoryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the clone directory created response
func (o *CloneDirectoryCreated) Code() int {
	return 201
}

func (o *CloneDirectoryCreated) Error() string {
	return fmt.Sprintf("[POST /public/views/cloneDirectory][%d] cloneDirectoryCreated", 201)
}

func (o *CloneDirectoryCreated) String() string {
	return fmt.Sprintf("[POST /public/views/cloneDirectory][%d] cloneDirectoryCreated", 201)
}

func (o *CloneDirectoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneDirectoryDefault creates a CloneDirectoryDefault with default headers values
func NewCloneDirectoryDefault(code int) *CloneDirectoryDefault {
	return &CloneDirectoryDefault{
		_statusCode: code,
	}
}

/*
CloneDirectoryDefault describes a response with status code -1, with default header values.

Error
*/
type CloneDirectoryDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this clone directory default response has a 2xx status code
func (o *CloneDirectoryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this clone directory default response has a 3xx status code
func (o *CloneDirectoryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this clone directory default response has a 4xx status code
func (o *CloneDirectoryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this clone directory default response has a 5xx status code
func (o *CloneDirectoryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this clone directory default response a status code equal to that given
func (o *CloneDirectoryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the clone directory default response
func (o *CloneDirectoryDefault) Code() int {
	return o._statusCode
}

func (o *CloneDirectoryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/cloneDirectory][%d] CloneDirectory default %s", o._statusCode, payload)
}

func (o *CloneDirectoryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/views/cloneDirectory][%d] CloneDirectory default %s", o._statusCode, payload)
}

func (o *CloneDirectoryDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CloneDirectoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

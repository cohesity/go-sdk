// Code generated by go-swagger; DO NOT EDIT.

package command

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

// LinuxCommandExecuteReader is a Reader for the LinuxCommandExecute structure.
type LinuxCommandExecuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LinuxCommandExecuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLinuxCommandExecuteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLinuxCommandExecuteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLinuxCommandExecuteOK creates a LinuxCommandExecuteOK with default headers values
func NewLinuxCommandExecuteOK() *LinuxCommandExecuteOK {
	return &LinuxCommandExecuteOK{}
}

/*
LinuxCommandExecuteOK describes a response with status code 200, with default header values.

Success
*/
type LinuxCommandExecuteOK struct {
	Payload []*models.LinuxCommandExecuteNodeResp
}

// IsSuccess returns true when this linux command execute o k response has a 2xx status code
func (o *LinuxCommandExecuteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this linux command execute o k response has a 3xx status code
func (o *LinuxCommandExecuteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this linux command execute o k response has a 4xx status code
func (o *LinuxCommandExecuteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this linux command execute o k response has a 5xx status code
func (o *LinuxCommandExecuteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this linux command execute o k response a status code equal to that given
func (o *LinuxCommandExecuteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the linux command execute o k response
func (o *LinuxCommandExecuteOK) Code() int {
	return 200
}

func (o *LinuxCommandExecuteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /linuxCommand/execute][%d] linuxCommandExecuteOK %s", 200, payload)
}

func (o *LinuxCommandExecuteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /linuxCommand/execute][%d] linuxCommandExecuteOK %s", 200, payload)
}

func (o *LinuxCommandExecuteOK) GetPayload() []*models.LinuxCommandExecuteNodeResp {
	return o.Payload
}

func (o *LinuxCommandExecuteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLinuxCommandExecuteDefault creates a LinuxCommandExecuteDefault with default headers values
func NewLinuxCommandExecuteDefault(code int) *LinuxCommandExecuteDefault {
	return &LinuxCommandExecuteDefault{
		_statusCode: code,
	}
}

/*
LinuxCommandExecuteDefault describes a response with status code -1, with default header values.

(empty)
*/
type LinuxCommandExecuteDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this linux command execute default response has a 2xx status code
func (o *LinuxCommandExecuteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this linux command execute default response has a 3xx status code
func (o *LinuxCommandExecuteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this linux command execute default response has a 4xx status code
func (o *LinuxCommandExecuteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this linux command execute default response has a 5xx status code
func (o *LinuxCommandExecuteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this linux command execute default response a status code equal to that given
func (o *LinuxCommandExecuteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the linux command execute default response
func (o *LinuxCommandExecuteDefault) Code() int {
	return o._statusCode
}

func (o *LinuxCommandExecuteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /linuxCommand/execute][%d] LinuxCommandExecute default %s", o._statusCode, payload)
}

func (o *LinuxCommandExecuteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /linuxCommand/execute][%d] LinuxCommandExecute default %s", o._statusCode, payload)
}

func (o *LinuxCommandExecuteDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *LinuxCommandExecuteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

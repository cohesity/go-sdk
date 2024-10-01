// Code generated by go-swagger; DO NOT EDIT.

package app_instance

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

// LaunchAppInstanceReader is a Reader for the LaunchAppInstance structure.
type LaunchAppInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LaunchAppInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewLaunchAppInstanceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLaunchAppInstanceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLaunchAppInstanceAccepted creates a LaunchAppInstanceAccepted with default headers values
func NewLaunchAppInstanceAccepted() *LaunchAppInstanceAccepted {
	return &LaunchAppInstanceAccepted{}
}

/*
LaunchAppInstanceAccepted describes a response with status code 202, with default header values.

Success
*/
type LaunchAppInstanceAccepted struct {
	Payload *models.AppInstanceIDParameter
}

// IsSuccess returns true when this launch app instance accepted response has a 2xx status code
func (o *LaunchAppInstanceAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this launch app instance accepted response has a 3xx status code
func (o *LaunchAppInstanceAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this launch app instance accepted response has a 4xx status code
func (o *LaunchAppInstanceAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this launch app instance accepted response has a 5xx status code
func (o *LaunchAppInstanceAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this launch app instance accepted response a status code equal to that given
func (o *LaunchAppInstanceAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the launch app instance accepted response
func (o *LaunchAppInstanceAccepted) Code() int {
	return 202
}

func (o *LaunchAppInstanceAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/appInstances][%d] launchAppInstanceAccepted %s", 202, payload)
}

func (o *LaunchAppInstanceAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/appInstances][%d] launchAppInstanceAccepted %s", 202, payload)
}

func (o *LaunchAppInstanceAccepted) GetPayload() *models.AppInstanceIDParameter {
	return o.Payload
}

func (o *LaunchAppInstanceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppInstanceIDParameter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLaunchAppInstanceDefault creates a LaunchAppInstanceDefault with default headers values
func NewLaunchAppInstanceDefault(code int) *LaunchAppInstanceDefault {
	return &LaunchAppInstanceDefault{
		_statusCode: code,
	}
}

/*
LaunchAppInstanceDefault describes a response with status code -1, with default header values.

Error
*/
type LaunchAppInstanceDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this launch app instance default response has a 2xx status code
func (o *LaunchAppInstanceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this launch app instance default response has a 3xx status code
func (o *LaunchAppInstanceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this launch app instance default response has a 4xx status code
func (o *LaunchAppInstanceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this launch app instance default response has a 5xx status code
func (o *LaunchAppInstanceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this launch app instance default response a status code equal to that given
func (o *LaunchAppInstanceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the launch app instance default response
func (o *LaunchAppInstanceDefault) Code() int {
	return o._statusCode
}

func (o *LaunchAppInstanceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/appInstances][%d] LaunchAppInstance default %s", o._statusCode, payload)
}

func (o *LaunchAppInstanceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/appInstances][%d] LaunchAppInstance default %s", o._statusCode, payload)
}

func (o *LaunchAppInstanceDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *LaunchAppInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

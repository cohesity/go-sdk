// Code generated by go-swagger; DO NOT EDIT.

package protection_jobs

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

// RunProtectionJobReader is a Reader for the RunProtectionJob structure.
type RunProtectionJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunProtectionJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRunProtectionJobNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRunProtectionJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunProtectionJobNoContent creates a RunProtectionJobNoContent with default headers values
func NewRunProtectionJobNoContent() *RunProtectionJobNoContent {
	return &RunProtectionJobNoContent{}
}

/*
RunProtectionJobNoContent describes a response with status code 204, with default header values.

No Content
*/
type RunProtectionJobNoContent struct {
}

// IsSuccess returns true when this run protection job no content response has a 2xx status code
func (o *RunProtectionJobNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this run protection job no content response has a 3xx status code
func (o *RunProtectionJobNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this run protection job no content response has a 4xx status code
func (o *RunProtectionJobNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this run protection job no content response has a 5xx status code
func (o *RunProtectionJobNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this run protection job no content response a status code equal to that given
func (o *RunProtectionJobNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the run protection job no content response
func (o *RunProtectionJobNoContent) Code() int {
	return 204
}

func (o *RunProtectionJobNoContent) Error() string {
	return fmt.Sprintf("[POST /public/protectionJobs/run/{id}][%d] runProtectionJobNoContent", 204)
}

func (o *RunProtectionJobNoContent) String() string {
	return fmt.Sprintf("[POST /public/protectionJobs/run/{id}][%d] runProtectionJobNoContent", 204)
}

func (o *RunProtectionJobNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRunProtectionJobDefault creates a RunProtectionJobDefault with default headers values
func NewRunProtectionJobDefault(code int) *RunProtectionJobDefault {
	return &RunProtectionJobDefault{
		_statusCode: code,
	}
}

/*
RunProtectionJobDefault describes a response with status code -1, with default header values.

Error
*/
type RunProtectionJobDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this run protection job default response has a 2xx status code
func (o *RunProtectionJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this run protection job default response has a 3xx status code
func (o *RunProtectionJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this run protection job default response has a 4xx status code
func (o *RunProtectionJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this run protection job default response has a 5xx status code
func (o *RunProtectionJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this run protection job default response a status code equal to that given
func (o *RunProtectionJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the run protection job default response
func (o *RunProtectionJobDefault) Code() int {
	return o._statusCode
}

func (o *RunProtectionJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionJobs/run/{id}][%d] RunProtectionJob default %s", o._statusCode, payload)
}

func (o *RunProtectionJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionJobs/run/{id}][%d] RunProtectionJob default %s", o._statusCode, payload)
}

func (o *RunProtectionJobDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *RunProtectionJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

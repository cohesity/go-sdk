// Code generated by go-swagger; DO NOT EDIT.

package protection_runs

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

// CancelProtectionJobRunReader is a Reader for the CancelProtectionJobRun structure.
type CancelProtectionJobRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelProtectionJobRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCancelProtectionJobRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelProtectionJobRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelProtectionJobRunNoContent creates a CancelProtectionJobRunNoContent with default headers values
func NewCancelProtectionJobRunNoContent() *CancelProtectionJobRunNoContent {
	return &CancelProtectionJobRunNoContent{}
}

/*
CancelProtectionJobRunNoContent describes a response with status code 204, with default header values.

No Content
*/
type CancelProtectionJobRunNoContent struct {
}

// IsSuccess returns true when this cancel protection job run no content response has a 2xx status code
func (o *CancelProtectionJobRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel protection job run no content response has a 3xx status code
func (o *CancelProtectionJobRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel protection job run no content response has a 4xx status code
func (o *CancelProtectionJobRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel protection job run no content response has a 5xx status code
func (o *CancelProtectionJobRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel protection job run no content response a status code equal to that given
func (o *CancelProtectionJobRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the cancel protection job run no content response
func (o *CancelProtectionJobRunNoContent) Code() int {
	return 204
}

func (o *CancelProtectionJobRunNoContent) Error() string {
	return fmt.Sprintf("[POST /public/protectionRuns/cancel/{id}][%d] cancelProtectionJobRunNoContent", 204)
}

func (o *CancelProtectionJobRunNoContent) String() string {
	return fmt.Sprintf("[POST /public/protectionRuns/cancel/{id}][%d] cancelProtectionJobRunNoContent", 204)
}

func (o *CancelProtectionJobRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelProtectionJobRunDefault creates a CancelProtectionJobRunDefault with default headers values
func NewCancelProtectionJobRunDefault(code int) *CancelProtectionJobRunDefault {
	return &CancelProtectionJobRunDefault{
		_statusCode: code,
	}
}

/*
CancelProtectionJobRunDefault describes a response with status code -1, with default header values.

Error
*/
type CancelProtectionJobRunDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this cancel protection job run default response has a 2xx status code
func (o *CancelProtectionJobRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cancel protection job run default response has a 3xx status code
func (o *CancelProtectionJobRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cancel protection job run default response has a 4xx status code
func (o *CancelProtectionJobRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cancel protection job run default response has a 5xx status code
func (o *CancelProtectionJobRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cancel protection job run default response a status code equal to that given
func (o *CancelProtectionJobRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cancel protection job run default response
func (o *CancelProtectionJobRunDefault) Code() int {
	return o._statusCode
}

func (o *CancelProtectionJobRunDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionRuns/cancel/{id}][%d] CancelProtectionJobRun default %s", o._statusCode, payload)
}

func (o *CancelProtectionJobRunDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/protectionRuns/cancel/{id}][%d] CancelProtectionJobRun default %s", o._statusCode, payload)
}

func (o *CancelProtectionJobRunDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CancelProtectionJobRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

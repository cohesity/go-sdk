// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// UpdateBackgroundActivityScheduleReader is a Reader for the UpdateBackgroundActivitySchedule structure.
type UpdateBackgroundActivityScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBackgroundActivityScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBackgroundActivityScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateBackgroundActivityScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBackgroundActivityScheduleOK creates a UpdateBackgroundActivityScheduleOK with default headers values
func NewUpdateBackgroundActivityScheduleOK() *UpdateBackgroundActivityScheduleOK {
	return &UpdateBackgroundActivityScheduleOK{}
}

/*
UpdateBackgroundActivityScheduleOK describes a response with status code 200, with default header values.

Success
*/
type UpdateBackgroundActivityScheduleOK struct {
	Payload *models.BandwidthLimit
}

// IsSuccess returns true when this update background activity schedule o k response has a 2xx status code
func (o *UpdateBackgroundActivityScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update background activity schedule o k response has a 3xx status code
func (o *UpdateBackgroundActivityScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update background activity schedule o k response has a 4xx status code
func (o *UpdateBackgroundActivityScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update background activity schedule o k response has a 5xx status code
func (o *UpdateBackgroundActivityScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update background activity schedule o k response a status code equal to that given
func (o *UpdateBackgroundActivityScheduleOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update background activity schedule o k response
func (o *UpdateBackgroundActivityScheduleOK) Code() int {
	return 200
}

func (o *UpdateBackgroundActivityScheduleOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster/backgroundActivitySchedule][%d] updateBackgroundActivityScheduleOK %s", 200, payload)
}

func (o *UpdateBackgroundActivityScheduleOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster/backgroundActivitySchedule][%d] updateBackgroundActivityScheduleOK %s", 200, payload)
}

func (o *UpdateBackgroundActivityScheduleOK) GetPayload() *models.BandwidthLimit {
	return o.Payload
}

func (o *UpdateBackgroundActivityScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BandwidthLimit)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBackgroundActivityScheduleDefault creates a UpdateBackgroundActivityScheduleDefault with default headers values
func NewUpdateBackgroundActivityScheduleDefault(code int) *UpdateBackgroundActivityScheduleDefault {
	return &UpdateBackgroundActivityScheduleDefault{
		_statusCode: code,
	}
}

/*
UpdateBackgroundActivityScheduleDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateBackgroundActivityScheduleDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update background activity schedule default response has a 2xx status code
func (o *UpdateBackgroundActivityScheduleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update background activity schedule default response has a 3xx status code
func (o *UpdateBackgroundActivityScheduleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update background activity schedule default response has a 4xx status code
func (o *UpdateBackgroundActivityScheduleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update background activity schedule default response has a 5xx status code
func (o *UpdateBackgroundActivityScheduleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update background activity schedule default response a status code equal to that given
func (o *UpdateBackgroundActivityScheduleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update background activity schedule default response
func (o *UpdateBackgroundActivityScheduleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBackgroundActivityScheduleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster/backgroundActivitySchedule][%d] UpdateBackgroundActivitySchedule default %s", o._statusCode, payload)
}

func (o *UpdateBackgroundActivityScheduleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster/backgroundActivitySchedule][%d] UpdateBackgroundActivitySchedule default %s", o._statusCode, payload)
}

func (o *UpdateBackgroundActivityScheduleDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateBackgroundActivityScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

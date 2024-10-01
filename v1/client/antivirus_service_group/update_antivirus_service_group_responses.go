// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service_group

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

// UpdateAntivirusServiceGroupReader is a Reader for the UpdateAntivirusServiceGroup structure.
type UpdateAntivirusServiceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAntivirusServiceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAntivirusServiceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateAntivirusServiceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAntivirusServiceGroupOK creates a UpdateAntivirusServiceGroupOK with default headers values
func NewUpdateAntivirusServiceGroupOK() *UpdateAntivirusServiceGroupOK {
	return &UpdateAntivirusServiceGroupOK{}
}

/*
UpdateAntivirusServiceGroupOK describes a response with status code 200, with default header values.

Success
*/
type UpdateAntivirusServiceGroupOK struct {
	Payload *models.AntivirusServiceGroup
}

// IsSuccess returns true when this update antivirus service group o k response has a 2xx status code
func (o *UpdateAntivirusServiceGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update antivirus service group o k response has a 3xx status code
func (o *UpdateAntivirusServiceGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update antivirus service group o k response has a 4xx status code
func (o *UpdateAntivirusServiceGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update antivirus service group o k response has a 5xx status code
func (o *UpdateAntivirusServiceGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update antivirus service group o k response a status code equal to that given
func (o *UpdateAntivirusServiceGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update antivirus service group o k response
func (o *UpdateAntivirusServiceGroupOK) Code() int {
	return 200
}

func (o *UpdateAntivirusServiceGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups][%d] updateAntivirusServiceGroupOK %s", 200, payload)
}

func (o *UpdateAntivirusServiceGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups][%d] updateAntivirusServiceGroupOK %s", 200, payload)
}

func (o *UpdateAntivirusServiceGroupOK) GetPayload() *models.AntivirusServiceGroup {
	return o.Payload
}

func (o *UpdateAntivirusServiceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AntivirusServiceGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAntivirusServiceGroupDefault creates a UpdateAntivirusServiceGroupDefault with default headers values
func NewUpdateAntivirusServiceGroupDefault(code int) *UpdateAntivirusServiceGroupDefault {
	return &UpdateAntivirusServiceGroupDefault{
		_statusCode: code,
	}
}

/*
UpdateAntivirusServiceGroupDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateAntivirusServiceGroupDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update antivirus service group default response has a 2xx status code
func (o *UpdateAntivirusServiceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update antivirus service group default response has a 3xx status code
func (o *UpdateAntivirusServiceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update antivirus service group default response has a 4xx status code
func (o *UpdateAntivirusServiceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update antivirus service group default response has a 5xx status code
func (o *UpdateAntivirusServiceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update antivirus service group default response a status code equal to that given
func (o *UpdateAntivirusServiceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update antivirus service group default response
func (o *UpdateAntivirusServiceGroupDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAntivirusServiceGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups][%d] UpdateAntivirusServiceGroup default %s", o._statusCode, payload)
}

func (o *UpdateAntivirusServiceGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/antivirusGroups][%d] UpdateAntivirusServiceGroup default %s", o._statusCode, payload)
}

func (o *UpdateAntivirusServiceGroupDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateAntivirusServiceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

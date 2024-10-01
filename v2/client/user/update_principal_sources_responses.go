// Code generated by go-swagger; DO NOT EDIT.

package user

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

// UpdatePrincipalSourcesReader is a Reader for the UpdatePrincipalSources structure.
type UpdatePrincipalSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePrincipalSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePrincipalSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdatePrincipalSourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePrincipalSourcesOK creates a UpdatePrincipalSourcesOK with default headers values
func NewUpdatePrincipalSourcesOK() *UpdatePrincipalSourcesOK {
	return &UpdatePrincipalSourcesOK{}
}

/*
UpdatePrincipalSourcesOK describes a response with status code 200, with default header values.

Success
*/
type UpdatePrincipalSourcesOK struct {
	Payload *models.AssignedSources
}

// IsSuccess returns true when this update principal sources o k response has a 2xx status code
func (o *UpdatePrincipalSourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update principal sources o k response has a 3xx status code
func (o *UpdatePrincipalSourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update principal sources o k response has a 4xx status code
func (o *UpdatePrincipalSourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update principal sources o k response has a 5xx status code
func (o *UpdatePrincipalSourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update principal sources o k response a status code equal to that given
func (o *UpdatePrincipalSourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update principal sources o k response
func (o *UpdatePrincipalSourcesOK) Code() int {
	return 200
}

func (o *UpdatePrincipalSourcesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /security-principals/{sid}/sources][%d] updatePrincipalSourcesOK %s", 200, payload)
}

func (o *UpdatePrincipalSourcesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /security-principals/{sid}/sources][%d] updatePrincipalSourcesOK %s", 200, payload)
}

func (o *UpdatePrincipalSourcesOK) GetPayload() *models.AssignedSources {
	return o.Payload
}

func (o *UpdatePrincipalSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignedSources)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePrincipalSourcesDefault creates a UpdatePrincipalSourcesDefault with default headers values
func NewUpdatePrincipalSourcesDefault(code int) *UpdatePrincipalSourcesDefault {
	return &UpdatePrincipalSourcesDefault{
		_statusCode: code,
	}
}

/*
UpdatePrincipalSourcesDefault describes a response with status code -1, with default header values.

Error
*/
type UpdatePrincipalSourcesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update principal sources default response has a 2xx status code
func (o *UpdatePrincipalSourcesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update principal sources default response has a 3xx status code
func (o *UpdatePrincipalSourcesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update principal sources default response has a 4xx status code
func (o *UpdatePrincipalSourcesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update principal sources default response has a 5xx status code
func (o *UpdatePrincipalSourcesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update principal sources default response a status code equal to that given
func (o *UpdatePrincipalSourcesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update principal sources default response
func (o *UpdatePrincipalSourcesDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePrincipalSourcesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /security-principals/{sid}/sources][%d] UpdatePrincipalSources default %s", o._statusCode, payload)
}

func (o *UpdatePrincipalSourcesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /security-principals/{sid}/sources][%d] UpdatePrincipalSources default %s", o._statusCode, payload)
}

func (o *UpdatePrincipalSourcesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePrincipalSourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

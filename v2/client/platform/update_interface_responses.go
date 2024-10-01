// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// UpdateInterfaceReader is a Reader for the UpdateInterface structure.
type UpdateInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateInterfaceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateInterfaceOK creates a UpdateInterfaceOK with default headers values
func NewUpdateInterfaceOK() *UpdateInterfaceOK {
	return &UpdateInterfaceOK{}
}

/*
UpdateInterfaceOK describes a response with status code 200, with default header values.

Success
*/
type UpdateInterfaceOK struct {
	Payload *models.InterfaceParams
}

// IsSuccess returns true when this update interface o k response has a 2xx status code
func (o *UpdateInterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update interface o k response has a 3xx status code
func (o *UpdateInterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update interface o k response has a 4xx status code
func (o *UpdateInterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update interface o k response has a 5xx status code
func (o *UpdateInterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update interface o k response a status code equal to that given
func (o *UpdateInterfaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update interface o k response
func (o *UpdateInterfaceOK) Code() int {
	return 200
}

func (o *UpdateInterfaceOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/interfaces/{id}][%d] updateInterfaceOK %s", 200, payload)
}

func (o *UpdateInterfaceOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/interfaces/{id}][%d] updateInterfaceOK %s", 200, payload)
}

func (o *UpdateInterfaceOK) GetPayload() *models.InterfaceParams {
	return o.Payload
}

func (o *UpdateInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateInterfaceDefault creates a UpdateInterfaceDefault with default headers values
func NewUpdateInterfaceDefault(code int) *UpdateInterfaceDefault {
	return &UpdateInterfaceDefault{
		_statusCode: code,
	}
}

/*
UpdateInterfaceDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateInterfaceDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update interface default response has a 2xx status code
func (o *UpdateInterfaceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update interface default response has a 3xx status code
func (o *UpdateInterfaceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update interface default response has a 4xx status code
func (o *UpdateInterfaceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update interface default response has a 5xx status code
func (o *UpdateInterfaceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update interface default response a status code equal to that given
func (o *UpdateInterfaceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update interface default response
func (o *UpdateInterfaceDefault) Code() int {
	return o._statusCode
}

func (o *UpdateInterfaceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/interfaces/{id}][%d] UpdateInterface default %s", o._statusCode, payload)
}

func (o *UpdateInterfaceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /network/interfaces/{id}][%d] UpdateInterface default %s", o._statusCode, payload)
}

func (o *UpdateInterfaceDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateInterfaceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

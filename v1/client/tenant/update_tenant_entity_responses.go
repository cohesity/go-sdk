// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v1/models"
)

// UpdateTenantEntityReader is a Reader for the UpdateTenantEntity structure.
type UpdateTenantEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTenantEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTenantEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateTenantEntityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTenantEntityOK creates a UpdateTenantEntityOK with default headers values
func NewUpdateTenantEntityOK() *UpdateTenantEntityOK {
	return &UpdateTenantEntityOK{}
}

/*
UpdateTenantEntityOK describes a response with status code 200, with default header values.

Tenant Entity Update Response.
*/
type UpdateTenantEntityOK struct {
	Payload *models.TenantEntityUpdate
}

// IsSuccess returns true when this update tenant entity o k response has a 2xx status code
func (o *UpdateTenantEntityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update tenant entity o k response has a 3xx status code
func (o *UpdateTenantEntityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tenant entity o k response has a 4xx status code
func (o *UpdateTenantEntityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update tenant entity o k response has a 5xx status code
func (o *UpdateTenantEntityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update tenant entity o k response a status code equal to that given
func (o *UpdateTenantEntityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update tenant entity o k response
func (o *UpdateTenantEntityOK) Code() int {
	return 200
}

func (o *UpdateTenantEntityOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/entity][%d] updateTenantEntityOK %s", 200, payload)
}

func (o *UpdateTenantEntityOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/entity][%d] updateTenantEntityOK %s", 200, payload)
}

func (o *UpdateTenantEntityOK) GetPayload() *models.TenantEntityUpdate {
	return o.Payload
}

func (o *UpdateTenantEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantEntityUpdate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTenantEntityDefault creates a UpdateTenantEntityDefault with default headers values
func NewUpdateTenantEntityDefault(code int) *UpdateTenantEntityDefault {
	return &UpdateTenantEntityDefault{
		_statusCode: code,
	}
}

/*
UpdateTenantEntityDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateTenantEntityDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update tenant entity default response has a 2xx status code
func (o *UpdateTenantEntityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update tenant entity default response has a 3xx status code
func (o *UpdateTenantEntityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update tenant entity default response has a 4xx status code
func (o *UpdateTenantEntityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update tenant entity default response has a 5xx status code
func (o *UpdateTenantEntityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update tenant entity default response a status code equal to that given
func (o *UpdateTenantEntityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update tenant entity default response
func (o *UpdateTenantEntityDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTenantEntityDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/entity][%d] UpdateTenantEntity default %s", o._statusCode, payload)
}

func (o *UpdateTenantEntityDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/entity][%d] UpdateTenantEntity default %s", o._statusCode, payload)
}

func (o *UpdateTenantEntityDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateTenantEntityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// CreateViewUserQuotaReader is a Reader for the CreateViewUserQuota structure.
type CreateViewUserQuotaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateViewUserQuotaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateViewUserQuotaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateViewUserQuotaDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateViewUserQuotaOK creates a CreateViewUserQuotaOK with default headers values
func NewCreateViewUserQuotaOK() *CreateViewUserQuotaOK {
	return &CreateViewUserQuotaOK{}
}

/*
CreateViewUserQuotaOK describes a response with status code 200, with default header values.

Success
*/
type CreateViewUserQuotaOK struct {
	Payload *models.UserQuotaAndUsage
}

// IsSuccess returns true when this create view user quota o k response has a 2xx status code
func (o *CreateViewUserQuotaOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create view user quota o k response has a 3xx status code
func (o *CreateViewUserQuotaOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create view user quota o k response has a 4xx status code
func (o *CreateViewUserQuotaOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create view user quota o k response has a 5xx status code
func (o *CreateViewUserQuotaOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create view user quota o k response a status code equal to that given
func (o *CreateViewUserQuotaOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create view user quota o k response
func (o *CreateViewUserQuotaOK) Code() int {
	return 200
}

func (o *CreateViewUserQuotaOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewUserQuotas][%d] createViewUserQuotaOK %s", 200, payload)
}

func (o *CreateViewUserQuotaOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewUserQuotas][%d] createViewUserQuotaOK %s", 200, payload)
}

func (o *CreateViewUserQuotaOK) GetPayload() *models.UserQuotaAndUsage {
	return o.Payload
}

func (o *CreateViewUserQuotaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQuotaAndUsage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateViewUserQuotaDefault creates a CreateViewUserQuotaDefault with default headers values
func NewCreateViewUserQuotaDefault(code int) *CreateViewUserQuotaDefault {
	return &CreateViewUserQuotaDefault{
		_statusCode: code,
	}
}

/*
CreateViewUserQuotaDefault describes a response with status code -1, with default header values.

Error
*/
type CreateViewUserQuotaDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create view user quota default response has a 2xx status code
func (o *CreateViewUserQuotaDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create view user quota default response has a 3xx status code
func (o *CreateViewUserQuotaDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create view user quota default response has a 4xx status code
func (o *CreateViewUserQuotaDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create view user quota default response has a 5xx status code
func (o *CreateViewUserQuotaDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create view user quota default response a status code equal to that given
func (o *CreateViewUserQuotaDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create view user quota default response
func (o *CreateViewUserQuotaDefault) Code() int {
	return o._statusCode
}

func (o *CreateViewUserQuotaDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewUserQuotas][%d] CreateViewUserQuota default %s", o._statusCode, payload)
}

func (o *CreateViewUserQuotaDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/viewUserQuotas][%d] CreateViewUserQuota default %s", o._statusCode, payload)
}

func (o *CreateViewUserQuotaDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateViewUserQuotaDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

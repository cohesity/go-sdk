// Code generated by go-swagger; DO NOT EDIT.

package recovery

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

// CreateRecoveryReader is a Reader for the CreateRecovery structure.
type CreateRecoveryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRecoveryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRecoveryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateRecoveryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRecoveryCreated creates a CreateRecoveryCreated with default headers values
func NewCreateRecoveryCreated() *CreateRecoveryCreated {
	return &CreateRecoveryCreated{}
}

/*
CreateRecoveryCreated describes a response with status code 201, with default header values.

Success
*/
type CreateRecoveryCreated struct {
	Payload *models.Recovery
}

// IsSuccess returns true when this create recovery created response has a 2xx status code
func (o *CreateRecoveryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create recovery created response has a 3xx status code
func (o *CreateRecoveryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create recovery created response has a 4xx status code
func (o *CreateRecoveryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create recovery created response has a 5xx status code
func (o *CreateRecoveryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create recovery created response a status code equal to that given
func (o *CreateRecoveryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create recovery created response
func (o *CreateRecoveryCreated) Code() int {
	return 201
}

func (o *CreateRecoveryCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries][%d] createRecoveryCreated %s", 201, payload)
}

func (o *CreateRecoveryCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries][%d] createRecoveryCreated %s", 201, payload)
}

func (o *CreateRecoveryCreated) GetPayload() *models.Recovery {
	return o.Payload
}

func (o *CreateRecoveryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recovery)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecoveryDefault creates a CreateRecoveryDefault with default headers values
func NewCreateRecoveryDefault(code int) *CreateRecoveryDefault {
	return &CreateRecoveryDefault{
		_statusCode: code,
	}
}

/*
CreateRecoveryDefault describes a response with status code -1, with default header values.

Error
*/
type CreateRecoveryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create recovery default response has a 2xx status code
func (o *CreateRecoveryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create recovery default response has a 3xx status code
func (o *CreateRecoveryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create recovery default response has a 4xx status code
func (o *CreateRecoveryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create recovery default response has a 5xx status code
func (o *CreateRecoveryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create recovery default response a status code equal to that given
func (o *CreateRecoveryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create recovery default response
func (o *CreateRecoveryDefault) Code() int {
	return o._statusCode
}

func (o *CreateRecoveryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries][%d] CreateRecovery default %s", o._statusCode, payload)
}

func (o *CreateRecoveryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/recoveries][%d] CreateRecovery default %s", o._statusCode, payload)
}

func (o *CreateRecoveryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateRecoveryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

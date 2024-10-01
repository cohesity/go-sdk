// Code generated by go-swagger; DO NOT EDIT.

package security

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

// ModifyCiphersReader is a Reader for the ModifyCiphers structure.
type ModifyCiphersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyCiphersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyCiphersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModifyCiphersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModifyCiphersOK creates a ModifyCiphersOK with default headers values
func NewModifyCiphersOK() *ModifyCiphersOK {
	return &ModifyCiphersOK{}
}

/*
ModifyCiphersOK describes a response with status code 200, with default header values.

Success
*/
type ModifyCiphersOK struct {
	Payload *models.CiphersResp
}

// IsSuccess returns true when this modify ciphers o k response has a 2xx status code
func (o *ModifyCiphersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this modify ciphers o k response has a 3xx status code
func (o *ModifyCiphersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this modify ciphers o k response has a 4xx status code
func (o *ModifyCiphersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this modify ciphers o k response has a 5xx status code
func (o *ModifyCiphersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this modify ciphers o k response a status code equal to that given
func (o *ModifyCiphersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the modify ciphers o k response
func (o *ModifyCiphersOK) Code() int {
	return 200
}

func (o *ModifyCiphersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/ciphers][%d] modifyCiphersOK %s", 200, payload)
}

func (o *ModifyCiphersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/ciphers][%d] modifyCiphersOK %s", 200, payload)
}

func (o *ModifyCiphersOK) GetPayload() *models.CiphersResp {
	return o.Payload
}

func (o *ModifyCiphersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CiphersResp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModifyCiphersDefault creates a ModifyCiphersDefault with default headers values
func NewModifyCiphersDefault(code int) *ModifyCiphersDefault {
	return &ModifyCiphersDefault{
		_statusCode: code,
	}
}

/*
ModifyCiphersDefault describes a response with status code -1, with default header values.

Error
*/
type ModifyCiphersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this modify ciphers default response has a 2xx status code
func (o *ModifyCiphersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this modify ciphers default response has a 3xx status code
func (o *ModifyCiphersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this modify ciphers default response has a 4xx status code
func (o *ModifyCiphersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this modify ciphers default response has a 5xx status code
func (o *ModifyCiphersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this modify ciphers default response a status code equal to that given
func (o *ModifyCiphersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the modify ciphers default response
func (o *ModifyCiphersDefault) Code() int {
	return o._statusCode
}

func (o *ModifyCiphersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/ciphers][%d] ModifyCiphers default %s", o._statusCode, payload)
}

func (o *ModifyCiphersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /security/ciphers][%d] ModifyCiphers default %s", o._statusCode, payload)
}

func (o *ModifyCiphersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModifyCiphersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

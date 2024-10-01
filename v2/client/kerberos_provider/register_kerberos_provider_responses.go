// Code generated by go-swagger; DO NOT EDIT.

package kerberos_provider

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

// RegisterKerberosProviderReader is a Reader for the RegisterKerberosProvider structure.
type RegisterKerberosProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterKerberosProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterKerberosProviderCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterKerberosProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterKerberosProviderCreated creates a RegisterKerberosProviderCreated with default headers values
func NewRegisterKerberosProviderCreated() *RegisterKerberosProviderCreated {
	return &RegisterKerberosProviderCreated{}
}

/*
RegisterKerberosProviderCreated describes a response with status code 201, with default header values.

Success
*/
type RegisterKerberosProviderCreated struct {
	Payload *models.KerberosProvider
}

// IsSuccess returns true when this register kerberos provider created response has a 2xx status code
func (o *RegisterKerberosProviderCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register kerberos provider created response has a 3xx status code
func (o *RegisterKerberosProviderCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register kerberos provider created response has a 4xx status code
func (o *RegisterKerberosProviderCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this register kerberos provider created response has a 5xx status code
func (o *RegisterKerberosProviderCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this register kerberos provider created response a status code equal to that given
func (o *RegisterKerberosProviderCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the register kerberos provider created response
func (o *RegisterKerberosProviderCreated) Code() int {
	return 201
}

func (o *RegisterKerberosProviderCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kerberos-providers/register][%d] registerKerberosProviderCreated %s", 201, payload)
}

func (o *RegisterKerberosProviderCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kerberos-providers/register][%d] registerKerberosProviderCreated %s", 201, payload)
}

func (o *RegisterKerberosProviderCreated) GetPayload() *models.KerberosProvider {
	return o.Payload
}

func (o *RegisterKerberosProviderCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterKerberosProviderDefault creates a RegisterKerberosProviderDefault with default headers values
func NewRegisterKerberosProviderDefault(code int) *RegisterKerberosProviderDefault {
	return &RegisterKerberosProviderDefault{
		_statusCode: code,
	}
}

/*
RegisterKerberosProviderDefault describes a response with status code -1, with default header values.

Error
*/
type RegisterKerberosProviderDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this register kerberos provider default response has a 2xx status code
func (o *RegisterKerberosProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this register kerberos provider default response has a 3xx status code
func (o *RegisterKerberosProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this register kerberos provider default response has a 4xx status code
func (o *RegisterKerberosProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this register kerberos provider default response has a 5xx status code
func (o *RegisterKerberosProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this register kerberos provider default response a status code equal to that given
func (o *RegisterKerberosProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the register kerberos provider default response
func (o *RegisterKerberosProviderDefault) Code() int {
	return o._statusCode
}

func (o *RegisterKerberosProviderDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kerberos-providers/register][%d] RegisterKerberosProvider default %s", o._statusCode, payload)
}

func (o *RegisterKerberosProviderDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /kerberos-providers/register][%d] RegisterKerberosProvider default %s", o._statusCode, payload)
}

func (o *RegisterKerberosProviderDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterKerberosProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

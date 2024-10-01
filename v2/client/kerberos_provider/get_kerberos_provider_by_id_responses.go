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

// GetKerberosProviderByIDReader is a Reader for the GetKerberosProviderByID structure.
type GetKerberosProviderByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKerberosProviderByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKerberosProviderByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKerberosProviderByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKerberosProviderByIDOK creates a GetKerberosProviderByIDOK with default headers values
func NewGetKerberosProviderByIDOK() *GetKerberosProviderByIDOK {
	return &GetKerberosProviderByIDOK{}
}

/*
GetKerberosProviderByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetKerberosProviderByIDOK struct {
	Payload *models.KerberosProvider
}

// IsSuccess returns true when this get kerberos provider by Id o k response has a 2xx status code
func (o *GetKerberosProviderByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kerberos provider by Id o k response has a 3xx status code
func (o *GetKerberosProviderByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kerberos provider by Id o k response has a 4xx status code
func (o *GetKerberosProviderByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kerberos provider by Id o k response has a 5xx status code
func (o *GetKerberosProviderByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kerberos provider by Id o k response a status code equal to that given
func (o *GetKerberosProviderByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kerberos provider by Id o k response
func (o *GetKerberosProviderByIDOK) Code() int {
	return 200
}

func (o *GetKerberosProviderByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kerberos-providers/{id}][%d] getKerberosProviderByIdOK %s", 200, payload)
}

func (o *GetKerberosProviderByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kerberos-providers/{id}][%d] getKerberosProviderByIdOK %s", 200, payload)
}

func (o *GetKerberosProviderByIDOK) GetPayload() *models.KerberosProvider {
	return o.Payload
}

func (o *GetKerberosProviderByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KerberosProvider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKerberosProviderByIDDefault creates a GetKerberosProviderByIDDefault with default headers values
func NewGetKerberosProviderByIDDefault(code int) *GetKerberosProviderByIDDefault {
	return &GetKerberosProviderByIDDefault{
		_statusCode: code,
	}
}

/*
GetKerberosProviderByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetKerberosProviderByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get kerberos provider by Id default response has a 2xx status code
func (o *GetKerberosProviderByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get kerberos provider by Id default response has a 3xx status code
func (o *GetKerberosProviderByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get kerberos provider by Id default response has a 4xx status code
func (o *GetKerberosProviderByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get kerberos provider by Id default response has a 5xx status code
func (o *GetKerberosProviderByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get kerberos provider by Id default response a status code equal to that given
func (o *GetKerberosProviderByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get kerberos provider by Id default response
func (o *GetKerberosProviderByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetKerberosProviderByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kerberos-providers/{id}][%d] GetKerberosProviderById default %s", o._statusCode, payload)
}

func (o *GetKerberosProviderByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kerberos-providers/{id}][%d] GetKerberosProviderById default %s", o._statusCode, payload)
}

func (o *GetKerberosProviderByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKerberosProviderByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

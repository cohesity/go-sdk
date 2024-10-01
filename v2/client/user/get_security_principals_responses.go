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

	"github.com/cohesity/go-sdk/v2/models"
)

// GetSecurityPrincipalsReader is a Reader for the GetSecurityPrincipals structure.
type GetSecurityPrincipalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityPrincipalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityPrincipalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSecurityPrincipalsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSecurityPrincipalsOK creates a GetSecurityPrincipalsOK with default headers values
func NewGetSecurityPrincipalsOK() *GetSecurityPrincipalsOK {
	return &GetSecurityPrincipalsOK{}
}

/*
GetSecurityPrincipalsOK describes a response with status code 200, with default header values.

Success
*/
type GetSecurityPrincipalsOK struct {
	Payload *models.SecurityPrincipals
}

// IsSuccess returns true when this get security principals o k response has a 2xx status code
func (o *GetSecurityPrincipalsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get security principals o k response has a 3xx status code
func (o *GetSecurityPrincipalsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get security principals o k response has a 4xx status code
func (o *GetSecurityPrincipalsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get security principals o k response has a 5xx status code
func (o *GetSecurityPrincipalsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get security principals o k response a status code equal to that given
func (o *GetSecurityPrincipalsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get security principals o k response
func (o *GetSecurityPrincipalsOK) Code() int {
	return 200
}

func (o *GetSecurityPrincipalsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security-principals][%d] getSecurityPrincipalsOK %s", 200, payload)
}

func (o *GetSecurityPrincipalsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security-principals][%d] getSecurityPrincipalsOK %s", 200, payload)
}

func (o *GetSecurityPrincipalsOK) GetPayload() *models.SecurityPrincipals {
	return o.Payload
}

func (o *GetSecurityPrincipalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityPrincipals)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityPrincipalsDefault creates a GetSecurityPrincipalsDefault with default headers values
func NewGetSecurityPrincipalsDefault(code int) *GetSecurityPrincipalsDefault {
	return &GetSecurityPrincipalsDefault{
		_statusCode: code,
	}
}

/*
GetSecurityPrincipalsDefault describes a response with status code -1, with default header values.

Error
*/
type GetSecurityPrincipalsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get security principals default response has a 2xx status code
func (o *GetSecurityPrincipalsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get security principals default response has a 3xx status code
func (o *GetSecurityPrincipalsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get security principals default response has a 4xx status code
func (o *GetSecurityPrincipalsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get security principals default response has a 5xx status code
func (o *GetSecurityPrincipalsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get security principals default response a status code equal to that given
func (o *GetSecurityPrincipalsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get security principals default response
func (o *GetSecurityPrincipalsDefault) Code() int {
	return o._statusCode
}

func (o *GetSecurityPrincipalsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security-principals][%d] GetSecurityPrincipals default %s", o._statusCode, payload)
}

func (o *GetSecurityPrincipalsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /security-principals][%d] GetSecurityPrincipals default %s", o._statusCode, payload)
}

func (o *GetSecurityPrincipalsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetSecurityPrincipalsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// GetExternalClientSubnetsReader is a Reader for the GetExternalClientSubnets structure.
type GetExternalClientSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalClientSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalClientSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetExternalClientSubnetsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExternalClientSubnetsOK creates a GetExternalClientSubnetsOK with default headers values
func NewGetExternalClientSubnetsOK() *GetExternalClientSubnetsOK {
	return &GetExternalClientSubnetsOK{}
}

/*
GetExternalClientSubnetsOK describes a response with status code 200, with default header values.

Success
*/
type GetExternalClientSubnetsOK struct {
	Payload *models.ExternalClientSubnets
}

// IsSuccess returns true when this get external client subnets o k response has a 2xx status code
func (o *GetExternalClientSubnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get external client subnets o k response has a 3xx status code
func (o *GetExternalClientSubnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get external client subnets o k response has a 4xx status code
func (o *GetExternalClientSubnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get external client subnets o k response has a 5xx status code
func (o *GetExternalClientSubnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get external client subnets o k response a status code equal to that given
func (o *GetExternalClientSubnetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get external client subnets o k response
func (o *GetExternalClientSubnetsOK) Code() int {
	return 200
}

func (o *GetExternalClientSubnetsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/externalClientSubnets][%d] getExternalClientSubnetsOK %s", 200, payload)
}

func (o *GetExternalClientSubnetsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/externalClientSubnets][%d] getExternalClientSubnetsOK %s", 200, payload)
}

func (o *GetExternalClientSubnetsOK) GetPayload() *models.ExternalClientSubnets {
	return o.Payload
}

func (o *GetExternalClientSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalClientSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalClientSubnetsDefault creates a GetExternalClientSubnetsDefault with default headers values
func NewGetExternalClientSubnetsDefault(code int) *GetExternalClientSubnetsDefault {
	return &GetExternalClientSubnetsDefault{
		_statusCode: code,
	}
}

/*
GetExternalClientSubnetsDefault describes a response with status code -1, with default header values.

Error
*/
type GetExternalClientSubnetsDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get external client subnets default response has a 2xx status code
func (o *GetExternalClientSubnetsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get external client subnets default response has a 3xx status code
func (o *GetExternalClientSubnetsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get external client subnets default response has a 4xx status code
func (o *GetExternalClientSubnetsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get external client subnets default response has a 5xx status code
func (o *GetExternalClientSubnetsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get external client subnets default response a status code equal to that given
func (o *GetExternalClientSubnetsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get external client subnets default response
func (o *GetExternalClientSubnetsDefault) Code() int {
	return o._statusCode
}

func (o *GetExternalClientSubnetsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/externalClientSubnets][%d] GetExternalClientSubnets default %s", o._statusCode, payload)
}

func (o *GetExternalClientSubnetsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/externalClientSubnets][%d] GetExternalClientSubnets default %s", o._statusCode, payload)
}

func (o *GetExternalClientSubnetsDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetExternalClientSubnetsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package key_management_system

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

// GetKmsConfigurationsReader is a Reader for the GetKmsConfigurations structure.
type GetKmsConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKmsConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKmsConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKmsConfigurationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKmsConfigurationsOK creates a GetKmsConfigurationsOK with default headers values
func NewGetKmsConfigurationsOK() *GetKmsConfigurationsOK {
	return &GetKmsConfigurationsOK{}
}

/*
GetKmsConfigurationsOK describes a response with status code 200, with default header values.

Success
*/
type GetKmsConfigurationsOK struct {
	Payload *models.KmsConfigurations
}

// IsSuccess returns true when this get kms configurations o k response has a 2xx status code
func (o *GetKmsConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get kms configurations o k response has a 3xx status code
func (o *GetKmsConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get kms configurations o k response has a 4xx status code
func (o *GetKmsConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get kms configurations o k response has a 5xx status code
func (o *GetKmsConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get kms configurations o k response a status code equal to that given
func (o *GetKmsConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get kms configurations o k response
func (o *GetKmsConfigurationsOK) Code() int {
	return 200
}

func (o *GetKmsConfigurationsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kms][%d] getKmsConfigurationsOK %s", 200, payload)
}

func (o *GetKmsConfigurationsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kms][%d] getKmsConfigurationsOK %s", 200, payload)
}

func (o *GetKmsConfigurationsOK) GetPayload() *models.KmsConfigurations {
	return o.Payload
}

func (o *GetKmsConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KmsConfigurations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKmsConfigurationsDefault creates a GetKmsConfigurationsDefault with default headers values
func NewGetKmsConfigurationsDefault(code int) *GetKmsConfigurationsDefault {
	return &GetKmsConfigurationsDefault{
		_statusCode: code,
	}
}

/*
GetKmsConfigurationsDefault describes a response with status code -1, with default header values.

Error
*/
type GetKmsConfigurationsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get kms configurations default response has a 2xx status code
func (o *GetKmsConfigurationsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get kms configurations default response has a 3xx status code
func (o *GetKmsConfigurationsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get kms configurations default response has a 4xx status code
func (o *GetKmsConfigurationsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get kms configurations default response has a 5xx status code
func (o *GetKmsConfigurationsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get kms configurations default response a status code equal to that given
func (o *GetKmsConfigurationsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get kms configurations default response
func (o *GetKmsConfigurationsDefault) Code() int {
	return o._statusCode
}

func (o *GetKmsConfigurationsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kms][%d] GetKmsConfigurations default %s", o._statusCode, payload)
}

func (o *GetKmsConfigurationsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /kms][%d] GetKmsConfigurations default %s", o._statusCode, payload)
}

func (o *GetKmsConfigurationsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetKmsConfigurationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

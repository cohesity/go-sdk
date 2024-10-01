// Code generated by go-swagger; DO NOT EDIT.

package protection_policies

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

// GetProtectionPolicyByIDReader is a Reader for the GetProtectionPolicyByID structure.
type GetProtectionPolicyByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProtectionPolicyByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProtectionPolicyByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProtectionPolicyByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProtectionPolicyByIDOK creates a GetProtectionPolicyByIDOK with default headers values
func NewGetProtectionPolicyByIDOK() *GetProtectionPolicyByIDOK {
	return &GetProtectionPolicyByIDOK{}
}

/*
GetProtectionPolicyByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetProtectionPolicyByIDOK struct {
	Payload *models.ProtectionPolicy
}

// IsSuccess returns true when this get protection policy by Id o k response has a 2xx status code
func (o *GetProtectionPolicyByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get protection policy by Id o k response has a 3xx status code
func (o *GetProtectionPolicyByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get protection policy by Id o k response has a 4xx status code
func (o *GetProtectionPolicyByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get protection policy by Id o k response has a 5xx status code
func (o *GetProtectionPolicyByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get protection policy by Id o k response a status code equal to that given
func (o *GetProtectionPolicyByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get protection policy by Id o k response
func (o *GetProtectionPolicyByIDOK) Code() int {
	return 200
}

func (o *GetProtectionPolicyByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies/{id}][%d] getProtectionPolicyByIdOK %s", 200, payload)
}

func (o *GetProtectionPolicyByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies/{id}][%d] getProtectionPolicyByIdOK %s", 200, payload)
}

func (o *GetProtectionPolicyByIDOK) GetPayload() *models.ProtectionPolicy {
	return o.Payload
}

func (o *GetProtectionPolicyByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProtectionPolicyByIDDefault creates a GetProtectionPolicyByIDDefault with default headers values
func NewGetProtectionPolicyByIDDefault(code int) *GetProtectionPolicyByIDDefault {
	return &GetProtectionPolicyByIDDefault{
		_statusCode: code,
	}
}

/*
GetProtectionPolicyByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetProtectionPolicyByIDDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get protection policy by Id default response has a 2xx status code
func (o *GetProtectionPolicyByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get protection policy by Id default response has a 3xx status code
func (o *GetProtectionPolicyByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get protection policy by Id default response has a 4xx status code
func (o *GetProtectionPolicyByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get protection policy by Id default response has a 5xx status code
func (o *GetProtectionPolicyByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get protection policy by Id default response a status code equal to that given
func (o *GetProtectionPolicyByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get protection policy by Id default response
func (o *GetProtectionPolicyByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProtectionPolicyByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies/{id}][%d] GetProtectionPolicyById default %s", o._statusCode, payload)
}

func (o *GetProtectionPolicyByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies/{id}][%d] GetProtectionPolicyById default %s", o._statusCode, payload)
}

func (o *GetProtectionPolicyByIDDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetProtectionPolicyByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetProtectionPoliciesReader is a Reader for the GetProtectionPolicies structure.
type GetProtectionPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProtectionPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProtectionPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProtectionPoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProtectionPoliciesOK creates a GetProtectionPoliciesOK with default headers values
func NewGetProtectionPoliciesOK() *GetProtectionPoliciesOK {
	return &GetProtectionPoliciesOK{}
}

/*
GetProtectionPoliciesOK describes a response with status code 200, with default header values.

Success
*/
type GetProtectionPoliciesOK struct {
	Payload []*models.ProtectionPolicy
}

// IsSuccess returns true when this get protection policies o k response has a 2xx status code
func (o *GetProtectionPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get protection policies o k response has a 3xx status code
func (o *GetProtectionPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get protection policies o k response has a 4xx status code
func (o *GetProtectionPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get protection policies o k response has a 5xx status code
func (o *GetProtectionPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get protection policies o k response a status code equal to that given
func (o *GetProtectionPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get protection policies o k response
func (o *GetProtectionPoliciesOK) Code() int {
	return 200
}

func (o *GetProtectionPoliciesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies][%d] getProtectionPoliciesOK %s", 200, payload)
}

func (o *GetProtectionPoliciesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies][%d] getProtectionPoliciesOK %s", 200, payload)
}

func (o *GetProtectionPoliciesOK) GetPayload() []*models.ProtectionPolicy {
	return o.Payload
}

func (o *GetProtectionPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProtectionPoliciesDefault creates a GetProtectionPoliciesDefault with default headers values
func NewGetProtectionPoliciesDefault(code int) *GetProtectionPoliciesDefault {
	return &GetProtectionPoliciesDefault{
		_statusCode: code,
	}
}

/*
GetProtectionPoliciesDefault describes a response with status code -1, with default header values.

Error
*/
type GetProtectionPoliciesDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get protection policies default response has a 2xx status code
func (o *GetProtectionPoliciesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get protection policies default response has a 3xx status code
func (o *GetProtectionPoliciesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get protection policies default response has a 4xx status code
func (o *GetProtectionPoliciesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get protection policies default response has a 5xx status code
func (o *GetProtectionPoliciesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get protection policies default response a status code equal to that given
func (o *GetProtectionPoliciesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get protection policies default response
func (o *GetProtectionPoliciesDefault) Code() int {
	return o._statusCode
}

func (o *GetProtectionPoliciesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies][%d] GetProtectionPolicies default %s", o._statusCode, payload)
}

func (o *GetProtectionPoliciesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionPolicies][%d] GetProtectionPolicies default %s", o._statusCode, payload)
}

func (o *GetProtectionPoliciesDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetProtectionPoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

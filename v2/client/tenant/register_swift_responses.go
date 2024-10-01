// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// RegisterSwiftReader is a Reader for the RegisterSwift structure.
type RegisterSwiftReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterSwiftReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRegisterSwiftNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRegisterSwiftDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRegisterSwiftNoContent creates a RegisterSwiftNoContent with default headers values
func NewRegisterSwiftNoContent() *RegisterSwiftNoContent {
	return &RegisterSwiftNoContent{}
}

/*
RegisterSwiftNoContent describes a response with status code 204, with default header values.

No Content
*/
type RegisterSwiftNoContent struct {
}

// IsSuccess returns true when this register swift no content response has a 2xx status code
func (o *RegisterSwiftNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this register swift no content response has a 3xx status code
func (o *RegisterSwiftNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this register swift no content response has a 4xx status code
func (o *RegisterSwiftNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this register swift no content response has a 5xx status code
func (o *RegisterSwiftNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this register swift no content response a status code equal to that given
func (o *RegisterSwiftNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the register swift no content response
func (o *RegisterSwiftNoContent) Code() int {
	return 204
}

func (o *RegisterSwiftNoContent) Error() string {
	return fmt.Sprintf("[POST /tenants/swift/register][%d] registerSwiftNoContent", 204)
}

func (o *RegisterSwiftNoContent) String() string {
	return fmt.Sprintf("[POST /tenants/swift/register][%d] registerSwiftNoContent", 204)
}

func (o *RegisterSwiftNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRegisterSwiftDefault creates a RegisterSwiftDefault with default headers values
func NewRegisterSwiftDefault(code int) *RegisterSwiftDefault {
	return &RegisterSwiftDefault{
		_statusCode: code,
	}
}

/*
RegisterSwiftDefault describes a response with status code -1, with default header values.

Error
*/
type RegisterSwiftDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this register swift default response has a 2xx status code
func (o *RegisterSwiftDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this register swift default response has a 3xx status code
func (o *RegisterSwiftDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this register swift default response has a 4xx status code
func (o *RegisterSwiftDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this register swift default response has a 5xx status code
func (o *RegisterSwiftDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this register swift default response a status code equal to that given
func (o *RegisterSwiftDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the register swift default response
func (o *RegisterSwiftDefault) Code() int {
	return o._statusCode
}

func (o *RegisterSwiftDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenants/swift/register][%d] RegisterSwift default %s", o._statusCode, payload)
}

func (o *RegisterSwiftDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /tenants/swift/register][%d] RegisterSwift default %s", o._statusCode, payload)
}

func (o *RegisterSwiftDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterSwiftDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

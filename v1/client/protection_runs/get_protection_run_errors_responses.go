// Code generated by go-swagger; DO NOT EDIT.

package protection_runs

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

// GetProtectionRunErrorsReader is a Reader for the GetProtectionRunErrors structure.
type GetProtectionRunErrorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProtectionRunErrorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProtectionRunErrorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProtectionRunErrorsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProtectionRunErrorsOK creates a GetProtectionRunErrorsOK with default headers values
func NewGetProtectionRunErrorsOK() *GetProtectionRunErrorsOK {
	return &GetProtectionRunErrorsOK{}
}

/*
GetProtectionRunErrorsOK describes a response with status code 200, with default header values.

Success
*/
type GetProtectionRunErrorsOK struct {
	Payload *models.ProtectionRunErrors
}

// IsSuccess returns true when this get protection run errors o k response has a 2xx status code
func (o *GetProtectionRunErrorsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get protection run errors o k response has a 3xx status code
func (o *GetProtectionRunErrorsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get protection run errors o k response has a 4xx status code
func (o *GetProtectionRunErrorsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get protection run errors o k response has a 5xx status code
func (o *GetProtectionRunErrorsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get protection run errors o k response a status code equal to that given
func (o *GetProtectionRunErrorsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get protection run errors o k response
func (o *GetProtectionRunErrorsOK) Code() int {
	return 200
}

func (o *GetProtectionRunErrorsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionRuns/errors][%d] getProtectionRunErrorsOK %s", 200, payload)
}

func (o *GetProtectionRunErrorsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionRuns/errors][%d] getProtectionRunErrorsOK %s", 200, payload)
}

func (o *GetProtectionRunErrorsOK) GetPayload() *models.ProtectionRunErrors {
	return o.Payload
}

func (o *GetProtectionRunErrorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionRunErrors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProtectionRunErrorsDefault creates a GetProtectionRunErrorsDefault with default headers values
func NewGetProtectionRunErrorsDefault(code int) *GetProtectionRunErrorsDefault {
	return &GetProtectionRunErrorsDefault{
		_statusCode: code,
	}
}

/*
GetProtectionRunErrorsDefault describes a response with status code -1, with default header values.

Error
*/
type GetProtectionRunErrorsDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get protection run errors default response has a 2xx status code
func (o *GetProtectionRunErrorsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get protection run errors default response has a 3xx status code
func (o *GetProtectionRunErrorsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get protection run errors default response has a 4xx status code
func (o *GetProtectionRunErrorsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get protection run errors default response has a 5xx status code
func (o *GetProtectionRunErrorsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get protection run errors default response a status code equal to that given
func (o *GetProtectionRunErrorsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get protection run errors default response
func (o *GetProtectionRunErrorsDefault) Code() int {
	return o._statusCode
}

func (o *GetProtectionRunErrorsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionRuns/errors][%d] GetProtectionRunErrors default %s", o._statusCode, payload)
}

func (o *GetProtectionRunErrorsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionRuns/errors][%d] GetProtectionRunErrors default %s", o._statusCode, payload)
}

func (o *GetProtectionRunErrorsDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetProtectionRunErrorsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

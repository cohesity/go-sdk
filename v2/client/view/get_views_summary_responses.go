// Code generated by go-swagger; DO NOT EDIT.

package view

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

// GetViewsSummaryReader is a Reader for the GetViewsSummary structure.
type GetViewsSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetViewsSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetViewsSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetViewsSummaryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetViewsSummaryOK creates a GetViewsSummaryOK with default headers values
func NewGetViewsSummaryOK() *GetViewsSummaryOK {
	return &GetViewsSummaryOK{}
}

/*
GetViewsSummaryOK describes a response with status code 200, with default header values.

Success
*/
type GetViewsSummaryOK struct {
	Payload *models.ViewsSummary
}

// IsSuccess returns true when this get views summary o k response has a 2xx status code
func (o *GetViewsSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get views summary o k response has a 3xx status code
func (o *GetViewsSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get views summary o k response has a 4xx status code
func (o *GetViewsSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get views summary o k response has a 5xx status code
func (o *GetViewsSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get views summary o k response a status code equal to that given
func (o *GetViewsSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get views summary o k response
func (o *GetViewsSummaryOK) Code() int {
	return 200
}

func (o *GetViewsSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views-summary][%d] getViewsSummaryOK %s", 200, payload)
}

func (o *GetViewsSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views-summary][%d] getViewsSummaryOK %s", 200, payload)
}

func (o *GetViewsSummaryOK) GetPayload() *models.ViewsSummary {
	return o.Payload
}

func (o *GetViewsSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ViewsSummary)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetViewsSummaryDefault creates a GetViewsSummaryDefault with default headers values
func NewGetViewsSummaryDefault(code int) *GetViewsSummaryDefault {
	return &GetViewsSummaryDefault{
		_statusCode: code,
	}
}

/*
GetViewsSummaryDefault describes a response with status code -1, with default header values.

Error
*/
type GetViewsSummaryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get views summary default response has a 2xx status code
func (o *GetViewsSummaryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get views summary default response has a 3xx status code
func (o *GetViewsSummaryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get views summary default response has a 4xx status code
func (o *GetViewsSummaryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get views summary default response has a 5xx status code
func (o *GetViewsSummaryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get views summary default response a status code equal to that given
func (o *GetViewsSummaryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get views summary default response
func (o *GetViewsSummaryDefault) Code() int {
	return o._statusCode
}

func (o *GetViewsSummaryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views-summary][%d] GetViewsSummary default %s", o._statusCode, payload)
}

func (o *GetViewsSummaryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file-services/views-summary][%d] GetViewsSummary default %s", o._statusCode, payload)
}

func (o *GetViewsSummaryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetViewsSummaryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

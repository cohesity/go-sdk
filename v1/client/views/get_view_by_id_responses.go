// Code generated by go-swagger; DO NOT EDIT.

package views

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

// GetViewByIDReader is a Reader for the GetViewByID structure.
type GetViewByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetViewByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetViewByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetViewByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetViewByIDOK creates a GetViewByIDOK with default headers values
func NewGetViewByIDOK() *GetViewByIDOK {
	return &GetViewByIDOK{}
}

/*
GetViewByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetViewByIDOK struct {
	Payload *models.View
}

// IsSuccess returns true when this get view by Id o k response has a 2xx status code
func (o *GetViewByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get view by Id o k response has a 3xx status code
func (o *GetViewByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get view by Id o k response has a 4xx status code
func (o *GetViewByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get view by Id o k response has a 5xx status code
func (o *GetViewByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get view by Id o k response a status code equal to that given
func (o *GetViewByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get view by Id o k response
func (o *GetViewByIDOK) Code() int {
	return 200
}

func (o *GetViewByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/id/{id}][%d] getViewByIdOK %s", 200, payload)
}

func (o *GetViewByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/id/{id}][%d] getViewByIdOK %s", 200, payload)
}

func (o *GetViewByIDOK) GetPayload() *models.View {
	return o.Payload
}

func (o *GetViewByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.View)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetViewByIDDefault creates a GetViewByIDDefault with default headers values
func NewGetViewByIDDefault(code int) *GetViewByIDDefault {
	return &GetViewByIDDefault{
		_statusCode: code,
	}
}

/*
GetViewByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetViewByIDDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get view by Id default response has a 2xx status code
func (o *GetViewByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get view by Id default response has a 3xx status code
func (o *GetViewByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get view by Id default response has a 4xx status code
func (o *GetViewByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get view by Id default response has a 5xx status code
func (o *GetViewByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get view by Id default response a status code equal to that given
func (o *GetViewByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get view by Id default response
func (o *GetViewByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetViewByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/id/{id}][%d] GetViewById default %s", o._statusCode, payload)
}

func (o *GetViewByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/id/{id}][%d] GetViewById default %s", o._statusCode, payload)
}

func (o *GetViewByIDDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetViewByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

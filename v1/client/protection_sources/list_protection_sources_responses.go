// Code generated by go-swagger; DO NOT EDIT.

package protection_sources

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

// ListProtectionSourcesReader is a Reader for the ListProtectionSources structure.
type ListProtectionSourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProtectionSourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProtectionSourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListProtectionSourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListProtectionSourcesOK creates a ListProtectionSourcesOK with default headers values
func NewListProtectionSourcesOK() *ListProtectionSourcesOK {
	return &ListProtectionSourcesOK{}
}

/*
ListProtectionSourcesOK describes a response with status code 200, with default header values.

Success
*/
type ListProtectionSourcesOK struct {
	Payload []*models.ProtectionSourceNode
}

// IsSuccess returns true when this list protection sources o k response has a 2xx status code
func (o *ListProtectionSourcesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list protection sources o k response has a 3xx status code
func (o *ListProtectionSourcesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list protection sources o k response has a 4xx status code
func (o *ListProtectionSourcesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list protection sources o k response has a 5xx status code
func (o *ListProtectionSourcesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list protection sources o k response a status code equal to that given
func (o *ListProtectionSourcesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list protection sources o k response
func (o *ListProtectionSourcesOK) Code() int {
	return 200
}

func (o *ListProtectionSourcesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionSources][%d] listProtectionSourcesOK %s", 200, payload)
}

func (o *ListProtectionSourcesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionSources][%d] listProtectionSourcesOK %s", 200, payload)
}

func (o *ListProtectionSourcesOK) GetPayload() []*models.ProtectionSourceNode {
	return o.Payload
}

func (o *ListProtectionSourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProtectionSourcesDefault creates a ListProtectionSourcesDefault with default headers values
func NewListProtectionSourcesDefault(code int) *ListProtectionSourcesDefault {
	return &ListProtectionSourcesDefault{
		_statusCode: code,
	}
}

/*
ListProtectionSourcesDefault describes a response with status code -1, with default header values.

Error
*/
type ListProtectionSourcesDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this list protection sources default response has a 2xx status code
func (o *ListProtectionSourcesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list protection sources default response has a 3xx status code
func (o *ListProtectionSourcesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list protection sources default response has a 4xx status code
func (o *ListProtectionSourcesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list protection sources default response has a 5xx status code
func (o *ListProtectionSourcesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list protection sources default response a status code equal to that given
func (o *ListProtectionSourcesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list protection sources default response
func (o *ListProtectionSourcesDefault) Code() int {
	return o._statusCode
}

func (o *ListProtectionSourcesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionSources][%d] ListProtectionSources default %s", o._statusCode, payload)
}

func (o *ListProtectionSourcesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionSources][%d] ListProtectionSources default %s", o._statusCode, payload)
}

func (o *ListProtectionSourcesDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *ListProtectionSourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

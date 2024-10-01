// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// ListFreeNodesReader is a Reader for the ListFreeNodes structure.
type ListFreeNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFreeNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFreeNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListFreeNodesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListFreeNodesOK creates a ListFreeNodesOK with default headers values
func NewListFreeNodesOK() *ListFreeNodesOK {
	return &ListFreeNodesOK{}
}

/*
ListFreeNodesOK describes a response with status code 200, with default header values.

Success
*/
type ListFreeNodesOK struct {
	Payload *models.FreeNodes
}

// IsSuccess returns true when this list free nodes o k response has a 2xx status code
func (o *ListFreeNodesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list free nodes o k response has a 3xx status code
func (o *ListFreeNodesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list free nodes o k response has a 4xx status code
func (o *ListFreeNodesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list free nodes o k response has a 5xx status code
func (o *ListFreeNodesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list free nodes o k response a status code equal to that given
func (o *ListFreeNodesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list free nodes o k response
func (o *ListFreeNodesOK) Code() int {
	return 200
}

func (o *ListFreeNodesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/nodes/free][%d] listFreeNodesOK %s", 200, payload)
}

func (o *ListFreeNodesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/nodes/free][%d] listFreeNodesOK %s", 200, payload)
}

func (o *ListFreeNodesOK) GetPayload() *models.FreeNodes {
	return o.Payload
}

func (o *ListFreeNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FreeNodes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFreeNodesDefault creates a ListFreeNodesDefault with default headers values
func NewListFreeNodesDefault(code int) *ListFreeNodesDefault {
	return &ListFreeNodesDefault{
		_statusCode: code,
	}
}

/*
ListFreeNodesDefault describes a response with status code -1, with default header values.

Error
*/
type ListFreeNodesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this list free nodes default response has a 2xx status code
func (o *ListFreeNodesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this list free nodes default response has a 3xx status code
func (o *ListFreeNodesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this list free nodes default response has a 4xx status code
func (o *ListFreeNodesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this list free nodes default response has a 5xx status code
func (o *ListFreeNodesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this list free nodes default response a status code equal to that given
func (o *ListFreeNodesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the list free nodes default response
func (o *ListFreeNodesDefault) Code() int {
	return o._statusCode
}

func (o *ListFreeNodesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/nodes/free][%d] ListFreeNodes default %s", o._statusCode, payload)
}

func (o *ListFreeNodesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/nodes/free][%d] ListFreeNodes default %s", o._statusCode, payload)
}

func (o *ListFreeNodesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListFreeNodesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

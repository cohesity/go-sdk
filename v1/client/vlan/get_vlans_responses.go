// Code generated by go-swagger; DO NOT EDIT.

package vlan

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

// GetVlansReader is a Reader for the GetVlans structure.
type GetVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVlansDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVlansOK creates a GetVlansOK with default headers values
func NewGetVlansOK() *GetVlansOK {
	return &GetVlansOK{}
}

/*
GetVlansOK describes a response with status code 200, with default header values.

Success
*/
type GetVlansOK struct {
	Payload []*models.Vlan
}

// IsSuccess returns true when this get vlans o k response has a 2xx status code
func (o *GetVlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get vlans o k response has a 3xx status code
func (o *GetVlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get vlans o k response has a 4xx status code
func (o *GetVlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get vlans o k response has a 5xx status code
func (o *GetVlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get vlans o k response a status code equal to that given
func (o *GetVlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get vlans o k response
func (o *GetVlansOK) Code() int {
	return 200
}

func (o *GetVlansOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/vlans][%d] getVlansOK %s", 200, payload)
}

func (o *GetVlansOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/vlans][%d] getVlansOK %s", 200, payload)
}

func (o *GetVlansOK) GetPayload() []*models.Vlan {
	return o.Payload
}

func (o *GetVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVlansDefault creates a GetVlansDefault with default headers values
func NewGetVlansDefault(code int) *GetVlansDefault {
	return &GetVlansDefault{
		_statusCode: code,
	}
}

/*
GetVlansDefault describes a response with status code -1, with default header values.

Error
*/
type GetVlansDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get vlans default response has a 2xx status code
func (o *GetVlansDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get vlans default response has a 3xx status code
func (o *GetVlansDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get vlans default response has a 4xx status code
func (o *GetVlansDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get vlans default response has a 5xx status code
func (o *GetVlansDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get vlans default response a status code equal to that given
func (o *GetVlansDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get vlans default response
func (o *GetVlansDefault) Code() int {
	return o._statusCode
}

func (o *GetVlansDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/vlans][%d] GetVlans default %s", o._statusCode, payload)
}

func (o *GetVlansDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/vlans][%d] GetVlans default %s", o._statusCode, payload)
}

func (o *GetVlansDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetVlansDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

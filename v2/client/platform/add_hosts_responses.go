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

// AddHostsReader is a Reader for the AddHosts structure.
type AddHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddHostsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddHostsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddHostsCreated creates a AddHostsCreated with default headers values
func NewAddHostsCreated() *AddHostsCreated {
	return &AddHostsCreated{}
}

/*
AddHostsCreated describes a response with status code 201, with default header values.

Success
*/
type AddHostsCreated struct {
	Payload *models.HostMappings
}

// IsSuccess returns true when this add hosts created response has a 2xx status code
func (o *AddHostsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add hosts created response has a 3xx status code
func (o *AddHostsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add hosts created response has a 4xx status code
func (o *AddHostsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add hosts created response has a 5xx status code
func (o *AddHostsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add hosts created response a status code equal to that given
func (o *AddHostsCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add hosts created response
func (o *AddHostsCreated) Code() int {
	return 201
}

func (o *AddHostsCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters/host-mappings][%d] addHostsCreated %s", 201, payload)
}

func (o *AddHostsCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters/host-mappings][%d] addHostsCreated %s", 201, payload)
}

func (o *AddHostsCreated) GetPayload() *models.HostMappings {
	return o.Payload
}

func (o *AddHostsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostMappings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddHostsDefault creates a AddHostsDefault with default headers values
func NewAddHostsDefault(code int) *AddHostsDefault {
	return &AddHostsDefault{
		_statusCode: code,
	}
}

/*
AddHostsDefault describes a response with status code -1, with default header values.

Error
*/
type AddHostsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this add hosts default response has a 2xx status code
func (o *AddHostsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this add hosts default response has a 3xx status code
func (o *AddHostsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this add hosts default response has a 4xx status code
func (o *AddHostsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this add hosts default response has a 5xx status code
func (o *AddHostsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this add hosts default response a status code equal to that given
func (o *AddHostsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the add hosts default response
func (o *AddHostsDefault) Code() int {
	return o._statusCode
}

func (o *AddHostsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters/host-mappings][%d] AddHosts default %s", o._statusCode, payload)
}

func (o *AddHostsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /clusters/host-mappings][%d] AddHosts default %s", o._statusCode, payload)
}

func (o *AddHostsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *AddHostsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

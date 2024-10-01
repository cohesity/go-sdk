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

// CreateClusterVlanReader is a Reader for the CreateClusterVlan structure.
type CreateClusterVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateClusterVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateClusterVlanCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateClusterVlanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateClusterVlanCreated creates a CreateClusterVlanCreated with default headers values
func NewCreateClusterVlanCreated() *CreateClusterVlanCreated {
	return &CreateClusterVlanCreated{}
}

/*
CreateClusterVlanCreated describes a response with status code 201, with default header values.

Success
*/
type CreateClusterVlanCreated struct {
	Payload *models.ClusterVlanParams
}

// IsSuccess returns true when this create cluster vlan created response has a 2xx status code
func (o *CreateClusterVlanCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cluster vlan created response has a 3xx status code
func (o *CreateClusterVlanCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cluster vlan created response has a 4xx status code
func (o *CreateClusterVlanCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cluster vlan created response has a 5xx status code
func (o *CreateClusterVlanCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create cluster vlan created response a status code equal to that given
func (o *CreateClusterVlanCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create cluster vlan created response
func (o *CreateClusterVlanCreated) Code() int {
	return 201
}

func (o *CreateClusterVlanCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/vlans][%d] createClusterVlanCreated %s", 201, payload)
}

func (o *CreateClusterVlanCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/vlans][%d] createClusterVlanCreated %s", 201, payload)
}

func (o *CreateClusterVlanCreated) GetPayload() *models.ClusterVlanParams {
	return o.Payload
}

func (o *CreateClusterVlanCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterVlanParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateClusterVlanDefault creates a CreateClusterVlanDefault with default headers values
func NewCreateClusterVlanDefault(code int) *CreateClusterVlanDefault {
	return &CreateClusterVlanDefault{
		_statusCode: code,
	}
}

/*
CreateClusterVlanDefault describes a response with status code -1, with default header values.

Error
*/
type CreateClusterVlanDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create cluster vlan default response has a 2xx status code
func (o *CreateClusterVlanDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cluster vlan default response has a 3xx status code
func (o *CreateClusterVlanDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cluster vlan default response has a 4xx status code
func (o *CreateClusterVlanDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cluster vlan default response has a 5xx status code
func (o *CreateClusterVlanDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cluster vlan default response a status code equal to that given
func (o *CreateClusterVlanDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create cluster vlan default response
func (o *CreateClusterVlanDefault) Code() int {
	return o._statusCode
}

func (o *CreateClusterVlanDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/vlans][%d] CreateClusterVlan default %s", o._statusCode, payload)
}

func (o *CreateClusterVlanDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /network/vlans][%d] CreateClusterVlan default %s", o._statusCode, payload)
}

func (o *CreateClusterVlanDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateClusterVlanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

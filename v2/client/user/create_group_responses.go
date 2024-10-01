// Code generated by go-swagger; DO NOT EDIT.

package user

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

// CreateGroupReader is a Reader for the CreateGroup structure.
type CreateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateGroupCreated creates a CreateGroupCreated with default headers values
func NewCreateGroupCreated() *CreateGroupCreated {
	return &CreateGroupCreated{}
}

/*
CreateGroupCreated describes a response with status code 201, with default header values.

Success
*/
type CreateGroupCreated struct {
	Payload *models.Groups
}

// IsSuccess returns true when this create group created response has a 2xx status code
func (o *CreateGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create group created response has a 3xx status code
func (o *CreateGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create group created response has a 4xx status code
func (o *CreateGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create group created response has a 5xx status code
func (o *CreateGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create group created response a status code equal to that given
func (o *CreateGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create group created response
func (o *CreateGroupCreated) Code() int {
	return 201
}

func (o *CreateGroupCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /groups][%d] createGroupCreated %s", 201, payload)
}

func (o *CreateGroupCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /groups][%d] createGroupCreated %s", 201, payload)
}

func (o *CreateGroupCreated) GetPayload() *models.Groups {
	return o.Payload
}

func (o *CreateGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Groups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateGroupDefault creates a CreateGroupDefault with default headers values
func NewCreateGroupDefault(code int) *CreateGroupDefault {
	return &CreateGroupDefault{
		_statusCode: code,
	}
}

/*
CreateGroupDefault describes a response with status code -1, with default header values.

Error
*/
type CreateGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create group default response has a 2xx status code
func (o *CreateGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create group default response has a 3xx status code
func (o *CreateGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create group default response has a 4xx status code
func (o *CreateGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create group default response has a 5xx status code
func (o *CreateGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create group default response a status code equal to that given
func (o *CreateGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create group default response
func (o *CreateGroupDefault) Code() int {
	return o._statusCode
}

func (o *CreateGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /groups][%d] CreateGroup default %s", o._statusCode, payload)
}

func (o *CreateGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /groups][%d] CreateGroup default %s", o._statusCode, payload)
}

func (o *CreateGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

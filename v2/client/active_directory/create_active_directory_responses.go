// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// CreateActiveDirectoryReader is a Reader for the CreateActiveDirectory structure.
type CreateActiveDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateActiveDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateActiveDirectoryCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateActiveDirectoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateActiveDirectoryCreated creates a CreateActiveDirectoryCreated with default headers values
func NewCreateActiveDirectoryCreated() *CreateActiveDirectoryCreated {
	return &CreateActiveDirectoryCreated{}
}

/*
CreateActiveDirectoryCreated describes a response with status code 201, with default header values.

Success
*/
type CreateActiveDirectoryCreated struct {
	Payload *models.ActiveDirectory
}

// IsSuccess returns true when this create active directory created response has a 2xx status code
func (o *CreateActiveDirectoryCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create active directory created response has a 3xx status code
func (o *CreateActiveDirectoryCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create active directory created response has a 4xx status code
func (o *CreateActiveDirectoryCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create active directory created response has a 5xx status code
func (o *CreateActiveDirectoryCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create active directory created response a status code equal to that given
func (o *CreateActiveDirectoryCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create active directory created response
func (o *CreateActiveDirectoryCreated) Code() int {
	return 201
}

func (o *CreateActiveDirectoryCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /active-directories][%d] createActiveDirectoryCreated %s", 201, payload)
}

func (o *CreateActiveDirectoryCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /active-directories][%d] createActiveDirectoryCreated %s", 201, payload)
}

func (o *CreateActiveDirectoryCreated) GetPayload() *models.ActiveDirectory {
	return o.Payload
}

func (o *CreateActiveDirectoryCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveDirectory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateActiveDirectoryDefault creates a CreateActiveDirectoryDefault with default headers values
func NewCreateActiveDirectoryDefault(code int) *CreateActiveDirectoryDefault {
	return &CreateActiveDirectoryDefault{
		_statusCode: code,
	}
}

/*
CreateActiveDirectoryDefault describes a response with status code -1, with default header values.

Error
*/
type CreateActiveDirectoryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this create active directory default response has a 2xx status code
func (o *CreateActiveDirectoryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create active directory default response has a 3xx status code
func (o *CreateActiveDirectoryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create active directory default response has a 4xx status code
func (o *CreateActiveDirectoryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create active directory default response has a 5xx status code
func (o *CreateActiveDirectoryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create active directory default response a status code equal to that given
func (o *CreateActiveDirectoryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create active directory default response
func (o *CreateActiveDirectoryDefault) Code() int {
	return o._statusCode
}

func (o *CreateActiveDirectoryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /active-directories][%d] CreateActiveDirectory default %s", o._statusCode, payload)
}

func (o *CreateActiveDirectoryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /active-directories][%d] CreateActiveDirectory default %s", o._statusCode, payload)
}

func (o *CreateActiveDirectoryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateActiveDirectoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

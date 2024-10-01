// Code generated by go-swagger; DO NOT EDIT.

package remote_restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v1/models"
)

// CreateCloudDomainMigrationRequestReader is a Reader for the CreateCloudDomainMigrationRequest structure.
type CreateCloudDomainMigrationRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCloudDomainMigrationRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreateCloudDomainMigrationRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateCloudDomainMigrationRequestDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateCloudDomainMigrationRequestNoContent creates a CreateCloudDomainMigrationRequestNoContent with default headers values
func NewCreateCloudDomainMigrationRequestNoContent() *CreateCloudDomainMigrationRequestNoContent {
	return &CreateCloudDomainMigrationRequestNoContent{}
}

/*
CreateCloudDomainMigrationRequestNoContent describes a response with status code 204, with default header values.

No Content
*/
type CreateCloudDomainMigrationRequestNoContent struct {
}

// IsSuccess returns true when this create cloud domain migration request no content response has a 2xx status code
func (o *CreateCloudDomainMigrationRequestNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create cloud domain migration request no content response has a 3xx status code
func (o *CreateCloudDomainMigrationRequestNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create cloud domain migration request no content response has a 4xx status code
func (o *CreateCloudDomainMigrationRequestNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this create cloud domain migration request no content response has a 5xx status code
func (o *CreateCloudDomainMigrationRequestNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this create cloud domain migration request no content response a status code equal to that given
func (o *CreateCloudDomainMigrationRequestNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the create cloud domain migration request no content response
func (o *CreateCloudDomainMigrationRequestNoContent) Code() int {
	return 204
}

func (o *CreateCloudDomainMigrationRequestNoContent) Error() string {
	return fmt.Sprintf("[POST /public/remoteVaults/cloudDomainMigration][%d] createCloudDomainMigrationRequestNoContent", 204)
}

func (o *CreateCloudDomainMigrationRequestNoContent) String() string {
	return fmt.Sprintf("[POST /public/remoteVaults/cloudDomainMigration][%d] createCloudDomainMigrationRequestNoContent", 204)
}

func (o *CreateCloudDomainMigrationRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateCloudDomainMigrationRequestDefault creates a CreateCloudDomainMigrationRequestDefault with default headers values
func NewCreateCloudDomainMigrationRequestDefault(code int) *CreateCloudDomainMigrationRequestDefault {
	return &CreateCloudDomainMigrationRequestDefault{
		_statusCode: code,
	}
}

/*
CreateCloudDomainMigrationRequestDefault describes a response with status code -1, with default header values.

Error
*/
type CreateCloudDomainMigrationRequestDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create cloud domain migration request default response has a 2xx status code
func (o *CreateCloudDomainMigrationRequestDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create cloud domain migration request default response has a 3xx status code
func (o *CreateCloudDomainMigrationRequestDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create cloud domain migration request default response has a 4xx status code
func (o *CreateCloudDomainMigrationRequestDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create cloud domain migration request default response has a 5xx status code
func (o *CreateCloudDomainMigrationRequestDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create cloud domain migration request default response a status code equal to that given
func (o *CreateCloudDomainMigrationRequestDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create cloud domain migration request default response
func (o *CreateCloudDomainMigrationRequestDefault) Code() int {
	return o._statusCode
}

func (o *CreateCloudDomainMigrationRequestDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/remoteVaults/cloudDomainMigration][%d] CreateCloudDomainMigrationRequest default %s", o._statusCode, payload)
}

func (o *CreateCloudDomainMigrationRequestDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/remoteVaults/cloudDomainMigration][%d] CreateCloudDomainMigrationRequest default %s", o._statusCode, payload)
}

func (o *CreateCloudDomainMigrationRequestDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateCloudDomainMigrationRequestDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

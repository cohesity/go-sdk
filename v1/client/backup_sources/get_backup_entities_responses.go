// Code generated by go-swagger; DO NOT EDIT.

package backup_sources

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

// GetBackupEntitiesReader is a Reader for the GetBackupEntities structure.
type GetBackupEntitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupEntitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupEntitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBackupEntitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBackupEntitiesOK creates a GetBackupEntitiesOK with default headers values
func NewGetBackupEntitiesOK() *GetBackupEntitiesOK {
	return &GetBackupEntitiesOK{}
}

/*
GetBackupEntitiesOK describes a response with status code 200, with default header values.

Success
*/
type GetBackupEntitiesOK struct {
	Payload []*models.PrivateEntityProto
}

// IsSuccess returns true when this get backup entities o k response has a 2xx status code
func (o *GetBackupEntitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get backup entities o k response has a 3xx status code
func (o *GetBackupEntitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get backup entities o k response has a 4xx status code
func (o *GetBackupEntitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get backup entities o k response has a 5xx status code
func (o *GetBackupEntitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get backup entities o k response a status code equal to that given
func (o *GetBackupEntitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get backup entities o k response
func (o *GetBackupEntitiesOK) Code() int {
	return 200
}

func (o *GetBackupEntitiesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupEntities][%d] getBackupEntitiesOK %s", 200, payload)
}

func (o *GetBackupEntitiesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupEntities][%d] getBackupEntitiesOK %s", 200, payload)
}

func (o *GetBackupEntitiesOK) GetPayload() []*models.PrivateEntityProto {
	return o.Payload
}

func (o *GetBackupEntitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupEntitiesDefault creates a GetBackupEntitiesDefault with default headers values
func NewGetBackupEntitiesDefault(code int) *GetBackupEntitiesDefault {
	return &GetBackupEntitiesDefault{
		_statusCode: code,
	}
}

/*
GetBackupEntitiesDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetBackupEntitiesDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get backup entities default response has a 2xx status code
func (o *GetBackupEntitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get backup entities default response has a 3xx status code
func (o *GetBackupEntitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get backup entities default response has a 4xx status code
func (o *GetBackupEntitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get backup entities default response has a 5xx status code
func (o *GetBackupEntitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get backup entities default response a status code equal to that given
func (o *GetBackupEntitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get backup entities default response
func (o *GetBackupEntitiesDefault) Code() int {
	return o._statusCode
}

func (o *GetBackupEntitiesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupEntities][%d] GetBackupEntities default %s", o._statusCode, payload)
}

func (o *GetBackupEntitiesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /backupEntities][%d] GetBackupEntities default %s", o._statusCode, payload)
}

func (o *GetBackupEntitiesDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetBackupEntitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

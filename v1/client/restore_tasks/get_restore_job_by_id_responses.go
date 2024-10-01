// Code generated by go-swagger; DO NOT EDIT.

package restore_tasks

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

// GetRestoreJobByIDReader is a Reader for the GetRestoreJobByID structure.
type GetRestoreJobByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestoreJobByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRestoreJobByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRestoreJobByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRestoreJobByIDOK creates a GetRestoreJobByIDOK with default headers values
func NewGetRestoreJobByIDOK() *GetRestoreJobByIDOK {
	return &GetRestoreJobByIDOK{}
}

/*
GetRestoreJobByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetRestoreJobByIDOK struct {
	Payload []*models.RestoreTaskWrapper
}

// IsSuccess returns true when this get restore job by Id o k response has a 2xx status code
func (o *GetRestoreJobByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get restore job by Id o k response has a 3xx status code
func (o *GetRestoreJobByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get restore job by Id o k response has a 4xx status code
func (o *GetRestoreJobByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get restore job by Id o k response has a 5xx status code
func (o *GetRestoreJobByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get restore job by Id o k response a status code equal to that given
func (o *GetRestoreJobByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get restore job by Id o k response
func (o *GetRestoreJobByIDOK) Code() int {
	return 200
}

func (o *GetRestoreJobByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /restoretasks/{id}][%d] getRestoreJobByIdOK %s", 200, payload)
}

func (o *GetRestoreJobByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /restoretasks/{id}][%d] getRestoreJobByIdOK %s", 200, payload)
}

func (o *GetRestoreJobByIDOK) GetPayload() []*models.RestoreTaskWrapper {
	return o.Payload
}

func (o *GetRestoreJobByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRestoreJobByIDDefault creates a GetRestoreJobByIDDefault with default headers values
func NewGetRestoreJobByIDDefault(code int) *GetRestoreJobByIDDefault {
	return &GetRestoreJobByIDDefault{
		_statusCode: code,
	}
}

/*
GetRestoreJobByIDDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetRestoreJobByIDDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get restore job by Id default response has a 2xx status code
func (o *GetRestoreJobByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get restore job by Id default response has a 3xx status code
func (o *GetRestoreJobByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get restore job by Id default response has a 4xx status code
func (o *GetRestoreJobByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get restore job by Id default response has a 5xx status code
func (o *GetRestoreJobByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get restore job by Id default response a status code equal to that given
func (o *GetRestoreJobByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get restore job by Id default response
func (o *GetRestoreJobByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetRestoreJobByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /restoretasks/{id}][%d] GetRestoreJobById default %s", o._statusCode, payload)
}

func (o *GetRestoreJobByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /restoretasks/{id}][%d] GetRestoreJobById default %s", o._statusCode, payload)
}

func (o *GetRestoreJobByIDDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetRestoreJobByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

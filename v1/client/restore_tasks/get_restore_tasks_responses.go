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

// GetRestoreTasksReader is a Reader for the GetRestoreTasks structure.
type GetRestoreTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestoreTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRestoreTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRestoreTasksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRestoreTasksOK creates a GetRestoreTasksOK with default headers values
func NewGetRestoreTasksOK() *GetRestoreTasksOK {
	return &GetRestoreTasksOK{}
}

/*
GetRestoreTasksOK describes a response with status code 200, with default header values.

Success
*/
type GetRestoreTasksOK struct {
	Payload []*models.RestoreTask
}

// IsSuccess returns true when this get restore tasks o k response has a 2xx status code
func (o *GetRestoreTasksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get restore tasks o k response has a 3xx status code
func (o *GetRestoreTasksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get restore tasks o k response has a 4xx status code
func (o *GetRestoreTasksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get restore tasks o k response has a 5xx status code
func (o *GetRestoreTasksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get restore tasks o k response a status code equal to that given
func (o *GetRestoreTasksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get restore tasks o k response
func (o *GetRestoreTasksOK) Code() int {
	return 200
}

func (o *GetRestoreTasksOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/tasks][%d] getRestoreTasksOK %s", 200, payload)
}

func (o *GetRestoreTasksOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/tasks][%d] getRestoreTasksOK %s", 200, payload)
}

func (o *GetRestoreTasksOK) GetPayload() []*models.RestoreTask {
	return o.Payload
}

func (o *GetRestoreTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRestoreTasksDefault creates a GetRestoreTasksDefault with default headers values
func NewGetRestoreTasksDefault(code int) *GetRestoreTasksDefault {
	return &GetRestoreTasksDefault{
		_statusCode: code,
	}
}

/*
GetRestoreTasksDefault describes a response with status code -1, with default header values.

Error
*/
type GetRestoreTasksDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get restore tasks default response has a 2xx status code
func (o *GetRestoreTasksDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get restore tasks default response has a 3xx status code
func (o *GetRestoreTasksDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get restore tasks default response has a 4xx status code
func (o *GetRestoreTasksDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get restore tasks default response has a 5xx status code
func (o *GetRestoreTasksDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get restore tasks default response a status code equal to that given
func (o *GetRestoreTasksDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get restore tasks default response
func (o *GetRestoreTasksDefault) Code() int {
	return o._statusCode
}

func (o *GetRestoreTasksDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/tasks][%d] GetRestoreTasks default %s", o._statusCode, payload)
}

func (o *GetRestoreTasksDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/tasks][%d] GetRestoreTasks default %s", o._statusCode, payload)
}

func (o *GetRestoreTasksDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetRestoreTasksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

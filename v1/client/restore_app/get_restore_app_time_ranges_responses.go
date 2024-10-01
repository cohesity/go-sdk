// Code generated by go-swagger; DO NOT EDIT.

package restore_app

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

// GetRestoreAppTimeRangesReader is a Reader for the GetRestoreAppTimeRanges structure.
type GetRestoreAppTimeRangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRestoreAppTimeRangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRestoreAppTimeRangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRestoreAppTimeRangesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRestoreAppTimeRangesOK creates a GetRestoreAppTimeRangesOK with default headers values
func NewGetRestoreAppTimeRangesOK() *GetRestoreAppTimeRangesOK {
	return &GetRestoreAppTimeRangesOK{}
}

/*
GetRestoreAppTimeRangesOK describes a response with status code 200, with default header values.

Success
*/
type GetRestoreAppTimeRangesOK struct {
	Payload *models.GetRestoreAppTimeRangesResult
}

// IsSuccess returns true when this get restore app time ranges o k response has a 2xx status code
func (o *GetRestoreAppTimeRangesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get restore app time ranges o k response has a 3xx status code
func (o *GetRestoreAppTimeRangesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get restore app time ranges o k response has a 4xx status code
func (o *GetRestoreAppTimeRangesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get restore app time ranges o k response has a 5xx status code
func (o *GetRestoreAppTimeRangesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get restore app time ranges o k response a status code equal to that given
func (o *GetRestoreAppTimeRangesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get restore app time ranges o k response
func (o *GetRestoreAppTimeRangesOK) Code() int {
	return 200
}

func (o *GetRestoreAppTimeRangesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreApp/timeRanges][%d] getRestoreAppTimeRangesOK %s", 200, payload)
}

func (o *GetRestoreAppTimeRangesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreApp/timeRanges][%d] getRestoreAppTimeRangesOK %s", 200, payload)
}

func (o *GetRestoreAppTimeRangesOK) GetPayload() *models.GetRestoreAppTimeRangesResult {
	return o.Payload
}

func (o *GetRestoreAppTimeRangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRestoreAppTimeRangesResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRestoreAppTimeRangesDefault creates a GetRestoreAppTimeRangesDefault with default headers values
func NewGetRestoreAppTimeRangesDefault(code int) *GetRestoreAppTimeRangesDefault {
	return &GetRestoreAppTimeRangesDefault{
		_statusCode: code,
	}
}

/*
GetRestoreAppTimeRangesDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetRestoreAppTimeRangesDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get restore app time ranges default response has a 2xx status code
func (o *GetRestoreAppTimeRangesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get restore app time ranges default response has a 3xx status code
func (o *GetRestoreAppTimeRangesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get restore app time ranges default response has a 4xx status code
func (o *GetRestoreAppTimeRangesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get restore app time ranges default response has a 5xx status code
func (o *GetRestoreAppTimeRangesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get restore app time ranges default response a status code equal to that given
func (o *GetRestoreAppTimeRangesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get restore app time ranges default response
func (o *GetRestoreAppTimeRangesDefault) Code() int {
	return o._statusCode
}

func (o *GetRestoreAppTimeRangesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreApp/timeRanges][%d] GetRestoreAppTimeRanges default %s", o._statusCode, payload)
}

func (o *GetRestoreAppTimeRangesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /restoreApp/timeRanges][%d] GetRestoreAppTimeRanges default %s", o._statusCode, payload)
}

func (o *GetRestoreAppTimeRangesDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetRestoreAppTimeRangesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

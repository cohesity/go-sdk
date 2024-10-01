// Code generated by go-swagger; DO NOT EDIT.

package views

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

// GetFileLockStatusReader is a Reader for the GetFileLockStatus structure.
type GetFileLockStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileLockStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileLockStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFileLockStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFileLockStatusOK creates a GetFileLockStatusOK with default headers values
func NewGetFileLockStatusOK() *GetFileLockStatusOK {
	return &GetFileLockStatusOK{}
}

/*
GetFileLockStatusOK describes a response with status code 200, with default header values.

Get lock file status response
*/
type GetFileLockStatusOK struct {
	Payload *models.FileLockStatus
}

// IsSuccess returns true when this get file lock status o k response has a 2xx status code
func (o *GetFileLockStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file lock status o k response has a 3xx status code
func (o *GetFileLockStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file lock status o k response has a 4xx status code
func (o *GetFileLockStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file lock status o k response has a 5xx status code
func (o *GetFileLockStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file lock status o k response a status code equal to that given
func (o *GetFileLockStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file lock status o k response
func (o *GetFileLockStatusOK) Code() int {
	return 200
}

func (o *GetFileLockStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/{name}/fileLocks][%d] getFileLockStatusOK %s", 200, payload)
}

func (o *GetFileLockStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/{name}/fileLocks][%d] getFileLockStatusOK %s", 200, payload)
}

func (o *GetFileLockStatusOK) GetPayload() *models.FileLockStatus {
	return o.Payload
}

func (o *GetFileLockStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileLockStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileLockStatusDefault creates a GetFileLockStatusDefault with default headers values
func NewGetFileLockStatusDefault(code int) *GetFileLockStatusDefault {
	return &GetFileLockStatusDefault{
		_statusCode: code,
	}
}

/*
GetFileLockStatusDefault describes a response with status code -1, with default header values.

Error
*/
type GetFileLockStatusDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get file lock status default response has a 2xx status code
func (o *GetFileLockStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get file lock status default response has a 3xx status code
func (o *GetFileLockStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get file lock status default response has a 4xx status code
func (o *GetFileLockStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get file lock status default response has a 5xx status code
func (o *GetFileLockStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get file lock status default response a status code equal to that given
func (o *GetFileLockStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get file lock status default response
func (o *GetFileLockStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetFileLockStatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/{name}/fileLocks][%d] GetFileLockStatus default %s", o._statusCode, payload)
}

func (o *GetFileLockStatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/views/{name}/fileLocks][%d] GetFileLockStatus default %s", o._statusCode, payload)
}

func (o *GetFileLockStatusDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetFileLockStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

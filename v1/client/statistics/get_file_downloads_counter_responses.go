// Code generated by go-swagger; DO NOT EDIT.

package statistics

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

// GetFileDownloadsCounterReader is a Reader for the GetFileDownloadsCounter structure.
type GetFileDownloadsCounterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileDownloadsCounterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileDownloadsCounterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFileDownloadsCounterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFileDownloadsCounterOK creates a GetFileDownloadsCounterOK with default headers values
func NewGetFileDownloadsCounterOK() *GetFileDownloadsCounterOK {
	return &GetFileDownloadsCounterOK{}
}

/*
GetFileDownloadsCounterOK describes a response with status code 200, with default header values.

Success
*/
type GetFileDownloadsCounterOK struct {
	Payload []*models.Counter
}

// IsSuccess returns true when this get file downloads counter o k response has a 2xx status code
func (o *GetFileDownloadsCounterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file downloads counter o k response has a 3xx status code
func (o *GetFileDownloadsCounterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file downloads counter o k response has a 4xx status code
func (o *GetFileDownloadsCounterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file downloads counter o k response has a 5xx status code
func (o *GetFileDownloadsCounterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file downloads counter o k response a status code equal to that given
func (o *GetFileDownloadsCounterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file downloads counter o k response
func (o *GetFileDownloadsCounterOK) Code() int {
	return 200
}

func (o *GetFileDownloadsCounterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stats/fileDownloads][%d] getFileDownloadsCounterOK %s", 200, payload)
}

func (o *GetFileDownloadsCounterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stats/fileDownloads][%d] getFileDownloadsCounterOK %s", 200, payload)
}

func (o *GetFileDownloadsCounterOK) GetPayload() []*models.Counter {
	return o.Payload
}

func (o *GetFileDownloadsCounterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileDownloadsCounterDefault creates a GetFileDownloadsCounterDefault with default headers values
func NewGetFileDownloadsCounterDefault(code int) *GetFileDownloadsCounterDefault {
	return &GetFileDownloadsCounterDefault{
		_statusCode: code,
	}
}

/*
GetFileDownloadsCounterDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetFileDownloadsCounterDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get file downloads counter default response has a 2xx status code
func (o *GetFileDownloadsCounterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get file downloads counter default response has a 3xx status code
func (o *GetFileDownloadsCounterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get file downloads counter default response has a 4xx status code
func (o *GetFileDownloadsCounterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get file downloads counter default response has a 5xx status code
func (o *GetFileDownloadsCounterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get file downloads counter default response a status code equal to that given
func (o *GetFileDownloadsCounterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get file downloads counter default response
func (o *GetFileDownloadsCounterDefault) Code() int {
	return o._statusCode
}

func (o *GetFileDownloadsCounterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stats/fileDownloads][%d] GetFileDownloadsCounter default %s", o._statusCode, payload)
}

func (o *GetFileDownloadsCounterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /stats/fileDownloads][%d] GetFileDownloadsCounter default %s", o._statusCode, payload)
}

func (o *GetFileDownloadsCounterDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetFileDownloadsCounterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

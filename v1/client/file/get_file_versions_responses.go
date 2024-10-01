// Code generated by go-swagger; DO NOT EDIT.

package file

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

// GetFileVersionsReader is a Reader for the GetFileVersions structure.
type GetFileVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFileVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFileVersionsOK creates a GetFileVersionsOK with default headers values
func NewGetFileVersionsOK() *GetFileVersionsOK {
	return &GetFileVersionsOK{}
}

/*
GetFileVersionsOK describes a response with status code 200, with default header values.

Success
*/
type GetFileVersionsOK struct {
	Payload *models.GetCrackedFileVersionsResult
}

// IsSuccess returns true when this get file versions o k response has a 2xx status code
func (o *GetFileVersionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file versions o k response has a 3xx status code
func (o *GetFileVersionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file versions o k response has a 4xx status code
func (o *GetFileVersionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file versions o k response has a 5xx status code
func (o *GetFileVersionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file versions o k response a status code equal to that given
func (o *GetFileVersionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file versions o k response
func (o *GetFileVersionsOK) Code() int {
	return 200
}

func (o *GetFileVersionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/versions][%d] getFileVersionsOK %s", 200, payload)
}

func (o *GetFileVersionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/versions][%d] getFileVersionsOK %s", 200, payload)
}

func (o *GetFileVersionsOK) GetPayload() *models.GetCrackedFileVersionsResult {
	return o.Payload
}

func (o *GetFileVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCrackedFileVersionsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileVersionsDefault creates a GetFileVersionsDefault with default headers values
func NewGetFileVersionsDefault(code int) *GetFileVersionsDefault {
	return &GetFileVersionsDefault{
		_statusCode: code,
	}
}

/*
GetFileVersionsDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetFileVersionsDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get file versions default response has a 2xx status code
func (o *GetFileVersionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get file versions default response has a 3xx status code
func (o *GetFileVersionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get file versions default response has a 4xx status code
func (o *GetFileVersionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get file versions default response has a 5xx status code
func (o *GetFileVersionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get file versions default response a status code equal to that given
func (o *GetFileVersionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get file versions default response
func (o *GetFileVersionsDefault) Code() int {
	return o._statusCode
}

func (o *GetFileVersionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/versions][%d] GetFileVersions default %s", o._statusCode, payload)
}

func (o *GetFileVersionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /file/versions][%d] GetFileVersions default %s", o._statusCode, payload)
}

func (o *GetFileVersionsDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetFileVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

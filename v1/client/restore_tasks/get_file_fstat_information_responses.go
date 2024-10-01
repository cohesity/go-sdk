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

// GetFileFstatInformationReader is a Reader for the GetFileFstatInformation structure.
type GetFileFstatInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFileFstatInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFileFstatInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFileFstatInformationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFileFstatInformationOK creates a GetFileFstatInformationOK with default headers values
func NewGetFileFstatInformationOK() *GetFileFstatInformationOK {
	return &GetFileFstatInformationOK{}
}

/*
GetFileFstatInformationOK describes a response with status code 200, with default header values.

Success
*/
type GetFileFstatInformationOK struct {
	Payload *models.FileFstatResult
}

// IsSuccess returns true when this get file fstat information o k response has a 2xx status code
func (o *GetFileFstatInformationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get file fstat information o k response has a 3xx status code
func (o *GetFileFstatInformationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get file fstat information o k response has a 4xx status code
func (o *GetFileFstatInformationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get file fstat information o k response has a 5xx status code
func (o *GetFileFstatInformationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get file fstat information o k response a status code equal to that given
func (o *GetFileFstatInformationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get file fstat information o k response
func (o *GetFileFstatInformationOK) Code() int {
	return 200
}

func (o *GetFileFstatInformationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/files/fstats][%d] getFileFstatInformationOK %s", 200, payload)
}

func (o *GetFileFstatInformationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/files/fstats][%d] getFileFstatInformationOK %s", 200, payload)
}

func (o *GetFileFstatInformationOK) GetPayload() *models.FileFstatResult {
	return o.Payload
}

func (o *GetFileFstatInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileFstatResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFileFstatInformationDefault creates a GetFileFstatInformationDefault with default headers values
func NewGetFileFstatInformationDefault(code int) *GetFileFstatInformationDefault {
	return &GetFileFstatInformationDefault{
		_statusCode: code,
	}
}

/*
GetFileFstatInformationDefault describes a response with status code -1, with default header values.

Error
*/
type GetFileFstatInformationDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get file fstat information default response has a 2xx status code
func (o *GetFileFstatInformationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get file fstat information default response has a 3xx status code
func (o *GetFileFstatInformationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get file fstat information default response has a 4xx status code
func (o *GetFileFstatInformationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get file fstat information default response has a 5xx status code
func (o *GetFileFstatInformationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get file fstat information default response a status code equal to that given
func (o *GetFileFstatInformationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get file fstat information default response
func (o *GetFileFstatInformationDefault) Code() int {
	return o._statusCode
}

func (o *GetFileFstatInformationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/files/fstats][%d] GetFileFstatInformation default %s", o._statusCode, payload)
}

func (o *GetFileFstatInformationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/restore/files/fstats][%d] GetFileFstatInformation default %s", o._statusCode, payload)
}

func (o *GetFileFstatInformationDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetFileFstatInformationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

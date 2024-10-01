// Code generated by go-swagger; DO NOT EDIT.

package protection_jobs

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

// GetProtectionJobByIDReader is a Reader for the GetProtectionJobByID structure.
type GetProtectionJobByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProtectionJobByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProtectionJobByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProtectionJobByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProtectionJobByIDOK creates a GetProtectionJobByIDOK with default headers values
func NewGetProtectionJobByIDOK() *GetProtectionJobByIDOK {
	return &GetProtectionJobByIDOK{}
}

/*
GetProtectionJobByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetProtectionJobByIDOK struct {
	Payload *models.ProtectionJob
}

// IsSuccess returns true when this get protection job by Id o k response has a 2xx status code
func (o *GetProtectionJobByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get protection job by Id o k response has a 3xx status code
func (o *GetProtectionJobByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get protection job by Id o k response has a 4xx status code
func (o *GetProtectionJobByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get protection job by Id o k response has a 5xx status code
func (o *GetProtectionJobByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get protection job by Id o k response a status code equal to that given
func (o *GetProtectionJobByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get protection job by Id o k response
func (o *GetProtectionJobByIDOK) Code() int {
	return 200
}

func (o *GetProtectionJobByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}][%d] getProtectionJobByIdOK %s", 200, payload)
}

func (o *GetProtectionJobByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}][%d] getProtectionJobByIdOK %s", 200, payload)
}

func (o *GetProtectionJobByIDOK) GetPayload() *models.ProtectionJob {
	return o.Payload
}

func (o *GetProtectionJobByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProtectionJobByIDDefault creates a GetProtectionJobByIDDefault with default headers values
func NewGetProtectionJobByIDDefault(code int) *GetProtectionJobByIDDefault {
	return &GetProtectionJobByIDDefault{
		_statusCode: code,
	}
}

/*
GetProtectionJobByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetProtectionJobByIDDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get protection job by Id default response has a 2xx status code
func (o *GetProtectionJobByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get protection job by Id default response has a 3xx status code
func (o *GetProtectionJobByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get protection job by Id default response has a 4xx status code
func (o *GetProtectionJobByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get protection job by Id default response has a 5xx status code
func (o *GetProtectionJobByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get protection job by Id default response a status code equal to that given
func (o *GetProtectionJobByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get protection job by Id default response
func (o *GetProtectionJobByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetProtectionJobByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}][%d] GetProtectionJobById default %s", o._statusCode, payload)
}

func (o *GetProtectionJobByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}][%d] GetProtectionJobById default %s", o._statusCode, payload)
}

func (o *GetProtectionJobByIDDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetProtectionJobByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetProtectionJobAuditReader is a Reader for the GetProtectionJobAudit structure.
type GetProtectionJobAuditReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProtectionJobAuditReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProtectionJobAuditOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProtectionJobAuditDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProtectionJobAuditOK creates a GetProtectionJobAuditOK with default headers values
func NewGetProtectionJobAuditOK() *GetProtectionJobAuditOK {
	return &GetProtectionJobAuditOK{}
}

/*
GetProtectionJobAuditOK describes a response with status code 200, with default header values.

Success
*/
type GetProtectionJobAuditOK struct {
	Payload []*models.ProtectionJobAuditTrail
}

// IsSuccess returns true when this get protection job audit o k response has a 2xx status code
func (o *GetProtectionJobAuditOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get protection job audit o k response has a 3xx status code
func (o *GetProtectionJobAuditOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get protection job audit o k response has a 4xx status code
func (o *GetProtectionJobAuditOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get protection job audit o k response has a 5xx status code
func (o *GetProtectionJobAuditOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get protection job audit o k response a status code equal to that given
func (o *GetProtectionJobAuditOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get protection job audit o k response
func (o *GetProtectionJobAuditOK) Code() int {
	return 200
}

func (o *GetProtectionJobAuditOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}/auditTrail][%d] getProtectionJobAuditOK %s", 200, payload)
}

func (o *GetProtectionJobAuditOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}/auditTrail][%d] getProtectionJobAuditOK %s", 200, payload)
}

func (o *GetProtectionJobAuditOK) GetPayload() []*models.ProtectionJobAuditTrail {
	return o.Payload
}

func (o *GetProtectionJobAuditOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProtectionJobAuditDefault creates a GetProtectionJobAuditDefault with default headers values
func NewGetProtectionJobAuditDefault(code int) *GetProtectionJobAuditDefault {
	return &GetProtectionJobAuditDefault{
		_statusCode: code,
	}
}

/*
GetProtectionJobAuditDefault describes a response with status code -1, with default header values.

Error
*/
type GetProtectionJobAuditDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get protection job audit default response has a 2xx status code
func (o *GetProtectionJobAuditDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get protection job audit default response has a 3xx status code
func (o *GetProtectionJobAuditDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get protection job audit default response has a 4xx status code
func (o *GetProtectionJobAuditDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get protection job audit default response has a 5xx status code
func (o *GetProtectionJobAuditDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get protection job audit default response a status code equal to that given
func (o *GetProtectionJobAuditDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get protection job audit default response
func (o *GetProtectionJobAuditDefault) Code() int {
	return o._statusCode
}

func (o *GetProtectionJobAuditDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}/auditTrail][%d] GetProtectionJobAudit default %s", o._statusCode, payload)
}

func (o *GetProtectionJobAuditDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/protectionJobs/{id}/auditTrail][%d] GetProtectionJobAudit default %s", o._statusCode, payload)
}

func (o *GetProtectionJobAuditDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetProtectionJobAuditDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

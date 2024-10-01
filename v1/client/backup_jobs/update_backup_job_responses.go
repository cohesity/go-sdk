// Code generated by go-swagger; DO NOT EDIT.

package backup_jobs

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

// UpdateBackupJobReader is a Reader for the UpdateBackupJob structure.
type UpdateBackupJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateBackupJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateBackupJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateBackupJobDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateBackupJobOK creates a UpdateBackupJobOK with default headers values
func NewUpdateBackupJobOK() *UpdateBackupJobOK {
	return &UpdateBackupJobOK{}
}

/*
UpdateBackupJobOK describes a response with status code 200, with default header values.

Success
*/
type UpdateBackupJobOK struct {
	Payload *models.BackupJobWrapper
}

// IsSuccess returns true when this update backup job o k response has a 2xx status code
func (o *UpdateBackupJobOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update backup job o k response has a 3xx status code
func (o *UpdateBackupJobOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update backup job o k response has a 4xx status code
func (o *UpdateBackupJobOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update backup job o k response has a 5xx status code
func (o *UpdateBackupJobOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update backup job o k response a status code equal to that given
func (o *UpdateBackupJobOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update backup job o k response
func (o *UpdateBackupJobOK) Code() int {
	return 200
}

func (o *UpdateBackupJobOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /backupjobs/{id}][%d] updateBackupJobOK %s", 200, payload)
}

func (o *UpdateBackupJobOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /backupjobs/{id}][%d] updateBackupJobOK %s", 200, payload)
}

func (o *UpdateBackupJobOK) GetPayload() *models.BackupJobWrapper {
	return o.Payload
}

func (o *UpdateBackupJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupJobWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateBackupJobDefault creates a UpdateBackupJobDefault with default headers values
func NewUpdateBackupJobDefault(code int) *UpdateBackupJobDefault {
	return &UpdateBackupJobDefault{
		_statusCode: code,
	}
}

/*
UpdateBackupJobDefault describes a response with status code -1, with default header values.

(empty)
*/
type UpdateBackupJobDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this update backup job default response has a 2xx status code
func (o *UpdateBackupJobDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update backup job default response has a 3xx status code
func (o *UpdateBackupJobDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update backup job default response has a 4xx status code
func (o *UpdateBackupJobDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update backup job default response has a 5xx status code
func (o *UpdateBackupJobDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update backup job default response a status code equal to that given
func (o *UpdateBackupJobDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update backup job default response
func (o *UpdateBackupJobDefault) Code() int {
	return o._statusCode
}

func (o *UpdateBackupJobDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /backupjobs/{id}][%d] UpdateBackupJob default %s", o._statusCode, payload)
}

func (o *UpdateBackupJobDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /backupjobs/{id}][%d] UpdateBackupJob default %s", o._statusCode, payload)
}

func (o *UpdateBackupJobDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *UpdateBackupJobDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package failover

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// SourceBackupDeactivationReader is a Reader for the SourceBackupDeactivation structure.
type SourceBackupDeactivationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SourceBackupDeactivationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSourceBackupDeactivationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSourceBackupDeactivationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSourceBackupDeactivationCreated creates a SourceBackupDeactivationCreated with default headers values
func NewSourceBackupDeactivationCreated() *SourceBackupDeactivationCreated {
	return &SourceBackupDeactivationCreated{}
}

/*
SourceBackupDeactivationCreated describes a response with status code 201, with default header values.

No Content
*/
type SourceBackupDeactivationCreated struct {
}

// IsSuccess returns true when this source backup deactivation created response has a 2xx status code
func (o *SourceBackupDeactivationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this source backup deactivation created response has a 3xx status code
func (o *SourceBackupDeactivationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this source backup deactivation created response has a 4xx status code
func (o *SourceBackupDeactivationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this source backup deactivation created response has a 5xx status code
func (o *SourceBackupDeactivationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this source backup deactivation created response a status code equal to that given
func (o *SourceBackupDeactivationCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the source backup deactivation created response
func (o *SourceBackupDeactivationCreated) Code() int {
	return 201
}

func (o *SourceBackupDeactivationCreated) Error() string {
	return fmt.Sprintf("[POST /data-protect/failover/{id}/backup-deactivation][%d] sourceBackupDeactivationCreated", 201)
}

func (o *SourceBackupDeactivationCreated) String() string {
	return fmt.Sprintf("[POST /data-protect/failover/{id}/backup-deactivation][%d] sourceBackupDeactivationCreated", 201)
}

func (o *SourceBackupDeactivationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSourceBackupDeactivationDefault creates a SourceBackupDeactivationDefault with default headers values
func NewSourceBackupDeactivationDefault(code int) *SourceBackupDeactivationDefault {
	return &SourceBackupDeactivationDefault{
		_statusCode: code,
	}
}

/*
SourceBackupDeactivationDefault describes a response with status code -1, with default header values.

Error
*/
type SourceBackupDeactivationDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this source backup deactivation default response has a 2xx status code
func (o *SourceBackupDeactivationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this source backup deactivation default response has a 3xx status code
func (o *SourceBackupDeactivationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this source backup deactivation default response has a 4xx status code
func (o *SourceBackupDeactivationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this source backup deactivation default response has a 5xx status code
func (o *SourceBackupDeactivationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this source backup deactivation default response a status code equal to that given
func (o *SourceBackupDeactivationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the source backup deactivation default response
func (o *SourceBackupDeactivationDefault) Code() int {
	return o._statusCode
}

func (o *SourceBackupDeactivationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/failover/{id}/backup-deactivation][%d] SourceBackupDeactivation default %s", o._statusCode, payload)
}

func (o *SourceBackupDeactivationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/failover/{id}/backup-deactivation][%d] SourceBackupDeactivation default %s", o._statusCode, payload)
}

func (o *SourceBackupDeactivationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *SourceBackupDeactivationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package antivirus_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v2/models"
)

// DeleteInfectedFilesReader is a Reader for the DeleteInfectedFiles structure.
type DeleteInfectedFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInfectedFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteInfectedFilesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteInfectedFilesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteInfectedFilesCreated creates a DeleteInfectedFilesCreated with default headers values
func NewDeleteInfectedFilesCreated() *DeleteInfectedFilesCreated {
	return &DeleteInfectedFilesCreated{}
}

/*
DeleteInfectedFilesCreated describes a response with status code 201, with default header values.

Success
*/
type DeleteInfectedFilesCreated struct {
	Payload *models.DeleteInfectedFiles
}

// IsSuccess returns true when this delete infected files created response has a 2xx status code
func (o *DeleteInfectedFilesCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete infected files created response has a 3xx status code
func (o *DeleteInfectedFilesCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete infected files created response has a 4xx status code
func (o *DeleteInfectedFilesCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete infected files created response has a 5xx status code
func (o *DeleteInfectedFilesCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this delete infected files created response a status code equal to that given
func (o *DeleteInfectedFilesCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the delete infected files created response
func (o *DeleteInfectedFilesCreated) Code() int {
	return 201
}

func (o *DeleteInfectedFilesCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /antivirus-service/infected-files][%d] deleteInfectedFilesCreated %s", 201, payload)
}

func (o *DeleteInfectedFilesCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /antivirus-service/infected-files][%d] deleteInfectedFilesCreated %s", 201, payload)
}

func (o *DeleteInfectedFilesCreated) GetPayload() *models.DeleteInfectedFiles {
	return o.Payload
}

func (o *DeleteInfectedFilesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeleteInfectedFiles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteInfectedFilesDefault creates a DeleteInfectedFilesDefault with default headers values
func NewDeleteInfectedFilesDefault(code int) *DeleteInfectedFilesDefault {
	return &DeleteInfectedFilesDefault{
		_statusCode: code,
	}
}

/*
DeleteInfectedFilesDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteInfectedFilesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete infected files default response has a 2xx status code
func (o *DeleteInfectedFilesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete infected files default response has a 3xx status code
func (o *DeleteInfectedFilesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete infected files default response has a 4xx status code
func (o *DeleteInfectedFilesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete infected files default response has a 5xx status code
func (o *DeleteInfectedFilesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete infected files default response a status code equal to that given
func (o *DeleteInfectedFilesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete infected files default response
func (o *DeleteInfectedFilesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInfectedFilesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /antivirus-service/infected-files][%d] DeleteInfectedFiles default %s", o._statusCode, payload)
}

func (o *DeleteInfectedFilesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /antivirus-service/infected-files][%d] DeleteInfectedFiles default %s", o._statusCode, payload)
}

func (o *DeleteInfectedFilesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteInfectedFilesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

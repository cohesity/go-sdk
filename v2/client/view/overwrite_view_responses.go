// Code generated by go-swagger; DO NOT EDIT.

package view

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

// OverwriteViewReader is a Reader for the OverwriteView structure.
type OverwriteViewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OverwriteViewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOverwriteViewNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOverwriteViewDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOverwriteViewNoContent creates a OverwriteViewNoContent with default headers values
func NewOverwriteViewNoContent() *OverwriteViewNoContent {
	return &OverwriteViewNoContent{}
}

/*
OverwriteViewNoContent describes a response with status code 204, with default header values.

No Content
*/
type OverwriteViewNoContent struct {
}

// IsSuccess returns true when this overwrite view no content response has a 2xx status code
func (o *OverwriteViewNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this overwrite view no content response has a 3xx status code
func (o *OverwriteViewNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this overwrite view no content response has a 4xx status code
func (o *OverwriteViewNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this overwrite view no content response has a 5xx status code
func (o *OverwriteViewNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this overwrite view no content response a status code equal to that given
func (o *OverwriteViewNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the overwrite view no content response
func (o *OverwriteViewNoContent) Code() int {
	return 204
}

func (o *OverwriteViewNoContent) Error() string {
	return fmt.Sprintf("[POST /file-services/views/{id}/overwrite][%d] overwriteViewNoContent", 204)
}

func (o *OverwriteViewNoContent) String() string {
	return fmt.Sprintf("[POST /file-services/views/{id}/overwrite][%d] overwriteViewNoContent", 204)
}

func (o *OverwriteViewNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOverwriteViewDefault creates a OverwriteViewDefault with default headers values
func NewOverwriteViewDefault(code int) *OverwriteViewDefault {
	return &OverwriteViewDefault{
		_statusCode: code,
	}
}

/*
OverwriteViewDefault describes a response with status code -1, with default header values.

Error
*/
type OverwriteViewDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this overwrite view default response has a 2xx status code
func (o *OverwriteViewDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this overwrite view default response has a 3xx status code
func (o *OverwriteViewDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this overwrite view default response has a 4xx status code
func (o *OverwriteViewDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this overwrite view default response has a 5xx status code
func (o *OverwriteViewDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this overwrite view default response a status code equal to that given
func (o *OverwriteViewDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the overwrite view default response
func (o *OverwriteViewDefault) Code() int {
	return o._statusCode
}

func (o *OverwriteViewDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file-services/views/{id}/overwrite][%d] OverwriteView default %s", o._statusCode, payload)
}

func (o *OverwriteViewDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /file-services/views/{id}/overwrite][%d] OverwriteView default %s", o._statusCode, payload)
}

func (o *OverwriteViewDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *OverwriteViewDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

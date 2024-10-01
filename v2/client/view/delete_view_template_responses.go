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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// DeleteViewTemplateReader is a Reader for the DeleteViewTemplate structure.
type DeleteViewTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteViewTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteViewTemplateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteViewTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteViewTemplateNoContent creates a DeleteViewTemplateNoContent with default headers values
func NewDeleteViewTemplateNoContent() *DeleteViewTemplateNoContent {
	return &DeleteViewTemplateNoContent{}
}

/*
DeleteViewTemplateNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteViewTemplateNoContent struct {
}

// IsSuccess returns true when this delete view template no content response has a 2xx status code
func (o *DeleteViewTemplateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete view template no content response has a 3xx status code
func (o *DeleteViewTemplateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete view template no content response has a 4xx status code
func (o *DeleteViewTemplateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete view template no content response has a 5xx status code
func (o *DeleteViewTemplateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete view template no content response a status code equal to that given
func (o *DeleteViewTemplateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete view template no content response
func (o *DeleteViewTemplateNoContent) Code() int {
	return 204
}

func (o *DeleteViewTemplateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /file-services/view-template/{id}][%d] deleteViewTemplateNoContent", 204)
}

func (o *DeleteViewTemplateNoContent) String() string {
	return fmt.Sprintf("[DELETE /file-services/view-template/{id}][%d] deleteViewTemplateNoContent", 204)
}

func (o *DeleteViewTemplateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteViewTemplateDefault creates a DeleteViewTemplateDefault with default headers values
func NewDeleteViewTemplateDefault(code int) *DeleteViewTemplateDefault {
	return &DeleteViewTemplateDefault{
		_statusCode: code,
	}
}

/*
DeleteViewTemplateDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteViewTemplateDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete view template default response has a 2xx status code
func (o *DeleteViewTemplateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete view template default response has a 3xx status code
func (o *DeleteViewTemplateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete view template default response has a 4xx status code
func (o *DeleteViewTemplateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete view template default response has a 5xx status code
func (o *DeleteViewTemplateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete view template default response a status code equal to that given
func (o *DeleteViewTemplateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete view template default response
func (o *DeleteViewTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteViewTemplateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /file-services/view-template/{id}][%d] DeleteViewTemplate default %s", o._statusCode, payload)
}

func (o *DeleteViewTemplateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /file-services/view-template/{id}][%d] DeleteViewTemplate default %s", o._statusCode, payload)
}

func (o *DeleteViewTemplateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteViewTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

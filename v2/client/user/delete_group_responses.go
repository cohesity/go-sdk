// Code generated by go-swagger; DO NOT EDIT.

package user

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

// DeleteGroupReader is a Reader for the DeleteGroup structure.
type DeleteGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGroupNoContent creates a DeleteGroupNoContent with default headers values
func NewDeleteGroupNoContent() *DeleteGroupNoContent {
	return &DeleteGroupNoContent{}
}

/*
DeleteGroupNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteGroupNoContent struct {
}

// IsSuccess returns true when this delete group no content response has a 2xx status code
func (o *DeleteGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete group no content response has a 3xx status code
func (o *DeleteGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete group no content response has a 4xx status code
func (o *DeleteGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete group no content response has a 5xx status code
func (o *DeleteGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete group no content response a status code equal to that given
func (o *DeleteGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete group no content response
func (o *DeleteGroupNoContent) Code() int {
	return 204
}

func (o *DeleteGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /groups/{sid}][%d] deleteGroupNoContent", 204)
}

func (o *DeleteGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /groups/{sid}][%d] deleteGroupNoContent", 204)
}

func (o *DeleteGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupDefault creates a DeleteGroupDefault with default headers values
func NewDeleteGroupDefault(code int) *DeleteGroupDefault {
	return &DeleteGroupDefault{
		_statusCode: code,
	}
}

/*
DeleteGroupDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteGroupDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete group default response has a 2xx status code
func (o *DeleteGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete group default response has a 3xx status code
func (o *DeleteGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete group default response has a 4xx status code
func (o *DeleteGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete group default response has a 5xx status code
func (o *DeleteGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete group default response a status code equal to that given
func (o *DeleteGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete group default response
func (o *DeleteGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groups/{sid}][%d] DeleteGroup default %s", o._statusCode, payload)
}

func (o *DeleteGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /groups/{sid}][%d] DeleteGroup default %s", o._statusCode, payload)
}

func (o *DeleteGroupDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package principals

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

// DeleteUserAPIKeyReader is a Reader for the DeleteUserAPIKey structure.
type DeleteUserAPIKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserAPIKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserAPIKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteUserAPIKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserAPIKeyNoContent creates a DeleteUserAPIKeyNoContent with default headers values
func NewDeleteUserAPIKeyNoContent() *DeleteUserAPIKeyNoContent {
	return &DeleteUserAPIKeyNoContent{}
}

/*
DeleteUserAPIKeyNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteUserAPIKeyNoContent struct {
}

// IsSuccess returns true when this delete user Api key no content response has a 2xx status code
func (o *DeleteUserAPIKeyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user Api key no content response has a 3xx status code
func (o *DeleteUserAPIKeyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user Api key no content response has a 4xx status code
func (o *DeleteUserAPIKeyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user Api key no content response has a 5xx status code
func (o *DeleteUserAPIKeyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user Api key no content response a status code equal to that given
func (o *DeleteUserAPIKeyNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete user Api key no content response
func (o *DeleteUserAPIKeyNoContent) Code() int {
	return 204
}

func (o *DeleteUserAPIKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/users/{sid}/apiKeys/{id}][%d] deleteUserApiKeyNoContent", 204)
}

func (o *DeleteUserAPIKeyNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/users/{sid}/apiKeys/{id}][%d] deleteUserApiKeyNoContent", 204)
}

func (o *DeleteUserAPIKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserAPIKeyDefault creates a DeleteUserAPIKeyDefault with default headers values
func NewDeleteUserAPIKeyDefault(code int) *DeleteUserAPIKeyDefault {
	return &DeleteUserAPIKeyDefault{
		_statusCode: code,
	}
}

/*
DeleteUserAPIKeyDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteUserAPIKeyDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this delete user Api key default response has a 2xx status code
func (o *DeleteUserAPIKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete user Api key default response has a 3xx status code
func (o *DeleteUserAPIKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete user Api key default response has a 4xx status code
func (o *DeleteUserAPIKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete user Api key default response has a 5xx status code
func (o *DeleteUserAPIKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete user Api key default response a status code equal to that given
func (o *DeleteUserAPIKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete user Api key default response
func (o *DeleteUserAPIKeyDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserAPIKeyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/users/{sid}/apiKeys/{id}][%d] DeleteUserApiKey default %s", o._statusCode, payload)
}

func (o *DeleteUserAPIKeyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/users/{sid}/apiKeys/{id}][%d] DeleteUserApiKey default %s", o._statusCode, payload)
}

func (o *DeleteUserAPIKeyDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DeleteUserAPIKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

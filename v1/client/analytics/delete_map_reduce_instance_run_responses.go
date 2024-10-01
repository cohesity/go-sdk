// Code generated by go-swagger; DO NOT EDIT.

package analytics

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

// DeleteMapReduceInstanceRunReader is a Reader for the DeleteMapReduceInstanceRun structure.
type DeleteMapReduceInstanceRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMapReduceInstanceRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMapReduceInstanceRunNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteMapReduceInstanceRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteMapReduceInstanceRunNoContent creates a DeleteMapReduceInstanceRunNoContent with default headers values
func NewDeleteMapReduceInstanceRunNoContent() *DeleteMapReduceInstanceRunNoContent {
	return &DeleteMapReduceInstanceRunNoContent{}
}

/*
DeleteMapReduceInstanceRunNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteMapReduceInstanceRunNoContent struct {
}

// IsSuccess returns true when this delete map reduce instance run no content response has a 2xx status code
func (o *DeleteMapReduceInstanceRunNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete map reduce instance run no content response has a 3xx status code
func (o *DeleteMapReduceInstanceRunNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete map reduce instance run no content response has a 4xx status code
func (o *DeleteMapReduceInstanceRunNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete map reduce instance run no content response has a 5xx status code
func (o *DeleteMapReduceInstanceRunNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete map reduce instance run no content response a status code equal to that given
func (o *DeleteMapReduceInstanceRunNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete map reduce instance run no content response
func (o *DeleteMapReduceInstanceRunNoContent) Code() int {
	return 204
}

func (o *DeleteMapReduceInstanceRunNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/analytics/mrAppRun/{id}][%d] deleteMapReduceInstanceRunNoContent", 204)
}

func (o *DeleteMapReduceInstanceRunNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/analytics/mrAppRun/{id}][%d] deleteMapReduceInstanceRunNoContent", 204)
}

func (o *DeleteMapReduceInstanceRunNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMapReduceInstanceRunDefault creates a DeleteMapReduceInstanceRunDefault with default headers values
func NewDeleteMapReduceInstanceRunDefault(code int) *DeleteMapReduceInstanceRunDefault {
	return &DeleteMapReduceInstanceRunDefault{
		_statusCode: code,
	}
}

/*
DeleteMapReduceInstanceRunDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteMapReduceInstanceRunDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this delete map reduce instance run default response has a 2xx status code
func (o *DeleteMapReduceInstanceRunDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete map reduce instance run default response has a 3xx status code
func (o *DeleteMapReduceInstanceRunDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete map reduce instance run default response has a 4xx status code
func (o *DeleteMapReduceInstanceRunDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete map reduce instance run default response has a 5xx status code
func (o *DeleteMapReduceInstanceRunDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete map reduce instance run default response a status code equal to that given
func (o *DeleteMapReduceInstanceRunDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete map reduce instance run default response
func (o *DeleteMapReduceInstanceRunDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMapReduceInstanceRunDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/analytics/mrAppRun/{id}][%d] DeleteMapReduceInstanceRun default %s", o._statusCode, payload)
}

func (o *DeleteMapReduceInstanceRunDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/analytics/mrAppRun/{id}][%d] DeleteMapReduceInstanceRun default %s", o._statusCode, payload)
}

func (o *DeleteMapReduceInstanceRunDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DeleteMapReduceInstanceRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

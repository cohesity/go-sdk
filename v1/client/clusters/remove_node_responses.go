// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// RemoveNodeReader is a Reader for the RemoveNode structure.
type RemoveNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveNodeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveNodeNoContent creates a RemoveNodeNoContent with default headers values
func NewRemoveNodeNoContent() *RemoveNodeNoContent {
	return &RemoveNodeNoContent{}
}

/*
RemoveNodeNoContent describes a response with status code 204, with default header values.

No Content
*/
type RemoveNodeNoContent struct {
}

// IsSuccess returns true when this remove node no content response has a 2xx status code
func (o *RemoveNodeNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove node no content response has a 3xx status code
func (o *RemoveNodeNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove node no content response has a 4xx status code
func (o *RemoveNodeNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove node no content response has a 5xx status code
func (o *RemoveNodeNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this remove node no content response a status code equal to that given
func (o *RemoveNodeNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the remove node no content response
func (o *RemoveNodeNoContent) Code() int {
	return 204
}

func (o *RemoveNodeNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/clusters/nodes/{id}][%d] removeNodeNoContent", 204)
}

func (o *RemoveNodeNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/clusters/nodes/{id}][%d] removeNodeNoContent", 204)
}

func (o *RemoveNodeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveNodeDefault creates a RemoveNodeDefault with default headers values
func NewRemoveNodeDefault(code int) *RemoveNodeDefault {
	return &RemoveNodeDefault{
		_statusCode: code,
	}
}

/*
RemoveNodeDefault describes a response with status code -1, with default header values.

Error
*/
type RemoveNodeDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this remove node default response has a 2xx status code
func (o *RemoveNodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this remove node default response has a 3xx status code
func (o *RemoveNodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this remove node default response has a 4xx status code
func (o *RemoveNodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this remove node default response has a 5xx status code
func (o *RemoveNodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this remove node default response a status code equal to that given
func (o *RemoveNodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the remove node default response
func (o *RemoveNodeDefault) Code() int {
	return o._statusCode
}

func (o *RemoveNodeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/clusters/nodes/{id}][%d] RemoveNode default %s", o._statusCode, payload)
}

func (o *RemoveNodeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/clusters/nodes/{id}][%d] RemoveNode default %s", o._statusCode, payload)
}

func (o *RemoveNodeDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *RemoveNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

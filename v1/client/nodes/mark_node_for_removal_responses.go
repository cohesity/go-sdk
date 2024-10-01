// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// MarkNodeForRemovalReader is a Reader for the MarkNodeForRemoval structure.
type MarkNodeForRemovalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MarkNodeForRemovalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMarkNodeForRemovalAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMarkNodeForRemovalDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMarkNodeForRemovalAccepted creates a MarkNodeForRemovalAccepted with default headers values
func NewMarkNodeForRemovalAccepted() *MarkNodeForRemovalAccepted {
	return &MarkNodeForRemovalAccepted{}
}

/*
MarkNodeForRemovalAccepted describes a response with status code 202, with default header values.

Success
*/
type MarkNodeForRemovalAccepted struct {
	Payload *models.NodeDeleteResult
}

// IsSuccess returns true when this mark node for removal accepted response has a 2xx status code
func (o *MarkNodeForRemovalAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this mark node for removal accepted response has a 3xx status code
func (o *MarkNodeForRemovalAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this mark node for removal accepted response has a 4xx status code
func (o *MarkNodeForRemovalAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this mark node for removal accepted response has a 5xx status code
func (o *MarkNodeForRemovalAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this mark node for removal accepted response a status code equal to that given
func (o *MarkNodeForRemovalAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the mark node for removal accepted response
func (o *MarkNodeForRemovalAccepted) Code() int {
	return 202
}

func (o *MarkNodeForRemovalAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/{id}][%d] markNodeForRemovalAccepted %s", 202, payload)
}

func (o *MarkNodeForRemovalAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/{id}][%d] markNodeForRemovalAccepted %s", 202, payload)
}

func (o *MarkNodeForRemovalAccepted) GetPayload() *models.NodeDeleteResult {
	return o.Payload
}

func (o *MarkNodeForRemovalAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeDeleteResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMarkNodeForRemovalDefault creates a MarkNodeForRemovalDefault with default headers values
func NewMarkNodeForRemovalDefault(code int) *MarkNodeForRemovalDefault {
	return &MarkNodeForRemovalDefault{
		_statusCode: code,
	}
}

/*
MarkNodeForRemovalDefault describes a response with status code -1, with default header values.

(empty)
*/
type MarkNodeForRemovalDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this mark node for removal default response has a 2xx status code
func (o *MarkNodeForRemovalDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this mark node for removal default response has a 3xx status code
func (o *MarkNodeForRemovalDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this mark node for removal default response has a 4xx status code
func (o *MarkNodeForRemovalDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this mark node for removal default response has a 5xx status code
func (o *MarkNodeForRemovalDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this mark node for removal default response a status code equal to that given
func (o *MarkNodeForRemovalDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the mark node for removal default response
func (o *MarkNodeForRemovalDefault) Code() int {
	return o._statusCode
}

func (o *MarkNodeForRemovalDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/{id}][%d] MarkNodeForRemoval default %s", o._statusCode, payload)
}

func (o *MarkNodeForRemovalDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /nodes/{id}][%d] MarkNodeForRemoval default %s", o._statusCode, payload)
}

func (o *MarkNodeForRemovalDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *MarkNodeForRemovalDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

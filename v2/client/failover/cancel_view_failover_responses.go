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

	"github.com/cohesity/go-sdk/v2/models"
)

// CancelViewFailoverReader is a Reader for the CancelViewFailover structure.
type CancelViewFailoverReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelViewFailoverReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCancelViewFailoverNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCancelViewFailoverDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCancelViewFailoverNoContent creates a CancelViewFailoverNoContent with default headers values
func NewCancelViewFailoverNoContent() *CancelViewFailoverNoContent {
	return &CancelViewFailoverNoContent{}
}

/*
CancelViewFailoverNoContent describes a response with status code 204, with default header values.

No Content
*/
type CancelViewFailoverNoContent struct {
}

// IsSuccess returns true when this cancel view failover no content response has a 2xx status code
func (o *CancelViewFailoverNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cancel view failover no content response has a 3xx status code
func (o *CancelViewFailoverNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cancel view failover no content response has a 4xx status code
func (o *CancelViewFailoverNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this cancel view failover no content response has a 5xx status code
func (o *CancelViewFailoverNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this cancel view failover no content response a status code equal to that given
func (o *CancelViewFailoverNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the cancel view failover no content response
func (o *CancelViewFailoverNoContent) Code() int {
	return 204
}

func (o *CancelViewFailoverNoContent) Error() string {
	return fmt.Sprintf("[POST /data-protect/failover/views/{id}/cancel][%d] cancelViewFailoverNoContent", 204)
}

func (o *CancelViewFailoverNoContent) String() string {
	return fmt.Sprintf("[POST /data-protect/failover/views/{id}/cancel][%d] cancelViewFailoverNoContent", 204)
}

func (o *CancelViewFailoverNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelViewFailoverDefault creates a CancelViewFailoverDefault with default headers values
func NewCancelViewFailoverDefault(code int) *CancelViewFailoverDefault {
	return &CancelViewFailoverDefault{
		_statusCode: code,
	}
}

/*
CancelViewFailoverDefault describes a response with status code -1, with default header values.

Error
*/
type CancelViewFailoverDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this cancel view failover default response has a 2xx status code
func (o *CancelViewFailoverDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cancel view failover default response has a 3xx status code
func (o *CancelViewFailoverDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cancel view failover default response has a 4xx status code
func (o *CancelViewFailoverDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cancel view failover default response has a 5xx status code
func (o *CancelViewFailoverDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cancel view failover default response a status code equal to that given
func (o *CancelViewFailoverDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cancel view failover default response
func (o *CancelViewFailoverDefault) Code() int {
	return o._statusCode
}

func (o *CancelViewFailoverDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/failover/views/{id}/cancel][%d] CancelViewFailover default %s", o._statusCode, payload)
}

func (o *CancelViewFailoverDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/failover/views/{id}/cancel][%d] CancelViewFailover default %s", o._statusCode, payload)
}

func (o *CancelViewFailoverDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *CancelViewFailoverDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

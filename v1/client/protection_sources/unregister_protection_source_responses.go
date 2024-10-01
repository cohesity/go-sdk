// Code generated by go-swagger; DO NOT EDIT.

package protection_sources

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

// UnregisterProtectionSourceReader is a Reader for the UnregisterProtectionSource structure.
type UnregisterProtectionSourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnregisterProtectionSourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnregisterProtectionSourceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnregisterProtectionSourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnregisterProtectionSourceNoContent creates a UnregisterProtectionSourceNoContent with default headers values
func NewUnregisterProtectionSourceNoContent() *UnregisterProtectionSourceNoContent {
	return &UnregisterProtectionSourceNoContent{}
}

/*
UnregisterProtectionSourceNoContent describes a response with status code 204, with default header values.

No Content
*/
type UnregisterProtectionSourceNoContent struct {
}

// IsSuccess returns true when this unregister protection source no content response has a 2xx status code
func (o *UnregisterProtectionSourceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unregister protection source no content response has a 3xx status code
func (o *UnregisterProtectionSourceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unregister protection source no content response has a 4xx status code
func (o *UnregisterProtectionSourceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this unregister protection source no content response has a 5xx status code
func (o *UnregisterProtectionSourceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this unregister protection source no content response a status code equal to that given
func (o *UnregisterProtectionSourceNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the unregister protection source no content response
func (o *UnregisterProtectionSourceNoContent) Code() int {
	return 204
}

func (o *UnregisterProtectionSourceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/protectionSources/{id}][%d] unregisterProtectionSourceNoContent", 204)
}

func (o *UnregisterProtectionSourceNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/protectionSources/{id}][%d] unregisterProtectionSourceNoContent", 204)
}

func (o *UnregisterProtectionSourceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnregisterProtectionSourceDefault creates a UnregisterProtectionSourceDefault with default headers values
func NewUnregisterProtectionSourceDefault(code int) *UnregisterProtectionSourceDefault {
	return &UnregisterProtectionSourceDefault{
		_statusCode: code,
	}
}

/*
UnregisterProtectionSourceDefault describes a response with status code -1, with default header values.

Error
*/
type UnregisterProtectionSourceDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this unregister protection source default response has a 2xx status code
func (o *UnregisterProtectionSourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unregister protection source default response has a 3xx status code
func (o *UnregisterProtectionSourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unregister protection source default response has a 4xx status code
func (o *UnregisterProtectionSourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unregister protection source default response has a 5xx status code
func (o *UnregisterProtectionSourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unregister protection source default response a status code equal to that given
func (o *UnregisterProtectionSourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unregister protection source default response
func (o *UnregisterProtectionSourceDefault) Code() int {
	return o._statusCode
}

func (o *UnregisterProtectionSourceDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/protectionSources/{id}][%d] UnregisterProtectionSource default %s", o._statusCode, payload)
}

func (o *UnregisterProtectionSourceDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/protectionSources/{id}][%d] UnregisterProtectionSource default %s", o._statusCode, payload)
}

func (o *UnregisterProtectionSourceDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UnregisterProtectionSourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

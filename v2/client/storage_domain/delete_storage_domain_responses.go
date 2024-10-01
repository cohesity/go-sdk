// Code generated by go-swagger; DO NOT EDIT.

package storage_domain

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

// DeleteStorageDomainReader is a Reader for the DeleteStorageDomain structure.
type DeleteStorageDomainReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageDomainReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStorageDomainNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteStorageDomainDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteStorageDomainNoContent creates a DeleteStorageDomainNoContent with default headers values
func NewDeleteStorageDomainNoContent() *DeleteStorageDomainNoContent {
	return &DeleteStorageDomainNoContent{}
}

/*
DeleteStorageDomainNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteStorageDomainNoContent struct {
}

// IsSuccess returns true when this delete storage domain no content response has a 2xx status code
func (o *DeleteStorageDomainNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete storage domain no content response has a 3xx status code
func (o *DeleteStorageDomainNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete storage domain no content response has a 4xx status code
func (o *DeleteStorageDomainNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete storage domain no content response has a 5xx status code
func (o *DeleteStorageDomainNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete storage domain no content response a status code equal to that given
func (o *DeleteStorageDomainNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete storage domain no content response
func (o *DeleteStorageDomainNoContent) Code() int {
	return 204
}

func (o *DeleteStorageDomainNoContent) Error() string {
	return fmt.Sprintf("[DELETE /storage-domains/{id}][%d] deleteStorageDomainNoContent", 204)
}

func (o *DeleteStorageDomainNoContent) String() string {
	return fmt.Sprintf("[DELETE /storage-domains/{id}][%d] deleteStorageDomainNoContent", 204)
}

func (o *DeleteStorageDomainNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageDomainDefault creates a DeleteStorageDomainDefault with default headers values
func NewDeleteStorageDomainDefault(code int) *DeleteStorageDomainDefault {
	return &DeleteStorageDomainDefault{
		_statusCode: code,
	}
}

/*
DeleteStorageDomainDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteStorageDomainDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete storage domain default response has a 2xx status code
func (o *DeleteStorageDomainDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete storage domain default response has a 3xx status code
func (o *DeleteStorageDomainDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete storage domain default response has a 4xx status code
func (o *DeleteStorageDomainDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete storage domain default response has a 5xx status code
func (o *DeleteStorageDomainDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete storage domain default response a status code equal to that given
func (o *DeleteStorageDomainDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete storage domain default response
func (o *DeleteStorageDomainDefault) Code() int {
	return o._statusCode
}

func (o *DeleteStorageDomainDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage-domains/{id}][%d] DeleteStorageDomain default %s", o._statusCode, payload)
}

func (o *DeleteStorageDomainDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /storage-domains/{id}][%d] DeleteStorageDomain default %s", o._statusCode, payload)
}

func (o *DeleteStorageDomainDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteStorageDomainDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

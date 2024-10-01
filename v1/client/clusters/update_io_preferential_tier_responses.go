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

	"github.com/cohesity/go-sdk/v1/models"
)

// UpdateIoPreferentialTierReader is a Reader for the UpdateIoPreferentialTier structure.
type UpdateIoPreferentialTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIoPreferentialTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIoPreferentialTierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateIoPreferentialTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateIoPreferentialTierOK creates a UpdateIoPreferentialTierOK with default headers values
func NewUpdateIoPreferentialTierOK() *UpdateIoPreferentialTierOK {
	return &UpdateIoPreferentialTierOK{}
}

/*
UpdateIoPreferentialTierOK describes a response with status code 200, with default header values.

Success
*/
type UpdateIoPreferentialTierOK struct {
	Payload *models.PrivateIoPreferentialTier
}

// IsSuccess returns true when this update io preferential tier o k response has a 2xx status code
func (o *UpdateIoPreferentialTierOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update io preferential tier o k response has a 3xx status code
func (o *UpdateIoPreferentialTierOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update io preferential tier o k response has a 4xx status code
func (o *UpdateIoPreferentialTierOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update io preferential tier o k response has a 5xx status code
func (o *UpdateIoPreferentialTierOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update io preferential tier o k response a status code equal to that given
func (o *UpdateIoPreferentialTierOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update io preferential tier o k response
func (o *UpdateIoPreferentialTierOK) Code() int {
	return 200
}

func (o *UpdateIoPreferentialTierOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ioPreferentialTier][%d] updateIoPreferentialTierOK %s", 200, payload)
}

func (o *UpdateIoPreferentialTierOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ioPreferentialTier][%d] updateIoPreferentialTierOK %s", 200, payload)
}

func (o *UpdateIoPreferentialTierOK) GetPayload() *models.PrivateIoPreferentialTier {
	return o.Payload
}

func (o *UpdateIoPreferentialTierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateIoPreferentialTier)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIoPreferentialTierDefault creates a UpdateIoPreferentialTierDefault with default headers values
func NewUpdateIoPreferentialTierDefault(code int) *UpdateIoPreferentialTierDefault {
	return &UpdateIoPreferentialTierDefault{
		_statusCode: code,
	}
}

/*
UpdateIoPreferentialTierDefault describes a response with status code -1, with default header values.

(empty)
*/
type UpdateIoPreferentialTierDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this update io preferential tier default response has a 2xx status code
func (o *UpdateIoPreferentialTierDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update io preferential tier default response has a 3xx status code
func (o *UpdateIoPreferentialTierDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update io preferential tier default response has a 4xx status code
func (o *UpdateIoPreferentialTierDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update io preferential tier default response has a 5xx status code
func (o *UpdateIoPreferentialTierDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update io preferential tier default response a status code equal to that given
func (o *UpdateIoPreferentialTierDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update io preferential tier default response
func (o *UpdateIoPreferentialTierDefault) Code() int {
	return o._statusCode
}

func (o *UpdateIoPreferentialTierDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ioPreferentialTier][%d] UpdateIoPreferentialTier default %s", o._statusCode, payload)
}

func (o *UpdateIoPreferentialTierDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ioPreferentialTier][%d] UpdateIoPreferentialTier default %s", o._statusCode, payload)
}

func (o *UpdateIoPreferentialTierDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *UpdateIoPreferentialTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

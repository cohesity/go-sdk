// Code generated by go-swagger; DO NOT EDIT.

package ip

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

// UnconfigureIPReader is a Reader for the UnconfigureIP structure.
type UnconfigureIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnconfigureIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUnconfigureIPNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnconfigureIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnconfigureIPNoContent creates a UnconfigureIPNoContent with default headers values
func NewUnconfigureIPNoContent() *UnconfigureIPNoContent {
	return &UnconfigureIPNoContent{}
}

/*
UnconfigureIPNoContent describes a response with status code 204, with default header values.

No Content
*/
type UnconfigureIPNoContent struct {
}

// IsSuccess returns true when this unconfigure Ip no content response has a 2xx status code
func (o *UnconfigureIPNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unconfigure Ip no content response has a 3xx status code
func (o *UnconfigureIPNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unconfigure Ip no content response has a 4xx status code
func (o *UnconfigureIPNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this unconfigure Ip no content response has a 5xx status code
func (o *UnconfigureIPNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this unconfigure Ip no content response a status code equal to that given
func (o *UnconfigureIPNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the unconfigure Ip no content response
func (o *UnconfigureIPNoContent) Code() int {
	return 204
}

func (o *UnconfigureIPNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/ip][%d] unconfigureIpNoContent", 204)
}

func (o *UnconfigureIPNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/ip][%d] unconfigureIpNoContent", 204)
}

func (o *UnconfigureIPNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnconfigureIPDefault creates a UnconfigureIPDefault with default headers values
func NewUnconfigureIPDefault(code int) *UnconfigureIPDefault {
	return &UnconfigureIPDefault{
		_statusCode: code,
	}
}

/*
UnconfigureIPDefault describes a response with status code -1, with default header values.

Error
*/
type UnconfigureIPDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this unconfigure Ip default response has a 2xx status code
func (o *UnconfigureIPDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unconfigure Ip default response has a 3xx status code
func (o *UnconfigureIPDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unconfigure Ip default response has a 4xx status code
func (o *UnconfigureIPDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unconfigure Ip default response has a 5xx status code
func (o *UnconfigureIPDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unconfigure Ip default response a status code equal to that given
func (o *UnconfigureIPDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the unconfigure Ip default response
func (o *UnconfigureIPDefault) Code() int {
	return o._statusCode
}

func (o *UnconfigureIPDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/ip][%d] UnconfigureIp default %s", o._statusCode, payload)
}

func (o *UnconfigureIPDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/ip][%d] UnconfigureIp default %s", o._statusCode, payload)
}

func (o *UnconfigureIPDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UnconfigureIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

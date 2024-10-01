// Code generated by go-swagger; DO NOT EDIT.

package platform

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

// GetIpmiUsersReader is a Reader for the GetIpmiUsers structure.
type GetIpmiUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpmiUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIpmiUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIpmiUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIpmiUsersOK creates a GetIpmiUsersOK with default headers values
func NewGetIpmiUsersOK() *GetIpmiUsersOK {
	return &GetIpmiUsersOK{}
}

/*
GetIpmiUsersOK describes a response with status code 200, with default header values.

Success
*/
type GetIpmiUsersOK struct {
	Payload *models.IpmiUsers
}

// IsSuccess returns true when this get ipmi users o k response has a 2xx status code
func (o *GetIpmiUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ipmi users o k response has a 3xx status code
func (o *GetIpmiUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ipmi users o k response has a 4xx status code
func (o *GetIpmiUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ipmi users o k response has a 5xx status code
func (o *GetIpmiUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ipmi users o k response a status code equal to that given
func (o *GetIpmiUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get ipmi users o k response
func (o *GetIpmiUsersOK) Code() int {
	return 200
}

func (o *GetIpmiUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/users][%d] getIpmiUsersOK %s", 200, payload)
}

func (o *GetIpmiUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/users][%d] getIpmiUsersOK %s", 200, payload)
}

func (o *GetIpmiUsersOK) GetPayload() *models.IpmiUsers {
	return o.Payload
}

func (o *GetIpmiUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpmiUsers)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIpmiUsersDefault creates a GetIpmiUsersDefault with default headers values
func NewGetIpmiUsersDefault(code int) *GetIpmiUsersDefault {
	return &GetIpmiUsersDefault{
		_statusCode: code,
	}
}

/*
GetIpmiUsersDefault describes a response with status code -1, with default header values.

Error
*/
type GetIpmiUsersDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get ipmi users default response has a 2xx status code
func (o *GetIpmiUsersDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get ipmi users default response has a 3xx status code
func (o *GetIpmiUsersDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get ipmi users default response has a 4xx status code
func (o *GetIpmiUsersDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get ipmi users default response has a 5xx status code
func (o *GetIpmiUsersDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get ipmi users default response a status code equal to that given
func (o *GetIpmiUsersDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get ipmi users default response
func (o *GetIpmiUsersDefault) Code() int {
	return o._statusCode
}

func (o *GetIpmiUsersDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/users][%d] GetIpmiUsers default %s", o._statusCode, payload)
}

func (o *GetIpmiUsersDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/users][%d] GetIpmiUsers default %s", o._statusCode, payload)
}

func (o *GetIpmiUsersDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIpmiUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

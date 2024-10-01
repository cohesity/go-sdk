// Code generated by go-swagger; DO NOT EDIT.

package network

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

// CreateIPConfigReader is a Reader for the CreateIPConfig structure.
type CreateIPConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIPConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIPConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateIPConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateIPConfigOK creates a CreateIPConfigOK with default headers values
func NewCreateIPConfigOK() *CreateIPConfigOK {
	return &CreateIPConfigOK{}
}

/*
CreateIPConfigOK describes a response with status code 200, with default header values.

Success
*/
type CreateIPConfigOK struct {
	Payload *models.CreateIPConfigResult
}

// IsSuccess returns true when this create Ip config o k response has a 2xx status code
func (o *CreateIPConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create Ip config o k response has a 3xx status code
func (o *CreateIPConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create Ip config o k response has a 4xx status code
func (o *CreateIPConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this create Ip config o k response has a 5xx status code
func (o *CreateIPConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this create Ip config o k response a status code equal to that given
func (o *CreateIPConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the create Ip config o k response
func (o *CreateIPConfigOK) Code() int {
	return 200
}

func (o *CreateIPConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/network/ipConfig][%d] createIpConfigOK %s", 200, payload)
}

func (o *CreateIPConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/network/ipConfig][%d] createIpConfigOK %s", 200, payload)
}

func (o *CreateIPConfigOK) GetPayload() *models.CreateIPConfigResult {
	return o.Payload
}

func (o *CreateIPConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateIPConfigResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIPConfigDefault creates a CreateIPConfigDefault with default headers values
func NewCreateIPConfigDefault(code int) *CreateIPConfigDefault {
	return &CreateIPConfigDefault{
		_statusCode: code,
	}
}

/*
CreateIPConfigDefault describes a response with status code -1, with default header values.

Error
*/
type CreateIPConfigDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this create Ip config default response has a 2xx status code
func (o *CreateIPConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this create Ip config default response has a 3xx status code
func (o *CreateIPConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this create Ip config default response has a 4xx status code
func (o *CreateIPConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this create Ip config default response has a 5xx status code
func (o *CreateIPConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this create Ip config default response a status code equal to that given
func (o *CreateIPConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the create Ip config default response
func (o *CreateIPConfigDefault) Code() int {
	return o._statusCode
}

func (o *CreateIPConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/network/ipConfig][%d] CreateIpConfig default %s", o._statusCode, payload)
}

func (o *CreateIPConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/network/ipConfig][%d] CreateIpConfig default %s", o._statusCode, payload)
}

func (o *CreateIPConfigDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *CreateIPConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

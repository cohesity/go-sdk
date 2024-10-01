// Code generated by go-swagger; DO NOT EDIT.

package source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// GenerateM365DeviceCodeReader is a Reader for the GenerateM365DeviceCode structure.
type GenerateM365DeviceCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateM365DeviceCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGenerateM365DeviceCodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateM365DeviceCodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateM365DeviceCodeCreated creates a GenerateM365DeviceCodeCreated with default headers values
func NewGenerateM365DeviceCodeCreated() *GenerateM365DeviceCodeCreated {
	return &GenerateM365DeviceCodeCreated{}
}

/*
GenerateM365DeviceCodeCreated describes a response with status code 201, with default header values.

Success
*/
type GenerateM365DeviceCodeCreated struct {
	Payload *models.GenerateM365DeviceCodeResponseParams
}

// IsSuccess returns true when this generate m365 device code created response has a 2xx status code
func (o *GenerateM365DeviceCodeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate m365 device code created response has a 3xx status code
func (o *GenerateM365DeviceCodeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate m365 device code created response has a 4xx status code
func (o *GenerateM365DeviceCodeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate m365 device code created response has a 5xx status code
func (o *GenerateM365DeviceCodeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this generate m365 device code created response a status code equal to that given
func (o *GenerateM365DeviceCodeCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the generate m365 device code created response
func (o *GenerateM365DeviceCodeCreated) Code() int {
	return 201
}

func (o *GenerateM365DeviceCodeCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/device-code][%d] generateM365DeviceCodeCreated %s", 201, payload)
}

func (o *GenerateM365DeviceCodeCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/device-code][%d] generateM365DeviceCodeCreated %s", 201, payload)
}

func (o *GenerateM365DeviceCodeCreated) GetPayload() *models.GenerateM365DeviceCodeResponseParams {
	return o.Payload
}

func (o *GenerateM365DeviceCodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenerateM365DeviceCodeResponseParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateM365DeviceCodeDefault creates a GenerateM365DeviceCodeDefault with default headers values
func NewGenerateM365DeviceCodeDefault(code int) *GenerateM365DeviceCodeDefault {
	return &GenerateM365DeviceCodeDefault{
		_statusCode: code,
	}
}

/*
GenerateM365DeviceCodeDefault describes a response with status code -1, with default header values.

Error
*/
type GenerateM365DeviceCodeDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this generate m365 device code default response has a 2xx status code
func (o *GenerateM365DeviceCodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate m365 device code default response has a 3xx status code
func (o *GenerateM365DeviceCodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate m365 device code default response has a 4xx status code
func (o *GenerateM365DeviceCodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate m365 device code default response has a 5xx status code
func (o *GenerateM365DeviceCodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate m365 device code default response a status code equal to that given
func (o *GenerateM365DeviceCodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the generate m365 device code default response
func (o *GenerateM365DeviceCodeDefault) Code() int {
	return o._statusCode
}

func (o *GenerateM365DeviceCodeDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/device-code][%d] GenerateM365DeviceCode default %s", o._statusCode, payload)
}

func (o *GenerateM365DeviceCodeDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/device-code][%d] GenerateM365DeviceCode default %s", o._statusCode, payload)
}

func (o *GenerateM365DeviceCodeDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateM365DeviceCodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

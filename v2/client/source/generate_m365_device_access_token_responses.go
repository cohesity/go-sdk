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

	"github.com/cohesity/go-sdk/v2/models"
)

// GenerateM365DeviceAccessTokenReader is a Reader for the GenerateM365DeviceAccessToken structure.
type GenerateM365DeviceAccessTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateM365DeviceAccessTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewGenerateM365DeviceAccessTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGenerateM365DeviceAccessTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateM365DeviceAccessTokenCreated creates a GenerateM365DeviceAccessTokenCreated with default headers values
func NewGenerateM365DeviceAccessTokenCreated() *GenerateM365DeviceAccessTokenCreated {
	return &GenerateM365DeviceAccessTokenCreated{}
}

/*
GenerateM365DeviceAccessTokenCreated describes a response with status code 201, with default header values.

Success
*/
type GenerateM365DeviceAccessTokenCreated struct {
	Payload *models.GenerateM365DeviceAccessTokenResponseParams
}

// IsSuccess returns true when this generate m365 device access token created response has a 2xx status code
func (o *GenerateM365DeviceAccessTokenCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate m365 device access token created response has a 3xx status code
func (o *GenerateM365DeviceAccessTokenCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate m365 device access token created response has a 4xx status code
func (o *GenerateM365DeviceAccessTokenCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate m365 device access token created response has a 5xx status code
func (o *GenerateM365DeviceAccessTokenCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this generate m365 device access token created response a status code equal to that given
func (o *GenerateM365DeviceAccessTokenCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the generate m365 device access token created response
func (o *GenerateM365DeviceAccessTokenCreated) Code() int {
	return 201
}

func (o *GenerateM365DeviceAccessTokenCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/token][%d] generateM365DeviceAccessTokenCreated %s", 201, payload)
}

func (o *GenerateM365DeviceAccessTokenCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/token][%d] generateM365DeviceAccessTokenCreated %s", 201, payload)
}

func (o *GenerateM365DeviceAccessTokenCreated) GetPayload() *models.GenerateM365DeviceAccessTokenResponseParams {
	return o.Payload
}

func (o *GenerateM365DeviceAccessTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenerateM365DeviceAccessTokenResponseParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateM365DeviceAccessTokenDefault creates a GenerateM365DeviceAccessTokenDefault with default headers values
func NewGenerateM365DeviceAccessTokenDefault(code int) *GenerateM365DeviceAccessTokenDefault {
	return &GenerateM365DeviceAccessTokenDefault{
		_statusCode: code,
	}
}

/*
GenerateM365DeviceAccessTokenDefault describes a response with status code -1, with default header values.

Error
*/
type GenerateM365DeviceAccessTokenDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this generate m365 device access token default response has a 2xx status code
func (o *GenerateM365DeviceAccessTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this generate m365 device access token default response has a 3xx status code
func (o *GenerateM365DeviceAccessTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this generate m365 device access token default response has a 4xx status code
func (o *GenerateM365DeviceAccessTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this generate m365 device access token default response has a 5xx status code
func (o *GenerateM365DeviceAccessTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this generate m365 device access token default response a status code equal to that given
func (o *GenerateM365DeviceAccessTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the generate m365 device access token default response
func (o *GenerateM365DeviceAccessTokenDefault) Code() int {
	return o._statusCode
}

func (o *GenerateM365DeviceAccessTokenDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/token][%d] GenerateM365DeviceAccessToken default %s", o._statusCode, payload)
}

func (o *GenerateM365DeviceAccessTokenDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /data-protect/sources/microsoft365/auth/token][%d] GenerateM365DeviceAccessToken default %s", o._statusCode, payload)
}

func (o *GenerateM365DeviceAccessTokenDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GenerateM365DeviceAccessTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

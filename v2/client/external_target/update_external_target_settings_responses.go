// Code generated by go-swagger; DO NOT EDIT.

package external_target

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

// UpdateExternalTargetSettingsReader is a Reader for the UpdateExternalTargetSettings structure.
type UpdateExternalTargetSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateExternalTargetSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateExternalTargetSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateExternalTargetSettingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateExternalTargetSettingsOK creates a UpdateExternalTargetSettingsOK with default headers values
func NewUpdateExternalTargetSettingsOK() *UpdateExternalTargetSettingsOK {
	return &UpdateExternalTargetSettingsOK{}
}

/*
UpdateExternalTargetSettingsOK describes a response with status code 200, with default header values.

Success
*/
type UpdateExternalTargetSettingsOK struct {
	Payload *models.GlobalBandwidthSettings
}

// IsSuccess returns true when this update external target settings o k response has a 2xx status code
func (o *UpdateExternalTargetSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update external target settings o k response has a 3xx status code
func (o *UpdateExternalTargetSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update external target settings o k response has a 4xx status code
func (o *UpdateExternalTargetSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update external target settings o k response has a 5xx status code
func (o *UpdateExternalTargetSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update external target settings o k response a status code equal to that given
func (o *UpdateExternalTargetSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update external target settings o k response
func (o *UpdateExternalTargetSettingsOK) Code() int {
	return 200
}

func (o *UpdateExternalTargetSettingsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/external-targets/settings][%d] updateExternalTargetSettingsOK %s", 200, payload)
}

func (o *UpdateExternalTargetSettingsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/external-targets/settings][%d] updateExternalTargetSettingsOK %s", 200, payload)
}

func (o *UpdateExternalTargetSettingsOK) GetPayload() *models.GlobalBandwidthSettings {
	return o.Payload
}

func (o *UpdateExternalTargetSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GlobalBandwidthSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateExternalTargetSettingsDefault creates a UpdateExternalTargetSettingsDefault with default headers values
func NewUpdateExternalTargetSettingsDefault(code int) *UpdateExternalTargetSettingsDefault {
	return &UpdateExternalTargetSettingsDefault{
		_statusCode: code,
	}
}

/*
UpdateExternalTargetSettingsDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateExternalTargetSettingsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update external target settings default response has a 2xx status code
func (o *UpdateExternalTargetSettingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update external target settings default response has a 3xx status code
func (o *UpdateExternalTargetSettingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update external target settings default response has a 4xx status code
func (o *UpdateExternalTargetSettingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update external target settings default response has a 5xx status code
func (o *UpdateExternalTargetSettingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update external target settings default response a status code equal to that given
func (o *UpdateExternalTargetSettingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update external target settings default response
func (o *UpdateExternalTargetSettingsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateExternalTargetSettingsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/external-targets/settings][%d] UpdateExternalTargetSettings default %s", o._statusCode, payload)
}

func (o *UpdateExternalTargetSettingsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /data-protect/external-targets/settings][%d] UpdateExternalTargetSettings default %s", o._statusCode, payload)
}

func (o *UpdateExternalTargetSettingsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateExternalTargetSettingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

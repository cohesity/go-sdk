// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// TriggerTrustedDomainsDiscoveryReader is a Reader for the TriggerTrustedDomainsDiscovery structure.
type TriggerTrustedDomainsDiscoveryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TriggerTrustedDomainsDiscoveryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewTriggerTrustedDomainsDiscoveryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTriggerTrustedDomainsDiscoveryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTriggerTrustedDomainsDiscoveryAccepted creates a TriggerTrustedDomainsDiscoveryAccepted with default headers values
func NewTriggerTrustedDomainsDiscoveryAccepted() *TriggerTrustedDomainsDiscoveryAccepted {
	return &TriggerTrustedDomainsDiscoveryAccepted{}
}

/*
TriggerTrustedDomainsDiscoveryAccepted describes a response with status code 202, with default header values.

Request Accepted
*/
type TriggerTrustedDomainsDiscoveryAccepted struct {
}

// IsSuccess returns true when this trigger trusted domains discovery accepted response has a 2xx status code
func (o *TriggerTrustedDomainsDiscoveryAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this trigger trusted domains discovery accepted response has a 3xx status code
func (o *TriggerTrustedDomainsDiscoveryAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this trigger trusted domains discovery accepted response has a 4xx status code
func (o *TriggerTrustedDomainsDiscoveryAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this trigger trusted domains discovery accepted response has a 5xx status code
func (o *TriggerTrustedDomainsDiscoveryAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this trigger trusted domains discovery accepted response a status code equal to that given
func (o *TriggerTrustedDomainsDiscoveryAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the trigger trusted domains discovery accepted response
func (o *TriggerTrustedDomainsDiscoveryAccepted) Code() int {
	return 202
}

func (o *TriggerTrustedDomainsDiscoveryAccepted) Error() string {
	return fmt.Sprintf("[PUT /trusted-domains][%d] triggerTrustedDomainsDiscoveryAccepted", 202)
}

func (o *TriggerTrustedDomainsDiscoveryAccepted) String() string {
	return fmt.Sprintf("[PUT /trusted-domains][%d] triggerTrustedDomainsDiscoveryAccepted", 202)
}

func (o *TriggerTrustedDomainsDiscoveryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTriggerTrustedDomainsDiscoveryDefault creates a TriggerTrustedDomainsDiscoveryDefault with default headers values
func NewTriggerTrustedDomainsDiscoveryDefault(code int) *TriggerTrustedDomainsDiscoveryDefault {
	return &TriggerTrustedDomainsDiscoveryDefault{
		_statusCode: code,
	}
}

/*
TriggerTrustedDomainsDiscoveryDefault describes a response with status code -1, with default header values.

Error
*/
type TriggerTrustedDomainsDiscoveryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this trigger trusted domains discovery default response has a 2xx status code
func (o *TriggerTrustedDomainsDiscoveryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this trigger trusted domains discovery default response has a 3xx status code
func (o *TriggerTrustedDomainsDiscoveryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this trigger trusted domains discovery default response has a 4xx status code
func (o *TriggerTrustedDomainsDiscoveryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this trigger trusted domains discovery default response has a 5xx status code
func (o *TriggerTrustedDomainsDiscoveryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this trigger trusted domains discovery default response a status code equal to that given
func (o *TriggerTrustedDomainsDiscoveryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the trigger trusted domains discovery default response
func (o *TriggerTrustedDomainsDiscoveryDefault) Code() int {
	return o._statusCode
}

func (o *TriggerTrustedDomainsDiscoveryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trusted-domains][%d] TriggerTrustedDomainsDiscovery default %s", o._statusCode, payload)
}

func (o *TriggerTrustedDomainsDiscoveryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /trusted-domains][%d] TriggerTrustedDomainsDiscovery default %s", o._statusCode, payload)
}

func (o *TriggerTrustedDomainsDiscoveryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *TriggerTrustedDomainsDiscoveryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

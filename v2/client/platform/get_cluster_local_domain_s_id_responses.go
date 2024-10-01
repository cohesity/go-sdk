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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// GetClusterLocalDomainSIDReader is a Reader for the GetClusterLocalDomainSID structure.
type GetClusterLocalDomainSIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterLocalDomainSIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterLocalDomainSIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterLocalDomainSIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterLocalDomainSIDOK creates a GetClusterLocalDomainSIDOK with default headers values
func NewGetClusterLocalDomainSIDOK() *GetClusterLocalDomainSIDOK {
	return &GetClusterLocalDomainSIDOK{}
}

/*
GetClusterLocalDomainSIDOK describes a response with status code 200, with default header values.

Success
*/
type GetClusterLocalDomainSIDOK struct {
	Payload *models.ClusterLocalDomainSID
}

// IsSuccess returns true when this get cluster local domain s Id o k response has a 2xx status code
func (o *GetClusterLocalDomainSIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get cluster local domain s Id o k response has a 3xx status code
func (o *GetClusterLocalDomainSIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get cluster local domain s Id o k response has a 4xx status code
func (o *GetClusterLocalDomainSIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get cluster local domain s Id o k response has a 5xx status code
func (o *GetClusterLocalDomainSIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get cluster local domain s Id o k response a status code equal to that given
func (o *GetClusterLocalDomainSIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get cluster local domain s Id o k response
func (o *GetClusterLocalDomainSIDOK) Code() int {
	return 200
}

func (o *GetClusterLocalDomainSIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/local-domain-sid][%d] getClusterLocalDomainSIdOK %s", 200, payload)
}

func (o *GetClusterLocalDomainSIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/local-domain-sid][%d] getClusterLocalDomainSIdOK %s", 200, payload)
}

func (o *GetClusterLocalDomainSIDOK) GetPayload() *models.ClusterLocalDomainSID {
	return o.Payload
}

func (o *GetClusterLocalDomainSIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterLocalDomainSID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterLocalDomainSIDDefault creates a GetClusterLocalDomainSIDDefault with default headers values
func NewGetClusterLocalDomainSIDDefault(code int) *GetClusterLocalDomainSIDDefault {
	return &GetClusterLocalDomainSIDDefault{
		_statusCode: code,
	}
}

/*
GetClusterLocalDomainSIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetClusterLocalDomainSIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get cluster local domain s ID default response has a 2xx status code
func (o *GetClusterLocalDomainSIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get cluster local domain s ID default response has a 3xx status code
func (o *GetClusterLocalDomainSIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get cluster local domain s ID default response has a 4xx status code
func (o *GetClusterLocalDomainSIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get cluster local domain s ID default response has a 5xx status code
func (o *GetClusterLocalDomainSIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get cluster local domain s ID default response a status code equal to that given
func (o *GetClusterLocalDomainSIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get cluster local domain s ID default response
func (o *GetClusterLocalDomainSIDDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterLocalDomainSIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/local-domain-sid][%d] GetClusterLocalDomainSID default %s", o._statusCode, payload)
}

func (o *GetClusterLocalDomainSIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/local-domain-sid][%d] GetClusterLocalDomainSID default %s", o._statusCode, payload)
}

func (o *GetClusterLocalDomainSIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClusterLocalDomainSIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetIpmiLanConfigReader is a Reader for the GetIpmiLanConfig structure.
type GetIpmiLanConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIpmiLanConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIpmiLanConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIpmiLanConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIpmiLanConfigOK creates a GetIpmiLanConfigOK with default headers values
func NewGetIpmiLanConfigOK() *GetIpmiLanConfigOK {
	return &GetIpmiLanConfigOK{}
}

/*
GetIpmiLanConfigOK describes a response with status code 200, with default header values.

Success
*/
type GetIpmiLanConfigOK struct {
	Payload *models.IpmiLanParams
}

// IsSuccess returns true when this get ipmi lan config o k response has a 2xx status code
func (o *GetIpmiLanConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ipmi lan config o k response has a 3xx status code
func (o *GetIpmiLanConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ipmi lan config o k response has a 4xx status code
func (o *GetIpmiLanConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ipmi lan config o k response has a 5xx status code
func (o *GetIpmiLanConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ipmi lan config o k response a status code equal to that given
func (o *GetIpmiLanConfigOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get ipmi lan config o k response
func (o *GetIpmiLanConfigOK) Code() int {
	return 200
}

func (o *GetIpmiLanConfigOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/lan][%d] getIpmiLanConfigOK %s", 200, payload)
}

func (o *GetIpmiLanConfigOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/lan][%d] getIpmiLanConfigOK %s", 200, payload)
}

func (o *GetIpmiLanConfigOK) GetPayload() *models.IpmiLanParams {
	return o.Payload
}

func (o *GetIpmiLanConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IpmiLanParams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIpmiLanConfigDefault creates a GetIpmiLanConfigDefault with default headers values
func NewGetIpmiLanConfigDefault(code int) *GetIpmiLanConfigDefault {
	return &GetIpmiLanConfigDefault{
		_statusCode: code,
	}
}

/*
GetIpmiLanConfigDefault describes a response with status code -1, with default header values.

Error
*/
type GetIpmiLanConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get ipmi lan config default response has a 2xx status code
func (o *GetIpmiLanConfigDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get ipmi lan config default response has a 3xx status code
func (o *GetIpmiLanConfigDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get ipmi lan config default response has a 4xx status code
func (o *GetIpmiLanConfigDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get ipmi lan config default response has a 5xx status code
func (o *GetIpmiLanConfigDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get ipmi lan config default response a status code equal to that given
func (o *GetIpmiLanConfigDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get ipmi lan config default response
func (o *GetIpmiLanConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetIpmiLanConfigDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/lan][%d] GetIpmiLanConfig default %s", o._statusCode, payload)
}

func (o *GetIpmiLanConfigDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/ipmi/lan][%d] GetIpmiLanConfig default %s", o._statusCode, payload)
}

func (o *GetIpmiLanConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIpmiLanConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

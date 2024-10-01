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

// GetNetworkInterfacesReader is a Reader for the GetNetworkInterfaces structure.
type GetNetworkInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkInterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNetworkInterfacesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetworkInterfacesOK creates a GetNetworkInterfacesOK with default headers values
func NewGetNetworkInterfacesOK() *GetNetworkInterfacesOK {
	return &GetNetworkInterfacesOK{}
}

/*
GetNetworkInterfacesOK describes a response with status code 200, with default header values.

Success
*/
type GetNetworkInterfacesOK struct {
	Payload *models.ClusterInterfaces
}

// IsSuccess returns true when this get network interfaces o k response has a 2xx status code
func (o *GetNetworkInterfacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network interfaces o k response has a 3xx status code
func (o *GetNetworkInterfacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network interfaces o k response has a 4xx status code
func (o *GetNetworkInterfacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network interfaces o k response has a 5xx status code
func (o *GetNetworkInterfacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network interfaces o k response a status code equal to that given
func (o *GetNetworkInterfacesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network interfaces o k response
func (o *GetNetworkInterfacesOK) Code() int {
	return 200
}

func (o *GetNetworkInterfacesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network-interfaces][%d] getNetworkInterfacesOK %s", 200, payload)
}

func (o *GetNetworkInterfacesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network-interfaces][%d] getNetworkInterfacesOK %s", 200, payload)
}

func (o *GetNetworkInterfacesOK) GetPayload() *models.ClusterInterfaces {
	return o.Payload
}

func (o *GetNetworkInterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterInterfaces)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetworkInterfacesDefault creates a GetNetworkInterfacesDefault with default headers values
func NewGetNetworkInterfacesDefault(code int) *GetNetworkInterfacesDefault {
	return &GetNetworkInterfacesDefault{
		_statusCode: code,
	}
}

/*
GetNetworkInterfacesDefault describes a response with status code -1, with default header values.

Error
*/
type GetNetworkInterfacesDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get network interfaces default response has a 2xx status code
func (o *GetNetworkInterfacesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get network interfaces default response has a 3xx status code
func (o *GetNetworkInterfacesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get network interfaces default response has a 4xx status code
func (o *GetNetworkInterfacesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get network interfaces default response has a 5xx status code
func (o *GetNetworkInterfacesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get network interfaces default response a status code equal to that given
func (o *GetNetworkInterfacesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get network interfaces default response
func (o *GetNetworkInterfacesDefault) Code() int {
	return o._statusCode
}

func (o *GetNetworkInterfacesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network-interfaces][%d] GetNetworkInterfaces default %s", o._statusCode, payload)
}

func (o *GetNetworkInterfacesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network-interfaces][%d] GetNetworkInterfaces default %s", o._statusCode, payload)
}

func (o *GetNetworkInterfacesDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNetworkInterfacesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// GetBasicClusterInfoReader is a Reader for the GetBasicClusterInfo structure.
type GetBasicClusterInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBasicClusterInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBasicClusterInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetBasicClusterInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetBasicClusterInfoOK creates a GetBasicClusterInfoOK with default headers values
func NewGetBasicClusterInfoOK() *GetBasicClusterInfoOK {
	return &GetBasicClusterInfoOK{}
}

/*
GetBasicClusterInfoOK describes a response with status code 200, with default header values.

Success
*/
type GetBasicClusterInfoOK struct {
	Payload *models.BasicClusterInfo
}

// IsSuccess returns true when this get basic cluster info o k response has a 2xx status code
func (o *GetBasicClusterInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get basic cluster info o k response has a 3xx status code
func (o *GetBasicClusterInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get basic cluster info o k response has a 4xx status code
func (o *GetBasicClusterInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get basic cluster info o k response has a 5xx status code
func (o *GetBasicClusterInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get basic cluster info o k response a status code equal to that given
func (o *GetBasicClusterInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get basic cluster info o k response
func (o *GetBasicClusterInfoOK) Code() int {
	return 200
}

func (o *GetBasicClusterInfoOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/basicClusterInfo][%d] getBasicClusterInfoOK %s", 200, payload)
}

func (o *GetBasicClusterInfoOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/basicClusterInfo][%d] getBasicClusterInfoOK %s", 200, payload)
}

func (o *GetBasicClusterInfoOK) GetPayload() *models.BasicClusterInfo {
	return o.Payload
}

func (o *GetBasicClusterInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicClusterInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBasicClusterInfoDefault creates a GetBasicClusterInfoDefault with default headers values
func NewGetBasicClusterInfoDefault(code int) *GetBasicClusterInfoDefault {
	return &GetBasicClusterInfoDefault{
		_statusCode: code,
	}
}

/*
GetBasicClusterInfoDefault describes a response with status code -1, with default header values.

Error
*/
type GetBasicClusterInfoDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get basic cluster info default response has a 2xx status code
func (o *GetBasicClusterInfoDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get basic cluster info default response has a 3xx status code
func (o *GetBasicClusterInfoDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get basic cluster info default response has a 4xx status code
func (o *GetBasicClusterInfoDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get basic cluster info default response has a 5xx status code
func (o *GetBasicClusterInfoDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get basic cluster info default response a status code equal to that given
func (o *GetBasicClusterInfoDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get basic cluster info default response
func (o *GetBasicClusterInfoDefault) Code() int {
	return o._statusCode
}

func (o *GetBasicClusterInfoDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/basicClusterInfo][%d] GetBasicClusterInfo default %s", o._statusCode, payload)
}

func (o *GetBasicClusterInfoDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/basicClusterInfo][%d] GetBasicClusterInfo default %s", o._statusCode, payload)
}

func (o *GetBasicClusterInfoDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetBasicClusterInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

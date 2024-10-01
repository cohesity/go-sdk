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

// GetIsDMaaSClusterReader is a Reader for the GetIsDMaaSCluster structure.
type GetIsDMaaSClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIsDMaaSClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIsDMaaSClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIsDMaaSClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIsDMaaSClusterOK creates a GetIsDMaaSClusterOK with default headers values
func NewGetIsDMaaSClusterOK() *GetIsDMaaSClusterOK {
	return &GetIsDMaaSClusterOK{}
}

/*
GetIsDMaaSClusterOK describes a response with status code 200, with default header values.

Success
*/
type GetIsDMaaSClusterOK struct {
	Payload *models.DMaaSInfo
}

// IsSuccess returns true when this get is d maa s cluster o k response has a 2xx status code
func (o *GetIsDMaaSClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get is d maa s cluster o k response has a 3xx status code
func (o *GetIsDMaaSClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get is d maa s cluster o k response has a 4xx status code
func (o *GetIsDMaaSClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get is d maa s cluster o k response has a 5xx status code
func (o *GetIsDMaaSClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get is d maa s cluster o k response a status code equal to that given
func (o *GetIsDMaaSClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get is d maa s cluster o k response
func (o *GetIsDMaaSClusterOK) Code() int {
	return 200
}

func (o *GetIsDMaaSClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/is-dmaas][%d] getIsDMaaSClusterOK %s", 200, payload)
}

func (o *GetIsDMaaSClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/is-dmaas][%d] getIsDMaaSClusterOK %s", 200, payload)
}

func (o *GetIsDMaaSClusterOK) GetPayload() *models.DMaaSInfo {
	return o.Payload
}

func (o *GetIsDMaaSClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DMaaSInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIsDMaaSClusterDefault creates a GetIsDMaaSClusterDefault with default headers values
func NewGetIsDMaaSClusterDefault(code int) *GetIsDMaaSClusterDefault {
	return &GetIsDMaaSClusterDefault{
		_statusCode: code,
	}
}

/*
GetIsDMaaSClusterDefault describes a response with status code -1, with default header values.

Error
*/
type GetIsDMaaSClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get is d maa s cluster default response has a 2xx status code
func (o *GetIsDMaaSClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get is d maa s cluster default response has a 3xx status code
func (o *GetIsDMaaSClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get is d maa s cluster default response has a 4xx status code
func (o *GetIsDMaaSClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get is d maa s cluster default response has a 5xx status code
func (o *GetIsDMaaSClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get is d maa s cluster default response a status code equal to that given
func (o *GetIsDMaaSClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get is d maa s cluster default response
func (o *GetIsDMaaSClusterDefault) Code() int {
	return o._statusCode
}

func (o *GetIsDMaaSClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/is-dmaas][%d] GetIsDMaaSCluster default %s", o._statusCode, payload)
}

func (o *GetIsDMaaSClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /clusters/is-dmaas][%d] GetIsDMaaSCluster default %s", o._statusCode, payload)
}

func (o *GetIsDMaaSClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetIsDMaaSClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

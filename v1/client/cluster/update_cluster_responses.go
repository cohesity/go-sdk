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

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/*
UpdateClusterOK describes a response with status code 200, with default header values.

Success
*/
type UpdateClusterOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this update cluster o k response has a 2xx status code
func (o *UpdateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cluster o k response has a 3xx status code
func (o *UpdateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster o k response has a 4xx status code
func (o *UpdateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cluster o k response has a 5xx status code
func (o *UpdateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster o k response a status code equal to that given
func (o *UpdateClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update cluster o k response
func (o *UpdateClusterOK) Code() int {
	return 200
}

func (o *UpdateClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster][%d] updateClusterOK %s", 200, payload)
}

func (o *UpdateClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster][%d] updateClusterOK %s", 200, payload)
}

func (o *UpdateClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterDefault creates a UpdateClusterDefault with default headers values
func NewUpdateClusterDefault(code int) *UpdateClusterDefault {
	return &UpdateClusterDefault{
		_statusCode: code,
	}
}

/*
UpdateClusterDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateClusterDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update cluster default response has a 2xx status code
func (o *UpdateClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update cluster default response has a 3xx status code
func (o *UpdateClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update cluster default response has a 4xx status code
func (o *UpdateClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update cluster default response has a 5xx status code
func (o *UpdateClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update cluster default response a status code equal to that given
func (o *UpdateClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update cluster default response
func (o *UpdateClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster][%d] UpdateCluster default %s", o._statusCode, payload)
}

func (o *UpdateClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/cluster][%d] UpdateCluster default %s", o._statusCode, payload)
}

func (o *UpdateClusterDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

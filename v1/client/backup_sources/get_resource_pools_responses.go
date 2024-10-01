// Code generated by go-swagger; DO NOT EDIT.

package backup_sources

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

// GetResourcePoolsReader is a Reader for the GetResourcePools structure.
type GetResourcePoolsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourcePoolsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourcePoolsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetResourcePoolsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetResourcePoolsOK creates a GetResourcePoolsOK with default headers values
func NewGetResourcePoolsOK() *GetResourcePoolsOK {
	return &GetResourcePoolsOK{}
}

/*
GetResourcePoolsOK describes a response with status code 200, with default header values.

Success
*/
type GetResourcePoolsOK struct {
	Payload []*models.ResourcePoolResult
}

// IsSuccess returns true when this get resource pools o k response has a 2xx status code
func (o *GetResourcePoolsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get resource pools o k response has a 3xx status code
func (o *GetResourcePoolsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get resource pools o k response has a 4xx status code
func (o *GetResourcePoolsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get resource pools o k response has a 5xx status code
func (o *GetResourcePoolsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get resource pools o k response a status code equal to that given
func (o *GetResourcePoolsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get resource pools o k response
func (o *GetResourcePoolsOK) Code() int {
	return 200
}

func (o *GetResourcePoolsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /resourcePools][%d] getResourcePoolsOK %s", 200, payload)
}

func (o *GetResourcePoolsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /resourcePools][%d] getResourcePoolsOK %s", 200, payload)
}

func (o *GetResourcePoolsOK) GetPayload() []*models.ResourcePoolResult {
	return o.Payload
}

func (o *GetResourcePoolsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourcePoolsDefault creates a GetResourcePoolsDefault with default headers values
func NewGetResourcePoolsDefault(code int) *GetResourcePoolsDefault {
	return &GetResourcePoolsDefault{
		_statusCode: code,
	}
}

/*
GetResourcePoolsDefault describes a response with status code -1, with default header values.

(empty)
*/
type GetResourcePoolsDefault struct {
	_statusCode int

	Payload *models.PrivateError
}

// IsSuccess returns true when this get resource pools default response has a 2xx status code
func (o *GetResourcePoolsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get resource pools default response has a 3xx status code
func (o *GetResourcePoolsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get resource pools default response has a 4xx status code
func (o *GetResourcePoolsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get resource pools default response has a 5xx status code
func (o *GetResourcePoolsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get resource pools default response a status code equal to that given
func (o *GetResourcePoolsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get resource pools default response
func (o *GetResourcePoolsDefault) Code() int {
	return o._statusCode
}

func (o *GetResourcePoolsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /resourcePools][%d] GetResourcePools default %s", o._statusCode, payload)
}

func (o *GetResourcePoolsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /resourcePools][%d] GetResourcePools default %s", o._statusCode, payload)
}

func (o *GetResourcePoolsDefault) GetPayload() *models.PrivateError {
	return o.Payload
}

func (o *GetResourcePoolsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PrivateError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

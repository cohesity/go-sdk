// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// GetTenantByIDReader is a Reader for the GetTenantByID structure.
type GetTenantByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTenantByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTenantByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTenantByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTenantByIDOK creates a GetTenantByIDOK with default headers values
func NewGetTenantByIDOK() *GetTenantByIDOK {
	return &GetTenantByIDOK{}
}

/*
GetTenantByIDOK describes a response with status code 200, with default header values.

Success
*/
type GetTenantByIDOK struct {
	Payload *models.TenantInfo
}

// IsSuccess returns true when this get tenant by Id o k response has a 2xx status code
func (o *GetTenantByIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get tenant by Id o k response has a 3xx status code
func (o *GetTenantByIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get tenant by Id o k response has a 4xx status code
func (o *GetTenantByIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get tenant by Id o k response has a 5xx status code
func (o *GetTenantByIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get tenant by Id o k response a status code equal to that given
func (o *GetTenantByIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get tenant by Id o k response
func (o *GetTenantByIDOK) Code() int {
	return 200
}

func (o *GetTenantByIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants/{id}][%d] getTenantByIdOK %s", 200, payload)
}

func (o *GetTenantByIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants/{id}][%d] getTenantByIdOK %s", 200, payload)
}

func (o *GetTenantByIDOK) GetPayload() *models.TenantInfo {
	return o.Payload
}

func (o *GetTenantByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTenantByIDDefault creates a GetTenantByIDDefault with default headers values
func NewGetTenantByIDDefault(code int) *GetTenantByIDDefault {
	return &GetTenantByIDDefault{
		_statusCode: code,
	}
}

/*
GetTenantByIDDefault describes a response with status code -1, with default header values.

Error
*/
type GetTenantByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get tenant by ID default response has a 2xx status code
func (o *GetTenantByIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get tenant by ID default response has a 3xx status code
func (o *GetTenantByIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get tenant by ID default response has a 4xx status code
func (o *GetTenantByIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get tenant by ID default response has a 5xx status code
func (o *GetTenantByIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get tenant by ID default response a status code equal to that given
func (o *GetTenantByIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get tenant by ID default response
func (o *GetTenantByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetTenantByIDDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants/{id}][%d] GetTenantByID default %s", o._statusCode, payload)
}

func (o *GetTenantByIDDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /tenants/{id}][%d] GetTenantByID default %s", o._statusCode, payload)
}

func (o *GetTenantByIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTenantByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

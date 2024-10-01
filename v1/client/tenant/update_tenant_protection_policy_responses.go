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

	"github.com/cohesity/go-sdk/v1/models"
)

// UpdateTenantProtectionPolicyReader is a Reader for the UpdateTenantProtectionPolicy structure.
type UpdateTenantProtectionPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTenantProtectionPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateTenantProtectionPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateTenantProtectionPolicyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateTenantProtectionPolicyOK creates a UpdateTenantProtectionPolicyOK with default headers values
func NewUpdateTenantProtectionPolicyOK() *UpdateTenantProtectionPolicyOK {
	return &UpdateTenantProtectionPolicyOK{}
}

/*
UpdateTenantProtectionPolicyOK describes a response with status code 200, with default header values.

Tenant Protection Policy Mapping Response.
*/
type UpdateTenantProtectionPolicyOK struct {
	Payload *models.TenantProtectionPolicyUpdate
}

// IsSuccess returns true when this update tenant protection policy o k response has a 2xx status code
func (o *UpdateTenantProtectionPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update tenant protection policy o k response has a 3xx status code
func (o *UpdateTenantProtectionPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update tenant protection policy o k response has a 4xx status code
func (o *UpdateTenantProtectionPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update tenant protection policy o k response has a 5xx status code
func (o *UpdateTenantProtectionPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update tenant protection policy o k response a status code equal to that given
func (o *UpdateTenantProtectionPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update tenant protection policy o k response
func (o *UpdateTenantProtectionPolicyOK) Code() int {
	return 200
}

func (o *UpdateTenantProtectionPolicyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/policy][%d] updateTenantProtectionPolicyOK %s", 200, payload)
}

func (o *UpdateTenantProtectionPolicyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/policy][%d] updateTenantProtectionPolicyOK %s", 200, payload)
}

func (o *UpdateTenantProtectionPolicyOK) GetPayload() *models.TenantProtectionPolicyUpdate {
	return o.Payload
}

func (o *UpdateTenantProtectionPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantProtectionPolicyUpdate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateTenantProtectionPolicyDefault creates a UpdateTenantProtectionPolicyDefault with default headers values
func NewUpdateTenantProtectionPolicyDefault(code int) *UpdateTenantProtectionPolicyDefault {
	return &UpdateTenantProtectionPolicyDefault{
		_statusCode: code,
	}
}

/*
UpdateTenantProtectionPolicyDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateTenantProtectionPolicyDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this update tenant protection policy default response has a 2xx status code
func (o *UpdateTenantProtectionPolicyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update tenant protection policy default response has a 3xx status code
func (o *UpdateTenantProtectionPolicyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update tenant protection policy default response has a 4xx status code
func (o *UpdateTenantProtectionPolicyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update tenant protection policy default response has a 5xx status code
func (o *UpdateTenantProtectionPolicyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update tenant protection policy default response a status code equal to that given
func (o *UpdateTenantProtectionPolicyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update tenant protection policy default response
func (o *UpdateTenantProtectionPolicyDefault) Code() int {
	return o._statusCode
}

func (o *UpdateTenantProtectionPolicyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/policy][%d] UpdateTenantProtectionPolicy default %s", o._statusCode, payload)
}

func (o *UpdateTenantProtectionPolicyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /public/tenants/policy][%d] UpdateTenantProtectionPolicy default %s", o._statusCode, payload)
}

func (o *UpdateTenantProtectionPolicyDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *UpdateTenantProtectionPolicyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

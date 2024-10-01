// Code generated by go-swagger; DO NOT EDIT.

package audit

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

// GetAuditLogsCategoriesReader is a Reader for the GetAuditLogsCategories structure.
type GetAuditLogsCategoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditLogsCategoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditLogsCategoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAuditLogsCategoriesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAuditLogsCategoriesOK creates a GetAuditLogsCategoriesOK with default headers values
func NewGetAuditLogsCategoriesOK() *GetAuditLogsCategoriesOK {
	return &GetAuditLogsCategoriesOK{}
}

/*
GetAuditLogsCategoriesOK describes a response with status code 200, with default header values.

Success
*/
type GetAuditLogsCategoriesOK struct {
	Payload []string
}

// IsSuccess returns true when this get audit logs categories o k response has a 2xx status code
func (o *GetAuditLogsCategoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get audit logs categories o k response has a 3xx status code
func (o *GetAuditLogsCategoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get audit logs categories o k response has a 4xx status code
func (o *GetAuditLogsCategoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get audit logs categories o k response has a 5xx status code
func (o *GetAuditLogsCategoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get audit logs categories o k response a status code equal to that given
func (o *GetAuditLogsCategoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get audit logs categories o k response
func (o *GetAuditLogsCategoriesOK) Code() int {
	return 200
}

func (o *GetAuditLogsCategoriesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/auditLogs/categories][%d] getAuditLogsCategoriesOK %s", 200, payload)
}

func (o *GetAuditLogsCategoriesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/auditLogs/categories][%d] getAuditLogsCategoriesOK %s", 200, payload)
}

func (o *GetAuditLogsCategoriesOK) GetPayload() []string {
	return o.Payload
}

func (o *GetAuditLogsCategoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditLogsCategoriesDefault creates a GetAuditLogsCategoriesDefault with default headers values
func NewGetAuditLogsCategoriesDefault(code int) *GetAuditLogsCategoriesDefault {
	return &GetAuditLogsCategoriesDefault{
		_statusCode: code,
	}
}

/*
GetAuditLogsCategoriesDefault describes a response with status code -1, with default header values.

Error
*/
type GetAuditLogsCategoriesDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get audit logs categories default response has a 2xx status code
func (o *GetAuditLogsCategoriesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get audit logs categories default response has a 3xx status code
func (o *GetAuditLogsCategoriesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get audit logs categories default response has a 4xx status code
func (o *GetAuditLogsCategoriesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get audit logs categories default response has a 5xx status code
func (o *GetAuditLogsCategoriesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get audit logs categories default response a status code equal to that given
func (o *GetAuditLogsCategoriesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get audit logs categories default response
func (o *GetAuditLogsCategoriesDefault) Code() int {
	return o._statusCode
}

func (o *GetAuditLogsCategoriesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/auditLogs/categories][%d] GetAuditLogsCategories default %s", o._statusCode, payload)
}

func (o *GetAuditLogsCategoriesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/auditLogs/categories][%d] GetAuditLogsCategories default %s", o._statusCode, payload)
}

func (o *GetAuditLogsCategoriesDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetAuditLogsCategoriesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

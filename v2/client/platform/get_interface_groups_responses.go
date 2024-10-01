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

// GetInterfaceGroupsReader is a Reader for the GetInterfaceGroups structure.
type GetInterfaceGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInterfaceGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInterfaceGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetInterfaceGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetInterfaceGroupsOK creates a GetInterfaceGroupsOK with default headers values
func NewGetInterfaceGroupsOK() *GetInterfaceGroupsOK {
	return &GetInterfaceGroupsOK{}
}

/*
GetInterfaceGroupsOK describes a response with status code 200, with default header values.

Success
*/
type GetInterfaceGroupsOK struct {
	Payload *models.InterfaceGroups
}

// IsSuccess returns true when this get interface groups o k response has a 2xx status code
func (o *GetInterfaceGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get interface groups o k response has a 3xx status code
func (o *GetInterfaceGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get interface groups o k response has a 4xx status code
func (o *GetInterfaceGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get interface groups o k response has a 5xx status code
func (o *GetInterfaceGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get interface groups o k response a status code equal to that given
func (o *GetInterfaceGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get interface groups o k response
func (o *GetInterfaceGroupsOK) Code() int {
	return 200
}

func (o *GetInterfaceGroupsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/interface-groups][%d] getInterfaceGroupsOK %s", 200, payload)
}

func (o *GetInterfaceGroupsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/interface-groups][%d] getInterfaceGroupsOK %s", 200, payload)
}

func (o *GetInterfaceGroupsOK) GetPayload() *models.InterfaceGroups {
	return o.Payload
}

func (o *GetInterfaceGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceGroups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInterfaceGroupsDefault creates a GetInterfaceGroupsDefault with default headers values
func NewGetInterfaceGroupsDefault(code int) *GetInterfaceGroupsDefault {
	return &GetInterfaceGroupsDefault{
		_statusCode: code,
	}
}

/*
GetInterfaceGroupsDefault describes a response with status code -1, with default header values.

Error
*/
type GetInterfaceGroupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get interface groups default response has a 2xx status code
func (o *GetInterfaceGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get interface groups default response has a 3xx status code
func (o *GetInterfaceGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get interface groups default response has a 4xx status code
func (o *GetInterfaceGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get interface groups default response has a 5xx status code
func (o *GetInterfaceGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get interface groups default response a status code equal to that given
func (o *GetInterfaceGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get interface groups default response
func (o *GetInterfaceGroupsDefault) Code() int {
	return o._statusCode
}

func (o *GetInterfaceGroupsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/interface-groups][%d] GetInterfaceGroups default %s", o._statusCode, payload)
}

func (o *GetInterfaceGroupsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /network/interface-groups][%d] GetInterfaceGroups default %s", o._statusCode, payload)
}

func (o *GetInterfaceGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetInterfaceGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package views

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

// GetNfsConnectionsReader is a Reader for the GetNfsConnections structure.
type GetNfsConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNfsConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNfsConnectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNfsConnectionsOK creates a GetNfsConnectionsOK with default headers values
func NewGetNfsConnectionsOK() *GetNfsConnectionsOK {
	return &GetNfsConnectionsOK{}
}

/*
GetNfsConnectionsOK describes a response with status code 200, with default header values.

Success
*/
type GetNfsConnectionsOK struct {
	Payload []*models.NfsConnection
}

// IsSuccess returns true when this get nfs connections o k response has a 2xx status code
func (o *GetNfsConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get nfs connections o k response has a 3xx status code
func (o *GetNfsConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get nfs connections o k response has a 4xx status code
func (o *GetNfsConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get nfs connections o k response has a 5xx status code
func (o *GetNfsConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get nfs connections o k response a status code equal to that given
func (o *GetNfsConnectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get nfs connections o k response
func (o *GetNfsConnectionsOK) Code() int {
	return 200
}

func (o *GetNfsConnectionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/nfsConnections][%d] getNfsConnectionsOK %s", 200, payload)
}

func (o *GetNfsConnectionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/nfsConnections][%d] getNfsConnectionsOK %s", 200, payload)
}

func (o *GetNfsConnectionsOK) GetPayload() []*models.NfsConnection {
	return o.Payload
}

func (o *GetNfsConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsConnectionsDefault creates a GetNfsConnectionsDefault with default headers values
func NewGetNfsConnectionsDefault(code int) *GetNfsConnectionsDefault {
	return &GetNfsConnectionsDefault{
		_statusCode: code,
	}
}

/*
GetNfsConnectionsDefault describes a response with status code -1, with default header values.

Error
*/
type GetNfsConnectionsDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this get nfs connections default response has a 2xx status code
func (o *GetNfsConnectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get nfs connections default response has a 3xx status code
func (o *GetNfsConnectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get nfs connections default response has a 4xx status code
func (o *GetNfsConnectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get nfs connections default response has a 5xx status code
func (o *GetNfsConnectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get nfs connections default response a status code equal to that given
func (o *GetNfsConnectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get nfs connections default response
func (o *GetNfsConnectionsDefault) Code() int {
	return o._statusCode
}

func (o *GetNfsConnectionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/nfsConnections][%d] GetNfsConnections default %s", o._statusCode, payload)
}

func (o *GetNfsConnectionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/nfsConnections][%d] GetNfsConnections default %s", o._statusCode, payload)
}

func (o *GetNfsConnectionsDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *GetNfsConnectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

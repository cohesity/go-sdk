// Code generated by go-swagger; DO NOT EDIT.

package failover

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

// GetFailoverOpsReader is a Reader for the GetFailoverOps structure.
type GetFailoverOpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFailoverOpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFailoverOpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFailoverOpsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFailoverOpsOK creates a GetFailoverOpsOK with default headers values
func NewGetFailoverOpsOK() *GetFailoverOpsOK {
	return &GetFailoverOpsOK{}
}

/*
GetFailoverOpsOK describes a response with status code 200, with default header values.

Success
*/
type GetFailoverOpsOK struct {
	Payload *models.GetFailoverOpsResponse
}

// IsSuccess returns true when this get failover ops o k response has a 2xx status code
func (o *GetFailoverOpsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get failover ops o k response has a 3xx status code
func (o *GetFailoverOpsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get failover ops o k response has a 4xx status code
func (o *GetFailoverOpsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get failover ops o k response has a 5xx status code
func (o *GetFailoverOpsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get failover ops o k response a status code equal to that given
func (o *GetFailoverOpsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get failover ops o k response
func (o *GetFailoverOpsOK) Code() int {
	return 200
}

func (o *GetFailoverOpsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/failover/views/{id}/operations][%d] getFailoverOpsOK %s", 200, payload)
}

func (o *GetFailoverOpsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/failover/views/{id}/operations][%d] getFailoverOpsOK %s", 200, payload)
}

func (o *GetFailoverOpsOK) GetPayload() *models.GetFailoverOpsResponse {
	return o.Payload
}

func (o *GetFailoverOpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFailoverOpsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFailoverOpsDefault creates a GetFailoverOpsDefault with default headers values
func NewGetFailoverOpsDefault(code int) *GetFailoverOpsDefault {
	return &GetFailoverOpsDefault{
		_statusCode: code,
	}
}

/*
GetFailoverOpsDefault describes a response with status code -1, with default header values.

Error
*/
type GetFailoverOpsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get failover ops default response has a 2xx status code
func (o *GetFailoverOpsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get failover ops default response has a 3xx status code
func (o *GetFailoverOpsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get failover ops default response has a 4xx status code
func (o *GetFailoverOpsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get failover ops default response has a 5xx status code
func (o *GetFailoverOpsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get failover ops default response a status code equal to that given
func (o *GetFailoverOpsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get failover ops default response
func (o *GetFailoverOpsDefault) Code() int {
	return o._statusCode
}

func (o *GetFailoverOpsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/failover/views/{id}/operations][%d] GetFailoverOps default %s", o._statusCode, payload)
}

func (o *GetFailoverOpsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /data-protect/failover/views/{id}/operations][%d] GetFailoverOps default %s", o._statusCode, payload)
}

func (o *GetFailoverOpsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFailoverOpsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package alert

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

// GetAlertSummaryReader is a Reader for the GetAlertSummary structure.
type GetAlertSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAlertSummaryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertSummaryOK creates a GetAlertSummaryOK with default headers values
func NewGetAlertSummaryOK() *GetAlertSummaryOK {
	return &GetAlertSummaryOK{}
}

/*
GetAlertSummaryOK describes a response with status code 200, with default header values.

Success
*/
type GetAlertSummaryOK struct {
	Payload *models.AlertsSummaryResponse
}

// IsSuccess returns true when this get alert summary o k response has a 2xx status code
func (o *GetAlertSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get alert summary o k response has a 3xx status code
func (o *GetAlertSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get alert summary o k response has a 4xx status code
func (o *GetAlertSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get alert summary o k response has a 5xx status code
func (o *GetAlertSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get alert summary o k response a status code equal to that given
func (o *GetAlertSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get alert summary o k response
func (o *GetAlertSummaryOK) Code() int {
	return 200
}

func (o *GetAlertSummaryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /alerts-summary][%d] getAlertSummaryOK %s", 200, payload)
}

func (o *GetAlertSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /alerts-summary][%d] getAlertSummaryOK %s", 200, payload)
}

func (o *GetAlertSummaryOK) GetPayload() *models.AlertsSummaryResponse {
	return o.Payload
}

func (o *GetAlertSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertsSummaryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertSummaryDefault creates a GetAlertSummaryDefault with default headers values
func NewGetAlertSummaryDefault(code int) *GetAlertSummaryDefault {
	return &GetAlertSummaryDefault{
		_statusCode: code,
	}
}

/*
GetAlertSummaryDefault describes a response with status code -1, with default header values.

Error
*/
type GetAlertSummaryDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get alert summary default response has a 2xx status code
func (o *GetAlertSummaryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get alert summary default response has a 3xx status code
func (o *GetAlertSummaryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get alert summary default response has a 4xx status code
func (o *GetAlertSummaryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get alert summary default response has a 5xx status code
func (o *GetAlertSummaryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get alert summary default response a status code equal to that given
func (o *GetAlertSummaryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get alert summary default response
func (o *GetAlertSummaryDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertSummaryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /alerts-summary][%d] GetAlertSummary default %s", o._statusCode, payload)
}

func (o *GetAlertSummaryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /alerts-summary][%d] GetAlertSummary default %s", o._statusCode, payload)
}

func (o *GetAlertSummaryDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAlertSummaryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

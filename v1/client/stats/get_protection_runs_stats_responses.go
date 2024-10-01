// Code generated by go-swagger; DO NOT EDIT.

package stats

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

// GetProtectionRunsStatsReader is a Reader for the GetProtectionRunsStats structure.
type GetProtectionRunsStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProtectionRunsStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProtectionRunsStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProtectionRunsStatsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProtectionRunsStatsOK creates a GetProtectionRunsStatsOK with default headers values
func NewGetProtectionRunsStatsOK() *GetProtectionRunsStatsOK {
	return &GetProtectionRunsStatsOK{}
}

/*
GetProtectionRunsStatsOK describes a response with status code 200, with default header values.

Success
*/
type GetProtectionRunsStatsOK struct {
	Payload *models.ProtectionRunsStats
}

// IsSuccess returns true when this get protection runs stats o k response has a 2xx status code
func (o *GetProtectionRunsStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get protection runs stats o k response has a 3xx status code
func (o *GetProtectionRunsStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get protection runs stats o k response has a 4xx status code
func (o *GetProtectionRunsStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get protection runs stats o k response has a 5xx status code
func (o *GetProtectionRunsStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get protection runs stats o k response a status code equal to that given
func (o *GetProtectionRunsStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get protection runs stats o k response
func (o *GetProtectionRunsStatsOK) Code() int {
	return 200
}

func (o *GetProtectionRunsStatsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/stats/protectionRuns][%d] getProtectionRunsStatsOK %s", 200, payload)
}

func (o *GetProtectionRunsStatsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/stats/protectionRuns][%d] getProtectionRunsStatsOK %s", 200, payload)
}

func (o *GetProtectionRunsStatsOK) GetPayload() *models.ProtectionRunsStats {
	return o.Payload
}

func (o *GetProtectionRunsStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionRunsStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProtectionRunsStatsDefault creates a GetProtectionRunsStatsDefault with default headers values
func NewGetProtectionRunsStatsDefault(code int) *GetProtectionRunsStatsDefault {
	return &GetProtectionRunsStatsDefault{
		_statusCode: code,
	}
}

/*
GetProtectionRunsStatsDefault describes a response with status code -1, with default header values.

Error
*/
type GetProtectionRunsStatsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get protection runs stats default response has a 2xx status code
func (o *GetProtectionRunsStatsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get protection runs stats default response has a 3xx status code
func (o *GetProtectionRunsStatsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get protection runs stats default response has a 4xx status code
func (o *GetProtectionRunsStatsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get protection runs stats default response has a 5xx status code
func (o *GetProtectionRunsStatsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get protection runs stats default response a status code equal to that given
func (o *GetProtectionRunsStatsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get protection runs stats default response
func (o *GetProtectionRunsStatsDefault) Code() int {
	return o._statusCode
}

func (o *GetProtectionRunsStatsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/stats/protectionRuns][%d] GetProtectionRunsStats default %s", o._statusCode, payload)
}

func (o *GetProtectionRunsStatsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /public/stats/protectionRuns][%d] GetProtectionRunsStats default %s", o._statusCode, payload)
}

func (o *GetProtectionRunsStatsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetProtectionRunsStatsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

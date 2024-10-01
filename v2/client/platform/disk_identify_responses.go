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

	"github.com/satwik-miyyapuram-cohesity/go-sdk/v2/models"
)

// DiskIdentifyReader is a Reader for the DiskIdentify structure.
type DiskIdentifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiskIdentifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiskIdentifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDiskIdentifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDiskIdentifyOK creates a DiskIdentifyOK with default headers values
func NewDiskIdentifyOK() *DiskIdentifyOK {
	return &DiskIdentifyOK{}
}

/*
DiskIdentifyOK describes a response with status code 200, with default header values.

Success
*/
type DiskIdentifyOK struct {
	Payload *models.DiskIdentify
}

// IsSuccess returns true when this disk identify o k response has a 2xx status code
func (o *DiskIdentifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this disk identify o k response has a 3xx status code
func (o *DiskIdentifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this disk identify o k response has a 4xx status code
func (o *DiskIdentifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this disk identify o k response has a 5xx status code
func (o *DiskIdentifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this disk identify o k response a status code equal to that given
func (o *DiskIdentifyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the disk identify o k response
func (o *DiskIdentifyOK) Code() int {
	return 200
}

func (o *DiskIdentifyOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disks/identify][%d] diskIdentifyOK %s", 200, payload)
}

func (o *DiskIdentifyOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disks/identify][%d] diskIdentifyOK %s", 200, payload)
}

func (o *DiskIdentifyOK) GetPayload() *models.DiskIdentify {
	return o.Payload
}

func (o *DiskIdentifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiskIdentify)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDiskIdentifyDefault creates a DiskIdentifyDefault with default headers values
func NewDiskIdentifyDefault(code int) *DiskIdentifyDefault {
	return &DiskIdentifyDefault{
		_statusCode: code,
	}
}

/*
DiskIdentifyDefault describes a response with status code -1, with default header values.

Error
*/
type DiskIdentifyDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this disk identify default response has a 2xx status code
func (o *DiskIdentifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this disk identify default response has a 3xx status code
func (o *DiskIdentifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this disk identify default response has a 4xx status code
func (o *DiskIdentifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this disk identify default response has a 5xx status code
func (o *DiskIdentifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this disk identify default response a status code equal to that given
func (o *DiskIdentifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the disk identify default response
func (o *DiskIdentifyDefault) Code() int {
	return o._statusCode
}

func (o *DiskIdentifyDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disks/identify][%d] DiskIdentify default %s", o._statusCode, payload)
}

func (o *DiskIdentifyDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /disks/identify][%d] DiskIdentify default %s", o._statusCode, payload)
}

func (o *DiskIdentifyDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DiskIdentifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

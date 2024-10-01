// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// ExpandPhysicalClusterReader is a Reader for the ExpandPhysicalCluster structure.
type ExpandPhysicalClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpandPhysicalClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewExpandPhysicalClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExpandPhysicalClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExpandPhysicalClusterAccepted creates a ExpandPhysicalClusterAccepted with default headers values
func NewExpandPhysicalClusterAccepted() *ExpandPhysicalClusterAccepted {
	return &ExpandPhysicalClusterAccepted{}
}

/*
ExpandPhysicalClusterAccepted describes a response with status code 202, with default header values.

Success
*/
type ExpandPhysicalClusterAccepted struct {
	Payload *models.CreateClusterResult
}

// IsSuccess returns true when this expand physical cluster accepted response has a 2xx status code
func (o *ExpandPhysicalClusterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this expand physical cluster accepted response has a 3xx status code
func (o *ExpandPhysicalClusterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this expand physical cluster accepted response has a 4xx status code
func (o *ExpandPhysicalClusterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this expand physical cluster accepted response has a 5xx status code
func (o *ExpandPhysicalClusterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this expand physical cluster accepted response a status code equal to that given
func (o *ExpandPhysicalClusterAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the expand physical cluster accepted response
func (o *ExpandPhysicalClusterAccepted) Code() int {
	return 202
}

func (o *ExpandPhysicalClusterAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/physicalEdition/nodes][%d] expandPhysicalClusterAccepted %s", 202, payload)
}

func (o *ExpandPhysicalClusterAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/physicalEdition/nodes][%d] expandPhysicalClusterAccepted %s", 202, payload)
}

func (o *ExpandPhysicalClusterAccepted) GetPayload() *models.CreateClusterResult {
	return o.Payload
}

func (o *ExpandPhysicalClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateClusterResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExpandPhysicalClusterDefault creates a ExpandPhysicalClusterDefault with default headers values
func NewExpandPhysicalClusterDefault(code int) *ExpandPhysicalClusterDefault {
	return &ExpandPhysicalClusterDefault{
		_statusCode: code,
	}
}

/*
ExpandPhysicalClusterDefault describes a response with status code -1, with default header values.

Error
*/
type ExpandPhysicalClusterDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this expand physical cluster default response has a 2xx status code
func (o *ExpandPhysicalClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this expand physical cluster default response has a 3xx status code
func (o *ExpandPhysicalClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this expand physical cluster default response has a 4xx status code
func (o *ExpandPhysicalClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this expand physical cluster default response has a 5xx status code
func (o *ExpandPhysicalClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this expand physical cluster default response a status code equal to that given
func (o *ExpandPhysicalClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the expand physical cluster default response
func (o *ExpandPhysicalClusterDefault) Code() int {
	return o._statusCode
}

func (o *ExpandPhysicalClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/physicalEdition/nodes][%d] ExpandPhysicalCluster default %s", o._statusCode, payload)
}

func (o *ExpandPhysicalClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/physicalEdition/nodes][%d] ExpandPhysicalCluster default %s", o._statusCode, payload)
}

func (o *ExpandPhysicalClusterDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *ExpandPhysicalClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

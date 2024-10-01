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

// ExpandCloudClusterReader is a Reader for the ExpandCloudCluster structure.
type ExpandCloudClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpandCloudClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewExpandCloudClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExpandCloudClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExpandCloudClusterAccepted creates a ExpandCloudClusterAccepted with default headers values
func NewExpandCloudClusterAccepted() *ExpandCloudClusterAccepted {
	return &ExpandCloudClusterAccepted{}
}

/*
ExpandCloudClusterAccepted describes a response with status code 202, with default header values.

Success
*/
type ExpandCloudClusterAccepted struct {
	Payload *models.CreateClusterResult
}

// IsSuccess returns true when this expand cloud cluster accepted response has a 2xx status code
func (o *ExpandCloudClusterAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this expand cloud cluster accepted response has a 3xx status code
func (o *ExpandCloudClusterAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this expand cloud cluster accepted response has a 4xx status code
func (o *ExpandCloudClusterAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this expand cloud cluster accepted response has a 5xx status code
func (o *ExpandCloudClusterAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this expand cloud cluster accepted response a status code equal to that given
func (o *ExpandCloudClusterAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the expand cloud cluster accepted response
func (o *ExpandCloudClusterAccepted) Code() int {
	return 202
}

func (o *ExpandCloudClusterAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/cloudEdition/nodes][%d] expandCloudClusterAccepted %s", 202, payload)
}

func (o *ExpandCloudClusterAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/cloudEdition/nodes][%d] expandCloudClusterAccepted %s", 202, payload)
}

func (o *ExpandCloudClusterAccepted) GetPayload() *models.CreateClusterResult {
	return o.Payload
}

func (o *ExpandCloudClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateClusterResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExpandCloudClusterDefault creates a ExpandCloudClusterDefault with default headers values
func NewExpandCloudClusterDefault(code int) *ExpandCloudClusterDefault {
	return &ExpandCloudClusterDefault{
		_statusCode: code,
	}
}

/*
ExpandCloudClusterDefault describes a response with status code -1, with default header values.

Error
*/
type ExpandCloudClusterDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this expand cloud cluster default response has a 2xx status code
func (o *ExpandCloudClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this expand cloud cluster default response has a 3xx status code
func (o *ExpandCloudClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this expand cloud cluster default response has a 4xx status code
func (o *ExpandCloudClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this expand cloud cluster default response has a 5xx status code
func (o *ExpandCloudClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this expand cloud cluster default response a status code equal to that given
func (o *ExpandCloudClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the expand cloud cluster default response
func (o *ExpandCloudClusterDefault) Code() int {
	return o._statusCode
}

func (o *ExpandCloudClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/cloudEdition/nodes][%d] ExpandCloudCluster default %s", o._statusCode, payload)
}

func (o *ExpandCloudClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /public/clusters/cloudEdition/nodes][%d] ExpandCloudCluster default %s", o._statusCode, payload)
}

func (o *ExpandCloudClusterDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *ExpandCloudClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

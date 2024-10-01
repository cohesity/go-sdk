// Code generated by go-swagger; DO NOT EDIT.

package node_group

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

// GetNodeGroupByNameReader is a Reader for the GetNodeGroupByName structure.
type GetNodeGroupByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeGroupByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeGroupByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetNodeGroupByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeGroupByNameOK creates a GetNodeGroupByNameOK with default headers values
func NewGetNodeGroupByNameOK() *GetNodeGroupByNameOK {
	return &GetNodeGroupByNameOK{}
}

/*
GetNodeGroupByNameOK describes a response with status code 200, with default header values.

Success
*/
type GetNodeGroupByNameOK struct {
	Payload *models.NodeGroupResponse
}

// IsSuccess returns true when this get node group by name o k response has a 2xx status code
func (o *GetNodeGroupByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get node group by name o k response has a 3xx status code
func (o *GetNodeGroupByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get node group by name o k response has a 4xx status code
func (o *GetNodeGroupByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get node group by name o k response has a 5xx status code
func (o *GetNodeGroupByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get node group by name o k response a status code equal to that given
func (o *GetNodeGroupByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get node group by name o k response
func (o *GetNodeGroupByNameOK) Code() int {
	return 200
}

func (o *GetNodeGroupByNameOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /node-groups/{groupName}][%d] getNodeGroupByNameOK %s", 200, payload)
}

func (o *GetNodeGroupByNameOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /node-groups/{groupName}][%d] getNodeGroupByNameOK %s", 200, payload)
}

func (o *GetNodeGroupByNameOK) GetPayload() *models.NodeGroupResponse {
	return o.Payload
}

func (o *GetNodeGroupByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NodeGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeGroupByNameDefault creates a GetNodeGroupByNameDefault with default headers values
func NewGetNodeGroupByNameDefault(code int) *GetNodeGroupByNameDefault {
	return &GetNodeGroupByNameDefault{
		_statusCode: code,
	}
}

/*
GetNodeGroupByNameDefault describes a response with status code -1, with default header values.

Error
*/
type GetNodeGroupByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get node group by name default response has a 2xx status code
func (o *GetNodeGroupByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get node group by name default response has a 3xx status code
func (o *GetNodeGroupByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get node group by name default response has a 4xx status code
func (o *GetNodeGroupByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get node group by name default response has a 5xx status code
func (o *GetNodeGroupByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get node group by name default response a status code equal to that given
func (o *GetNodeGroupByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get node group by name default response
func (o *GetNodeGroupByNameDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeGroupByNameDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /node-groups/{groupName}][%d] GetNodeGroupByName default %s", o._statusCode, payload)
}

func (o *GetNodeGroupByNameDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /node-groups/{groupName}][%d] GetNodeGroupByName default %s", o._statusCode, payload)
}

func (o *GetNodeGroupByNameDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetNodeGroupByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package user

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

// DeleteGroupsReader is a Reader for the DeleteGroups structure.
type DeleteGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGroupsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteGroupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteGroupsNoContent creates a DeleteGroupsNoContent with default headers values
func NewDeleteGroupsNoContent() *DeleteGroupsNoContent {
	return &DeleteGroupsNoContent{}
}

/*
DeleteGroupsNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteGroupsNoContent struct {
}

// IsSuccess returns true when this delete groups no content response has a 2xx status code
func (o *DeleteGroupsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete groups no content response has a 3xx status code
func (o *DeleteGroupsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete groups no content response has a 4xx status code
func (o *DeleteGroupsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete groups no content response has a 5xx status code
func (o *DeleteGroupsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete groups no content response a status code equal to that given
func (o *DeleteGroupsNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete groups no content response
func (o *DeleteGroupsNoContent) Code() int {
	return 204
}

func (o *DeleteGroupsNoContent) Error() string {
	return fmt.Sprintf("[POST /groups/delete][%d] deleteGroupsNoContent", 204)
}

func (o *DeleteGroupsNoContent) String() string {
	return fmt.Sprintf("[POST /groups/delete][%d] deleteGroupsNoContent", 204)
}

func (o *DeleteGroupsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteGroupsDefault creates a DeleteGroupsDefault with default headers values
func NewDeleteGroupsDefault(code int) *DeleteGroupsDefault {
	return &DeleteGroupsDefault{
		_statusCode: code,
	}
}

/*
DeleteGroupsDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteGroupsDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this delete groups default response has a 2xx status code
func (o *DeleteGroupsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete groups default response has a 3xx status code
func (o *DeleteGroupsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete groups default response has a 4xx status code
func (o *DeleteGroupsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete groups default response has a 5xx status code
func (o *DeleteGroupsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete groups default response a status code equal to that given
func (o *DeleteGroupsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete groups default response
func (o *DeleteGroupsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteGroupsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /groups/delete][%d] DeleteGroups default %s", o._statusCode, payload)
}

func (o *DeleteGroupsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /groups/delete][%d] DeleteGroups default %s", o._statusCode, payload)
}

func (o *DeleteGroupsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteGroupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

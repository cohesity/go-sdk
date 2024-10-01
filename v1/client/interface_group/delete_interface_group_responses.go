// Code generated by go-swagger; DO NOT EDIT.

package interface_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cohesity/go-sdk/v1/models"
)

// DeleteInterfaceGroupReader is a Reader for the DeleteInterfaceGroup structure.
type DeleteInterfaceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInterfaceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInterfaceGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteInterfaceGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteInterfaceGroupNoContent creates a DeleteInterfaceGroupNoContent with default headers values
func NewDeleteInterfaceGroupNoContent() *DeleteInterfaceGroupNoContent {
	return &DeleteInterfaceGroupNoContent{}
}

/*
DeleteInterfaceGroupNoContent describes a response with status code 204, with default header values.

No Content
*/
type DeleteInterfaceGroupNoContent struct {
}

// IsSuccess returns true when this delete interface group no content response has a 2xx status code
func (o *DeleteInterfaceGroupNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete interface group no content response has a 3xx status code
func (o *DeleteInterfaceGroupNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete interface group no content response has a 4xx status code
func (o *DeleteInterfaceGroupNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete interface group no content response has a 5xx status code
func (o *DeleteInterfaceGroupNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete interface group no content response a status code equal to that given
func (o *DeleteInterfaceGroupNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete interface group no content response
func (o *DeleteInterfaceGroupNoContent) Code() int {
	return 204
}

func (o *DeleteInterfaceGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /public/interfaceGroups/{name}][%d] deleteInterfaceGroupNoContent", 204)
}

func (o *DeleteInterfaceGroupNoContent) String() string {
	return fmt.Sprintf("[DELETE /public/interfaceGroups/{name}][%d] deleteInterfaceGroupNoContent", 204)
}

func (o *DeleteInterfaceGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInterfaceGroupDefault creates a DeleteInterfaceGroupDefault with default headers values
func NewDeleteInterfaceGroupDefault(code int) *DeleteInterfaceGroupDefault {
	return &DeleteInterfaceGroupDefault{
		_statusCode: code,
	}
}

/*
DeleteInterfaceGroupDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteInterfaceGroupDefault struct {
	_statusCode int

	Payload *models.RequestError
}

// IsSuccess returns true when this delete interface group default response has a 2xx status code
func (o *DeleteInterfaceGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete interface group default response has a 3xx status code
func (o *DeleteInterfaceGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete interface group default response has a 4xx status code
func (o *DeleteInterfaceGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete interface group default response has a 5xx status code
func (o *DeleteInterfaceGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete interface group default response a status code equal to that given
func (o *DeleteInterfaceGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the delete interface group default response
func (o *DeleteInterfaceGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInterfaceGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/interfaceGroups/{name}][%d] DeleteInterfaceGroup default %s", o._statusCode, payload)
}

func (o *DeleteInterfaceGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /public/interfaceGroups/{name}][%d] DeleteInterfaceGroup default %s", o._statusCode, payload)
}

func (o *DeleteInterfaceGroupDefault) GetPayload() *models.RequestError {
	return o.Payload
}

func (o *DeleteInterfaceGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

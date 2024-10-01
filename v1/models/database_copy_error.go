// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatabaseCopyError Database copy error in a DAG.
//
// swagger:model DatabaseCopyError
type DatabaseCopyError struct {

	// Error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`

	// Exchange event log event id of error.
	EventID *int32 `json:"eventId,omitempty"`

	// If DB copy is suspended, reasons for suspension.
	SuspendComment *string `json:"suspendComment,omitempty"`
}

// Validate validates this database copy error
func (m *DatabaseCopyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this database copy error based on context it is used
func (m *DatabaseCopyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseCopyError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseCopyError) UnmarshalBinary(b []byte) error {
	var res DatabaseCopyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

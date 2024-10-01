// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InfectedFileState Specifies the state of infected file.
//
// swagger:model InfectedFileState
type InfectedFileState struct {

	// Specifies the state of infected file.
	// Enum: ["Quarantined","Unquarantined"]
	Type string `json:"type,omitempty"`
}

// Validate validates this infected file state
func (m *InfectedFileState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var infectedFileStateTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Quarantined","Unquarantined"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		infectedFileStateTypeTypePropEnum = append(infectedFileStateTypeTypePropEnum, v)
	}
}

const (

	// InfectedFileStateTypeQuarantined captures enum value "Quarantined"
	InfectedFileStateTypeQuarantined string = "Quarantined"

	// InfectedFileStateTypeUnquarantined captures enum value "Unquarantined"
	InfectedFileStateTypeUnquarantined string = "Unquarantined"
)

// prop value enum
func (m *InfectedFileState) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, infectedFileStateTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InfectedFileState) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this infected file state based on context it is used
func (m *InfectedFileState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InfectedFileState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InfectedFileState) UnmarshalBinary(b []byte) error {
	var res InfectedFileState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

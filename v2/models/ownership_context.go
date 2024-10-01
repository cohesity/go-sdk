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

// OwnershipContext Ownershp Context
//
// Specifies how the External Target or KMS will be consumed internal.
//
// swagger:model OwnershipContext
type OwnershipContext struct {

	// Specifies the ownership context of Vault or KMS.
	// Enum: ["Local","FortKnox"]
	Enum string `json:"enum,omitempty"`
}

// Validate validates this ownership context
func (m *OwnershipContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ownershipContextTypeEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Local","FortKnox"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ownershipContextTypeEnumPropEnum = append(ownershipContextTypeEnumPropEnum, v)
	}
}

const (

	// OwnershipContextEnumLocal captures enum value "Local"
	OwnershipContextEnumLocal string = "Local"

	// OwnershipContextEnumFortKnox captures enum value "FortKnox"
	OwnershipContextEnumFortKnox string = "FortKnox"
)

// prop value enum
func (m *OwnershipContext) validateEnumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ownershipContextTypeEnumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OwnershipContext) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnumEnum("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ownership context based on context it is used
func (m *OwnershipContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OwnershipContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OwnershipContext) UnmarshalBinary(b []byte) error {
	var res OwnershipContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

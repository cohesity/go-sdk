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

// HyperFlexProtectionSource HyperFlex Storae Snapshot.
//
// Specifies a Storage Snapshot Provider in a HyperFlex environment.
//
// swagger:model HyperFlexProtectionSource
type HyperFlexProtectionSource struct {

	// Specifies a unique name of the Protection Source
	Name *string `json:"name,omitempty"`

	// Specifies the product version of the protection source.
	ProductVersion *string `json:"productVersion,omitempty"`

	// Specifies the type of managed Object in a HyperFlex protection source
	// like kServer.
	// Examples of a HyperFlex types include 'kServer'.
	// 'kServer' indicates HyperFlex server entity.
	// Enum: ["kServer"]
	Type *string `json:"type,omitempty"`

	// Specifies the uuid of the protection source.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this hyper flex protection source
func (m *HyperFlexProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hyperFlexProtectionSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kServer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperFlexProtectionSourceTypeTypePropEnum = append(hyperFlexProtectionSourceTypeTypePropEnum, v)
	}
}

const (

	// HyperFlexProtectionSourceTypeKServer captures enum value "kServer"
	HyperFlexProtectionSourceTypeKServer string = "kServer"
)

// prop value enum
func (m *HyperFlexProtectionSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hyperFlexProtectionSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HyperFlexProtectionSource) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this hyper flex protection source based on context it is used
func (m *HyperFlexProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HyperFlexProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperFlexProtectionSource) UnmarshalBinary(b []byte) error {
	var res HyperFlexProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// CsrKeyType CsrKeyType
//
// # Csr Key Type
//
// swagger:model CsrKeyType
type CsrKeyType struct {

	// Specifies the csr key type.
	// Enum: ["rsa","ecdsa"]
	CsrKeyType string `json:"csrKeyType,omitempty"`
}

// Validate validates this csr key type
func (m *CsrKeyType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCsrKeyType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var csrKeyTypeTypeCsrKeyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rsa","ecdsa"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		csrKeyTypeTypeCsrKeyTypePropEnum = append(csrKeyTypeTypeCsrKeyTypePropEnum, v)
	}
}

const (

	// CsrKeyTypeCsrKeyTypeRsa captures enum value "rsa"
	CsrKeyTypeCsrKeyTypeRsa string = "rsa"

	// CsrKeyTypeCsrKeyTypeEcdsa captures enum value "ecdsa"
	CsrKeyTypeCsrKeyTypeEcdsa string = "ecdsa"
)

// prop value enum
func (m *CsrKeyType) validateCsrKeyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, csrKeyTypeTypeCsrKeyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CsrKeyType) validateCsrKeyType(formats strfmt.Registry) error {
	if swag.IsZero(m.CsrKeyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCsrKeyTypeEnum("csrKeyType", "body", m.CsrKeyType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this csr key type based on context it is used
func (m *CsrKeyType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CsrKeyType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CsrKeyType) UnmarshalBinary(b []byte) error {
	var res CsrKeyType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

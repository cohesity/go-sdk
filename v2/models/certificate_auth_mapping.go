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

// CertificateAuthMapping Certificate Auth Mapping
//
// Describes certificate mapping options for certificate based
// authentication.
//
// swagger:model CertificateAuthMapping
type CertificateAuthMapping struct {

	// Specifies the auth field options
	// Enum: ["CommonName","EmailAddress","UserPrincipalName"]
	Enum string `json:"enum,omitempty"`
}

// Validate validates this certificate auth mapping
func (m *CertificateAuthMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnum(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var certificateAuthMappingTypeEnumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CommonName","EmailAddress","UserPrincipalName"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		certificateAuthMappingTypeEnumPropEnum = append(certificateAuthMappingTypeEnumPropEnum, v)
	}
}

const (

	// CertificateAuthMappingEnumCommonName captures enum value "CommonName"
	CertificateAuthMappingEnumCommonName string = "CommonName"

	// CertificateAuthMappingEnumEmailAddress captures enum value "EmailAddress"
	CertificateAuthMappingEnumEmailAddress string = "EmailAddress"

	// CertificateAuthMappingEnumUserPrincipalName captures enum value "UserPrincipalName"
	CertificateAuthMappingEnumUserPrincipalName string = "UserPrincipalName"
)

// prop value enum
func (m *CertificateAuthMapping) validateEnumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, certificateAuthMappingTypeEnumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CertificateAuthMapping) validateEnum(formats strfmt.Registry) error {
	if swag.IsZero(m.Enum) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnumEnum("enum", "body", m.Enum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate auth mapping based on context it is used
func (m *CertificateAuthMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateAuthMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateAuthMapping) UnmarshalBinary(b []byte) error {
	var res CertificateAuthMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

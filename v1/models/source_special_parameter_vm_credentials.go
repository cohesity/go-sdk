// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SourceSpecialParameterVMCredentials VM Credentials.
//
// Specifies the administrator credentials to log in to the
// guest Windows system of a VM that hosts the Microsoft Exchange Server.
// If truncateExchangeLog is set to true and the specified source
// is a VM, administrator credentials to log in to the guest Windows system
// of the VM must be provided to truncate the logs.
// This field is only applicable to Sources in the kVMware environment.
// This field is deprecated. Use the field in VmCredentials inside
// source specific parameter.
// deprecated: true
//
// swagger:model sourceSpecialParameterVmCredentials
type SourceSpecialParameterVMCredentials struct {
	Credentials
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SourceSpecialParameterVMCredentials) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Credentials
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Credentials = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SourceSpecialParameterVMCredentials) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Credentials)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this source special parameter Vm credentials
func (m *SourceSpecialParameterVMCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Credentials
	if err := m.Credentials.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this source special parameter Vm credentials based on the context it is used
func (m *SourceSpecialParameterVMCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Credentials
	if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SourceSpecialParameterVMCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceSpecialParameterVMCredentials) UnmarshalBinary(b []byte) error {
	var res SourceSpecialParameterVMCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

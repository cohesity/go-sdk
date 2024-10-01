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

// UpdateProtectionSourceParametersNasMountCredentials NAS Server Credentials.
//
// Specifies the server credentials to connect to a NetApp server.
// This field is required for mounting SMB volumes on NetApp servers.
//
// swagger:model updateProtectionSourceParametersNasMountCredentials
type UpdateProtectionSourceParametersNasMountCredentials struct {
	NasMountCredentialParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UpdateProtectionSourceParametersNasMountCredentials) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NasMountCredentialParams
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NasMountCredentialParams = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UpdateProtectionSourceParametersNasMountCredentials) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.NasMountCredentialParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update protection source parameters nas mount credentials
func (m *UpdateProtectionSourceParametersNasMountCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NasMountCredentialParams
	if err := m.NasMountCredentialParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this update protection source parameters nas mount credentials based on the context it is used
func (m *UpdateProtectionSourceParametersNasMountCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NasMountCredentialParams
	if err := m.NasMountCredentialParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProtectionSourceParametersNasMountCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProtectionSourceParametersNasMountCredentials) UnmarshalBinary(b []byte) error {
	var res UpdateProtectionSourceParametersNasMountCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

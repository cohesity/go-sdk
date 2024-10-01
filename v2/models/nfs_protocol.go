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

// NfsProtocol NFS Protocol type.
//
// NFS Protocol type.
//
// swagger:model NfsProtocol
type NfsProtocol struct {

	// Specifies Nfs Protocol type.
	// Enum: ["kNfs3","kNfs4_1"]
	NfsProtocol string `json:"nfsProtocol,omitempty"`
}

// Validate validates this nfs protocol
func (m *NfsProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNfsProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nfsProtocolTypeNfsProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNfs3","kNfs4_1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nfsProtocolTypeNfsProtocolPropEnum = append(nfsProtocolTypeNfsProtocolPropEnum, v)
	}
}

const (

	// NfsProtocolNfsProtocolKNfs3 captures enum value "kNfs3"
	NfsProtocolNfsProtocolKNfs3 string = "kNfs3"

	// NfsProtocolNfsProtocolKNfs41 captures enum value "kNfs4_1"
	NfsProtocolNfsProtocolKNfs41 string = "kNfs4_1"
)

// prop value enum
func (m *NfsProtocol) validateNfsProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nfsProtocolTypeNfsProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NfsProtocol) validateNfsProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateNfsProtocolEnum("nfsProtocol", "body", m.NfsProtocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nfs protocol based on context it is used
func (m *NfsProtocol) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NfsProtocol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NfsProtocol) UnmarshalBinary(b []byte) error {
	var res NfsProtocol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

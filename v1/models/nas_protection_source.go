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

// NasProtectionSource Generic NAS Protection Source.
//
// Specifies a Protection Source in a Generic NAS environment.
//
// swagger:model NasProtectionSource
type NasProtectionSource struct {

	// Specifies a description about the Protection Source.
	Description *string `json:"description,omitempty"`

	// Specifies the mount path of this NAS. For example, for a NFS mount point,
	// this should be in the format of IP or hostname:/foo/bar.
	MountPath *string `json:"mountPath,omitempty"`

	// Specifies the name of the NetApp Object.
	Name *string `json:"name,omitempty"`

	// Specifies the protocol used by the NAS server.
	// Specifies the protocol used by a NAS server.
	// 'kNoProtocol' indicates no protocol set.
	// 'kNfs3' indicates NFS v3 protocol.
	// 'kNfs4_1' indicates NFS v4.1 protocol.
	// 'kCifs1' indicates CIFS v1.0 protocol.
	// 'kCifs2' indicates CIFS v2.0 protocol.
	// 'kCifs3' indicates CIFS v3.0 protocol.
	// Enum: ["kNoProtocol","kNfs3","kNfs4_1","kCifs1","kCifs2","kCifs3"]
	Protocol *string `json:"protocol,omitempty"`

	// Specifies whether to skip validation of the given mount point.
	SkipValidation *bool `json:"skipValidation,omitempty"`

	// Specifies the type of a Protection Source Object in a generic NAS Source
	// such as 'kGroup', or 'kHost'.
	// Specifies the kind of NAS mount.
	// 'kGroup' indicates top level node that holds individual NAS hosts.
	// 'kHost' indicates a single NAS path that can be mounted.
	// Enum: ["kGroup","kHost"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this nas protection source
func (m *NasProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nasProtectionSourceTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNoProtocol","kNfs3","kNfs4_1","kCifs1","kCifs2","kCifs3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nasProtectionSourceTypeProtocolPropEnum = append(nasProtectionSourceTypeProtocolPropEnum, v)
	}
}

const (

	// NasProtectionSourceProtocolKNoProtocol captures enum value "kNoProtocol"
	NasProtectionSourceProtocolKNoProtocol string = "kNoProtocol"

	// NasProtectionSourceProtocolKNfs3 captures enum value "kNfs3"
	NasProtectionSourceProtocolKNfs3 string = "kNfs3"

	// NasProtectionSourceProtocolKNfs41 captures enum value "kNfs4_1"
	NasProtectionSourceProtocolKNfs41 string = "kNfs4_1"

	// NasProtectionSourceProtocolKCifs1 captures enum value "kCifs1"
	NasProtectionSourceProtocolKCifs1 string = "kCifs1"

	// NasProtectionSourceProtocolKCifs2 captures enum value "kCifs2"
	NasProtectionSourceProtocolKCifs2 string = "kCifs2"

	// NasProtectionSourceProtocolKCifs3 captures enum value "kCifs3"
	NasProtectionSourceProtocolKCifs3 string = "kCifs3"
)

// prop value enum
func (m *NasProtectionSource) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nasProtectionSourceTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NasProtectionSource) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

var nasProtectionSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kGroup","kHost"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nasProtectionSourceTypeTypePropEnum = append(nasProtectionSourceTypeTypePropEnum, v)
	}
}

const (

	// NasProtectionSourceTypeKGroup captures enum value "kGroup"
	NasProtectionSourceTypeKGroup string = "kGroup"

	// NasProtectionSourceTypeKHost captures enum value "kHost"
	NasProtectionSourceTypeKHost string = "kHost"
)

// prop value enum
func (m *NasProtectionSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nasProtectionSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NasProtectionSource) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this nas protection source based on context it is used
func (m *NasProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NasProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NasProtectionSource) UnmarshalBinary(b []byte) error {
	var res NasProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

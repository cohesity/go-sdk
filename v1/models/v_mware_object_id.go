// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMwareObjectID VMware Object Id.
//
// Specifies a unique Protection Source id across Cohesity Clusters.
// It is derived from the id of the VMware Protection Source.
//
// swagger:model VMwareObjectId
type VMwareObjectID struct {

	// Specifies a UUID for the BIOS of a VMware object. This field will be
	// populated only for VMware VMs and if the VMware source had been registered
	// with the option to track VMs by their BIOS UUID.
	BiosUUID *string `json:"biosUuid,omitempty"`

	// Specifies the Managed Object Reference Item.
	MorItem *string `json:"morItem,omitempty"`

	// Specifies the Managed Object Reference Type.
	MorType *string `json:"morType,omitempty"`

	// Specifies a Universally Unique Identifier (UUID) of a VMware Object.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this v mware object Id
func (m *VMwareObjectID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v mware object Id based on context it is used
func (m *VMwareObjectID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMwareObjectID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMwareObjectID) UnmarshalBinary(b []byte) error {
	var res VMwareObjectID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VlanParams Contains vlan params associated with the backup/restore operation.
//
// swagger:model VlanParams
type VlanParams struct {

	// If this is set to true, then even if VLANs are configured on the system,
	// the partition VIPs will be used for the restore.
	DisableVlan *bool `json:"disableVlan,omitempty"`

	// Interface group to use for restore. If this is not specified, primary
	// interface group for the cluster will be used.
	InterfaceName *string `json:"interfaceName,omitempty"`

	// If this is set, then the Cohesity host name or the IP address associated
	// with this vlan is used for mounting Cohesity's view on the remote host.
	VlanID *int32 `json:"vlanId,omitempty"`
}

// Validate validates this vlan params
func (m *VlanParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this vlan params based on context it is used
func (m *VlanParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VlanParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VlanParams) UnmarshalBinary(b []byte) error {
	var res VlanParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

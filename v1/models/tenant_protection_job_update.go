// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TenantProtectionJobUpdate Tenant Protection Job Update.
//
// Specifies protection Job update details response about a tenant.
//
// swagger:model TenantProtectionJobUpdate
type TenantProtectionJobUpdate struct {

	// Specifies the ProtectionJobIds vec for respective tenant.
	ProtectionJobIds []int64 `json:"protectionJobIds"`

	// Specifies the unique id of the tenant.
	TenantID *string `json:"tenantId,omitempty"`
}

// Validate validates this tenant protection job update
func (m *TenantProtectionJobUpdate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tenant protection job update based on context it is used
func (m *TenantProtectionJobUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TenantProtectionJobUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TenantProtectionJobUpdate) UnmarshalBinary(b []byte) error {
	var res TenantProtectionJobUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

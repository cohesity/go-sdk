// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EntityPermissionInfo Message describing permissions that are defined for an entity.
//
// swagger:model EntityPermissionInfo
type EntityPermissionInfo struct {

	// Entity id.
	EntityID *int64 `json:"entityId,omitempty"`

	// If the permission was inferred. Inferred means, an explicit permission
	// was not set on the current entity. But this entity inherited permissions
	// by virtue of a permission assignment to an ancestor. This bit is
	// populated only when sending back the permission info result to iris.
	IsInferred *bool `json:"isInferred,omitempty"`

	// Specifies whether this object was originally registered by the SP or not.
	//
	// This is required so that we can differentiate between objects registered
	// before and after upgrade. For all objects registered pre-6.7,
	// registering_tenant_id would be nil and this can lead us to believe that
	// the object was originally registered by the SP, which might not be the
	// case. To avoid this, we will use registering_tenant_id in combination with
	// this field.
	//
	// This gets set when a registered entity is assigned to a tenant by the SP.
	IsRegisteredBySp *bool `json:"isRegisteredBySp,omitempty"`

	// Time when permissions for this entity was last updated.
	LastModificationTimeUsecs *int64 `json:"lastModificationTimeUsecs,omitempty"`

	// List of permissions defined on the entity.
	Permissions []*EntityPermissionInfoPermission `json:"permissions"`

	// Tenant ID of the user that registered this object. This will only be
	// populated if this is a root entity.
	RegisteringTenantID *string `json:"registeringTenantId,omitempty"`

	// Tenant ID, if any, that owns this object.
	TenantID *string `json:"tenantId,omitempty"`
}

// Validate validates this entity permission info
func (m *EntityPermissionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityPermissionInfo) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this entity permission info based on the context it is used
func (m *EntityPermissionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EntityPermissionInfo) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Permissions); i++ {

		if m.Permissions[i] != nil {

			if swag.IsZero(m.Permissions[i]) { // not required
				return nil
			}

			if err := m.Permissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EntityPermissionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityPermissionInfo) UnmarshalBinary(b []byte) error {
	var res EntityPermissionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// DeactivateJobEntitiesInfoProto Message that encapsulates information about the deactivate of entities
// protected by the backup job for any of the environments we support.
//
// Please note that for VMware env, we do not need any environment specific
// information. Thus extensions are not yet defined.
//
// swagger:model DeactivateJobEntitiesInfoProto
type DeactivateJobEntitiesInfoProto struct {

	// Contains information about the deactivated entities.
	DeactivatedEntityVec []*DeactivateJobEntitiesInfoProtoDeactivatedEntity `json:"deactivatedEntityVec"`

	// The type of environment this deactivate backup job info pertains to.
	Type *int32 `json:"type,omitempty"`
}

// Validate validates this deactivate job entities info proto
func (m *DeactivateJobEntitiesInfoProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeactivatedEntityVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeactivateJobEntitiesInfoProto) validateDeactivatedEntityVec(formats strfmt.Registry) error {
	if swag.IsZero(m.DeactivatedEntityVec) { // not required
		return nil
	}

	for i := 0; i < len(m.DeactivatedEntityVec); i++ {
		if swag.IsZero(m.DeactivatedEntityVec[i]) { // not required
			continue
		}

		if m.DeactivatedEntityVec[i] != nil {
			if err := m.DeactivatedEntityVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deactivatedEntityVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deactivatedEntityVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this deactivate job entities info proto based on the context it is used
func (m *DeactivateJobEntitiesInfoProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeactivatedEntityVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeactivateJobEntitiesInfoProto) contextValidateDeactivatedEntityVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DeactivatedEntityVec); i++ {

		if m.DeactivatedEntityVec[i] != nil {

			if swag.IsZero(m.DeactivatedEntityVec[i]) { // not required
				return nil
			}

			if err := m.DeactivatedEntityVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deactivatedEntityVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("deactivatedEntityVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeactivateJobEntitiesInfoProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeactivateJobEntitiesInfoProto) UnmarshalBinary(b []byte) error {
	var res DeactivateJobEntitiesInfoProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

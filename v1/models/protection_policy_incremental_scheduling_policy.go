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

// ProtectionPolicyIncrementalSchedulingPolicy CBT-based Job Policy.
//
// Specifies the CBT-based backup schedule of a Protection Job and
// how long Snapshots captured by this schedule are
// retained on the Cohesity Cluster.
//
// swagger:model protectionPolicyIncrementalSchedulingPolicy
type ProtectionPolicyIncrementalSchedulingPolicy struct {
	SchedulingPolicy
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionPolicyIncrementalSchedulingPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SchedulingPolicy
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SchedulingPolicy = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionPolicyIncrementalSchedulingPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SchedulingPolicy)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection policy incremental scheduling policy
func (m *ProtectionPolicyIncrementalSchedulingPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SchedulingPolicy
	if err := m.SchedulingPolicy.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this protection policy incremental scheduling policy based on the context it is used
func (m *ProtectionPolicyIncrementalSchedulingPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SchedulingPolicy
	if err := m.SchedulingPolicy.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionPolicyIncrementalSchedulingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionPolicyIncrementalSchedulingPolicy) UnmarshalBinary(b []byte) error {
	var res ProtectionPolicyIncrementalSchedulingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

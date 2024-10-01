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

// CopySnapshotParams Message to specify the target along with the expiry time to which the
// snapshots captured by the backup run should be copied to.
//
// swagger:model CopySnapshotParams
type CopySnapshotParams struct {

	// If this is true, then only snapshots from a fully successful
	// run will be considered for being copied. If
	// this is false, then snapshots from a partially successful
	// run will also be eligible to be copied.
	CopyOnlyFullySuccessfulRuns *bool `json:"copyOnlyFullySuccessfulRuns,omitempty"`

	// If this is set to true, place all the objects specified as per entity_vec
	// above on legal hold. If entity_vec is empty or not specified, places the
	// backup run on legal hold.
	EnableLegalHold *bool `json:"enableLegalHold,omitempty"`

	// If this is set to true, remove legal hold from all the objects specified
	// as per entity_vec above. If entity_vec is empty or not specified, removes
	// legal hold from the backup run.
	ReleaseLegalHold *bool `json:"releaseLegalHold,omitempty"`

	// Retention policy for the snapshots after they are copied to the target.
	//
	// a) If the snapshots are being copied for the first time to the target:
	//
	// . If this field not specified, the snapshots will never expire.
	//
	// . If this is specified, the expiry time will be determined as:
	// (duration specified in retention_policy) + (time copy finishes).
	//
	//
	// b) If the snapshots were already copied (either partially or fully) to the
	// target earlier:
	//
	// . If this field is not specified, any existing expiry time of the
	// snapshots at the target will not be updated.
	//
	// . If this is specified and retention_policy.num_days_to_keep is not 0,
	// the new expiry time will be determined as:
	// (duration specified in retention_policy) +
	// max(time copy finishes,
	// time snapshots were previously going to expire).
	//
	// Note that in this case, retention_policy.num_days_to_keep can be
	// negative too, which would imply that the retention at at the target
	// is to be reduced.
	//
	// . If this is specified and retention_policy.num_days_to_keep is 0, the
	// snapshots will be set to expire immediately at the target.
	RetentionPolicy *RetentionPolicyProto `json:"retentionPolicy,omitempty"`

	// The target location where the snapshots should be copied to. Target can be
	// one of kLocal, kRemote and kArchival.
	SnapshotTarget *SnapshotTarget `json:"snapshotTarget,omitempty"`
}

// Validate validates this copy snapshot params
func (m *CopySnapshotParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRetentionPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopySnapshotParams) validateRetentionPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.RetentionPolicy) { // not required
		return nil
	}

	if m.RetentionPolicy != nil {
		if err := m.RetentionPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retentionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retentionPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *CopySnapshotParams) validateSnapshotTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotTarget) { // not required
		return nil
	}

	if m.SnapshotTarget != nil {
		if err := m.SnapshotTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotTarget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotTarget")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this copy snapshot params based on the context it is used
func (m *CopySnapshotParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRetentionPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopySnapshotParams) contextValidateRetentionPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.RetentionPolicy != nil {

		if swag.IsZero(m.RetentionPolicy) { // not required
			return nil
		}

		if err := m.RetentionPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retentionPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retentionPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *CopySnapshotParams) contextValidateSnapshotTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotTarget != nil {

		if swag.IsZero(m.SnapshotTarget) { // not required
			return nil
		}

		if err := m.SnapshotTarget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotTarget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotTarget")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CopySnapshotParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopySnapshotParams) UnmarshalBinary(b []byte) error {
	var res CopySnapshotParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

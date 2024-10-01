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

// UpdateProtectionJobRun Update Protection Job Run.
//
// Specifies a Job Run and the expiration time to update. The expiration
// time defines the retention period for the Job Run and its snapshots.
//
// swagger:model UpdateProtectionJobRun
type UpdateProtectionJobRun struct {

	// Specifies the retention for archival, replication or extended local
	// retention.
	CopyRunTargets []*RunJobSnapshotTarget `json:"copyRunTargets"`

	// job Uid
	JobUID *UpdateProtectionJobRunJobUID `json:"jobUid,omitempty"`

	// Specifies the start time of the Job Run to update. The start time
	// is specified as a Unix epoch Timestamp (in microseconds).
	// This uniquely identifies a snapshot. This parameter is required.
	RunStartTimeUsecs *int64 `json:"runStartTimeUsecs,omitempty"`

	// Specifies the run type of the selected job.
	RunType *string `json:"runType,omitempty"`

	// Ids of the Protection Sources. If this is specified, retention time will
	// only be updated for the sources specified.
	SourceIds []int64 `json:"sourceIds"`
}

// Validate validates this update protection job run
func (m *UpdateProtectionJobRun) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCopyRunTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProtectionJobRun) validateCopyRunTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.CopyRunTargets) { // not required
		return nil
	}

	for i := 0; i < len(m.CopyRunTargets); i++ {
		if swag.IsZero(m.CopyRunTargets[i]) { // not required
			continue
		}

		if m.CopyRunTargets[i] != nil {
			if err := m.CopyRunTargets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyRunTargets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyRunTargets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateProtectionJobRun) validateJobUID(formats strfmt.Registry) error {
	if swag.IsZero(m.JobUID) { // not required
		return nil
	}

	if m.JobUID != nil {
		if err := m.JobUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jobUid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update protection job run based on the context it is used
func (m *UpdateProtectionJobRun) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCopyRunTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJobUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateProtectionJobRun) contextValidateCopyRunTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CopyRunTargets); i++ {

		if m.CopyRunTargets[i] != nil {

			if swag.IsZero(m.CopyRunTargets[i]) { // not required
				return nil
			}

			if err := m.CopyRunTargets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("copyRunTargets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("copyRunTargets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateProtectionJobRun) contextValidateJobUID(ctx context.Context, formats strfmt.Registry) error {

	if m.JobUID != nil {

		if swag.IsZero(m.JobUID) { // not required
			return nil
		}

		if err := m.JobUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("jobUid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateProtectionJobRun) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateProtectionJobRun) UnmarshalBinary(b []byte) error {
	var res UpdateProtectionJobRun
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

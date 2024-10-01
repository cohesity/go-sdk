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

// CancelProtectionJobRunParam Cancel a Protection Job Run.
//
// swagger:model CancelProtectionJobRunParam
type CancelProtectionJobRunParam struct {

	// CopyTaskUid is the Uid of a copy task. If a particular copy task is to be
	// cancelled, this field should be set to the id of that particular copy task.
	// For example, if replication task is to be canceled, CopyTaskUid of the
	// replication task has to be specified.
	CopyTaskUID *UniversalID `json:"copyTaskUid,omitempty"`

	// Run Id of a Protection Job Run that needs to be cancelled. If this Run
	// id does not match the id of an active Run in the Protection job, the job
	// Run is not cancelled and an error will be
	// returned.
	JobRunID *int64 `json:"jobRunId,omitempty"`

	// List of entity ids for which we need to cancel the backup tasks. If this
	// is provided it will not cancel the complete run but will cancel only
	// subset of backup tasks (if backup tasks are cancelled correspoding copy
	// task will also get cancelled). If the backup tasks are completed
	// successfully it will not cancel those backup tasks.
	TaskIDList []int64 `json:"taskIdList"`
}

// Validate validates this cancel protection job run param
func (m *CancelProtectionJobRunParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCopyTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelProtectionJobRunParam) validateCopyTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.CopyTaskUID) { // not required
		return nil
	}

	if m.CopyTaskUID != nil {
		if err := m.CopyTaskUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("copyTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("copyTaskUid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cancel protection job run param based on the context it is used
func (m *CancelProtectionJobRunParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCopyTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CancelProtectionJobRunParam) contextValidateCopyTaskUID(ctx context.Context, formats strfmt.Registry) error {

	if m.CopyTaskUID != nil {

		if swag.IsZero(m.CopyTaskUID) { // not required
			return nil
		}

		if err := m.CopyTaskUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("copyTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("copyTaskUid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CancelProtectionJobRunParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CancelProtectionJobRunParam) UnmarshalBinary(b []byte) error {
	var res CancelProtectionJobRunParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

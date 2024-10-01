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

// RemoteRestoreSnapshotStatus Status of Restore Snapshot Task.
//
// Specifies the status of a restore Snapshot task.
//
// swagger:model RemoteRestoreSnapshotStatus
type RemoteRestoreSnapshotStatus struct {

	// Specifies when the snapshot task started. This time is recorded
	// as a Unix epoch Timestamp (in microseconds).
	SnapshotTaskStartTimeUsecs *int64 `json:"SnapshotTaskStartTimeUsecs,omitempty"`

	// archive task Uid
	ArchiveTaskUID *RemoteRestoreSnapshotStatusArchiveTaskUID `json:"archiveTaskUid,omitempty"`

	// Specifies the error message if the indexing task fails.
	Error *string `json:"error,omitempty"`

	// Specifies the time when the Snapshot expires on the remote Vault.
	// This field is recorded as a Unix epoch Timestamp (in microseconds).
	ExpiryTimeUsecs *int64 `json:"expiryTimeUsecs,omitempty"`

	// Specifies the id of the Job Run that originally captured the Snapshot.
	JobRunID *int64 `json:"jobRunId,omitempty"`

	// Specifies the path to the progress monitor task that tracks the progress
	// of building the index.
	ProgressMonitorTask *string `json:"progressMonitorTask,omitempty"`

	// Specifies when the snapshot task completed. This time is recorded
	// as a Unix epoch Timestamp (in microseconds).
	SnapshotTaskEndTimeUsecs *int64 `json:"snapshotTaskEndTimeUsecs,omitempty"`

	// Specifies the status of the indexing task.
	// 'kJobRunning' indicates that the Job/task is currently running.
	// 'kJobFinished' indicates that the Job/task completed and finished.
	// 'kJobFailed' indicates that the Job/task failed and did not complete.
	// 'kJobCanceled' indicates that the Job/task was canceled.
	// 'kJobPaused' indicates the Job/task is paused.
	// Enum: ["kJobRunning","kJobFinished","kJobFailed","kJobCanceled","kJobPaused"]
	SnapshotTaskStatus *string `json:"snapshotTaskStatus,omitempty"`

	// snapshot task Uid
	SnapshotTaskUID *RemoteRestoreSnapshotStatusSnapshotTaskUID `json:"snapshotTaskUid,omitempty"`

	// Specify the time the Snapshot was captured.
	// This time is recorded as a Unix epoch Timestamp (in microseconds).
	SnapshotTimeUsecs *int64 `json:"snapshotTimeUsecs,omitempty"`
}

// Validate validates this remote restore snapshot status
func (m *RemoteRestoreSnapshotStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchiveTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotTaskStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteRestoreSnapshotStatus) validateArchiveTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchiveTaskUID) { // not required
		return nil
	}

	if m.ArchiveTaskUID != nil {
		if err := m.ArchiveTaskUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archiveTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archiveTaskUid")
			}
			return err
		}
	}

	return nil
}

var remoteRestoreSnapshotStatusTypeSnapshotTaskStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kJobRunning","kJobFinished","kJobFailed","kJobCanceled","kJobPaused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		remoteRestoreSnapshotStatusTypeSnapshotTaskStatusPropEnum = append(remoteRestoreSnapshotStatusTypeSnapshotTaskStatusPropEnum, v)
	}
}

const (

	// RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobRunning captures enum value "kJobRunning"
	RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobRunning string = "kJobRunning"

	// RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobFinished captures enum value "kJobFinished"
	RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobFinished string = "kJobFinished"

	// RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobFailed captures enum value "kJobFailed"
	RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobFailed string = "kJobFailed"

	// RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobCanceled captures enum value "kJobCanceled"
	RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobCanceled string = "kJobCanceled"

	// RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobPaused captures enum value "kJobPaused"
	RemoteRestoreSnapshotStatusSnapshotTaskStatusKJobPaused string = "kJobPaused"
)

// prop value enum
func (m *RemoteRestoreSnapshotStatus) validateSnapshotTaskStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, remoteRestoreSnapshotStatusTypeSnapshotTaskStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RemoteRestoreSnapshotStatus) validateSnapshotTaskStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotTaskStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSnapshotTaskStatusEnum("snapshotTaskStatus", "body", *m.SnapshotTaskStatus); err != nil {
		return err
	}

	return nil
}

func (m *RemoteRestoreSnapshotStatus) validateSnapshotTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotTaskUID) { // not required
		return nil
	}

	if m.SnapshotTaskUID != nil {
		if err := m.SnapshotTaskUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotTaskUid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this remote restore snapshot status based on the context it is used
func (m *RemoteRestoreSnapshotStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchiveTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteRestoreSnapshotStatus) contextValidateArchiveTaskUID(ctx context.Context, formats strfmt.Registry) error {

	if m.ArchiveTaskUID != nil {

		if swag.IsZero(m.ArchiveTaskUID) { // not required
			return nil
		}

		if err := m.ArchiveTaskUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archiveTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archiveTaskUid")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteRestoreSnapshotStatus) contextValidateSnapshotTaskUID(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotTaskUID != nil {

		if swag.IsZero(m.SnapshotTaskUID) { // not required
			return nil
		}

		if err := m.SnapshotTaskUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotTaskUid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteRestoreSnapshotStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteRestoreSnapshotStatus) UnmarshalBinary(b []byte) error {
	var res RemoteRestoreSnapshotStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

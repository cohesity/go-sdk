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

// RemoteRestoreIndexingStatus Status of Indexing Task.
//
// Specifies the status of an indexing task.
//
// swagger:model RemoteRestoreIndexingStatus
type RemoteRestoreIndexingStatus struct {

	// Specifies the end time of the time range that is being indexed.
	// The indexing task is creating an index of the Job Runs that occurred
	// between the startTimeUsecs and this endTimeUsecs.
	// This field is recorded as a Unix epoch Timestamp (in microseconds).
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// Specifies the error message if the indexing Job/task fails.
	Error *string `json:"error,omitempty"`

	// Specifies when the indexing task completed. This time is recorded
	// as a Unix epoch Timestamp (in microseconds).
	// This field is not set if the indexing task is still in progress.
	IndexingTaskEndTimeUsecs *int64 `json:"indexingTaskEndTimeUsecs,omitempty"`

	// Specifies when the indexing task started. This time is recorded
	// as a Unix epoch Timestamp (in microseconds).
	IndexingTaskStartTimeUsecs *int64 `json:"indexingTaskStartTimeUsecs,omitempty"`

	// Specifies the status of the indexing Job/task.
	// 'kJobRunning' indicates that the Job/task is currently running.
	// 'kJobFinished' indicates that the Job/task completed and finished.
	// 'kJobFailed' indicates that the Job/task failed and did not complete.
	// 'kJobCanceled' indicates that the Job/task was canceled.
	// 'kJobPaused' indicates the Job/task is paused.
	// Enum: ["kJobRunning","kJobFinished","kJobFailed","kJobCanceled","kJobPaused"]
	IndexingTaskStatus *string `json:"indexingTaskStatus,omitempty"`

	// Indexing Task Uid.
	//
	// Specifies the unique id of the indexing task assigned by this Cluster.
	IndexingTaskUID struct {
		UniversalID
	} `json:"indexingTaskUid,omitempty"`

	// For all the Snapshots retrieved by this Job, specifies the latest time
	// when a Snapshot expires.
	LatestExpiryTimeUsecs *int64 `json:"latestExpiryTimeUsecs,omitempty"`

	// Specifies the path to progress monitor task to track the progress
	// of building the index.
	ProgressMonitorTask *string `json:"progressMonitorTask,omitempty"`

	// Specifies the start time of the time range that is being indexed.
	// The indexing task is creating an index of the Job Runs that occurred
	// between this startTimeUsecs and the endTimeUsecs.
	// This field is recorded as a Unix epoch Timestamp (in microseconds).
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`
}

// Validate validates this remote restore indexing status
func (m *RemoteRestoreIndexingStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIndexingTaskStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexingTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var remoteRestoreIndexingStatusTypeIndexingTaskStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kJobRunning","kJobFinished","kJobFailed","kJobCanceled","kJobPaused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		remoteRestoreIndexingStatusTypeIndexingTaskStatusPropEnum = append(remoteRestoreIndexingStatusTypeIndexingTaskStatusPropEnum, v)
	}
}

const (

	// RemoteRestoreIndexingStatusIndexingTaskStatusKJobRunning captures enum value "kJobRunning"
	RemoteRestoreIndexingStatusIndexingTaskStatusKJobRunning string = "kJobRunning"

	// RemoteRestoreIndexingStatusIndexingTaskStatusKJobFinished captures enum value "kJobFinished"
	RemoteRestoreIndexingStatusIndexingTaskStatusKJobFinished string = "kJobFinished"

	// RemoteRestoreIndexingStatusIndexingTaskStatusKJobFailed captures enum value "kJobFailed"
	RemoteRestoreIndexingStatusIndexingTaskStatusKJobFailed string = "kJobFailed"

	// RemoteRestoreIndexingStatusIndexingTaskStatusKJobCanceled captures enum value "kJobCanceled"
	RemoteRestoreIndexingStatusIndexingTaskStatusKJobCanceled string = "kJobCanceled"

	// RemoteRestoreIndexingStatusIndexingTaskStatusKJobPaused captures enum value "kJobPaused"
	RemoteRestoreIndexingStatusIndexingTaskStatusKJobPaused string = "kJobPaused"
)

// prop value enum
func (m *RemoteRestoreIndexingStatus) validateIndexingTaskStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, remoteRestoreIndexingStatusTypeIndexingTaskStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RemoteRestoreIndexingStatus) validateIndexingTaskStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingTaskStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateIndexingTaskStatusEnum("indexingTaskStatus", "body", *m.IndexingTaskStatus); err != nil {
		return err
	}

	return nil
}

func (m *RemoteRestoreIndexingStatus) validateIndexingTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.IndexingTaskUID) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this remote restore indexing status based on the context it is used
func (m *RemoteRestoreIndexingStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIndexingTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteRestoreIndexingStatus) contextValidateIndexingTaskUID(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteRestoreIndexingStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteRestoreIndexingStatus) UnmarshalBinary(b []byte) error {
	var res RemoteRestoreIndexingStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

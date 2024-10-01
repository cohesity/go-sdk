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

// PlannedRunPollStatus Specifies whether run has been scheduled or not and also returns the unique run id along with failoverId upon scheduling the run.
//
// swagger:model PlannedRunPollStatus
type PlannedRunPollStatus struct {

	// Specifies the unique failover Id which will be generated by orchestrator. This Id will be used to uniquely identify current failover operation.
	FailoverID *string `json:"failoverId,omitempty"`

	// If cancelNonFailoverRuns was passed as true during creation of run for current failover then this will return the status of other run cacellations. If other runs are still pending for cancellations then this will be returned as true otherwise it will be return as false.
	WaitingOnOtherRunCancellations *bool `json:"waitingOnOtherRunCancellations,omitempty"`

	// If run has been scheduled then this field will be populated with unique run id.
	RunID *string `json:"runId,omitempty"`

	// Specifies the protection group id to which this run belongs.
	ProtectionGroupID *string `json:"protectionGroupId,omitempty"`

	// Status of the backup job. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","Paused"]
	BackupTaskStatus *string `json:"backupTaskStatus,omitempty"`

	// Status of the OnPrem deploy task. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","Paused"]
	OnPremDeployTaskStatus *string `json:"onPremDeployTaskStatus,omitempty"`
}

// Validate validates this planned run poll status
func (m *PlannedRunPollStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupTaskStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnPremDeployTaskStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var plannedRunPollStatusTypeBackupTaskStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		plannedRunPollStatusTypeBackupTaskStatusPropEnum = append(plannedRunPollStatusTypeBackupTaskStatusPropEnum, v)
	}
}

const (

	// PlannedRunPollStatusBackupTaskStatusAccepted captures enum value "Accepted"
	PlannedRunPollStatusBackupTaskStatusAccepted string = "Accepted"

	// PlannedRunPollStatusBackupTaskStatusRunning captures enum value "Running"
	PlannedRunPollStatusBackupTaskStatusRunning string = "Running"

	// PlannedRunPollStatusBackupTaskStatusCanceled captures enum value "Canceled"
	PlannedRunPollStatusBackupTaskStatusCanceled string = "Canceled"

	// PlannedRunPollStatusBackupTaskStatusCanceling captures enum value "Canceling"
	PlannedRunPollStatusBackupTaskStatusCanceling string = "Canceling"

	// PlannedRunPollStatusBackupTaskStatusFailed captures enum value "Failed"
	PlannedRunPollStatusBackupTaskStatusFailed string = "Failed"

	// PlannedRunPollStatusBackupTaskStatusMissed captures enum value "Missed"
	PlannedRunPollStatusBackupTaskStatusMissed string = "Missed"

	// PlannedRunPollStatusBackupTaskStatusSucceeded captures enum value "Succeeded"
	PlannedRunPollStatusBackupTaskStatusSucceeded string = "Succeeded"

	// PlannedRunPollStatusBackupTaskStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	PlannedRunPollStatusBackupTaskStatusSucceededWithWarning string = "SucceededWithWarning"

	// PlannedRunPollStatusBackupTaskStatusOnHold captures enum value "OnHold"
	PlannedRunPollStatusBackupTaskStatusOnHold string = "OnHold"

	// PlannedRunPollStatusBackupTaskStatusFinalizing captures enum value "Finalizing"
	PlannedRunPollStatusBackupTaskStatusFinalizing string = "Finalizing"

	// PlannedRunPollStatusBackupTaskStatusSkipped captures enum value "Skipped"
	PlannedRunPollStatusBackupTaskStatusSkipped string = "Skipped"

	// PlannedRunPollStatusBackupTaskStatusPaused captures enum value "Paused"
	PlannedRunPollStatusBackupTaskStatusPaused string = "Paused"
)

// prop value enum
func (m *PlannedRunPollStatus) validateBackupTaskStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, plannedRunPollStatusTypeBackupTaskStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlannedRunPollStatus) validateBackupTaskStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupTaskStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupTaskStatusEnum("backupTaskStatus", "body", *m.BackupTaskStatus); err != nil {
		return err
	}

	return nil
}

var plannedRunPollStatusTypeOnPremDeployTaskStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped","Paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		plannedRunPollStatusTypeOnPremDeployTaskStatusPropEnum = append(plannedRunPollStatusTypeOnPremDeployTaskStatusPropEnum, v)
	}
}

const (

	// PlannedRunPollStatusOnPremDeployTaskStatusAccepted captures enum value "Accepted"
	PlannedRunPollStatusOnPremDeployTaskStatusAccepted string = "Accepted"

	// PlannedRunPollStatusOnPremDeployTaskStatusRunning captures enum value "Running"
	PlannedRunPollStatusOnPremDeployTaskStatusRunning string = "Running"

	// PlannedRunPollStatusOnPremDeployTaskStatusCanceled captures enum value "Canceled"
	PlannedRunPollStatusOnPremDeployTaskStatusCanceled string = "Canceled"

	// PlannedRunPollStatusOnPremDeployTaskStatusCanceling captures enum value "Canceling"
	PlannedRunPollStatusOnPremDeployTaskStatusCanceling string = "Canceling"

	// PlannedRunPollStatusOnPremDeployTaskStatusFailed captures enum value "Failed"
	PlannedRunPollStatusOnPremDeployTaskStatusFailed string = "Failed"

	// PlannedRunPollStatusOnPremDeployTaskStatusMissed captures enum value "Missed"
	PlannedRunPollStatusOnPremDeployTaskStatusMissed string = "Missed"

	// PlannedRunPollStatusOnPremDeployTaskStatusSucceeded captures enum value "Succeeded"
	PlannedRunPollStatusOnPremDeployTaskStatusSucceeded string = "Succeeded"

	// PlannedRunPollStatusOnPremDeployTaskStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	PlannedRunPollStatusOnPremDeployTaskStatusSucceededWithWarning string = "SucceededWithWarning"

	// PlannedRunPollStatusOnPremDeployTaskStatusOnHold captures enum value "OnHold"
	PlannedRunPollStatusOnPremDeployTaskStatusOnHold string = "OnHold"

	// PlannedRunPollStatusOnPremDeployTaskStatusFinalizing captures enum value "Finalizing"
	PlannedRunPollStatusOnPremDeployTaskStatusFinalizing string = "Finalizing"

	// PlannedRunPollStatusOnPremDeployTaskStatusSkipped captures enum value "Skipped"
	PlannedRunPollStatusOnPremDeployTaskStatusSkipped string = "Skipped"

	// PlannedRunPollStatusOnPremDeployTaskStatusPaused captures enum value "Paused"
	PlannedRunPollStatusOnPremDeployTaskStatusPaused string = "Paused"
)

// prop value enum
func (m *PlannedRunPollStatus) validateOnPremDeployTaskStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, plannedRunPollStatusTypeOnPremDeployTaskStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PlannedRunPollStatus) validateOnPremDeployTaskStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.OnPremDeployTaskStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnPremDeployTaskStatusEnum("onPremDeployTaskStatus", "body", *m.OnPremDeployTaskStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this planned run poll status based on context it is used
func (m *PlannedRunPollStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlannedRunPollStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlannedRunPollStatus) UnmarshalBinary(b []byte) error {
	var res PlannedRunPollStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

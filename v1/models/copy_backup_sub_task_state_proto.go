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

// CopyBackupSubTaskStateProto This message contains persistent information about a sub task that copies
// the snapshot of a leaf level entity in a backup run to a particular target.
//
// swagger:model CopyBackupSubTaskStateProto
type CopyBackupSubTaskStateProto struct {

	// Target for action executor on which the slave task will execute actions.
	// e.g Bridge/Bridge_Proxy.
	ActionExecutorTargetType *int32 `json:"actionExecutorTargetType,omitempty"`

	// If this is a first object level replication task (or if previous object
	// level replication tasks failed) this field will contain the uid of the
	// last replication task (if it exists) performed using the old style (i.e.,
	// all backup tasks in same replication task) protocol for the same target.
	// View associated with the 'ancestor_replication_uid_hint' can be used by
	// Madrox to efficiently compute diffs for this object level replication
	// task.
	AncestorReplicationUIDHint *UniversalIDProto `json:"ancestorReplicationUidHint,omitempty"`

	// Id of the task that performed the backup of the above entity. Due to
	// retries on errors, a leaf level entity can have more than one attempt and
	// thus have more than one snapshot. If job policy dictates that snapshots
	// from even failed attempt/task needs to be copied, then there can be more
	// than one copy sub-task for an entity. This field is used to make sure that
	// only one copy sub-task is created for a backup task.
	BackupTaskID *int64 `json:"backupTaskId,omitempty"`

	// If this is a replication task, this will contain the constituent id of the
	// Bridge instance where this task has been scheduled (if the task has been
	// scheduled).
	BridgeConstituentID *int64 `json:"bridgeConstituentId,omitempty"`

	// Whether this copy task has a pending cancellation request.
	CancellationRequested *bool `json:"cancellationRequested,omitempty"`

	// If this is a cloud deploy sub-task, then the state of the cloud deploy
	// task is stored here.
	CloudDeployTaskState *CloudDeployTaskStateProto `json:"cloudDeployTaskState,omitempty"`

	// Id of the conversion cleanup task. The cleanup task is supposed to delete
	// residual state of the conversion task.
	ConversionCleanupTaskID *int64 `json:"conversionCleanupTaskId,omitempty"`

	// Id of the conversion task. In the event of DR-to-cloud scenario while
	// replicating the VM (that was failed over to cloud) back to on-prem, we
	// convert the on cloud format of the VM back to on-prem format, e.g.
	// on Azure we backup a VM (that came from on-prem VMWare environment) using
	// physical agent that backs up logical volumes as VHD files. While
	// replicating the backup to on-prem we convert the VM backup to VMWare VMDK
	// format.
	ConversionTaskID *int64 `json:"conversionTaskId,omitempty"`

	// The timeout duration taken from the policy(based on full/incremental
	// backup). A cancellation will automatically gets triggered if the copy
	// sub-task has been running for more than this timeout duration.
	CopySubTaskTimeoutMins *int64 `json:"copySubTaskTimeoutMins,omitempty"`

	// Whether to disable creation of MegaFiles on the Rx cluster for a
	// replication sub-task, even if the Rx cluster supports it.
	DisableRxMegafile *bool `json:"disableRxMegafile,omitempty"`

	// The time this task finished.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// Leaf level source entity whose snapshot is being copied by this task.
	Entity *PrivateEntityProto `json:"entity,omitempty"`

	// Leaf level source entities whose snapshot is being copied by this task,
	// This will be populated only if view replication is enabled.
	EntityVec []*PrivateEntityProto `json:"entityVec"`

	// Will contain any error encountered by this task.
	Error *PrivateErrorProto `json:"error,omitempty"`

	// Boolean to tell whether to fail the progress monitor if there an
	// kAlreadyExists error in creating the progress monitor for cloud spin
	// tasks.
	FailProgressMonitorOnAlreadyExistsError *bool `json:"failProgressMonitorOnAlreadyExistsError,omitempty"`

	// Whether this sub_task is the last to be spawned. Applicable for
	// "view_level_copy" replications, where the last replication is supposed to
	// start after the protection run view is finalized.
	IsFinalRequest *bool `json:"isFinalRequest,omitempty"`

	// Whether this is an out of band (OOB) copy sub-task triggered by the user.
	IsOutOfBandSubTask *bool `json:"isOutOfBandSubTask,omitempty"`

	// The instance id of the backup run whose snapshot is to be copied.
	JobInstanceID *int64 `json:"jobInstanceId,omitempty"`

	// The globally unique id of the backup job whose snapshot is to be copied.
	JobUID *UniversalIDProto `json:"jobUid,omitempty"`

	// Whether this task performs granular copy of snapshots at an object level.
	// This field can only be set for snapshot target of type 'kRemote'. If set
	// to true, this task will create and schedule child copy tasks as soon as
	// the backup of an object finishes. Else, this task is expected to wait for
	// the backup run to finish before starting the copy task.
	ObjectLevelCopy *bool `json:"objectLevelCopy,omitempty"`

	// If this is an onprem deploy sub-task, then the state of the onprem deploy
	// task is stored here.
	OnpremDeployTaskState *OnPremDeployTaskStateProto `json:"onpremDeployTaskState,omitempty"`

	// The globally unique id of the parent copy backup run task that spawned
	// this child copy task.
	ParentTaskUID *UniversalIDProto `json:"parentTaskUid,omitempty"`

	// Full path of the Pulse task tracking the progress of this sub task.
	ProgressMonitorTaskPath *string `json:"progressMonitorTaskPath,omitempty"`

	// Iris-facing task state. This field is stamped during the export.
	PublicStatus *int32 `json:"publicStatus,omitempty"`

	// If this is a replication task, this field will contain some basic info
	// about the replication task.
	ReplicationInfo *ReplicationInfoBase `json:"replicationInfo,omitempty"`

	// The start time of the backup run whose snapshot is to be copied.
	RunStartTimeUsecs *int64 `json:"runStartTimeUsecs,omitempty"`

	// Set to true if entity being copied needs to be converted to on-prem format
	// for "DR to cloud" failback use case.
	ShouldConvert *bool `json:"shouldConvert,omitempty"`

	// The target location where the snapshots should be copied to.
	SnapshotTarget *SnapshotTarget `json:"snapshotTarget,omitempty"`

	// The time this task started.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// The status of this task.
	Status *int32 `json:"status,omitempty"`

	// The globally unique id of this task.
	TaskUID *UniversalIDProto `json:"taskUid,omitempty"`

	// This field will only be populated when this proto is returned in an
	// external HTTP RPC response (e.g., for Iris). The copy sub-task returned
	// in the external call is created by merging all the copy sub-tasks for the
	// backup task and target. This field will contain the sum total of the
	// duration of each of the copy sub-tasks that were merged.
	TotalDurationUsecs *int64 `json:"totalDurationUsecs,omitempty"`
}

// Validate validates this copy backup sub task state proto
func (m *CopyBackupSubTaskStateProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAncestorReplicationUIDHint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudDeployTaskState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityVec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnpremDeployTaskState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaskUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyBackupSubTaskStateProto) validateAncestorReplicationUIDHint(formats strfmt.Registry) error {
	if swag.IsZero(m.AncestorReplicationUIDHint) { // not required
		return nil
	}

	if m.AncestorReplicationUIDHint != nil {
		if err := m.AncestorReplicationUIDHint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ancestorReplicationUidHint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ancestorReplicationUidHint")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateCloudDeployTaskState(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudDeployTaskState) { // not required
		return nil
	}

	if m.CloudDeployTaskState != nil {
		if err := m.CloudDeployTaskState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudDeployTaskState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudDeployTaskState")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateEntityVec(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityVec) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityVec); i++ {
		if swag.IsZero(m.EntityVec[i]) { // not required
			continue
		}

		if m.EntityVec[i] != nil {
			if err := m.EntityVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entityVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entityVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateJobUID(formats strfmt.Registry) error {
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

func (m *CopyBackupSubTaskStateProto) validateOnpremDeployTaskState(formats strfmt.Registry) error {
	if swag.IsZero(m.OnpremDeployTaskState) { // not required
		return nil
	}

	if m.OnpremDeployTaskState != nil {
		if err := m.OnpremDeployTaskState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onpremDeployTaskState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onpremDeployTaskState")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateParentTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentTaskUID) { // not required
		return nil
	}

	if m.ParentTaskUID != nil {
		if err := m.ParentTaskUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentTaskUid")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateReplicationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationInfo) { // not required
		return nil
	}

	if m.ReplicationInfo != nil {
		if err := m.ReplicationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) validateSnapshotTarget(formats strfmt.Registry) error {
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

func (m *CopyBackupSubTaskStateProto) validateTaskUID(formats strfmt.Registry) error {
	if swag.IsZero(m.TaskUID) { // not required
		return nil
	}

	if m.TaskUID != nil {
		if err := m.TaskUID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taskUid")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this copy backup sub task state proto based on the context it is used
func (m *CopyBackupSubTaskStateProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAncestorReplicationUIDHint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudDeployTaskState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJobUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOnpremDeployTaskState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReplicationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaskUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateAncestorReplicationUIDHint(ctx context.Context, formats strfmt.Registry) error {

	if m.AncestorReplicationUIDHint != nil {

		if swag.IsZero(m.AncestorReplicationUIDHint) { // not required
			return nil
		}

		if err := m.AncestorReplicationUIDHint.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ancestorReplicationUidHint")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ancestorReplicationUidHint")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateCloudDeployTaskState(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudDeployTaskState != nil {

		if swag.IsZero(m.CloudDeployTaskState) { // not required
			return nil
		}

		if err := m.CloudDeployTaskState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudDeployTaskState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudDeployTaskState")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.Entity != nil {

		if swag.IsZero(m.Entity) { // not required
			return nil
		}

		if err := m.Entity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateEntityVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityVec); i++ {

		if m.EntityVec[i] != nil {

			if swag.IsZero(m.EntityVec[i]) { // not required
				return nil
			}

			if err := m.EntityVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entityVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entityVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {

		if swag.IsZero(m.Error) { // not required
			return nil
		}

		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateJobUID(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CopyBackupSubTaskStateProto) contextValidateOnpremDeployTaskState(ctx context.Context, formats strfmt.Registry) error {

	if m.OnpremDeployTaskState != nil {

		if swag.IsZero(m.OnpremDeployTaskState) { // not required
			return nil
		}

		if err := m.OnpremDeployTaskState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onpremDeployTaskState")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("onpremDeployTaskState")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateParentTaskUID(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentTaskUID != nil {

		if swag.IsZero(m.ParentTaskUID) { // not required
			return nil
		}

		if err := m.ParentTaskUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentTaskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parentTaskUid")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateReplicationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ReplicationInfo != nil {

		if swag.IsZero(m.ReplicationInfo) { // not required
			return nil
		}

		if err := m.ReplicationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replicationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("replicationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CopyBackupSubTaskStateProto) contextValidateSnapshotTarget(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CopyBackupSubTaskStateProto) contextValidateTaskUID(ctx context.Context, formats strfmt.Registry) error {

	if m.TaskUID != nil {

		if swag.IsZero(m.TaskUID) { // not required
			return nil
		}

		if err := m.TaskUID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("taskUid")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("taskUid")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CopyBackupSubTaskStateProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CopyBackupSubTaskStateProto) UnmarshalBinary(b []byte) error {
	var res CopyBackupSubTaskStateProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

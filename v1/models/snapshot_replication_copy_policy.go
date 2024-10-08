// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SnapshotReplicationCopyPolicy Snapshot Copy Replication Policy.
//
// Specifies settings for copying Snapshots to Remote Clusters. This also
// specifies the retention policy that should be applied to Snapshots after
// they have been copied to the specified target.
//
// swagger:model SnapshotReplicationCopyPolicy
type SnapshotReplicationCopyPolicy struct {

	// Specified the Id for a snapshot copy policy. This is generated when the
	// policy is created.
	ID *string `json:"Id,omitempty"`

	// The backup run type to which this copy policy applies to. If set, this
	// will ensure that the first run in scheduled period of given type will be
	// copied. If this isn't set, copy tasks will be generated as per
	// other filters in the protection policy. Currently, it can only be set to
	// Full.
	// 'kRegular' indicates a incremental (CBT) backup. Incremental backups
	// utilizing CBT (if supported) are captured of the target protection objects.
	// The first run of a kRegular schedule captures all the blocks.
	// 'kFull' indicates a full (no CBT) backup. A complete backup
	// (all blocks) of the target protection objects are always captured and
	// Change Block Tracking (CBT) is not utilized.
	// 'kLog' indicates a Database Log backup. Capture the database
	// transaction logs to allow rolling back to a specific point in time.
	// 'kSystem' indicates a system backup. System backups are used to do
	// bare metal recovery of the system to a specific point in time.
	// Enum: ["kRegular","kFull","kLog","kSystem"]
	BackupRunType *string `json:"backupRunType,omitempty"`

	// Specifies the cloud replication target to copy the Snapshots to.
	CloudTarget *CloudDeployTargetDetails `json:"cloudTarget,omitempty"`

	// Specifies if Snapshots are copied from the first completely successful
	// Job Run or the first partially successful Job Run occurring at the start
	// of the replication schedule.
	// If true, Snapshots are copied from the first Job Run occurring at the start
	// of the replication schedule, even if first Job Run was not completely
	// successful i.e. Snapshots were not captured for all Objects in the Job.
	// If false, Snapshots are copied from the first Job Run occurring at the
	// start of the replication schedule that was completely successful i.e.
	// Snapshots for all the Objects in the Job were successfully captured.
	CopyPartial *bool `json:"copyPartial,omitempty"`

	// Specifies WORM retention type for the copy target snapshots. When a WORM
	// retention type is specified, the snapshots of the Protection Groups using
	// this policy will be kept for the last N days as specified in the duration
	// of the datalock. During that time, the snapshots cannot be deleted.
	DatalockConfig *DataLockConfig `json:"datalockConfig,omitempty"`

	// Specifies the number of days to retain copied Snapshots on the target.
	DaysToKeep *int64 `json:"daysToKeep,omitempty"`

	// Specifies the number of days to retain log run if Log Schedule exists on
	// the external target.
	DaysToKeepLog *int64 `json:"daysToKeepLog,omitempty"`

	// Specifies a factor to multiply the periodicity by, to determine the copy
	// schedule.
	// For example if set to 2 and the periodicity is hourly, then Snapshots
	// from the first eligible Job Run for every 2 hour period is copied.
	Multiplier *int32 `json:"multiplier,omitempty"`

	// Specifies the frequency that Snapshots should be copied to the
	// specified target. Used in combination with multiplier.
	// 'kEvery' means that the Snapshot copy occurs after the number of
	// Job Runs equals the number specified in the multiplier.
	// 'kHour' means that the Snapshot copy occurs hourly at the frequency
	// set in the multiplier, for example if multiplier is 2, the copy occurs
	// every 2 hours.
	// 'kDay' means that the Snapshot copy occurs daily at the frequency
	// set in the multiplier.
	// 'kWeek' means that the Snapshot copy occurs weekly at the frequency
	// set in the multiplier.
	// 'kMonth' means that the Snapshot copy occurs monthly at the frequency
	// set in the multiplier.
	// 'kYear' means that the Snapshot copy occurs yearly at the frequency
	// set in the multiplier.
	// Enum: ["kEvery","kHour","kDay","kWeek","kMonth","kYear"]
	Periodicity *string `json:"periodicity,omitempty"`

	// Run timeouts after which a run will get cancelled.
	RunTimeouts []*CancellationTimeoutParams `json:"runTimeouts"`

	// Specifies a the source cluster id from which the data must be replicated.
	SourceClusterID *int64 `json:"sourceClusterId,omitempty"`

	// Specifies the replication target to copy the Snapshots to.
	Target struct {
		ReplicationTargetSettings
	} `json:"target,omitempty"`
}

// Validate validates this snapshot replication copy policy
func (m *SnapshotReplicationCopyPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRunType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalockConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodicity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunTimeouts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var snapshotReplicationCopyPolicyTypeBackupRunTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRegular","kFull","kLog","kSystem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotReplicationCopyPolicyTypeBackupRunTypePropEnum = append(snapshotReplicationCopyPolicyTypeBackupRunTypePropEnum, v)
	}
}

const (

	// SnapshotReplicationCopyPolicyBackupRunTypeKRegular captures enum value "kRegular"
	SnapshotReplicationCopyPolicyBackupRunTypeKRegular string = "kRegular"

	// SnapshotReplicationCopyPolicyBackupRunTypeKFull captures enum value "kFull"
	SnapshotReplicationCopyPolicyBackupRunTypeKFull string = "kFull"

	// SnapshotReplicationCopyPolicyBackupRunTypeKLog captures enum value "kLog"
	SnapshotReplicationCopyPolicyBackupRunTypeKLog string = "kLog"

	// SnapshotReplicationCopyPolicyBackupRunTypeKSystem captures enum value "kSystem"
	SnapshotReplicationCopyPolicyBackupRunTypeKSystem string = "kSystem"
)

// prop value enum
func (m *SnapshotReplicationCopyPolicy) validateBackupRunTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotReplicationCopyPolicyTypeBackupRunTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapshotReplicationCopyPolicy) validateBackupRunType(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRunType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupRunTypeEnum("backupRunType", "body", *m.BackupRunType); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) validateCloudTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudTarget) { // not required
		return nil
	}

	if m.CloudTarget != nil {
		if err := m.CloudTarget.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudTarget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudTarget")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) validateDatalockConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DatalockConfig) { // not required
		return nil
	}

	if m.DatalockConfig != nil {
		if err := m.DatalockConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalockConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datalockConfig")
			}
			return err
		}
	}

	return nil
}

var snapshotReplicationCopyPolicyTypePeriodicityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kEvery","kHour","kDay","kWeek","kMonth","kYear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotReplicationCopyPolicyTypePeriodicityPropEnum = append(snapshotReplicationCopyPolicyTypePeriodicityPropEnum, v)
	}
}

const (

	// SnapshotReplicationCopyPolicyPeriodicityKEvery captures enum value "kEvery"
	SnapshotReplicationCopyPolicyPeriodicityKEvery string = "kEvery"

	// SnapshotReplicationCopyPolicyPeriodicityKHour captures enum value "kHour"
	SnapshotReplicationCopyPolicyPeriodicityKHour string = "kHour"

	// SnapshotReplicationCopyPolicyPeriodicityKDay captures enum value "kDay"
	SnapshotReplicationCopyPolicyPeriodicityKDay string = "kDay"

	// SnapshotReplicationCopyPolicyPeriodicityKWeek captures enum value "kWeek"
	SnapshotReplicationCopyPolicyPeriodicityKWeek string = "kWeek"

	// SnapshotReplicationCopyPolicyPeriodicityKMonth captures enum value "kMonth"
	SnapshotReplicationCopyPolicyPeriodicityKMonth string = "kMonth"

	// SnapshotReplicationCopyPolicyPeriodicityKYear captures enum value "kYear"
	SnapshotReplicationCopyPolicyPeriodicityKYear string = "kYear"
)

// prop value enum
func (m *SnapshotReplicationCopyPolicy) validatePeriodicityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotReplicationCopyPolicyTypePeriodicityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapshotReplicationCopyPolicy) validatePeriodicity(formats strfmt.Registry) error {
	if swag.IsZero(m.Periodicity) { // not required
		return nil
	}

	// value enum
	if err := m.validatePeriodicityEnum("periodicity", "body", *m.Periodicity); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) validateRunTimeouts(formats strfmt.Registry) error {
	if swag.IsZero(m.RunTimeouts) { // not required
		return nil
	}

	for i := 0; i < len(m.RunTimeouts); i++ {
		if swag.IsZero(m.RunTimeouts[i]) { // not required
			continue
		}

		if m.RunTimeouts[i] != nil {
			if err := m.RunTimeouts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) validateTarget(formats strfmt.Registry) error {
	if swag.IsZero(m.Target) { // not required
		return nil
	}

	return nil
}

// ContextValidate validate this snapshot replication copy policy based on the context it is used
func (m *SnapshotReplicationCopyPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCloudTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatalockConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRunTimeouts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTarget(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotReplicationCopyPolicy) contextValidateCloudTarget(ctx context.Context, formats strfmt.Registry) error {

	if m.CloudTarget != nil {

		if swag.IsZero(m.CloudTarget) { // not required
			return nil
		}

		if err := m.CloudTarget.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudTarget")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloudTarget")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) contextValidateDatalockConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DatalockConfig != nil {

		if swag.IsZero(m.DatalockConfig) { // not required
			return nil
		}

		if err := m.DatalockConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalockConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datalockConfig")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) contextValidateRunTimeouts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RunTimeouts); i++ {

		if m.RunTimeouts[i] != nil {

			if swag.IsZero(m.RunTimeouts[i]) { // not required
				return nil
			}

			if err := m.RunTimeouts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("runTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SnapshotReplicationCopyPolicy) contextValidateTarget(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotReplicationCopyPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotReplicationCopyPolicy) UnmarshalBinary(b []byte) error {
	var res SnapshotReplicationCopyPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

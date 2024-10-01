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

// ProtectionPolicyRequest Protection Policy Request.
//
// Specifies information about a Protection Policy.
//
// swagger:model ProtectionPolicyRequest
type ProtectionPolicyRequest struct {

	// Backup Run timeouts after which a run will get cancelled.
	BackupRunTimeouts []*CancellationTimeoutParams `json:"backupRunTimeouts"`

	// Array of QuietTime Periods.
	//
	// If specified, this field defines black periods when new Job Runs
	// are not started. If a Job Run has been scheduled but not yet
	// executed and the QuietTime period starts, the behavior depends on
	// the policy field AbortInBlackoutPeriod.
	BlackoutPeriods []*BlackoutPeriod `json:"blackoutPeriods"`

	// CDP Job Policy.
	//
	// Specifies the CDP backup schedule of a Protection Job and how long
	// log files captured by this schedule are retained on the
	// Cohesity Cluster.
	CdpSchedulingPolicy *SchedulingPolicy `json:"cdpSchedulingPolicy,omitempty"`

	// Array of Cloud Deploy Policies.
	//
	// Specifies settings for copying Snapshots to Cloud.
	// CloudDeploy target where backup snapshots may be converted and stored.
	// It also defines the retention of copied Snapshots on the Cloud.
	CloudDeployPolicies []*SnapshotCloudCopyPolicy `json:"cloudDeployPolicies"`

	// Specifies WORM retention type for the incremental/full snapshots. When a
	// WORM retention type is specified, the snapshots of the Protection Groups
	// using this policy will be kept for the last N days as specified in the
	// duration of the datalock. During that time, the snapshots cannot be
	// deleted.
	DatalockConfig *DataLockConfig `json:"datalockConfig,omitempty"`

	// Specifies WORM retention type for the log snapshots. When a WORM retention
	// type is specified, the snapshots of the Protection Groups using this
	// policy will be kept for the last N days as specified in the duration of
	// the datalock. During that time, the snapshots cannot be deleted.
	DatalockConfigLog *DataLockConfig `json:"datalockConfigLog,omitempty"`

	// Specifies WORM retention type for the BMR snapshots. When a WORM retention
	// type is specified, the snapshots of the Protection Groups using this
	// policy will be kept for the last N days as specified in the duration of
	// the datalock. During that time, the snapshots cannot be deleted.
	DatalockConfigSystem *DataLockConfig `json:"datalockConfigSystem,omitempty"`

	// Specifies how many days to retain Snapshots on the Cohesity Cluster.
	DaysToKeep *int64 `json:"daysToKeep,omitempty"`

	// Specifies the number of days to retain log run if Log Schedule exists.
	DaysToKeepLog *int64 `json:"daysToKeepLog,omitempty"`

	// Specifies the number of days to retain system backups made for bare metal
	// recovery. This field is applicable if systemSchedulingPolicy is specified.
	DaysToKeepSystem *int64 `json:"daysToKeepSystem,omitempty"`

	// Description of the Protection Policy.
	Description *string `json:"description,omitempty"`

	// Specifies additional retention policies that should be applied to the
	// backup snapshots. A backup snapshot will be retained up to a time that is
	// the maximum of all retention policies that are applicable to it.
	ExtendedRetentionPolicies []*ExtendedRetentionPolicy `json:"extendedRetentionPolicies"`

	// full scheduling policy
	FullSchedulingPolicy *ProtectionPolicyRequestFullSchedulingPolicy `json:"fullSchedulingPolicy,omitempty"`

	// incremental scheduling policy
	IncrementalSchedulingPolicy *ProtectionPolicyRequestIncrementalSchedulingPolicy `json:"incrementalSchedulingPolicy,omitempty"`

	// Log Job Policy.
	//
	// Specifies the Log backup schedule of a Protection Job and how long
	// log files captured by this schedule are retained on the
	// Cohesity Cluster.
	LogSchedulingPolicy *SchedulingPolicy `json:"logSchedulingPolicy,omitempty"`

	// Specifies the name of the Protection Policy.
	Name *string `json:"name,omitempty"`

	// Species the number of policies linked to a global policy.
	NumLinkedPolicies *int64 `json:"numLinkedPolicies,omitempty"`

	// Specifies the number of mins/hours/days in seconds to retain
	// CDP backups if CDP schedule exists.
	NumSecsToKeep *int32 `json:"numSecsToKeep,omitempty"`

	// Specifies the parent global policy Id. This must be specified when
	// creating a policy from global policy template.
	ParentPolicyID *string `json:"parentPolicyId,omitempty"`

	// Specifies the number of times to retry capturing Snapshots before
	// the Job Run fails.
	Retries *int32 `json:"retries,omitempty"`

	// Specifies the number of minutes before retrying a failed Protection Job.
	RetryIntervalMins *int32 `json:"retryIntervalMins,omitempty"`

	// Specifies the RPO Policy related settings.
	RpoPolicySettings *RpoPolicySettings `json:"rpoPolicySettings,omitempty"`

	// Specifies the period of time before skipping the execution of new
	// Job Runs if an existing queued Job Run of the same Protection Job
	// has not started. For example if this field is set to 30 minutes and
	// a Job Run is scheduled to start at 5:00 AM every day but does not start
	// due to conflicts (such as too many Jobs are running). If the new Job Run
	// does not start by 5:30AM, the Cohesity Cluster will skip the new Job Run.
	// If the original Job Run completes before 5:30AM the next day, a new
	// Job Run is created and starts executing.
	// This field is optional.
	SkipIntervalMins *int32 `json:"skipIntervalMins,omitempty"`

	// Array of External Targets.
	//
	// Specifies settings for copying Snapshots to  Archival External Targets
	// (such as AWS or Tape).
	// It also defines the retention of copied Snapshots on an External Targets
	// such as AWS and Tape.
	SnapshotArchivalCopyPolicies []*SnapshotArchivalCopyPolicy `json:"snapshotArchivalCopyPolicies"`

	// Array of Remote Clusters.
	//
	// Specifies settings for copying Snapshots to Remote Clusters.
	// It also defines the retention of copied Snapshots on a Remote Cluster.
	SnapshotReplicationCopyPolicies []*SnapshotReplicationCopyPolicy `json:"snapshotReplicationCopyPolicies"`

	// Storage Array Snapshot Job Policy.
	//
	// Specifies the storage array snapshot backup schedule of a Protection
	// Job and how long snapshot captured by this schedule are retained on the
	// Storage Array.
	StorageArraySnapshotSchedulingPolicy *SchedulingPolicy `json:"storageArraySnapshotSchedulingPolicy,omitempty"`

	// Agent driven System Job Policy.
	//
	// Specifies the system backup schedule for agents running on servers to run
	// low frequency backup jobs. Images created as part of the backup can be
	// used to perform a "bare metal" recovery.
	SystemSchedulingPolicy *SchedulingPolicy `json:"systemSchedulingPolicy,omitempty"`

	// Specifies the type of the protection policy.
	// 'kRegular' means a regular Protection Policy.
	// 'kRPO' means an RPO Protection Policy.
	// Enum: ["kRegular","kRPO"]
	Type *string `json:"type,omitempty"`

	// Specifies WORM retention type for the snapshots. When a WORM retention
	// type is specified, the snapshots of the Protection Jobs using this policy
	// will be kept until the maximum of the snapshot retention time. During
	// that time, the snapshots cannot be deleted.
	// This field is deprecated. Use DataLockConfig for incremental runs,
	// DataLockConfigLog for log runs, DataLockConfigSystem for BMR runs, and
	// DataLockConfig in extended retention and for copy targets config.
	// deprecated: true
	// 'kNone' implies there is no WORM retention set.
	// 'kCompliance' implies WORM retention is set for compliance reason.
	// 'kAdministrative' implies WORM retention is set for administrative purposes.
	// Enum: ["kNone","kCompliance","kAdministrative"]
	WormRetentionType *string `json:"wormRetentionType,omitempty"`
}

// Validate validates this protection policy request
func (m *ProtectionPolicyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRunTimeouts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlackoutPeriods(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCdpSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudDeployPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalockConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalockConfigLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalockConfigSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtendedRetentionPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIncrementalSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRpoPolicySettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotArchivalCopyPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotReplicationCopyPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArraySnapshotSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemSchedulingPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWormRetentionType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionPolicyRequest) validateBackupRunTimeouts(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRunTimeouts) { // not required
		return nil
	}

	for i := 0; i < len(m.BackupRunTimeouts); i++ {
		if swag.IsZero(m.BackupRunTimeouts[i]) { // not required
			continue
		}

		if m.BackupRunTimeouts[i] != nil {
			if err := m.BackupRunTimeouts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupRunTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupRunTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) validateBlackoutPeriods(formats strfmt.Registry) error {
	if swag.IsZero(m.BlackoutPeriods) { // not required
		return nil
	}

	for i := 0; i < len(m.BlackoutPeriods); i++ {
		if swag.IsZero(m.BlackoutPeriods[i]) { // not required
			continue
		}

		if m.BlackoutPeriods[i] != nil {
			if err := m.BlackoutPeriods[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blackoutPeriods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blackoutPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) validateCdpSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.CdpSchedulingPolicy) { // not required
		return nil
	}

	if m.CdpSchedulingPolicy != nil {
		if err := m.CdpSchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdpSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdpSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateCloudDeployPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.CloudDeployPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudDeployPolicies); i++ {
		if swag.IsZero(m.CloudDeployPolicies[i]) { // not required
			continue
		}

		if m.CloudDeployPolicies[i] != nil {
			if err := m.CloudDeployPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudDeployPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cloudDeployPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) validateDatalockConfig(formats strfmt.Registry) error {
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

func (m *ProtectionPolicyRequest) validateDatalockConfigLog(formats strfmt.Registry) error {
	if swag.IsZero(m.DatalockConfigLog) { // not required
		return nil
	}

	if m.DatalockConfigLog != nil {
		if err := m.DatalockConfigLog.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalockConfigLog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datalockConfigLog")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateDatalockConfigSystem(formats strfmt.Registry) error {
	if swag.IsZero(m.DatalockConfigSystem) { // not required
		return nil
	}

	if m.DatalockConfigSystem != nil {
		if err := m.DatalockConfigSystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalockConfigSystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datalockConfigSystem")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateExtendedRetentionPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtendedRetentionPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.ExtendedRetentionPolicies); i++ {
		if swag.IsZero(m.ExtendedRetentionPolicies[i]) { // not required
			continue
		}

		if m.ExtendedRetentionPolicies[i] != nil {
			if err := m.ExtendedRetentionPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extendedRetentionPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extendedRetentionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) validateFullSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.FullSchedulingPolicy) { // not required
		return nil
	}

	if m.FullSchedulingPolicy != nil {
		if err := m.FullSchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateIncrementalSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IncrementalSchedulingPolicy) { // not required
		return nil
	}

	if m.IncrementalSchedulingPolicy != nil {
		if err := m.IncrementalSchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incrementalSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incrementalSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateLogSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.LogSchedulingPolicy) { // not required
		return nil
	}

	if m.LogSchedulingPolicy != nil {
		if err := m.LogSchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateRpoPolicySettings(formats strfmt.Registry) error {
	if swag.IsZero(m.RpoPolicySettings) { // not required
		return nil
	}

	if m.RpoPolicySettings != nil {
		if err := m.RpoPolicySettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpoPolicySettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpoPolicySettings")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateSnapshotArchivalCopyPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotArchivalCopyPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotArchivalCopyPolicies); i++ {
		if swag.IsZero(m.SnapshotArchivalCopyPolicies[i]) { // not required
			continue
		}

		if m.SnapshotArchivalCopyPolicies[i] != nil {
			if err := m.SnapshotArchivalCopyPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotArchivalCopyPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshotArchivalCopyPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) validateSnapshotReplicationCopyPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotReplicationCopyPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.SnapshotReplicationCopyPolicies); i++ {
		if swag.IsZero(m.SnapshotReplicationCopyPolicies[i]) { // not required
			continue
		}

		if m.SnapshotReplicationCopyPolicies[i] != nil {
			if err := m.SnapshotReplicationCopyPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotReplicationCopyPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshotReplicationCopyPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) validateStorageArraySnapshotSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageArraySnapshotSchedulingPolicy) { // not required
		return nil
	}

	if m.StorageArraySnapshotSchedulingPolicy != nil {
		if err := m.StorageArraySnapshotSchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) validateSystemSchedulingPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.SystemSchedulingPolicy) { // not required
		return nil
	}

	if m.SystemSchedulingPolicy != nil {
		if err := m.SystemSchedulingPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

var protectionPolicyRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRegular","kRPO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectionPolicyRequestTypeTypePropEnum = append(protectionPolicyRequestTypeTypePropEnum, v)
	}
}

const (

	// ProtectionPolicyRequestTypeKRegular captures enum value "kRegular"
	ProtectionPolicyRequestTypeKRegular string = "kRegular"

	// ProtectionPolicyRequestTypeKRPO captures enum value "kRPO"
	ProtectionPolicyRequestTypeKRPO string = "kRPO"
)

// prop value enum
func (m *ProtectionPolicyRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectionPolicyRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectionPolicyRequest) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var protectionPolicyRequestTypeWormRetentionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNone","kCompliance","kAdministrative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protectionPolicyRequestTypeWormRetentionTypePropEnum = append(protectionPolicyRequestTypeWormRetentionTypePropEnum, v)
	}
}

const (

	// ProtectionPolicyRequestWormRetentionTypeKNone captures enum value "kNone"
	ProtectionPolicyRequestWormRetentionTypeKNone string = "kNone"

	// ProtectionPolicyRequestWormRetentionTypeKCompliance captures enum value "kCompliance"
	ProtectionPolicyRequestWormRetentionTypeKCompliance string = "kCompliance"

	// ProtectionPolicyRequestWormRetentionTypeKAdministrative captures enum value "kAdministrative"
	ProtectionPolicyRequestWormRetentionTypeKAdministrative string = "kAdministrative"
)

// prop value enum
func (m *ProtectionPolicyRequest) validateWormRetentionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, protectionPolicyRequestTypeWormRetentionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ProtectionPolicyRequest) validateWormRetentionType(formats strfmt.Registry) error {
	if swag.IsZero(m.WormRetentionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateWormRetentionTypeEnum("wormRetentionType", "body", *m.WormRetentionType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this protection policy request based on the context it is used
func (m *ProtectionPolicyRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBackupRunTimeouts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlackoutPeriods(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCdpSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCloudDeployPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatalockConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatalockConfigLog(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatalockConfigSystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExtendedRetentionPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIncrementalSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRpoPolicySettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotArchivalCopyPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotReplicationCopyPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageArraySnapshotSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSystemSchedulingPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProtectionPolicyRequest) contextValidateBackupRunTimeouts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BackupRunTimeouts); i++ {

		if m.BackupRunTimeouts[i] != nil {

			if swag.IsZero(m.BackupRunTimeouts[i]) { // not required
				return nil
			}

			if err := m.BackupRunTimeouts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("backupRunTimeouts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("backupRunTimeouts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateBlackoutPeriods(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BlackoutPeriods); i++ {

		if m.BlackoutPeriods[i] != nil {

			if swag.IsZero(m.BlackoutPeriods[i]) { // not required
				return nil
			}

			if err := m.BlackoutPeriods[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blackoutPeriods" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("blackoutPeriods" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateCdpSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.CdpSchedulingPolicy != nil {

		if swag.IsZero(m.CdpSchedulingPolicy) { // not required
			return nil
		}

		if err := m.CdpSchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdpSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdpSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateCloudDeployPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CloudDeployPolicies); i++ {

		if m.CloudDeployPolicies[i] != nil {

			if swag.IsZero(m.CloudDeployPolicies[i]) { // not required
				return nil
			}

			if err := m.CloudDeployPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudDeployPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cloudDeployPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateDatalockConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProtectionPolicyRequest) contextValidateDatalockConfigLog(ctx context.Context, formats strfmt.Registry) error {

	if m.DatalockConfigLog != nil {

		if swag.IsZero(m.DatalockConfigLog) { // not required
			return nil
		}

		if err := m.DatalockConfigLog.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalockConfigLog")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datalockConfigLog")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateDatalockConfigSystem(ctx context.Context, formats strfmt.Registry) error {

	if m.DatalockConfigSystem != nil {

		if swag.IsZero(m.DatalockConfigSystem) { // not required
			return nil
		}

		if err := m.DatalockConfigSystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datalockConfigSystem")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datalockConfigSystem")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateExtendedRetentionPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExtendedRetentionPolicies); i++ {

		if m.ExtendedRetentionPolicies[i] != nil {

			if swag.IsZero(m.ExtendedRetentionPolicies[i]) { // not required
				return nil
			}

			if err := m.ExtendedRetentionPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extendedRetentionPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("extendedRetentionPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateFullSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.FullSchedulingPolicy != nil {

		if swag.IsZero(m.FullSchedulingPolicy) { // not required
			return nil
		}

		if err := m.FullSchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fullSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fullSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateIncrementalSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IncrementalSchedulingPolicy != nil {

		if swag.IsZero(m.IncrementalSchedulingPolicy) { // not required
			return nil
		}

		if err := m.IncrementalSchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("incrementalSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("incrementalSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateLogSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.LogSchedulingPolicy != nil {

		if swag.IsZero(m.LogSchedulingPolicy) { // not required
			return nil
		}

		if err := m.LogSchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateRpoPolicySettings(ctx context.Context, formats strfmt.Registry) error {

	if m.RpoPolicySettings != nil {

		if swag.IsZero(m.RpoPolicySettings) { // not required
			return nil
		}

		if err := m.RpoPolicySettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rpoPolicySettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rpoPolicySettings")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateSnapshotArchivalCopyPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapshotArchivalCopyPolicies); i++ {

		if m.SnapshotArchivalCopyPolicies[i] != nil {

			if swag.IsZero(m.SnapshotArchivalCopyPolicies[i]) { // not required
				return nil
			}

			if err := m.SnapshotArchivalCopyPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotArchivalCopyPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshotArchivalCopyPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateSnapshotReplicationCopyPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SnapshotReplicationCopyPolicies); i++ {

		if m.SnapshotReplicationCopyPolicies[i] != nil {

			if swag.IsZero(m.SnapshotReplicationCopyPolicies[i]) { // not required
				return nil
			}

			if err := m.SnapshotReplicationCopyPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snapshotReplicationCopyPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("snapshotReplicationCopyPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateStorageArraySnapshotSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageArraySnapshotSchedulingPolicy != nil {

		if swag.IsZero(m.StorageArraySnapshotSchedulingPolicy) { // not required
			return nil
		}

		if err := m.StorageArraySnapshotSchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *ProtectionPolicyRequest) contextValidateSystemSchedulingPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.SystemSchedulingPolicy != nil {

		if swag.IsZero(m.SystemSchedulingPolicy) { // not required
			return nil
		}

		if err := m.SystemSchedulingPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("systemSchedulingPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("systemSchedulingPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionPolicyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionPolicyRequest) UnmarshalBinary(b []byte) error {
	var res ProtectionPolicyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

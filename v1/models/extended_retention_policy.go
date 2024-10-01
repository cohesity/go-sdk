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

// ExtendedRetentionPolicy Extended Retention Policy.
//
// Specifies additional retention policies to apply to backup snapshots.
//
// swagger:model ExtendedRetentionPolicy
type ExtendedRetentionPolicy struct {

	// Specified the Id for a snapshot copy policy. This is generated when the
	// policy is created.
	ID *string `json:"Id,omitempty"`

	// The backup run type to which this extended retention applies to. If this is
	// not set, the extended retention will be applicable to all non-log backup
	// types. Currently, the only value that can be set here is kFull.
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

	// Specifies WORM retention type for snapshots under extended retention.
	DatalockConfig *DataLockConfig `json:"datalockConfig,omitempty"`

	// Specifies the number of days to retain copied Snapshots on the target.
	DaysToKeep *int64 `json:"daysToKeep,omitempty"`

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
}

// Validate validates this extended retention policy
func (m *ExtendedRetentionPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackupRunType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalockConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodicity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var extendedRetentionPolicyTypeBackupRunTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRegular","kFull","kLog","kSystem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extendedRetentionPolicyTypeBackupRunTypePropEnum = append(extendedRetentionPolicyTypeBackupRunTypePropEnum, v)
	}
}

const (

	// ExtendedRetentionPolicyBackupRunTypeKRegular captures enum value "kRegular"
	ExtendedRetentionPolicyBackupRunTypeKRegular string = "kRegular"

	// ExtendedRetentionPolicyBackupRunTypeKFull captures enum value "kFull"
	ExtendedRetentionPolicyBackupRunTypeKFull string = "kFull"

	// ExtendedRetentionPolicyBackupRunTypeKLog captures enum value "kLog"
	ExtendedRetentionPolicyBackupRunTypeKLog string = "kLog"

	// ExtendedRetentionPolicyBackupRunTypeKSystem captures enum value "kSystem"
	ExtendedRetentionPolicyBackupRunTypeKSystem string = "kSystem"
)

// prop value enum
func (m *ExtendedRetentionPolicy) validateBackupRunTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, extendedRetentionPolicyTypeBackupRunTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExtendedRetentionPolicy) validateBackupRunType(formats strfmt.Registry) error {
	if swag.IsZero(m.BackupRunType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBackupRunTypeEnum("backupRunType", "body", *m.BackupRunType); err != nil {
		return err
	}

	return nil
}

func (m *ExtendedRetentionPolicy) validateDatalockConfig(formats strfmt.Registry) error {
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

var extendedRetentionPolicyTypePeriodicityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kEvery","kHour","kDay","kWeek","kMonth","kYear"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extendedRetentionPolicyTypePeriodicityPropEnum = append(extendedRetentionPolicyTypePeriodicityPropEnum, v)
	}
}

const (

	// ExtendedRetentionPolicyPeriodicityKEvery captures enum value "kEvery"
	ExtendedRetentionPolicyPeriodicityKEvery string = "kEvery"

	// ExtendedRetentionPolicyPeriodicityKHour captures enum value "kHour"
	ExtendedRetentionPolicyPeriodicityKHour string = "kHour"

	// ExtendedRetentionPolicyPeriodicityKDay captures enum value "kDay"
	ExtendedRetentionPolicyPeriodicityKDay string = "kDay"

	// ExtendedRetentionPolicyPeriodicityKWeek captures enum value "kWeek"
	ExtendedRetentionPolicyPeriodicityKWeek string = "kWeek"

	// ExtendedRetentionPolicyPeriodicityKMonth captures enum value "kMonth"
	ExtendedRetentionPolicyPeriodicityKMonth string = "kMonth"

	// ExtendedRetentionPolicyPeriodicityKYear captures enum value "kYear"
	ExtendedRetentionPolicyPeriodicityKYear string = "kYear"
)

// prop value enum
func (m *ExtendedRetentionPolicy) validatePeriodicityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, extendedRetentionPolicyTypePeriodicityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExtendedRetentionPolicy) validatePeriodicity(formats strfmt.Registry) error {
	if swag.IsZero(m.Periodicity) { // not required
		return nil
	}

	// value enum
	if err := m.validatePeriodicityEnum("periodicity", "body", *m.Periodicity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this extended retention policy based on the context it is used
func (m *ExtendedRetentionPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatalockConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendedRetentionPolicy) contextValidateDatalockConfig(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *ExtendedRetentionPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExtendedRetentionPolicy) UnmarshalBinary(b []byte) error {
	var res ExtendedRetentionPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

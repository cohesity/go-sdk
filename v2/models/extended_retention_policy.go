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

	// Specifies the schedule settings for extended retention.
	// Required: true
	Schedule *ExtendedRetentionSchedule `json:"schedule"`

	// Specifies the Retention period of a backup in days, months or years.
	// Required: true
	Retention *Retention `json:"retention"`

	// The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.
	// 'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.
	// 'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.
	// 'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.
	// 'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
	// Enum: ["Regular","Full","Log","System"]
	RunType *string `json:"runType,omitempty"`

	// Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
	ConfigID *string `json:"configId,omitempty"`
}

// Validate validates this extended retention policy
func (m *ExtendedRetentionPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendedRetentionPolicy) validateSchedule(formats strfmt.Registry) error {

	if err := validate.Required("schedule", "body", m.Schedule); err != nil {
		return err
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedRetentionPolicy) validateRetention(formats strfmt.Registry) error {

	if err := validate.Required("retention", "body", m.Retention); err != nil {
		return err
	}

	if m.Retention != nil {
		if err := m.Retention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
			}
			return err
		}
	}

	return nil
}

var extendedRetentionPolicyTypeRunTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Regular","Full","Log","System"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extendedRetentionPolicyTypeRunTypePropEnum = append(extendedRetentionPolicyTypeRunTypePropEnum, v)
	}
}

const (

	// ExtendedRetentionPolicyRunTypeRegular captures enum value "Regular"
	ExtendedRetentionPolicyRunTypeRegular string = "Regular"

	// ExtendedRetentionPolicyRunTypeFull captures enum value "Full"
	ExtendedRetentionPolicyRunTypeFull string = "Full"

	// ExtendedRetentionPolicyRunTypeLog captures enum value "Log"
	ExtendedRetentionPolicyRunTypeLog string = "Log"

	// ExtendedRetentionPolicyRunTypeSystem captures enum value "System"
	ExtendedRetentionPolicyRunTypeSystem string = "System"
)

// prop value enum
func (m *ExtendedRetentionPolicy) validateRunTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, extendedRetentionPolicyTypeRunTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExtendedRetentionPolicy) validateRunType(formats strfmt.Registry) error {
	if swag.IsZero(m.RunType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRunTypeEnum("runType", "body", *m.RunType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this extended retention policy based on the context it is used
func (m *ExtendedRetentionPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExtendedRetentionPolicy) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {

		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *ExtendedRetentionPolicy) contextValidateRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.Retention != nil {

		if err := m.Retention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("retention")
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

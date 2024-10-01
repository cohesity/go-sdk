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

// FileLevelDataLockConfig File Level Data Lock Configurations.
//
// Specifies a config to lock files in a view - to protect from malicious or
// an accidental attempt to delete or modify the files in this view.
//
// swagger:model FileLevelDataLockConfig
type FileLevelDataLockConfig struct {

	// Specifies the duration to lock a file that has not been accessed or
	// modified (ie. has been idle) for a certain duration of time in
	// milliseconds. Do not set if it is required to disable auto lock.
	AutoLockAfterDurationIdle *uint64 `json:"autoLockAfterDurationIdle,omitempty"`

	// Specifies a global default retention duration for files in this view, if
	// file lock is enabled for this view. Also, it is a required field if file
	// lock is enabled. Set to -1 if the required default retention period is
	// forever.
	DefaultFileRetentionDurationMsecs *int64 `json:"defaultFileRetentionDurationMsecs,omitempty"`

	// Specifies a definite timestamp in milliseconds for retaining the file.
	ExpiryTimestampMsecs *uint64 `json:"expiryTimestampMsecs,omitempty"`

	// Specifies the supported mechanisms to explicity lock a file from NFS/SMB
	// interface. Supported locking protocols: kSetReadOnly, kSetAtime.
	// 'kSetReadOnly' is compatible with Isilon/Netapp behaviour. This locks the
	// file and the retention duration is determined in this order:
	// 1) atime, if set by user/application and within min and max retention
	// duration.
	// 2) Min retention duration, if set.
	// 3) Otherwise, file is switched to expired data automatically.
	// 'kSetAtime' is compatible with Data Domain behaviour.
	// Enum: ["kSetReadOnly","kSetAtime"]
	LockingProtocol *string `json:"lockingProtocol,omitempty"`

	// Specifies a maximum duration in milliseconds for which any file in this
	// view can be retained for. Set to -1 if the required retention duration is
	// forever. If set, it should be greater than or equal to the default
	// retention period as well as the min retention period.
	MaxRetentionDurationMsecs *int64 `json:"maxRetentionDurationMsecs,omitempty"`

	// Specifies a minimum retention duration in milliseconds after a file gets
	// locked. The file cannot be modified or deleted during this timeframe. Set
	// to -1 if the required retention duration is forever. This should be set
	// less than or equal to the default retention duration.
	MinRetentionDurationMsecs *int64 `json:"minRetentionDurationMsecs,omitempty"`

	// Specifies the mode of file level datalock.
	// Enterprise mode can be upgraded to Compliance mode, but Compliance mode
	// cannot be downgraded to Enterprise mode, unless view's
	// FileLevelDataLockConfig has coexisting_lock_mode set.
	// kCompliance: This mode would disallow all user to delete/modify file or
	// view under any condition when it 's in locked status except for deleting
	// view when the view is empty.
	// kEnterprise: This mode would follow the rules as compliance mode for
	// normal users. But it would allow the storage admin
	// (1) to delete view or file anytime no matter
	// it is in locked status or expired.
	// (2) to rename the view
	// (3) to bring back the retention period when it's in locked mode
	// A lock mode of a file in a view can be in one of the following:
	//
	// 'kCompliance': Default mode of datalock, in this mode, Data Security Admin
	// cannot modify/delete this view when datalock is in effect. Data Security
	// Admin can delete this view when datalock is expired.
	// 'kEnterprise' : In this mode, Data Security Admin can change view name or
	// delete view when datalock is in effect. Datalock in this mode can be
	// upgraded to 'kCompliance' mode.
	// Enum: ["kCompliance","kEnterprise"]
	Mode *string `json:"mode,omitempty"`
}

// Validate validates this file level data lock config
func (m *FileLevelDataLockConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLockingProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var fileLevelDataLockConfigTypeLockingProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kSetReadOnly","kSetAtime"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileLevelDataLockConfigTypeLockingProtocolPropEnum = append(fileLevelDataLockConfigTypeLockingProtocolPropEnum, v)
	}
}

const (

	// FileLevelDataLockConfigLockingProtocolKSetReadOnly captures enum value "kSetReadOnly"
	FileLevelDataLockConfigLockingProtocolKSetReadOnly string = "kSetReadOnly"

	// FileLevelDataLockConfigLockingProtocolKSetAtime captures enum value "kSetAtime"
	FileLevelDataLockConfigLockingProtocolKSetAtime string = "kSetAtime"
)

// prop value enum
func (m *FileLevelDataLockConfig) validateLockingProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileLevelDataLockConfigTypeLockingProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileLevelDataLockConfig) validateLockingProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.LockingProtocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateLockingProtocolEnum("lockingProtocol", "body", *m.LockingProtocol); err != nil {
		return err
	}

	return nil
}

var fileLevelDataLockConfigTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kCompliance","kEnterprise"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fileLevelDataLockConfigTypeModePropEnum = append(fileLevelDataLockConfigTypeModePropEnum, v)
	}
}

const (

	// FileLevelDataLockConfigModeKCompliance captures enum value "kCompliance"
	FileLevelDataLockConfigModeKCompliance string = "kCompliance"

	// FileLevelDataLockConfigModeKEnterprise captures enum value "kEnterprise"
	FileLevelDataLockConfigModeKEnterprise string = "kEnterprise"
)

// prop value enum
func (m *FileLevelDataLockConfig) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fileLevelDataLockConfigTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FileLevelDataLockConfig) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", *m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this file level data lock config based on context it is used
func (m *FileLevelDataLockConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileLevelDataLockConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileLevelDataLockConfig) UnmarshalBinary(b []byte) error {
	var res FileLevelDataLockConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

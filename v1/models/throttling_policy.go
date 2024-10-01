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

// ThrottlingPolicy Message that specifies the throttling policy for a particular registered
// entity.
//
// swagger:model ThrottlingPolicy
type ThrottlingPolicy struct {

	// This specifies the datastore streams config applied to all datastores
	// that are part of the registered entity. This policy can be overriden per
	// datastore by specifying it in DatastoreThrottlingPolicy below.
	DatastoreStreamsConfig *ThrottlingPolicyDatastoreStreamsConfig `json:"datastoreStreamsConfig,omitempty"`

	// This field can be used to override the throttling policy for individual
	// datastores.
	DatastoreThrottlingPolicies []*ThrottlingPolicyDatastoreThrottlingPolicy `json:"datastoreThrottlingPolicies"`

	// The registered entity for which this throttling policy applies.
	// NOTE: This field is optional and need not be set by Iris.
	Entity *EntityProto `json:"entity,omitempty"`

	// Whether datastore streams can be configured on all datastores that are
	// part of the registered entity. If set to true, then the config within
	// 'DatastoreStreamsConfig' would be applicable to all those datastores.
	IsDatastoreStreamsConfigEnabled *bool `json:"isDatastoreStreamsConfigEnabled,omitempty"`

	// Whether we will use storage snapshot managmement max snap config to all
	// volumes/luns that are part of the registered entity.
	IsMaxSnapshotsConfigEnabled *bool `json:"isMaxSnapshotsConfigEnabled,omitempty"`

	// Whether we will use storage snapshot managmement max space config to all
	// volumes/luns that are part of the registered entity.
	IsMaxSpaceConfigEnabled *bool `json:"isMaxSpaceConfigEnabled,omitempty"`

	// Whether no. of backups can be configured on the registered entity. If set
	// to true, then the config within 'RegisteredSourceThrottlingConfig' would
	// be applicable to the registered entity.
	IsRegisteredSourceThrottlingConfigEnabled *bool `json:"isRegisteredSourceThrottlingConfigEnabled,omitempty"`

	// Whether we will adaptively throttle read operations from the datastores
	// that are part of the registered entity.
	// Note: This is only applicable to latency throttling.
	IsThrottlingEnabled *bool `json:"isThrottlingEnabled,omitempty"`

	// This specifies the thresholds that should be applied to all datastores
	// that are part of the registered entity. The thresholds for a datastore can
	// be overriden by specifying it in datastore_throttling_policy below.
	LatencyThresholds *ThrottlingPolicyLatencyThresholds `json:"latencyThresholds,omitempty"`

	// This specifies the registered source throttling config applied to
	// registered entity.
	RegisteredSourceThrottlingConfig *ThrottlingPolicyRegisteredSourceThrottlingConfig `json:"registeredSourceThrottlingConfig,omitempty"`

	// This specifies the storage snapshot managmement max snap config applied
	// to all volumes/lun that are part of the registered entity. This policy
	// can be overriden per volume/lun  by specifying it in
	// StorageArraySnapshotThrottlingPolicy below.
	StorageArraySnapshotMaxSnapshotConfig *ThrottlingPolicyStorageArraySnapshotMaxSnapshotConfig `json:"storageArraySnapshotMaxSnapshotConfig,omitempty"`

	// This specifies the storage snapshot managmement max space config applied
	// to all volumes/lun that are part of the registered entity. This policy
	// can be overriden per volume/lun  by specifying it in
	// StorageArraySnapshotThrottlingPolicy below.
	StorageArraySnapshotMaxSpaceConfig *ThrottlingPolicyStorageArraySnapshotMaxSpaceConfig `json:"storageArraySnapshotMaxSpaceConfig,omitempty"`

	// This field is used for throttling policy for individual volume/lun.
	StorageArraySnapshotThrottlingPolicies []*ThrottlingPolicyStorageArraySnapshotThrottlingPolicy `json:"storageArraySnapshotThrottlingPolicies"`
}

// Validate validates this throttling policy
func (m *ThrottlingPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatastoreStreamsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatastoreThrottlingPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatencyThresholds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredSourceThrottlingConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArraySnapshotMaxSnapshotConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArraySnapshotMaxSpaceConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageArraySnapshotThrottlingPolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottlingPolicy) validateDatastoreStreamsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DatastoreStreamsConfig) { // not required
		return nil
	}

	if m.DatastoreStreamsConfig != nil {
		if err := m.DatastoreStreamsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreStreamsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreStreamsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) validateDatastoreThrottlingPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.DatastoreThrottlingPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.DatastoreThrottlingPolicies); i++ {
		if swag.IsZero(m.DatastoreThrottlingPolicies[i]) { // not required
			continue
		}

		if m.DatastoreThrottlingPolicies[i] != nil {
			if err := m.DatastoreThrottlingPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datastoreThrottlingPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datastoreThrottlingPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThrottlingPolicy) validateEntity(formats strfmt.Registry) error {
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

func (m *ThrottlingPolicy) validateLatencyThresholds(formats strfmt.Registry) error {
	if swag.IsZero(m.LatencyThresholds) { // not required
		return nil
	}

	if m.LatencyThresholds != nil {
		if err := m.LatencyThresholds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latencyThresholds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latencyThresholds")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) validateRegisteredSourceThrottlingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.RegisteredSourceThrottlingConfig) { // not required
		return nil
	}

	if m.RegisteredSourceThrottlingConfig != nil {
		if err := m.RegisteredSourceThrottlingConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registeredSourceThrottlingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registeredSourceThrottlingConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) validateStorageArraySnapshotMaxSnapshotConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageArraySnapshotMaxSnapshotConfig) { // not required
		return nil
	}

	if m.StorageArraySnapshotMaxSnapshotConfig != nil {
		if err := m.StorageArraySnapshotMaxSnapshotConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotMaxSnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotMaxSnapshotConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) validateStorageArraySnapshotMaxSpaceConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageArraySnapshotMaxSpaceConfig) { // not required
		return nil
	}

	if m.StorageArraySnapshotMaxSpaceConfig != nil {
		if err := m.StorageArraySnapshotMaxSpaceConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotMaxSpaceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotMaxSpaceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) validateStorageArraySnapshotThrottlingPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageArraySnapshotThrottlingPolicies) { // not required
		return nil
	}

	for i := 0; i < len(m.StorageArraySnapshotThrottlingPolicies); i++ {
		if swag.IsZero(m.StorageArraySnapshotThrottlingPolicies[i]) { // not required
			continue
		}

		if m.StorageArraySnapshotThrottlingPolicies[i] != nil {
			if err := m.StorageArraySnapshotThrottlingPolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageArraySnapshotThrottlingPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageArraySnapshotThrottlingPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this throttling policy based on the context it is used
func (m *ThrottlingPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatastoreStreamsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatastoreThrottlingPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatencyThresholds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegisteredSourceThrottlingConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageArraySnapshotMaxSnapshotConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageArraySnapshotMaxSpaceConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageArraySnapshotThrottlingPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ThrottlingPolicy) contextValidateDatastoreStreamsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DatastoreStreamsConfig != nil {

		if swag.IsZero(m.DatastoreStreamsConfig) { // not required
			return nil
		}

		if err := m.DatastoreStreamsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datastoreStreamsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datastoreStreamsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) contextValidateDatastoreThrottlingPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DatastoreThrottlingPolicies); i++ {

		if m.DatastoreThrottlingPolicies[i] != nil {

			if swag.IsZero(m.DatastoreThrottlingPolicies[i]) { // not required
				return nil
			}

			if err := m.DatastoreThrottlingPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("datastoreThrottlingPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("datastoreThrottlingPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ThrottlingPolicy) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ThrottlingPolicy) contextValidateLatencyThresholds(ctx context.Context, formats strfmt.Registry) error {

	if m.LatencyThresholds != nil {

		if swag.IsZero(m.LatencyThresholds) { // not required
			return nil
		}

		if err := m.LatencyThresholds.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latencyThresholds")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latencyThresholds")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) contextValidateRegisteredSourceThrottlingConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.RegisteredSourceThrottlingConfig != nil {

		if swag.IsZero(m.RegisteredSourceThrottlingConfig) { // not required
			return nil
		}

		if err := m.RegisteredSourceThrottlingConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("registeredSourceThrottlingConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("registeredSourceThrottlingConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) contextValidateStorageArraySnapshotMaxSnapshotConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageArraySnapshotMaxSnapshotConfig != nil {

		if swag.IsZero(m.StorageArraySnapshotMaxSnapshotConfig) { // not required
			return nil
		}

		if err := m.StorageArraySnapshotMaxSnapshotConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotMaxSnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotMaxSnapshotConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) contextValidateStorageArraySnapshotMaxSpaceConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageArraySnapshotMaxSpaceConfig != nil {

		if swag.IsZero(m.StorageArraySnapshotMaxSpaceConfig) { // not required
			return nil
		}

		if err := m.StorageArraySnapshotMaxSpaceConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageArraySnapshotMaxSpaceConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageArraySnapshotMaxSpaceConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ThrottlingPolicy) contextValidateStorageArraySnapshotThrottlingPolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StorageArraySnapshotThrottlingPolicies); i++ {

		if m.StorageArraySnapshotThrottlingPolicies[i] != nil {

			if swag.IsZero(m.StorageArraySnapshotThrottlingPolicies[i]) { // not required
				return nil
			}

			if err := m.StorageArraySnapshotThrottlingPolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("storageArraySnapshotThrottlingPolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("storageArraySnapshotThrottlingPolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ThrottlingPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThrottlingPolicy) UnmarshalBinary(b []byte) error {
	var res ThrottlingPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

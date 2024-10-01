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

// ReplicationInfoBase replication info base
//
// swagger:model ReplicationInfoBase
type ReplicationInfoBase struct {

	// Indicates the time remaining for the blackout window to end.
	BlackoutRemainingTimeUsecs *int64 `json:"blackoutRemainingTimeUsecs,omitempty"`

	// Number of bytes sent over the wire for this replication so far. This value
	// could be larger than 'logical_bytes_transferred' due to redundant work
	// performed by Madrox across crashes.
	BytesTransferred *int64 `json:"bytesTransferred,omitempty"`

	// Indicates if nfs, smb permissions should be enabled while cloning from
	// backup view to madrox Tx internal view.
	EnableNfsSmbPermissions *bool `json:"enableNfsSmbPermissions,omitempty"`

	// Time when this replication ended. If not set, then the replication has
	// not ended yet.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// If the replication has completed, the following indicates whether there
	// was an error in its completion.
	Error *PrivateErrorProto `json:"error,omitempty"`

	// Sum of logical size of all the files which have changes that need to be
	// replicated.
	EstimatedLogicalBytesToTransfer *int64 `json:"estimatedLogicalBytesToTransfer,omitempty"`

	// Any existing data lock retention of the local copy of the backup run on
	// the Rx cluster if it was replicated there earlier.
	ExistingDataLockConstraints *DataLockConstraintsProto `json:"existingDataLockConstraints,omitempty"`

	// Any existing expiry time of the backup run on the Rx cluster if it was
	// replicated there earlier.
	ExistingExpiryTimeUsecs *int64 `json:"existingExpiryTimeUsecs,omitempty"`

	// Any existing data lock retention of the local copy of the backup run on
	// the Rx cluster if it was replicated there earlier. This is to hold
	// worm_retention info if Rx is not update with
	// 'datalock_retention_per_target'.
	ExistingWormRetentionOnRx *WormRetentionProto `json:"existingWormRetentionOnRx,omitempty"`

	// A map that contains all the backup tasks that should not be replicated to
	// the Rx cluster (e.g., because they might have been replicated earlier).
	// The key of the task is the id of the backup task, and the value is a bool
	// that indicates whether that backup task should be filtered or not.
	FilteredBackupTaskIDMap []*ReplicationInfoBaseFilteredBackupTaskIDMapEntry `json:"filteredBackupTaskIdMap"`

	// True, if this is a full replication. False, if incremental.
	IsFullReplication *bool `json:"isFullReplication,omitempty"`

	// True, if 'datalock_retention_per_target' is supported on Rx cluster.
	IsTargetDatalockCapable *bool `json:"isTargetDatalockCapable,omitempty"`

	// Number of logical bytes transferred for this replication so far. This
	// value can never exceed the total logical size of the replicated view.
	LogicalBytesTransferred *int64 `json:"logicalBytesTransferred,omitempty"`

	// The total amount of logical data to be transferred for this replication.
	LogicalSizeBytes *int64 `json:"logicalSizeBytes,omitempty"`

	// Number of metadata actions that have been completed for this replication.
	MetadataActionsCompleted *int64 `json:"metadataActionsCompleted,omitempty"`

	// Percentage of replication that has been completed. Note that this may be
	// < 100 if error is set to something other than kNoError. It implies we
	// saw an error after completing this percentage of replication.
	PctCompleted *int32 `json:"pctCompleted,omitempty"`

	// The name of the remote cluster.
	RemoteClusterName *string `json:"remoteClusterName,omitempty"`

	// The incarnation id of the Rx cluster.
	RxClusterIncarnationID *int64 `json:"rxClusterIncarnationId,omitempty"`

	// Time when this replication was started. If not set, then replication
	// has not been started yet.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// The name of the view that is being replicated.
	ViewName *string `json:"viewName,omitempty"`
}

// Validate validates this replication info base
func (m *ReplicationInfoBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingDataLockConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingWormRetentionOnRx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilteredBackupTaskIDMap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationInfoBase) validateError(formats strfmt.Registry) error {
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

func (m *ReplicationInfoBase) validateExistingDataLockConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.ExistingDataLockConstraints) { // not required
		return nil
	}

	if m.ExistingDataLockConstraints != nil {
		if err := m.ExistingDataLockConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingDataLockConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingDataLockConstraints")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationInfoBase) validateExistingWormRetentionOnRx(formats strfmt.Registry) error {
	if swag.IsZero(m.ExistingWormRetentionOnRx) { // not required
		return nil
	}

	if m.ExistingWormRetentionOnRx != nil {
		if err := m.ExistingWormRetentionOnRx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingWormRetentionOnRx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingWormRetentionOnRx")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationInfoBase) validateFilteredBackupTaskIDMap(formats strfmt.Registry) error {
	if swag.IsZero(m.FilteredBackupTaskIDMap) { // not required
		return nil
	}

	for i := 0; i < len(m.FilteredBackupTaskIDMap); i++ {
		if swag.IsZero(m.FilteredBackupTaskIDMap[i]) { // not required
			continue
		}

		if m.FilteredBackupTaskIDMap[i] != nil {
			if err := m.FilteredBackupTaskIDMap[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filteredBackupTaskIdMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filteredBackupTaskIdMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this replication info base based on the context it is used
func (m *ReplicationInfoBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExistingDataLockConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExistingWormRetentionOnRx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilteredBackupTaskIDMap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplicationInfoBase) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ReplicationInfoBase) contextValidateExistingDataLockConstraints(ctx context.Context, formats strfmt.Registry) error {

	if m.ExistingDataLockConstraints != nil {

		if swag.IsZero(m.ExistingDataLockConstraints) { // not required
			return nil
		}

		if err := m.ExistingDataLockConstraints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingDataLockConstraints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingDataLockConstraints")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationInfoBase) contextValidateExistingWormRetentionOnRx(ctx context.Context, formats strfmt.Registry) error {

	if m.ExistingWormRetentionOnRx != nil {

		if swag.IsZero(m.ExistingWormRetentionOnRx) { // not required
			return nil
		}

		if err := m.ExistingWormRetentionOnRx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingWormRetentionOnRx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("existingWormRetentionOnRx")
			}
			return err
		}
	}

	return nil
}

func (m *ReplicationInfoBase) contextValidateFilteredBackupTaskIDMap(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FilteredBackupTaskIDMap); i++ {

		if m.FilteredBackupTaskIDMap[i] != nil {

			if swag.IsZero(m.FilteredBackupTaskIDMap[i]) { // not required
				return nil
			}

			if err := m.FilteredBackupTaskIDMap[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filteredBackupTaskIdMap" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filteredBackupTaskIdMap" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationInfoBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationInfoBase) UnmarshalBinary(b []byte) error {
	var res ReplicationInfoBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

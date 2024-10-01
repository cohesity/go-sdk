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

// BackupTaskStateProtoAppEntityState If this backup task is backing up application entities (such as
// SQL/Oracle dbs), this message contains the state of those entities.
// Note that this information is populated for Iris on the fly, and is not
// persisted in Magneto's WAL.
//
// swagger:model BackupTaskStateProto_AppEntityState
type BackupTaskStateProtoAppEntityState struct {

	// The application entity being backed up.
	AppEntity *PrivateEntityProto `json:"appEntity,omitempty"`

	// The end time of the backup.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// The error encountered (if any) during backup. Only valid if
	// the 'public_status' is kFinished.
	Error *PrivateErrorProto `json:"error,omitempty"`

	// The progress monitor path.
	ProgressMonitorTaskPath *string `json:"progressMonitorTaskPath,omitempty"`

	// Iris-facing task state.
	PublicStatus *int32 `json:"publicStatus,omitempty"`

	// The relative snapshot dir where the application level data is stored.
	RelativeSnapshotDir *string `json:"relativeSnapshotDir,omitempty"`

	// The start time of the backup.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// The total number of bytes read from source.
	TotalBytesReadFromSource *int64 `json:"totalBytesReadFromSource,omitempty"`

	// The total number of logical bytes to be backed up.
	TotalLogicalBytes *int64 `json:"totalLogicalBytes,omitempty"`

	// The warnings encountered (if any) during backup.
	Warnings []*PrivateErrorProto `json:"warnings"`
}

// Validate validates this backup task state proto app entity state
func (m *BackupTaskStateProtoAppEntityState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateProtoAppEntityState) validateAppEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.AppEntity) { // not required
		return nil
	}

	if m.AppEntity != nil {
		if err := m.AppEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appEntity")
			}
			return err
		}
	}

	return nil
}

func (m *BackupTaskStateProtoAppEntityState) validateError(formats strfmt.Registry) error {
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

func (m *BackupTaskStateProtoAppEntityState) validateWarnings(formats strfmt.Registry) error {
	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	for i := 0; i < len(m.Warnings); i++ {
		if swag.IsZero(m.Warnings[i]) { // not required
			continue
		}

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this backup task state proto app entity state based on the context it is used
func (m *BackupTaskStateProtoAppEntityState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAppEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarnings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupTaskStateProtoAppEntityState) contextValidateAppEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.AppEntity != nil {

		if swag.IsZero(m.AppEntity) { // not required
			return nil
		}

		if err := m.AppEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("appEntity")
			}
			return err
		}
	}

	return nil
}

func (m *BackupTaskStateProtoAppEntityState) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BackupTaskStateProtoAppEntityState) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Warnings); i++ {

		if m.Warnings[i] != nil {

			if swag.IsZero(m.Warnings[i]) { // not required
				return nil
			}

			if err := m.Warnings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupTaskStateProtoAppEntityState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupTaskStateProtoAppEntityState) UnmarshalBinary(b []byte) error {
	var res BackupTaskStateProtoAppEntityState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProtectionJobRunInformationJobUID Protection Job Uid.
//
// Specifies the globally unique id of the original Protection Job
// that archived the data to the Vault. This id is assigned by the
// original Cluster that archived the data.
//
// swagger:model remoteProtectionJobRunInformationJobUid
type RemoteProtectionJobRunInformationJobUID struct {
	UniversalID
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RemoteProtectionJobRunInformationJobUID) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 UniversalID
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.UniversalID = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RemoteProtectionJobRunInformationJobUID) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.UniversalID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this remote protection job run information job Uid
func (m *RemoteProtectionJobRunInformationJobUID) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UniversalID
	if err := m.UniversalID.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this remote protection job run information job Uid based on the context it is used
func (m *RemoteProtectionJobRunInformationJobUID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with UniversalID
	if err := m.UniversalID.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProtectionJobRunInformationJobUID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProtectionJobRunInformationJobUID) UnmarshalBinary(b []byte) error {
	var res RemoteProtectionJobRunInformationJobUID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

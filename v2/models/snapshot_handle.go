// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SnapshotHandle Information about the snapshots that exist on the system.
//
// swagger:model SnapshotHandle
type SnapshotHandle struct {

	// Specifies a distinct value that's unique to a source.
	JobUID *int64 `json:"jobUid,omitempty"`

	// Run start time of the Magneto job run which has taken this snapshot.
	RunStartTimeUsecs *int64 `json:"runStartTimeUsecs,omitempty"`
}

// Validate validates this snapshot handle
func (m *SnapshotHandle) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this snapshot handle based on context it is used
func (m *SnapshotHandle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotHandle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotHandle) UnmarshalBinary(b []byte) error {
	var res SnapshotHandle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

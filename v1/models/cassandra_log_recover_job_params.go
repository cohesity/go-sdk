// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CassandraLogRecoverJobParams Contains additional params for cassandra log recovery workflow.
//
// swagger:model CassandraLogRecoverJobParams
type CassandraLogRecoverJobParams struct {

	// This is the end time from when logs should be replayed.
	EndTimeForLogReplayInUsecs *int64 `json:"endTimeForLogReplayInUsecs,omitempty"`

	// The view box name where commit logs are present.
	LogBackupViewBoxName *string `json:"logBackupViewBoxName,omitempty"`

	// The view name from where commit logs should be restored.
	LogBackupViewName *string `json:"logBackupViewName,omitempty"`

	// Objects are of the form keyspace.table. If a full keyspace is selected to
	// be restored, it is expanded before passing to imanis.
	ObjectNames []string `json:"objectNames"`

	// This is the start time from when logs should be replayed.
	StartTimeForLogReplayInUsecs *int64 `json:"startTimeForLogReplayInUsecs,omitempty"`
}

// Validate validates this cassandra log recover job params
func (m *CassandraLogRecoverJobParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cassandra log recover job params based on context it is used
func (m *CassandraLogRecoverJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CassandraLogRecoverJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CassandraLogRecoverJobParams) UnmarshalBinary(b []byte) error {
	var res CassandraLogRecoverJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

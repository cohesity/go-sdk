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

// ReplicationTargets ReplicationTargets
//
// # Replication Targets
//
// swagger:model ReplicationTargets
type ReplicationTargets struct {

	// Specifies the replication target.
	// Enum: ["RemoteCluster","AWS","Azure"]
	ReplicationTargets string `json:"replicationTargets,omitempty"`
}

// Validate validates this replication targets
func (m *ReplicationTargets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicationTargets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var replicationTargetsTypeReplicationTargetsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RemoteCluster","AWS","Azure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		replicationTargetsTypeReplicationTargetsPropEnum = append(replicationTargetsTypeReplicationTargetsPropEnum, v)
	}
}

const (

	// ReplicationTargetsReplicationTargetsRemoteCluster captures enum value "RemoteCluster"
	ReplicationTargetsReplicationTargetsRemoteCluster string = "RemoteCluster"

	// ReplicationTargetsReplicationTargetsAWS captures enum value "AWS"
	ReplicationTargetsReplicationTargetsAWS string = "AWS"

	// ReplicationTargetsReplicationTargetsAzure captures enum value "Azure"
	ReplicationTargetsReplicationTargetsAzure string = "Azure"
)

// prop value enum
func (m *ReplicationTargets) validateReplicationTargetsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, replicationTargetsTypeReplicationTargetsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReplicationTargets) validateReplicationTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationTargets) { // not required
		return nil
	}

	// value enum
	if err := m.validateReplicationTargetsEnum("replicationTargets", "body", m.ReplicationTargets); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this replication targets based on context it is used
func (m *ReplicationTargets) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReplicationTargets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplicationTargets) UnmarshalBinary(b []byte) error {
	var res ReplicationTargets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

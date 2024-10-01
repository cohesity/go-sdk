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

// CassandraKeyspace Cassandra Keyspace Info.
//
// Specifies an Object containing information about a Cassandra Keyspace.
//
// swagger:model CassandraKeyspace
type CassandraKeyspace struct {

	// Number of documents in this bucket.
	ChildrenCount *uint32 `json:"childrenCount,omitempty"`

	// If the replication strategy is set as kNetwork, then
	// dc_list will have a list of data centers to which the
	// keyspace is being replicated to.
	DcList []string `json:"dcList"`

	// Replication stragegy for the keyspace.
	// Specifies the type of an Cassandra source entity.
	// Enum: ["kSimple","kNetwork","kUnsupported"]
	ReplicationStrategy *string `json:"replicationStrategy,omitempty"`

	// Specifies Type of Keyspace.
	// Specifies the type of an Cassandra keyspace entity.
	// Enum: ["kRegular","kGraph","kSystem"]
	Type *string `json:"type,omitempty"`
}

// Validate validates this cassandra keyspace
func (m *CassandraKeyspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReplicationStrategy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cassandraKeyspaceTypeReplicationStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kSimple","kNetwork","kUnsupported"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cassandraKeyspaceTypeReplicationStrategyPropEnum = append(cassandraKeyspaceTypeReplicationStrategyPropEnum, v)
	}
}

const (

	// CassandraKeyspaceReplicationStrategyKSimple captures enum value "kSimple"
	CassandraKeyspaceReplicationStrategyKSimple string = "kSimple"

	// CassandraKeyspaceReplicationStrategyKNetwork captures enum value "kNetwork"
	CassandraKeyspaceReplicationStrategyKNetwork string = "kNetwork"

	// CassandraKeyspaceReplicationStrategyKUnsupported captures enum value "kUnsupported"
	CassandraKeyspaceReplicationStrategyKUnsupported string = "kUnsupported"
)

// prop value enum
func (m *CassandraKeyspace) validateReplicationStrategyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cassandraKeyspaceTypeReplicationStrategyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CassandraKeyspace) validateReplicationStrategy(formats strfmt.Registry) error {
	if swag.IsZero(m.ReplicationStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateReplicationStrategyEnum("replicationStrategy", "body", *m.ReplicationStrategy); err != nil {
		return err
	}

	return nil
}

var cassandraKeyspaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRegular","kGraph","kSystem"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cassandraKeyspaceTypeTypePropEnum = append(cassandraKeyspaceTypeTypePropEnum, v)
	}
}

const (

	// CassandraKeyspaceTypeKRegular captures enum value "kRegular"
	CassandraKeyspaceTypeKRegular string = "kRegular"

	// CassandraKeyspaceTypeKGraph captures enum value "kGraph"
	CassandraKeyspaceTypeKGraph string = "kGraph"

	// CassandraKeyspaceTypeKSystem captures enum value "kSystem"
	CassandraKeyspaceTypeKSystem string = "kSystem"
)

// prop value enum
func (m *CassandraKeyspace) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cassandraKeyspaceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CassandraKeyspace) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cassandra keyspace based on context it is used
func (m *CassandraKeyspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CassandraKeyspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CassandraKeyspace) UnmarshalBinary(b []byte) error {
	var res CassandraKeyspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// MongoDBAuthType Enumeration of all the MongoDB Authentication types.
//
// Enumeration of all the MongoDB Authentication types.
//
// swagger:model MongoDBAuthType
type MongoDBAuthType struct {

	// Enumeration of all the MongoDB Authentication.
	// Enum: ["SCRAM","LDAP","NONE","KERBEROS"]
	MongoDBAuthType string `json:"MongoDBAuthType,omitempty"`
}

// Validate validates this mongo d b auth type
func (m *MongoDBAuthType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMongoDBAuthType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mongoDBAuthTypeTypeMongoDBAuthTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SCRAM","LDAP","NONE","KERBEROS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDBAuthTypeTypeMongoDBAuthTypePropEnum = append(mongoDBAuthTypeTypeMongoDBAuthTypePropEnum, v)
	}
}

const (

	// MongoDBAuthTypeMongoDBAuthTypeSCRAM captures enum value "SCRAM"
	MongoDBAuthTypeMongoDBAuthTypeSCRAM string = "SCRAM"

	// MongoDBAuthTypeMongoDBAuthTypeLDAP captures enum value "LDAP"
	MongoDBAuthTypeMongoDBAuthTypeLDAP string = "LDAP"

	// MongoDBAuthTypeMongoDBAuthTypeNONE captures enum value "NONE"
	MongoDBAuthTypeMongoDBAuthTypeNONE string = "NONE"

	// MongoDBAuthTypeMongoDBAuthTypeKERBEROS captures enum value "KERBEROS"
	MongoDBAuthTypeMongoDBAuthTypeKERBEROS string = "KERBEROS"
)

// prop value enum
func (m *MongoDBAuthType) validateMongoDBAuthTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDBAuthTypeTypeMongoDBAuthTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDBAuthType) validateMongoDBAuthType(formats strfmt.Registry) error {
	if swag.IsZero(m.MongoDBAuthType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMongoDBAuthTypeEnum("MongoDBAuthType", "body", m.MongoDBAuthType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mongo d b auth type based on context it is used
func (m *MongoDBAuthType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MongoDBAuthType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDBAuthType) UnmarshalBinary(b []byte) error {
	var res MongoDBAuthType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

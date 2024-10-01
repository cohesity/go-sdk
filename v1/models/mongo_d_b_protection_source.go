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

// MongoDBProtectionSource MongoDB Protection Source.
//
// Specifies an Object representing MongoDB.
//
// swagger:model MongoDBProtectionSource
type MongoDBProtectionSource struct {

	// Information of a mongodb cluster, only valid for an entity of type
	// kCluster.
	ClusterInfo *MongoDBCluster `json:"clusterInfo,omitempty"`

	// Information about a mongodb collection, only valid for an entity of type
	// kCollection.
	CollectionInfo *MongoDBCollection `json:"collectionInfo,omitempty"`

	// Information of a mongodb database, only valid for an entity of type
	// kDatabase.
	DatabaseInfo *MongoDBDatabase `json:"databaseInfo,omitempty"`

	// Specifies the instance name of the MongoDB entity.
	Name *string `json:"name,omitempty"`

	// Specifies the type of the managed Object in MongoDB Protection Source.
	// Specifies the type of an MongoDB source entity.
	// 'kCluster' indicates a mongodb cluster distributed over several physical
	// nodes.
	// 'kDatabase' indicates a Database within the MongoDB environment.
	// 'kCollection' indicates a Collection in the MongoDB enironment.
	// Enum: ["kCluster","kDatabase","kCollection"]
	Type *string `json:"type,omitempty"`

	// Specifies the UUID for the MongoDB entity.
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this mongo d b protection source
func (m *MongoDBProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCollectionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabaseInfo(formats); err != nil {
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

func (m *MongoDBProtectionSource) validateClusterInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterInfo) { // not required
		return nil
	}

	if m.ClusterInfo != nil {
		if err := m.ClusterInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDBProtectionSource) validateCollectionInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CollectionInfo) { // not required
		return nil
	}

	if m.CollectionInfo != nil {
		if err := m.CollectionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDBProtectionSource) validateDatabaseInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseInfo) { // not required
		return nil
	}

	if m.DatabaseInfo != nil {
		if err := m.DatabaseInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseInfo")
			}
			return err
		}
	}

	return nil
}

var mongoDBProtectionSourceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kCluster","kDatabase","kCollection"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mongoDBProtectionSourceTypeTypePropEnum = append(mongoDBProtectionSourceTypeTypePropEnum, v)
	}
}

const (

	// MongoDBProtectionSourceTypeKCluster captures enum value "kCluster"
	MongoDBProtectionSourceTypeKCluster string = "kCluster"

	// MongoDBProtectionSourceTypeKDatabase captures enum value "kDatabase"
	MongoDBProtectionSourceTypeKDatabase string = "kDatabase"

	// MongoDBProtectionSourceTypeKCollection captures enum value "kCollection"
	MongoDBProtectionSourceTypeKCollection string = "kCollection"
)

// prop value enum
func (m *MongoDBProtectionSource) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mongoDBProtectionSourceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MongoDBProtectionSource) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mongo d b protection source based on the context it is used
func (m *MongoDBProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCollectionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatabaseInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDBProtectionSource) contextValidateClusterInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterInfo != nil {

		if swag.IsZero(m.ClusterInfo) { // not required
			return nil
		}

		if err := m.ClusterInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clusterInfo")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDBProtectionSource) contextValidateCollectionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CollectionInfo != nil {

		if swag.IsZero(m.CollectionInfo) { // not required
			return nil
		}

		if err := m.CollectionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collectionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("collectionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *MongoDBProtectionSource) contextValidateDatabaseInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.DatabaseInfo != nil {

		if swag.IsZero(m.DatabaseInfo) { // not required
			return nil
		}

		if err := m.DatabaseInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MongoDBProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDBProtectionSource) UnmarshalBinary(b []byte) error {
	var res MongoDBProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

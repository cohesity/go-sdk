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

// MongoDBBackupJobParams Contains any additional mongodb environment specific backup params at the
// job level.
//
// swagger:model MongoDBBackupJobParams
type MongoDBBackupJobParams struct {

	// Additional parameters required for Mongo backup
	MongodbAdditionalInfo *MongoDBAdditionalParams `json:"mongodbAdditionalInfo,omitempty"`
}

// Validate validates this mongo d b backup job params
func (m *MongoDBBackupJobParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMongodbAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDBBackupJobParams) validateMongodbAdditionalInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MongodbAdditionalInfo) { // not required
		return nil
	}

	if m.MongodbAdditionalInfo != nil {
		if err := m.MongodbAdditionalInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongodbAdditionalInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongodbAdditionalInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mongo d b backup job params based on the context it is used
func (m *MongoDBBackupJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMongodbAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MongoDBBackupJobParams) contextValidateMongodbAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MongodbAdditionalInfo != nil {

		if swag.IsZero(m.MongodbAdditionalInfo) { // not required
			return nil
		}

		if err := m.MongodbAdditionalInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mongodbAdditionalInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mongodbAdditionalInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MongoDBBackupJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MongoDBBackupJobParams) UnmarshalBinary(b []byte) error {
	var res MongoDBBackupJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

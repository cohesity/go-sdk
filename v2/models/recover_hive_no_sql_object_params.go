// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RecoverHiveNoSQLObjectParams Recover Hive Object Params.
//
// Specifies the fully qualified object name and other attributes of each object to be recovered.
//
// swagger:model RecoverHiveNoSqlObjectParams
type RecoverHiveNoSQLObjectParams struct {

	// Specifies the fully qualified name of the object to be restored.
	// Required: true
	ObjectName *string `json:"objectName"`
}

// Validate validates this recover hive no Sql object params
func (m *RecoverHiveNoSQLObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjectName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecoverHiveNoSQLObjectParams) validateObjectName(formats strfmt.Registry) error {

	if err := validate.Required("objectName", "body", m.ObjectName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recover hive no Sql object params based on context it is used
func (m *RecoverHiveNoSQLObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecoverHiveNoSQLObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecoverHiveNoSQLObjectParams) UnmarshalBinary(b []byte) error {
	var res RecoverHiveNoSQLObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

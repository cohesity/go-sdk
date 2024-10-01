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

// AwsRdsPostgresProtectionGroupObjectParams AWS RDS Postgres Ingest Protection Group Object Params.
//
// Specifies the object parameters to create an AWS RDS Postgres Ingest Protection Group.
//
// swagger:model AwsRdsPostgresProtectionGroupObjectParams
type AwsRdsPostgresProtectionGroupObjectParams struct {

	// Specifies the id of the object.
	// Required: true
	ID *int64 `json:"id"`

	// Specifies the name of the database.
	// Read Only: true
	Name *string `json:"name,omitempty"`
}

// Validate validates this aws rds postgres protection group object params
func (m *AwsRdsPostgresProtectionGroupObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRdsPostgresProtectionGroupObjectParams) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this aws rds postgres protection group object params based on the context it is used
func (m *AwsRdsPostgresProtectionGroupObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AwsRdsPostgresProtectionGroupObjectParams) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AwsRdsPostgresProtectionGroupObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AwsRdsPostgresProtectionGroupObjectParams) UnmarshalBinary(b []byte) error {
	var res AwsRdsPostgresProtectionGroupObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// CommonMssqlNativeObjectProtectionParams Specifies the common params to create a Native based MSSQL Object Protection.
//
// swagger:model CommonMssqlNativeObjectProtectionParams
type CommonMssqlNativeObjectProtectionParams struct {

	// Specifies the number of streams to be used.
	NumStreams *int32 `json:"numStreams,omitempty"`

	// Specifies the WithClause to be used.
	WithClause *string `json:"withClause,omitempty"`

	CommonMSSQLProtectionGroupParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CommonMssqlNativeObjectProtectionParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		NumStreams *int32 `json:"numStreams,omitempty"`

		WithClause *string `json:"withClause,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.NumStreams = dataAO0.NumStreams

	m.WithClause = dataAO0.WithClause

	// AO1
	var aO1 CommonMSSQLProtectionGroupParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.CommonMSSQLProtectionGroupParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CommonMssqlNativeObjectProtectionParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		NumStreams *int32 `json:"numStreams,omitempty"`

		WithClause *string `json:"withClause,omitempty"`
	}

	dataAO0.NumStreams = m.NumStreams

	dataAO0.WithClause = m.WithClause

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.CommonMSSQLProtectionGroupParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this common mssql native object protection params
func (m *CommonMssqlNativeObjectProtectionParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonMSSQLProtectionGroupParams
	if err := m.CommonMSSQLProtectionGroupParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this common mssql native object protection params based on the context it is used
func (m *CommonMssqlNativeObjectProtectionParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CommonMSSQLProtectionGroupParams
	if err := m.CommonMSSQLProtectionGroupParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *CommonMssqlNativeObjectProtectionParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonMssqlNativeObjectProtectionParams) UnmarshalBinary(b []byte) error {
	var res CommonMssqlNativeObjectProtectionParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

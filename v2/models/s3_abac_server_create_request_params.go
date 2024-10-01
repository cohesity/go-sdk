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

// S3AbacServerCreateRequestParams Specifies Request Parameters for adding a S3 ABAC server.
//
// swagger:model S3AbacServerCreateRequestParams
type S3AbacServerCreateRequestParams struct {

	// Specifies the hostname of S3 ABAC server.
	// Required: true
	Hostname *string `json:"hostname"`

	// Specifies the port of S3 ABAC server.
	// Required: true
	Port *int64 `json:"port"`

	S3AbacServerUpdateRequestParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *S3AbacServerCreateRequestParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Hostname *string `json:"hostname"`

		Port *int64 `json:"port"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Hostname = dataAO0.Hostname

	m.Port = dataAO0.Port

	// AO1
	var aO1 S3AbacServerUpdateRequestParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.S3AbacServerUpdateRequestParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m S3AbacServerCreateRequestParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		Hostname *string `json:"hostname"`

		Port *int64 `json:"port"`
	}

	dataAO0.Hostname = m.Hostname

	dataAO0.Port = m.Port

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.S3AbacServerUpdateRequestParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this s3 abac server create request params
func (m *S3AbacServerCreateRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with S3AbacServerUpdateRequestParams
	if err := m.S3AbacServerUpdateRequestParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *S3AbacServerCreateRequestParams) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

func (m *S3AbacServerCreateRequestParams) validatePort(formats strfmt.Registry) error {

	if err := validate.Required("port", "body", m.Port); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this s3 abac server create request params based on the context it is used
func (m *S3AbacServerCreateRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with S3AbacServerUpdateRequestParams
	if err := m.S3AbacServerUpdateRequestParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *S3AbacServerCreateRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *S3AbacServerCreateRequestParams) UnmarshalBinary(b []byte) error {
	var res S3AbacServerCreateRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// UptieringFileAgePolicy Specifies the file's selection rule by file age for up tiering data
// tiering task eg.
// 1. select files last accessed 2 weeks ago.
// 2. select files last modified 1 month ago.
//
// swagger:model UptieringFileAgePolicy
type UptieringFileAgePolicy struct {

	// Specifies the condition for the file age.
	// Enum: ["LastAccessed","LastModified"]
	Condition *string `json:"condition,omitempty"`

	// Specifies the number of msecs used for file selection.
	AgeMsecs *int64 `json:"ageMsecs,omitempty"`

	// Specifies number of file access in last ageMsecs.
	NumFileAccess *uint32 `json:"numFileAccess,omitempty"`
}

// Validate validates this uptiering file age policy
func (m *UptieringFileAgePolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCondition(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var uptieringFileAgePolicyTypeConditionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LastAccessed","LastModified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		uptieringFileAgePolicyTypeConditionPropEnum = append(uptieringFileAgePolicyTypeConditionPropEnum, v)
	}
}

const (

	// UptieringFileAgePolicyConditionLastAccessed captures enum value "LastAccessed"
	UptieringFileAgePolicyConditionLastAccessed string = "LastAccessed"

	// UptieringFileAgePolicyConditionLastModified captures enum value "LastModified"
	UptieringFileAgePolicyConditionLastModified string = "LastModified"
)

// prop value enum
func (m *UptieringFileAgePolicy) validateConditionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, uptieringFileAgePolicyTypeConditionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UptieringFileAgePolicy) validateCondition(formats strfmt.Registry) error {
	if swag.IsZero(m.Condition) { // not required
		return nil
	}

	// value enum
	if err := m.validateConditionEnum("condition", "body", *m.Condition); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this uptiering file age policy based on context it is used
func (m *UptieringFileAgePolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UptieringFileAgePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UptieringFileAgePolicy) UnmarshalBinary(b []byte) error {
	var res UptieringFileAgePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

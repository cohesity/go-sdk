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

// ProtectionSourcesSummaryStatsProtectionSource Specifies the leaf Protection Source Object (such as VM).
// Snapshot summary statistics are reported for this Protection Source
// Object.
//
// swagger:model protectionSourcesSummaryStatsProtectionSource
type ProtectionSourcesSummaryStatsProtectionSource struct {
	ProtectionSource
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionSourcesSummaryStatsProtectionSource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProtectionSource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProtectionSource = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionSourcesSummaryStatsProtectionSource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ProtectionSource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection sources summary stats protection source
func (m *ProtectionSourcesSummaryStatsProtectionSource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProtectionSource
	if err := m.ProtectionSource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this protection sources summary stats protection source based on the context it is used
func (m *ProtectionSourcesSummaryStatsProtectionSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProtectionSource
	if err := m.ProtectionSource.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionSourcesSummaryStatsProtectionSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionSourcesSummaryStatsProtectionSource) UnmarshalBinary(b []byte) error {
	var res ProtectionSourcesSummaryStatsProtectionSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

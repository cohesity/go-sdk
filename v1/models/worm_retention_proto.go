// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WormRetentionProto Message that specifies the WORM attributes. WORM attributes can be
// associated with any of the following:
// 1. backup policy: compliance or administrative policy with worm retention.
// 2. backup runs: worm retention inherited from policy at successful backup
// run completion..
// 3. backup tasks do not inherit WORM retention. Instead they check for WORM
// property on the corresponding backup run.
// There are no WORM attributes associated with the backup job.
//
// swagger:model WormRetentionProto
type WormRetentionProto struct {

	// Whether objects in the external target associated with this policy need to
	// be made immutable.
	EnableWormOnExternalTarget *bool `json:"enableWormOnExternalTarget,omitempty"`

	// The type of WORM policy set on this run. This field is irrelevant
	// while objects are on legal hold. Objects put on legal hold automatically
	// raise the WORM level on the object behaviorally to the strictest level
	// i.e. kCompliance.
	PolicyType *int32 `json:"policyType,omitempty"`

	// Retention time for datalock in seconds. This will be always relative time.
	RetentionSecs *int64 `json:"retentionSecs,omitempty"`

	// Version number for this proto.
	Version *int32 `json:"version,omitempty"`
}

// Validate validates this worm retention proto
func (m *WormRetentionProto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this worm retention proto based on context it is used
func (m *WormRetentionProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WormRetentionProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WormRetentionProto) UnmarshalBinary(b []byte) error {
	var res WormRetentionProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VMNumReplicas Struct containing vm-name and replica count as members.
//
// swagger:model VmNumReplicas
type VMNumReplicas struct {

	// Replica count.
	NumReplicas *int64 `json:"numReplicas,omitempty"`

	// Vm-name.
	VMName *string `json:"vmName,omitempty"`
}

// Validate validates this Vm num replicas
func (m *VMNumReplicas) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Vm num replicas based on context it is used
func (m *VMNumReplicas) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMNumReplicas) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMNumReplicas) UnmarshalBinary(b []byte) error {
	var res VMNumReplicas
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

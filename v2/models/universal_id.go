// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UniversalID universal Id
//
// swagger:model UniversalId
type UniversalID struct {

	// The id of the cluster at which the object was created.
	ClusterID *int64 `json:"clusterId,omitempty"`

	// The incarnation id of the above cluster.
	ClusterIncarnationID *int64 `json:"clusterIncarnationId,omitempty"`

	// The object id - this is unique within the above cluster.
	ObjectID *int64 `json:"objectId,omitempty"`
}

// Validate validates this universal Id
func (m *UniversalID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this universal Id based on context it is used
func (m *UniversalID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UniversalID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UniversalID) UnmarshalBinary(b []byte) error {
	var res UniversalID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

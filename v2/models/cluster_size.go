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

// ClusterSize Cluster Size
//
// Cluster Size of cloud platforms.
//
// swagger:model ClusterSize
type ClusterSize struct {

	// Specifies the cluster size.
	// Enum: ["Small","Medium","Large","XLarge","NextGen"]
	ClusterSize string `json:"clusterSize,omitempty"`
}

// Validate validates this cluster size
func (m *ClusterSize) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var clusterSizeTypeClusterSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Small","Medium","Large","XLarge","NextGen"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterSizeTypeClusterSizePropEnum = append(clusterSizeTypeClusterSizePropEnum, v)
	}
}

const (

	// ClusterSizeClusterSizeSmall captures enum value "Small"
	ClusterSizeClusterSizeSmall string = "Small"

	// ClusterSizeClusterSizeMedium captures enum value "Medium"
	ClusterSizeClusterSizeMedium string = "Medium"

	// ClusterSizeClusterSizeLarge captures enum value "Large"
	ClusterSizeClusterSizeLarge string = "Large"

	// ClusterSizeClusterSizeXLarge captures enum value "XLarge"
	ClusterSizeClusterSizeXLarge string = "XLarge"

	// ClusterSizeClusterSizeNextGen captures enum value "NextGen"
	ClusterSizeClusterSizeNextGen string = "NextGen"
)

// prop value enum
func (m *ClusterSize) validateClusterSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterSizeTypeClusterSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ClusterSize) validateClusterSize(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterSize) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterSizeEnum("clusterSize", "body", m.ClusterSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster size based on context it is used
func (m *ClusterSize) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSize) UnmarshalBinary(b []byte) error {
	var res ClusterSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

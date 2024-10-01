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

// McmCohesionClaimRequestParams Request params to claim cohesion appliance to Helios.
//
// Specifies the request params to claim cohesion appliance to Helios.
//
// swagger:model McmCohesionClaimRequestParams
type McmCohesionClaimRequestParams struct {

	// Specifies the Id of the cohesion appliance with AWS.
	// Required: true
	ApplianceID *string `json:"applianceId"`

	// Specifies the name of the cohesion appliance.
	// Required: true
	ApplianceName *string `json:"applianceName"`

	// Specifies the cluster id.
	// Required: true
	ClusterID *int64 `json:"clusterId"`

	// Specifies the cluster incarnation id.
	// Required: true
	ClusterIncarnationID *int64 `json:"clusterIncarnationId"`

	// Claim token used for authentication.
	// Required: true
	ClaimToken *string `json:"claimToken"`
}

// Validate validates this mcm cohesion claim request params
func (m *McmCohesionClaimRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplianceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplianceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterIncarnationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClaimToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *McmCohesionClaimRequestParams) validateApplianceID(formats strfmt.Registry) error {

	if err := validate.Required("applianceId", "body", m.ApplianceID); err != nil {
		return err
	}

	return nil
}

func (m *McmCohesionClaimRequestParams) validateApplianceName(formats strfmt.Registry) error {

	if err := validate.Required("applianceName", "body", m.ApplianceName); err != nil {
		return err
	}

	return nil
}

func (m *McmCohesionClaimRequestParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("clusterId", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *McmCohesionClaimRequestParams) validateClusterIncarnationID(formats strfmt.Registry) error {

	if err := validate.Required("clusterIncarnationId", "body", m.ClusterIncarnationID); err != nil {
		return err
	}

	return nil
}

func (m *McmCohesionClaimRequestParams) validateClaimToken(formats strfmt.Registry) error {

	if err := validate.Required("claimToken", "body", m.ClaimToken); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this mcm cohesion claim request params based on context it is used
func (m *McmCohesionClaimRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *McmCohesionClaimRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *McmCohesionClaimRequestParams) UnmarshalBinary(b []byte) error {
	var res McmCohesionClaimRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

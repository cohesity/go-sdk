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

// ExternallyTriggeredJobParams Specifies the externally triggered job paramters.
//
// swagger:model ExternallyTriggeredJobParams
type ExternallyTriggeredJobParams struct {

	// Specifies the client type of the externally triggered backup job.
	// Enum: ["Generic","SBT"]
	ClientType *string `json:"clientType,omitempty"`

	// Specifies all the tags of the externally triggered job.
	Tags []string `json:"tags,omitempty"`

	// Specifies the SBT parameters for the externally triggered backup job.
	SbtParams *ExternallyTriggeredSbtParams `json:"sbtParams,omitempty"`
}

// Validate validates this externally triggered job params
func (m *ExternallyTriggeredJobParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSbtParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var externallyTriggeredJobParamsTypeClientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Generic","SBT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externallyTriggeredJobParamsTypeClientTypePropEnum = append(externallyTriggeredJobParamsTypeClientTypePropEnum, v)
	}
}

const (

	// ExternallyTriggeredJobParamsClientTypeGeneric captures enum value "Generic"
	ExternallyTriggeredJobParamsClientTypeGeneric string = "Generic"

	// ExternallyTriggeredJobParamsClientTypeSBT captures enum value "SBT"
	ExternallyTriggeredJobParamsClientTypeSBT string = "SBT"
)

// prop value enum
func (m *ExternallyTriggeredJobParams) validateClientTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externallyTriggeredJobParamsTypeClientTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternallyTriggeredJobParams) validateClientType(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientType) { // not required
		return nil
	}

	// value enum
	if err := m.validateClientTypeEnum("clientType", "body", *m.ClientType); err != nil {
		return err
	}

	return nil
}

func (m *ExternallyTriggeredJobParams) validateSbtParams(formats strfmt.Registry) error {
	if swag.IsZero(m.SbtParams) { // not required
		return nil
	}

	if m.SbtParams != nil {
		if err := m.SbtParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbtParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbtParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this externally triggered job params based on the context it is used
func (m *ExternallyTriggeredJobParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSbtParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExternallyTriggeredJobParams) contextValidateSbtParams(ctx context.Context, formats strfmt.Registry) error {

	if m.SbtParams != nil {

		if swag.IsZero(m.SbtParams) { // not required
			return nil
		}

		if err := m.SbtParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sbtParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sbtParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExternallyTriggeredJobParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternallyTriggeredJobParams) UnmarshalBinary(b []byte) error {
	var res ExternallyTriggeredJobParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

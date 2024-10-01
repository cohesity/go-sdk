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

// EndpointCheckResult Endpoint Check Result.
//
// Specify the connectivity check results on each endpoint.
//
// swagger:model EndpointCheckResult
type EndpointCheckResult struct {

	// Specifies the name of the check.
	CheckName *string `json:"checkName,omitempty"`

	// Specifies the status of the check.
	// Enum: ["kStateUnknown","kStateOk","kStateFail","kStateInProgress"]
	Status *string `json:"status,omitempty"`

	// Specifies the error message to help troubleshoot.
	ErrorMsg *string `json:"errorMsg,omitempty"`
}

// Validate validates this endpoint check result
func (m *EndpointCheckResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var endpointCheckResultTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kStateUnknown","kStateOk","kStateFail","kStateInProgress"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointCheckResultTypeStatusPropEnum = append(endpointCheckResultTypeStatusPropEnum, v)
	}
}

const (

	// EndpointCheckResultStatusKStateUnknown captures enum value "kStateUnknown"
	EndpointCheckResultStatusKStateUnknown string = "kStateUnknown"

	// EndpointCheckResultStatusKStateOk captures enum value "kStateOk"
	EndpointCheckResultStatusKStateOk string = "kStateOk"

	// EndpointCheckResultStatusKStateFail captures enum value "kStateFail"
	EndpointCheckResultStatusKStateFail string = "kStateFail"

	// EndpointCheckResultStatusKStateInProgress captures enum value "kStateInProgress"
	EndpointCheckResultStatusKStateInProgress string = "kStateInProgress"
)

// prop value enum
func (m *EndpointCheckResult) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, endpointCheckResultTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EndpointCheckResult) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this endpoint check result based on context it is used
func (m *EndpointCheckResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EndpointCheckResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointCheckResult) UnmarshalBinary(b []byte) error {
	var res EndpointCheckResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

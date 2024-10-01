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

// CommonTdmTaskRequestParams Specifies the common parameters used in TDM Task Creation.
//
// swagger:model CommonTdmTaskRequestParams
type CommonTdmTaskRequestParams struct {

	// Specifies the name of the task.
	Name *string `json:"name,omitempty"`

	// Specifies the TDM Task action.
	// Required: true
	// Enum: ["clone","teardown","refresh","snapshot","rewind"]
	Action *string `json:"action"`
}

// Validate validates this common tdm task request params
func (m *CommonTdmTaskRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var commonTdmTaskRequestParamsTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["clone","teardown","refresh","snapshot","rewind"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonTdmTaskRequestParamsTypeActionPropEnum = append(commonTdmTaskRequestParamsTypeActionPropEnum, v)
	}
}

const (

	// CommonTdmTaskRequestParamsActionClone captures enum value "clone"
	CommonTdmTaskRequestParamsActionClone string = "clone"

	// CommonTdmTaskRequestParamsActionTeardown captures enum value "teardown"
	CommonTdmTaskRequestParamsActionTeardown string = "teardown"

	// CommonTdmTaskRequestParamsActionRefresh captures enum value "refresh"
	CommonTdmTaskRequestParamsActionRefresh string = "refresh"

	// CommonTdmTaskRequestParamsActionSnapshot captures enum value "snapshot"
	CommonTdmTaskRequestParamsActionSnapshot string = "snapshot"

	// CommonTdmTaskRequestParamsActionRewind captures enum value "rewind"
	CommonTdmTaskRequestParamsActionRewind string = "rewind"
)

// prop value enum
func (m *CommonTdmTaskRequestParams) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonTdmTaskRequestParamsTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonTdmTaskRequestParams) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	// value enum
	if err := m.validateActionEnum("action", "body", *m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this common tdm task request params based on context it is used
func (m *CommonTdmTaskRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommonTdmTaskRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonTdmTaskRequestParams) UnmarshalBinary(b []byte) error {
	var res CommonTdmTaskRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

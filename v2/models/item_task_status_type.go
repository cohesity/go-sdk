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

// ItemTaskStatusType Mailbox task items' task status types for search.
//
// Mailbox task items' task status types for search.
//
// swagger:model ItemTaskStatusType
type ItemTaskStatusType struct {

	// Specifies the Mailbox task item's task status type for search.
	// Enum: ["NotStarted","InProgress","Completed","WaitingOnOthers","Deferred"]
	Type string `json:"type,omitempty"`
}

// Validate validates this item task status type
func (m *ItemTaskStatusType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var itemTaskStatusTypeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotStarted","InProgress","Completed","WaitingOnOthers","Deferred"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		itemTaskStatusTypeTypeTypePropEnum = append(itemTaskStatusTypeTypeTypePropEnum, v)
	}
}

const (

	// ItemTaskStatusTypeTypeNotStarted captures enum value "NotStarted"
	ItemTaskStatusTypeTypeNotStarted string = "NotStarted"

	// ItemTaskStatusTypeTypeInProgress captures enum value "InProgress"
	ItemTaskStatusTypeTypeInProgress string = "InProgress"

	// ItemTaskStatusTypeTypeCompleted captures enum value "Completed"
	ItemTaskStatusTypeTypeCompleted string = "Completed"

	// ItemTaskStatusTypeTypeWaitingOnOthers captures enum value "WaitingOnOthers"
	ItemTaskStatusTypeTypeWaitingOnOthers string = "WaitingOnOthers"

	// ItemTaskStatusTypeTypeDeferred captures enum value "Deferred"
	ItemTaskStatusTypeTypeDeferred string = "Deferred"
)

// prop value enum
func (m *ItemTaskStatusType) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, itemTaskStatusTypeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ItemTaskStatusType) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this item task status type based on context it is used
func (m *ItemTaskStatusType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ItemTaskStatusType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ItemTaskStatusType) UnmarshalBinary(b []byte) error {
	var res ItemTaskStatusType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

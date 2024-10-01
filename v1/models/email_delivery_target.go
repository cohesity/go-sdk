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

// EmailDeliveryTarget EmailDeliveryTarget
// Specifies the email address and the locale setting for it.
//
// swagger:model EmailDeliveryTarget
type EmailDeliveryTarget struct {

	// email address
	EmailAddress *string `json:"emailAddress,omitempty"`

	// Specifies the language in which the emails sent to the above defined
	// mail address should be in.
	Locale *string `json:"locale,omitempty"`

	// Specifies the recipient type on how the emails are to received.
	// The email recipient type.
	// 'kTo' indicates the primary receiver type
	// 'kCc' indicates the carbon copy receiver type
	// Enum: ["kTo","kCc"]
	RecipientType *string `json:"recipientType,omitempty"`
}

// Validate validates this email delivery target
func (m *EmailDeliveryTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipientType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var emailDeliveryTargetTypeRecipientTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kTo","kCc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailDeliveryTargetTypeRecipientTypePropEnum = append(emailDeliveryTargetTypeRecipientTypePropEnum, v)
	}
}

const (

	// EmailDeliveryTargetRecipientTypeKTo captures enum value "kTo"
	EmailDeliveryTargetRecipientTypeKTo string = "kTo"

	// EmailDeliveryTargetRecipientTypeKCc captures enum value "kCc"
	EmailDeliveryTargetRecipientTypeKCc string = "kCc"
)

// prop value enum
func (m *EmailDeliveryTarget) validateRecipientTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emailDeliveryTargetTypeRecipientTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmailDeliveryTarget) validateRecipientType(formats strfmt.Registry) error {
	if swag.IsZero(m.RecipientType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecipientTypeEnum("recipientType", "body", *m.RecipientType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this email delivery target based on context it is used
func (m *EmailDeliveryTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailDeliveryTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailDeliveryTarget) UnmarshalBinary(b []byte) error {
	var res EmailDeliveryTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

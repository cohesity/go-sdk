// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HeliosStorageDomainUUIDs Storage Domain UUIDs
//
// Specifies a list of storage domain UUIDs on Helios. Only matches found in these storage domains will be returned.
//
// swagger:model HeliosStorageDomainUUIDs
type HeliosStorageDomainUUIDs []string

// Validate validates this helios storage domain u UI ds
func (m HeliosStorageDomainUUIDs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.UniqueItems("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this helios storage domain u UI ds based on context it is used
func (m HeliosStorageDomainUUIDs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

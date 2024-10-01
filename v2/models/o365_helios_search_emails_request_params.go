// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// O365HeliosSearchEmailsRequestParams O365 Search Emails Request Params
//
// Specifies email search request params specific to O365 environment.
//
// swagger:model O365HeliosSearchEmailsRequestParams
type O365HeliosSearchEmailsRequestParams struct {
	GlobalClusterIdentifier

	// Specifies the domain Ids in which mailboxes are registered.
	DomainIds []int64 `json:"domainIds"`

	// Specifies the mailbox Ids which contains the emails/folders.
	MailboxIds []int64 `json:"mailboxIds"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *O365HeliosSearchEmailsRequestParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GlobalClusterIdentifier
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GlobalClusterIdentifier = aO0

	// AO1
	var dataAO1 struct {
		DomainIds []int64 `json:"domainIds"`

		MailboxIds []int64 `json:"mailboxIds"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.DomainIds = dataAO1.DomainIds

	m.MailboxIds = dataAO1.MailboxIds

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m O365HeliosSearchEmailsRequestParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.GlobalClusterIdentifier)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		DomainIds []int64 `json:"domainIds"`

		MailboxIds []int64 `json:"mailboxIds"`
	}

	dataAO1.DomainIds = m.DomainIds

	dataAO1.MailboxIds = m.MailboxIds

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this o365 helios search emails request params
func (m *O365HeliosSearchEmailsRequestParams) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GlobalClusterIdentifier
	if err := m.GlobalClusterIdentifier.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this o365 helios search emails request params based on the context it is used
func (m *O365HeliosSearchEmailsRequestParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GlobalClusterIdentifier
	if err := m.GlobalClusterIdentifier.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *O365HeliosSearchEmailsRequestParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *O365HeliosSearchEmailsRequestParams) UnmarshalBinary(b []byte) error {
	var res O365HeliosSearchEmailsRequestParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

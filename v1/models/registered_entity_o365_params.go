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

// RegisteredEntityO365Params Contains all params specified by the user while registering O365 entity.
//
// swagger:model RegisteredEntityO365Params
type RegisteredEntityO365Params struct {

	// Specifies the array of container entities which are to be discovered
	// within the O365 domain. Currently the discovery of Users, Sites, Groups,
	// Teams & Public Folders are supported. If this is not specified then all
	// entities will be discovered.
	//
	// Example: For discovery of only kUser & kSite entities, this vec should be
	// set to [kUsers, kSites].
	DiscoverableEntityTypeVec []int32 `json:"discoverableEntityTypeVec"`

	// If set to true, sharepoint sites will be tagged with group site or teams
	// channel site.
	EnableSiteTagging *bool `json:"enableSiteTagging,omitempty"`

	// Keeps track of background entity fetch on the o365 source.
	// Note: Applicable to new source registrations.
	PendingBackgroundEntityFetch *bool `json:"pendingBackgroundEntityFetch,omitempty"`

	// team entity additional params
	TeamEntityAdditionalParams *RegisteredEntityO365ParamsTeamEntityAdditionalParams `json:"teamEntityAdditionalParams,omitempty"`

	// If some credentials are not updated, we will use the previous existing
	// value for them thus avoiding specifying every secret credential in case
	// of update.
	UseExistingCredentials *bool `json:"useExistingCredentials,omitempty"`

	// user entity discovery params
	UserEntityDiscoveryParams *RegisteredEntityO365ParamsUserEntityDiscoveryParams `json:"userEntityDiscoveryParams,omitempty"`
}

// Validate validates this registered entity o365 params
func (m *RegisteredEntityO365Params) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTeamEntityAdditionalParams(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserEntityDiscoveryParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredEntityO365Params) validateTeamEntityAdditionalParams(formats strfmt.Registry) error {
	if swag.IsZero(m.TeamEntityAdditionalParams) { // not required
		return nil
	}

	if m.TeamEntityAdditionalParams != nil {
		if err := m.TeamEntityAdditionalParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamEntityAdditionalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamEntityAdditionalParams")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredEntityO365Params) validateUserEntityDiscoveryParams(formats strfmt.Registry) error {
	if swag.IsZero(m.UserEntityDiscoveryParams) { // not required
		return nil
	}

	if m.UserEntityDiscoveryParams != nil {
		if err := m.UserEntityDiscoveryParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userEntityDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userEntityDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registered entity o365 params based on the context it is used
func (m *RegisteredEntityO365Params) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTeamEntityAdditionalParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUserEntityDiscoveryParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegisteredEntityO365Params) contextValidateTeamEntityAdditionalParams(ctx context.Context, formats strfmt.Registry) error {

	if m.TeamEntityAdditionalParams != nil {

		if swag.IsZero(m.TeamEntityAdditionalParams) { // not required
			return nil
		}

		if err := m.TeamEntityAdditionalParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("teamEntityAdditionalParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("teamEntityAdditionalParams")
			}
			return err
		}
	}

	return nil
}

func (m *RegisteredEntityO365Params) contextValidateUserEntityDiscoveryParams(ctx context.Context, formats strfmt.Registry) error {

	if m.UserEntityDiscoveryParams != nil {

		if swag.IsZero(m.UserEntityDiscoveryParams) { // not required
			return nil
		}

		if err := m.UserEntityDiscoveryParams.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userEntityDiscoveryParams")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("userEntityDiscoveryParams")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegisteredEntityO365Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisteredEntityO365Params) UnmarshalBinary(b []byte) error {
	var res RegisteredEntityO365Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// IndexingCloudConfig IndexingCloudConfig
//
// Config required for indexing in DMaaS.
//
// swagger:model IndexingCloudConfig
type IndexingCloudConfig struct {

	// Tenant ID to which this config belongs.
	// Required: true
	TenantID *string `json:"tenantId"`

	// Name of the cloud region.
	// Required: true
	Region *string `json:"region"`

	// AWS Elasticsearch config.
	EsConfig *ESConfigForIndexing `json:"esConfig,omitempty"`

	// Azure Elasticsearch config.
	AzureEsConfig *AzureESConfigForIndexing `json:"azureEsConfig,omitempty"`

	// S3 config.
	S3Config *S3ConfigForIndexing `json:"s3Config,omitempty"`

	// Azure storage config.
	AzureStorageConfig *AzureStorageConfigForIndexing `json:"azureStorageConfig,omitempty"`

	// Whether this tenant is being migrated to this cluster or freshly onboarded.
	IsMigratedTenant *bool `json:"isMigratedTenant,omitempty"`
}

// Validate validates this indexing cloud config
func (m *IndexingCloudConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureEsConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3Config(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureStorageConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexingCloudConfig) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

func (m *IndexingCloudConfig) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *IndexingCloudConfig) validateEsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EsConfig) { // not required
		return nil
	}

	if m.EsConfig != nil {
		if err := m.EsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("esConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("esConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IndexingCloudConfig) validateAzureEsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureEsConfig) { // not required
		return nil
	}

	if m.AzureEsConfig != nil {
		if err := m.AzureEsConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureEsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureEsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IndexingCloudConfig) validateS3Config(formats strfmt.Registry) error {
	if swag.IsZero(m.S3Config) { // not required
		return nil
	}

	if m.S3Config != nil {
		if err := m.S3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3Config")
			}
			return err
		}
	}

	return nil
}

func (m *IndexingCloudConfig) validateAzureStorageConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureStorageConfig) { // not required
		return nil
	}

	if m.AzureStorageConfig != nil {
		if err := m.AzureStorageConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureStorageConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureStorageConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this indexing cloud config based on the context it is used
func (m *IndexingCloudConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureEsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS3Config(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureStorageConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IndexingCloudConfig) contextValidateEsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.EsConfig != nil {

		if swag.IsZero(m.EsConfig) { // not required
			return nil
		}

		if err := m.EsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("esConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("esConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IndexingCloudConfig) contextValidateAzureEsConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureEsConfig != nil {

		if swag.IsZero(m.AzureEsConfig) { // not required
			return nil
		}

		if err := m.AzureEsConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureEsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureEsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *IndexingCloudConfig) contextValidateS3Config(ctx context.Context, formats strfmt.Registry) error {

	if m.S3Config != nil {

		if swag.IsZero(m.S3Config) { // not required
			return nil
		}

		if err := m.S3Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3Config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s3Config")
			}
			return err
		}
	}

	return nil
}

func (m *IndexingCloudConfig) contextValidateAzureStorageConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureStorageConfig != nil {

		if swag.IsZero(m.AzureStorageConfig) { // not required
			return nil
		}

		if err := m.AzureStorageConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureStorageConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureStorageConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IndexingCloudConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IndexingCloudConfig) UnmarshalBinary(b []byte) error {
	var res IndexingCloudConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

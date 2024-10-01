// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VaultProviderStatsInfo Vault Provider Stats Info.
//
// Specifies the stats for each vault.
//
// swagger:model VaultProviderStatsInfo
type VaultProviderStatsInfo struct {

	// Specifies the relative change of size of entities on the vault.
	ChangeRate *int64 `json:"changeRate,omitempty"`

	// Specifies the cluster id.
	ClusterID *int64 `json:"clusterId,omitempty"`

	// Specifies the cluster incarnation id.
	ClusterIncarnationID *int64 `json:"clusterIncarnationId,omitempty"`

	// Specifies the cluster name.
	ClusterName *string `json:"clusterName,omitempty"`

	// Specifies the average read bandwidth over the last 24 hours.
	ReadBandwidth *int64 `json:"readBandwidth,omitempty"`

	// Specifies the stats by environments.
	StatsByEnv []*VaultProviderStatsByEnv `json:"statsByEnv"`

	// Specifies the cloud vendor type.
	// Enum: ["kAws","kAzure","kGcp","kOracle","kNas","kQStar","kS3C","kOther"]
	VaultGroup *string `json:"vaultGroup,omitempty"`

	// Specifies the Vault id.
	VaultID *int64 `json:"vaultId,omitempty"`

	// Specifies the External Target type.
	// Enum: ["kNearline","kGlacier","kS3","kAzureStandard","kS3Compatible","kQStarTape","kGoogleStandard","kGoogleDRA","kAmazonS3StandardIA","kAWSGovCloud","kNAS","kColdline","kAzureGovCloud","kAzureArchive","kAzure","kGoogle","kAmazon","kOracle","kOracleTierStandard","kOracleTierArchive","kAmazonC2S"]
	VaultType *string `json:"vaultType,omitempty"`

	// Specifies the Vault name.
	Vaultname *string `json:"vaultname,omitempty"`

	// Specifies the average write bandwidth over the last 24 hours.
	WriteBandwidth *int64 `json:"writeBandwidth,omitempty"`
}

// Validate validates this vault provider stats info
func (m *VaultProviderStatsInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatsByEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVaultType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaultProviderStatsInfo) validateStatsByEnv(formats strfmt.Registry) error {
	if swag.IsZero(m.StatsByEnv) { // not required
		return nil
	}

	for i := 0; i < len(m.StatsByEnv); i++ {
		if swag.IsZero(m.StatsByEnv[i]) { // not required
			continue
		}

		if m.StatsByEnv[i] != nil {
			if err := m.StatsByEnv[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statsByEnv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statsByEnv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var vaultProviderStatsInfoTypeVaultGroupPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAws","kAzure","kGcp","kOracle","kNas","kQStar","kS3C","kOther"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vaultProviderStatsInfoTypeVaultGroupPropEnum = append(vaultProviderStatsInfoTypeVaultGroupPropEnum, v)
	}
}

const (

	// VaultProviderStatsInfoVaultGroupKAws captures enum value "kAws"
	VaultProviderStatsInfoVaultGroupKAws string = "kAws"

	// VaultProviderStatsInfoVaultGroupKAzure captures enum value "kAzure"
	VaultProviderStatsInfoVaultGroupKAzure string = "kAzure"

	// VaultProviderStatsInfoVaultGroupKGcp captures enum value "kGcp"
	VaultProviderStatsInfoVaultGroupKGcp string = "kGcp"

	// VaultProviderStatsInfoVaultGroupKOracle captures enum value "kOracle"
	VaultProviderStatsInfoVaultGroupKOracle string = "kOracle"

	// VaultProviderStatsInfoVaultGroupKNas captures enum value "kNas"
	VaultProviderStatsInfoVaultGroupKNas string = "kNas"

	// VaultProviderStatsInfoVaultGroupKQStar captures enum value "kQStar"
	VaultProviderStatsInfoVaultGroupKQStar string = "kQStar"

	// VaultProviderStatsInfoVaultGroupKS3C captures enum value "kS3C"
	VaultProviderStatsInfoVaultGroupKS3C string = "kS3C"

	// VaultProviderStatsInfoVaultGroupKOther captures enum value "kOther"
	VaultProviderStatsInfoVaultGroupKOther string = "kOther"
)

// prop value enum
func (m *VaultProviderStatsInfo) validateVaultGroupEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vaultProviderStatsInfoTypeVaultGroupPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VaultProviderStatsInfo) validateVaultGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.VaultGroup) { // not required
		return nil
	}

	// value enum
	if err := m.validateVaultGroupEnum("vaultGroup", "body", *m.VaultGroup); err != nil {
		return err
	}

	return nil
}

var vaultProviderStatsInfoTypeVaultTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNearline","kGlacier","kS3","kAzureStandard","kS3Compatible","kQStarTape","kGoogleStandard","kGoogleDRA","kAmazonS3StandardIA","kAWSGovCloud","kNAS","kColdline","kAzureGovCloud","kAzureArchive","kAzure","kGoogle","kAmazon","kOracle","kOracleTierStandard","kOracleTierArchive","kAmazonC2S"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vaultProviderStatsInfoTypeVaultTypePropEnum = append(vaultProviderStatsInfoTypeVaultTypePropEnum, v)
	}
}

const (

	// VaultProviderStatsInfoVaultTypeKNearline captures enum value "kNearline"
	VaultProviderStatsInfoVaultTypeKNearline string = "kNearline"

	// VaultProviderStatsInfoVaultTypeKGlacier captures enum value "kGlacier"
	VaultProviderStatsInfoVaultTypeKGlacier string = "kGlacier"

	// VaultProviderStatsInfoVaultTypeKS3 captures enum value "kS3"
	VaultProviderStatsInfoVaultTypeKS3 string = "kS3"

	// VaultProviderStatsInfoVaultTypeKAzureStandard captures enum value "kAzureStandard"
	VaultProviderStatsInfoVaultTypeKAzureStandard string = "kAzureStandard"

	// VaultProviderStatsInfoVaultTypeKS3Compatible captures enum value "kS3Compatible"
	VaultProviderStatsInfoVaultTypeKS3Compatible string = "kS3Compatible"

	// VaultProviderStatsInfoVaultTypeKQStarTape captures enum value "kQStarTape"
	VaultProviderStatsInfoVaultTypeKQStarTape string = "kQStarTape"

	// VaultProviderStatsInfoVaultTypeKGoogleStandard captures enum value "kGoogleStandard"
	VaultProviderStatsInfoVaultTypeKGoogleStandard string = "kGoogleStandard"

	// VaultProviderStatsInfoVaultTypeKGoogleDRA captures enum value "kGoogleDRA"
	VaultProviderStatsInfoVaultTypeKGoogleDRA string = "kGoogleDRA"

	// VaultProviderStatsInfoVaultTypeKAmazonS3StandardIA captures enum value "kAmazonS3StandardIA"
	VaultProviderStatsInfoVaultTypeKAmazonS3StandardIA string = "kAmazonS3StandardIA"

	// VaultProviderStatsInfoVaultTypeKAWSGovCloud captures enum value "kAWSGovCloud"
	VaultProviderStatsInfoVaultTypeKAWSGovCloud string = "kAWSGovCloud"

	// VaultProviderStatsInfoVaultTypeKNAS captures enum value "kNAS"
	VaultProviderStatsInfoVaultTypeKNAS string = "kNAS"

	// VaultProviderStatsInfoVaultTypeKColdline captures enum value "kColdline"
	VaultProviderStatsInfoVaultTypeKColdline string = "kColdline"

	// VaultProviderStatsInfoVaultTypeKAzureGovCloud captures enum value "kAzureGovCloud"
	VaultProviderStatsInfoVaultTypeKAzureGovCloud string = "kAzureGovCloud"

	// VaultProviderStatsInfoVaultTypeKAzureArchive captures enum value "kAzureArchive"
	VaultProviderStatsInfoVaultTypeKAzureArchive string = "kAzureArchive"

	// VaultProviderStatsInfoVaultTypeKAzure captures enum value "kAzure"
	VaultProviderStatsInfoVaultTypeKAzure string = "kAzure"

	// VaultProviderStatsInfoVaultTypeKGoogle captures enum value "kGoogle"
	VaultProviderStatsInfoVaultTypeKGoogle string = "kGoogle"

	// VaultProviderStatsInfoVaultTypeKAmazon captures enum value "kAmazon"
	VaultProviderStatsInfoVaultTypeKAmazon string = "kAmazon"

	// VaultProviderStatsInfoVaultTypeKOracle captures enum value "kOracle"
	VaultProviderStatsInfoVaultTypeKOracle string = "kOracle"

	// VaultProviderStatsInfoVaultTypeKOracleTierStandard captures enum value "kOracleTierStandard"
	VaultProviderStatsInfoVaultTypeKOracleTierStandard string = "kOracleTierStandard"

	// VaultProviderStatsInfoVaultTypeKOracleTierArchive captures enum value "kOracleTierArchive"
	VaultProviderStatsInfoVaultTypeKOracleTierArchive string = "kOracleTierArchive"

	// VaultProviderStatsInfoVaultTypeKAmazonC2S captures enum value "kAmazonC2S"
	VaultProviderStatsInfoVaultTypeKAmazonC2S string = "kAmazonC2S"
)

// prop value enum
func (m *VaultProviderStatsInfo) validateVaultTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vaultProviderStatsInfoTypeVaultTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VaultProviderStatsInfo) validateVaultType(formats strfmt.Registry) error {
	if swag.IsZero(m.VaultType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVaultTypeEnum("vaultType", "body", *m.VaultType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this vault provider stats info based on the context it is used
func (m *VaultProviderStatsInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatsByEnv(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VaultProviderStatsInfo) contextValidateStatsByEnv(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StatsByEnv); i++ {

		if m.StatsByEnv[i] != nil {

			if swag.IsZero(m.StatsByEnv[i]) { // not required
				return nil
			}

			if err := m.StatsByEnv[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statsByEnv" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statsByEnv" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VaultProviderStatsInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VaultProviderStatsInfo) UnmarshalBinary(b []byte) error {
	var res VaultProviderStatsInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

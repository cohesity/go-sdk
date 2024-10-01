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

// AzureSQLSkuOptions Azure SQL SKU Options.
//
// Specifies the SQL SKU parameters which are specific to Azure related Object Protection & Recovery.
//
// swagger:model AzureSqlSkuOptions
type AzureSQLSkuOptions struct {

	// Specifies the sku name for azure sql databases and by default Hyperscale is selected.
	// Enum: ["Basic","Standard","Premium","GeneralPurpose","BusinessCritical","Hyperscale","DataWarehouse","Stretch"]
	Name *string `json:"name,omitempty"`

	// Specifies the sku tier selection for azure sql databases and by default HS_Gen5 is selected. The tiers must match their sku name selected.
	// Enum: ["S0","S1","S2","S3","S4","S6","S7","S9","S12","P1","P2","P4","P6","P11","P15","BC_DC","BC_Gen5","BC_M","GP_DC","GP_Fsv2","GP_Gen5","GP_S_Gen5","HS_DC","HS_Gen5","HS_S_Gen5","HS_MOPRMS","HS_PRMS","Free","Basic","Standard","Premium","DataWarehouse","Stretch"]
	TierType *string `json:"tierType,omitempty"`

	// Specifies the capacity of the sku. For azure sql dbs, this is the number of cores. Default value is 4.
	Capacity *int32 `json:"capacity,omitempty"`
}

// Validate validates this azure Sql sku options
func (m *AzureSQLSkuOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTierType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var azureSqlSkuOptionsTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Basic","Standard","Premium","GeneralPurpose","BusinessCritical","Hyperscale","DataWarehouse","Stretch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureSqlSkuOptionsTypeNamePropEnum = append(azureSqlSkuOptionsTypeNamePropEnum, v)
	}
}

const (

	// AzureSQLSkuOptionsNameBasic captures enum value "Basic"
	AzureSQLSkuOptionsNameBasic string = "Basic"

	// AzureSQLSkuOptionsNameStandard captures enum value "Standard"
	AzureSQLSkuOptionsNameStandard string = "Standard"

	// AzureSQLSkuOptionsNamePremium captures enum value "Premium"
	AzureSQLSkuOptionsNamePremium string = "Premium"

	// AzureSQLSkuOptionsNameGeneralPurpose captures enum value "GeneralPurpose"
	AzureSQLSkuOptionsNameGeneralPurpose string = "GeneralPurpose"

	// AzureSQLSkuOptionsNameBusinessCritical captures enum value "BusinessCritical"
	AzureSQLSkuOptionsNameBusinessCritical string = "BusinessCritical"

	// AzureSQLSkuOptionsNameHyperscale captures enum value "Hyperscale"
	AzureSQLSkuOptionsNameHyperscale string = "Hyperscale"

	// AzureSQLSkuOptionsNameDataWarehouse captures enum value "DataWarehouse"
	AzureSQLSkuOptionsNameDataWarehouse string = "DataWarehouse"

	// AzureSQLSkuOptionsNameStretch captures enum value "Stretch"
	AzureSQLSkuOptionsNameStretch string = "Stretch"
)

// prop value enum
func (m *AzureSQLSkuOptions) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureSqlSkuOptionsTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureSQLSkuOptions) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", *m.Name); err != nil {
		return err
	}

	return nil
}

var azureSqlSkuOptionsTypeTierTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["S0","S1","S2","S3","S4","S6","S7","S9","S12","P1","P2","P4","P6","P11","P15","BC_DC","BC_Gen5","BC_M","GP_DC","GP_Fsv2","GP_Gen5","GP_S_Gen5","HS_DC","HS_Gen5","HS_S_Gen5","HS_MOPRMS","HS_PRMS","Free","Basic","Standard","Premium","DataWarehouse","Stretch"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		azureSqlSkuOptionsTypeTierTypePropEnum = append(azureSqlSkuOptionsTypeTierTypePropEnum, v)
	}
}

const (

	// AzureSQLSkuOptionsTierTypeS0 captures enum value "S0"
	AzureSQLSkuOptionsTierTypeS0 string = "S0"

	// AzureSQLSkuOptionsTierTypeS1 captures enum value "S1"
	AzureSQLSkuOptionsTierTypeS1 string = "S1"

	// AzureSQLSkuOptionsTierTypeS2 captures enum value "S2"
	AzureSQLSkuOptionsTierTypeS2 string = "S2"

	// AzureSQLSkuOptionsTierTypeS3 captures enum value "S3"
	AzureSQLSkuOptionsTierTypeS3 string = "S3"

	// AzureSQLSkuOptionsTierTypeS4 captures enum value "S4"
	AzureSQLSkuOptionsTierTypeS4 string = "S4"

	// AzureSQLSkuOptionsTierTypeS6 captures enum value "S6"
	AzureSQLSkuOptionsTierTypeS6 string = "S6"

	// AzureSQLSkuOptionsTierTypeS7 captures enum value "S7"
	AzureSQLSkuOptionsTierTypeS7 string = "S7"

	// AzureSQLSkuOptionsTierTypeS9 captures enum value "S9"
	AzureSQLSkuOptionsTierTypeS9 string = "S9"

	// AzureSQLSkuOptionsTierTypeS12 captures enum value "S12"
	AzureSQLSkuOptionsTierTypeS12 string = "S12"

	// AzureSQLSkuOptionsTierTypeP1 captures enum value "P1"
	AzureSQLSkuOptionsTierTypeP1 string = "P1"

	// AzureSQLSkuOptionsTierTypeP2 captures enum value "P2"
	AzureSQLSkuOptionsTierTypeP2 string = "P2"

	// AzureSQLSkuOptionsTierTypeP4 captures enum value "P4"
	AzureSQLSkuOptionsTierTypeP4 string = "P4"

	// AzureSQLSkuOptionsTierTypeP6 captures enum value "P6"
	AzureSQLSkuOptionsTierTypeP6 string = "P6"

	// AzureSQLSkuOptionsTierTypeP11 captures enum value "P11"
	AzureSQLSkuOptionsTierTypeP11 string = "P11"

	// AzureSQLSkuOptionsTierTypeP15 captures enum value "P15"
	AzureSQLSkuOptionsTierTypeP15 string = "P15"

	// AzureSQLSkuOptionsTierTypeBCDC captures enum value "BC_DC"
	AzureSQLSkuOptionsTierTypeBCDC string = "BC_DC"

	// AzureSQLSkuOptionsTierTypeBCGen5 captures enum value "BC_Gen5"
	AzureSQLSkuOptionsTierTypeBCGen5 string = "BC_Gen5"

	// AzureSQLSkuOptionsTierTypeBCM captures enum value "BC_M"
	AzureSQLSkuOptionsTierTypeBCM string = "BC_M"

	// AzureSQLSkuOptionsTierTypeGPDC captures enum value "GP_DC"
	AzureSQLSkuOptionsTierTypeGPDC string = "GP_DC"

	// AzureSQLSkuOptionsTierTypeGPFsv2 captures enum value "GP_Fsv2"
	AzureSQLSkuOptionsTierTypeGPFsv2 string = "GP_Fsv2"

	// AzureSQLSkuOptionsTierTypeGPGen5 captures enum value "GP_Gen5"
	AzureSQLSkuOptionsTierTypeGPGen5 string = "GP_Gen5"

	// AzureSQLSkuOptionsTierTypeGPSGen5 captures enum value "GP_S_Gen5"
	AzureSQLSkuOptionsTierTypeGPSGen5 string = "GP_S_Gen5"

	// AzureSQLSkuOptionsTierTypeHSDC captures enum value "HS_DC"
	AzureSQLSkuOptionsTierTypeHSDC string = "HS_DC"

	// AzureSQLSkuOptionsTierTypeHSGen5 captures enum value "HS_Gen5"
	AzureSQLSkuOptionsTierTypeHSGen5 string = "HS_Gen5"

	// AzureSQLSkuOptionsTierTypeHSSGen5 captures enum value "HS_S_Gen5"
	AzureSQLSkuOptionsTierTypeHSSGen5 string = "HS_S_Gen5"

	// AzureSQLSkuOptionsTierTypeHSMOPRMS captures enum value "HS_MOPRMS"
	AzureSQLSkuOptionsTierTypeHSMOPRMS string = "HS_MOPRMS"

	// AzureSQLSkuOptionsTierTypeHSPRMS captures enum value "HS_PRMS"
	AzureSQLSkuOptionsTierTypeHSPRMS string = "HS_PRMS"

	// AzureSQLSkuOptionsTierTypeFree captures enum value "Free"
	AzureSQLSkuOptionsTierTypeFree string = "Free"

	// AzureSQLSkuOptionsTierTypeBasic captures enum value "Basic"
	AzureSQLSkuOptionsTierTypeBasic string = "Basic"

	// AzureSQLSkuOptionsTierTypeStandard captures enum value "Standard"
	AzureSQLSkuOptionsTierTypeStandard string = "Standard"

	// AzureSQLSkuOptionsTierTypePremium captures enum value "Premium"
	AzureSQLSkuOptionsTierTypePremium string = "Premium"

	// AzureSQLSkuOptionsTierTypeDataWarehouse captures enum value "DataWarehouse"
	AzureSQLSkuOptionsTierTypeDataWarehouse string = "DataWarehouse"

	// AzureSQLSkuOptionsTierTypeStretch captures enum value "Stretch"
	AzureSQLSkuOptionsTierTypeStretch string = "Stretch"
)

// prop value enum
func (m *AzureSQLSkuOptions) validateTierTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, azureSqlSkuOptionsTypeTierTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AzureSQLSkuOptions) validateTierType(formats strfmt.Registry) error {
	if swag.IsZero(m.TierType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTierTypeEnum("tierType", "body", *m.TierType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this azure Sql sku options based on context it is used
func (m *AzureSQLSkuOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureSQLSkuOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureSQLSkuOptions) UnmarshalBinary(b []byte) error {
	var res AzureSQLSkuOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

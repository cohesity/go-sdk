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

// SearchFileRequestParamsBase Search files request params.
//
// Specifies the request parameters to search for files and file folders.
//
// swagger:model SearchFileRequestParamsBase
type SearchFileRequestParamsBase struct {

	// Specifies the search string to filter the files. User can specify a wildcard character '*' as a suffix to a string where all files name are matched with the prefix string.
	SearchString *string `json:"searchString,omitempty"`

	// Specifies a list of file types. Only files within the given types will be returned.
	// Unique: true
	Types []string `json:"types"`

	// Specifies a list of the source environments. Only files from these types of source will be returned.
	SourceEnvironments []string `json:"sourceEnvironments"`
}

// Validate validates this search file request params base
func (m *SearchFileRequestParamsBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceEnvironments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var searchFileRequestParamsBaseTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["File","Directory","Symlink"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchFileRequestParamsBaseTypesItemsEnum = append(searchFileRequestParamsBaseTypesItemsEnum, v)
	}
}

func (m *SearchFileRequestParamsBase) validateTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchFileRequestParamsBaseTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchFileRequestParamsBase) validateTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.Types) { // not required
		return nil
	}

	if err := validate.UniqueItems("types", "body", m.Types); err != nil {
		return err
	}

	for i := 0; i < len(m.Types); i++ {

		// value enum
		if err := m.validateTypesItemsEnum("types"+"."+strconv.Itoa(i), "body", m.Types[i]); err != nil {
			return err
		}

	}

	return nil
}

var searchFileRequestParamsBaseSourceEnvironmentsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kSQL","kView","kRemoteAdapter","kPhysical","kPhysicalFiles","kPure","kIbmFlashSystem","kAzure","kNetapp","kGenericNas","kAcropolis","kIsilon","kGPFS","kKVM","kAWS","kExchange","kOracle","kGCP","kFlashBlade","kO365","kHyperFlex","kKubernetes","kElastifile","kUDA","kSfdc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		searchFileRequestParamsBaseSourceEnvironmentsItemsEnum = append(searchFileRequestParamsBaseSourceEnvironmentsItemsEnum, v)
	}
}

func (m *SearchFileRequestParamsBase) validateSourceEnvironmentsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, searchFileRequestParamsBaseSourceEnvironmentsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SearchFileRequestParamsBase) validateSourceEnvironments(formats strfmt.Registry) error {
	if swag.IsZero(m.SourceEnvironments) { // not required
		return nil
	}

	for i := 0; i < len(m.SourceEnvironments); i++ {

		// value enum
		if err := m.validateSourceEnvironmentsItemsEnum("sourceEnvironments"+"."+strconv.Itoa(i), "body", m.SourceEnvironments[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this search file request params base based on context it is used
func (m *SearchFileRequestParamsBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchFileRequestParamsBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchFileRequestParamsBase) UnmarshalBinary(b []byte) error {
	var res SearchFileRequestParamsBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

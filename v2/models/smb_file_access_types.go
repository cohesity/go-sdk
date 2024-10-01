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

// SmbFileAccessTypes SMB File Access types
//
// Specifies the SMB file access types.
//
// swagger:model SmbFileAccessTypes
type SmbFileAccessTypes struct {

	// Specifies the list of SMB file access types.
	// Enum: ["FileReadData","FileWriteData","FileAppendData","FileReadEa","FileWriteEa","FileExecute","FileDeleteChild","FileReadAttributes","FileWriteAttributes","Delete","ReadControl","WriteDac","WriteOwner","Synchronize","AccessSystemSecurity","MaximumAllowed","GenericAll","GenericExecute","GenericWrite","GenericRead"]
	Value string `json:"value,omitempty"`
}

// Validate validates this smb file access types
func (m *SmbFileAccessTypes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var smbFileAccessTypesTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FileReadData","FileWriteData","FileAppendData","FileReadEa","FileWriteEa","FileExecute","FileDeleteChild","FileReadAttributes","FileWriteAttributes","Delete","ReadControl","WriteDac","WriteOwner","Synchronize","AccessSystemSecurity","MaximumAllowed","GenericAll","GenericExecute","GenericWrite","GenericRead"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smbFileAccessTypesTypeValuePropEnum = append(smbFileAccessTypesTypeValuePropEnum, v)
	}
}

const (

	// SmbFileAccessTypesValueFileReadData captures enum value "FileReadData"
	SmbFileAccessTypesValueFileReadData string = "FileReadData"

	// SmbFileAccessTypesValueFileWriteData captures enum value "FileWriteData"
	SmbFileAccessTypesValueFileWriteData string = "FileWriteData"

	// SmbFileAccessTypesValueFileAppendData captures enum value "FileAppendData"
	SmbFileAccessTypesValueFileAppendData string = "FileAppendData"

	// SmbFileAccessTypesValueFileReadEa captures enum value "FileReadEa"
	SmbFileAccessTypesValueFileReadEa string = "FileReadEa"

	// SmbFileAccessTypesValueFileWriteEa captures enum value "FileWriteEa"
	SmbFileAccessTypesValueFileWriteEa string = "FileWriteEa"

	// SmbFileAccessTypesValueFileExecute captures enum value "FileExecute"
	SmbFileAccessTypesValueFileExecute string = "FileExecute"

	// SmbFileAccessTypesValueFileDeleteChild captures enum value "FileDeleteChild"
	SmbFileAccessTypesValueFileDeleteChild string = "FileDeleteChild"

	// SmbFileAccessTypesValueFileReadAttributes captures enum value "FileReadAttributes"
	SmbFileAccessTypesValueFileReadAttributes string = "FileReadAttributes"

	// SmbFileAccessTypesValueFileWriteAttributes captures enum value "FileWriteAttributes"
	SmbFileAccessTypesValueFileWriteAttributes string = "FileWriteAttributes"

	// SmbFileAccessTypesValueDelete captures enum value "Delete"
	SmbFileAccessTypesValueDelete string = "Delete"

	// SmbFileAccessTypesValueReadControl captures enum value "ReadControl"
	SmbFileAccessTypesValueReadControl string = "ReadControl"

	// SmbFileAccessTypesValueWriteDac captures enum value "WriteDac"
	SmbFileAccessTypesValueWriteDac string = "WriteDac"

	// SmbFileAccessTypesValueWriteOwner captures enum value "WriteOwner"
	SmbFileAccessTypesValueWriteOwner string = "WriteOwner"

	// SmbFileAccessTypesValueSynchronize captures enum value "Synchronize"
	SmbFileAccessTypesValueSynchronize string = "Synchronize"

	// SmbFileAccessTypesValueAccessSystemSecurity captures enum value "AccessSystemSecurity"
	SmbFileAccessTypesValueAccessSystemSecurity string = "AccessSystemSecurity"

	// SmbFileAccessTypesValueMaximumAllowed captures enum value "MaximumAllowed"
	SmbFileAccessTypesValueMaximumAllowed string = "MaximumAllowed"

	// SmbFileAccessTypesValueGenericAll captures enum value "GenericAll"
	SmbFileAccessTypesValueGenericAll string = "GenericAll"

	// SmbFileAccessTypesValueGenericExecute captures enum value "GenericExecute"
	SmbFileAccessTypesValueGenericExecute string = "GenericExecute"

	// SmbFileAccessTypesValueGenericWrite captures enum value "GenericWrite"
	SmbFileAccessTypesValueGenericWrite string = "GenericWrite"

	// SmbFileAccessTypesValueGenericRead captures enum value "GenericRead"
	SmbFileAccessTypesValueGenericRead string = "GenericRead"
)

// prop value enum
func (m *SmbFileAccessTypes) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smbFileAccessTypesTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmbFileAccessTypes) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	// value enum
	if err := m.validateValueEnum("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this smb file access types based on context it is used
func (m *SmbFileAccessTypes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SmbFileAccessTypes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmbFileAccessTypes) UnmarshalBinary(b []byte) error {
	var res SmbFileAccessTypes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

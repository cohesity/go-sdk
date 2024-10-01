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

// DownloadChatsParams Specifies the Download chat/posts specific parameters.
//
// swagger:model DownloadChatsParams
type DownloadChatsParams struct {

	// Specifies the html template for the downloaded chats.
	HTMLTemplate *string `json:"htmlTemplate,omitempty"`

	// Specifies channel IDs whose posts needs to be downloaded. If channelIds is nil or empty then full teams' posts will be downloaded.
	ChannelIds []string `json:"channelIds"`

	// Specifies the file type for the downloaded content.
	// Required: true
	// Enum: ["kHTML","kJSON"]
	DownloadFileType *string `json:"downloadFileType"`
}

// Validate validates this download chats params
func (m *DownloadChatsParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownloadFileType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var downloadChatsParamsTypeDownloadFileTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kHTML","kJSON"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		downloadChatsParamsTypeDownloadFileTypePropEnum = append(downloadChatsParamsTypeDownloadFileTypePropEnum, v)
	}
}

const (

	// DownloadChatsParamsDownloadFileTypeKHTML captures enum value "kHTML"
	DownloadChatsParamsDownloadFileTypeKHTML string = "kHTML"

	// DownloadChatsParamsDownloadFileTypeKJSON captures enum value "kJSON"
	DownloadChatsParamsDownloadFileTypeKJSON string = "kJSON"
)

// prop value enum
func (m *DownloadChatsParams) validateDownloadFileTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, downloadChatsParamsTypeDownloadFileTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DownloadChatsParams) validateDownloadFileType(formats strfmt.Registry) error {

	if err := validate.Required("downloadFileType", "body", m.DownloadFileType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDownloadFileTypeEnum("downloadFileType", "body", *m.DownloadFileType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this download chats params based on context it is used
func (m *DownloadChatsParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DownloadChatsParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadChatsParams) UnmarshalBinary(b []byte) error {
	var res DownloadChatsParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

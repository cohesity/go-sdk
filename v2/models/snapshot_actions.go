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

// SnapshotActions Snapshot Actions
//
// # Snapshot Actions
//
// swagger:model SnapshotActions
type SnapshotActions struct {

	// Snapshot Actions
	// Enum: ["RecoverVMs","RecoverFiles","InstantVolumeMount","RecoverVmDisks","MountVolumes","RecoverVApps","RecoverRDS","RecoverAurora","RecoverS3Buckets","RecoverApps","RecoverNasVolume","RecoverPhysicalVolumes","RecoverSystem","RecoverSanVolumes","RecoverNamespaces","RecoverObjects","DownloadFilesAndFolders","RecoverPublicFolders","RecoverVAppTemplates","RecoverMailbox","RecoverOneDrive","RecoverMsTeam","RecoverMsGroup","RecoverSharePoint","ConvertToPst","RecoverSfdcRecords","RecoverAzureSQL","DownloadChats","RecoverRDSPostgres"]
	SnapshotActions string `json:"snapshotActions,omitempty"`
}

// Validate validates this snapshot actions
func (m *SnapshotActions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSnapshotActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var snapshotActionsTypeSnapshotActionsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverVMs","RecoverFiles","InstantVolumeMount","RecoverVmDisks","MountVolumes","RecoverVApps","RecoverRDS","RecoverAurora","RecoverS3Buckets","RecoverApps","RecoverNasVolume","RecoverPhysicalVolumes","RecoverSystem","RecoverSanVolumes","RecoverNamespaces","RecoverObjects","DownloadFilesAndFolders","RecoverPublicFolders","RecoverVAppTemplates","RecoverMailbox","RecoverOneDrive","RecoverMsTeam","RecoverMsGroup","RecoverSharePoint","ConvertToPst","RecoverSfdcRecords","RecoverAzureSQL","DownloadChats","RecoverRDSPostgres"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		snapshotActionsTypeSnapshotActionsPropEnum = append(snapshotActionsTypeSnapshotActionsPropEnum, v)
	}
}

const (

	// SnapshotActionsSnapshotActionsRecoverVMs captures enum value "RecoverVMs"
	SnapshotActionsSnapshotActionsRecoverVMs string = "RecoverVMs"

	// SnapshotActionsSnapshotActionsRecoverFiles captures enum value "RecoverFiles"
	SnapshotActionsSnapshotActionsRecoverFiles string = "RecoverFiles"

	// SnapshotActionsSnapshotActionsInstantVolumeMount captures enum value "InstantVolumeMount"
	SnapshotActionsSnapshotActionsInstantVolumeMount string = "InstantVolumeMount"

	// SnapshotActionsSnapshotActionsRecoverVMDisks captures enum value "RecoverVmDisks"
	SnapshotActionsSnapshotActionsRecoverVMDisks string = "RecoverVmDisks"

	// SnapshotActionsSnapshotActionsMountVolumes captures enum value "MountVolumes"
	SnapshotActionsSnapshotActionsMountVolumes string = "MountVolumes"

	// SnapshotActionsSnapshotActionsRecoverVApps captures enum value "RecoverVApps"
	SnapshotActionsSnapshotActionsRecoverVApps string = "RecoverVApps"

	// SnapshotActionsSnapshotActionsRecoverRDS captures enum value "RecoverRDS"
	SnapshotActionsSnapshotActionsRecoverRDS string = "RecoverRDS"

	// SnapshotActionsSnapshotActionsRecoverAurora captures enum value "RecoverAurora"
	SnapshotActionsSnapshotActionsRecoverAurora string = "RecoverAurora"

	// SnapshotActionsSnapshotActionsRecoverS3Buckets captures enum value "RecoverS3Buckets"
	SnapshotActionsSnapshotActionsRecoverS3Buckets string = "RecoverS3Buckets"

	// SnapshotActionsSnapshotActionsRecoverApps captures enum value "RecoverApps"
	SnapshotActionsSnapshotActionsRecoverApps string = "RecoverApps"

	// SnapshotActionsSnapshotActionsRecoverNasVolume captures enum value "RecoverNasVolume"
	SnapshotActionsSnapshotActionsRecoverNasVolume string = "RecoverNasVolume"

	// SnapshotActionsSnapshotActionsRecoverPhysicalVolumes captures enum value "RecoverPhysicalVolumes"
	SnapshotActionsSnapshotActionsRecoverPhysicalVolumes string = "RecoverPhysicalVolumes"

	// SnapshotActionsSnapshotActionsRecoverSystem captures enum value "RecoverSystem"
	SnapshotActionsSnapshotActionsRecoverSystem string = "RecoverSystem"

	// SnapshotActionsSnapshotActionsRecoverSanVolumes captures enum value "RecoverSanVolumes"
	SnapshotActionsSnapshotActionsRecoverSanVolumes string = "RecoverSanVolumes"

	// SnapshotActionsSnapshotActionsRecoverNamespaces captures enum value "RecoverNamespaces"
	SnapshotActionsSnapshotActionsRecoverNamespaces string = "RecoverNamespaces"

	// SnapshotActionsSnapshotActionsRecoverObjects captures enum value "RecoverObjects"
	SnapshotActionsSnapshotActionsRecoverObjects string = "RecoverObjects"

	// SnapshotActionsSnapshotActionsDownloadFilesAndFolders captures enum value "DownloadFilesAndFolders"
	SnapshotActionsSnapshotActionsDownloadFilesAndFolders string = "DownloadFilesAndFolders"

	// SnapshotActionsSnapshotActionsRecoverPublicFolders captures enum value "RecoverPublicFolders"
	SnapshotActionsSnapshotActionsRecoverPublicFolders string = "RecoverPublicFolders"

	// SnapshotActionsSnapshotActionsRecoverVAppTemplates captures enum value "RecoverVAppTemplates"
	SnapshotActionsSnapshotActionsRecoverVAppTemplates string = "RecoverVAppTemplates"

	// SnapshotActionsSnapshotActionsRecoverMailbox captures enum value "RecoverMailbox"
	SnapshotActionsSnapshotActionsRecoverMailbox string = "RecoverMailbox"

	// SnapshotActionsSnapshotActionsRecoverOneDrive captures enum value "RecoverOneDrive"
	SnapshotActionsSnapshotActionsRecoverOneDrive string = "RecoverOneDrive"

	// SnapshotActionsSnapshotActionsRecoverMsTeam captures enum value "RecoverMsTeam"
	SnapshotActionsSnapshotActionsRecoverMsTeam string = "RecoverMsTeam"

	// SnapshotActionsSnapshotActionsRecoverMsGroup captures enum value "RecoverMsGroup"
	SnapshotActionsSnapshotActionsRecoverMsGroup string = "RecoverMsGroup"

	// SnapshotActionsSnapshotActionsRecoverSharePoint captures enum value "RecoverSharePoint"
	SnapshotActionsSnapshotActionsRecoverSharePoint string = "RecoverSharePoint"

	// SnapshotActionsSnapshotActionsConvertToPst captures enum value "ConvertToPst"
	SnapshotActionsSnapshotActionsConvertToPst string = "ConvertToPst"

	// SnapshotActionsSnapshotActionsRecoverSfdcRecords captures enum value "RecoverSfdcRecords"
	SnapshotActionsSnapshotActionsRecoverSfdcRecords string = "RecoverSfdcRecords"

	// SnapshotActionsSnapshotActionsRecoverAzureSQL captures enum value "RecoverAzureSQL"
	SnapshotActionsSnapshotActionsRecoverAzureSQL string = "RecoverAzureSQL"

	// SnapshotActionsSnapshotActionsDownloadChats captures enum value "DownloadChats"
	SnapshotActionsSnapshotActionsDownloadChats string = "DownloadChats"

	// SnapshotActionsSnapshotActionsRecoverRDSPostgres captures enum value "RecoverRDSPostgres"
	SnapshotActionsSnapshotActionsRecoverRDSPostgres string = "RecoverRDSPostgres"
)

// prop value enum
func (m *SnapshotActions) validateSnapshotActionsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, snapshotActionsTypeSnapshotActionsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SnapshotActions) validateSnapshotActions(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotActions) { // not required
		return nil
	}

	// value enum
	if err := m.validateSnapshotActionsEnum("snapshotActions", "body", m.SnapshotActions); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this snapshot actions based on context it is used
func (m *SnapshotActions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotActions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotActions) UnmarshalBinary(b []byte) error {
	var res SnapshotActions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

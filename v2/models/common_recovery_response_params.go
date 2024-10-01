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

// CommonRecoveryResponseParams Common Recovery Response Params.
//
// # Specifies the common response parameters to create a Recovery
//
// swagger:model CommonRecoveryResponseParams
type CommonRecoveryResponseParams struct {

	// Specifies the id of the Recovery.
	// Pattern: ^\d+:\d+:\d+$
	ID *string `json:"id,omitempty"`

	// Specifies the name of the Recovery.
	Name *string `json:"name,omitempty"`

	// Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
	StartTimeUsecs *int64 `json:"startTimeUsecs,omitempty"`

	// Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
	EndTimeUsecs *int64 `json:"endTimeUsecs,omitempty"`

	// Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
	// Enum: ["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped"]
	Status *string `json:"status,omitempty"`

	// Progress monitor task id for Recovery.
	ProgressTaskID *string `json:"progressTaskId,omitempty"`

	// Specifies the type of snapshot environment for which the Recovery was performed.
	// Enum: ["kVMware","kHyperV","kAzure","kGCP","kKVM","kAcropolis","kAWS","kPhysical","kGPFS","kElastifile","kNetapp","kGenericNas","kIsilon","kFlashBlade","kPure","kIbmFlashSystem","kSQL","kExchange","kAD","kOracle","kView","kRemoteAdapter","kO365","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kUDA","kSfdc"]
	SnapshotEnvironment string `json:"snapshotEnvironment,omitempty"`

	// Specifies the type of recover action.
	// Enum: ["RecoverVMs","RecoverFiles","InstantVolumeMount","RecoverVmDisks","RecoverVApps","RecoverVAppTemplates","UptierSnapshot","RecoverRDS","RecoverAurora","RecoverS3Buckets","RecoverRDSPostgres","RecoverAzureSQL","RecoverApps","CloneApps","RecoverNasVolume","RecoverPhysicalVolumes","RecoverSystem","RecoverExchangeDbs","CloneAppView","RecoverSanVolumes","RecoverSanGroup","RecoverMailbox","RecoverOneDrive","RecoverSharePoint","RecoverPublicFolders","RecoverMsGroup","RecoverMsTeam","ConvertToPst","DownloadChats","RecoverNamespaces","RecoverObjects","RecoverSfdcObjects","RecoverSfdcOrg","RecoverSfdcRecords","DownloadFilesAndFolders","CloneVMs","CloneView","CloneRefreshApp","CloneVMsToView","ConvertAndDeployVMs","DeployVMs"]
	RecoveryAction string `json:"recoveryAction,omitempty"`

	// Specifies the list of tenants that have permissions for this recovery.
	Permissions []*Tenant `json:"permissions"`

	// Specifies the information about the creation of the recovery.
	CreationInfo *CreationInfo `json:"creationInfo,omitempty"`

	// Specifies whether it's possible to tear down the objects created by the recovery.
	CanTearDown *bool `json:"canTearDown,omitempty"`

	// Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. 'DestroyScheduled' indicates that the tear down is ready to schedule. 'Destroying' indicates that the tear down is still running. 'Destroyed' indicates that the tear down succeeded. 'DestroyError' indicates that the tear down failed.
	// Enum: ["DestroyScheduled","Destroying","Destroyed","DestroyError"]
	TearDownStatus *string `json:"tearDownStatus,omitempty"`

	// Specifies the error message about the tear down operation if it fails.
	TearDownMessage *string `json:"tearDownMessage,omitempty"`

	// Specifies messages about the recovery.
	Messages []string `json:"messages"`

	// Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery.
	IsParentRecovery *bool `json:"isParentRecovery,omitempty"`

	// If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery.
	// Pattern: ^\d+:\d+:\d+$
	ParentRecoveryID *string `json:"parentRecoveryId,omitempty"`

	// Specifies the list of persistent state of a retrieve of an archive task.
	RetrieveArchiveTasks []*RetrieveArchiveTask `json:"retrieveArchiveTasks"`

	// Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case.
	IsMultiStageRestore *bool `json:"isMultiStageRestore,omitempty"`
}

// Validate validates this common recovery response params
func (m *CommonRecoveryResponseParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotEnvironment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTearDownStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentRecoveryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetrieveArchiveTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonRecoveryResponseParams) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", *m.ID, `^\d+:\d+:\d+$`); err != nil {
		return err
	}

	return nil
}

var commonRecoveryResponseParamsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Accepted","Running","Canceled","Canceling","Failed","Missed","Succeeded","SucceededWithWarning","OnHold","Finalizing","Skipped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonRecoveryResponseParamsTypeStatusPropEnum = append(commonRecoveryResponseParamsTypeStatusPropEnum, v)
	}
}

const (

	// CommonRecoveryResponseParamsStatusAccepted captures enum value "Accepted"
	CommonRecoveryResponseParamsStatusAccepted string = "Accepted"

	// CommonRecoveryResponseParamsStatusRunning captures enum value "Running"
	CommonRecoveryResponseParamsStatusRunning string = "Running"

	// CommonRecoveryResponseParamsStatusCanceled captures enum value "Canceled"
	CommonRecoveryResponseParamsStatusCanceled string = "Canceled"

	// CommonRecoveryResponseParamsStatusCanceling captures enum value "Canceling"
	CommonRecoveryResponseParamsStatusCanceling string = "Canceling"

	// CommonRecoveryResponseParamsStatusFailed captures enum value "Failed"
	CommonRecoveryResponseParamsStatusFailed string = "Failed"

	// CommonRecoveryResponseParamsStatusMissed captures enum value "Missed"
	CommonRecoveryResponseParamsStatusMissed string = "Missed"

	// CommonRecoveryResponseParamsStatusSucceeded captures enum value "Succeeded"
	CommonRecoveryResponseParamsStatusSucceeded string = "Succeeded"

	// CommonRecoveryResponseParamsStatusSucceededWithWarning captures enum value "SucceededWithWarning"
	CommonRecoveryResponseParamsStatusSucceededWithWarning string = "SucceededWithWarning"

	// CommonRecoveryResponseParamsStatusOnHold captures enum value "OnHold"
	CommonRecoveryResponseParamsStatusOnHold string = "OnHold"

	// CommonRecoveryResponseParamsStatusFinalizing captures enum value "Finalizing"
	CommonRecoveryResponseParamsStatusFinalizing string = "Finalizing"

	// CommonRecoveryResponseParamsStatusSkipped captures enum value "Skipped"
	CommonRecoveryResponseParamsStatusSkipped string = "Skipped"
)

// prop value enum
func (m *CommonRecoveryResponseParams) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonRecoveryResponseParamsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonRecoveryResponseParams) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

var commonRecoveryResponseParamsTypeSnapshotEnvironmentPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kVMware","kHyperV","kAzure","kGCP","kKVM","kAcropolis","kAWS","kPhysical","kGPFS","kElastifile","kNetapp","kGenericNas","kIsilon","kFlashBlade","kPure","kIbmFlashSystem","kSQL","kExchange","kAD","kOracle","kView","kRemoteAdapter","kO365","kKubernetes","kCassandra","kMongoDB","kCouchbase","kHdfs","kHive","kHBase","kUDA","kSfdc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonRecoveryResponseParamsTypeSnapshotEnvironmentPropEnum = append(commonRecoveryResponseParamsTypeSnapshotEnvironmentPropEnum, v)
	}
}

const (

	// CommonRecoveryResponseParamsSnapshotEnvironmentKVMware captures enum value "kVMware"
	CommonRecoveryResponseParamsSnapshotEnvironmentKVMware string = "kVMware"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKHyperV captures enum value "kHyperV"
	CommonRecoveryResponseParamsSnapshotEnvironmentKHyperV string = "kHyperV"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKAzure captures enum value "kAzure"
	CommonRecoveryResponseParamsSnapshotEnvironmentKAzure string = "kAzure"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKGCP captures enum value "kGCP"
	CommonRecoveryResponseParamsSnapshotEnvironmentKGCP string = "kGCP"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKKVM captures enum value "kKVM"
	CommonRecoveryResponseParamsSnapshotEnvironmentKKVM string = "kKVM"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKAcropolis captures enum value "kAcropolis"
	CommonRecoveryResponseParamsSnapshotEnvironmentKAcropolis string = "kAcropolis"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKAWS captures enum value "kAWS"
	CommonRecoveryResponseParamsSnapshotEnvironmentKAWS string = "kAWS"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKPhysical captures enum value "kPhysical"
	CommonRecoveryResponseParamsSnapshotEnvironmentKPhysical string = "kPhysical"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKGPFS captures enum value "kGPFS"
	CommonRecoveryResponseParamsSnapshotEnvironmentKGPFS string = "kGPFS"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKElastifile captures enum value "kElastifile"
	CommonRecoveryResponseParamsSnapshotEnvironmentKElastifile string = "kElastifile"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKNetapp captures enum value "kNetapp"
	CommonRecoveryResponseParamsSnapshotEnvironmentKNetapp string = "kNetapp"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKGenericNas captures enum value "kGenericNas"
	CommonRecoveryResponseParamsSnapshotEnvironmentKGenericNas string = "kGenericNas"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKIsilon captures enum value "kIsilon"
	CommonRecoveryResponseParamsSnapshotEnvironmentKIsilon string = "kIsilon"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKFlashBlade captures enum value "kFlashBlade"
	CommonRecoveryResponseParamsSnapshotEnvironmentKFlashBlade string = "kFlashBlade"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKPure captures enum value "kPure"
	CommonRecoveryResponseParamsSnapshotEnvironmentKPure string = "kPure"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKIbmFlashSystem captures enum value "kIbmFlashSystem"
	CommonRecoveryResponseParamsSnapshotEnvironmentKIbmFlashSystem string = "kIbmFlashSystem"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKSQL captures enum value "kSQL"
	CommonRecoveryResponseParamsSnapshotEnvironmentKSQL string = "kSQL"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKExchange captures enum value "kExchange"
	CommonRecoveryResponseParamsSnapshotEnvironmentKExchange string = "kExchange"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKAD captures enum value "kAD"
	CommonRecoveryResponseParamsSnapshotEnvironmentKAD string = "kAD"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKOracle captures enum value "kOracle"
	CommonRecoveryResponseParamsSnapshotEnvironmentKOracle string = "kOracle"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKView captures enum value "kView"
	CommonRecoveryResponseParamsSnapshotEnvironmentKView string = "kView"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKRemoteAdapter captures enum value "kRemoteAdapter"
	CommonRecoveryResponseParamsSnapshotEnvironmentKRemoteAdapter string = "kRemoteAdapter"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKO365 captures enum value "kO365"
	CommonRecoveryResponseParamsSnapshotEnvironmentKO365 string = "kO365"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKKubernetes captures enum value "kKubernetes"
	CommonRecoveryResponseParamsSnapshotEnvironmentKKubernetes string = "kKubernetes"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKCassandra captures enum value "kCassandra"
	CommonRecoveryResponseParamsSnapshotEnvironmentKCassandra string = "kCassandra"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKMongoDB captures enum value "kMongoDB"
	CommonRecoveryResponseParamsSnapshotEnvironmentKMongoDB string = "kMongoDB"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKCouchbase captures enum value "kCouchbase"
	CommonRecoveryResponseParamsSnapshotEnvironmentKCouchbase string = "kCouchbase"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKHdfs captures enum value "kHdfs"
	CommonRecoveryResponseParamsSnapshotEnvironmentKHdfs string = "kHdfs"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKHive captures enum value "kHive"
	CommonRecoveryResponseParamsSnapshotEnvironmentKHive string = "kHive"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKHBase captures enum value "kHBase"
	CommonRecoveryResponseParamsSnapshotEnvironmentKHBase string = "kHBase"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKUDA captures enum value "kUDA"
	CommonRecoveryResponseParamsSnapshotEnvironmentKUDA string = "kUDA"

	// CommonRecoveryResponseParamsSnapshotEnvironmentKSfdc captures enum value "kSfdc"
	CommonRecoveryResponseParamsSnapshotEnvironmentKSfdc string = "kSfdc"
)

// prop value enum
func (m *CommonRecoveryResponseParams) validateSnapshotEnvironmentEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonRecoveryResponseParamsTypeSnapshotEnvironmentPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonRecoveryResponseParams) validateSnapshotEnvironment(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotEnvironment) { // not required
		return nil
	}

	// value enum
	if err := m.validateSnapshotEnvironmentEnum("snapshotEnvironment", "body", m.SnapshotEnvironment); err != nil {
		return err
	}

	return nil
}

var commonRecoveryResponseParamsTypeRecoveryActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RecoverVMs","RecoverFiles","InstantVolumeMount","RecoverVmDisks","RecoverVApps","RecoverVAppTemplates","UptierSnapshot","RecoverRDS","RecoverAurora","RecoverS3Buckets","RecoverRDSPostgres","RecoverAzureSQL","RecoverApps","CloneApps","RecoverNasVolume","RecoverPhysicalVolumes","RecoverSystem","RecoverExchangeDbs","CloneAppView","RecoverSanVolumes","RecoverSanGroup","RecoverMailbox","RecoverOneDrive","RecoverSharePoint","RecoverPublicFolders","RecoverMsGroup","RecoverMsTeam","ConvertToPst","DownloadChats","RecoverNamespaces","RecoverObjects","RecoverSfdcObjects","RecoverSfdcOrg","RecoverSfdcRecords","DownloadFilesAndFolders","CloneVMs","CloneView","CloneRefreshApp","CloneVMsToView","ConvertAndDeployVMs","DeployVMs"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonRecoveryResponseParamsTypeRecoveryActionPropEnum = append(commonRecoveryResponseParamsTypeRecoveryActionPropEnum, v)
	}
}

const (

	// CommonRecoveryResponseParamsRecoveryActionRecoverVMs captures enum value "RecoverVMs"
	CommonRecoveryResponseParamsRecoveryActionRecoverVMs string = "RecoverVMs"

	// CommonRecoveryResponseParamsRecoveryActionRecoverFiles captures enum value "RecoverFiles"
	CommonRecoveryResponseParamsRecoveryActionRecoverFiles string = "RecoverFiles"

	// CommonRecoveryResponseParamsRecoveryActionInstantVolumeMount captures enum value "InstantVolumeMount"
	CommonRecoveryResponseParamsRecoveryActionInstantVolumeMount string = "InstantVolumeMount"

	// CommonRecoveryResponseParamsRecoveryActionRecoverVMDisks captures enum value "RecoverVmDisks"
	CommonRecoveryResponseParamsRecoveryActionRecoverVMDisks string = "RecoverVmDisks"

	// CommonRecoveryResponseParamsRecoveryActionRecoverVApps captures enum value "RecoverVApps"
	CommonRecoveryResponseParamsRecoveryActionRecoverVApps string = "RecoverVApps"

	// CommonRecoveryResponseParamsRecoveryActionRecoverVAppTemplates captures enum value "RecoverVAppTemplates"
	CommonRecoveryResponseParamsRecoveryActionRecoverVAppTemplates string = "RecoverVAppTemplates"

	// CommonRecoveryResponseParamsRecoveryActionUptierSnapshot captures enum value "UptierSnapshot"
	CommonRecoveryResponseParamsRecoveryActionUptierSnapshot string = "UptierSnapshot"

	// CommonRecoveryResponseParamsRecoveryActionRecoverRDS captures enum value "RecoverRDS"
	CommonRecoveryResponseParamsRecoveryActionRecoverRDS string = "RecoverRDS"

	// CommonRecoveryResponseParamsRecoveryActionRecoverAurora captures enum value "RecoverAurora"
	CommonRecoveryResponseParamsRecoveryActionRecoverAurora string = "RecoverAurora"

	// CommonRecoveryResponseParamsRecoveryActionRecoverS3Buckets captures enum value "RecoverS3Buckets"
	CommonRecoveryResponseParamsRecoveryActionRecoverS3Buckets string = "RecoverS3Buckets"

	// CommonRecoveryResponseParamsRecoveryActionRecoverRDSPostgres captures enum value "RecoverRDSPostgres"
	CommonRecoveryResponseParamsRecoveryActionRecoverRDSPostgres string = "RecoverRDSPostgres"

	// CommonRecoveryResponseParamsRecoveryActionRecoverAzureSQL captures enum value "RecoverAzureSQL"
	CommonRecoveryResponseParamsRecoveryActionRecoverAzureSQL string = "RecoverAzureSQL"

	// CommonRecoveryResponseParamsRecoveryActionRecoverApps captures enum value "RecoverApps"
	CommonRecoveryResponseParamsRecoveryActionRecoverApps string = "RecoverApps"

	// CommonRecoveryResponseParamsRecoveryActionCloneApps captures enum value "CloneApps"
	CommonRecoveryResponseParamsRecoveryActionCloneApps string = "CloneApps"

	// CommonRecoveryResponseParamsRecoveryActionRecoverNasVolume captures enum value "RecoverNasVolume"
	CommonRecoveryResponseParamsRecoveryActionRecoverNasVolume string = "RecoverNasVolume"

	// CommonRecoveryResponseParamsRecoveryActionRecoverPhysicalVolumes captures enum value "RecoverPhysicalVolumes"
	CommonRecoveryResponseParamsRecoveryActionRecoverPhysicalVolumes string = "RecoverPhysicalVolumes"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSystem captures enum value "RecoverSystem"
	CommonRecoveryResponseParamsRecoveryActionRecoverSystem string = "RecoverSystem"

	// CommonRecoveryResponseParamsRecoveryActionRecoverExchangeDbs captures enum value "RecoverExchangeDbs"
	CommonRecoveryResponseParamsRecoveryActionRecoverExchangeDbs string = "RecoverExchangeDbs"

	// CommonRecoveryResponseParamsRecoveryActionCloneAppView captures enum value "CloneAppView"
	CommonRecoveryResponseParamsRecoveryActionCloneAppView string = "CloneAppView"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSanVolumes captures enum value "RecoverSanVolumes"
	CommonRecoveryResponseParamsRecoveryActionRecoverSanVolumes string = "RecoverSanVolumes"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSanGroup captures enum value "RecoverSanGroup"
	CommonRecoveryResponseParamsRecoveryActionRecoverSanGroup string = "RecoverSanGroup"

	// CommonRecoveryResponseParamsRecoveryActionRecoverMailbox captures enum value "RecoverMailbox"
	CommonRecoveryResponseParamsRecoveryActionRecoverMailbox string = "RecoverMailbox"

	// CommonRecoveryResponseParamsRecoveryActionRecoverOneDrive captures enum value "RecoverOneDrive"
	CommonRecoveryResponseParamsRecoveryActionRecoverOneDrive string = "RecoverOneDrive"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSharePoint captures enum value "RecoverSharePoint"
	CommonRecoveryResponseParamsRecoveryActionRecoverSharePoint string = "RecoverSharePoint"

	// CommonRecoveryResponseParamsRecoveryActionRecoverPublicFolders captures enum value "RecoverPublicFolders"
	CommonRecoveryResponseParamsRecoveryActionRecoverPublicFolders string = "RecoverPublicFolders"

	// CommonRecoveryResponseParamsRecoveryActionRecoverMsGroup captures enum value "RecoverMsGroup"
	CommonRecoveryResponseParamsRecoveryActionRecoverMsGroup string = "RecoverMsGroup"

	// CommonRecoveryResponseParamsRecoveryActionRecoverMsTeam captures enum value "RecoverMsTeam"
	CommonRecoveryResponseParamsRecoveryActionRecoverMsTeam string = "RecoverMsTeam"

	// CommonRecoveryResponseParamsRecoveryActionConvertToPst captures enum value "ConvertToPst"
	CommonRecoveryResponseParamsRecoveryActionConvertToPst string = "ConvertToPst"

	// CommonRecoveryResponseParamsRecoveryActionDownloadChats captures enum value "DownloadChats"
	CommonRecoveryResponseParamsRecoveryActionDownloadChats string = "DownloadChats"

	// CommonRecoveryResponseParamsRecoveryActionRecoverNamespaces captures enum value "RecoverNamespaces"
	CommonRecoveryResponseParamsRecoveryActionRecoverNamespaces string = "RecoverNamespaces"

	// CommonRecoveryResponseParamsRecoveryActionRecoverObjects captures enum value "RecoverObjects"
	CommonRecoveryResponseParamsRecoveryActionRecoverObjects string = "RecoverObjects"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSfdcObjects captures enum value "RecoverSfdcObjects"
	CommonRecoveryResponseParamsRecoveryActionRecoverSfdcObjects string = "RecoverSfdcObjects"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSfdcOrg captures enum value "RecoverSfdcOrg"
	CommonRecoveryResponseParamsRecoveryActionRecoverSfdcOrg string = "RecoverSfdcOrg"

	// CommonRecoveryResponseParamsRecoveryActionRecoverSfdcRecords captures enum value "RecoverSfdcRecords"
	CommonRecoveryResponseParamsRecoveryActionRecoverSfdcRecords string = "RecoverSfdcRecords"

	// CommonRecoveryResponseParamsRecoveryActionDownloadFilesAndFolders captures enum value "DownloadFilesAndFolders"
	CommonRecoveryResponseParamsRecoveryActionDownloadFilesAndFolders string = "DownloadFilesAndFolders"

	// CommonRecoveryResponseParamsRecoveryActionCloneVMs captures enum value "CloneVMs"
	CommonRecoveryResponseParamsRecoveryActionCloneVMs string = "CloneVMs"

	// CommonRecoveryResponseParamsRecoveryActionCloneView captures enum value "CloneView"
	CommonRecoveryResponseParamsRecoveryActionCloneView string = "CloneView"

	// CommonRecoveryResponseParamsRecoveryActionCloneRefreshApp captures enum value "CloneRefreshApp"
	CommonRecoveryResponseParamsRecoveryActionCloneRefreshApp string = "CloneRefreshApp"

	// CommonRecoveryResponseParamsRecoveryActionCloneVMsToView captures enum value "CloneVMsToView"
	CommonRecoveryResponseParamsRecoveryActionCloneVMsToView string = "CloneVMsToView"

	// CommonRecoveryResponseParamsRecoveryActionConvertAndDeployVMs captures enum value "ConvertAndDeployVMs"
	CommonRecoveryResponseParamsRecoveryActionConvertAndDeployVMs string = "ConvertAndDeployVMs"

	// CommonRecoveryResponseParamsRecoveryActionDeployVMs captures enum value "DeployVMs"
	CommonRecoveryResponseParamsRecoveryActionDeployVMs string = "DeployVMs"
)

// prop value enum
func (m *CommonRecoveryResponseParams) validateRecoveryActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonRecoveryResponseParamsTypeRecoveryActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonRecoveryResponseParams) validateRecoveryAction(formats strfmt.Registry) error {
	if swag.IsZero(m.RecoveryAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryActionEnum("recoveryAction", "body", m.RecoveryAction); err != nil {
		return err
	}

	return nil
}

func (m *CommonRecoveryResponseParams) validatePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	for i := 0; i < len(m.Permissions); i++ {
		if swag.IsZero(m.Permissions[i]) { // not required
			continue
		}

		if m.Permissions[i] != nil {
			if err := m.Permissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonRecoveryResponseParams) validateCreationInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationInfo) { // not required
		return nil
	}

	if m.CreationInfo != nil {
		if err := m.CreationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creationInfo")
			}
			return err
		}
	}

	return nil
}

var commonRecoveryResponseParamsTypeTearDownStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DestroyScheduled","Destroying","Destroyed","DestroyError"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		commonRecoveryResponseParamsTypeTearDownStatusPropEnum = append(commonRecoveryResponseParamsTypeTearDownStatusPropEnum, v)
	}
}

const (

	// CommonRecoveryResponseParamsTearDownStatusDestroyScheduled captures enum value "DestroyScheduled"
	CommonRecoveryResponseParamsTearDownStatusDestroyScheduled string = "DestroyScheduled"

	// CommonRecoveryResponseParamsTearDownStatusDestroying captures enum value "Destroying"
	CommonRecoveryResponseParamsTearDownStatusDestroying string = "Destroying"

	// CommonRecoveryResponseParamsTearDownStatusDestroyed captures enum value "Destroyed"
	CommonRecoveryResponseParamsTearDownStatusDestroyed string = "Destroyed"

	// CommonRecoveryResponseParamsTearDownStatusDestroyError captures enum value "DestroyError"
	CommonRecoveryResponseParamsTearDownStatusDestroyError string = "DestroyError"
)

// prop value enum
func (m *CommonRecoveryResponseParams) validateTearDownStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, commonRecoveryResponseParamsTypeTearDownStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CommonRecoveryResponseParams) validateTearDownStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.TearDownStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateTearDownStatusEnum("tearDownStatus", "body", *m.TearDownStatus); err != nil {
		return err
	}

	return nil
}

func (m *CommonRecoveryResponseParams) validateParentRecoveryID(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentRecoveryID) { // not required
		return nil
	}

	if err := validate.Pattern("parentRecoveryId", "body", *m.ParentRecoveryID, `^\d+:\d+:\d+$`); err != nil {
		return err
	}

	return nil
}

func (m *CommonRecoveryResponseParams) validateRetrieveArchiveTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.RetrieveArchiveTasks) { // not required
		return nil
	}

	for i := 0; i < len(m.RetrieveArchiveTasks); i++ {
		if swag.IsZero(m.RetrieveArchiveTasks[i]) { // not required
			continue
		}

		if m.RetrieveArchiveTasks[i] != nil {
			if err := m.RetrieveArchiveTasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("retrieveArchiveTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("retrieveArchiveTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this common recovery response params based on the context it is used
func (m *CommonRecoveryResponseParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreationInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRetrieveArchiveTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommonRecoveryResponseParams) contextValidatePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Permissions); i++ {

		if m.Permissions[i] != nil {

			if swag.IsZero(m.Permissions[i]) { // not required
				return nil
			}

			if err := m.Permissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommonRecoveryResponseParams) contextValidateCreationInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CreationInfo != nil {

		if swag.IsZero(m.CreationInfo) { // not required
			return nil
		}

		if err := m.CreationInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("creationInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("creationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *CommonRecoveryResponseParams) contextValidateRetrieveArchiveTasks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RetrieveArchiveTasks); i++ {

		if m.RetrieveArchiveTasks[i] != nil {

			if swag.IsZero(m.RetrieveArchiveTasks[i]) { // not required
				return nil
			}

			if err := m.RetrieveArchiveTasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("retrieveArchiveTasks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("retrieveArchiveTasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommonRecoveryResponseParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommonRecoveryResponseParams) UnmarshalBinary(b []byte) error {
	var res CommonRecoveryResponseParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

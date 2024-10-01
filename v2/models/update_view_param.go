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

// UpdateViewParam Specifies the settings that define a View. Common fields between create,
// edit and response struct.
//
// swagger:model UpdateViewParam
type UpdateViewParam struct {

	// Specifies the name of the View.
	Name *string `json:"name,omitempty"`

	// Specifies the category of the View.
	// Enum: ["BackupTarget","FileServices","ObjectServices"]
	Category *string `json:"category,omitempty"`

	// Specifies the supported Protocols for the View.
	ProtocolAccess []*ViewProtocol `json:"protocolAccess,omitempty"`

	// Specifies whether view level client subnet whitelist overrides cluster and
	// global setting.
	OverrideGlobalSubnetWhitelist *bool `json:"overrideGlobalSubnetWhitelist,omitempty"`

	// Array of Subnets.
	// Specifies a list of Subnets with IP addresses that have permissions to
	// access the View. (Overrides or extends the Subnets specified at the global
	// Cohesity Cluster level.)
	SubnetWhitelist []*Subnet `json:"subnetWhitelist,omitempty"`

	// Specifies whether view level client netgroup whitelist overrides cluster and
	// global setting.
	OverrideGlobalNetgroupWhitelist *bool `json:"overrideGlobalNetgroupWhitelist,omitempty"`

	// Specifies the security mode used for this view.
	// Currently we support the following modes: Native, Unified and NTFS style.
	// 'NativeMode' indicates a native security mode.
	// 'UnifiedMode' indicates a unified security mode.
	// 'NtfsMode' indicates a NTFS style security mode.
	// Enum: ["NativeMode","UnifiedMode","NtfsMode"]
	SecurityMode *string `json:"securityMode,omitempty"`

	// Specifies an optional text description about the View.
	Description *string `json:"description,omitempty"`

	// Specifies if this View can be mounted using the NFS protocol
	// on Windows systems. If true, this View can be NFS mounted on Windows
	// systems.
	AllowMountOnWindows *bool `json:"allowMountOnWindows,omitempty"`

	// Specifies if this view should allow minion or not. If true, this will
	// allow minion.
	EnableMinion *bool `json:"enableMinion,omitempty"`

	// Specifies if Filer Audit Logging is enabled for this view.
	EnableFilerAuditLogging *bool `json:"enableFilerAuditLogging,omitempty"`

	// Optional tenant id who has access to this View.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies whether to enable live indexing for the view.
	EnableLiveIndexing *bool `json:"enableLiveIndexing,omitempty"`

	// Specifies whether to enable offline file caching of the view.
	EnableOfflineCaching *bool `json:"enableOfflineCaching,omitempty"`

	// Array of Security Identifiers (SIDs)
	// Specifies the list of security identifiers (SIDs) for the restricted
	// Principals who have access to this View.
	AccessSids []string `json:"accessSids,omitempty"`

	// Specifies whether view lock is enabled. If enabled the view cannot be
	// modified or deleted until unlock. By default it is disabled.
	ViewLockEnabled *bool `json:"viewLockEnabled,omitempty"`

	// Specifies if the view is a read only view. User will no longer be able to write to this view if this is set to true.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Specifies if metadata accelerator is enabled for this view. Only
	// supported while creating a view.
	EnableMetadataAccelerator *bool `json:"enableMetadataAccelerator,omitempty"`

	// Specifies whether the view is for externally triggered backup
	// target. If so, Magneto will ignore the backup schedule for the
	// view protection job of this view. By default it is disabled.
	IsExternallyTriggeredBackupTarget *bool `json:"isExternallyTriggeredBackupTarget,omitempty"`

	// If small files are accessed sequentially from a client,
	// this specifies whether to detect and
	// prefetch files based on the lexicographic index to
	// improve file read performance.
	LexicographicPrefetch *bool `json:"lexicographicPrefetch,omitempty"`

	// Specifies the antivirus scan config settings for this View.
	AntivirusScanConfig *AntivirusScanConfig `json:"antivirusScanConfig,omitempty"`

	// Optional filtering criteria that should be satisfied by all the files
	// created in this view. It does not affect existing files.
	FileExtensionFilter *FileExtensionFilter `json:"fileExtensionFilter,omitempty"`

	// Optional config that enables file locking for this view. It cannot be
	// disabled during the edit of a view, if it has been enabled during the
	// creation of the view. Also, it cannot be enabled if it was disabled
	// during the creation of the view.
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`

	// Specifies the Lifecycle policy of this filer (NFS/SMB) view.
	FilerLifecycleManagement *FilerLifecycleManagement `json:"filerLifecycleManagement,omitempty"`

	// Specifies an optional logical quota limit (in bytes) for the usage allowed
	// on this View.
	// (Logical data is when the data is fully hydrated and expanded.)
	// This limit overrides the limit inherited from the Storage Domain
	// (View Box) (if set).
	// If logicalQuota is nil, the limit is inherited from the
	// Storage Domain (View Box) (if set).
	// A new write is not allowed if the Storage Domain (View Box) will exceed the
	// specified quota.
	// However, it takes time for the Cohesity Cluster to calculate
	// the usage across Nodes, so the limit may be exceeded by a small amount.
	// In addition, if the limit is increased or data is removed,
	// there may be a delay before the Cohesity Cluster allows more data
	// to be written to the View, as the Cluster is calculating the usage
	// across Nodes.
	LogicalQuota *QuotaPolicy `json:"logicalQuota,omitempty"`

	// Array of Netgroups.
	// Specifies a list of netgroups with domains that have permissions to
	// access the View. (Overrides or extends the Netgroup specified at the global
	// Cohesity Cluster level.)
	NetgroupWhitelist *NisNetgroups `json:"netgroupWhitelist,omitempty"`

	// Specifies the Quality of Service (QoS) Policy for the View.
	Qos *QoS `json:"qos,omitempty"`

	// Specifies self service config of this view.
	SelfServiceSnapshotConfig *SelfServiceSnapshotConfig `json:"selfServiceSnapshotConfig,omitempty"`

	// Specifies if inline deduplication and compression settings inherited from
	// the Storage Domain (View Box) should be disabled for this View.
	StoragePolicyOverride *StoragePolicyOverride `json:"storagePolicyOverride,omitempty"`

	// Specifies the pinning config of this view.
	ViewPinningConfig *ViewPinningConfig `json:"viewPinningConfig,omitempty"`

	NfsConfig

	SmbConfig

	S3Config

	SwiftConfig
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *UpdateViewParam) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		Name *string `json:"name,omitempty"`

		Category *string `json:"category,omitempty"`

		ProtocolAccess []*ViewProtocol `json:"protocolAccess,omitempty"`

		OverrideGlobalSubnetWhitelist *bool `json:"overrideGlobalSubnetWhitelist,omitempty"`

		SubnetWhitelist []*Subnet `json:"subnetWhitelist,omitempty"`

		OverrideGlobalNetgroupWhitelist *bool `json:"overrideGlobalNetgroupWhitelist,omitempty"`

		SecurityMode *string `json:"securityMode,omitempty"`

		Description *string `json:"description,omitempty"`

		AllowMountOnWindows *bool `json:"allowMountOnWindows,omitempty"`

		EnableMinion *bool `json:"enableMinion,omitempty"`

		EnableFilerAuditLogging *bool `json:"enableFilerAuditLogging,omitempty"`

		TenantID *string `json:"tenantId,omitempty"`

		EnableLiveIndexing *bool `json:"enableLiveIndexing,omitempty"`

		EnableOfflineCaching *bool `json:"enableOfflineCaching,omitempty"`

		AccessSids []string `json:"accessSids,omitempty"`

		ViewLockEnabled *bool `json:"viewLockEnabled,omitempty"`

		IsReadOnly *bool `json:"isReadOnly,omitempty"`

		EnableMetadataAccelerator *bool `json:"enableMetadataAccelerator,omitempty"`

		IsExternallyTriggeredBackupTarget *bool `json:"isExternallyTriggeredBackupTarget,omitempty"`

		LexicographicPrefetch *bool `json:"lexicographicPrefetch,omitempty"`

		AntivirusScanConfig *AntivirusScanConfig `json:"antivirusScanConfig,omitempty"`

		FileExtensionFilter *FileExtensionFilter `json:"fileExtensionFilter,omitempty"`

		FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`

		FilerLifecycleManagement *FilerLifecycleManagement `json:"filerLifecycleManagement,omitempty"`

		LogicalQuota *QuotaPolicy `json:"logicalQuota,omitempty"`

		NetgroupWhitelist *NisNetgroups `json:"netgroupWhitelist,omitempty"`

		Qos *QoS `json:"qos,omitempty"`

		SelfServiceSnapshotConfig *SelfServiceSnapshotConfig `json:"selfServiceSnapshotConfig,omitempty"`

		StoragePolicyOverride *StoragePolicyOverride `json:"storagePolicyOverride,omitempty"`

		ViewPinningConfig *ViewPinningConfig `json:"viewPinningConfig,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.Name = dataAO0.Name

	m.Category = dataAO0.Category

	m.ProtocolAccess = dataAO0.ProtocolAccess

	m.OverrideGlobalSubnetWhitelist = dataAO0.OverrideGlobalSubnetWhitelist

	m.SubnetWhitelist = dataAO0.SubnetWhitelist

	m.OverrideGlobalNetgroupWhitelist = dataAO0.OverrideGlobalNetgroupWhitelist

	m.SecurityMode = dataAO0.SecurityMode

	m.Description = dataAO0.Description

	m.AllowMountOnWindows = dataAO0.AllowMountOnWindows

	m.EnableMinion = dataAO0.EnableMinion

	m.EnableFilerAuditLogging = dataAO0.EnableFilerAuditLogging

	m.TenantID = dataAO0.TenantID

	m.EnableLiveIndexing = dataAO0.EnableLiveIndexing

	m.EnableOfflineCaching = dataAO0.EnableOfflineCaching

	m.AccessSids = dataAO0.AccessSids

	m.ViewLockEnabled = dataAO0.ViewLockEnabled

	m.IsReadOnly = dataAO0.IsReadOnly

	m.EnableMetadataAccelerator = dataAO0.EnableMetadataAccelerator

	m.IsExternallyTriggeredBackupTarget = dataAO0.IsExternallyTriggeredBackupTarget

	m.LexicographicPrefetch = dataAO0.LexicographicPrefetch

	m.AntivirusScanConfig = dataAO0.AntivirusScanConfig

	m.FileExtensionFilter = dataAO0.FileExtensionFilter

	m.FileLockConfig = dataAO0.FileLockConfig

	m.FilerLifecycleManagement = dataAO0.FilerLifecycleManagement

	m.LogicalQuota = dataAO0.LogicalQuota

	m.NetgroupWhitelist = dataAO0.NetgroupWhitelist

	m.Qos = dataAO0.Qos

	m.SelfServiceSnapshotConfig = dataAO0.SelfServiceSnapshotConfig

	m.StoragePolicyOverride = dataAO0.StoragePolicyOverride

	m.ViewPinningConfig = dataAO0.ViewPinningConfig

	// AO1
	var aO1 NfsConfig
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NfsConfig = aO1

	// AO2
	var aO2 SmbConfig
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.SmbConfig = aO2

	// AO3
	var aO3 S3Config
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.S3Config = aO3

	// AO4
	var aO4 SwiftConfig
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.SwiftConfig = aO4

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m UpdateViewParam) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 5)

	var dataAO0 struct {
		Name *string `json:"name,omitempty"`

		Category *string `json:"category,omitempty"`

		ProtocolAccess []*ViewProtocol `json:"protocolAccess,omitempty"`

		OverrideGlobalSubnetWhitelist *bool `json:"overrideGlobalSubnetWhitelist,omitempty"`

		SubnetWhitelist []*Subnet `json:"subnetWhitelist,omitempty"`

		OverrideGlobalNetgroupWhitelist *bool `json:"overrideGlobalNetgroupWhitelist,omitempty"`

		SecurityMode *string `json:"securityMode,omitempty"`

		Description *string `json:"description,omitempty"`

		AllowMountOnWindows *bool `json:"allowMountOnWindows,omitempty"`

		EnableMinion *bool `json:"enableMinion,omitempty"`

		EnableFilerAuditLogging *bool `json:"enableFilerAuditLogging,omitempty"`

		TenantID *string `json:"tenantId,omitempty"`

		EnableLiveIndexing *bool `json:"enableLiveIndexing,omitempty"`

		EnableOfflineCaching *bool `json:"enableOfflineCaching,omitempty"`

		AccessSids []string `json:"accessSids,omitempty"`

		ViewLockEnabled *bool `json:"viewLockEnabled,omitempty"`

		IsReadOnly *bool `json:"isReadOnly,omitempty"`

		EnableMetadataAccelerator *bool `json:"enableMetadataAccelerator,omitempty"`

		IsExternallyTriggeredBackupTarget *bool `json:"isExternallyTriggeredBackupTarget,omitempty"`

		LexicographicPrefetch *bool `json:"lexicographicPrefetch,omitempty"`

		AntivirusScanConfig *AntivirusScanConfig `json:"antivirusScanConfig,omitempty"`

		FileExtensionFilter *FileExtensionFilter `json:"fileExtensionFilter,omitempty"`

		FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`

		FilerLifecycleManagement *FilerLifecycleManagement `json:"filerLifecycleManagement,omitempty"`

		LogicalQuota *QuotaPolicy `json:"logicalQuota,omitempty"`

		NetgroupWhitelist *NisNetgroups `json:"netgroupWhitelist,omitempty"`

		Qos *QoS `json:"qos,omitempty"`

		SelfServiceSnapshotConfig *SelfServiceSnapshotConfig `json:"selfServiceSnapshotConfig,omitempty"`

		StoragePolicyOverride *StoragePolicyOverride `json:"storagePolicyOverride,omitempty"`

		ViewPinningConfig *ViewPinningConfig `json:"viewPinningConfig,omitempty"`
	}

	dataAO0.Name = m.Name

	dataAO0.Category = m.Category

	dataAO0.ProtocolAccess = m.ProtocolAccess

	dataAO0.OverrideGlobalSubnetWhitelist = m.OverrideGlobalSubnetWhitelist

	dataAO0.SubnetWhitelist = m.SubnetWhitelist

	dataAO0.OverrideGlobalNetgroupWhitelist = m.OverrideGlobalNetgroupWhitelist

	dataAO0.SecurityMode = m.SecurityMode

	dataAO0.Description = m.Description

	dataAO0.AllowMountOnWindows = m.AllowMountOnWindows

	dataAO0.EnableMinion = m.EnableMinion

	dataAO0.EnableFilerAuditLogging = m.EnableFilerAuditLogging

	dataAO0.TenantID = m.TenantID

	dataAO0.EnableLiveIndexing = m.EnableLiveIndexing

	dataAO0.EnableOfflineCaching = m.EnableOfflineCaching

	dataAO0.AccessSids = m.AccessSids

	dataAO0.ViewLockEnabled = m.ViewLockEnabled

	dataAO0.IsReadOnly = m.IsReadOnly

	dataAO0.EnableMetadataAccelerator = m.EnableMetadataAccelerator

	dataAO0.IsExternallyTriggeredBackupTarget = m.IsExternallyTriggeredBackupTarget

	dataAO0.LexicographicPrefetch = m.LexicographicPrefetch

	dataAO0.AntivirusScanConfig = m.AntivirusScanConfig

	dataAO0.FileExtensionFilter = m.FileExtensionFilter

	dataAO0.FileLockConfig = m.FileLockConfig

	dataAO0.FilerLifecycleManagement = m.FilerLifecycleManagement

	dataAO0.LogicalQuota = m.LogicalQuota

	dataAO0.NetgroupWhitelist = m.NetgroupWhitelist

	dataAO0.Qos = m.Qos

	dataAO0.SelfServiceSnapshotConfig = m.SelfServiceSnapshotConfig

	dataAO0.StoragePolicyOverride = m.StoragePolicyOverride

	dataAO0.ViewPinningConfig = m.ViewPinningConfig

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.NfsConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.SmbConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.S3Config)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.SwiftConfig)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this update view param
func (m *UpdateViewParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAntivirusScanConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileExtensionFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileLockConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilerLifecycleManagement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetgroupWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfServiceSnapshotConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicyOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewPinningConfig(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with NfsConfig
	if err := m.NfsConfig.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SmbConfig
	if err := m.SmbConfig.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with S3Config
	if err := m.S3Config.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SwiftConfig
	if err := m.SwiftConfig.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateViewParamTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BackupTarget","FileServices","ObjectServices"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateViewParamTypeCategoryPropEnum = append(updateViewParamTypeCategoryPropEnum, v)
	}
}

// property enum
func (m *UpdateViewParam) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateViewParamTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateViewParam) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", *m.Category); err != nil {
		return err
	}

	return nil
}

func (m *UpdateViewParam) validateProtocolAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.ProtocolAccess) { // not required
		return nil
	}

	for i := 0; i < len(m.ProtocolAccess); i++ {
		if swag.IsZero(m.ProtocolAccess[i]) { // not required
			continue
		}

		if m.ProtocolAccess[i] != nil {
			if err := m.ProtocolAccess[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocolAccess" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protocolAccess" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateViewParam) validateSubnetWhitelist(formats strfmt.Registry) error {

	if swag.IsZero(m.SubnetWhitelist) { // not required
		return nil
	}

	for i := 0; i < len(m.SubnetWhitelist); i++ {
		if swag.IsZero(m.SubnetWhitelist[i]) { // not required
			continue
		}

		if m.SubnetWhitelist[i] != nil {
			if err := m.SubnetWhitelist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var updateViewParamTypeSecurityModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NativeMode","UnifiedMode","NtfsMode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateViewParamTypeSecurityModePropEnum = append(updateViewParamTypeSecurityModePropEnum, v)
	}
}

// property enum
func (m *UpdateViewParam) validateSecurityModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateViewParamTypeSecurityModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateViewParam) validateSecurityMode(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateSecurityModeEnum("securityMode", "body", *m.SecurityMode); err != nil {
		return err
	}

	return nil
}

func (m *UpdateViewParam) validateAntivirusScanConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.AntivirusScanConfig) { // not required
		return nil
	}

	if m.AntivirusScanConfig != nil {
		if err := m.AntivirusScanConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antivirusScanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("antivirusScanConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateFileExtensionFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.FileExtensionFilter) { // not required
		return nil
	}

	if m.FileExtensionFilter != nil {
		if err := m.FileExtensionFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileExtensionFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileExtensionFilter")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateFileLockConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.FileLockConfig) { // not required
		return nil
	}

	if m.FileLockConfig != nil {
		if err := m.FileLockConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileLockConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileLockConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateFilerLifecycleManagement(formats strfmt.Registry) error {

	if swag.IsZero(m.FilerLifecycleManagement) { // not required
		return nil
	}

	if m.FilerLifecycleManagement != nil {
		if err := m.FilerLifecycleManagement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filerLifecycleManagement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filerLifecycleManagement")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateLogicalQuota(formats strfmt.Registry) error {

	if swag.IsZero(m.LogicalQuota) { // not required
		return nil
	}

	if m.LogicalQuota != nil {
		if err := m.LogicalQuota.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logicalQuota")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateNetgroupWhitelist(formats strfmt.Registry) error {

	if swag.IsZero(m.NetgroupWhitelist) { // not required
		return nil
	}

	if m.NetgroupWhitelist != nil {
		if err := m.NetgroupWhitelist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netgroupWhitelist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netgroupWhitelist")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateQos(formats strfmt.Registry) error {

	if swag.IsZero(m.Qos) { // not required
		return nil
	}

	if m.Qos != nil {
		if err := m.Qos.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateSelfServiceSnapshotConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfServiceSnapshotConfig) { // not required
		return nil
	}

	if m.SelfServiceSnapshotConfig != nil {
		if err := m.SelfServiceSnapshotConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selfServiceSnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selfServiceSnapshotConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateStoragePolicyOverride(formats strfmt.Registry) error {

	if swag.IsZero(m.StoragePolicyOverride) { // not required
		return nil
	}

	if m.StoragePolicyOverride != nil {
		if err := m.StoragePolicyOverride.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storagePolicyOverride")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storagePolicyOverride")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) validateViewPinningConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.ViewPinningConfig) { // not required
		return nil
	}

	if m.ViewPinningConfig != nil {
		if err := m.ViewPinningConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewPinningConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewPinningConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update view param based on the context it is used
func (m *UpdateViewParam) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProtocolAccess(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAntivirusScanConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileExtensionFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFileLockConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFilerLifecycleManagement(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogicalQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetgroupWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfServiceSnapshotConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoragePolicyOverride(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewPinningConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with NfsConfig
	if err := m.NfsConfig.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SmbConfig
	if err := m.SmbConfig.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with S3Config
	if err := m.S3Config.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SwiftConfig
	if err := m.SwiftConfig.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateViewParam) contextValidateProtocolAccess(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProtocolAccess); i++ {

		if m.ProtocolAccess[i] != nil {

			if swag.IsZero(m.ProtocolAccess[i]) { // not required
				return nil
			}

			if err := m.ProtocolAccess[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("protocolAccess" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("protocolAccess" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateViewParam) contextValidateSubnetWhitelist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SubnetWhitelist); i++ {

		if m.SubnetWhitelist[i] != nil {

			if swag.IsZero(m.SubnetWhitelist[i]) { // not required
				return nil
			}

			if err := m.SubnetWhitelist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("subnetWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateViewParam) contextValidateAntivirusScanConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.AntivirusScanConfig != nil {

		if swag.IsZero(m.AntivirusScanConfig) { // not required
			return nil
		}

		if err := m.AntivirusScanConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("antivirusScanConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("antivirusScanConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateFileExtensionFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.FileExtensionFilter != nil {

		if swag.IsZero(m.FileExtensionFilter) { // not required
			return nil
		}

		if err := m.FileExtensionFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileExtensionFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileExtensionFilter")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateFileLockConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.FileLockConfig != nil {

		if swag.IsZero(m.FileLockConfig) { // not required
			return nil
		}

		if err := m.FileLockConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileLockConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fileLockConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateFilerLifecycleManagement(ctx context.Context, formats strfmt.Registry) error {

	if m.FilerLifecycleManagement != nil {

		if swag.IsZero(m.FilerLifecycleManagement) { // not required
			return nil
		}

		if err := m.FilerLifecycleManagement.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filerLifecycleManagement")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("filerLifecycleManagement")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateLogicalQuota(ctx context.Context, formats strfmt.Registry) error {

	if m.LogicalQuota != nil {

		if swag.IsZero(m.LogicalQuota) { // not required
			return nil
		}

		if err := m.LogicalQuota.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("logicalQuota")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("logicalQuota")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateNetgroupWhitelist(ctx context.Context, formats strfmt.Registry) error {

	if m.NetgroupWhitelist != nil {

		if swag.IsZero(m.NetgroupWhitelist) { // not required
			return nil
		}

		if err := m.NetgroupWhitelist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("netgroupWhitelist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("netgroupWhitelist")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateQos(ctx context.Context, formats strfmt.Registry) error {

	if m.Qos != nil {

		if swag.IsZero(m.Qos) { // not required
			return nil
		}

		if err := m.Qos.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("qos")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateSelfServiceSnapshotConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.SelfServiceSnapshotConfig != nil {

		if swag.IsZero(m.SelfServiceSnapshotConfig) { // not required
			return nil
		}

		if err := m.SelfServiceSnapshotConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selfServiceSnapshotConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("selfServiceSnapshotConfig")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateStoragePolicyOverride(ctx context.Context, formats strfmt.Registry) error {

	if m.StoragePolicyOverride != nil {

		if swag.IsZero(m.StoragePolicyOverride) { // not required
			return nil
		}

		if err := m.StoragePolicyOverride.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storagePolicyOverride")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storagePolicyOverride")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateViewParam) contextValidateViewPinningConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.ViewPinningConfig != nil {

		if swag.IsZero(m.ViewPinningConfig) { // not required
			return nil
		}

		if err := m.ViewPinningConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewPinningConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewPinningConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateViewParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateViewParam) UnmarshalBinary(b []byte) error {
	var res UpdateViewParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

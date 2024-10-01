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

// View View.
//
// Specifies settings for defining a storage location (called a View)
// with NFS and SMB mount paths in a Storage Domain (View Box) on the Cohesity
// Cluster.
//
// swagger:model View
type View struct {

	// Array of Security Identifiers (SIDs)
	//
	// Specifies the list of security identifiers (SIDs) for the restricted
	// Principals who have access to this View.
	AccessSids []string `json:"accessSids"`

	// Aliases created for the view. A view alias allows a directory path inside
	// a view to be mounted using the alias name.
	Aliases []*ViewAlias `json:"aliases"`

	// Array of SMB Paths.
	//
	// Specifies the possible paths that can be used to mount this View
	// as a SMB share. If Active Directory has multiple account names;
	// each machine account has its own path.
	AllSmbMountPaths []string `json:"allSmbMountPaths"`

	// Specifies the antivirus scan config settings for this View.
	AntivirusScanConfig *AntivirusScanConfig `json:"antivirusScanConfig,omitempty"`

	// Specifies the NFS mount path of the View (without the hostname
	// information).
	// This path is used to support NFS mounting of the paths specified in the
	// nfsExportPathList on Windows systems.
	BasicMountPath *string `json:"basicMountPath,omitempty"`

	// Specifies whether to support case insensitive file/folder names. This
	// parameter can only be set during create and cannot be changed.
	CaseInsensitiveNamesEnabled *bool `json:"caseInsensitiveNamesEnabled,omitempty"`

	// Specifies the time that the View was created in milliseconds.
	CreateTimeMsecs *int64 `json:"createTimeMsecs,omitempty"`

	// DataLock (Write Once Read Many) lock expiry epoch time in microseconds. If
	// a view is marked as a DataLock view, only a Data Security Officer (a user
	// having Data Security Privilege) can delete the view until the lock expiry
	// time.
	DataLockExpiryUsecs *int64 `json:"dataLockExpiryUsecs,omitempty"`

	// Specifies an optional text description about the View.
	Description *string `json:"description,omitempty"`

	// Specifies whether fast durable handle is enabled. If enabled, view open
	// handle will be kept in memory, which results in a higher performance. But
	// the handles cannot be recovered if node or service crashes.
	EnableFastDurableHandle *bool `json:"enableFastDurableHandle,omitempty"`

	// Specifies if Filer Audit Logging is enabled for this view.
	EnableFilerAuditLogging *bool `json:"enableFilerAuditLogging,omitempty"`

	// Specifies whether to enable live indexing for the view.
	EnableLiveIndexing *bool `json:"enableLiveIndexing,omitempty"`

	// Specifies whether view is blur enabled.
	EnableMetadataAccelerator *bool `json:"enableMetadataAccelerator,omitempty"`

	// If set, mixed mode (NFS and SMB) access is enabled for this view.
	// This field is deprecated. Use the field SecurityMode.
	// deprecated: true
	EnableMixedModePermissions *bool `json:"enableMixedModePermissions,omitempty"`

	// If set, it enables discovery of view for NFS.
	EnableNfsViewDiscovery *bool `json:"enableNfsViewDiscovery,omitempty"`

	// Specifies whether to enable offline file caching of the view.
	EnableOfflineCaching *bool `json:"enableOfflineCaching,omitempty"`

	// Specifies if access-based enumeration should be enabled.
	// If 'true', only files and folders that the user has permissions to
	// access are visible on the SMB share for that user.
	EnableSmbAccessBasedEnumeration *bool `json:"enableSmbAccessBasedEnumeration,omitempty"`

	// Specifies the SMB encryption for the View. If set, it enables the SMB
	// encryption for the View. Encryption is supported only by SMB 3.x dialects.
	// Dialects that do not support would still access data in unencrypted
	// format.
	EnableSmbEncryption *bool `json:"enableSmbEncryption,omitempty"`

	// Specifies whether SMB opportunistic lock is enabled.
	EnableSmbOplock *bool `json:"enableSmbOplock,omitempty"`

	// If set, it enables discovery of view for SMB.
	EnableSmbViewDiscovery *bool `json:"enableSmbViewDiscovery,omitempty"`

	// Specifies the SMB encryption for all the sessions for the View.
	// If set, encryption is enforced for all the sessions for the View. When
	// enabled all future and existing unencrypted sessions are disallowed.
	EnforceSmbEncryption *bool `json:"enforceSmbEncryption,omitempty"`

	// Optional filtering criteria that should be satisfied by all the files
	// created in this view. It does not affect existing files.
	FileExtensionFilter *FileExtensionFilter `json:"fileExtensionFilter,omitempty"`

	// Optional config that enables file locking for this view. It cannot be
	// disabled during the edit of a view, if it has been enabled during the
	// creation of the view. Also, it cannot be enabled if it was disabled
	// during the creation of the view.
	FileLockConfig *FileLevelDataLockConfig `json:"fileLockConfig,omitempty"`

	// Specifies the Intent of the View.
	// readOnly: true
	Intent *ViewIntent `json:"intent,omitempty"`

	// Specifies whether view is for externally triggered backup target.
	IsExternallyTriggeredBackupTarget *bool `json:"isExternallyTriggeredBackupTarget,omitempty"`

	// Specifies if the view is a read only view. User will no longer be able to
	// write to this view if this is set to true.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Specifies if a view contains migrated data.
	IsTargetForMigratedData *bool `json:"isTargetForMigratedData,omitempty"`

	// logical quota
	LogicalQuota *ViewLogicalQuota `json:"logicalQuota,omitempty"`

	// LogicalUsageBytes is the logical usage in bytes for the view.
	LogicalUsageBytes *int64 `json:"logicalUsageBytes,omitempty"`

	// Specifies the name of the View.
	Name *string `json:"name,omitempty"`

	// Array of Netgroups.
	//
	// Specifies a list of Netgroups that have permissions to access the
	// View. (Overrides the Netgroups specified at the global Cohesity
	// Cluster level.)
	NetgroupWhitelist []*NisNetgroup `json:"netgroupWhitelist"`

	// Specifies the NFS all squash config.
	NfsAllSquash *NfsSquash `json:"nfsAllSquash,omitempty"`

	// Specifies the path for mounting this View as an NFS share.
	NfsMountPath *string `json:"nfsMountPath,omitempty"`

	// Specifies the NFS root permission config of the view file system.
	NfsRootPermissions *NfsRootPermissions `json:"nfsRootPermissions,omitempty"`

	// Specifies the NFS root squash config.
	NfsRootSquash *NfsSquash `json:"nfsRootSquash,omitempty"`

	// Specifies whether view level client netgroup allowlist overrides cluster
	// and global setting.
	OverrideGlobalNetgroupWhitelist *bool `json:"overrideGlobalNetgroupWhitelist,omitempty"`

	// Specifies whether view level client subnet allowlist overrides cluster and
	// global setting.
	OverrideGlobalWhitelist *bool `json:"overrideGlobalWhitelist,omitempty"`

	// Specifies the Sid of the view owner.
	OwnerSid *string `json:"ownerSid,omitempty"`

	// Specifies the supported Protocols for the View.
	// 'kAll' enables protocol access to following three views: NFS, SMB and S3.
	// 'kNFSOnly' enables protocol access to NFS only.
	// 'kSMBOnly' enables protocol access to SMB only.
	// 'kS3Only' enables protocol access to S3 only.
	// 'kSwiftOnly' enables protocol access to Swift only.
	// 'kUnknown' indicates that the protocol access of a view does not match
	// any of the above. In this case, the constant is used as 'catch-all'.
	// Enum: ["kAll","kNFSOnly","kSMBOnly","kS3Only","kSwiftOnly","kUnknown"]
	ProtocolAccess *string `json:"protocolAccess,omitempty"`

	// Specifies the Quality of Service (QoS) Policy for the View.
	Qos *QoS `json:"qos,omitempty"`

	// Specifies the path to access this View as an S3 share.
	S3AccessPath *string `json:"s3AccessPath,omitempty"`

	// Specifies whether to support s3 folder support feature on the view. This
	// parameter can only be set during create and cannot be changed.
	S3FolderSupportEnabled *bool `json:"s3FolderSupportEnabled,omitempty"`

	// Specifies the S3 key mapping config of the view. This parameter can only
	// be set during create and cannot be changed.
	// Configuration of S3 key mapping.
	//
	// Specifies the type of S3 key mapping config.
	// Enum: ["kRandom","kShort","kLong","kHierarchical","kObjectId"]
	S3KeyMappingConfig *string `json:"s3KeyMappingConfig,omitempty"`

	// Specifies the security mode used for this view.
	// Currently we support the following modes: Native, Unified and NTFS style.
	// 'kNativeMode' indicates a native security mode.
	// 'kUnifiedMode' indicates a unified security mode.
	// 'kNtfsMode' indicates a NTFS style security mode.
	// Enum: ["kNativeMode","kUnifiedMode","kNtfsMode"]
	SecurityMode *string `json:"securityMode,omitempty"`

	// Specifies a list of share level permissions.
	SharePermissions []*SmbPermission `json:"sharePermissions"`

	// Specifies the main path for mounting this View as an SMB share.
	SmbMountPath *string `json:"smbMountPath,omitempty"`

	// Specifies the SMB permissions for the View.
	SmbPermissionsInfo *SmbPermissionsInfo `json:"smbPermissionsInfo,omitempty"`

	// Specifies statistics about the View.
	// readOnly: true
	Stats *ViewStats `json:"stats,omitempty"`

	// Specifies if inline deduplication and compression settings inherited from
	// the Storage Domain (View Box) should be disabled for this View.
	StoragePolicyOverride *StoragePolicyOverride `json:"storagePolicyOverride,omitempty"`

	// Array of Subnets.
	//
	// Specifies a list of Subnets with IP addresses that have permissions to
	// access the View. (Overrides the Subnets specified at the global
	// Cohesity Cluster level.)
	SubnetWhitelist []*Subnet `json:"subnetWhitelist"`

	// Specifies a list of user sids who have Superuser access to this view.
	SuperUserSids []string `json:"superUserSids"`

	// Specifies the Keystone project domain.
	SwiftProjectDomain *string `json:"swiftProjectDomain,omitempty"`

	// Specifies the Keystone project name.
	SwiftProjectName *string `json:"swiftProjectName,omitempty"`

	// Specifies the Keystone user domain.
	SwiftUserDomain *string `json:"swiftUserDomain,omitempty"`

	// Specifies the Keystone username.
	SwiftUsername *string `json:"swiftUsername,omitempty"`

	// Optional tenant id who has access to this View.
	TenantID *string `json:"tenantId,omitempty"`

	// Specifies the id of the Storage Domain (View Box) where the View is stored.
	ViewBoxID *int64 `json:"viewBoxId,omitempty"`

	// Specifies the name of the Storage Domain (View Box) where the View is stored.
	ViewBoxName *string `json:"viewBoxName,omitempty"`

	// Specifies an id of the View assigned by the Cohesity Cluster.
	ViewID *int64 `json:"viewId,omitempty"`

	// Specifies whether view lock is enabled. If enabled the view cannot be
	// modified or deleted until unlock. By default it is disabled.
	ViewLockEnabled *bool `json:"viewLockEnabled,omitempty"`

	// Specifies the pinning config of this view.
	ViewPinningConfig *ViewPinningConfig `json:"viewPinningConfig,omitempty"`

	// Specifies information about the Protection Jobs protecting this View.
	ViewProtection *ViewProtection `json:"viewProtection,omitempty"`
}

// Validate validates this view
func (m *View) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAliases(formats); err != nil {
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

	if err := m.validateIntent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogicalQuota(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetgroupWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsAllSquash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsRootPermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNfsRootSquash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQos(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3KeyMappingConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharePermissions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmbPermissionsInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStoragePolicyOverride(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewPinningConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewProtection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *View) validateAliases(formats strfmt.Registry) error {
	if swag.IsZero(m.Aliases) { // not required
		return nil
	}

	for i := 0; i < len(m.Aliases); i++ {
		if swag.IsZero(m.Aliases[i]) { // not required
			continue
		}

		if m.Aliases[i] != nil {
			if err := m.Aliases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *View) validateAntivirusScanConfig(formats strfmt.Registry) error {
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

func (m *View) validateFileExtensionFilter(formats strfmt.Registry) error {
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

func (m *View) validateFileLockConfig(formats strfmt.Registry) error {
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

func (m *View) validateIntent(formats strfmt.Registry) error {
	if swag.IsZero(m.Intent) { // not required
		return nil
	}

	if m.Intent != nil {
		if err := m.Intent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("intent")
			}
			return err
		}
	}

	return nil
}

func (m *View) validateLogicalQuota(formats strfmt.Registry) error {
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

func (m *View) validateNetgroupWhitelist(formats strfmt.Registry) error {
	if swag.IsZero(m.NetgroupWhitelist) { // not required
		return nil
	}

	for i := 0; i < len(m.NetgroupWhitelist); i++ {
		if swag.IsZero(m.NetgroupWhitelist[i]) { // not required
			continue
		}

		if m.NetgroupWhitelist[i] != nil {
			if err := m.NetgroupWhitelist[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netgroupWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netgroupWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *View) validateNfsAllSquash(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsAllSquash) { // not required
		return nil
	}

	if m.NfsAllSquash != nil {
		if err := m.NfsAllSquash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsAllSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsAllSquash")
			}
			return err
		}
	}

	return nil
}

func (m *View) validateNfsRootPermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsRootPermissions) { // not required
		return nil
	}

	if m.NfsRootPermissions != nil {
		if err := m.NfsRootPermissions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootPermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootPermissions")
			}
			return err
		}
	}

	return nil
}

func (m *View) validateNfsRootSquash(formats strfmt.Registry) error {
	if swag.IsZero(m.NfsRootSquash) { // not required
		return nil
	}

	if m.NfsRootSquash != nil {
		if err := m.NfsRootSquash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootSquash")
			}
			return err
		}
	}

	return nil
}

var viewTypeProtocolAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kAll","kNFSOnly","kSMBOnly","kS3Only","kSwiftOnly","kUnknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewTypeProtocolAccessPropEnum = append(viewTypeProtocolAccessPropEnum, v)
	}
}

const (

	// ViewProtocolAccessKAll captures enum value "kAll"
	ViewProtocolAccessKAll string = "kAll"

	// ViewProtocolAccessKNFSOnly captures enum value "kNFSOnly"
	ViewProtocolAccessKNFSOnly string = "kNFSOnly"

	// ViewProtocolAccessKSMBOnly captures enum value "kSMBOnly"
	ViewProtocolAccessKSMBOnly string = "kSMBOnly"

	// ViewProtocolAccessKS3Only captures enum value "kS3Only"
	ViewProtocolAccessKS3Only string = "kS3Only"

	// ViewProtocolAccessKSwiftOnly captures enum value "kSwiftOnly"
	ViewProtocolAccessKSwiftOnly string = "kSwiftOnly"

	// ViewProtocolAccessKUnknown captures enum value "kUnknown"
	ViewProtocolAccessKUnknown string = "kUnknown"
)

// prop value enum
func (m *View) validateProtocolAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewTypeProtocolAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *View) validateProtocolAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtocolAccess) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolAccessEnum("protocolAccess", "body", *m.ProtocolAccess); err != nil {
		return err
	}

	return nil
}

func (m *View) validateQos(formats strfmt.Registry) error {
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

var viewTypeS3KeyMappingConfigPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kRandom","kShort","kLong","kHierarchical","kObjectId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewTypeS3KeyMappingConfigPropEnum = append(viewTypeS3KeyMappingConfigPropEnum, v)
	}
}

const (

	// ViewS3KeyMappingConfigKRandom captures enum value "kRandom"
	ViewS3KeyMappingConfigKRandom string = "kRandom"

	// ViewS3KeyMappingConfigKShort captures enum value "kShort"
	ViewS3KeyMappingConfigKShort string = "kShort"

	// ViewS3KeyMappingConfigKLong captures enum value "kLong"
	ViewS3KeyMappingConfigKLong string = "kLong"

	// ViewS3KeyMappingConfigKHierarchical captures enum value "kHierarchical"
	ViewS3KeyMappingConfigKHierarchical string = "kHierarchical"

	// ViewS3KeyMappingConfigKObjectID captures enum value "kObjectId"
	ViewS3KeyMappingConfigKObjectID string = "kObjectId"
)

// prop value enum
func (m *View) validateS3KeyMappingConfigEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewTypeS3KeyMappingConfigPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *View) validateS3KeyMappingConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.S3KeyMappingConfig) { // not required
		return nil
	}

	// value enum
	if err := m.validateS3KeyMappingConfigEnum("s3KeyMappingConfig", "body", *m.S3KeyMappingConfig); err != nil {
		return err
	}

	return nil
}

var viewTypeSecurityModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["kNativeMode","kUnifiedMode","kNtfsMode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		viewTypeSecurityModePropEnum = append(viewTypeSecurityModePropEnum, v)
	}
}

const (

	// ViewSecurityModeKNativeMode captures enum value "kNativeMode"
	ViewSecurityModeKNativeMode string = "kNativeMode"

	// ViewSecurityModeKUnifiedMode captures enum value "kUnifiedMode"
	ViewSecurityModeKUnifiedMode string = "kUnifiedMode"

	// ViewSecurityModeKNtfsMode captures enum value "kNtfsMode"
	ViewSecurityModeKNtfsMode string = "kNtfsMode"
)

// prop value enum
func (m *View) validateSecurityModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, viewTypeSecurityModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *View) validateSecurityMode(formats strfmt.Registry) error {
	if swag.IsZero(m.SecurityMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateSecurityModeEnum("securityMode", "body", *m.SecurityMode); err != nil {
		return err
	}

	return nil
}

func (m *View) validateSharePermissions(formats strfmt.Registry) error {
	if swag.IsZero(m.SharePermissions) { // not required
		return nil
	}

	for i := 0; i < len(m.SharePermissions); i++ {
		if swag.IsZero(m.SharePermissions[i]) { // not required
			continue
		}

		if m.SharePermissions[i] != nil {
			if err := m.SharePermissions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *View) validateSmbPermissionsInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.SmbPermissionsInfo) { // not required
		return nil
	}

	if m.SmbPermissionsInfo != nil {
		if err := m.SmbPermissionsInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smbPermissionsInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smbPermissionsInfo")
			}
			return err
		}
	}

	return nil
}

func (m *View) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *View) validateStoragePolicyOverride(formats strfmt.Registry) error {
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

func (m *View) validateSubnetWhitelist(formats strfmt.Registry) error {
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

func (m *View) validateViewPinningConfig(formats strfmt.Registry) error {
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

func (m *View) validateViewProtection(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewProtection) { // not required
		return nil
	}

	if m.ViewProtection != nil {
		if err := m.ViewProtection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewProtection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewProtection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this view based on the context it is used
func (m *View) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAliases(ctx, formats); err != nil {
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

	if err := m.contextValidateIntent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLogicalQuota(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetgroupWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsAllSquash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsRootPermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNfsRootSquash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQos(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharePermissions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSmbPermissionsInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStoragePolicyOverride(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubnetWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewPinningConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateViewProtection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *View) contextValidateAliases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Aliases); i++ {

		if m.Aliases[i] != nil {

			if swag.IsZero(m.Aliases[i]) { // not required
				return nil
			}

			if err := m.Aliases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aliases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aliases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *View) contextValidateAntivirusScanConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateFileExtensionFilter(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateFileLockConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateIntent(ctx context.Context, formats strfmt.Registry) error {

	if m.Intent != nil {

		if swag.IsZero(m.Intent) { // not required
			return nil
		}

		if err := m.Intent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("intent")
			}
			return err
		}
	}

	return nil
}

func (m *View) contextValidateLogicalQuota(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateNetgroupWhitelist(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetgroupWhitelist); i++ {

		if m.NetgroupWhitelist[i] != nil {

			if swag.IsZero(m.NetgroupWhitelist[i]) { // not required
				return nil
			}

			if err := m.NetgroupWhitelist[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("netgroupWhitelist" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("netgroupWhitelist" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *View) contextValidateNfsAllSquash(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsAllSquash != nil {

		if swag.IsZero(m.NfsAllSquash) { // not required
			return nil
		}

		if err := m.NfsAllSquash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsAllSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsAllSquash")
			}
			return err
		}
	}

	return nil
}

func (m *View) contextValidateNfsRootPermissions(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsRootPermissions != nil {

		if swag.IsZero(m.NfsRootPermissions) { // not required
			return nil
		}

		if err := m.NfsRootPermissions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootPermissions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootPermissions")
			}
			return err
		}
	}

	return nil
}

func (m *View) contextValidateNfsRootSquash(ctx context.Context, formats strfmt.Registry) error {

	if m.NfsRootSquash != nil {

		if swag.IsZero(m.NfsRootSquash) { // not required
			return nil
		}

		if err := m.NfsRootSquash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nfsRootSquash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nfsRootSquash")
			}
			return err
		}
	}

	return nil
}

func (m *View) contextValidateQos(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateSharePermissions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SharePermissions); i++ {

		if m.SharePermissions[i] != nil {

			if swag.IsZero(m.SharePermissions[i]) { // not required
				return nil
			}

			if err := m.SharePermissions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sharePermissions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *View) contextValidateSmbPermissionsInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.SmbPermissionsInfo != nil {

		if swag.IsZero(m.SmbPermissionsInfo) { // not required
			return nil
		}

		if err := m.SmbPermissionsInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("smbPermissionsInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("smbPermissionsInfo")
			}
			return err
		}
	}

	return nil
}

func (m *View) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {

	if m.Stats != nil {

		if swag.IsZero(m.Stats) { // not required
			return nil
		}

		if err := m.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

func (m *View) contextValidateStoragePolicyOverride(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateSubnetWhitelist(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateViewPinningConfig(ctx context.Context, formats strfmt.Registry) error {

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

func (m *View) contextValidateViewProtection(ctx context.Context, formats strfmt.Registry) error {

	if m.ViewProtection != nil {

		if swag.IsZero(m.ViewProtection) { // not required
			return nil
		}

		if err := m.ViewProtection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("viewProtection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("viewProtection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *View) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *View) UnmarshalBinary(b []byte) error {
	var res View
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

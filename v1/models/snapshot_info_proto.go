// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SnapshotInfoProto Message that encapsulates information about snapshots for any of the
// environments we support. Environment specific snapshot infos are defined as
// extensions to this proto.
//
// Each available extension is listed below along with the location of the
// proto file (relative to magneto/connectors) where it is defined. The only
// exception is view.proto and physical.proto which reside in magneto/base.
//
// SnapshotInfoProto extension                     Location              Extn
// =============================================================================
// vmware::SnapshotInfo::vmware_snapshot_info     vmware/vmware.proto       100
// sql::SnapshotInfo::sql_snapshot_info           sql/sql.proto             101
// view::SnapshotInfo::view_snapshot_info         base/view.proto           102
// physical::SnapshotInfo::physical_snapshot_info base/physical.proto       103
// san::SnapshotInfo::san_snapshot_info           san/san.proto             104
// file::SnapshotInfo::file_snapshot_info         file/file.proto           105
// hyperv::SnapshotInfo::hyperv_snapshot_info     hyperv/hyperv.proto       106
// acropolis::SnapshotInfo::
// acropolis_snapshot_info                        acropolis/acropolis.proto 107
// kvm::SnapshotInfo::kvm_snapshot_info           kvm/kvm.proto             108
// app_file::SnapshotInfo::app_file_snapshot_info app_file/app_file.proto   109
// oracle::SnapshotInfo::oracle_snapshot_info     oracle/oracle.proto       110
// aws::SnapshotInfo::aws_snapshot_info           aws/aws.proto             111
// outlook::SnapshotInfo::outlook_snapshot_info   outlook/outlook.proto     112
// azure::SnapshotInfo::azure_snapshot_info       azure/azure.proto         113
// gcp::SnapshotInfo::gcp_snapshot_info           gcp/gcp.proto             114
// ad::SnapshotInfo::ad_snapshot_info             ad/ad.proto               115
// MSGraph::SnapshotInfo::one_drive_snapshot_info ms_graph/graph.proto      116
// kubernetes::SnapshotInfo:: kubernetes_snapshot_info
// kubernetes/kubernetes.proto
// 117
// aws::RDSSnapshotInfo::rds_snapshot_info        aws/aws.proto             118
// o365::SnapshotInfo::o365_snapshot_info         o365/o365.proto           119
// exchange::SnapshotInfo::exchange_snapshot_info exchange/exchange.proto   120
// o365::SharepointSnapshotInfo::sharepoint_snapshot_info
// o365/o365.proto           121
// MSGraph::SharepointListSnapshotInfo::sharepoint_list_snapshot_info
// ms_graph/graph.proto      122
// cdp::SnapshotInfo::cdp_snapshot_info           base/cdp.proto            123
// imanis::SnapshotInfo::nosql_snapshot_info      imanis/nosql.proto        124
// o365::PublicFolderSnapshotInfo::public_folder_snapshot_info
// o365/o365.proto           125
// SnapshotInfo::uda_snapshot_info                uda.proto                 126
// o365::TeamsSnapshotInfo::teams_snapshot_info   o365/o365.proto           127
// o365::O365GroupSnapshotInfo::o365_group_snapshot_info
// o365/o365.proto           128
// SnapshotInfo::sfdc_snapshot_info               sfdc_service.proto        129
// san::GroupSnapshotInfo::group_snapshot_info    san/san.proto             130
// rds::OracleRmanSnapshotInfo::rds_oracle_rman_snapshot_info
// rds/rds.proto             131
// o365::ChatsSnapshotInfo::chats_snapshot_info   o365/o365.proto           132
// =============================================================================
//
// swagger:model SnapshotInfoProto
type SnapshotInfoProto struct {

	// The name of the rocksdb directory for writing high change directories.
	// It is stored in 'config' directory of the current view.
	ChangeRocksdbName *string `json:"changeRocksdbName,omitempty"`

	// The name of the rocksdb directory for errors seen during this backup,
	// stored in 'config' directory of the current view.
	ErrorRocksdbName *string `json:"errorRocksdbName,omitempty"`

	// This field is only applicable for NAS and file backup jobs. It indicates
	// whether the file walk portion of the backup has completed.
	// TODO(amandeep): Rename this as this can be used for multiple workflows.
	FileWalkDone *bool `json:"fileWalkDone,omitempty"`

	// Front end size information. An example use case is for billing purposes
	// in "[Backup | Data Management] as a Service" offering.
	FrontEndSizeInfo *SizeInfo `json:"frontEndSizeInfo,omitempty"`

	// The metadata view name in which the backup metadata was created.
	// NOTE: This is populated only for CADv2 NAS backup.
	MetadataViewName *string `json:"metadataViewName,omitempty"`

	// Number of application instances backed up by this task. For example, if
	// the environment type is kSQL, this number is for the SQL server instances.
	NumAppInstances *int32 `json:"numAppInstances,omitempty"`

	// Number of application objects in total backed up by this task. For
	// example, if the environment type is kSQL, this number is for all of the
	// SQL server databases
	NumAppObjects *int32 `json:"numAppObjects,omitempty"`

	// Captures the execution status of post backup script.
	PostBackupScriptStatus *ScriptExecutionStatus `json:"postBackupScriptStatus,omitempty"`

	// Captures the execution status of post snapshot script.
	PostSnapshotScriptStatus *ScriptExecutionStatus `json:"postSnapshotScriptStatus,omitempty"`

	// Captures the execution status of pre backup script.
	PreBackupScriptStatus *ScriptExecutionStatus `json:"preBackupScriptStatus,omitempty"`

	// If the permit of this task is released on pausing backup, this boolean
	// informs the task to re-acquire it.
	ReacquirePermit *bool `json:"reacquirePermit,omitempty"`

	// This is the path relative to 'root_path' under which the snapshot lives.
	// This does not begin with a '/' and is of the form foo/bar/baz.
	RelativeSnapshotDir *string `json:"relativeSnapshotDir,omitempty"`

	// The root path under which the snapshot is stored. This is of the form
	// "/ViewBox/ViewName/fs".
	RootPath *string `json:"rootPath,omitempty"`

	// If this backup task stores any auxiliary state in Scribe table, this field
	// will be populated with the column key in that table where such state is
	// stored. Data stored in the column is extension of SnapshotScribeInfoProto
	// message.
	ScribeTableColumn *string `json:"scribeTableColumn,omitempty"`

	// If this backup task stores any auxiliary state in Scribe table, this field
	// will be populated with the row key in that table where such state is
	// stored.
	ScribeTableRow *string `json:"scribeTableRow,omitempty"`

	// This is the timestamp at which the slave task started.
	SlaveTaskStartTimeUsecs *int64 `json:"slaveTaskStartTimeUsecs,omitempty"`

	// Snapshot expiry time.
	SnapshotExpiryTime *uint64 `json:"snapshotExpiryTime,omitempty"`

	// Captures the snapshot type for some objects such as VM.
	SnapshotType *ObjectSnapshotType `json:"snapshotType,omitempty"`

	// The source snapshot create time.
	SourceSnapshotCreateTimeUsecs *int64 `json:"sourceSnapshotCreateTimeUsecs,omitempty"`

	// This field is only relevant for NAS backups where we are backing up from a
	// ReadOnly or DataProtect volume, or an RW volume with Snapdiff. For RW
	// volumes, Cohesity will create the snapshot on NetApp, while for DP
	// volumes, Cohesity will use the existing snapshot.
	SourceSnapshotName *string `json:"sourceSnapshotName,omitempty"`

	// Indicates the state of the source snapshot if it is being managed by the
	// master op. 'source_snapshot_name' will be set to indicate the snapshot
	// name. At the moment, this feature is enabled only for Netapp & Isilon
	// adapters to support continuous snapshotting feature.
	SourceSnapshotStatus *int32 `json:"sourceSnapshotStatus,omitempty"`

	// Specifies the parameters required for Storage Snapshot provider.
	StorageSnapshotProvider *StorageSnapshotProviderParams `json:"storageSnapshotProvider,omitempty"`

	// Specifies the target type for the task. The field is only valid if the
	// task has got a permit.
	TargetType *int32 `json:"targetType,omitempty"`

	// Contains the information regarding number of bytes that are read from the
	// source (such as VM) so far.
	TotalBytesReadFromSource *int64 `json:"totalBytesReadFromSource,omitempty"`

	// Total amount of data successfully tiered from the NAS source.
	TotalBytesTiered *int64 `json:"totalBytesTiered,omitempty"`

	// Contains the total number of bytes that will be read from the source
	// (such as VM) for this snapshot.
	TotalBytesToReadFromSource *int64 `json:"totalBytesToReadFromSource,omitempty"`

	// The total number of file and directory entities that have changed since
	// last backup. Only applicable to file based backups.
	TotalChangedEntityCount *int64 `json:"totalChangedEntityCount,omitempty"`

	// The total number of file and directory entities visited in this
	// backup. Only applicable to file based backups.
	TotalEntityCount *int64 `json:"totalEntityCount,omitempty"`

	// Logical size of the source whose snapshot is being taken. This is the
	// amount of data we would have read from the source had this been a full
	// backup.
	TotalLogicalBackupSizeBytes *int64 `json:"totalLogicalBackupSizeBytes,omitempty"`

	// Contains the information regarding number of bytes that the source (such
	// as VM) has taken up on the primary storage.
	TotalPrimaryPhysicalSizeBytes *int64 `json:"totalPrimaryPhysicalSizeBytes,omitempty"`

	// Total number of bytes to be zero-filled in Snap FS as part of this backup.
	// Currently applicable only for VMware backups.
	TotalZeroFillBytes *int64 `json:"totalZeroFillBytes,omitempty"`

	// The type of environment this snapshot info pertains to.
	Type *int32 `json:"type,omitempty"`

	// Whether during the backup, the backup view's case insensitivity property
	// has been altered. If so, Madrox needs to take corresponding actions during
	// replication.
	ViewCaseInsensitivityAltered *bool `json:"viewCaseInsensitivityAltered,omitempty"`

	// The data view name under which the snapshot was created.
	// NOTE: This is populated only for View, Puppeteer, NAS and Oracle backup.
	ViewName *string `json:"viewName,omitempty"`

	// The view name under which the snapshot of the migrated data was created.
	// NOTE: This is populated only for data migration tasks.
	ViewNameToGc *string `json:"viewNameToGc,omitempty"`

	// Warnings if any. These warnings will be propogated to the UI by master.
	Warnings []*ErrorProto `json:"warnings"`

	// Factor by which weight of a zero-fill sub-task should be scaled down. This
	// is used while creating sub-task monitors.
	ZeroFillTaskWeightScaleDownFactor *int64 `json:"zeroFillTaskWeightScaleDownFactor,omitempty"`
}

// Validate validates this snapshot info proto
func (m *SnapshotInfoProto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrontEndSizeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostBackupScriptStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePostSnapshotScriptStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreBackupScriptStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageSnapshotProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarnings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotInfoProto) validateFrontEndSizeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.FrontEndSizeInfo) { // not required
		return nil
	}

	if m.FrontEndSizeInfo != nil {
		if err := m.FrontEndSizeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frontEndSizeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("frontEndSizeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) validatePostBackupScriptStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PostBackupScriptStatus) { // not required
		return nil
	}

	if m.PostBackupScriptStatus != nil {
		if err := m.PostBackupScriptStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postBackupScriptStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postBackupScriptStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) validatePostSnapshotScriptStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PostSnapshotScriptStatus) { // not required
		return nil
	}

	if m.PostSnapshotScriptStatus != nil {
		if err := m.PostSnapshotScriptStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postSnapshotScriptStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postSnapshotScriptStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) validatePreBackupScriptStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PreBackupScriptStatus) { // not required
		return nil
	}

	if m.PreBackupScriptStatus != nil {
		if err := m.PreBackupScriptStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preBackupScriptStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preBackupScriptStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) validateSnapshotType(formats strfmt.Registry) error {
	if swag.IsZero(m.SnapshotType) { // not required
		return nil
	}

	if m.SnapshotType != nil {
		if err := m.SnapshotType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotType")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) validateStorageSnapshotProvider(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageSnapshotProvider) { // not required
		return nil
	}

	if m.StorageSnapshotProvider != nil {
		if err := m.StorageSnapshotProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageSnapshotProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageSnapshotProvider")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) validateWarnings(formats strfmt.Registry) error {
	if swag.IsZero(m.Warnings) { // not required
		return nil
	}

	for i := 0; i < len(m.Warnings); i++ {
		if swag.IsZero(m.Warnings[i]) { // not required
			continue
		}

		if m.Warnings[i] != nil {
			if err := m.Warnings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this snapshot info proto based on the context it is used
func (m *SnapshotInfoProto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFrontEndSizeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostBackupScriptStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePostSnapshotScriptStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreBackupScriptStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSnapshotType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageSnapshotProvider(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarnings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotInfoProto) contextValidateFrontEndSizeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.FrontEndSizeInfo != nil {

		if swag.IsZero(m.FrontEndSizeInfo) { // not required
			return nil
		}

		if err := m.FrontEndSizeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("frontEndSizeInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("frontEndSizeInfo")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) contextValidatePostBackupScriptStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PostBackupScriptStatus != nil {

		if swag.IsZero(m.PostBackupScriptStatus) { // not required
			return nil
		}

		if err := m.PostBackupScriptStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postBackupScriptStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postBackupScriptStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) contextValidatePostSnapshotScriptStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PostSnapshotScriptStatus != nil {

		if swag.IsZero(m.PostSnapshotScriptStatus) { // not required
			return nil
		}

		if err := m.PostSnapshotScriptStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postSnapshotScriptStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postSnapshotScriptStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) contextValidatePreBackupScriptStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.PreBackupScriptStatus != nil {

		if swag.IsZero(m.PreBackupScriptStatus) { // not required
			return nil
		}

		if err := m.PreBackupScriptStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preBackupScriptStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preBackupScriptStatus")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) contextValidateSnapshotType(ctx context.Context, formats strfmt.Registry) error {

	if m.SnapshotType != nil {

		if swag.IsZero(m.SnapshotType) { // not required
			return nil
		}

		if err := m.SnapshotType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("snapshotType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("snapshotType")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) contextValidateStorageSnapshotProvider(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageSnapshotProvider != nil {

		if swag.IsZero(m.StorageSnapshotProvider) { // not required
			return nil
		}

		if err := m.StorageSnapshotProvider.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageSnapshotProvider")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageSnapshotProvider")
			}
			return err
		}
	}

	return nil
}

func (m *SnapshotInfoProto) contextValidateWarnings(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Warnings); i++ {

		if m.Warnings[i] != nil {

			if swag.IsZero(m.Warnings[i]) { // not required
				return nil
			}

			if err := m.Warnings[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warnings" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("warnings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotInfoProto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotInfoProto) UnmarshalBinary(b []byte) error {
	var res SnapshotInfoProto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

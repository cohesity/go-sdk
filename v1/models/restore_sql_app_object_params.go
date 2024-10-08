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

// RestoreSQLAppObjectParams The application object restore params for SQL objects (SQL databases).
//
// swagger:model RestoreSqlAppObjectParams
type RestoreSQLAppObjectParams struct {

	// Set to true if tail logs are to be captured before the restore
	// operation. This is only applicable if we are restoring the SQL database
	// to its original source, and the database is not being renamed.
	CaptureTailLogs *bool `json:"captureTailLogs,omitempty"`

	// Whether restore should continue after encountering a page checksum error.
	ContinueAfterError *bool `json:"continueAfterError,omitempty"`

	// Which directory to put the database data files. Missing directory will be
	// automatically created. Cannot be empty if not restoring to the original
	// SQL instance.
	DataFileDestination *string `json:"dataFileDestination,omitempty"`

	// Policy to overwrite an existing DB during a restore operation.
	DbRestoreOverwritePolicy *int32 `json:"dbRestoreOverwritePolicy,omitempty"`

	// Whether restore checksums are enabled.
	EnableChecksum *bool `json:"enableChecksum,omitempty"`

	// The name of the SQL instance that we restore database to. If target_host
	// is not empty, this also cannot be empty.
	InstanceName *string `json:"instanceName,omitempty"`

	// The following field is set if auto_sync for multi-stage SQL restore
	// task is enabled. This field is valid only if is_multi_state_restore
	// is set to true.
	IsAutoSyncEnabled *bool `json:"isAutoSyncEnabled,omitempty"`

	// The following field is set if we are creating a multi-stage SQL restore
	// task needed for features such as Hot-Standby.
	IsMultiStageRestore *bool `json:"isMultiStageRestore,omitempty"`

	// Set to true to keep cdc on restored database.
	KeepCdc *bool `json:"keepCdc,omitempty"`

	// Which directory to put the database log files. Missing directory will be
	// automatically created. Cannot be empty if not restoring to the original
	// SQL instance.
	LogFileDestination *string `json:"logFileDestination,omitempty"`

	// The following field is set if this is a sub task for a multi-stage SQL
	// restore task. It captures the options specified for this sub-task.
	//
	// Note that this field is set internally by Magneto, and should not be set
	// by Iris.
	MultiStageRestoreOptions *SQLUpdateRestoreTaskOptions `json:"multiStageRestoreOptions,omitempty"`

	// The new name of the database, if it is going to be renamed. app_entity in
	// RestoreAppObject has to be non-empty for the renaming, otherwise it does
	// not make sense to rename all databases in the owner.
	NewDatabaseName *string `json:"newDatabaseName,omitempty"`

	// If this is set to true, we will replay the entire last log without STOPAT.
	ReplayEntireLastLog *bool `json:"replayEntireLastLog,omitempty"`

	// The time to which the SQL database needs to be restored. This allows for
	// granular recovery of SQL databases. If this is not set, the SQL database
	// will be recovered to the full/incremental snapshot (specified in the
	// owner's restore object in AppOwnerRestoreInfo).
	RestoreTimeSecs *int64 `json:"restoreTimeSecs,omitempty"`

	// Resume restore if sql instance/database exist in restore/recovering state.
	// The database might be in restore/recovering state if previous restore
	// failed or previous  restore was attempted  with norecovery option.
	ResumeRestore *bool `json:"resumeRestore,omitempty"`

	// Which directory to put the secondary data files of the database. Secondary
	// data files are optional and are user defined. The recommended file name
	// extension for these is ".ndf".
	//
	// If this option is specified, the directory will be automatically created
	// if its missing.
	SecondaryDataFileDestination *string `json:"secondaryDataFileDestination,omitempty"`

	// Specify the secondary data files and corresponding direcories of the DB.
	// Secondary data files are optional and are user defined. The recommended
	// file extension for secondary files is ".ndf".
	//
	// If this option is specified and the destination folders do not exist they
	// will be automatically created.
	SecondaryDataFileDestinationVec []*FilesToDirectoryMapping `json:"secondaryDataFileDestinationVec"`

	// 'with_clause' contains 'with clause' to be used in native sql restore
	// command. This is only applicable for db restore of native sql backup. Here
	// user can specify multiple restore options. Example: "WITH BUFFERCOUNT =
	// 575, MAXTRANSFERSIZE = 2097152". If this is not specified, we use the
	// value specified in magneto_sql_native_restore_with_clause gflag.
	WithClause *string `json:"withClause,omitempty"`

	// Set to true if we want to recover the database in "NO_RECOVERY" mode
	// which does not bring it online after restore.
	WithNoRecovery *bool `json:"withNoRecovery,omitempty"`
}

// Validate validates this restore Sql app object params
func (m *RestoreSQLAppObjectParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMultiStageRestoreOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryDataFileDestinationVec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreSQLAppObjectParams) validateMultiStageRestoreOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.MultiStageRestoreOptions) { // not required
		return nil
	}

	if m.MultiStageRestoreOptions != nil {
		if err := m.MultiStageRestoreOptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multiStageRestoreOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multiStageRestoreOptions")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreSQLAppObjectParams) validateSecondaryDataFileDestinationVec(formats strfmt.Registry) error {
	if swag.IsZero(m.SecondaryDataFileDestinationVec) { // not required
		return nil
	}

	for i := 0; i < len(m.SecondaryDataFileDestinationVec); i++ {
		if swag.IsZero(m.SecondaryDataFileDestinationVec[i]) { // not required
			continue
		}

		if m.SecondaryDataFileDestinationVec[i] != nil {
			if err := m.SecondaryDataFileDestinationVec[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondaryDataFileDestinationVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secondaryDataFileDestinationVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this restore Sql app object params based on the context it is used
func (m *RestoreSQLAppObjectParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMultiStageRestoreOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSecondaryDataFileDestinationVec(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreSQLAppObjectParams) contextValidateMultiStageRestoreOptions(ctx context.Context, formats strfmt.Registry) error {

	if m.MultiStageRestoreOptions != nil {

		if swag.IsZero(m.MultiStageRestoreOptions) { // not required
			return nil
		}

		if err := m.MultiStageRestoreOptions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("multiStageRestoreOptions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("multiStageRestoreOptions")
			}
			return err
		}
	}

	return nil
}

func (m *RestoreSQLAppObjectParams) contextValidateSecondaryDataFileDestinationVec(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SecondaryDataFileDestinationVec); i++ {

		if m.SecondaryDataFileDestinationVec[i] != nil {

			if swag.IsZero(m.SecondaryDataFileDestinationVec[i]) { // not required
				return nil
			}

			if err := m.SecondaryDataFileDestinationVec[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondaryDataFileDestinationVec" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("secondaryDataFileDestinationVec" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreSQLAppObjectParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreSQLAppObjectParams) UnmarshalBinary(b []byte) error {
	var res RestoreSQLAppObjectParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeLogInfo Information about an Exchange Database log file group and checkpoint file.
//
// swagger:model ExchangeLogInfo
type ExchangeLogInfo struct {

	// Checkpoint file name that lives in log folder. This file contains
	// information about log generation commited to EDB. This file is named as
	// <file_prefix>.chk. Eg: E00.chk.
	CheckpointFileName *string `json:"checkpointFileName,omitempty"`

	// Is circular logging enabled for this database. If this is true, then
	// incremental backup should not be performed. Only full backup should
	// be supported.
	CircularLogging *bool `json:"circularLogging,omitempty"`

	// Log file prefix for a database. Eg: E00.
	// The log file themselves have a file extension of .LOG. Log files are named
	// with a numeric sequence number. Example E00000006A0.log. There may be
	// temp log file such as  E00tmp.log that is being populated.
	FilePrefix *string `json:"filePrefix,omitempty"`

	// Windows file system type string such as 'NTFS,ReFS,CSVFS' of the 'volume'.
	// Exchange VSS has restriction that CSVFS and non-CSVFS volumes cannot be
	// mixed in same snapshot set.
	FileSystem *string `json:"fileSystem,omitempty"`

	// LOG files folder path.  This file path will be common for all nodes within
	// a DAG.
	FolderPath *string `json:"folderPath,omitempty"`

	// This is per log file size in bytes. This is NOT the total size
	// of all log files. Default size is 1MiB.
	SingleLogSize *uint32 `json:"singleLogSize,omitempty"`

	// Approximate total size of LOG directory. It gets total number of files in
	// 'folder_path' and multiplies it with 1MiB. Some files (.CHK) are smaller.
	// Accurate size computation may take a long time, so its not performed. This
	// approximate size is sufficient for reporting purposes. If -1, it could not
	// be computed or resulted in error.
	TotalSize *int64 `json:"totalSize,omitempty"`

	// Volume device id where the log folder mount point lives. This is unique to
	// an Exchange server in DAG, so each DAG node will have different volume id
	// that should be included in VSS snapshot.
	// Example: \\?\Volume{63f6b242-8ec1-11e8-80e0-00505685e2a9}\
	Volume *string `json:"volume,omitempty"`
}

// Validate validates this exchange log info
func (m *ExchangeLogInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exchange log info based on context it is used
func (m *ExchangeLogInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeLogInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeLogInfo) UnmarshalBinary(b []byte) error {
	var res ExchangeLogInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

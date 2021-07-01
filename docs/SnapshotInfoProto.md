# SnapshotInfoProto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorRocksdbName** | Pointer to **NullableString** | The name of the rocksdb directory for errors seen during this backup, stored in &#39;config&#39; directory of the current view. | [optional] 
**FileWalkDone** | Pointer to **NullableBool** | This field is only applicable for NAS and file backup jobs. It indicates whether the file walk portion of the backup has completed. | [optional] 
**NumAppInstances** | Pointer to **NullableInt32** | Number of application instances backed up by this task. For example, if the environment type is kSQL, this number is for the SQL server instances. | [optional] 
**NumAppObjects** | Pointer to **NullableInt32** | Number of application objects in total backed up by this task. For example, if the environment type is kSQL, this number is for all of the SQL server databases | [optional] 
**PostBackupScriptStatus** | Pointer to [**ScriptExecutionStatus**](ScriptExecutionStatus.md) |  | [optional] 
**PreBackupScriptStatus** | Pointer to [**ScriptExecutionStatus**](ScriptExecutionStatus.md) |  | [optional] 
**RelativeSnapshotDir** | Pointer to **NullableString** | This is the path relative to &#39;root_path&#39; under which the snapshot lives. This does not begin with a &#39;/&#39; and is of the form foo/bar/baz. | [optional] 
**RootPath** | Pointer to **NullableString** | The root path under which the snapshot is stored. This is of the form \&quot;/ViewBox/ViewName/fs\&quot;. | [optional] 
**ScribeTableColumn** | Pointer to **NullableString** | If this backup task stores any auxiliary state in Scribe table, this field will be populated with the column key in that table where such state is stored. Data stored in the column is extension of SnapshotScribeInfoProto message. | [optional] 
**ScribeTableRow** | Pointer to **NullableString** | If this backup task stores any auxiliary state in Scribe table, this field will be populated with the row key in that table where such state is stored. | [optional] 
**SlaveTaskStartTimeUsecs** | Pointer to **NullableInt64** | This is the timestamp at which the slave task started. | [optional] 
**SnapshotType** | Pointer to [**ObjectSnapshotType**](ObjectSnapshotType.md) |  | [optional] 
**SourceSnapshotCreateTimeUsecs** | Pointer to **NullableInt64** | The source snapshot create time. | [optional] 
**SourceSnapshotName** | Pointer to **NullableString** | This filed is only applicable for NAS when we do backup from Readonly/DataProtect volume where we use already created snapshot on the source. | [optional] 
**StorageSnapshotProvider** | Pointer to [**StorageSnapshotProviderParams**](StorageSnapshotProviderParams.md) |  | [optional] 
**TargetType** | Pointer to **NullableInt32** | Specifies the target type for the task. The field is only valid if the task has got a permit. | [optional] 
**TotalBytesReadFromSource** | Pointer to **NullableInt64** | Contains the information regarding number of bytes that are read from the source (such as VM) so far. | [optional] 
**TotalBytesToReadFromSource** | Pointer to **NullableInt64** | Contains the total number of bytes that will be read from the source (such as VM) for this snapshot. | [optional] 
**TotalChangedEntityCount** | Pointer to **NullableInt64** | The total number of file and directory entities that have changed since last backup. Only applicable to file based backups. | [optional] 
**TotalEntityCount** | Pointer to **NullableInt64** | The total number of file and directory entities visited in this backup. Only applicable to file based backups. | [optional] 
**TotalLogicalBackupSizeBytes** | Pointer to **NullableInt64** | Logical size of the source whose snapshot is being taken. This is the amount of data we would have read from the source had this been a full backup. | [optional] 
**TotalPrimaryPhysicalSizeBytes** | Pointer to **NullableInt64** | Contains the information regarding number of bytes that the source (such as VM) has taken up on the primary storage. | [optional] 
**Type** | Pointer to **NullableInt32** | The type of environment this snapshot info pertains to. | [optional] 
**ViewCaseInsensitivityAltered** | Pointer to **NullableBool** | Whether during the backup, the backup view&#39;s case insensitivity property has been altered. If so, Madrox needs to take corresponding actions during replication. | [optional] 
**ViewName** | Pointer to **NullableString** | The view name under which the snapshot was created. NOTE: This is populated only for View, Puppeteer, NAS and Oracle backup. | [optional] 
**ViewNameToGc** | Pointer to **NullableString** | The view name under which the snapshot of the migrated data was created. NOTE: This is populated only for data migration tasks. | [optional] 
**Warnings** | Pointer to [**[]ErrorProto**](ErrorProto.md) | Warnings if any. These warnings will be propogated to the UI by master. | [optional] 

## Methods

### NewSnapshotInfoProto

`func NewSnapshotInfoProto() *SnapshotInfoProto`

NewSnapshotInfoProto instantiates a new SnapshotInfoProto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotInfoProtoWithDefaults

`func NewSnapshotInfoProtoWithDefaults() *SnapshotInfoProto`

NewSnapshotInfoProtoWithDefaults instantiates a new SnapshotInfoProto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorRocksdbName

`func (o *SnapshotInfoProto) GetErrorRocksdbName() string`

GetErrorRocksdbName returns the ErrorRocksdbName field if non-nil, zero value otherwise.

### GetErrorRocksdbNameOk

`func (o *SnapshotInfoProto) GetErrorRocksdbNameOk() (*string, bool)`

GetErrorRocksdbNameOk returns a tuple with the ErrorRocksdbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorRocksdbName

`func (o *SnapshotInfoProto) SetErrorRocksdbName(v string)`

SetErrorRocksdbName sets ErrorRocksdbName field to given value.

### HasErrorRocksdbName

`func (o *SnapshotInfoProto) HasErrorRocksdbName() bool`

HasErrorRocksdbName returns a boolean if a field has been set.

### SetErrorRocksdbNameNil

`func (o *SnapshotInfoProto) SetErrorRocksdbNameNil(b bool)`

 SetErrorRocksdbNameNil sets the value for ErrorRocksdbName to be an explicit nil

### UnsetErrorRocksdbName
`func (o *SnapshotInfoProto) UnsetErrorRocksdbName()`

UnsetErrorRocksdbName ensures that no value is present for ErrorRocksdbName, not even an explicit nil
### GetFileWalkDone

`func (o *SnapshotInfoProto) GetFileWalkDone() bool`

GetFileWalkDone returns the FileWalkDone field if non-nil, zero value otherwise.

### GetFileWalkDoneOk

`func (o *SnapshotInfoProto) GetFileWalkDoneOk() (*bool, bool)`

GetFileWalkDoneOk returns a tuple with the FileWalkDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileWalkDone

`func (o *SnapshotInfoProto) SetFileWalkDone(v bool)`

SetFileWalkDone sets FileWalkDone field to given value.

### HasFileWalkDone

`func (o *SnapshotInfoProto) HasFileWalkDone() bool`

HasFileWalkDone returns a boolean if a field has been set.

### SetFileWalkDoneNil

`func (o *SnapshotInfoProto) SetFileWalkDoneNil(b bool)`

 SetFileWalkDoneNil sets the value for FileWalkDone to be an explicit nil

### UnsetFileWalkDone
`func (o *SnapshotInfoProto) UnsetFileWalkDone()`

UnsetFileWalkDone ensures that no value is present for FileWalkDone, not even an explicit nil
### GetNumAppInstances

`func (o *SnapshotInfoProto) GetNumAppInstances() int32`

GetNumAppInstances returns the NumAppInstances field if non-nil, zero value otherwise.

### GetNumAppInstancesOk

`func (o *SnapshotInfoProto) GetNumAppInstancesOk() (*int32, bool)`

GetNumAppInstancesOk returns a tuple with the NumAppInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAppInstances

`func (o *SnapshotInfoProto) SetNumAppInstances(v int32)`

SetNumAppInstances sets NumAppInstances field to given value.

### HasNumAppInstances

`func (o *SnapshotInfoProto) HasNumAppInstances() bool`

HasNumAppInstances returns a boolean if a field has been set.

### SetNumAppInstancesNil

`func (o *SnapshotInfoProto) SetNumAppInstancesNil(b bool)`

 SetNumAppInstancesNil sets the value for NumAppInstances to be an explicit nil

### UnsetNumAppInstances
`func (o *SnapshotInfoProto) UnsetNumAppInstances()`

UnsetNumAppInstances ensures that no value is present for NumAppInstances, not even an explicit nil
### GetNumAppObjects

`func (o *SnapshotInfoProto) GetNumAppObjects() int32`

GetNumAppObjects returns the NumAppObjects field if non-nil, zero value otherwise.

### GetNumAppObjectsOk

`func (o *SnapshotInfoProto) GetNumAppObjectsOk() (*int32, bool)`

GetNumAppObjectsOk returns a tuple with the NumAppObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumAppObjects

`func (o *SnapshotInfoProto) SetNumAppObjects(v int32)`

SetNumAppObjects sets NumAppObjects field to given value.

### HasNumAppObjects

`func (o *SnapshotInfoProto) HasNumAppObjects() bool`

HasNumAppObjects returns a boolean if a field has been set.

### SetNumAppObjectsNil

`func (o *SnapshotInfoProto) SetNumAppObjectsNil(b bool)`

 SetNumAppObjectsNil sets the value for NumAppObjects to be an explicit nil

### UnsetNumAppObjects
`func (o *SnapshotInfoProto) UnsetNumAppObjects()`

UnsetNumAppObjects ensures that no value is present for NumAppObjects, not even an explicit nil
### GetPostBackupScriptStatus

`func (o *SnapshotInfoProto) GetPostBackupScriptStatus() ScriptExecutionStatus`

GetPostBackupScriptStatus returns the PostBackupScriptStatus field if non-nil, zero value otherwise.

### GetPostBackupScriptStatusOk

`func (o *SnapshotInfoProto) GetPostBackupScriptStatusOk() (*ScriptExecutionStatus, bool)`

GetPostBackupScriptStatusOk returns a tuple with the PostBackupScriptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostBackupScriptStatus

`func (o *SnapshotInfoProto) SetPostBackupScriptStatus(v ScriptExecutionStatus)`

SetPostBackupScriptStatus sets PostBackupScriptStatus field to given value.

### HasPostBackupScriptStatus

`func (o *SnapshotInfoProto) HasPostBackupScriptStatus() bool`

HasPostBackupScriptStatus returns a boolean if a field has been set.

### GetPreBackupScriptStatus

`func (o *SnapshotInfoProto) GetPreBackupScriptStatus() ScriptExecutionStatus`

GetPreBackupScriptStatus returns the PreBackupScriptStatus field if non-nil, zero value otherwise.

### GetPreBackupScriptStatusOk

`func (o *SnapshotInfoProto) GetPreBackupScriptStatusOk() (*ScriptExecutionStatus, bool)`

GetPreBackupScriptStatusOk returns a tuple with the PreBackupScriptStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBackupScriptStatus

`func (o *SnapshotInfoProto) SetPreBackupScriptStatus(v ScriptExecutionStatus)`

SetPreBackupScriptStatus sets PreBackupScriptStatus field to given value.

### HasPreBackupScriptStatus

`func (o *SnapshotInfoProto) HasPreBackupScriptStatus() bool`

HasPreBackupScriptStatus returns a boolean if a field has been set.

### GetRelativeSnapshotDir

`func (o *SnapshotInfoProto) GetRelativeSnapshotDir() string`

GetRelativeSnapshotDir returns the RelativeSnapshotDir field if non-nil, zero value otherwise.

### GetRelativeSnapshotDirOk

`func (o *SnapshotInfoProto) GetRelativeSnapshotDirOk() (*string, bool)`

GetRelativeSnapshotDirOk returns a tuple with the RelativeSnapshotDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeSnapshotDir

`func (o *SnapshotInfoProto) SetRelativeSnapshotDir(v string)`

SetRelativeSnapshotDir sets RelativeSnapshotDir field to given value.

### HasRelativeSnapshotDir

`func (o *SnapshotInfoProto) HasRelativeSnapshotDir() bool`

HasRelativeSnapshotDir returns a boolean if a field has been set.

### SetRelativeSnapshotDirNil

`func (o *SnapshotInfoProto) SetRelativeSnapshotDirNil(b bool)`

 SetRelativeSnapshotDirNil sets the value for RelativeSnapshotDir to be an explicit nil

### UnsetRelativeSnapshotDir
`func (o *SnapshotInfoProto) UnsetRelativeSnapshotDir()`

UnsetRelativeSnapshotDir ensures that no value is present for RelativeSnapshotDir, not even an explicit nil
### GetRootPath

`func (o *SnapshotInfoProto) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *SnapshotInfoProto) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *SnapshotInfoProto) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *SnapshotInfoProto) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### SetRootPathNil

`func (o *SnapshotInfoProto) SetRootPathNil(b bool)`

 SetRootPathNil sets the value for RootPath to be an explicit nil

### UnsetRootPath
`func (o *SnapshotInfoProto) UnsetRootPath()`

UnsetRootPath ensures that no value is present for RootPath, not even an explicit nil
### GetScribeTableColumn

`func (o *SnapshotInfoProto) GetScribeTableColumn() string`

GetScribeTableColumn returns the ScribeTableColumn field if non-nil, zero value otherwise.

### GetScribeTableColumnOk

`func (o *SnapshotInfoProto) GetScribeTableColumnOk() (*string, bool)`

GetScribeTableColumnOk returns a tuple with the ScribeTableColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScribeTableColumn

`func (o *SnapshotInfoProto) SetScribeTableColumn(v string)`

SetScribeTableColumn sets ScribeTableColumn field to given value.

### HasScribeTableColumn

`func (o *SnapshotInfoProto) HasScribeTableColumn() bool`

HasScribeTableColumn returns a boolean if a field has been set.

### SetScribeTableColumnNil

`func (o *SnapshotInfoProto) SetScribeTableColumnNil(b bool)`

 SetScribeTableColumnNil sets the value for ScribeTableColumn to be an explicit nil

### UnsetScribeTableColumn
`func (o *SnapshotInfoProto) UnsetScribeTableColumn()`

UnsetScribeTableColumn ensures that no value is present for ScribeTableColumn, not even an explicit nil
### GetScribeTableRow

`func (o *SnapshotInfoProto) GetScribeTableRow() string`

GetScribeTableRow returns the ScribeTableRow field if non-nil, zero value otherwise.

### GetScribeTableRowOk

`func (o *SnapshotInfoProto) GetScribeTableRowOk() (*string, bool)`

GetScribeTableRowOk returns a tuple with the ScribeTableRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScribeTableRow

`func (o *SnapshotInfoProto) SetScribeTableRow(v string)`

SetScribeTableRow sets ScribeTableRow field to given value.

### HasScribeTableRow

`func (o *SnapshotInfoProto) HasScribeTableRow() bool`

HasScribeTableRow returns a boolean if a field has been set.

### SetScribeTableRowNil

`func (o *SnapshotInfoProto) SetScribeTableRowNil(b bool)`

 SetScribeTableRowNil sets the value for ScribeTableRow to be an explicit nil

### UnsetScribeTableRow
`func (o *SnapshotInfoProto) UnsetScribeTableRow()`

UnsetScribeTableRow ensures that no value is present for ScribeTableRow, not even an explicit nil
### GetSlaveTaskStartTimeUsecs

`func (o *SnapshotInfoProto) GetSlaveTaskStartTimeUsecs() int64`

GetSlaveTaskStartTimeUsecs returns the SlaveTaskStartTimeUsecs field if non-nil, zero value otherwise.

### GetSlaveTaskStartTimeUsecsOk

`func (o *SnapshotInfoProto) GetSlaveTaskStartTimeUsecsOk() (*int64, bool)`

GetSlaveTaskStartTimeUsecsOk returns a tuple with the SlaveTaskStartTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveTaskStartTimeUsecs

`func (o *SnapshotInfoProto) SetSlaveTaskStartTimeUsecs(v int64)`

SetSlaveTaskStartTimeUsecs sets SlaveTaskStartTimeUsecs field to given value.

### HasSlaveTaskStartTimeUsecs

`func (o *SnapshotInfoProto) HasSlaveTaskStartTimeUsecs() bool`

HasSlaveTaskStartTimeUsecs returns a boolean if a field has been set.

### SetSlaveTaskStartTimeUsecsNil

`func (o *SnapshotInfoProto) SetSlaveTaskStartTimeUsecsNil(b bool)`

 SetSlaveTaskStartTimeUsecsNil sets the value for SlaveTaskStartTimeUsecs to be an explicit nil

### UnsetSlaveTaskStartTimeUsecs
`func (o *SnapshotInfoProto) UnsetSlaveTaskStartTimeUsecs()`

UnsetSlaveTaskStartTimeUsecs ensures that no value is present for SlaveTaskStartTimeUsecs, not even an explicit nil
### GetSnapshotType

`func (o *SnapshotInfoProto) GetSnapshotType() ObjectSnapshotType`

GetSnapshotType returns the SnapshotType field if non-nil, zero value otherwise.

### GetSnapshotTypeOk

`func (o *SnapshotInfoProto) GetSnapshotTypeOk() (*ObjectSnapshotType, bool)`

GetSnapshotTypeOk returns a tuple with the SnapshotType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotType

`func (o *SnapshotInfoProto) SetSnapshotType(v ObjectSnapshotType)`

SetSnapshotType sets SnapshotType field to given value.

### HasSnapshotType

`func (o *SnapshotInfoProto) HasSnapshotType() bool`

HasSnapshotType returns a boolean if a field has been set.

### GetSourceSnapshotCreateTimeUsecs

`func (o *SnapshotInfoProto) GetSourceSnapshotCreateTimeUsecs() int64`

GetSourceSnapshotCreateTimeUsecs returns the SourceSnapshotCreateTimeUsecs field if non-nil, zero value otherwise.

### GetSourceSnapshotCreateTimeUsecsOk

`func (o *SnapshotInfoProto) GetSourceSnapshotCreateTimeUsecsOk() (*int64, bool)`

GetSourceSnapshotCreateTimeUsecsOk returns a tuple with the SourceSnapshotCreateTimeUsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotCreateTimeUsecs

`func (o *SnapshotInfoProto) SetSourceSnapshotCreateTimeUsecs(v int64)`

SetSourceSnapshotCreateTimeUsecs sets SourceSnapshotCreateTimeUsecs field to given value.

### HasSourceSnapshotCreateTimeUsecs

`func (o *SnapshotInfoProto) HasSourceSnapshotCreateTimeUsecs() bool`

HasSourceSnapshotCreateTimeUsecs returns a boolean if a field has been set.

### SetSourceSnapshotCreateTimeUsecsNil

`func (o *SnapshotInfoProto) SetSourceSnapshotCreateTimeUsecsNil(b bool)`

 SetSourceSnapshotCreateTimeUsecsNil sets the value for SourceSnapshotCreateTimeUsecs to be an explicit nil

### UnsetSourceSnapshotCreateTimeUsecs
`func (o *SnapshotInfoProto) UnsetSourceSnapshotCreateTimeUsecs()`

UnsetSourceSnapshotCreateTimeUsecs ensures that no value is present for SourceSnapshotCreateTimeUsecs, not even an explicit nil
### GetSourceSnapshotName

`func (o *SnapshotInfoProto) GetSourceSnapshotName() string`

GetSourceSnapshotName returns the SourceSnapshotName field if non-nil, zero value otherwise.

### GetSourceSnapshotNameOk

`func (o *SnapshotInfoProto) GetSourceSnapshotNameOk() (*string, bool)`

GetSourceSnapshotNameOk returns a tuple with the SourceSnapshotName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotName

`func (o *SnapshotInfoProto) SetSourceSnapshotName(v string)`

SetSourceSnapshotName sets SourceSnapshotName field to given value.

### HasSourceSnapshotName

`func (o *SnapshotInfoProto) HasSourceSnapshotName() bool`

HasSourceSnapshotName returns a boolean if a field has been set.

### SetSourceSnapshotNameNil

`func (o *SnapshotInfoProto) SetSourceSnapshotNameNil(b bool)`

 SetSourceSnapshotNameNil sets the value for SourceSnapshotName to be an explicit nil

### UnsetSourceSnapshotName
`func (o *SnapshotInfoProto) UnsetSourceSnapshotName()`

UnsetSourceSnapshotName ensures that no value is present for SourceSnapshotName, not even an explicit nil
### GetStorageSnapshotProvider

`func (o *SnapshotInfoProto) GetStorageSnapshotProvider() StorageSnapshotProviderParams`

GetStorageSnapshotProvider returns the StorageSnapshotProvider field if non-nil, zero value otherwise.

### GetStorageSnapshotProviderOk

`func (o *SnapshotInfoProto) GetStorageSnapshotProviderOk() (*StorageSnapshotProviderParams, bool)`

GetStorageSnapshotProviderOk returns a tuple with the StorageSnapshotProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSnapshotProvider

`func (o *SnapshotInfoProto) SetStorageSnapshotProvider(v StorageSnapshotProviderParams)`

SetStorageSnapshotProvider sets StorageSnapshotProvider field to given value.

### HasStorageSnapshotProvider

`func (o *SnapshotInfoProto) HasStorageSnapshotProvider() bool`

HasStorageSnapshotProvider returns a boolean if a field has been set.

### GetTargetType

`func (o *SnapshotInfoProto) GetTargetType() int32`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *SnapshotInfoProto) GetTargetTypeOk() (*int32, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *SnapshotInfoProto) SetTargetType(v int32)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *SnapshotInfoProto) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *SnapshotInfoProto) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *SnapshotInfoProto) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTotalBytesReadFromSource

`func (o *SnapshotInfoProto) GetTotalBytesReadFromSource() int64`

GetTotalBytesReadFromSource returns the TotalBytesReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesReadFromSourceOk

`func (o *SnapshotInfoProto) GetTotalBytesReadFromSourceOk() (*int64, bool)`

GetTotalBytesReadFromSourceOk returns a tuple with the TotalBytesReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesReadFromSource

`func (o *SnapshotInfoProto) SetTotalBytesReadFromSource(v int64)`

SetTotalBytesReadFromSource sets TotalBytesReadFromSource field to given value.

### HasTotalBytesReadFromSource

`func (o *SnapshotInfoProto) HasTotalBytesReadFromSource() bool`

HasTotalBytesReadFromSource returns a boolean if a field has been set.

### SetTotalBytesReadFromSourceNil

`func (o *SnapshotInfoProto) SetTotalBytesReadFromSourceNil(b bool)`

 SetTotalBytesReadFromSourceNil sets the value for TotalBytesReadFromSource to be an explicit nil

### UnsetTotalBytesReadFromSource
`func (o *SnapshotInfoProto) UnsetTotalBytesReadFromSource()`

UnsetTotalBytesReadFromSource ensures that no value is present for TotalBytesReadFromSource, not even an explicit nil
### GetTotalBytesToReadFromSource

`func (o *SnapshotInfoProto) GetTotalBytesToReadFromSource() int64`

GetTotalBytesToReadFromSource returns the TotalBytesToReadFromSource field if non-nil, zero value otherwise.

### GetTotalBytesToReadFromSourceOk

`func (o *SnapshotInfoProto) GetTotalBytesToReadFromSourceOk() (*int64, bool)`

GetTotalBytesToReadFromSourceOk returns a tuple with the TotalBytesToReadFromSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesToReadFromSource

`func (o *SnapshotInfoProto) SetTotalBytesToReadFromSource(v int64)`

SetTotalBytesToReadFromSource sets TotalBytesToReadFromSource field to given value.

### HasTotalBytesToReadFromSource

`func (o *SnapshotInfoProto) HasTotalBytesToReadFromSource() bool`

HasTotalBytesToReadFromSource returns a boolean if a field has been set.

### SetTotalBytesToReadFromSourceNil

`func (o *SnapshotInfoProto) SetTotalBytesToReadFromSourceNil(b bool)`

 SetTotalBytesToReadFromSourceNil sets the value for TotalBytesToReadFromSource to be an explicit nil

### UnsetTotalBytesToReadFromSource
`func (o *SnapshotInfoProto) UnsetTotalBytesToReadFromSource()`

UnsetTotalBytesToReadFromSource ensures that no value is present for TotalBytesToReadFromSource, not even an explicit nil
### GetTotalChangedEntityCount

`func (o *SnapshotInfoProto) GetTotalChangedEntityCount() int64`

GetTotalChangedEntityCount returns the TotalChangedEntityCount field if non-nil, zero value otherwise.

### GetTotalChangedEntityCountOk

`func (o *SnapshotInfoProto) GetTotalChangedEntityCountOk() (*int64, bool)`

GetTotalChangedEntityCountOk returns a tuple with the TotalChangedEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalChangedEntityCount

`func (o *SnapshotInfoProto) SetTotalChangedEntityCount(v int64)`

SetTotalChangedEntityCount sets TotalChangedEntityCount field to given value.

### HasTotalChangedEntityCount

`func (o *SnapshotInfoProto) HasTotalChangedEntityCount() bool`

HasTotalChangedEntityCount returns a boolean if a field has been set.

### SetTotalChangedEntityCountNil

`func (o *SnapshotInfoProto) SetTotalChangedEntityCountNil(b bool)`

 SetTotalChangedEntityCountNil sets the value for TotalChangedEntityCount to be an explicit nil

### UnsetTotalChangedEntityCount
`func (o *SnapshotInfoProto) UnsetTotalChangedEntityCount()`

UnsetTotalChangedEntityCount ensures that no value is present for TotalChangedEntityCount, not even an explicit nil
### GetTotalEntityCount

`func (o *SnapshotInfoProto) GetTotalEntityCount() int64`

GetTotalEntityCount returns the TotalEntityCount field if non-nil, zero value otherwise.

### GetTotalEntityCountOk

`func (o *SnapshotInfoProto) GetTotalEntityCountOk() (*int64, bool)`

GetTotalEntityCountOk returns a tuple with the TotalEntityCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEntityCount

`func (o *SnapshotInfoProto) SetTotalEntityCount(v int64)`

SetTotalEntityCount sets TotalEntityCount field to given value.

### HasTotalEntityCount

`func (o *SnapshotInfoProto) HasTotalEntityCount() bool`

HasTotalEntityCount returns a boolean if a field has been set.

### SetTotalEntityCountNil

`func (o *SnapshotInfoProto) SetTotalEntityCountNil(b bool)`

 SetTotalEntityCountNil sets the value for TotalEntityCount to be an explicit nil

### UnsetTotalEntityCount
`func (o *SnapshotInfoProto) UnsetTotalEntityCount()`

UnsetTotalEntityCount ensures that no value is present for TotalEntityCount, not even an explicit nil
### GetTotalLogicalBackupSizeBytes

`func (o *SnapshotInfoProto) GetTotalLogicalBackupSizeBytes() int64`

GetTotalLogicalBackupSizeBytes returns the TotalLogicalBackupSizeBytes field if non-nil, zero value otherwise.

### GetTotalLogicalBackupSizeBytesOk

`func (o *SnapshotInfoProto) GetTotalLogicalBackupSizeBytesOk() (*int64, bool)`

GetTotalLogicalBackupSizeBytesOk returns a tuple with the TotalLogicalBackupSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLogicalBackupSizeBytes

`func (o *SnapshotInfoProto) SetTotalLogicalBackupSizeBytes(v int64)`

SetTotalLogicalBackupSizeBytes sets TotalLogicalBackupSizeBytes field to given value.

### HasTotalLogicalBackupSizeBytes

`func (o *SnapshotInfoProto) HasTotalLogicalBackupSizeBytes() bool`

HasTotalLogicalBackupSizeBytes returns a boolean if a field has been set.

### SetTotalLogicalBackupSizeBytesNil

`func (o *SnapshotInfoProto) SetTotalLogicalBackupSizeBytesNil(b bool)`

 SetTotalLogicalBackupSizeBytesNil sets the value for TotalLogicalBackupSizeBytes to be an explicit nil

### UnsetTotalLogicalBackupSizeBytes
`func (o *SnapshotInfoProto) UnsetTotalLogicalBackupSizeBytes()`

UnsetTotalLogicalBackupSizeBytes ensures that no value is present for TotalLogicalBackupSizeBytes, not even an explicit nil
### GetTotalPrimaryPhysicalSizeBytes

`func (o *SnapshotInfoProto) GetTotalPrimaryPhysicalSizeBytes() int64`

GetTotalPrimaryPhysicalSizeBytes returns the TotalPrimaryPhysicalSizeBytes field if non-nil, zero value otherwise.

### GetTotalPrimaryPhysicalSizeBytesOk

`func (o *SnapshotInfoProto) GetTotalPrimaryPhysicalSizeBytesOk() (*int64, bool)`

GetTotalPrimaryPhysicalSizeBytesOk returns a tuple with the TotalPrimaryPhysicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrimaryPhysicalSizeBytes

`func (o *SnapshotInfoProto) SetTotalPrimaryPhysicalSizeBytes(v int64)`

SetTotalPrimaryPhysicalSizeBytes sets TotalPrimaryPhysicalSizeBytes field to given value.

### HasTotalPrimaryPhysicalSizeBytes

`func (o *SnapshotInfoProto) HasTotalPrimaryPhysicalSizeBytes() bool`

HasTotalPrimaryPhysicalSizeBytes returns a boolean if a field has been set.

### SetTotalPrimaryPhysicalSizeBytesNil

`func (o *SnapshotInfoProto) SetTotalPrimaryPhysicalSizeBytesNil(b bool)`

 SetTotalPrimaryPhysicalSizeBytesNil sets the value for TotalPrimaryPhysicalSizeBytes to be an explicit nil

### UnsetTotalPrimaryPhysicalSizeBytes
`func (o *SnapshotInfoProto) UnsetTotalPrimaryPhysicalSizeBytes()`

UnsetTotalPrimaryPhysicalSizeBytes ensures that no value is present for TotalPrimaryPhysicalSizeBytes, not even an explicit nil
### GetType

`func (o *SnapshotInfoProto) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotInfoProto) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotInfoProto) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SnapshotInfoProto) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *SnapshotInfoProto) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *SnapshotInfoProto) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetViewCaseInsensitivityAltered

`func (o *SnapshotInfoProto) GetViewCaseInsensitivityAltered() bool`

GetViewCaseInsensitivityAltered returns the ViewCaseInsensitivityAltered field if non-nil, zero value otherwise.

### GetViewCaseInsensitivityAlteredOk

`func (o *SnapshotInfoProto) GetViewCaseInsensitivityAlteredOk() (*bool, bool)`

GetViewCaseInsensitivityAlteredOk returns a tuple with the ViewCaseInsensitivityAltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCaseInsensitivityAltered

`func (o *SnapshotInfoProto) SetViewCaseInsensitivityAltered(v bool)`

SetViewCaseInsensitivityAltered sets ViewCaseInsensitivityAltered field to given value.

### HasViewCaseInsensitivityAltered

`func (o *SnapshotInfoProto) HasViewCaseInsensitivityAltered() bool`

HasViewCaseInsensitivityAltered returns a boolean if a field has been set.

### SetViewCaseInsensitivityAlteredNil

`func (o *SnapshotInfoProto) SetViewCaseInsensitivityAlteredNil(b bool)`

 SetViewCaseInsensitivityAlteredNil sets the value for ViewCaseInsensitivityAltered to be an explicit nil

### UnsetViewCaseInsensitivityAltered
`func (o *SnapshotInfoProto) UnsetViewCaseInsensitivityAltered()`

UnsetViewCaseInsensitivityAltered ensures that no value is present for ViewCaseInsensitivityAltered, not even an explicit nil
### GetViewName

`func (o *SnapshotInfoProto) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SnapshotInfoProto) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SnapshotInfoProto) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SnapshotInfoProto) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### SetViewNameNil

`func (o *SnapshotInfoProto) SetViewNameNil(b bool)`

 SetViewNameNil sets the value for ViewName to be an explicit nil

### UnsetViewName
`func (o *SnapshotInfoProto) UnsetViewName()`

UnsetViewName ensures that no value is present for ViewName, not even an explicit nil
### GetViewNameToGc

`func (o *SnapshotInfoProto) GetViewNameToGc() string`

GetViewNameToGc returns the ViewNameToGc field if non-nil, zero value otherwise.

### GetViewNameToGcOk

`func (o *SnapshotInfoProto) GetViewNameToGcOk() (*string, bool)`

GetViewNameToGcOk returns a tuple with the ViewNameToGc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewNameToGc

`func (o *SnapshotInfoProto) SetViewNameToGc(v string)`

SetViewNameToGc sets ViewNameToGc field to given value.

### HasViewNameToGc

`func (o *SnapshotInfoProto) HasViewNameToGc() bool`

HasViewNameToGc returns a boolean if a field has been set.

### SetViewNameToGcNil

`func (o *SnapshotInfoProto) SetViewNameToGcNil(b bool)`

 SetViewNameToGcNil sets the value for ViewNameToGc to be an explicit nil

### UnsetViewNameToGc
`func (o *SnapshotInfoProto) UnsetViewNameToGc()`

UnsetViewNameToGc ensures that no value is present for ViewNameToGc, not even an explicit nil
### GetWarnings

`func (o *SnapshotInfoProto) GetWarnings() []ErrorProto`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SnapshotInfoProto) GetWarningsOk() (*[]ErrorProto, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SnapshotInfoProto) SetWarnings(v []ErrorProto)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *SnapshotInfoProto) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### SetWarningsNil

`func (o *SnapshotInfoProto) SetWarningsNil(b bool)`

 SetWarningsNil sets the value for Warnings to be an explicit nil

### UnsetWarnings
`func (o *SnapshotInfoProto) UnsetWarnings()`

UnsetWarnings ensures that no value is present for Warnings, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



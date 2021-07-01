# WindowsHostSnapshotParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyOnlyBackup** | Pointer to **NullableBool** | Specifies whether to backup regardless of the state of each file&#39;s backup history. Backup history will not be updated. Refer Microsoft documentation on VSS_BT_COPY. | [optional] 
**DisableMetadata** | Pointer to **NullableBool** | Specifies whether to disable fetching and storing of some metadata on Cohesity Cluster to save storage space. Otherwise, there will be some metadata fetched and stored on Cohesity Cluster. | [optional] 
**DisableNotification** | Pointer to **NullableBool** | Specifies whether to disable some notification steps when taking snapshots. | [optional] 
**ExcludedVssWriters** | Pointer to **[]string** | Specifies a list of Windows VSS writers that are excluded from backups. For example, \&quot;ASR Writer\&quot;, \&quot;System Writer\&quot;. Refer Microsoft documentaion for a complete list. | [optional] 

## Methods

### NewWindowsHostSnapshotParameters

`func NewWindowsHostSnapshotParameters() *WindowsHostSnapshotParameters`

NewWindowsHostSnapshotParameters instantiates a new WindowsHostSnapshotParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWindowsHostSnapshotParametersWithDefaults

`func NewWindowsHostSnapshotParametersWithDefaults() *WindowsHostSnapshotParameters`

NewWindowsHostSnapshotParametersWithDefaults instantiates a new WindowsHostSnapshotParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyOnlyBackup

`func (o *WindowsHostSnapshotParameters) GetCopyOnlyBackup() bool`

GetCopyOnlyBackup returns the CopyOnlyBackup field if non-nil, zero value otherwise.

### GetCopyOnlyBackupOk

`func (o *WindowsHostSnapshotParameters) GetCopyOnlyBackupOk() (*bool, bool)`

GetCopyOnlyBackupOk returns a tuple with the CopyOnlyBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyOnlyBackup

`func (o *WindowsHostSnapshotParameters) SetCopyOnlyBackup(v bool)`

SetCopyOnlyBackup sets CopyOnlyBackup field to given value.

### HasCopyOnlyBackup

`func (o *WindowsHostSnapshotParameters) HasCopyOnlyBackup() bool`

HasCopyOnlyBackup returns a boolean if a field has been set.

### SetCopyOnlyBackupNil

`func (o *WindowsHostSnapshotParameters) SetCopyOnlyBackupNil(b bool)`

 SetCopyOnlyBackupNil sets the value for CopyOnlyBackup to be an explicit nil

### UnsetCopyOnlyBackup
`func (o *WindowsHostSnapshotParameters) UnsetCopyOnlyBackup()`

UnsetCopyOnlyBackup ensures that no value is present for CopyOnlyBackup, not even an explicit nil
### GetDisableMetadata

`func (o *WindowsHostSnapshotParameters) GetDisableMetadata() bool`

GetDisableMetadata returns the DisableMetadata field if non-nil, zero value otherwise.

### GetDisableMetadataOk

`func (o *WindowsHostSnapshotParameters) GetDisableMetadataOk() (*bool, bool)`

GetDisableMetadataOk returns a tuple with the DisableMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableMetadata

`func (o *WindowsHostSnapshotParameters) SetDisableMetadata(v bool)`

SetDisableMetadata sets DisableMetadata field to given value.

### HasDisableMetadata

`func (o *WindowsHostSnapshotParameters) HasDisableMetadata() bool`

HasDisableMetadata returns a boolean if a field has been set.

### SetDisableMetadataNil

`func (o *WindowsHostSnapshotParameters) SetDisableMetadataNil(b bool)`

 SetDisableMetadataNil sets the value for DisableMetadata to be an explicit nil

### UnsetDisableMetadata
`func (o *WindowsHostSnapshotParameters) UnsetDisableMetadata()`

UnsetDisableMetadata ensures that no value is present for DisableMetadata, not even an explicit nil
### GetDisableNotification

`func (o *WindowsHostSnapshotParameters) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *WindowsHostSnapshotParameters) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *WindowsHostSnapshotParameters) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *WindowsHostSnapshotParameters) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### SetDisableNotificationNil

`func (o *WindowsHostSnapshotParameters) SetDisableNotificationNil(b bool)`

 SetDisableNotificationNil sets the value for DisableNotification to be an explicit nil

### UnsetDisableNotification
`func (o *WindowsHostSnapshotParameters) UnsetDisableNotification()`

UnsetDisableNotification ensures that no value is present for DisableNotification, not even an explicit nil
### GetExcludedVssWriters

`func (o *WindowsHostSnapshotParameters) GetExcludedVssWriters() []string`

GetExcludedVssWriters returns the ExcludedVssWriters field if non-nil, zero value otherwise.

### GetExcludedVssWritersOk

`func (o *WindowsHostSnapshotParameters) GetExcludedVssWritersOk() (*[]string, bool)`

GetExcludedVssWritersOk returns a tuple with the ExcludedVssWriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedVssWriters

`func (o *WindowsHostSnapshotParameters) SetExcludedVssWriters(v []string)`

SetExcludedVssWriters sets ExcludedVssWriters field to given value.

### HasExcludedVssWriters

`func (o *WindowsHostSnapshotParameters) HasExcludedVssWriters() bool`

HasExcludedVssWriters returns a boolean if a field has been set.

### SetExcludedVssWritersNil

`func (o *WindowsHostSnapshotParameters) SetExcludedVssWritersNil(b bool)`

 SetExcludedVssWritersNil sets the value for ExcludedVssWriters to be an explicit nil

### UnsetExcludedVssWriters
`func (o *WindowsHostSnapshotParameters) UnsetExcludedVssWriters()`

UnsetExcludedVssWriters ensures that no value is present for ExcludedVssWriters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



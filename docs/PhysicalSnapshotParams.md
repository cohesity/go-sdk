# PhysicalSnapshotParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchSnapshotMetadataDisabled** | Pointer to **NullableBool** | Whether fetching and storing of snapshot metadata was disabled. | [optional] 
**NotifyBackupCompleteDisabled** | Pointer to **NullableBool** | Whether notify backup complete step was disabled. | [optional] 
**VssCopyOnlyBackup** | Pointer to **NullableBool** | If copy_only_backup option is requrested at the time of the snapshot. | [optional] 
**VssExcludedWriters** | Pointer to **[]string** | List of VSS writers that were excluded. | [optional] 

## Methods

### NewPhysicalSnapshotParams

`func NewPhysicalSnapshotParams() *PhysicalSnapshotParams`

NewPhysicalSnapshotParams instantiates a new PhysicalSnapshotParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalSnapshotParamsWithDefaults

`func NewPhysicalSnapshotParamsWithDefaults() *PhysicalSnapshotParams`

NewPhysicalSnapshotParamsWithDefaults instantiates a new PhysicalSnapshotParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchSnapshotMetadataDisabled

`func (o *PhysicalSnapshotParams) GetFetchSnapshotMetadataDisabled() bool`

GetFetchSnapshotMetadataDisabled returns the FetchSnapshotMetadataDisabled field if non-nil, zero value otherwise.

### GetFetchSnapshotMetadataDisabledOk

`func (o *PhysicalSnapshotParams) GetFetchSnapshotMetadataDisabledOk() (*bool, bool)`

GetFetchSnapshotMetadataDisabledOk returns a tuple with the FetchSnapshotMetadataDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchSnapshotMetadataDisabled

`func (o *PhysicalSnapshotParams) SetFetchSnapshotMetadataDisabled(v bool)`

SetFetchSnapshotMetadataDisabled sets FetchSnapshotMetadataDisabled field to given value.

### HasFetchSnapshotMetadataDisabled

`func (o *PhysicalSnapshotParams) HasFetchSnapshotMetadataDisabled() bool`

HasFetchSnapshotMetadataDisabled returns a boolean if a field has been set.

### SetFetchSnapshotMetadataDisabledNil

`func (o *PhysicalSnapshotParams) SetFetchSnapshotMetadataDisabledNil(b bool)`

 SetFetchSnapshotMetadataDisabledNil sets the value for FetchSnapshotMetadataDisabled to be an explicit nil

### UnsetFetchSnapshotMetadataDisabled
`func (o *PhysicalSnapshotParams) UnsetFetchSnapshotMetadataDisabled()`

UnsetFetchSnapshotMetadataDisabled ensures that no value is present for FetchSnapshotMetadataDisabled, not even an explicit nil
### GetNotifyBackupCompleteDisabled

`func (o *PhysicalSnapshotParams) GetNotifyBackupCompleteDisabled() bool`

GetNotifyBackupCompleteDisabled returns the NotifyBackupCompleteDisabled field if non-nil, zero value otherwise.

### GetNotifyBackupCompleteDisabledOk

`func (o *PhysicalSnapshotParams) GetNotifyBackupCompleteDisabledOk() (*bool, bool)`

GetNotifyBackupCompleteDisabledOk returns a tuple with the NotifyBackupCompleteDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyBackupCompleteDisabled

`func (o *PhysicalSnapshotParams) SetNotifyBackupCompleteDisabled(v bool)`

SetNotifyBackupCompleteDisabled sets NotifyBackupCompleteDisabled field to given value.

### HasNotifyBackupCompleteDisabled

`func (o *PhysicalSnapshotParams) HasNotifyBackupCompleteDisabled() bool`

HasNotifyBackupCompleteDisabled returns a boolean if a field has been set.

### SetNotifyBackupCompleteDisabledNil

`func (o *PhysicalSnapshotParams) SetNotifyBackupCompleteDisabledNil(b bool)`

 SetNotifyBackupCompleteDisabledNil sets the value for NotifyBackupCompleteDisabled to be an explicit nil

### UnsetNotifyBackupCompleteDisabled
`func (o *PhysicalSnapshotParams) UnsetNotifyBackupCompleteDisabled()`

UnsetNotifyBackupCompleteDisabled ensures that no value is present for NotifyBackupCompleteDisabled, not even an explicit nil
### GetVssCopyOnlyBackup

`func (o *PhysicalSnapshotParams) GetVssCopyOnlyBackup() bool`

GetVssCopyOnlyBackup returns the VssCopyOnlyBackup field if non-nil, zero value otherwise.

### GetVssCopyOnlyBackupOk

`func (o *PhysicalSnapshotParams) GetVssCopyOnlyBackupOk() (*bool, bool)`

GetVssCopyOnlyBackupOk returns a tuple with the VssCopyOnlyBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVssCopyOnlyBackup

`func (o *PhysicalSnapshotParams) SetVssCopyOnlyBackup(v bool)`

SetVssCopyOnlyBackup sets VssCopyOnlyBackup field to given value.

### HasVssCopyOnlyBackup

`func (o *PhysicalSnapshotParams) HasVssCopyOnlyBackup() bool`

HasVssCopyOnlyBackup returns a boolean if a field has been set.

### SetVssCopyOnlyBackupNil

`func (o *PhysicalSnapshotParams) SetVssCopyOnlyBackupNil(b bool)`

 SetVssCopyOnlyBackupNil sets the value for VssCopyOnlyBackup to be an explicit nil

### UnsetVssCopyOnlyBackup
`func (o *PhysicalSnapshotParams) UnsetVssCopyOnlyBackup()`

UnsetVssCopyOnlyBackup ensures that no value is present for VssCopyOnlyBackup, not even an explicit nil
### GetVssExcludedWriters

`func (o *PhysicalSnapshotParams) GetVssExcludedWriters() []string`

GetVssExcludedWriters returns the VssExcludedWriters field if non-nil, zero value otherwise.

### GetVssExcludedWritersOk

`func (o *PhysicalSnapshotParams) GetVssExcludedWritersOk() (*[]string, bool)`

GetVssExcludedWritersOk returns a tuple with the VssExcludedWriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVssExcludedWriters

`func (o *PhysicalSnapshotParams) SetVssExcludedWriters(v []string)`

SetVssExcludedWriters sets VssExcludedWriters field to given value.

### HasVssExcludedWriters

`func (o *PhysicalSnapshotParams) HasVssExcludedWriters() bool`

HasVssExcludedWriters returns a boolean if a field has been set.

### SetVssExcludedWritersNil

`func (o *PhysicalSnapshotParams) SetVssExcludedWritersNil(b bool)`

 SetVssExcludedWritersNil sets the value for VssExcludedWriters to be an explicit nil

### UnsetVssExcludedWriters
`func (o *PhysicalSnapshotParams) UnsetVssExcludedWriters()`

UnsetVssExcludedWriters ensures that no value is present for VssExcludedWriters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



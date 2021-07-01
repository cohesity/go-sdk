# FlashBladeFileSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupEnabled** | Pointer to **NullableBool** | Specifies whether the .snapshot directory exists on the file system. Backup is enabled only if the directory exists. | [optional] 
**CreatedTimeMsecs** | Pointer to **NullableInt64** | Specifies the time when the filesystem was created in Unix epoch time in milliseconds. | [optional] 
**LogicalCapacityBytes** | Pointer to **NullableInt64** | Specifies the total capacity in bytes of the file system. | [optional] 
**LogicalUsedBytes** | Pointer to **NullableInt64** | Specifies the size of logical data currently represented on the file system in bytes. | [optional] 
**NfsInfo** | Pointer to [**FlashBladeNfsInfo**](FlashBladeNfsInfo.md) |  | [optional] 
**PhysicalUsedBytes** | Pointer to **NullableInt64** | Specifies the size of physical data currently consumed by the file system. This includes the space used for the snapshots. | [optional] 
**Protocols** | Pointer to **[]string** | List of Protocols.  Specifies the list of protocols enabled on the file system. &#39;kNfs&#39; indicates NFS exports are supported on Pure FlashBlade File System. &#39;kCifs2&#39; indicates CIFS/SMB Shares are supported on Pure FlashBlade File System. &#39;kHttp&#39; indicates object protocol over HTTP and HTTPS are supported. | [optional] 
**UniqueUsedBytes** | Pointer to **NullableInt64** | Specifies the size of physical data consumed by the file system itself not including the size of the snapshots. | [optional] 

## Methods

### NewFlashBladeFileSystem

`func NewFlashBladeFileSystem() *FlashBladeFileSystem`

NewFlashBladeFileSystem instantiates a new FlashBladeFileSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlashBladeFileSystemWithDefaults

`func NewFlashBladeFileSystemWithDefaults() *FlashBladeFileSystem`

NewFlashBladeFileSystemWithDefaults instantiates a new FlashBladeFileSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupEnabled

`func (o *FlashBladeFileSystem) GetBackupEnabled() bool`

GetBackupEnabled returns the BackupEnabled field if non-nil, zero value otherwise.

### GetBackupEnabledOk

`func (o *FlashBladeFileSystem) GetBackupEnabledOk() (*bool, bool)`

GetBackupEnabledOk returns a tuple with the BackupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupEnabled

`func (o *FlashBladeFileSystem) SetBackupEnabled(v bool)`

SetBackupEnabled sets BackupEnabled field to given value.

### HasBackupEnabled

`func (o *FlashBladeFileSystem) HasBackupEnabled() bool`

HasBackupEnabled returns a boolean if a field has been set.

### SetBackupEnabledNil

`func (o *FlashBladeFileSystem) SetBackupEnabledNil(b bool)`

 SetBackupEnabledNil sets the value for BackupEnabled to be an explicit nil

### UnsetBackupEnabled
`func (o *FlashBladeFileSystem) UnsetBackupEnabled()`

UnsetBackupEnabled ensures that no value is present for BackupEnabled, not even an explicit nil
### GetCreatedTimeMsecs

`func (o *FlashBladeFileSystem) GetCreatedTimeMsecs() int64`

GetCreatedTimeMsecs returns the CreatedTimeMsecs field if non-nil, zero value otherwise.

### GetCreatedTimeMsecsOk

`func (o *FlashBladeFileSystem) GetCreatedTimeMsecsOk() (*int64, bool)`

GetCreatedTimeMsecsOk returns a tuple with the CreatedTimeMsecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimeMsecs

`func (o *FlashBladeFileSystem) SetCreatedTimeMsecs(v int64)`

SetCreatedTimeMsecs sets CreatedTimeMsecs field to given value.

### HasCreatedTimeMsecs

`func (o *FlashBladeFileSystem) HasCreatedTimeMsecs() bool`

HasCreatedTimeMsecs returns a boolean if a field has been set.

### SetCreatedTimeMsecsNil

`func (o *FlashBladeFileSystem) SetCreatedTimeMsecsNil(b bool)`

 SetCreatedTimeMsecsNil sets the value for CreatedTimeMsecs to be an explicit nil

### UnsetCreatedTimeMsecs
`func (o *FlashBladeFileSystem) UnsetCreatedTimeMsecs()`

UnsetCreatedTimeMsecs ensures that no value is present for CreatedTimeMsecs, not even an explicit nil
### GetLogicalCapacityBytes

`func (o *FlashBladeFileSystem) GetLogicalCapacityBytes() int64`

GetLogicalCapacityBytes returns the LogicalCapacityBytes field if non-nil, zero value otherwise.

### GetLogicalCapacityBytesOk

`func (o *FlashBladeFileSystem) GetLogicalCapacityBytesOk() (*int64, bool)`

GetLogicalCapacityBytesOk returns a tuple with the LogicalCapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalCapacityBytes

`func (o *FlashBladeFileSystem) SetLogicalCapacityBytes(v int64)`

SetLogicalCapacityBytes sets LogicalCapacityBytes field to given value.

### HasLogicalCapacityBytes

`func (o *FlashBladeFileSystem) HasLogicalCapacityBytes() bool`

HasLogicalCapacityBytes returns a boolean if a field has been set.

### SetLogicalCapacityBytesNil

`func (o *FlashBladeFileSystem) SetLogicalCapacityBytesNil(b bool)`

 SetLogicalCapacityBytesNil sets the value for LogicalCapacityBytes to be an explicit nil

### UnsetLogicalCapacityBytes
`func (o *FlashBladeFileSystem) UnsetLogicalCapacityBytes()`

UnsetLogicalCapacityBytes ensures that no value is present for LogicalCapacityBytes, not even an explicit nil
### GetLogicalUsedBytes

`func (o *FlashBladeFileSystem) GetLogicalUsedBytes() int64`

GetLogicalUsedBytes returns the LogicalUsedBytes field if non-nil, zero value otherwise.

### GetLogicalUsedBytesOk

`func (o *FlashBladeFileSystem) GetLogicalUsedBytesOk() (*int64, bool)`

GetLogicalUsedBytesOk returns a tuple with the LogicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalUsedBytes

`func (o *FlashBladeFileSystem) SetLogicalUsedBytes(v int64)`

SetLogicalUsedBytes sets LogicalUsedBytes field to given value.

### HasLogicalUsedBytes

`func (o *FlashBladeFileSystem) HasLogicalUsedBytes() bool`

HasLogicalUsedBytes returns a boolean if a field has been set.

### SetLogicalUsedBytesNil

`func (o *FlashBladeFileSystem) SetLogicalUsedBytesNil(b bool)`

 SetLogicalUsedBytesNil sets the value for LogicalUsedBytes to be an explicit nil

### UnsetLogicalUsedBytes
`func (o *FlashBladeFileSystem) UnsetLogicalUsedBytes()`

UnsetLogicalUsedBytes ensures that no value is present for LogicalUsedBytes, not even an explicit nil
### GetNfsInfo

`func (o *FlashBladeFileSystem) GetNfsInfo() FlashBladeNfsInfo`

GetNfsInfo returns the NfsInfo field if non-nil, zero value otherwise.

### GetNfsInfoOk

`func (o *FlashBladeFileSystem) GetNfsInfoOk() (*FlashBladeNfsInfo, bool)`

GetNfsInfoOk returns a tuple with the NfsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNfsInfo

`func (o *FlashBladeFileSystem) SetNfsInfo(v FlashBladeNfsInfo)`

SetNfsInfo sets NfsInfo field to given value.

### HasNfsInfo

`func (o *FlashBladeFileSystem) HasNfsInfo() bool`

HasNfsInfo returns a boolean if a field has been set.

### GetPhysicalUsedBytes

`func (o *FlashBladeFileSystem) GetPhysicalUsedBytes() int64`

GetPhysicalUsedBytes returns the PhysicalUsedBytes field if non-nil, zero value otherwise.

### GetPhysicalUsedBytesOk

`func (o *FlashBladeFileSystem) GetPhysicalUsedBytesOk() (*int64, bool)`

GetPhysicalUsedBytesOk returns a tuple with the PhysicalUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalUsedBytes

`func (o *FlashBladeFileSystem) SetPhysicalUsedBytes(v int64)`

SetPhysicalUsedBytes sets PhysicalUsedBytes field to given value.

### HasPhysicalUsedBytes

`func (o *FlashBladeFileSystem) HasPhysicalUsedBytes() bool`

HasPhysicalUsedBytes returns a boolean if a field has been set.

### SetPhysicalUsedBytesNil

`func (o *FlashBladeFileSystem) SetPhysicalUsedBytesNil(b bool)`

 SetPhysicalUsedBytesNil sets the value for PhysicalUsedBytes to be an explicit nil

### UnsetPhysicalUsedBytes
`func (o *FlashBladeFileSystem) UnsetPhysicalUsedBytes()`

UnsetPhysicalUsedBytes ensures that no value is present for PhysicalUsedBytes, not even an explicit nil
### GetProtocols

`func (o *FlashBladeFileSystem) GetProtocols() []string`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *FlashBladeFileSystem) GetProtocolsOk() (*[]string, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *FlashBladeFileSystem) SetProtocols(v []string)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *FlashBladeFileSystem) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.

### SetProtocolsNil

`func (o *FlashBladeFileSystem) SetProtocolsNil(b bool)`

 SetProtocolsNil sets the value for Protocols to be an explicit nil

### UnsetProtocols
`func (o *FlashBladeFileSystem) UnsetProtocols()`

UnsetProtocols ensures that no value is present for Protocols, not even an explicit nil
### GetUniqueUsedBytes

`func (o *FlashBladeFileSystem) GetUniqueUsedBytes() int64`

GetUniqueUsedBytes returns the UniqueUsedBytes field if non-nil, zero value otherwise.

### GetUniqueUsedBytesOk

`func (o *FlashBladeFileSystem) GetUniqueUsedBytesOk() (*int64, bool)`

GetUniqueUsedBytesOk returns a tuple with the UniqueUsedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueUsedBytes

`func (o *FlashBladeFileSystem) SetUniqueUsedBytes(v int64)`

SetUniqueUsedBytes sets UniqueUsedBytes field to given value.

### HasUniqueUsedBytes

`func (o *FlashBladeFileSystem) HasUniqueUsedBytes() bool`

HasUniqueUsedBytes returns a boolean if a field has been set.

### SetUniqueUsedBytesNil

`func (o *FlashBladeFileSystem) SetUniqueUsedBytesNil(b bool)`

 SetUniqueUsedBytesNil sets the value for UniqueUsedBytes to be an explicit nil

### UnsetUniqueUsedBytes
`func (o *FlashBladeFileSystem) UnsetUniqueUsedBytes()`

UnsetUniqueUsedBytes ensures that no value is present for UniqueUsedBytes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



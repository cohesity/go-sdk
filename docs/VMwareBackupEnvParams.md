# VMwareBackupEnvParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowCrashConsistentSnapshot** | Pointer to **NullableBool** | Whether to fallback to take a crash-consistent snapshot incase taking an app-consistent snapshot fails. | [optional] 
**AllowNbdsslTransportFallback** | Pointer to **NullableBool** | Whether to fallback to use NBDSSL transport for backup in case using SAN transport backup fails. | [optional] 
**AllowVmsWithPhysicalRdmDisks** | Pointer to **NullableBool** | Physical RDM disks cannot be backed up using VADP. By default the backups of such VMs will fail. If this is set to true, then such VMs in this backup job will be backed up by excluding the physical RDM disks. | [optional] 
**VmwareDiskExclusionInfo** | Pointer to [**[]VMwareDiskExclusionProto**](VMwareDiskExclusionProto.md) | List of Virtual Disk(s) to be excluded from the backup job. These disks will be excluded for all VMs in this environment unless overriden by the disk exclusion list from BackupSourceParams.VMwareBackupSourceParams. | [optional] 

## Methods

### NewVMwareBackupEnvParams

`func NewVMwareBackupEnvParams() *VMwareBackupEnvParams`

NewVMwareBackupEnvParams instantiates a new VMwareBackupEnvParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMwareBackupEnvParamsWithDefaults

`func NewVMwareBackupEnvParamsWithDefaults() *VMwareBackupEnvParams`

NewVMwareBackupEnvParamsWithDefaults instantiates a new VMwareBackupEnvParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowCrashConsistentSnapshot

`func (o *VMwareBackupEnvParams) GetAllowCrashConsistentSnapshot() bool`

GetAllowCrashConsistentSnapshot returns the AllowCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetAllowCrashConsistentSnapshotOk

`func (o *VMwareBackupEnvParams) GetAllowCrashConsistentSnapshotOk() (*bool, bool)`

GetAllowCrashConsistentSnapshotOk returns a tuple with the AllowCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCrashConsistentSnapshot

`func (o *VMwareBackupEnvParams) SetAllowCrashConsistentSnapshot(v bool)`

SetAllowCrashConsistentSnapshot sets AllowCrashConsistentSnapshot field to given value.

### HasAllowCrashConsistentSnapshot

`func (o *VMwareBackupEnvParams) HasAllowCrashConsistentSnapshot() bool`

HasAllowCrashConsistentSnapshot returns a boolean if a field has been set.

### SetAllowCrashConsistentSnapshotNil

`func (o *VMwareBackupEnvParams) SetAllowCrashConsistentSnapshotNil(b bool)`

 SetAllowCrashConsistentSnapshotNil sets the value for AllowCrashConsistentSnapshot to be an explicit nil

### UnsetAllowCrashConsistentSnapshot
`func (o *VMwareBackupEnvParams) UnsetAllowCrashConsistentSnapshot()`

UnsetAllowCrashConsistentSnapshot ensures that no value is present for AllowCrashConsistentSnapshot, not even an explicit nil
### GetAllowNbdsslTransportFallback

`func (o *VMwareBackupEnvParams) GetAllowNbdsslTransportFallback() bool`

GetAllowNbdsslTransportFallback returns the AllowNbdsslTransportFallback field if non-nil, zero value otherwise.

### GetAllowNbdsslTransportFallbackOk

`func (o *VMwareBackupEnvParams) GetAllowNbdsslTransportFallbackOk() (*bool, bool)`

GetAllowNbdsslTransportFallbackOk returns a tuple with the AllowNbdsslTransportFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowNbdsslTransportFallback

`func (o *VMwareBackupEnvParams) SetAllowNbdsslTransportFallback(v bool)`

SetAllowNbdsslTransportFallback sets AllowNbdsslTransportFallback field to given value.

### HasAllowNbdsslTransportFallback

`func (o *VMwareBackupEnvParams) HasAllowNbdsslTransportFallback() bool`

HasAllowNbdsslTransportFallback returns a boolean if a field has been set.

### SetAllowNbdsslTransportFallbackNil

`func (o *VMwareBackupEnvParams) SetAllowNbdsslTransportFallbackNil(b bool)`

 SetAllowNbdsslTransportFallbackNil sets the value for AllowNbdsslTransportFallback to be an explicit nil

### UnsetAllowNbdsslTransportFallback
`func (o *VMwareBackupEnvParams) UnsetAllowNbdsslTransportFallback()`

UnsetAllowNbdsslTransportFallback ensures that no value is present for AllowNbdsslTransportFallback, not even an explicit nil
### GetAllowVmsWithPhysicalRdmDisks

`func (o *VMwareBackupEnvParams) GetAllowVmsWithPhysicalRdmDisks() bool`

GetAllowVmsWithPhysicalRdmDisks returns the AllowVmsWithPhysicalRdmDisks field if non-nil, zero value otherwise.

### GetAllowVmsWithPhysicalRdmDisksOk

`func (o *VMwareBackupEnvParams) GetAllowVmsWithPhysicalRdmDisksOk() (*bool, bool)`

GetAllowVmsWithPhysicalRdmDisksOk returns a tuple with the AllowVmsWithPhysicalRdmDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowVmsWithPhysicalRdmDisks

`func (o *VMwareBackupEnvParams) SetAllowVmsWithPhysicalRdmDisks(v bool)`

SetAllowVmsWithPhysicalRdmDisks sets AllowVmsWithPhysicalRdmDisks field to given value.

### HasAllowVmsWithPhysicalRdmDisks

`func (o *VMwareBackupEnvParams) HasAllowVmsWithPhysicalRdmDisks() bool`

HasAllowVmsWithPhysicalRdmDisks returns a boolean if a field has been set.

### SetAllowVmsWithPhysicalRdmDisksNil

`func (o *VMwareBackupEnvParams) SetAllowVmsWithPhysicalRdmDisksNil(b bool)`

 SetAllowVmsWithPhysicalRdmDisksNil sets the value for AllowVmsWithPhysicalRdmDisks to be an explicit nil

### UnsetAllowVmsWithPhysicalRdmDisks
`func (o *VMwareBackupEnvParams) UnsetAllowVmsWithPhysicalRdmDisks()`

UnsetAllowVmsWithPhysicalRdmDisks ensures that no value is present for AllowVmsWithPhysicalRdmDisks, not even an explicit nil
### GetVmwareDiskExclusionInfo

`func (o *VMwareBackupEnvParams) GetVmwareDiskExclusionInfo() []VMwareDiskExclusionProto`

GetVmwareDiskExclusionInfo returns the VmwareDiskExclusionInfo field if non-nil, zero value otherwise.

### GetVmwareDiskExclusionInfoOk

`func (o *VMwareBackupEnvParams) GetVmwareDiskExclusionInfoOk() (*[]VMwareDiskExclusionProto, bool)`

GetVmwareDiskExclusionInfoOk returns a tuple with the VmwareDiskExclusionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareDiskExclusionInfo

`func (o *VMwareBackupEnvParams) SetVmwareDiskExclusionInfo(v []VMwareDiskExclusionProto)`

SetVmwareDiskExclusionInfo sets VmwareDiskExclusionInfo field to given value.

### HasVmwareDiskExclusionInfo

`func (o *VMwareBackupEnvParams) HasVmwareDiskExclusionInfo() bool`

HasVmwareDiskExclusionInfo returns a boolean if a field has been set.

### SetVmwareDiskExclusionInfoNil

`func (o *VMwareBackupEnvParams) SetVmwareDiskExclusionInfoNil(b bool)`

 SetVmwareDiskExclusionInfoNil sets the value for VmwareDiskExclusionInfo to be an explicit nil

### UnsetVmwareDiskExclusionInfo
`func (o *VMwareBackupEnvParams) UnsetVmwareDiskExclusionInfo()`

UnsetVmwareDiskExclusionInfo ensures that no value is present for VmwareDiskExclusionInfo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



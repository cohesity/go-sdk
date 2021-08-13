# VmwareObjectProtectionRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Objects** | [**[]VmwareObjectProtectionRequest**](VmwareObjectProtectionRequest.md) | Specifies the objects to include in the backup. | 
**GlobalExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from the backup. | [optional] 
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. This parameter defaults to true and only changes the behavior of the operation if &#39;appConsistentSnapshot&#39; is set to &#39;true&#39;. | [optional] 
**SkipPhysicalRDMDisks** | Pointer to **NullableBool** | Specifies whether or not to skip backing up physical RDM disks. Physical RDM disks cannot be backed up, so if you attempt to backup a VM with physical RDM disks and this value is set to &#39;false&#39;, then those VM backups will fail. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | If this field is set to true, then the backup for the objects will be performed using dedicated storage area network (SAN) instead of LAN or managment network. | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport backup fails, then backup will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 

## Methods

### NewVmwareObjectProtectionRequestParams

`func NewVmwareObjectProtectionRequestParams(objects []VmwareObjectProtectionRequest, ) *VmwareObjectProtectionRequestParams`

NewVmwareObjectProtectionRequestParams instantiates a new VmwareObjectProtectionRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionRequestParamsWithDefaults

`func NewVmwareObjectProtectionRequestParamsWithDefaults() *VmwareObjectProtectionRequestParams`

NewVmwareObjectProtectionRequestParamsWithDefaults instantiates a new VmwareObjectProtectionRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjects

`func (o *VmwareObjectProtectionRequestParams) GetObjects() []VmwareObjectProtectionRequest`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *VmwareObjectProtectionRequestParams) GetObjectsOk() (*[]VmwareObjectProtectionRequest, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *VmwareObjectProtectionRequestParams) SetObjects(v []VmwareObjectProtectionRequest)`

SetObjects sets Objects field to given value.


### GetGlobalExcludeDisks

`func (o *VmwareObjectProtectionRequestParams) GetGlobalExcludeDisks() []DiskInfo`

GetGlobalExcludeDisks returns the GlobalExcludeDisks field if non-nil, zero value otherwise.

### GetGlobalExcludeDisksOk

`func (o *VmwareObjectProtectionRequestParams) GetGlobalExcludeDisksOk() (*[]DiskInfo, bool)`

GetGlobalExcludeDisksOk returns a tuple with the GlobalExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalExcludeDisks

`func (o *VmwareObjectProtectionRequestParams) SetGlobalExcludeDisks(v []DiskInfo)`

SetGlobalExcludeDisks sets GlobalExcludeDisks field to given value.

### HasGlobalExcludeDisks

`func (o *VmwareObjectProtectionRequestParams) HasGlobalExcludeDisks() bool`

HasGlobalExcludeDisks returns a boolean if a field has been set.

### SetGlobalExcludeDisksNil

`func (o *VmwareObjectProtectionRequestParams) SetGlobalExcludeDisksNil(b bool)`

 SetGlobalExcludeDisksNil sets the value for GlobalExcludeDisks to be an explicit nil

### UnsetGlobalExcludeDisks
`func (o *VmwareObjectProtectionRequestParams) UnsetGlobalExcludeDisks()`

UnsetGlobalExcludeDisks ensures that no value is present for GlobalExcludeDisks, not even an explicit nil
### GetAppConsistentSnapshot

`func (o *VmwareObjectProtectionRequestParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *VmwareObjectProtectionRequestParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *VmwareObjectProtectionRequestParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *VmwareObjectProtectionRequestParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *VmwareObjectProtectionRequestParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *VmwareObjectProtectionRequestParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *VmwareObjectProtectionRequestParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *VmwareObjectProtectionRequestParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *VmwareObjectProtectionRequestParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *VmwareObjectProtectionRequestParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *VmwareObjectProtectionRequestParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *VmwareObjectProtectionRequestParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetSkipPhysicalRDMDisks

`func (o *VmwareObjectProtectionRequestParams) GetSkipPhysicalRDMDisks() bool`

GetSkipPhysicalRDMDisks returns the SkipPhysicalRDMDisks field if non-nil, zero value otherwise.

### GetSkipPhysicalRDMDisksOk

`func (o *VmwareObjectProtectionRequestParams) GetSkipPhysicalRDMDisksOk() (*bool, bool)`

GetSkipPhysicalRDMDisksOk returns a tuple with the SkipPhysicalRDMDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPhysicalRDMDisks

`func (o *VmwareObjectProtectionRequestParams) SetSkipPhysicalRDMDisks(v bool)`

SetSkipPhysicalRDMDisks sets SkipPhysicalRDMDisks field to given value.

### HasSkipPhysicalRDMDisks

`func (o *VmwareObjectProtectionRequestParams) HasSkipPhysicalRDMDisks() bool`

HasSkipPhysicalRDMDisks returns a boolean if a field has been set.

### SetSkipPhysicalRDMDisksNil

`func (o *VmwareObjectProtectionRequestParams) SetSkipPhysicalRDMDisksNil(b bool)`

 SetSkipPhysicalRDMDisksNil sets the value for SkipPhysicalRDMDisks to be an explicit nil

### UnsetSkipPhysicalRDMDisks
`func (o *VmwareObjectProtectionRequestParams) UnsetSkipPhysicalRDMDisks()`

UnsetSkipPhysicalRDMDisks ensures that no value is present for SkipPhysicalRDMDisks, not even an explicit nil
### GetIndexingPolicy

`func (o *VmwareObjectProtectionRequestParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *VmwareObjectProtectionRequestParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *VmwareObjectProtectionRequestParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *VmwareObjectProtectionRequestParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetPrePostScript

`func (o *VmwareObjectProtectionRequestParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *VmwareObjectProtectionRequestParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *VmwareObjectProtectionRequestParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *VmwareObjectProtectionRequestParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetLeverageSanTransport

`func (o *VmwareObjectProtectionRequestParams) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *VmwareObjectProtectionRequestParams) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *VmwareObjectProtectionRequestParams) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *VmwareObjectProtectionRequestParams) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *VmwareObjectProtectionRequestParams) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *VmwareObjectProtectionRequestParams) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *VmwareObjectProtectionRequestParams) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *VmwareObjectProtectionRequestParams) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *VmwareObjectProtectionRequestParams) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *VmwareObjectProtectionRequestParams) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *VmwareObjectProtectionRequestParams) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *VmwareObjectProtectionRequestParams) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



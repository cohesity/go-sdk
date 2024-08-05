# VmwareObjectProtectionResponseParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from being protected. This is only applicable to VM objects. | [optional] 
**TruncateExchangeLogs** | Pointer to **NullableBool** | Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app. | [optional] 
**CdpInfo** | Pointer to [**VmwareCdpObject**](VmwareCdpObject.md) |  | [optional] 
**ExcludeObjectIds** | Pointer to **[]int64** | Specifies the list of IDs of the objects to not be protected in this backup. This field only applies if provided object id is non leaf entity such as Tag or a folder. This can be used to ignore specific objects under a parent object which has been included for protection. | [optional] 
**StandbyInfo** | Pointer to [**VmwareStandbyObject**](VmwareStandbyObject.md) |  | [optional] 
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport backup fails, then backup will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. This parameter defaults to true and only changes the behavior of the operation if &#39;appConsistentSnapshot&#39; is set to &#39;true&#39;. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | If this field is set to true, then the backup for the objects will be performed using dedicated storage area network (SAN) instead of LAN or managment network. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**SkipPhysicalRDMDisks** | Pointer to **NullableBool** | Specifies whether or not to skip backing up physical RDM disks. Physical RDM disks cannot be backed up, so if you attempt to backup a VM with physical RDM disks and this value is set to &#39;false&#39;, then those VM backups will fail. | [optional] 

## Methods

### NewVmwareObjectProtectionResponseParams

`func NewVmwareObjectProtectionResponseParams() *VmwareObjectProtectionResponseParams`

NewVmwareObjectProtectionResponseParams instantiates a new VmwareObjectProtectionResponseParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareObjectProtectionResponseParamsWithDefaults

`func NewVmwareObjectProtectionResponseParamsWithDefaults() *VmwareObjectProtectionResponseParams`

NewVmwareObjectProtectionResponseParamsWithDefaults instantiates a new VmwareObjectProtectionResponseParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeDisks

`func (o *VmwareObjectProtectionResponseParams) GetExcludeDisks() []DiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *VmwareObjectProtectionResponseParams) GetExcludeDisksOk() (*[]DiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *VmwareObjectProtectionResponseParams) SetExcludeDisks(v []DiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *VmwareObjectProtectionResponseParams) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### GetTruncateExchangeLogs

`func (o *VmwareObjectProtectionResponseParams) GetTruncateExchangeLogs() bool`

GetTruncateExchangeLogs returns the TruncateExchangeLogs field if non-nil, zero value otherwise.

### GetTruncateExchangeLogsOk

`func (o *VmwareObjectProtectionResponseParams) GetTruncateExchangeLogsOk() (*bool, bool)`

GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLogs

`func (o *VmwareObjectProtectionResponseParams) SetTruncateExchangeLogs(v bool)`

SetTruncateExchangeLogs sets TruncateExchangeLogs field to given value.

### HasTruncateExchangeLogs

`func (o *VmwareObjectProtectionResponseParams) HasTruncateExchangeLogs() bool`

HasTruncateExchangeLogs returns a boolean if a field has been set.

### SetTruncateExchangeLogsNil

`func (o *VmwareObjectProtectionResponseParams) SetTruncateExchangeLogsNil(b bool)`

 SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil

### UnsetTruncateExchangeLogs
`func (o *VmwareObjectProtectionResponseParams) UnsetTruncateExchangeLogs()`

UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil
### GetCdpInfo

`func (o *VmwareObjectProtectionResponseParams) GetCdpInfo() VmwareCdpObject`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *VmwareObjectProtectionResponseParams) GetCdpInfoOk() (*VmwareCdpObject, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *VmwareObjectProtectionResponseParams) SetCdpInfo(v VmwareCdpObject)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *VmwareObjectProtectionResponseParams) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.

### GetExcludeObjectIds

`func (o *VmwareObjectProtectionResponseParams) GetExcludeObjectIds() []*int64`

GetExcludeObjectIds returns the ExcludeObjectIds field if non-nil, zero value otherwise.

### GetExcludeObjectIdsOk

`func (o *VmwareObjectProtectionResponseParams) GetExcludeObjectIdsOk() (*[]*int64, bool)`

GetExcludeObjectIdsOk returns a tuple with the ExcludeObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeObjectIds

`func (o *VmwareObjectProtectionResponseParams) SetExcludeObjectIds(v []*int64)`

SetExcludeObjectIds sets ExcludeObjectIds field to given value.

### HasExcludeObjectIds

`func (o *VmwareObjectProtectionResponseParams) HasExcludeObjectIds() bool`

HasExcludeObjectIds returns a boolean if a field has been set.

### GetStandbyInfo

`func (o *VmwareObjectProtectionResponseParams) GetStandbyInfo() VmwareStandbyObject`

GetStandbyInfo returns the StandbyInfo field if non-nil, zero value otherwise.

### GetStandbyInfoOk

`func (o *VmwareObjectProtectionResponseParams) GetStandbyInfoOk() (*VmwareStandbyObject, bool)`

GetStandbyInfoOk returns a tuple with the StandbyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInfo

`func (o *VmwareObjectProtectionResponseParams) SetStandbyInfo(v VmwareStandbyObject)`

SetStandbyInfo sets StandbyInfo field to given value.

### HasStandbyInfo

`func (o *VmwareObjectProtectionResponseParams) HasStandbyInfo() bool`

HasStandbyInfo returns a boolean if a field has been set.

### GetAppConsistentSnapshot

`func (o *VmwareObjectProtectionResponseParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *VmwareObjectProtectionResponseParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *VmwareObjectProtectionResponseParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *VmwareObjectProtectionResponseParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *VmwareObjectProtectionResponseParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *VmwareObjectProtectionResponseParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *VmwareObjectProtectionResponseParams) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *VmwareObjectProtectionResponseParams) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *VmwareObjectProtectionResponseParams) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *VmwareObjectProtectionResponseParams) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *VmwareObjectProtectionResponseParams) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *VmwareObjectProtectionResponseParams) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *VmwareObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *VmwareObjectProtectionResponseParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *VmwareObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *VmwareObjectProtectionResponseParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *VmwareObjectProtectionResponseParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *VmwareObjectProtectionResponseParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *VmwareObjectProtectionResponseParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *VmwareObjectProtectionResponseParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *VmwareObjectProtectionResponseParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *VmwareObjectProtectionResponseParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetLeverageSanTransport

`func (o *VmwareObjectProtectionResponseParams) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *VmwareObjectProtectionResponseParams) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *VmwareObjectProtectionResponseParams) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *VmwareObjectProtectionResponseParams) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *VmwareObjectProtectionResponseParams) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *VmwareObjectProtectionResponseParams) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetPrePostScript

`func (o *VmwareObjectProtectionResponseParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *VmwareObjectProtectionResponseParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *VmwareObjectProtectionResponseParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *VmwareObjectProtectionResponseParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetSkipPhysicalRDMDisks

`func (o *VmwareObjectProtectionResponseParams) GetSkipPhysicalRDMDisks() bool`

GetSkipPhysicalRDMDisks returns the SkipPhysicalRDMDisks field if non-nil, zero value otherwise.

### GetSkipPhysicalRDMDisksOk

`func (o *VmwareObjectProtectionResponseParams) GetSkipPhysicalRDMDisksOk() (*bool, bool)`

GetSkipPhysicalRDMDisksOk returns a tuple with the SkipPhysicalRDMDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPhysicalRDMDisks

`func (o *VmwareObjectProtectionResponseParams) SetSkipPhysicalRDMDisks(v bool)`

SetSkipPhysicalRDMDisks sets SkipPhysicalRDMDisks field to given value.

### HasSkipPhysicalRDMDisks

`func (o *VmwareObjectProtectionResponseParams) HasSkipPhysicalRDMDisks() bool`

HasSkipPhysicalRDMDisks returns a boolean if a field has been set.

### SetSkipPhysicalRDMDisksNil

`func (o *VmwareObjectProtectionResponseParams) SetSkipPhysicalRDMDisksNil(b bool)`

 SetSkipPhysicalRDMDisksNil sets the value for SkipPhysicalRDMDisks to be an explicit nil

### UnsetSkipPhysicalRDMDisks
`func (o *VmwareObjectProtectionResponseParams) UnsetSkipPhysicalRDMDisks()`

UnsetSkipPhysicalRDMDisks ensures that no value is present for SkipPhysicalRDMDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



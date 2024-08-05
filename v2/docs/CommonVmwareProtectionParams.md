# CommonVmwareProtectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to quiesce apps and the file system in order to take app consistent snapshots. | [optional] 
**EnableNBDSSLFallback** | Pointer to **NullableBool** | If this field is set to true and SAN transport backup fails, then backup will fallback to use NBDSSL transport. This field only applies if &#39;leverageSanTransport&#39; is set to true. | [optional] 
**FallbackToCrashConsistentSnapshot** | Pointer to **NullableBool** | Specifies whether or not to fallback to a crash consistent snapshot in the event that an app consistent snapshot fails. This parameter defaults to true and only changes the behavior of the operation if &#39;appConsistentSnapshot&#39; is set to &#39;true&#39;. | [optional] 
**IndexingPolicy** | Pointer to [**IndexingPolicy**](IndexingPolicy.md) |  | [optional] 
**LeverageSanTransport** | Pointer to **NullableBool** | If this field is set to true, then the backup for the objects will be performed using dedicated storage area network (SAN) instead of LAN or managment network. | [optional] 
**PrePostScript** | Pointer to [**PrePostScriptParams**](PrePostScriptParams.md) |  | [optional] 
**SkipPhysicalRDMDisks** | Pointer to **NullableBool** | Specifies whether or not to skip backing up physical RDM disks. Physical RDM disks cannot be backed up, so if you attempt to backup a VM with physical RDM disks and this value is set to &#39;false&#39;, then those VM backups will fail. | [optional] 

## Methods

### NewCommonVmwareProtectionParams

`func NewCommonVmwareProtectionParams() *CommonVmwareProtectionParams`

NewCommonVmwareProtectionParams instantiates a new CommonVmwareProtectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonVmwareProtectionParamsWithDefaults

`func NewCommonVmwareProtectionParamsWithDefaults() *CommonVmwareProtectionParams`

NewCommonVmwareProtectionParamsWithDefaults instantiates a new CommonVmwareProtectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppConsistentSnapshot

`func (o *CommonVmwareProtectionParams) GetAppConsistentSnapshot() bool`

GetAppConsistentSnapshot returns the AppConsistentSnapshot field if non-nil, zero value otherwise.

### GetAppConsistentSnapshotOk

`func (o *CommonVmwareProtectionParams) GetAppConsistentSnapshotOk() (*bool, bool)`

GetAppConsistentSnapshotOk returns a tuple with the AppConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppConsistentSnapshot

`func (o *CommonVmwareProtectionParams) SetAppConsistentSnapshot(v bool)`

SetAppConsistentSnapshot sets AppConsistentSnapshot field to given value.

### HasAppConsistentSnapshot

`func (o *CommonVmwareProtectionParams) HasAppConsistentSnapshot() bool`

HasAppConsistentSnapshot returns a boolean if a field has been set.

### SetAppConsistentSnapshotNil

`func (o *CommonVmwareProtectionParams) SetAppConsistentSnapshotNil(b bool)`

 SetAppConsistentSnapshotNil sets the value for AppConsistentSnapshot to be an explicit nil

### UnsetAppConsistentSnapshot
`func (o *CommonVmwareProtectionParams) UnsetAppConsistentSnapshot()`

UnsetAppConsistentSnapshot ensures that no value is present for AppConsistentSnapshot, not even an explicit nil
### GetEnableNBDSSLFallback

`func (o *CommonVmwareProtectionParams) GetEnableNBDSSLFallback() bool`

GetEnableNBDSSLFallback returns the EnableNBDSSLFallback field if non-nil, zero value otherwise.

### GetEnableNBDSSLFallbackOk

`func (o *CommonVmwareProtectionParams) GetEnableNBDSSLFallbackOk() (*bool, bool)`

GetEnableNBDSSLFallbackOk returns a tuple with the EnableNBDSSLFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNBDSSLFallback

`func (o *CommonVmwareProtectionParams) SetEnableNBDSSLFallback(v bool)`

SetEnableNBDSSLFallback sets EnableNBDSSLFallback field to given value.

### HasEnableNBDSSLFallback

`func (o *CommonVmwareProtectionParams) HasEnableNBDSSLFallback() bool`

HasEnableNBDSSLFallback returns a boolean if a field has been set.

### SetEnableNBDSSLFallbackNil

`func (o *CommonVmwareProtectionParams) SetEnableNBDSSLFallbackNil(b bool)`

 SetEnableNBDSSLFallbackNil sets the value for EnableNBDSSLFallback to be an explicit nil

### UnsetEnableNBDSSLFallback
`func (o *CommonVmwareProtectionParams) UnsetEnableNBDSSLFallback()`

UnsetEnableNBDSSLFallback ensures that no value is present for EnableNBDSSLFallback, not even an explicit nil
### GetFallbackToCrashConsistentSnapshot

`func (o *CommonVmwareProtectionParams) GetFallbackToCrashConsistentSnapshot() bool`

GetFallbackToCrashConsistentSnapshot returns the FallbackToCrashConsistentSnapshot field if non-nil, zero value otherwise.

### GetFallbackToCrashConsistentSnapshotOk

`func (o *CommonVmwareProtectionParams) GetFallbackToCrashConsistentSnapshotOk() (*bool, bool)`

GetFallbackToCrashConsistentSnapshotOk returns a tuple with the FallbackToCrashConsistentSnapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackToCrashConsistentSnapshot

`func (o *CommonVmwareProtectionParams) SetFallbackToCrashConsistentSnapshot(v bool)`

SetFallbackToCrashConsistentSnapshot sets FallbackToCrashConsistentSnapshot field to given value.

### HasFallbackToCrashConsistentSnapshot

`func (o *CommonVmwareProtectionParams) HasFallbackToCrashConsistentSnapshot() bool`

HasFallbackToCrashConsistentSnapshot returns a boolean if a field has been set.

### SetFallbackToCrashConsistentSnapshotNil

`func (o *CommonVmwareProtectionParams) SetFallbackToCrashConsistentSnapshotNil(b bool)`

 SetFallbackToCrashConsistentSnapshotNil sets the value for FallbackToCrashConsistentSnapshot to be an explicit nil

### UnsetFallbackToCrashConsistentSnapshot
`func (o *CommonVmwareProtectionParams) UnsetFallbackToCrashConsistentSnapshot()`

UnsetFallbackToCrashConsistentSnapshot ensures that no value is present for FallbackToCrashConsistentSnapshot, not even an explicit nil
### GetIndexingPolicy

`func (o *CommonVmwareProtectionParams) GetIndexingPolicy() IndexingPolicy`

GetIndexingPolicy returns the IndexingPolicy field if non-nil, zero value otherwise.

### GetIndexingPolicyOk

`func (o *CommonVmwareProtectionParams) GetIndexingPolicyOk() (*IndexingPolicy, bool)`

GetIndexingPolicyOk returns a tuple with the IndexingPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingPolicy

`func (o *CommonVmwareProtectionParams) SetIndexingPolicy(v IndexingPolicy)`

SetIndexingPolicy sets IndexingPolicy field to given value.

### HasIndexingPolicy

`func (o *CommonVmwareProtectionParams) HasIndexingPolicy() bool`

HasIndexingPolicy returns a boolean if a field has been set.

### GetLeverageSanTransport

`func (o *CommonVmwareProtectionParams) GetLeverageSanTransport() bool`

GetLeverageSanTransport returns the LeverageSanTransport field if non-nil, zero value otherwise.

### GetLeverageSanTransportOk

`func (o *CommonVmwareProtectionParams) GetLeverageSanTransportOk() (*bool, bool)`

GetLeverageSanTransportOk returns a tuple with the LeverageSanTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeverageSanTransport

`func (o *CommonVmwareProtectionParams) SetLeverageSanTransport(v bool)`

SetLeverageSanTransport sets LeverageSanTransport field to given value.

### HasLeverageSanTransport

`func (o *CommonVmwareProtectionParams) HasLeverageSanTransport() bool`

HasLeverageSanTransport returns a boolean if a field has been set.

### SetLeverageSanTransportNil

`func (o *CommonVmwareProtectionParams) SetLeverageSanTransportNil(b bool)`

 SetLeverageSanTransportNil sets the value for LeverageSanTransport to be an explicit nil

### UnsetLeverageSanTransport
`func (o *CommonVmwareProtectionParams) UnsetLeverageSanTransport()`

UnsetLeverageSanTransport ensures that no value is present for LeverageSanTransport, not even an explicit nil
### GetPrePostScript

`func (o *CommonVmwareProtectionParams) GetPrePostScript() PrePostScriptParams`

GetPrePostScript returns the PrePostScript field if non-nil, zero value otherwise.

### GetPrePostScriptOk

`func (o *CommonVmwareProtectionParams) GetPrePostScriptOk() (*PrePostScriptParams, bool)`

GetPrePostScriptOk returns a tuple with the PrePostScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrePostScript

`func (o *CommonVmwareProtectionParams) SetPrePostScript(v PrePostScriptParams)`

SetPrePostScript sets PrePostScript field to given value.

### HasPrePostScript

`func (o *CommonVmwareProtectionParams) HasPrePostScript() bool`

HasPrePostScript returns a boolean if a field has been set.

### GetSkipPhysicalRDMDisks

`func (o *CommonVmwareProtectionParams) GetSkipPhysicalRDMDisks() bool`

GetSkipPhysicalRDMDisks returns the SkipPhysicalRDMDisks field if non-nil, zero value otherwise.

### GetSkipPhysicalRDMDisksOk

`func (o *CommonVmwareProtectionParams) GetSkipPhysicalRDMDisksOk() (*bool, bool)`

GetSkipPhysicalRDMDisksOk returns a tuple with the SkipPhysicalRDMDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPhysicalRDMDisks

`func (o *CommonVmwareProtectionParams) SetSkipPhysicalRDMDisks(v bool)`

SetSkipPhysicalRDMDisks sets SkipPhysicalRDMDisks field to given value.

### HasSkipPhysicalRDMDisks

`func (o *CommonVmwareProtectionParams) HasSkipPhysicalRDMDisks() bool`

HasSkipPhysicalRDMDisks returns a boolean if a field has been set.

### SetSkipPhysicalRDMDisksNil

`func (o *CommonVmwareProtectionParams) SetSkipPhysicalRDMDisksNil(b bool)`

 SetSkipPhysicalRDMDisksNil sets the value for SkipPhysicalRDMDisks to be an explicit nil

### UnsetSkipPhysicalRDMDisks
`func (o *CommonVmwareProtectionParams) UnsetSkipPhysicalRDMDisks()`

UnsetSkipPhysicalRDMDisks ensures that no value is present for SkipPhysicalRDMDisks, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



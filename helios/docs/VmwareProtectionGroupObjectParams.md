# VmwareProtectionGroupObjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the virtual machine. | [optional] [readonly] 
**IsAutoprotected** | Pointer to **NullableBool** | Specifies whether the vm is part of an Autoprotection. True implies that the vm or its parent directory is autoprotected and will remain part of the autoprotection with additional settings specified here. False implies the object is not part of an Autoprotection and will remain protected and its individual settings here even if a parent directory&#39;s Autoprotection setting is altered. Default is false. | [optional] 
**CdpInfo** | Pointer to [**VmwareCdpObject**](VmwareCdpObject.md) | Specifies the CDP related information for a given object. This field will only be populated when protection group is configured with policy having CDP retnetion settings. | [optional] 
**StandbyInfo** | Pointer to [**VmwareStandbyObject**](VmwareStandbyObject.md) | Specifies the standby related information for a given object. This field will only be populated when standby is configured in backup job settings. | [optional] 
**ExcludeDisks** | Pointer to [**[]DiskInfo**](DiskInfo.md) | Specifies a list of disks to exclude from being protected. This is only applicable to VM objects. | [optional] 
**TruncateExchangeLogs** | Pointer to **NullableBool** | Specifies whether or not to truncate MS Exchange logs while taking an app consistent snapshot of this object. This is only applicable to objects which have a registered MS Exchange app. | [optional] 

## Methods

### NewVmwareProtectionGroupObjectParams

`func NewVmwareProtectionGroupObjectParams(id NullableInt64, ) *VmwareProtectionGroupObjectParams`

NewVmwareProtectionGroupObjectParams instantiates a new VmwareProtectionGroupObjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupObjectParamsWithDefaults

`func NewVmwareProtectionGroupObjectParamsWithDefaults() *VmwareProtectionGroupObjectParams`

NewVmwareProtectionGroupObjectParamsWithDefaults instantiates a new VmwareProtectionGroupObjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmwareProtectionGroupObjectParams) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmwareProtectionGroupObjectParams) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmwareProtectionGroupObjectParams) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *VmwareProtectionGroupObjectParams) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VmwareProtectionGroupObjectParams) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *VmwareProtectionGroupObjectParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmwareProtectionGroupObjectParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmwareProtectionGroupObjectParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmwareProtectionGroupObjectParams) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VmwareProtectionGroupObjectParams) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmwareProtectionGroupObjectParams) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsAutoprotected

`func (o *VmwareProtectionGroupObjectParams) GetIsAutoprotected() bool`

GetIsAutoprotected returns the IsAutoprotected field if non-nil, zero value otherwise.

### GetIsAutoprotectedOk

`func (o *VmwareProtectionGroupObjectParams) GetIsAutoprotectedOk() (*bool, bool)`

GetIsAutoprotectedOk returns a tuple with the IsAutoprotected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprotected

`func (o *VmwareProtectionGroupObjectParams) SetIsAutoprotected(v bool)`

SetIsAutoprotected sets IsAutoprotected field to given value.

### HasIsAutoprotected

`func (o *VmwareProtectionGroupObjectParams) HasIsAutoprotected() bool`

HasIsAutoprotected returns a boolean if a field has been set.

### SetIsAutoprotectedNil

`func (o *VmwareProtectionGroupObjectParams) SetIsAutoprotectedNil(b bool)`

 SetIsAutoprotectedNil sets the value for IsAutoprotected to be an explicit nil

### UnsetIsAutoprotected
`func (o *VmwareProtectionGroupObjectParams) UnsetIsAutoprotected()`

UnsetIsAutoprotected ensures that no value is present for IsAutoprotected, not even an explicit nil
### GetCdpInfo

`func (o *VmwareProtectionGroupObjectParams) GetCdpInfo() VmwareCdpObject`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *VmwareProtectionGroupObjectParams) GetCdpInfoOk() (*VmwareCdpObject, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *VmwareProtectionGroupObjectParams) SetCdpInfo(v VmwareCdpObject)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *VmwareProtectionGroupObjectParams) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.

### GetStandbyInfo

`func (o *VmwareProtectionGroupObjectParams) GetStandbyInfo() VmwareStandbyObject`

GetStandbyInfo returns the StandbyInfo field if non-nil, zero value otherwise.

### GetStandbyInfoOk

`func (o *VmwareProtectionGroupObjectParams) GetStandbyInfoOk() (*VmwareStandbyObject, bool)`

GetStandbyInfoOk returns a tuple with the StandbyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInfo

`func (o *VmwareProtectionGroupObjectParams) SetStandbyInfo(v VmwareStandbyObject)`

SetStandbyInfo sets StandbyInfo field to given value.

### HasStandbyInfo

`func (o *VmwareProtectionGroupObjectParams) HasStandbyInfo() bool`

HasStandbyInfo returns a boolean if a field has been set.

### GetExcludeDisks

`func (o *VmwareProtectionGroupObjectParams) GetExcludeDisks() []DiskInfo`

GetExcludeDisks returns the ExcludeDisks field if non-nil, zero value otherwise.

### GetExcludeDisksOk

`func (o *VmwareProtectionGroupObjectParams) GetExcludeDisksOk() (*[]DiskInfo, bool)`

GetExcludeDisksOk returns a tuple with the ExcludeDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDisks

`func (o *VmwareProtectionGroupObjectParams) SetExcludeDisks(v []DiskInfo)`

SetExcludeDisks sets ExcludeDisks field to given value.

### HasExcludeDisks

`func (o *VmwareProtectionGroupObjectParams) HasExcludeDisks() bool`

HasExcludeDisks returns a boolean if a field has been set.

### GetTruncateExchangeLogs

`func (o *VmwareProtectionGroupObjectParams) GetTruncateExchangeLogs() bool`

GetTruncateExchangeLogs returns the TruncateExchangeLogs field if non-nil, zero value otherwise.

### GetTruncateExchangeLogsOk

`func (o *VmwareProtectionGroupObjectParams) GetTruncateExchangeLogsOk() (*bool, bool)`

GetTruncateExchangeLogsOk returns a tuple with the TruncateExchangeLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncateExchangeLogs

`func (o *VmwareProtectionGroupObjectParams) SetTruncateExchangeLogs(v bool)`

SetTruncateExchangeLogs sets TruncateExchangeLogs field to given value.

### HasTruncateExchangeLogs

`func (o *VmwareProtectionGroupObjectParams) HasTruncateExchangeLogs() bool`

HasTruncateExchangeLogs returns a boolean if a field has been set.

### SetTruncateExchangeLogsNil

`func (o *VmwareProtectionGroupObjectParams) SetTruncateExchangeLogsNil(b bool)`

 SetTruncateExchangeLogsNil sets the value for TruncateExchangeLogs to be an explicit nil

### UnsetTruncateExchangeLogs
`func (o *VmwareProtectionGroupObjectParams) UnsetTruncateExchangeLogs()`

UnsetTruncateExchangeLogs ensures that no value is present for TruncateExchangeLogs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



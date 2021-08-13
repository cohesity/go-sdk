# VmwareProtectionGroupObjectParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object being protected. This can be a leaf level or non leaf level object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the virtual machine. | [optional] [readonly] 
**IsAutoprotected** | Pointer to **NullableBool** | Specifies whether the vm is part of an Autoprotection. True implies that the vm or its parent directory is autoprotected and will remain part of the autoprotection with additional settings specified here. False implies the object is not part of an Autoprotection and will remain protected and its individual settings here even if a parent directory&#39;s Autoprotection setting is altered. Default is false. | [optional] 
**CdpInfo** | Pointer to [**VmwareCdpObject**](VmwareCdpObject.md) | Specifies the CDP related information for a given object. This field will only be populated when protection group is configured with policy having CDP retnetion settings. | [optional] 
**StandbyInfo** | Pointer to [**VmwareStandbyObject**](VmwareStandbyObject.md) | Specifies the standby related information for a given object. This field will only be populated when standby is configured in backup job settings. | [optional] 

## Methods

### NewVmwareProtectionGroupObjectParamsAllOf

`func NewVmwareProtectionGroupObjectParamsAllOf(id NullableInt64, ) *VmwareProtectionGroupObjectParamsAllOf`

NewVmwareProtectionGroupObjectParamsAllOf instantiates a new VmwareProtectionGroupObjectParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareProtectionGroupObjectParamsAllOfWithDefaults

`func NewVmwareProtectionGroupObjectParamsAllOfWithDefaults() *VmwareProtectionGroupObjectParamsAllOf`

NewVmwareProtectionGroupObjectParamsAllOfWithDefaults instantiates a new VmwareProtectionGroupObjectParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VmwareProtectionGroupObjectParamsAllOf) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VmwareProtectionGroupObjectParamsAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VmwareProtectionGroupObjectParamsAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsAutoprotected

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetIsAutoprotected() bool`

GetIsAutoprotected returns the IsAutoprotected field if non-nil, zero value otherwise.

### GetIsAutoprotectedOk

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetIsAutoprotectedOk() (*bool, bool)`

GetIsAutoprotectedOk returns a tuple with the IsAutoprotected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoprotected

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetIsAutoprotected(v bool)`

SetIsAutoprotected sets IsAutoprotected field to given value.

### HasIsAutoprotected

`func (o *VmwareProtectionGroupObjectParamsAllOf) HasIsAutoprotected() bool`

HasIsAutoprotected returns a boolean if a field has been set.

### SetIsAutoprotectedNil

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetIsAutoprotectedNil(b bool)`

 SetIsAutoprotectedNil sets the value for IsAutoprotected to be an explicit nil

### UnsetIsAutoprotected
`func (o *VmwareProtectionGroupObjectParamsAllOf) UnsetIsAutoprotected()`

UnsetIsAutoprotected ensures that no value is present for IsAutoprotected, not even an explicit nil
### GetCdpInfo

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetCdpInfo() VmwareCdpObject`

GetCdpInfo returns the CdpInfo field if non-nil, zero value otherwise.

### GetCdpInfoOk

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetCdpInfoOk() (*VmwareCdpObject, bool)`

GetCdpInfoOk returns a tuple with the CdpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpInfo

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetCdpInfo(v VmwareCdpObject)`

SetCdpInfo sets CdpInfo field to given value.

### HasCdpInfo

`func (o *VmwareProtectionGroupObjectParamsAllOf) HasCdpInfo() bool`

HasCdpInfo returns a boolean if a field has been set.

### GetStandbyInfo

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetStandbyInfo() VmwareStandbyObject`

GetStandbyInfo returns the StandbyInfo field if non-nil, zero value otherwise.

### GetStandbyInfoOk

`func (o *VmwareProtectionGroupObjectParamsAllOf) GetStandbyInfoOk() (*VmwareStandbyObject, bool)`

GetStandbyInfoOk returns a tuple with the StandbyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandbyInfo

`func (o *VmwareProtectionGroupObjectParamsAllOf) SetStandbyInfo(v VmwareStandbyObject)`

SetStandbyInfo sets StandbyInfo field to given value.

### HasStandbyInfo

`func (o *VmwareProtectionGroupObjectParamsAllOf) HasStandbyInfo() bool`

HasStandbyInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



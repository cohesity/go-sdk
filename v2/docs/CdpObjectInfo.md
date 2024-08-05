# CdpObjectInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowReEnableCdp** | Pointer to **NullableBool** | Specifies if re-enabling CDP is allowed or not through UI without any job or policy update through API. | [optional] 
**CdpEnabled** | Pointer to **NullableBool** | Specifies whether CDP is currently active or not. CDP might have been active on this object before, but it might not be anymore. | [optional] 
**LastRunInfo** | Pointer to [**CdpObjectLastRunInfo**](CdpObjectLastRunInfo.md) |  | [optional] 
**ProtectionGroupId** | Pointer to **NullableString** | Specifies the protection group id to which this CDP object belongs. | [optional] [readonly] 

## Methods

### NewCdpObjectInfo

`func NewCdpObjectInfo() *CdpObjectInfo`

NewCdpObjectInfo instantiates a new CdpObjectInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdpObjectInfoWithDefaults

`func NewCdpObjectInfoWithDefaults() *CdpObjectInfo`

NewCdpObjectInfoWithDefaults instantiates a new CdpObjectInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowReEnableCdp

`func (o *CdpObjectInfo) GetAllowReEnableCdp() bool`

GetAllowReEnableCdp returns the AllowReEnableCdp field if non-nil, zero value otherwise.

### GetAllowReEnableCdpOk

`func (o *CdpObjectInfo) GetAllowReEnableCdpOk() (*bool, bool)`

GetAllowReEnableCdpOk returns a tuple with the AllowReEnableCdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReEnableCdp

`func (o *CdpObjectInfo) SetAllowReEnableCdp(v bool)`

SetAllowReEnableCdp sets AllowReEnableCdp field to given value.

### HasAllowReEnableCdp

`func (o *CdpObjectInfo) HasAllowReEnableCdp() bool`

HasAllowReEnableCdp returns a boolean if a field has been set.

### SetAllowReEnableCdpNil

`func (o *CdpObjectInfo) SetAllowReEnableCdpNil(b bool)`

 SetAllowReEnableCdpNil sets the value for AllowReEnableCdp to be an explicit nil

### UnsetAllowReEnableCdp
`func (o *CdpObjectInfo) UnsetAllowReEnableCdp()`

UnsetAllowReEnableCdp ensures that no value is present for AllowReEnableCdp, not even an explicit nil
### GetCdpEnabled

`func (o *CdpObjectInfo) GetCdpEnabled() bool`

GetCdpEnabled returns the CdpEnabled field if non-nil, zero value otherwise.

### GetCdpEnabledOk

`func (o *CdpObjectInfo) GetCdpEnabledOk() (*bool, bool)`

GetCdpEnabledOk returns a tuple with the CdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdpEnabled

`func (o *CdpObjectInfo) SetCdpEnabled(v bool)`

SetCdpEnabled sets CdpEnabled field to given value.

### HasCdpEnabled

`func (o *CdpObjectInfo) HasCdpEnabled() bool`

HasCdpEnabled returns a boolean if a field has been set.

### SetCdpEnabledNil

`func (o *CdpObjectInfo) SetCdpEnabledNil(b bool)`

 SetCdpEnabledNil sets the value for CdpEnabled to be an explicit nil

### UnsetCdpEnabled
`func (o *CdpObjectInfo) UnsetCdpEnabled()`

UnsetCdpEnabled ensures that no value is present for CdpEnabled, not even an explicit nil
### GetLastRunInfo

`func (o *CdpObjectInfo) GetLastRunInfo() CdpObjectLastRunInfo`

GetLastRunInfo returns the LastRunInfo field if non-nil, zero value otherwise.

### GetLastRunInfoOk

`func (o *CdpObjectInfo) GetLastRunInfoOk() (*CdpObjectLastRunInfo, bool)`

GetLastRunInfoOk returns a tuple with the LastRunInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunInfo

`func (o *CdpObjectInfo) SetLastRunInfo(v CdpObjectLastRunInfo)`

SetLastRunInfo sets LastRunInfo field to given value.

### HasLastRunInfo

`func (o *CdpObjectInfo) HasLastRunInfo() bool`

HasLastRunInfo returns a boolean if a field has been set.

### GetProtectionGroupId

`func (o *CdpObjectInfo) GetProtectionGroupId() string`

GetProtectionGroupId returns the ProtectionGroupId field if non-nil, zero value otherwise.

### GetProtectionGroupIdOk

`func (o *CdpObjectInfo) GetProtectionGroupIdOk() (*string, bool)`

GetProtectionGroupIdOk returns a tuple with the ProtectionGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionGroupId

`func (o *CdpObjectInfo) SetProtectionGroupId(v string)`

SetProtectionGroupId sets ProtectionGroupId field to given value.

### HasProtectionGroupId

`func (o *CdpObjectInfo) HasProtectionGroupId() bool`

HasProtectionGroupId returns a boolean if a field has been set.

### SetProtectionGroupIdNil

`func (o *CdpObjectInfo) SetProtectionGroupIdNil(b bool)`

 SetProtectionGroupIdNil sets the value for ProtectionGroupId to be an explicit nil

### UnsetProtectionGroupId
`func (o *CdpObjectInfo) UnsetProtectionGroupId()`

UnsetProtectionGroupId ensures that no value is present for ProtectionGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



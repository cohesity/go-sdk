# CommonRecoverObjectSnapshotParamsArchivalTargetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivalTaskId** | Pointer to **NullableString** | Specifies the archival task id. This is a protection group UID which only applies when archival type is &#39;Tape&#39;. | [optional] 
**OwnershipContext** | Pointer to **NullableString** | Specifies the ownership context for the target. | [optional] 
**TargetId** | Pointer to **NullableInt64** | Specifies the archival target ID. | [optional] 
**TargetName** | Pointer to **NullableString** | Specifies the archival target name. | [optional] 
**TargetType** | Pointer to **NullableString** | Specifies the archival target type. | [optional] 
**TierSettings** | Pointer to [**ArchivalTargetTierInfo**](ArchivalTargetTierInfo.md) |  | [optional] 
**UsageType** | Pointer to **NullableString** | Specifies the usage type for the target. | [optional] 

## Methods

### NewCommonRecoverObjectSnapshotParamsArchivalTargetInfo

`func NewCommonRecoverObjectSnapshotParamsArchivalTargetInfo() *CommonRecoverObjectSnapshotParamsArchivalTargetInfo`

NewCommonRecoverObjectSnapshotParamsArchivalTargetInfo instantiates a new CommonRecoverObjectSnapshotParamsArchivalTargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonRecoverObjectSnapshotParamsArchivalTargetInfoWithDefaults

`func NewCommonRecoverObjectSnapshotParamsArchivalTargetInfoWithDefaults() *CommonRecoverObjectSnapshotParamsArchivalTargetInfo`

NewCommonRecoverObjectSnapshotParamsArchivalTargetInfoWithDefaults instantiates a new CommonRecoverObjectSnapshotParamsArchivalTargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTaskId

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetArchivalTaskId() string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetArchivalTaskIdOk() (*string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetArchivalTaskId(v string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetOwnershipContext

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetTargetId

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTierSettings() ArchivalTargetTierInfo`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetTierSettings(v ArchivalTargetTierInfo)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetUsageType

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *CommonRecoverObjectSnapshotParamsArchivalTargetInfo) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



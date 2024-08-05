# ArchivalTargetSummaryInfo

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

### NewArchivalTargetSummaryInfo

`func NewArchivalTargetSummaryInfo() *ArchivalTargetSummaryInfo`

NewArchivalTargetSummaryInfo instantiates a new ArchivalTargetSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchivalTargetSummaryInfoWithDefaults

`func NewArchivalTargetSummaryInfoWithDefaults() *ArchivalTargetSummaryInfo`

NewArchivalTargetSummaryInfoWithDefaults instantiates a new ArchivalTargetSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTaskId

`func (o *ArchivalTargetSummaryInfo) GetArchivalTaskId() string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *ArchivalTargetSummaryInfo) GetArchivalTaskIdOk() (*string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *ArchivalTargetSummaryInfo) SetArchivalTaskId(v string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *ArchivalTargetSummaryInfo) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *ArchivalTargetSummaryInfo) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *ArchivalTargetSummaryInfo) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetOwnershipContext

`func (o *ArchivalTargetSummaryInfo) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *ArchivalTargetSummaryInfo) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *ArchivalTargetSummaryInfo) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *ArchivalTargetSummaryInfo) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *ArchivalTargetSummaryInfo) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *ArchivalTargetSummaryInfo) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetTargetId

`func (o *ArchivalTargetSummaryInfo) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ArchivalTargetSummaryInfo) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ArchivalTargetSummaryInfo) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ArchivalTargetSummaryInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *ArchivalTargetSummaryInfo) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ArchivalTargetSummaryInfo) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *ArchivalTargetSummaryInfo) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ArchivalTargetSummaryInfo) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ArchivalTargetSummaryInfo) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ArchivalTargetSummaryInfo) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *ArchivalTargetSummaryInfo) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *ArchivalTargetSummaryInfo) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *ArchivalTargetSummaryInfo) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ArchivalTargetSummaryInfo) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ArchivalTargetSummaryInfo) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ArchivalTargetSummaryInfo) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ArchivalTargetSummaryInfo) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ArchivalTargetSummaryInfo) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *ArchivalTargetSummaryInfo) GetTierSettings() ArchivalTargetTierInfo`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *ArchivalTargetSummaryInfo) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *ArchivalTargetSummaryInfo) SetTierSettings(v ArchivalTargetTierInfo)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *ArchivalTargetSummaryInfo) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetUsageType

`func (o *ArchivalTargetSummaryInfo) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ArchivalTargetSummaryInfo) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ArchivalTargetSummaryInfo) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ArchivalTargetSummaryInfo) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *ArchivalTargetSummaryInfo) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *ArchivalTargetSummaryInfo) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



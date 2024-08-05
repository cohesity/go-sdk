# ObjectSnapshotExternalTargetInfo

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

### NewObjectSnapshotExternalTargetInfo

`func NewObjectSnapshotExternalTargetInfo() *ObjectSnapshotExternalTargetInfo`

NewObjectSnapshotExternalTargetInfo instantiates a new ObjectSnapshotExternalTargetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectSnapshotExternalTargetInfoWithDefaults

`func NewObjectSnapshotExternalTargetInfoWithDefaults() *ObjectSnapshotExternalTargetInfo`

NewObjectSnapshotExternalTargetInfoWithDefaults instantiates a new ObjectSnapshotExternalTargetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTaskId

`func (o *ObjectSnapshotExternalTargetInfo) GetArchivalTaskId() string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *ObjectSnapshotExternalTargetInfo) GetArchivalTaskIdOk() (*string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *ObjectSnapshotExternalTargetInfo) SetArchivalTaskId(v string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *ObjectSnapshotExternalTargetInfo) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *ObjectSnapshotExternalTargetInfo) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *ObjectSnapshotExternalTargetInfo) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetOwnershipContext

`func (o *ObjectSnapshotExternalTargetInfo) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *ObjectSnapshotExternalTargetInfo) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *ObjectSnapshotExternalTargetInfo) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *ObjectSnapshotExternalTargetInfo) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *ObjectSnapshotExternalTargetInfo) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *ObjectSnapshotExternalTargetInfo) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetTargetId

`func (o *ObjectSnapshotExternalTargetInfo) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ObjectSnapshotExternalTargetInfo) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ObjectSnapshotExternalTargetInfo) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ObjectSnapshotExternalTargetInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *ObjectSnapshotExternalTargetInfo) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ObjectSnapshotExternalTargetInfo) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *ObjectSnapshotExternalTargetInfo) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ObjectSnapshotExternalTargetInfo) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ObjectSnapshotExternalTargetInfo) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ObjectSnapshotExternalTargetInfo) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *ObjectSnapshotExternalTargetInfo) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *ObjectSnapshotExternalTargetInfo) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *ObjectSnapshotExternalTargetInfo) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ObjectSnapshotExternalTargetInfo) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ObjectSnapshotExternalTargetInfo) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ObjectSnapshotExternalTargetInfo) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ObjectSnapshotExternalTargetInfo) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ObjectSnapshotExternalTargetInfo) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *ObjectSnapshotExternalTargetInfo) GetTierSettings() ArchivalTargetTierInfo`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *ObjectSnapshotExternalTargetInfo) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *ObjectSnapshotExternalTargetInfo) SetTierSettings(v ArchivalTargetTierInfo)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *ObjectSnapshotExternalTargetInfo) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetUsageType

`func (o *ObjectSnapshotExternalTargetInfo) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ObjectSnapshotExternalTargetInfo) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ObjectSnapshotExternalTargetInfo) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ObjectSnapshotExternalTargetInfo) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *ObjectSnapshotExternalTargetInfo) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *ObjectSnapshotExternalTargetInfo) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



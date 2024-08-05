# ObjectArchivalSnapshotInfo

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
**LogicalSizeBytes** | Pointer to **NullableInt64** | Specifies the logical size of this snapshot in bytes. | [optional] 
**SnapshotId** | Pointer to **NullableString** | Specifies the id of the archival snapshot for the object. | [optional] 

## Methods

### NewObjectArchivalSnapshotInfo

`func NewObjectArchivalSnapshotInfo() *ObjectArchivalSnapshotInfo`

NewObjectArchivalSnapshotInfo instantiates a new ObjectArchivalSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectArchivalSnapshotInfoWithDefaults

`func NewObjectArchivalSnapshotInfoWithDefaults() *ObjectArchivalSnapshotInfo`

NewObjectArchivalSnapshotInfoWithDefaults instantiates a new ObjectArchivalSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivalTaskId

`func (o *ObjectArchivalSnapshotInfo) GetArchivalTaskId() string`

GetArchivalTaskId returns the ArchivalTaskId field if non-nil, zero value otherwise.

### GetArchivalTaskIdOk

`func (o *ObjectArchivalSnapshotInfo) GetArchivalTaskIdOk() (*string, bool)`

GetArchivalTaskIdOk returns a tuple with the ArchivalTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivalTaskId

`func (o *ObjectArchivalSnapshotInfo) SetArchivalTaskId(v string)`

SetArchivalTaskId sets ArchivalTaskId field to given value.

### HasArchivalTaskId

`func (o *ObjectArchivalSnapshotInfo) HasArchivalTaskId() bool`

HasArchivalTaskId returns a boolean if a field has been set.

### SetArchivalTaskIdNil

`func (o *ObjectArchivalSnapshotInfo) SetArchivalTaskIdNil(b bool)`

 SetArchivalTaskIdNil sets the value for ArchivalTaskId to be an explicit nil

### UnsetArchivalTaskId
`func (o *ObjectArchivalSnapshotInfo) UnsetArchivalTaskId()`

UnsetArchivalTaskId ensures that no value is present for ArchivalTaskId, not even an explicit nil
### GetOwnershipContext

`func (o *ObjectArchivalSnapshotInfo) GetOwnershipContext() string`

GetOwnershipContext returns the OwnershipContext field if non-nil, zero value otherwise.

### GetOwnershipContextOk

`func (o *ObjectArchivalSnapshotInfo) GetOwnershipContextOk() (*string, bool)`

GetOwnershipContextOk returns a tuple with the OwnershipContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnershipContext

`func (o *ObjectArchivalSnapshotInfo) SetOwnershipContext(v string)`

SetOwnershipContext sets OwnershipContext field to given value.

### HasOwnershipContext

`func (o *ObjectArchivalSnapshotInfo) HasOwnershipContext() bool`

HasOwnershipContext returns a boolean if a field has been set.

### SetOwnershipContextNil

`func (o *ObjectArchivalSnapshotInfo) SetOwnershipContextNil(b bool)`

 SetOwnershipContextNil sets the value for OwnershipContext to be an explicit nil

### UnsetOwnershipContext
`func (o *ObjectArchivalSnapshotInfo) UnsetOwnershipContext()`

UnsetOwnershipContext ensures that no value is present for OwnershipContext, not even an explicit nil
### GetTargetId

`func (o *ObjectArchivalSnapshotInfo) GetTargetId() int64`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ObjectArchivalSnapshotInfo) GetTargetIdOk() (*int64, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ObjectArchivalSnapshotInfo) SetTargetId(v int64)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ObjectArchivalSnapshotInfo) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *ObjectArchivalSnapshotInfo) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ObjectArchivalSnapshotInfo) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetTargetName

`func (o *ObjectArchivalSnapshotInfo) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *ObjectArchivalSnapshotInfo) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *ObjectArchivalSnapshotInfo) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *ObjectArchivalSnapshotInfo) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### SetTargetNameNil

`func (o *ObjectArchivalSnapshotInfo) SetTargetNameNil(b bool)`

 SetTargetNameNil sets the value for TargetName to be an explicit nil

### UnsetTargetName
`func (o *ObjectArchivalSnapshotInfo) UnsetTargetName()`

UnsetTargetName ensures that no value is present for TargetName, not even an explicit nil
### GetTargetType

`func (o *ObjectArchivalSnapshotInfo) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ObjectArchivalSnapshotInfo) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ObjectArchivalSnapshotInfo) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ObjectArchivalSnapshotInfo) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ObjectArchivalSnapshotInfo) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ObjectArchivalSnapshotInfo) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTierSettings

`func (o *ObjectArchivalSnapshotInfo) GetTierSettings() ArchivalTargetTierInfo`

GetTierSettings returns the TierSettings field if non-nil, zero value otherwise.

### GetTierSettingsOk

`func (o *ObjectArchivalSnapshotInfo) GetTierSettingsOk() (*ArchivalTargetTierInfo, bool)`

GetTierSettingsOk returns a tuple with the TierSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSettings

`func (o *ObjectArchivalSnapshotInfo) SetTierSettings(v ArchivalTargetTierInfo)`

SetTierSettings sets TierSettings field to given value.

### HasTierSettings

`func (o *ObjectArchivalSnapshotInfo) HasTierSettings() bool`

HasTierSettings returns a boolean if a field has been set.

### GetUsageType

`func (o *ObjectArchivalSnapshotInfo) GetUsageType() string`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *ObjectArchivalSnapshotInfo) GetUsageTypeOk() (*string, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *ObjectArchivalSnapshotInfo) SetUsageType(v string)`

SetUsageType sets UsageType field to given value.

### HasUsageType

`func (o *ObjectArchivalSnapshotInfo) HasUsageType() bool`

HasUsageType returns a boolean if a field has been set.

### SetUsageTypeNil

`func (o *ObjectArchivalSnapshotInfo) SetUsageTypeNil(b bool)`

 SetUsageTypeNil sets the value for UsageType to be an explicit nil

### UnsetUsageType
`func (o *ObjectArchivalSnapshotInfo) UnsetUsageType()`

UnsetUsageType ensures that no value is present for UsageType, not even an explicit nil
### GetLogicalSizeBytes

`func (o *ObjectArchivalSnapshotInfo) GetLogicalSizeBytes() int64`

GetLogicalSizeBytes returns the LogicalSizeBytes field if non-nil, zero value otherwise.

### GetLogicalSizeBytesOk

`func (o *ObjectArchivalSnapshotInfo) GetLogicalSizeBytesOk() (*int64, bool)`

GetLogicalSizeBytesOk returns a tuple with the LogicalSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalSizeBytes

`func (o *ObjectArchivalSnapshotInfo) SetLogicalSizeBytes(v int64)`

SetLogicalSizeBytes sets LogicalSizeBytes field to given value.

### HasLogicalSizeBytes

`func (o *ObjectArchivalSnapshotInfo) HasLogicalSizeBytes() bool`

HasLogicalSizeBytes returns a boolean if a field has been set.

### SetLogicalSizeBytesNil

`func (o *ObjectArchivalSnapshotInfo) SetLogicalSizeBytesNil(b bool)`

 SetLogicalSizeBytesNil sets the value for LogicalSizeBytes to be an explicit nil

### UnsetLogicalSizeBytes
`func (o *ObjectArchivalSnapshotInfo) UnsetLogicalSizeBytes()`

UnsetLogicalSizeBytes ensures that no value is present for LogicalSizeBytes, not even an explicit nil
### GetSnapshotId

`func (o *ObjectArchivalSnapshotInfo) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *ObjectArchivalSnapshotInfo) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *ObjectArchivalSnapshotInfo) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *ObjectArchivalSnapshotInfo) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### SetSnapshotIdNil

`func (o *ObjectArchivalSnapshotInfo) SetSnapshotIdNil(b bool)`

 SetSnapshotIdNil sets the value for SnapshotId to be an explicit nil

### UnsetSnapshotId
`func (o *ObjectArchivalSnapshotInfo) UnsetSnapshotId()`

UnsetSnapshotId ensures that no value is present for SnapshotId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



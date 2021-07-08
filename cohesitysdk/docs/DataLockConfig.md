# DataLockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DaysToKeep** | Pointer to **NullableInt64** | Specifies last N days to keep Snapshots under datalock in a protection group. | [optional] 
**WormRetentionType** | Pointer to **NullableString** | Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Jobs using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. &#39;kNone&#39; implies there is no WORM retention set. &#39;kCompliance&#39; implies WORM retention is set for compliance reason. &#39;kAdministrative&#39; implies WORM retention is set for administrative purposes. | [optional] 

## Methods

### NewDataLockConfig

`func NewDataLockConfig() *DataLockConfig`

NewDataLockConfig instantiates a new DataLockConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLockConfigWithDefaults

`func NewDataLockConfigWithDefaults() *DataLockConfig`

NewDataLockConfigWithDefaults instantiates a new DataLockConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaysToKeep

`func (o *DataLockConfig) GetDaysToKeep() int64`

GetDaysToKeep returns the DaysToKeep field if non-nil, zero value otherwise.

### GetDaysToKeepOk

`func (o *DataLockConfig) GetDaysToKeepOk() (*int64, bool)`

GetDaysToKeepOk returns a tuple with the DaysToKeep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToKeep

`func (o *DataLockConfig) SetDaysToKeep(v int64)`

SetDaysToKeep sets DaysToKeep field to given value.

### HasDaysToKeep

`func (o *DataLockConfig) HasDaysToKeep() bool`

HasDaysToKeep returns a boolean if a field has been set.

### SetDaysToKeepNil

`func (o *DataLockConfig) SetDaysToKeepNil(b bool)`

 SetDaysToKeepNil sets the value for DaysToKeep to be an explicit nil

### UnsetDaysToKeep
`func (o *DataLockConfig) UnsetDaysToKeep()`

UnsetDaysToKeep ensures that no value is present for DaysToKeep, not even an explicit nil
### GetWormRetentionType

`func (o *DataLockConfig) GetWormRetentionType() string`

GetWormRetentionType returns the WormRetentionType field if non-nil, zero value otherwise.

### GetWormRetentionTypeOk

`func (o *DataLockConfig) GetWormRetentionTypeOk() (*string, bool)`

GetWormRetentionTypeOk returns a tuple with the WormRetentionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWormRetentionType

`func (o *DataLockConfig) SetWormRetentionType(v string)`

SetWormRetentionType sets WormRetentionType field to given value.

### HasWormRetentionType

`func (o *DataLockConfig) HasWormRetentionType() bool`

HasWormRetentionType returns a boolean if a field has been set.

### SetWormRetentionTypeNil

`func (o *DataLockConfig) SetWormRetentionTypeNil(b bool)`

 SetWormRetentionTypeNil sets the value for WormRetentionType to be an explicit nil

### UnsetWormRetentionType
`func (o *DataLockConfig) UnsetWormRetentionType()`

UnsetWormRetentionType ensures that no value is present for WormRetentionType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



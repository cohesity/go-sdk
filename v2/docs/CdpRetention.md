# CdpRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataLockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**Duration** | **NullableInt32** | Specifies the duration for a cdp backup retention. | 
**Unit** | **NullableString** | Specificies the Retention Unit of a CDP backup measured in minutes or hours. | 

## Methods

### NewCdpRetention

`func NewCdpRetention(duration NullableInt32, unit NullableString, ) *CdpRetention`

NewCdpRetention instantiates a new CdpRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdpRetentionWithDefaults

`func NewCdpRetentionWithDefaults() *CdpRetention`

NewCdpRetentionWithDefaults instantiates a new CdpRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataLockConfig

`func (o *CdpRetention) GetDataLockConfig() DataLockConfig`

GetDataLockConfig returns the DataLockConfig field if non-nil, zero value otherwise.

### GetDataLockConfigOk

`func (o *CdpRetention) GetDataLockConfigOk() (*DataLockConfig, bool)`

GetDataLockConfigOk returns a tuple with the DataLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConfig

`func (o *CdpRetention) SetDataLockConfig(v DataLockConfig)`

SetDataLockConfig sets DataLockConfig field to given value.

### HasDataLockConfig

`func (o *CdpRetention) HasDataLockConfig() bool`

HasDataLockConfig returns a boolean if a field has been set.

### GetDuration

`func (o *CdpRetention) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CdpRetention) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CdpRetention) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### SetDurationNil

`func (o *CdpRetention) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *CdpRetention) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetUnit

`func (o *CdpRetention) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CdpRetention) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CdpRetention) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *CdpRetention) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *CdpRetention) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



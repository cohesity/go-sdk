# HeliosCdpRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unit** | Pointer to **NullableString** | Specificies the Retention Unit of a CDP backup measured in minutes or hours. | [optional] 
**Duration** | Pointer to **NullableInt32** | Specifies the duration for a cdp backup retention. | [optional] 
**DataLockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 

## Methods

### NewHeliosCdpRetention

`func NewHeliosCdpRetention() *HeliosCdpRetention`

NewHeliosCdpRetention instantiates a new HeliosCdpRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosCdpRetentionWithDefaults

`func NewHeliosCdpRetentionWithDefaults() *HeliosCdpRetention`

NewHeliosCdpRetentionWithDefaults instantiates a new HeliosCdpRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnit

`func (o *HeliosCdpRetention) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosCdpRetention) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosCdpRetention) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosCdpRetention) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosCdpRetention) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosCdpRetention) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetDuration

`func (o *HeliosCdpRetention) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HeliosCdpRetention) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HeliosCdpRetention) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HeliosCdpRetention) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *HeliosCdpRetention) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *HeliosCdpRetention) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetDataLockConfig

`func (o *HeliosCdpRetention) GetDataLockConfig() DataLockConfig`

GetDataLockConfig returns the DataLockConfig field if non-nil, zero value otherwise.

### GetDataLockConfigOk

`func (o *HeliosCdpRetention) GetDataLockConfigOk() (*DataLockConfig, bool)`

GetDataLockConfigOk returns a tuple with the DataLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConfig

`func (o *HeliosCdpRetention) SetDataLockConfig(v DataLockConfig)`

SetDataLockConfig sets DataLockConfig field to given value.

### HasDataLockConfig

`func (o *HeliosCdpRetention) HasDataLockConfig() bool`

HasDataLockConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Retention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataLockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**Duration** | **NullableInt64** | Specifies the duration for a backup retention. &lt;br&gt; Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 &#x3D; 210 days. | 
**Unit** | **NullableString** | Specificies the Retention Unit of a backup measured in days, months or years. &lt;br&gt; If unit is &#39;Months&#39;, then number specified in duration is multiplied to 30. &lt;br&gt; Example: If duration is 4 and unit is &#39;Months&#39; then number of retention days will be 30 * 4 &#x3D; 120 days. &lt;br&gt; If unit is &#39;Years&#39;, then number specified in duration is multiplied to 365. &lt;br&gt; If duration is 2 and unit is &#39;Months&#39; then number of retention days will be 365 * 2 &#x3D; 730 days. | 

## Methods

### NewRetention

`func NewRetention(duration NullableInt64, unit NullableString, ) *Retention`

NewRetention instantiates a new Retention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetentionWithDefaults

`func NewRetentionWithDefaults() *Retention`

NewRetentionWithDefaults instantiates a new Retention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataLockConfig

`func (o *Retention) GetDataLockConfig() DataLockConfig`

GetDataLockConfig returns the DataLockConfig field if non-nil, zero value otherwise.

### GetDataLockConfigOk

`func (o *Retention) GetDataLockConfigOk() (*DataLockConfig, bool)`

GetDataLockConfigOk returns a tuple with the DataLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConfig

`func (o *Retention) SetDataLockConfig(v DataLockConfig)`

SetDataLockConfig sets DataLockConfig field to given value.

### HasDataLockConfig

`func (o *Retention) HasDataLockConfig() bool`

HasDataLockConfig returns a boolean if a field has been set.

### GetDuration

`func (o *Retention) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Retention) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Retention) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### SetDurationNil

`func (o *Retention) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *Retention) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetUnit

`func (o *Retention) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Retention) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Retention) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *Retention) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *Retention) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



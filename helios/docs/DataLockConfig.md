# DataLockConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **NullableString** | Specifies the type of WORM retention type.  &#39;Compliance&#39; implies WORM retention is set for compliance reason.  &#39;Administrative&#39; implies WORM retention is set for administrative purposes. | 
**Unit** | **NullableString** | Specificies the Retention Unit of a dataLock measured in days, months or years. &lt;br&gt; If unit is &#39;Months&#39;, then number specified in duration is multiplied to 30. &lt;br&gt; Example: If duration is 4 and unit is &#39;Months&#39; then number of retention days will be 30 * 4 &#x3D; 120 days. &lt;br&gt; If unit is &#39;Years&#39;, then number specified in duration is multiplied to 365. &lt;br&gt; If duration is 2 and unit is &#39;Months&#39; then number of retention days will be 365 * 2 &#x3D; 730 days. | 
**Duration** | **NullableInt64** | Specifies the duration for a dataLock. &lt;br&gt; Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 &#x3D; 210 days of the backup. | 

## Methods

### NewDataLockConfig

`func NewDataLockConfig(mode NullableString, unit NullableString, duration NullableInt64, ) *DataLockConfig`

NewDataLockConfig instantiates a new DataLockConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLockConfigWithDefaults

`func NewDataLockConfigWithDefaults() *DataLockConfig`

NewDataLockConfigWithDefaults instantiates a new DataLockConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *DataLockConfig) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *DataLockConfig) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *DataLockConfig) SetMode(v string)`

SetMode sets Mode field to given value.


### SetModeNil

`func (o *DataLockConfig) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *DataLockConfig) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetUnit

`func (o *DataLockConfig) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DataLockConfig) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DataLockConfig) SetUnit(v string)`

SetUnit sets Unit field to given value.


### SetUnitNil

`func (o *DataLockConfig) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *DataLockConfig) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil
### GetDuration

`func (o *DataLockConfig) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DataLockConfig) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DataLockConfig) SetDuration(v int64)`

SetDuration sets Duration field to given value.


### SetDurationNil

`func (o *DataLockConfig) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *DataLockConfig) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



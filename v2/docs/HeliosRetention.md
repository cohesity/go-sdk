# HeliosRetention

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataLockConfig** | Pointer to [**DataLockConfig**](DataLockConfig.md) |  | [optional] 
**Duration** | Pointer to **NullableInt64** | Specifies the duration for a backup retention. &lt;br&gt; Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 &#x3D; 210 days. | [optional] 
**Tiers** | Pointer to [**[]HeliosTier**](HeliosTier.md) | Specifies the list of tiers where backup will be moved. This will be populated only if poilcy type is DMaaS. | [optional] 
**Unit** | Pointer to **NullableString** | Specificies the Retention Unit of a backup measured in days, months or years. &lt;br&gt; If unit is &#39;Months&#39;, then number specified in duration is multiplied to 30. &lt;br&gt; Example: If duration is 4 and unit is &#39;Months&#39; then number of retention days will be 30 * 4 &#x3D; 120 days. &lt;br&gt; If unit is &#39;Years&#39;, then number specified in duration is multiplied to 365. &lt;br&gt; If duration is 2 and unit is &#39;Months&#39; then number of retention days will be 365 * 2 &#x3D; 730 days. | [optional] 

## Methods

### NewHeliosRetention

`func NewHeliosRetention() *HeliosRetention`

NewHeliosRetention instantiates a new HeliosRetention object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosRetentionWithDefaults

`func NewHeliosRetentionWithDefaults() *HeliosRetention`

NewHeliosRetentionWithDefaults instantiates a new HeliosRetention object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataLockConfig

`func (o *HeliosRetention) GetDataLockConfig() DataLockConfig`

GetDataLockConfig returns the DataLockConfig field if non-nil, zero value otherwise.

### GetDataLockConfigOk

`func (o *HeliosRetention) GetDataLockConfigOk() (*DataLockConfig, bool)`

GetDataLockConfigOk returns a tuple with the DataLockConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLockConfig

`func (o *HeliosRetention) SetDataLockConfig(v DataLockConfig)`

SetDataLockConfig sets DataLockConfig field to given value.

### HasDataLockConfig

`func (o *HeliosRetention) HasDataLockConfig() bool`

HasDataLockConfig returns a boolean if a field has been set.

### GetDuration

`func (o *HeliosRetention) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *HeliosRetention) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *HeliosRetention) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *HeliosRetention) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *HeliosRetention) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *HeliosRetention) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetTiers

`func (o *HeliosRetention) GetTiers() []HeliosTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *HeliosRetention) GetTiersOk() (*[]HeliosTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *HeliosRetention) SetTiers(v []HeliosTier)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *HeliosRetention) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetUnit

`func (o *HeliosRetention) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *HeliosRetention) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *HeliosRetention) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *HeliosRetention) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### SetUnitNil

`func (o *HeliosRetention) SetUnitNil(b bool)`

 SetUnitNil sets the value for Unit to be an explicit nil

### UnsetUnit
`func (o *HeliosRetention) UnsetUnit()`

UnsetUnit ensures that no value is present for Unit, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



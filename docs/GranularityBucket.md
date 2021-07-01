# GranularityBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Granularity** | Pointer to **NullableInt32** | The base time period granularity that determines the frequency at which backup run snapshots will be copied.  NOTE: The granularity (in combination with the &#39;multiplier&#39; field below) that is specified should be such that the frequency of copying snapshots is lower than the frequency of actually creating the snapshots (i.e., lower than the frequency of the backup job runs). | [optional] 
**Multiplier** | Pointer to **NullableInt32** | A factor to multiply the granularity by. For example, if this is 2 and the granularity is kHour, then snapshots from the first eligible run from every 2 hour period will be copied. | [optional] 

## Methods

### NewGranularityBucket

`func NewGranularityBucket() *GranularityBucket`

NewGranularityBucket instantiates a new GranularityBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGranularityBucketWithDefaults

`func NewGranularityBucketWithDefaults() *GranularityBucket`

NewGranularityBucketWithDefaults instantiates a new GranularityBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGranularity

`func (o *GranularityBucket) GetGranularity() int32`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GranularityBucket) GetGranularityOk() (*int32, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GranularityBucket) SetGranularity(v int32)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *GranularityBucket) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### SetGranularityNil

`func (o *GranularityBucket) SetGranularityNil(b bool)`

 SetGranularityNil sets the value for Granularity to be an explicit nil

### UnsetGranularity
`func (o *GranularityBucket) UnsetGranularity()`

UnsetGranularity ensures that no value is present for Granularity, not even an explicit nil
### GetMultiplier

`func (o *GranularityBucket) GetMultiplier() int32`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *GranularityBucket) GetMultiplierOk() (*int32, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *GranularityBucket) SetMultiplier(v int32)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *GranularityBucket) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### SetMultiplierNil

`func (o *GranularityBucket) SetMultiplierNil(b bool)`

 SetMultiplierNil sets the value for Multiplier to be an explicit nil

### UnsetMultiplier
`func (o *GranularityBucket) UnsetMultiplier()`

UnsetMultiplier ensures that no value is present for Multiplier, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



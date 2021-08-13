# SourceAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastRefreshedTime** | Pointer to **NullableInt64** | Time at which the data about this protection source was last refreshed. | [optional] 
**RegistrationId** | Pointer to **NullableInt64** | Id of the registration as part of which this source was discovered. | [optional] 

## Methods

### NewSourceAllOf

`func NewSourceAllOf() *SourceAllOf`

NewSourceAllOf instantiates a new SourceAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceAllOfWithDefaults

`func NewSourceAllOfWithDefaults() *SourceAllOf`

NewSourceAllOfWithDefaults instantiates a new SourceAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastRefreshedTime

`func (o *SourceAllOf) GetLastRefreshedTime() int64`

GetLastRefreshedTime returns the LastRefreshedTime field if non-nil, zero value otherwise.

### GetLastRefreshedTimeOk

`func (o *SourceAllOf) GetLastRefreshedTimeOk() (*int64, bool)`

GetLastRefreshedTimeOk returns a tuple with the LastRefreshedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRefreshedTime

`func (o *SourceAllOf) SetLastRefreshedTime(v int64)`

SetLastRefreshedTime sets LastRefreshedTime field to given value.

### HasLastRefreshedTime

`func (o *SourceAllOf) HasLastRefreshedTime() bool`

HasLastRefreshedTime returns a boolean if a field has been set.

### SetLastRefreshedTimeNil

`func (o *SourceAllOf) SetLastRefreshedTimeNil(b bool)`

 SetLastRefreshedTimeNil sets the value for LastRefreshedTime to be an explicit nil

### UnsetLastRefreshedTime
`func (o *SourceAllOf) UnsetLastRefreshedTime()`

UnsetLastRefreshedTime ensures that no value is present for LastRefreshedTime, not even an explicit nil
### GetRegistrationId

`func (o *SourceAllOf) GetRegistrationId() int64`

GetRegistrationId returns the RegistrationId field if non-nil, zero value otherwise.

### GetRegistrationIdOk

`func (o *SourceAllOf) GetRegistrationIdOk() (*int64, bool)`

GetRegistrationIdOk returns a tuple with the RegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationId

`func (o *SourceAllOf) SetRegistrationId(v int64)`

SetRegistrationId sets RegistrationId field to given value.

### HasRegistrationId

`func (o *SourceAllOf) HasRegistrationId() bool`

HasRegistrationId returns a boolean if a field has been set.

### SetRegistrationIdNil

`func (o *SourceAllOf) SetRegistrationIdNil(b bool)`

 SetRegistrationIdNil sets the value for RegistrationId to be an explicit nil

### UnsetRegistrationId
`func (o *SourceAllOf) UnsetRegistrationId()`

UnsetRegistrationId ensures that no value is present for RegistrationId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



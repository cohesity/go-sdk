# RetryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retries** | Pointer to **NullableInt32** | Specifies the number of times to retry capturing Snapshots before the Protection Group Run fails. | [optional] 
**RetryIntervalMins** | Pointer to **NullableInt32** | Specifies the number of minutes before retrying a failed Protection Group. | [optional] 

## Methods

### NewRetryOptions

`func NewRetryOptions() *RetryOptions`

NewRetryOptions instantiates a new RetryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetryOptionsWithDefaults

`func NewRetryOptionsWithDefaults() *RetryOptions`

NewRetryOptionsWithDefaults instantiates a new RetryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetries

`func (o *RetryOptions) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *RetryOptions) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *RetryOptions) SetRetries(v int32)`

SetRetries sets Retries field to given value.

### HasRetries

`func (o *RetryOptions) HasRetries() bool`

HasRetries returns a boolean if a field has been set.

### SetRetriesNil

`func (o *RetryOptions) SetRetriesNil(b bool)`

 SetRetriesNil sets the value for Retries to be an explicit nil

### UnsetRetries
`func (o *RetryOptions) UnsetRetries()`

UnsetRetries ensures that no value is present for Retries, not even an explicit nil
### GetRetryIntervalMins

`func (o *RetryOptions) GetRetryIntervalMins() int32`

GetRetryIntervalMins returns the RetryIntervalMins field if non-nil, zero value otherwise.

### GetRetryIntervalMinsOk

`func (o *RetryOptions) GetRetryIntervalMinsOk() (*int32, bool)`

GetRetryIntervalMinsOk returns a tuple with the RetryIntervalMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryIntervalMins

`func (o *RetryOptions) SetRetryIntervalMins(v int32)`

SetRetryIntervalMins sets RetryIntervalMins field to given value.

### HasRetryIntervalMins

`func (o *RetryOptions) HasRetryIntervalMins() bool`

HasRetryIntervalMins returns a boolean if a field has been set.

### SetRetryIntervalMinsNil

`func (o *RetryOptions) SetRetryIntervalMinsNil(b bool)`

 SetRetryIntervalMinsNil sets the value for RetryIntervalMins to be an explicit nil

### UnsetRetryIntervalMins
`func (o *RetryOptions) UnsetRetryIntervalMins()`

UnsetRetryIntervalMins ensures that no value is present for RetryIntervalMins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



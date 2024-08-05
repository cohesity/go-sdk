# ComponentRemovalProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsRemovalStuck** | Pointer to **NullableBool** | Specifies if the removal is stuck. | [optional] 
**ProgressPercentage** | Pointer to **NullableInt32** | Specifies the overall progress percentage for the service. | [optional] 
**RemovalProgress** | Pointer to **NullableString** | Specifies removal progress details. | [optional] 
**ServiceName** | Pointer to **NullableString** | Specifies service name. | [optional] 
**TimeRemaining** | Pointer to **NullableInt64** | Specifies the total duration in seconds left to Ack the service. | [optional] 

## Methods

### NewComponentRemovalProgress

`func NewComponentRemovalProgress() *ComponentRemovalProgress`

NewComponentRemovalProgress instantiates a new ComponentRemovalProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentRemovalProgressWithDefaults

`func NewComponentRemovalProgressWithDefaults() *ComponentRemovalProgress`

NewComponentRemovalProgressWithDefaults instantiates a new ComponentRemovalProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsRemovalStuck

`func (o *ComponentRemovalProgress) GetIsRemovalStuck() bool`

GetIsRemovalStuck returns the IsRemovalStuck field if non-nil, zero value otherwise.

### GetIsRemovalStuckOk

`func (o *ComponentRemovalProgress) GetIsRemovalStuckOk() (*bool, bool)`

GetIsRemovalStuckOk returns a tuple with the IsRemovalStuck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemovalStuck

`func (o *ComponentRemovalProgress) SetIsRemovalStuck(v bool)`

SetIsRemovalStuck sets IsRemovalStuck field to given value.

### HasIsRemovalStuck

`func (o *ComponentRemovalProgress) HasIsRemovalStuck() bool`

HasIsRemovalStuck returns a boolean if a field has been set.

### SetIsRemovalStuckNil

`func (o *ComponentRemovalProgress) SetIsRemovalStuckNil(b bool)`

 SetIsRemovalStuckNil sets the value for IsRemovalStuck to be an explicit nil

### UnsetIsRemovalStuck
`func (o *ComponentRemovalProgress) UnsetIsRemovalStuck()`

UnsetIsRemovalStuck ensures that no value is present for IsRemovalStuck, not even an explicit nil
### GetProgressPercentage

`func (o *ComponentRemovalProgress) GetProgressPercentage() int32`

GetProgressPercentage returns the ProgressPercentage field if non-nil, zero value otherwise.

### GetProgressPercentageOk

`func (o *ComponentRemovalProgress) GetProgressPercentageOk() (*int32, bool)`

GetProgressPercentageOk returns a tuple with the ProgressPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercentage

`func (o *ComponentRemovalProgress) SetProgressPercentage(v int32)`

SetProgressPercentage sets ProgressPercentage field to given value.

### HasProgressPercentage

`func (o *ComponentRemovalProgress) HasProgressPercentage() bool`

HasProgressPercentage returns a boolean if a field has been set.

### SetProgressPercentageNil

`func (o *ComponentRemovalProgress) SetProgressPercentageNil(b bool)`

 SetProgressPercentageNil sets the value for ProgressPercentage to be an explicit nil

### UnsetProgressPercentage
`func (o *ComponentRemovalProgress) UnsetProgressPercentage()`

UnsetProgressPercentage ensures that no value is present for ProgressPercentage, not even an explicit nil
### GetRemovalProgress

`func (o *ComponentRemovalProgress) GetRemovalProgress() string`

GetRemovalProgress returns the RemovalProgress field if non-nil, zero value otherwise.

### GetRemovalProgressOk

`func (o *ComponentRemovalProgress) GetRemovalProgressOk() (*string, bool)`

GetRemovalProgressOk returns a tuple with the RemovalProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalProgress

`func (o *ComponentRemovalProgress) SetRemovalProgress(v string)`

SetRemovalProgress sets RemovalProgress field to given value.

### HasRemovalProgress

`func (o *ComponentRemovalProgress) HasRemovalProgress() bool`

HasRemovalProgress returns a boolean if a field has been set.

### SetRemovalProgressNil

`func (o *ComponentRemovalProgress) SetRemovalProgressNil(b bool)`

 SetRemovalProgressNil sets the value for RemovalProgress to be an explicit nil

### UnsetRemovalProgress
`func (o *ComponentRemovalProgress) UnsetRemovalProgress()`

UnsetRemovalProgress ensures that no value is present for RemovalProgress, not even an explicit nil
### GetServiceName

`func (o *ComponentRemovalProgress) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ComponentRemovalProgress) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ComponentRemovalProgress) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *ComponentRemovalProgress) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### SetServiceNameNil

`func (o *ComponentRemovalProgress) SetServiceNameNil(b bool)`

 SetServiceNameNil sets the value for ServiceName to be an explicit nil

### UnsetServiceName
`func (o *ComponentRemovalProgress) UnsetServiceName()`

UnsetServiceName ensures that no value is present for ServiceName, not even an explicit nil
### GetTimeRemaining

`func (o *ComponentRemovalProgress) GetTimeRemaining() int64`

GetTimeRemaining returns the TimeRemaining field if non-nil, zero value otherwise.

### GetTimeRemainingOk

`func (o *ComponentRemovalProgress) GetTimeRemainingOk() (*int64, bool)`

GetTimeRemainingOk returns a tuple with the TimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemaining

`func (o *ComponentRemovalProgress) SetTimeRemaining(v int64)`

SetTimeRemaining sets TimeRemaining field to given value.

### HasTimeRemaining

`func (o *ComponentRemovalProgress) HasTimeRemaining() bool`

HasTimeRemaining returns a boolean if a field has been set.

### SetTimeRemainingNil

`func (o *ComponentRemovalProgress) SetTimeRemainingNil(b bool)`

 SetTimeRemainingNil sets the value for TimeRemaining to be an explicit nil

### UnsetTimeRemaining
`func (o *ComponentRemovalProgress) UnsetTimeRemaining()`

UnsetTimeRemaining ensures that no value is present for TimeRemaining, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



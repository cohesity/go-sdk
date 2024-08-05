# ServiceUnitProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InProgress** | Pointer to **NullableBool** | Specifies whether a operation is in progress for the service. | [optional] 
**NodesProgress** | Pointer to [**[]NodeUnitProgress**](NodeUnitProgress.md) | Specifies the details of patch operation for each service at each patch level. | [optional] 
**Percentage** | Pointer to **NullableInt64** | Specifies the percentage of completion of the patch unit operation. | [optional] 
**Service** | Pointer to **NullableString** | Specifies the service which is patched. | [optional] 
**ServiceMessage** | Pointer to **NullableString** | Specifies a message about the patch unit operation. | [optional] 
**TimeRemainingSeconds** | Pointer to **NullableInt64** | Specifies the time remaining to complete the patch operation for the service. | [optional] 
**TimeTakenSeconds** | Pointer to **NullableInt64** | Specifies the time taken so far in this patch unit operation for the service. | [optional] 

## Methods

### NewServiceUnitProgress

`func NewServiceUnitProgress() *ServiceUnitProgress`

NewServiceUnitProgress instantiates a new ServiceUnitProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceUnitProgressWithDefaults

`func NewServiceUnitProgressWithDefaults() *ServiceUnitProgress`

NewServiceUnitProgressWithDefaults instantiates a new ServiceUnitProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInProgress

`func (o *ServiceUnitProgress) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *ServiceUnitProgress) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *ServiceUnitProgress) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *ServiceUnitProgress) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### SetInProgressNil

`func (o *ServiceUnitProgress) SetInProgressNil(b bool)`

 SetInProgressNil sets the value for InProgress to be an explicit nil

### UnsetInProgress
`func (o *ServiceUnitProgress) UnsetInProgress()`

UnsetInProgress ensures that no value is present for InProgress, not even an explicit nil
### GetNodesProgress

`func (o *ServiceUnitProgress) GetNodesProgress() []NodeUnitProgress`

GetNodesProgress returns the NodesProgress field if non-nil, zero value otherwise.

### GetNodesProgressOk

`func (o *ServiceUnitProgress) GetNodesProgressOk() (*[]NodeUnitProgress, bool)`

GetNodesProgressOk returns a tuple with the NodesProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodesProgress

`func (o *ServiceUnitProgress) SetNodesProgress(v []NodeUnitProgress)`

SetNodesProgress sets NodesProgress field to given value.

### HasNodesProgress

`func (o *ServiceUnitProgress) HasNodesProgress() bool`

HasNodesProgress returns a boolean if a field has been set.

### SetNodesProgressNil

`func (o *ServiceUnitProgress) SetNodesProgressNil(b bool)`

 SetNodesProgressNil sets the value for NodesProgress to be an explicit nil

### UnsetNodesProgress
`func (o *ServiceUnitProgress) UnsetNodesProgress()`

UnsetNodesProgress ensures that no value is present for NodesProgress, not even an explicit nil
### GetPercentage

`func (o *ServiceUnitProgress) GetPercentage() int64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ServiceUnitProgress) GetPercentageOk() (*int64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ServiceUnitProgress) SetPercentage(v int64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ServiceUnitProgress) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *ServiceUnitProgress) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *ServiceUnitProgress) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetService

`func (o *ServiceUnitProgress) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ServiceUnitProgress) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ServiceUnitProgress) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ServiceUnitProgress) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ServiceUnitProgress) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ServiceUnitProgress) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetServiceMessage

`func (o *ServiceUnitProgress) GetServiceMessage() string`

GetServiceMessage returns the ServiceMessage field if non-nil, zero value otherwise.

### GetServiceMessageOk

`func (o *ServiceUnitProgress) GetServiceMessageOk() (*string, bool)`

GetServiceMessageOk returns a tuple with the ServiceMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMessage

`func (o *ServiceUnitProgress) SetServiceMessage(v string)`

SetServiceMessage sets ServiceMessage field to given value.

### HasServiceMessage

`func (o *ServiceUnitProgress) HasServiceMessage() bool`

HasServiceMessage returns a boolean if a field has been set.

### SetServiceMessageNil

`func (o *ServiceUnitProgress) SetServiceMessageNil(b bool)`

 SetServiceMessageNil sets the value for ServiceMessage to be an explicit nil

### UnsetServiceMessage
`func (o *ServiceUnitProgress) UnsetServiceMessage()`

UnsetServiceMessage ensures that no value is present for ServiceMessage, not even an explicit nil
### GetTimeRemainingSeconds

`func (o *ServiceUnitProgress) GetTimeRemainingSeconds() int64`

GetTimeRemainingSeconds returns the TimeRemainingSeconds field if non-nil, zero value otherwise.

### GetTimeRemainingSecondsOk

`func (o *ServiceUnitProgress) GetTimeRemainingSecondsOk() (*int64, bool)`

GetTimeRemainingSecondsOk returns a tuple with the TimeRemainingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemainingSeconds

`func (o *ServiceUnitProgress) SetTimeRemainingSeconds(v int64)`

SetTimeRemainingSeconds sets TimeRemainingSeconds field to given value.

### HasTimeRemainingSeconds

`func (o *ServiceUnitProgress) HasTimeRemainingSeconds() bool`

HasTimeRemainingSeconds returns a boolean if a field has been set.

### SetTimeRemainingSecondsNil

`func (o *ServiceUnitProgress) SetTimeRemainingSecondsNil(b bool)`

 SetTimeRemainingSecondsNil sets the value for TimeRemainingSeconds to be an explicit nil

### UnsetTimeRemainingSeconds
`func (o *ServiceUnitProgress) UnsetTimeRemainingSeconds()`

UnsetTimeRemainingSeconds ensures that no value is present for TimeRemainingSeconds, not even an explicit nil
### GetTimeTakenSeconds

`func (o *ServiceUnitProgress) GetTimeTakenSeconds() int64`

GetTimeTakenSeconds returns the TimeTakenSeconds field if non-nil, zero value otherwise.

### GetTimeTakenSecondsOk

`func (o *ServiceUnitProgress) GetTimeTakenSecondsOk() (*int64, bool)`

GetTimeTakenSecondsOk returns a tuple with the TimeTakenSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenSeconds

`func (o *ServiceUnitProgress) SetTimeTakenSeconds(v int64)`

SetTimeTakenSeconds sets TimeTakenSeconds field to given value.

### HasTimeTakenSeconds

`func (o *ServiceUnitProgress) HasTimeTakenSeconds() bool`

HasTimeTakenSeconds returns a boolean if a field has been set.

### SetTimeTakenSecondsNil

`func (o *ServiceUnitProgress) SetTimeTakenSecondsNil(b bool)`

 SetTimeTakenSecondsNil sets the value for TimeTakenSeconds to be an explicit nil

### UnsetTimeTakenSeconds
`func (o *ServiceUnitProgress) UnsetTimeTakenSeconds()`

UnsetTimeTakenSeconds ensures that no value is present for TimeTakenSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



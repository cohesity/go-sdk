# KubernetesHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | **[]string** | Specifies the commands. | 
**Container** | Pointer to **NullableString** | Specifies the name of the container. | [optional] 
**FailOnError** | Pointer to **NullableBool** | Specifies whether to fail on error or not. | [optional] 
**Timeout** | Pointer to **NullableInt64** | Specifies timeout for the operation. | [optional] 

## Methods

### NewKubernetesHook

`func NewKubernetesHook(commands []string, ) *KubernetesHook`

NewKubernetesHook instantiates a new KubernetesHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesHookWithDefaults

`func NewKubernetesHookWithDefaults() *KubernetesHook`

NewKubernetesHookWithDefaults instantiates a new KubernetesHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *KubernetesHook) GetCommands() []string`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *KubernetesHook) GetCommandsOk() (*[]string, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *KubernetesHook) SetCommands(v []string)`

SetCommands sets Commands field to given value.


### GetContainer

`func (o *KubernetesHook) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *KubernetesHook) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *KubernetesHook) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *KubernetesHook) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *KubernetesHook) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *KubernetesHook) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetFailOnError

`func (o *KubernetesHook) GetFailOnError() bool`

GetFailOnError returns the FailOnError field if non-nil, zero value otherwise.

### GetFailOnErrorOk

`func (o *KubernetesHook) GetFailOnErrorOk() (*bool, bool)`

GetFailOnErrorOk returns a tuple with the FailOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnError

`func (o *KubernetesHook) SetFailOnError(v bool)`

SetFailOnError sets FailOnError field to given value.

### HasFailOnError

`func (o *KubernetesHook) HasFailOnError() bool`

HasFailOnError returns a boolean if a field has been set.

### SetFailOnErrorNil

`func (o *KubernetesHook) SetFailOnErrorNil(b bool)`

 SetFailOnErrorNil sets the value for FailOnError to be an explicit nil

### UnsetFailOnError
`func (o *KubernetesHook) UnsetFailOnError()`

UnsetFailOnError ensures that no value is present for FailOnError, not even an explicit nil
### GetTimeout

`func (o *KubernetesHook) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *KubernetesHook) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *KubernetesHook) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *KubernetesHook) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### SetTimeoutNil

`func (o *KubernetesHook) SetTimeoutNil(b bool)`

 SetTimeoutNil sets the value for Timeout to be an explicit nil

### UnsetTimeout
`func (o *KubernetesHook) UnsetTimeout()`

UnsetTimeout ensures that no value is present for Timeout, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



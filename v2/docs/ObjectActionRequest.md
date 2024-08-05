# ObjectActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Object. | 
**VmwareParams** | Pointer to [**VmwareObjectActionParams**](VmwareObjectActionParams.md) |  | [optional] 

## Methods

### NewObjectActionRequest

`func NewObjectActionRequest(environment NullableString, ) *ObjectActionRequest`

NewObjectActionRequest instantiates a new ObjectActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewObjectActionRequestWithDefaults

`func NewObjectActionRequestWithDefaults() *ObjectActionRequest`

NewObjectActionRequestWithDefaults instantiates a new ObjectActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ObjectActionRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ObjectActionRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ObjectActionRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *ObjectActionRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *ObjectActionRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetVmwareParams

`func (o *ObjectActionRequest) GetVmwareParams() VmwareObjectActionParams`

GetVmwareParams returns the VmwareParams field if non-nil, zero value otherwise.

### GetVmwareParamsOk

`func (o *ObjectActionRequest) GetVmwareParamsOk() (*VmwareObjectActionParams, bool)`

GetVmwareParamsOk returns a tuple with the VmwareParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareParams

`func (o *ObjectActionRequest) SetVmwareParams(v VmwareObjectActionParams)`

SetVmwareParams sets VmwareParams field to given value.

### HasVmwareParams

`func (o *ObjectActionRequest) HasVmwareParams() bool`

HasVmwareParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



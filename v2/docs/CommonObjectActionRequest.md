# CommonObjectActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | **NullableString** | Specifies the environment type of the Object. | 

## Methods

### NewCommonObjectActionRequest

`func NewCommonObjectActionRequest(environment NullableString, ) *CommonObjectActionRequest`

NewCommonObjectActionRequest instantiates a new CommonObjectActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonObjectActionRequestWithDefaults

`func NewCommonObjectActionRequestWithDefaults() *CommonObjectActionRequest`

NewCommonObjectActionRequestWithDefaults instantiates a new CommonObjectActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *CommonObjectActionRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CommonObjectActionRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CommonObjectActionRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### SetEnvironmentNil

`func (o *CommonObjectActionRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CommonObjectActionRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ScvmmRegistrationParamsAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentEndpoint** | Pointer to **NullableString** | Specifies the agent endpoint if it is different from the source endpoint. | [optional] 
**ThrottlingParams** | Pointer to [**ThrottlingParams**](ThrottlingParams.md) |  | [optional] 

## Methods

### NewScvmmRegistrationParamsAllOf

`func NewScvmmRegistrationParamsAllOf() *ScvmmRegistrationParamsAllOf`

NewScvmmRegistrationParamsAllOf instantiates a new ScvmmRegistrationParamsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScvmmRegistrationParamsAllOfWithDefaults

`func NewScvmmRegistrationParamsAllOfWithDefaults() *ScvmmRegistrationParamsAllOf`

NewScvmmRegistrationParamsAllOfWithDefaults instantiates a new ScvmmRegistrationParamsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentEndpoint

`func (o *ScvmmRegistrationParamsAllOf) GetAgentEndpoint() string`

GetAgentEndpoint returns the AgentEndpoint field if non-nil, zero value otherwise.

### GetAgentEndpointOk

`func (o *ScvmmRegistrationParamsAllOf) GetAgentEndpointOk() (*string, bool)`

GetAgentEndpointOk returns a tuple with the AgentEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEndpoint

`func (o *ScvmmRegistrationParamsAllOf) SetAgentEndpoint(v string)`

SetAgentEndpoint sets AgentEndpoint field to given value.

### HasAgentEndpoint

`func (o *ScvmmRegistrationParamsAllOf) HasAgentEndpoint() bool`

HasAgentEndpoint returns a boolean if a field has been set.

### SetAgentEndpointNil

`func (o *ScvmmRegistrationParamsAllOf) SetAgentEndpointNil(b bool)`

 SetAgentEndpointNil sets the value for AgentEndpoint to be an explicit nil

### UnsetAgentEndpoint
`func (o *ScvmmRegistrationParamsAllOf) UnsetAgentEndpoint()`

UnsetAgentEndpoint ensures that no value is present for AgentEndpoint, not even an explicit nil
### GetThrottlingParams

`func (o *ScvmmRegistrationParamsAllOf) GetThrottlingParams() ThrottlingParams`

GetThrottlingParams returns the ThrottlingParams field if non-nil, zero value otherwise.

### GetThrottlingParamsOk

`func (o *ScvmmRegistrationParamsAllOf) GetThrottlingParamsOk() (*ThrottlingParams, bool)`

GetThrottlingParamsOk returns a tuple with the ThrottlingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingParams

`func (o *ScvmmRegistrationParamsAllOf) SetThrottlingParams(v ThrottlingParams)`

SetThrottlingParams sets ThrottlingParams field to given value.

### HasThrottlingParams

`func (o *ScvmmRegistrationParamsAllOf) HasThrottlingParams() bool`

HasThrottlingParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



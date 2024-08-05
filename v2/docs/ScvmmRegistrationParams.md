# ScvmmRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Specifies the password to access target entity. | 
**Username** | **string** | Specifies the username to access target entity. | 
**Description** | Pointer to **NullableString** | Specifies the description of the source being registered. | [optional] 
**Endpoint** | **string** | Specifies the endpoint IPaddress, URL or hostname of the host. | 
**AgentEndpoint** | Pointer to **NullableString** | Specifies the agent endpoint if it is different from the source endpoint. | [optional] 
**ThrottlingParams** | Pointer to [**ThrottlingParams**](ThrottlingParams.md) |  | [optional] 

## Methods

### NewScvmmRegistrationParams

`func NewScvmmRegistrationParams(password string, username string, endpoint string, ) *ScvmmRegistrationParams`

NewScvmmRegistrationParams instantiates a new ScvmmRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScvmmRegistrationParamsWithDefaults

`func NewScvmmRegistrationParamsWithDefaults() *ScvmmRegistrationParams`

NewScvmmRegistrationParamsWithDefaults instantiates a new ScvmmRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ScvmmRegistrationParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ScvmmRegistrationParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ScvmmRegistrationParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ScvmmRegistrationParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ScvmmRegistrationParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ScvmmRegistrationParams) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDescription

`func (o *ScvmmRegistrationParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScvmmRegistrationParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScvmmRegistrationParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScvmmRegistrationParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ScvmmRegistrationParams) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ScvmmRegistrationParams) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEndpoint

`func (o *ScvmmRegistrationParams) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ScvmmRegistrationParams) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ScvmmRegistrationParams) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAgentEndpoint

`func (o *ScvmmRegistrationParams) GetAgentEndpoint() string`

GetAgentEndpoint returns the AgentEndpoint field if non-nil, zero value otherwise.

### GetAgentEndpointOk

`func (o *ScvmmRegistrationParams) GetAgentEndpointOk() (*string, bool)`

GetAgentEndpointOk returns a tuple with the AgentEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentEndpoint

`func (o *ScvmmRegistrationParams) SetAgentEndpoint(v string)`

SetAgentEndpoint sets AgentEndpoint field to given value.

### HasAgentEndpoint

`func (o *ScvmmRegistrationParams) HasAgentEndpoint() bool`

HasAgentEndpoint returns a boolean if a field has been set.

### SetAgentEndpointNil

`func (o *ScvmmRegistrationParams) SetAgentEndpointNil(b bool)`

 SetAgentEndpointNil sets the value for AgentEndpoint to be an explicit nil

### UnsetAgentEndpoint
`func (o *ScvmmRegistrationParams) UnsetAgentEndpoint()`

UnsetAgentEndpoint ensures that no value is present for AgentEndpoint, not even an explicit nil
### GetThrottlingParams

`func (o *ScvmmRegistrationParams) GetThrottlingParams() ThrottlingParams`

GetThrottlingParams returns the ThrottlingParams field if non-nil, zero value otherwise.

### GetThrottlingParamsOk

`func (o *ScvmmRegistrationParams) GetThrottlingParamsOk() (*ThrottlingParams, bool)`

GetThrottlingParamsOk returns a tuple with the ThrottlingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThrottlingParams

`func (o *ScvmmRegistrationParams) SetThrottlingParams(v ThrottlingParams)`

SetThrottlingParams sets ThrottlingParams field to given value.

### HasThrottlingParams

`func (o *ScvmmRegistrationParams) HasThrottlingParams() bool`

HasThrottlingParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# McmAgentImagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agents** | Pointer to [**[]McmAgentImage**](McmAgentImage.md) | Specifies a list of agents from Helios. | [optional] 

## Methods

### NewMcmAgentImagesResponse

`func NewMcmAgentImagesResponse() *McmAgentImagesResponse`

NewMcmAgentImagesResponse instantiates a new McmAgentImagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmAgentImagesResponseWithDefaults

`func NewMcmAgentImagesResponseWithDefaults() *McmAgentImagesResponse`

NewMcmAgentImagesResponseWithDefaults instantiates a new McmAgentImagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgents

`func (o *McmAgentImagesResponse) GetAgents() []McmAgentImage`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *McmAgentImagesResponse) GetAgentsOk() (*[]McmAgentImage, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *McmAgentImagesResponse) SetAgents(v []McmAgentImage)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *McmAgentImagesResponse) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### SetAgentsNil

`func (o *McmAgentImagesResponse) SetAgentsNil(b bool)`

 SetAgentsNil sets the value for Agents to be an explicit nil

### UnsetAgents
`func (o *McmAgentImagesResponse) UnsetAgents()`

UnsetAgents ensures that no value is present for Agents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



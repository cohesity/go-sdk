# McmAgentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentID** | Pointer to **NullableInt64** | Specifies the ids of the agents related to the source. | [optional] 
**AgentDisplayName** | Pointer to **NullableString** | Specifies the hostname/ip address of the source. | [optional] 

## Methods

### NewMcmAgentInfo

`func NewMcmAgentInfo() *McmAgentInfo`

NewMcmAgentInfo instantiates a new McmAgentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcmAgentInfoWithDefaults

`func NewMcmAgentInfoWithDefaults() *McmAgentInfo`

NewMcmAgentInfoWithDefaults instantiates a new McmAgentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentID

`func (o *McmAgentInfo) GetAgentID() int64`

GetAgentID returns the AgentID field if non-nil, zero value otherwise.

### GetAgentIDOk

`func (o *McmAgentInfo) GetAgentIDOk() (*int64, bool)`

GetAgentIDOk returns a tuple with the AgentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentID

`func (o *McmAgentInfo) SetAgentID(v int64)`

SetAgentID sets AgentID field to given value.

### HasAgentID

`func (o *McmAgentInfo) HasAgentID() bool`

HasAgentID returns a boolean if a field has been set.

### SetAgentIDNil

`func (o *McmAgentInfo) SetAgentIDNil(b bool)`

 SetAgentIDNil sets the value for AgentID to be an explicit nil

### UnsetAgentID
`func (o *McmAgentInfo) UnsetAgentID()`

UnsetAgentID ensures that no value is present for AgentID, not even an explicit nil
### GetAgentDisplayName

`func (o *McmAgentInfo) GetAgentDisplayName() string`

GetAgentDisplayName returns the AgentDisplayName field if non-nil, zero value otherwise.

### GetAgentDisplayNameOk

`func (o *McmAgentInfo) GetAgentDisplayNameOk() (*string, bool)`

GetAgentDisplayNameOk returns a tuple with the AgentDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentDisplayName

`func (o *McmAgentInfo) SetAgentDisplayName(v string)`

SetAgentDisplayName sets AgentDisplayName field to given value.

### HasAgentDisplayName

`func (o *McmAgentInfo) HasAgentDisplayName() bool`

HasAgentDisplayName returns a boolean if a field has been set.

### SetAgentDisplayNameNil

`func (o *McmAgentInfo) SetAgentDisplayNameNil(b bool)`

 SetAgentDisplayNameNil sets the value for AgentDisplayName to be an explicit nil

### UnsetAgentDisplayName
`func (o *McmAgentInfo) UnsetAgentDisplayName()`

UnsetAgentDisplayName ensures that no value is present for AgentDisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



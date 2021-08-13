# AgentUpgradeTaskActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **NullableString** | Specifies the action type.&lt;br&gt; &#39;Cancel&#39; indicates that the upgrade task needs to be cancelled. | [optional] 
**Id** | Pointer to **NullableInt64** | Specifies the ID of the task. | [optional] 

## Methods

### NewAgentUpgradeTaskActionRequest

`func NewAgentUpgradeTaskActionRequest() *AgentUpgradeTaskActionRequest`

NewAgentUpgradeTaskActionRequest instantiates a new AgentUpgradeTaskActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpgradeTaskActionRequestWithDefaults

`func NewAgentUpgradeTaskActionRequestWithDefaults() *AgentUpgradeTaskActionRequest`

NewAgentUpgradeTaskActionRequestWithDefaults instantiates a new AgentUpgradeTaskActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *AgentUpgradeTaskActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *AgentUpgradeTaskActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *AgentUpgradeTaskActionRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *AgentUpgradeTaskActionRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetActionNil

`func (o *AgentUpgradeTaskActionRequest) SetActionNil(b bool)`

 SetActionNil sets the value for Action to be an explicit nil

### UnsetAction
`func (o *AgentUpgradeTaskActionRequest) UnsetAction()`

UnsetAction ensures that no value is present for Action, not even an explicit nil
### GetId

`func (o *AgentUpgradeTaskActionRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentUpgradeTaskActionRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentUpgradeTaskActionRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AgentUpgradeTaskActionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgentUpgradeTaskActionRequest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgentUpgradeTaskActionRequest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



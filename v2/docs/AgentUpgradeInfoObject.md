# AgentUpgradeInfoObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** | Specifies the ID of the agent. | [optional] 
**Info** | Pointer to [**AgentInfoObject**](AgentInfoObject.md) |  | [optional] 

## Methods

### NewAgentUpgradeInfoObject

`func NewAgentUpgradeInfoObject() *AgentUpgradeInfoObject`

NewAgentUpgradeInfoObject instantiates a new AgentUpgradeInfoObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentUpgradeInfoObjectWithDefaults

`func NewAgentUpgradeInfoObjectWithDefaults() *AgentUpgradeInfoObject`

NewAgentUpgradeInfoObjectWithDefaults instantiates a new AgentUpgradeInfoObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentUpgradeInfoObject) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentUpgradeInfoObject) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentUpgradeInfoObject) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AgentUpgradeInfoObject) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *AgentUpgradeInfoObject) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *AgentUpgradeInfoObject) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInfo

`func (o *AgentUpgradeInfoObject) GetInfo() AgentInfoObject`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgentUpgradeInfoObject) GetInfoOk() (*AgentInfoObject, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgentUpgradeInfoObject) SetInfo(v AgentInfoObject)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgentUpgradeInfoObject) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



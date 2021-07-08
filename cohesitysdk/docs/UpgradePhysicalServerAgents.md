# UpgradePhysicalServerAgents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentIds** | **[]int64** | Array of Agent Ids.  Specifies a list of agentIds associated with the Physical Servers to upgrade with the agent release currently available from the Cohesity Cluster. | 

## Methods

### NewUpgradePhysicalServerAgents

`func NewUpgradePhysicalServerAgents(agentIds []int64, ) *UpgradePhysicalServerAgents`

NewUpgradePhysicalServerAgents instantiates a new UpgradePhysicalServerAgents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradePhysicalServerAgentsWithDefaults

`func NewUpgradePhysicalServerAgentsWithDefaults() *UpgradePhysicalServerAgents`

NewUpgradePhysicalServerAgentsWithDefaults instantiates a new UpgradePhysicalServerAgents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentIds

`func (o *UpgradePhysicalServerAgents) GetAgentIds() []int64`

GetAgentIds returns the AgentIds field if non-nil, zero value otherwise.

### GetAgentIdsOk

`func (o *UpgradePhysicalServerAgents) GetAgentIdsOk() (*[]int64, bool)`

GetAgentIdsOk returns a tuple with the AgentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentIds

`func (o *UpgradePhysicalServerAgents) SetAgentIds(v []int64)`

SetAgentIds sets AgentIds field to given value.


### SetAgentIdsNil

`func (o *UpgradePhysicalServerAgents) SetAgentIdsNil(b bool)`

 SetAgentIdsNil sets the value for AgentIds to be an explicit nil

### UnsetAgentIds
`func (o *UpgradePhysicalServerAgents) UnsetAgentIds()`

UnsetAgentIds ensures that no value is present for AgentIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# NodeUnitProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIp** | Pointer to **NullableString** | Specifies the IP address of the node. | [optional] 
**InProgress** | Pointer to **NullableBool** | Specifies whether a operation is in progress on the node. | [optional] 
**PatchLevelTransition** | Pointer to **NullableString** | Specifies the patch level transition of the patch operation. For Apply operation, patch level goes up for each operation. For Revert operation, patch level goes down. Patch level zero is the base level where no patch was applied. | [optional] 
**Percentage** | Pointer to **NullableInt64** | Specifies the percentage of completion of the patch operation on the node. | [optional] 
**TimeTakenSeconds** | Pointer to **NullableInt64** | Specifies the time taken so far in this patch unit operation on the node. | [optional] 
**NodeMessage** | Pointer to **NullableString** | Specifies a message about the patch operation on the node. | [optional] 

## Methods

### NewNodeUnitProgress

`func NewNodeUnitProgress() *NodeUnitProgress`

NewNodeUnitProgress instantiates a new NodeUnitProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeUnitProgressWithDefaults

`func NewNodeUnitProgressWithDefaults() *NodeUnitProgress`

NewNodeUnitProgressWithDefaults instantiates a new NodeUnitProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIp

`func (o *NodeUnitProgress) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *NodeUnitProgress) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *NodeUnitProgress) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *NodeUnitProgress) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *NodeUnitProgress) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *NodeUnitProgress) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil
### GetInProgress

`func (o *NodeUnitProgress) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *NodeUnitProgress) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *NodeUnitProgress) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *NodeUnitProgress) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### SetInProgressNil

`func (o *NodeUnitProgress) SetInProgressNil(b bool)`

 SetInProgressNil sets the value for InProgress to be an explicit nil

### UnsetInProgress
`func (o *NodeUnitProgress) UnsetInProgress()`

UnsetInProgress ensures that no value is present for InProgress, not even an explicit nil
### GetPatchLevelTransition

`func (o *NodeUnitProgress) GetPatchLevelTransition() string`

GetPatchLevelTransition returns the PatchLevelTransition field if non-nil, zero value otherwise.

### GetPatchLevelTransitionOk

`func (o *NodeUnitProgress) GetPatchLevelTransitionOk() (*string, bool)`

GetPatchLevelTransitionOk returns a tuple with the PatchLevelTransition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchLevelTransition

`func (o *NodeUnitProgress) SetPatchLevelTransition(v string)`

SetPatchLevelTransition sets PatchLevelTransition field to given value.

### HasPatchLevelTransition

`func (o *NodeUnitProgress) HasPatchLevelTransition() bool`

HasPatchLevelTransition returns a boolean if a field has been set.

### SetPatchLevelTransitionNil

`func (o *NodeUnitProgress) SetPatchLevelTransitionNil(b bool)`

 SetPatchLevelTransitionNil sets the value for PatchLevelTransition to be an explicit nil

### UnsetPatchLevelTransition
`func (o *NodeUnitProgress) UnsetPatchLevelTransition()`

UnsetPatchLevelTransition ensures that no value is present for PatchLevelTransition, not even an explicit nil
### GetPercentage

`func (o *NodeUnitProgress) GetPercentage() int64`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *NodeUnitProgress) GetPercentageOk() (*int64, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *NodeUnitProgress) SetPercentage(v int64)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *NodeUnitProgress) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### SetPercentageNil

`func (o *NodeUnitProgress) SetPercentageNil(b bool)`

 SetPercentageNil sets the value for Percentage to be an explicit nil

### UnsetPercentage
`func (o *NodeUnitProgress) UnsetPercentage()`

UnsetPercentage ensures that no value is present for Percentage, not even an explicit nil
### GetTimeTakenSeconds

`func (o *NodeUnitProgress) GetTimeTakenSeconds() int64`

GetTimeTakenSeconds returns the TimeTakenSeconds field if non-nil, zero value otherwise.

### GetTimeTakenSecondsOk

`func (o *NodeUnitProgress) GetTimeTakenSecondsOk() (*int64, bool)`

GetTimeTakenSecondsOk returns a tuple with the TimeTakenSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenSeconds

`func (o *NodeUnitProgress) SetTimeTakenSeconds(v int64)`

SetTimeTakenSeconds sets TimeTakenSeconds field to given value.

### HasTimeTakenSeconds

`func (o *NodeUnitProgress) HasTimeTakenSeconds() bool`

HasTimeTakenSeconds returns a boolean if a field has been set.

### SetTimeTakenSecondsNil

`func (o *NodeUnitProgress) SetTimeTakenSecondsNil(b bool)`

 SetTimeTakenSecondsNil sets the value for TimeTakenSeconds to be an explicit nil

### UnsetTimeTakenSeconds
`func (o *NodeUnitProgress) UnsetTimeTakenSeconds()`

UnsetTimeTakenSeconds ensures that no value is present for TimeTakenSeconds, not even an explicit nil
### GetNodeMessage

`func (o *NodeUnitProgress) GetNodeMessage() string`

GetNodeMessage returns the NodeMessage field if non-nil, zero value otherwise.

### GetNodeMessageOk

`func (o *NodeUnitProgress) GetNodeMessageOk() (*string, bool)`

GetNodeMessageOk returns a tuple with the NodeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeMessage

`func (o *NodeUnitProgress) SetNodeMessage(v string)`

SetNodeMessage sets NodeMessage field to given value.

### HasNodeMessage

`func (o *NodeUnitProgress) HasNodeMessage() bool`

HasNodeMessage returns a boolean if a field has been set.

### SetNodeMessageNil

`func (o *NodeUnitProgress) SetNodeMessageNil(b bool)`

 SetNodeMessageNil sets the value for NodeMessage to be an explicit nil

### UnsetNodeMessage
`func (o *NodeUnitProgress) UnsetNodeMessage()`

UnsetNodeMessage ensures that no value is present for NodeMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



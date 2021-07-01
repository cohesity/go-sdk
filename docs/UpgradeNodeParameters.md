# UpgradeNodeParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeIds** | Pointer to **[]int64** | Specifies a list of IDs of additional nodes to be upgraded. These must be free Nodes present on the same local network as the Node that the request was sent to. The ID of the Node the request was sent to should not be included in this list. This parameter can only be specified if upgradeAllFreeNodes is not. | [optional] 
**TargetSwVersion** | Pointer to **NullableString** | Specifies the target software version. The node that the request is sent to will search itself for the specified software package and if that package is found, it will be used for the upgrade. | [optional] 
**UpgradeAllFreeNodes** | Pointer to **NullableBool** | Specifies whether or not to attempt to upgrade all free nodes which are currently connected to the same local network as the node that the request was sent to. This parameter can only be specified if nodeIds is not. | [optional] 
**UpgradeSelf** | Pointer to **NullableBool** | Specifies that the node that the request is being sent to should be upgraded. By default, this is set to true. | [optional] 

## Methods

### NewUpgradeNodeParameters

`func NewUpgradeNodeParameters() *UpgradeNodeParameters`

NewUpgradeNodeParameters instantiates a new UpgradeNodeParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeNodeParametersWithDefaults

`func NewUpgradeNodeParametersWithDefaults() *UpgradeNodeParameters`

NewUpgradeNodeParametersWithDefaults instantiates a new UpgradeNodeParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeIds

`func (o *UpgradeNodeParameters) GetNodeIds() []int64`

GetNodeIds returns the NodeIds field if non-nil, zero value otherwise.

### GetNodeIdsOk

`func (o *UpgradeNodeParameters) GetNodeIdsOk() (*[]int64, bool)`

GetNodeIdsOk returns a tuple with the NodeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIds

`func (o *UpgradeNodeParameters) SetNodeIds(v []int64)`

SetNodeIds sets NodeIds field to given value.

### HasNodeIds

`func (o *UpgradeNodeParameters) HasNodeIds() bool`

HasNodeIds returns a boolean if a field has been set.

### SetNodeIdsNil

`func (o *UpgradeNodeParameters) SetNodeIdsNil(b bool)`

 SetNodeIdsNil sets the value for NodeIds to be an explicit nil

### UnsetNodeIds
`func (o *UpgradeNodeParameters) UnsetNodeIds()`

UnsetNodeIds ensures that no value is present for NodeIds, not even an explicit nil
### GetTargetSwVersion

`func (o *UpgradeNodeParameters) GetTargetSwVersion() string`

GetTargetSwVersion returns the TargetSwVersion field if non-nil, zero value otherwise.

### GetTargetSwVersionOk

`func (o *UpgradeNodeParameters) GetTargetSwVersionOk() (*string, bool)`

GetTargetSwVersionOk returns a tuple with the TargetSwVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSwVersion

`func (o *UpgradeNodeParameters) SetTargetSwVersion(v string)`

SetTargetSwVersion sets TargetSwVersion field to given value.

### HasTargetSwVersion

`func (o *UpgradeNodeParameters) HasTargetSwVersion() bool`

HasTargetSwVersion returns a boolean if a field has been set.

### SetTargetSwVersionNil

`func (o *UpgradeNodeParameters) SetTargetSwVersionNil(b bool)`

 SetTargetSwVersionNil sets the value for TargetSwVersion to be an explicit nil

### UnsetTargetSwVersion
`func (o *UpgradeNodeParameters) UnsetTargetSwVersion()`

UnsetTargetSwVersion ensures that no value is present for TargetSwVersion, not even an explicit nil
### GetUpgradeAllFreeNodes

`func (o *UpgradeNodeParameters) GetUpgradeAllFreeNodes() bool`

GetUpgradeAllFreeNodes returns the UpgradeAllFreeNodes field if non-nil, zero value otherwise.

### GetUpgradeAllFreeNodesOk

`func (o *UpgradeNodeParameters) GetUpgradeAllFreeNodesOk() (*bool, bool)`

GetUpgradeAllFreeNodesOk returns a tuple with the UpgradeAllFreeNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeAllFreeNodes

`func (o *UpgradeNodeParameters) SetUpgradeAllFreeNodes(v bool)`

SetUpgradeAllFreeNodes sets UpgradeAllFreeNodes field to given value.

### HasUpgradeAllFreeNodes

`func (o *UpgradeNodeParameters) HasUpgradeAllFreeNodes() bool`

HasUpgradeAllFreeNodes returns a boolean if a field has been set.

### SetUpgradeAllFreeNodesNil

`func (o *UpgradeNodeParameters) SetUpgradeAllFreeNodesNil(b bool)`

 SetUpgradeAllFreeNodesNil sets the value for UpgradeAllFreeNodes to be an explicit nil

### UnsetUpgradeAllFreeNodes
`func (o *UpgradeNodeParameters) UnsetUpgradeAllFreeNodes()`

UnsetUpgradeAllFreeNodes ensures that no value is present for UpgradeAllFreeNodes, not even an explicit nil
### GetUpgradeSelf

`func (o *UpgradeNodeParameters) GetUpgradeSelf() bool`

GetUpgradeSelf returns the UpgradeSelf field if non-nil, zero value otherwise.

### GetUpgradeSelfOk

`func (o *UpgradeNodeParameters) GetUpgradeSelfOk() (*bool, bool)`

GetUpgradeSelfOk returns a tuple with the UpgradeSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeSelf

`func (o *UpgradeNodeParameters) SetUpgradeSelf(v bool)`

SetUpgradeSelf sets UpgradeSelf field to given value.

### HasUpgradeSelf

`func (o *UpgradeNodeParameters) HasUpgradeSelf() bool`

HasUpgradeSelf returns a boolean if a field has been set.

### SetUpgradeSelfNil

`func (o *UpgradeNodeParameters) SetUpgradeSelfNil(b bool)`

 SetUpgradeSelfNil sets the value for UpgradeSelf to be an explicit nil

### UnsetUpgradeSelf
`func (o *UpgradeNodeParameters) UnsetUpgradeSelf()`

UnsetUpgradeSelf ensures that no value is present for UpgradeSelf, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



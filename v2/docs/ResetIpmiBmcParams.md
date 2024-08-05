# ResetIpmiBmcParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **NullableString** | Specifies the node id of the node for which ipmi bmc needs to be reset. This parameter is incompatible with &#39;nodeIp&#39;. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the node id of the node for which ipmi bmc needs to be reset. This parameter is incompatible with &#39;nodeId&#39;. | [optional] 

## Methods

### NewResetIpmiBmcParams

`func NewResetIpmiBmcParams() *ResetIpmiBmcParams`

NewResetIpmiBmcParams instantiates a new ResetIpmiBmcParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetIpmiBmcParamsWithDefaults

`func NewResetIpmiBmcParamsWithDefaults() *ResetIpmiBmcParams`

NewResetIpmiBmcParamsWithDefaults instantiates a new ResetIpmiBmcParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *ResetIpmiBmcParams) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ResetIpmiBmcParams) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ResetIpmiBmcParams) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ResetIpmiBmcParams) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *ResetIpmiBmcParams) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *ResetIpmiBmcParams) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeIp

`func (o *ResetIpmiBmcParams) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *ResetIpmiBmcParams) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *ResetIpmiBmcParams) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *ResetIpmiBmcParams) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *ResetIpmiBmcParams) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *ResetIpmiBmcParams) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



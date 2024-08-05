# NodeIpmiUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpmiUsername** | Pointer to **NullableString** | Specifies the ipmi user name of the node. | [optional] 
**NodeIp** | Pointer to **NullableString** | Specifies the ip address of the node. | [optional] 

## Methods

### NewNodeIpmiUser

`func NewNodeIpmiUser() *NodeIpmiUser`

NewNodeIpmiUser instantiates a new NodeIpmiUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeIpmiUserWithDefaults

`func NewNodeIpmiUserWithDefaults() *NodeIpmiUser`

NewNodeIpmiUserWithDefaults instantiates a new NodeIpmiUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpmiUsername

`func (o *NodeIpmiUser) GetIpmiUsername() string`

GetIpmiUsername returns the IpmiUsername field if non-nil, zero value otherwise.

### GetIpmiUsernameOk

`func (o *NodeIpmiUser) GetIpmiUsernameOk() (*string, bool)`

GetIpmiUsernameOk returns a tuple with the IpmiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiUsername

`func (o *NodeIpmiUser) SetIpmiUsername(v string)`

SetIpmiUsername sets IpmiUsername field to given value.

### HasIpmiUsername

`func (o *NodeIpmiUser) HasIpmiUsername() bool`

HasIpmiUsername returns a boolean if a field has been set.

### SetIpmiUsernameNil

`func (o *NodeIpmiUser) SetIpmiUsernameNil(b bool)`

 SetIpmiUsernameNil sets the value for IpmiUsername to be an explicit nil

### UnsetIpmiUsername
`func (o *NodeIpmiUser) UnsetIpmiUsername()`

UnsetIpmiUsername ensures that no value is present for IpmiUsername, not even an explicit nil
### GetNodeIp

`func (o *NodeIpmiUser) GetNodeIp() string`

GetNodeIp returns the NodeIp field if non-nil, zero value otherwise.

### GetNodeIpOk

`func (o *NodeIpmiUser) GetNodeIpOk() (*string, bool)`

GetNodeIpOk returns a tuple with the NodeIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIp

`func (o *NodeIpmiUser) SetNodeIp(v string)`

SetNodeIp sets NodeIp field to given value.

### HasNodeIp

`func (o *NodeIpmiUser) HasNodeIp() bool`

HasNodeIp returns a boolean if a field has been set.

### SetNodeIpNil

`func (o *NodeIpmiUser) SetNodeIpNil(b bool)`

 SetNodeIpNil sets the value for NodeIp to be an explicit nil

### UnsetNodeIp
`func (o *NodeIpmiUser) UnsetNodeIp()`

UnsetNodeIp ensures that no value is present for NodeIp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



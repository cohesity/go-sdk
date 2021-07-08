# NodeIpmiUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpmiPassword** | Pointer to **NullableString** | In request, IPMI password is entered by the user. In response, it is is set to empty and hence will be omitted. | [optional] 
**IpmiUser** | Pointer to **NullableString** |  | [optional] 
**NodeIp** | Pointer to **NullableString** |  | [optional] 

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

### GetIpmiPassword

`func (o *NodeIpmiUser) GetIpmiPassword() string`

GetIpmiPassword returns the IpmiPassword field if non-nil, zero value otherwise.

### GetIpmiPasswordOk

`func (o *NodeIpmiUser) GetIpmiPasswordOk() (*string, bool)`

GetIpmiPasswordOk returns a tuple with the IpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiPassword

`func (o *NodeIpmiUser) SetIpmiPassword(v string)`

SetIpmiPassword sets IpmiPassword field to given value.

### HasIpmiPassword

`func (o *NodeIpmiUser) HasIpmiPassword() bool`

HasIpmiPassword returns a boolean if a field has been set.

### SetIpmiPasswordNil

`func (o *NodeIpmiUser) SetIpmiPasswordNil(b bool)`

 SetIpmiPasswordNil sets the value for IpmiPassword to be an explicit nil

### UnsetIpmiPassword
`func (o *NodeIpmiUser) UnsetIpmiPassword()`

UnsetIpmiPassword ensures that no value is present for IpmiPassword, not even an explicit nil
### GetIpmiUser

`func (o *NodeIpmiUser) GetIpmiUser() string`

GetIpmiUser returns the IpmiUser field if non-nil, zero value otherwise.

### GetIpmiUserOk

`func (o *NodeIpmiUser) GetIpmiUserOk() (*string, bool)`

GetIpmiUserOk returns a tuple with the IpmiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiUser

`func (o *NodeIpmiUser) SetIpmiUser(v string)`

SetIpmiUser sets IpmiUser field to given value.

### HasIpmiUser

`func (o *NodeIpmiUser) HasIpmiUser() bool`

HasIpmiUser returns a boolean if a field has been set.

### SetIpmiUserNil

`func (o *NodeIpmiUser) SetIpmiUserNil(b bool)`

 SetIpmiUserNil sets the value for IpmiUser to be an explicit nil

### UnsetIpmiUser
`func (o *NodeIpmiUser) UnsetIpmiUser()`

UnsetIpmiUser ensures that no value is present for IpmiUser, not even an explicit nil
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



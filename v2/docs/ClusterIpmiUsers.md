# ClusterIpmiUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIpmiUsername** | Pointer to **NullableString** | IPMI user name at the cluster level for the cluster. | [optional] 
**NodeIpmiUsers** | Pointer to [**[]NodeIpmiUser**](NodeIpmiUser.md) | Specifies the ipmi user info for all the nodes in the cluster. | [optional] 

## Methods

### NewClusterIpmiUsers

`func NewClusterIpmiUsers() *ClusterIpmiUsers`

NewClusterIpmiUsers instantiates a new ClusterIpmiUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterIpmiUsersWithDefaults

`func NewClusterIpmiUsersWithDefaults() *ClusterIpmiUsers`

NewClusterIpmiUsersWithDefaults instantiates a new ClusterIpmiUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIpmiUsername

`func (o *ClusterIpmiUsers) GetClusterIpmiUsername() string`

GetClusterIpmiUsername returns the ClusterIpmiUsername field if non-nil, zero value otherwise.

### GetClusterIpmiUsernameOk

`func (o *ClusterIpmiUsers) GetClusterIpmiUsernameOk() (*string, bool)`

GetClusterIpmiUsernameOk returns a tuple with the ClusterIpmiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpmiUsername

`func (o *ClusterIpmiUsers) SetClusterIpmiUsername(v string)`

SetClusterIpmiUsername sets ClusterIpmiUsername field to given value.

### HasClusterIpmiUsername

`func (o *ClusterIpmiUsers) HasClusterIpmiUsername() bool`

HasClusterIpmiUsername returns a boolean if a field has been set.

### SetClusterIpmiUsernameNil

`func (o *ClusterIpmiUsers) SetClusterIpmiUsernameNil(b bool)`

 SetClusterIpmiUsernameNil sets the value for ClusterIpmiUsername to be an explicit nil

### UnsetClusterIpmiUsername
`func (o *ClusterIpmiUsers) UnsetClusterIpmiUsername()`

UnsetClusterIpmiUsername ensures that no value is present for ClusterIpmiUsername, not even an explicit nil
### GetNodeIpmiUsers

`func (o *ClusterIpmiUsers) GetNodeIpmiUsers() []NodeIpmiUser`

GetNodeIpmiUsers returns the NodeIpmiUsers field if non-nil, zero value otherwise.

### GetNodeIpmiUsersOk

`func (o *ClusterIpmiUsers) GetNodeIpmiUsersOk() (*[]NodeIpmiUser, bool)`

GetNodeIpmiUsersOk returns a tuple with the NodeIpmiUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpmiUsers

`func (o *ClusterIpmiUsers) SetNodeIpmiUsers(v []NodeIpmiUser)`

SetNodeIpmiUsers sets NodeIpmiUsers field to given value.

### HasNodeIpmiUsers

`func (o *ClusterIpmiUsers) HasNodeIpmiUsers() bool`

HasNodeIpmiUsers returns a boolean if a field has been set.

### SetNodeIpmiUsersNil

`func (o *ClusterIpmiUsers) SetNodeIpmiUsersNil(b bool)`

 SetNodeIpmiUsersNil sets the value for NodeIpmiUsers to be an explicit nil

### UnsetNodeIpmiUsers
`func (o *ClusterIpmiUsers) UnsetNodeIpmiUsers()`

UnsetNodeIpmiUsers ensures that no value is present for NodeIpmiUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



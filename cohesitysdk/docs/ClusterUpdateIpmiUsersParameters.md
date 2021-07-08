# ClusterUpdateIpmiUsersParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIpmiUser** | Pointer to **NullableString** | Cluster level IPMI User name. | [optional] 
**IpmiPassword** | Pointer to **NullableString** | Cluster level IPMI Password. | [optional] 
**NodeIpmiUsers** | Pointer to [**[]NodeIpmiUser**](NodeIpmiUser.md) | Node level IPMI User config. | [optional] 

## Methods

### NewClusterUpdateIpmiUsersParameters

`func NewClusterUpdateIpmiUsersParameters() *ClusterUpdateIpmiUsersParameters`

NewClusterUpdateIpmiUsersParameters instantiates a new ClusterUpdateIpmiUsersParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUpdateIpmiUsersParametersWithDefaults

`func NewClusterUpdateIpmiUsersParametersWithDefaults() *ClusterUpdateIpmiUsersParameters`

NewClusterUpdateIpmiUsersParametersWithDefaults instantiates a new ClusterUpdateIpmiUsersParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIpmiUser

`func (o *ClusterUpdateIpmiUsersParameters) GetClusterIpmiUser() string`

GetClusterIpmiUser returns the ClusterIpmiUser field if non-nil, zero value otherwise.

### GetClusterIpmiUserOk

`func (o *ClusterUpdateIpmiUsersParameters) GetClusterIpmiUserOk() (*string, bool)`

GetClusterIpmiUserOk returns a tuple with the ClusterIpmiUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpmiUser

`func (o *ClusterUpdateIpmiUsersParameters) SetClusterIpmiUser(v string)`

SetClusterIpmiUser sets ClusterIpmiUser field to given value.

### HasClusterIpmiUser

`func (o *ClusterUpdateIpmiUsersParameters) HasClusterIpmiUser() bool`

HasClusterIpmiUser returns a boolean if a field has been set.

### SetClusterIpmiUserNil

`func (o *ClusterUpdateIpmiUsersParameters) SetClusterIpmiUserNil(b bool)`

 SetClusterIpmiUserNil sets the value for ClusterIpmiUser to be an explicit nil

### UnsetClusterIpmiUser
`func (o *ClusterUpdateIpmiUsersParameters) UnsetClusterIpmiUser()`

UnsetClusterIpmiUser ensures that no value is present for ClusterIpmiUser, not even an explicit nil
### GetIpmiPassword

`func (o *ClusterUpdateIpmiUsersParameters) GetIpmiPassword() string`

GetIpmiPassword returns the IpmiPassword field if non-nil, zero value otherwise.

### GetIpmiPasswordOk

`func (o *ClusterUpdateIpmiUsersParameters) GetIpmiPasswordOk() (*string, bool)`

GetIpmiPasswordOk returns a tuple with the IpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpmiPassword

`func (o *ClusterUpdateIpmiUsersParameters) SetIpmiPassword(v string)`

SetIpmiPassword sets IpmiPassword field to given value.

### HasIpmiPassword

`func (o *ClusterUpdateIpmiUsersParameters) HasIpmiPassword() bool`

HasIpmiPassword returns a boolean if a field has been set.

### SetIpmiPasswordNil

`func (o *ClusterUpdateIpmiUsersParameters) SetIpmiPasswordNil(b bool)`

 SetIpmiPasswordNil sets the value for IpmiPassword to be an explicit nil

### UnsetIpmiPassword
`func (o *ClusterUpdateIpmiUsersParameters) UnsetIpmiPassword()`

UnsetIpmiPassword ensures that no value is present for IpmiPassword, not even an explicit nil
### GetNodeIpmiUsers

`func (o *ClusterUpdateIpmiUsersParameters) GetNodeIpmiUsers() []NodeIpmiUser`

GetNodeIpmiUsers returns the NodeIpmiUsers field if non-nil, zero value otherwise.

### GetNodeIpmiUsersOk

`func (o *ClusterUpdateIpmiUsersParameters) GetNodeIpmiUsersOk() (*[]NodeIpmiUser, bool)`

GetNodeIpmiUsersOk returns a tuple with the NodeIpmiUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpmiUsers

`func (o *ClusterUpdateIpmiUsersParameters) SetNodeIpmiUsers(v []NodeIpmiUser)`

SetNodeIpmiUsers sets NodeIpmiUsers field to given value.

### HasNodeIpmiUsers

`func (o *ClusterUpdateIpmiUsersParameters) HasNodeIpmiUsers() bool`

HasNodeIpmiUsers returns a boolean if a field has been set.

### SetNodeIpmiUsersNil

`func (o *ClusterUpdateIpmiUsersParameters) SetNodeIpmiUsersNil(b bool)`

 SetNodeIpmiUsersNil sets the value for NodeIpmiUsers to be an explicit nil

### UnsetNodeIpmiUsers
`func (o *ClusterUpdateIpmiUsersParameters) UnsetNodeIpmiUsers()`

UnsetNodeIpmiUsers ensures that no value is present for NodeIpmiUsers, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



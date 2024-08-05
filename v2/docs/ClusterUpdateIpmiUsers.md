# ClusterUpdateIpmiUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterIpmiPassword** | Pointer to **NullableString** | IPMI password at the cluster level for the cluster. | [optional] 
**ClusterIpmiUsername** | Pointer to **NullableString** | IPMI user name at the cluster level for the cluster. | [optional] 
**NodeIpmiPasswords** | Pointer to **[]string** | Specifies the ipmi passwords for all the nodes in the cluster. | [optional] 
**NodeIpmiUsernames** | Pointer to **[]string** | Specifies the ipmi usernames for all the nodes in the cluster. | [optional] 
**NodeIps** | Pointer to **[]string** | Specifies the ip addresses for all the nodes in the cluster. | [optional] 

## Methods

### NewClusterUpdateIpmiUsers

`func NewClusterUpdateIpmiUsers() *ClusterUpdateIpmiUsers`

NewClusterUpdateIpmiUsers instantiates a new ClusterUpdateIpmiUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterUpdateIpmiUsersWithDefaults

`func NewClusterUpdateIpmiUsersWithDefaults() *ClusterUpdateIpmiUsers`

NewClusterUpdateIpmiUsersWithDefaults instantiates a new ClusterUpdateIpmiUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterIpmiPassword

`func (o *ClusterUpdateIpmiUsers) GetClusterIpmiPassword() string`

GetClusterIpmiPassword returns the ClusterIpmiPassword field if non-nil, zero value otherwise.

### GetClusterIpmiPasswordOk

`func (o *ClusterUpdateIpmiUsers) GetClusterIpmiPasswordOk() (*string, bool)`

GetClusterIpmiPasswordOk returns a tuple with the ClusterIpmiPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpmiPassword

`func (o *ClusterUpdateIpmiUsers) SetClusterIpmiPassword(v string)`

SetClusterIpmiPassword sets ClusterIpmiPassword field to given value.

### HasClusterIpmiPassword

`func (o *ClusterUpdateIpmiUsers) HasClusterIpmiPassword() bool`

HasClusterIpmiPassword returns a boolean if a field has been set.

### SetClusterIpmiPasswordNil

`func (o *ClusterUpdateIpmiUsers) SetClusterIpmiPasswordNil(b bool)`

 SetClusterIpmiPasswordNil sets the value for ClusterIpmiPassword to be an explicit nil

### UnsetClusterIpmiPassword
`func (o *ClusterUpdateIpmiUsers) UnsetClusterIpmiPassword()`

UnsetClusterIpmiPassword ensures that no value is present for ClusterIpmiPassword, not even an explicit nil
### GetClusterIpmiUsername

`func (o *ClusterUpdateIpmiUsers) GetClusterIpmiUsername() string`

GetClusterIpmiUsername returns the ClusterIpmiUsername field if non-nil, zero value otherwise.

### GetClusterIpmiUsernameOk

`func (o *ClusterUpdateIpmiUsers) GetClusterIpmiUsernameOk() (*string, bool)`

GetClusterIpmiUsernameOk returns a tuple with the ClusterIpmiUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterIpmiUsername

`func (o *ClusterUpdateIpmiUsers) SetClusterIpmiUsername(v string)`

SetClusterIpmiUsername sets ClusterIpmiUsername field to given value.

### HasClusterIpmiUsername

`func (o *ClusterUpdateIpmiUsers) HasClusterIpmiUsername() bool`

HasClusterIpmiUsername returns a boolean if a field has been set.

### SetClusterIpmiUsernameNil

`func (o *ClusterUpdateIpmiUsers) SetClusterIpmiUsernameNil(b bool)`

 SetClusterIpmiUsernameNil sets the value for ClusterIpmiUsername to be an explicit nil

### UnsetClusterIpmiUsername
`func (o *ClusterUpdateIpmiUsers) UnsetClusterIpmiUsername()`

UnsetClusterIpmiUsername ensures that no value is present for ClusterIpmiUsername, not even an explicit nil
### GetNodeIpmiPasswords

`func (o *ClusterUpdateIpmiUsers) GetNodeIpmiPasswords() []string`

GetNodeIpmiPasswords returns the NodeIpmiPasswords field if non-nil, zero value otherwise.

### GetNodeIpmiPasswordsOk

`func (o *ClusterUpdateIpmiUsers) GetNodeIpmiPasswordsOk() (*[]string, bool)`

GetNodeIpmiPasswordsOk returns a tuple with the NodeIpmiPasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpmiPasswords

`func (o *ClusterUpdateIpmiUsers) SetNodeIpmiPasswords(v []string)`

SetNodeIpmiPasswords sets NodeIpmiPasswords field to given value.

### HasNodeIpmiPasswords

`func (o *ClusterUpdateIpmiUsers) HasNodeIpmiPasswords() bool`

HasNodeIpmiPasswords returns a boolean if a field has been set.

### SetNodeIpmiPasswordsNil

`func (o *ClusterUpdateIpmiUsers) SetNodeIpmiPasswordsNil(b bool)`

 SetNodeIpmiPasswordsNil sets the value for NodeIpmiPasswords to be an explicit nil

### UnsetNodeIpmiPasswords
`func (o *ClusterUpdateIpmiUsers) UnsetNodeIpmiPasswords()`

UnsetNodeIpmiPasswords ensures that no value is present for NodeIpmiPasswords, not even an explicit nil
### GetNodeIpmiUsernames

`func (o *ClusterUpdateIpmiUsers) GetNodeIpmiUsernames() []string`

GetNodeIpmiUsernames returns the NodeIpmiUsernames field if non-nil, zero value otherwise.

### GetNodeIpmiUsernamesOk

`func (o *ClusterUpdateIpmiUsers) GetNodeIpmiUsernamesOk() (*[]string, bool)`

GetNodeIpmiUsernamesOk returns a tuple with the NodeIpmiUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIpmiUsernames

`func (o *ClusterUpdateIpmiUsers) SetNodeIpmiUsernames(v []string)`

SetNodeIpmiUsernames sets NodeIpmiUsernames field to given value.

### HasNodeIpmiUsernames

`func (o *ClusterUpdateIpmiUsers) HasNodeIpmiUsernames() bool`

HasNodeIpmiUsernames returns a boolean if a field has been set.

### SetNodeIpmiUsernamesNil

`func (o *ClusterUpdateIpmiUsers) SetNodeIpmiUsernamesNil(b bool)`

 SetNodeIpmiUsernamesNil sets the value for NodeIpmiUsernames to be an explicit nil

### UnsetNodeIpmiUsernames
`func (o *ClusterUpdateIpmiUsers) UnsetNodeIpmiUsernames()`

UnsetNodeIpmiUsernames ensures that no value is present for NodeIpmiUsernames, not even an explicit nil
### GetNodeIps

`func (o *ClusterUpdateIpmiUsers) GetNodeIps() []string`

GetNodeIps returns the NodeIps field if non-nil, zero value otherwise.

### GetNodeIpsOk

`func (o *ClusterUpdateIpmiUsers) GetNodeIpsOk() (*[]string, bool)`

GetNodeIpsOk returns a tuple with the NodeIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeIps

`func (o *ClusterUpdateIpmiUsers) SetNodeIps(v []string)`

SetNodeIps sets NodeIps field to given value.

### HasNodeIps

`func (o *ClusterUpdateIpmiUsers) HasNodeIps() bool`

HasNodeIps returns a boolean if a field has been set.

### SetNodeIpsNil

`func (o *ClusterUpdateIpmiUsers) SetNodeIpsNil(b bool)`

 SetNodeIpsNil sets the value for NodeIps to be an explicit nil

### UnsetNodeIps
`func (o *ClusterUpdateIpmiUsers) UnsetNodeIps()`

UnsetNodeIps ensures that no value is present for NodeIps, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



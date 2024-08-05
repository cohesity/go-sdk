# ValidateRemoteClusterConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeAddresses** | **[]string** | Specifies the VIP or IP addresses of the Nodes on the Remote Cluster to connect with. Hostnames are not supported. | 
**Password** | **string** | Specifies the password for Cohesity user to use when connecting to the Remote Cluster. | 
**Username** | **string** | Specifies the Cohesity user name used to connect to the Remote Cluster. | 

## Methods

### NewValidateRemoteClusterConnectionParams

`func NewValidateRemoteClusterConnectionParams(nodeAddresses []string, password string, username string, ) *ValidateRemoteClusterConnectionParams`

NewValidateRemoteClusterConnectionParams instantiates a new ValidateRemoteClusterConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateRemoteClusterConnectionParamsWithDefaults

`func NewValidateRemoteClusterConnectionParamsWithDefaults() *ValidateRemoteClusterConnectionParams`

NewValidateRemoteClusterConnectionParamsWithDefaults instantiates a new ValidateRemoteClusterConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeAddresses

`func (o *ValidateRemoteClusterConnectionParams) GetNodeAddresses() []string`

GetNodeAddresses returns the NodeAddresses field if non-nil, zero value otherwise.

### GetNodeAddressesOk

`func (o *ValidateRemoteClusterConnectionParams) GetNodeAddressesOk() (*[]string, bool)`

GetNodeAddressesOk returns a tuple with the NodeAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddresses

`func (o *ValidateRemoteClusterConnectionParams) SetNodeAddresses(v []string)`

SetNodeAddresses sets NodeAddresses field to given value.


### GetPassword

`func (o *ValidateRemoteClusterConnectionParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ValidateRemoteClusterConnectionParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ValidateRemoteClusterConnectionParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *ValidateRemoteClusterConnectionParams) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ValidateRemoteClusterConnectionParams) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ValidateRemoteClusterConnectionParams) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



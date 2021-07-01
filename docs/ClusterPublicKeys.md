# ClusterPublicKeys

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshPublicKey** | Pointer to **NullableString** | Specifies the SSH public key used to login to Cluster nodes. | [optional] 

## Methods

### NewClusterPublicKeys

`func NewClusterPublicKeys() *ClusterPublicKeys`

NewClusterPublicKeys instantiates a new ClusterPublicKeys object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterPublicKeysWithDefaults

`func NewClusterPublicKeysWithDefaults() *ClusterPublicKeys`

NewClusterPublicKeysWithDefaults instantiates a new ClusterPublicKeys object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshPublicKey

`func (o *ClusterPublicKeys) GetSshPublicKey() string`

GetSshPublicKey returns the SshPublicKey field if non-nil, zero value otherwise.

### GetSshPublicKeyOk

`func (o *ClusterPublicKeys) GetSshPublicKeyOk() (*string, bool)`

GetSshPublicKeyOk returns a tuple with the SshPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKey

`func (o *ClusterPublicKeys) SetSshPublicKey(v string)`

SetSshPublicKey sets SshPublicKey field to given value.

### HasSshPublicKey

`func (o *ClusterPublicKeys) HasSshPublicKey() bool`

HasSshPublicKey returns a boolean if a field has been set.

### SetSshPublicKeyNil

`func (o *ClusterPublicKeys) SetSshPublicKeyNil(b bool)`

 SetSshPublicKeyNil sets the value for SshPublicKey to be an explicit nil

### UnsetSshPublicKey
`func (o *ClusterPublicKeys) UnsetSshPublicKey()`

UnsetSshPublicKey ensures that no value is present for SshPublicKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RemoteTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **NullableInt64** | Specifies the cluster id of the target replication cluster. | 
**ClusterName** | Pointer to **NullableString** | Specifies the cluster name of the target replication cluster. | [optional] [readonly] 

## Methods

### NewRemoteTargetConfig

`func NewRemoteTargetConfig(clusterId NullableInt64, ) *RemoteTargetConfig`

NewRemoteTargetConfig instantiates a new RemoteTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteTargetConfigWithDefaults

`func NewRemoteTargetConfigWithDefaults() *RemoteTargetConfig`

NewRemoteTargetConfigWithDefaults instantiates a new RemoteTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *RemoteTargetConfig) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *RemoteTargetConfig) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *RemoteTargetConfig) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### SetClusterIdNil

`func (o *RemoteTargetConfig) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *RemoteTargetConfig) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterName

`func (o *RemoteTargetConfig) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *RemoteTargetConfig) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *RemoteTargetConfig) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *RemoteTargetConfig) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *RemoteTargetConfig) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *RemoteTargetConfig) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



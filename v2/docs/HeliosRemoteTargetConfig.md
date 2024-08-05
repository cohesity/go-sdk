# HeliosRemoteTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **NullableInt64** | Specifies the cluster id of the target replication cluster. | 
**ClusterName** | Pointer to **NullableString** | Specifies the cluster name of the target replication cluster. | [optional] [readonly] 

## Methods

### NewHeliosRemoteTargetConfig

`func NewHeliosRemoteTargetConfig(clusterId NullableInt64, ) *HeliosRemoteTargetConfig`

NewHeliosRemoteTargetConfig instantiates a new HeliosRemoteTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeliosRemoteTargetConfigWithDefaults

`func NewHeliosRemoteTargetConfigWithDefaults() *HeliosRemoteTargetConfig`

NewHeliosRemoteTargetConfigWithDefaults instantiates a new HeliosRemoteTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *HeliosRemoteTargetConfig) GetClusterId() int64`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *HeliosRemoteTargetConfig) GetClusterIdOk() (*int64, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *HeliosRemoteTargetConfig) SetClusterId(v int64)`

SetClusterId sets ClusterId field to given value.


### SetClusterIdNil

`func (o *HeliosRemoteTargetConfig) SetClusterIdNil(b bool)`

 SetClusterIdNil sets the value for ClusterId to be an explicit nil

### UnsetClusterId
`func (o *HeliosRemoteTargetConfig) UnsetClusterId()`

UnsetClusterId ensures that no value is present for ClusterId, not even an explicit nil
### GetClusterName

`func (o *HeliosRemoteTargetConfig) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *HeliosRemoteTargetConfig) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *HeliosRemoteTargetConfig) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *HeliosRemoteTargetConfig) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### SetClusterNameNil

`func (o *HeliosRemoteTargetConfig) SetClusterNameNil(b bool)`

 SetClusterNameNil sets the value for ClusterName to be an explicit nil

### UnsetClusterName
`func (o *HeliosRemoteTargetConfig) UnsetClusterName()`

UnsetClusterName ensures that no value is present for ClusterName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CascadedTargetConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RemoteTargets** | [**TargetsConfiguration**](TargetsConfiguration.md) |  | 
**SourceClusterId** | **NullableInt64** | Specifies the source cluster id from where the remote operations will be performed to the next set of remote targets. | 

## Methods

### NewCascadedTargetConfiguration

`func NewCascadedTargetConfiguration(remoteTargets TargetsConfiguration, sourceClusterId NullableInt64, ) *CascadedTargetConfiguration`

NewCascadedTargetConfiguration instantiates a new CascadedTargetConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCascadedTargetConfigurationWithDefaults

`func NewCascadedTargetConfigurationWithDefaults() *CascadedTargetConfiguration`

NewCascadedTargetConfigurationWithDefaults instantiates a new CascadedTargetConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRemoteTargets

`func (o *CascadedTargetConfiguration) GetRemoteTargets() TargetsConfiguration`

GetRemoteTargets returns the RemoteTargets field if non-nil, zero value otherwise.

### GetRemoteTargetsOk

`func (o *CascadedTargetConfiguration) GetRemoteTargetsOk() (*TargetsConfiguration, bool)`

GetRemoteTargetsOk returns a tuple with the RemoteTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTargets

`func (o *CascadedTargetConfiguration) SetRemoteTargets(v TargetsConfiguration)`

SetRemoteTargets sets RemoteTargets field to given value.


### GetSourceClusterId

`func (o *CascadedTargetConfiguration) GetSourceClusterId() int64`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### GetSourceClusterIdOk

`func (o *CascadedTargetConfiguration) GetSourceClusterIdOk() (*int64, bool)`

GetSourceClusterIdOk returns a tuple with the SourceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceClusterId

`func (o *CascadedTargetConfiguration) SetSourceClusterId(v int64)`

SetSourceClusterId sets SourceClusterId field to given value.


### SetSourceClusterIdNil

`func (o *CascadedTargetConfiguration) SetSourceClusterIdNil(b bool)`

 SetSourceClusterIdNil sets the value for SourceClusterId to be an explicit nil

### UnsetSourceClusterId
`func (o *CascadedTargetConfiguration) UnsetSourceClusterId()`

UnsetSourceClusterId ensures that no value is present for SourceClusterId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



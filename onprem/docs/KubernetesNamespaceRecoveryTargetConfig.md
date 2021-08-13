# KubernetesNamespaceRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **NullableBool** | Specifies whether or not to recover the Namespaces to a different source than they were backed up from. | 
**NewSourceConfig** | Pointer to [**NullableKubernetesNamespaceRecoveryNewSourceConfig**](KubernetesNamespaceRecoveryNewSourceConfig.md) | Specifies the new source configuration if a Kubernetes Namespace is being restored to a different source than the one from which it was protected. | [optional] 

## Methods

### NewKubernetesNamespaceRecoveryTargetConfig

`func NewKubernetesNamespaceRecoveryTargetConfig(recoverToNewSource NullableBool, ) *KubernetesNamespaceRecoveryTargetConfig`

NewKubernetesNamespaceRecoveryTargetConfig instantiates a new KubernetesNamespaceRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNamespaceRecoveryTargetConfigWithDefaults

`func NewKubernetesNamespaceRecoveryTargetConfigWithDefaults() *KubernetesNamespaceRecoveryTargetConfig`

NewKubernetesNamespaceRecoveryTargetConfigWithDefaults instantiates a new KubernetesNamespaceRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *KubernetesNamespaceRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *KubernetesNamespaceRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *KubernetesNamespaceRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### SetRecoverToNewSourceNil

`func (o *KubernetesNamespaceRecoveryTargetConfig) SetRecoverToNewSourceNil(b bool)`

 SetRecoverToNewSourceNil sets the value for RecoverToNewSource to be an explicit nil

### UnsetRecoverToNewSource
`func (o *KubernetesNamespaceRecoveryTargetConfig) UnsetRecoverToNewSource()`

UnsetRecoverToNewSource ensures that no value is present for RecoverToNewSource, not even an explicit nil
### GetNewSourceConfig

`func (o *KubernetesNamespaceRecoveryTargetConfig) GetNewSourceConfig() KubernetesNamespaceRecoveryNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *KubernetesNamespaceRecoveryTargetConfig) GetNewSourceConfigOk() (*KubernetesNamespaceRecoveryNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *KubernetesNamespaceRecoveryTargetConfig) SetNewSourceConfig(v KubernetesNamespaceRecoveryNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *KubernetesNamespaceRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *KubernetesNamespaceRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *KubernetesNamespaceRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



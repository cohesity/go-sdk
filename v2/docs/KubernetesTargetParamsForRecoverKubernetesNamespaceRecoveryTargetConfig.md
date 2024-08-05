# KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableKubernetesNamespaceRecoveryTargetConfigNewSourceConfig**](KubernetesNamespaceRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **NullableBool** | Specifies whether or not to recover the Namespaces to a different source than they were backed up from. | 

## Methods

### NewKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig

`func NewKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig(recoverToNewSource NullableBool, ) *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig`

NewKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfigWithDefaults

`func NewKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfigWithDefaults() *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig`

NewKubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfigWithDefaults instantiates a new KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) GetNewSourceConfig() KubernetesNamespaceRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) GetNewSourceConfigOk() (*KubernetesNamespaceRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) SetNewSourceConfig(v KubernetesNamespaceRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### SetRecoverToNewSourceNil

`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) SetRecoverToNewSourceNil(b bool)`

 SetRecoverToNewSourceNil sets the value for RecoverToNewSource to be an explicit nil

### UnsetRecoverToNewSource
`func (o *KubernetesTargetParamsForRecoverKubernetesNamespaceRecoveryTargetConfig) UnsetRecoverToNewSource()`

UnsetRecoverToNewSource ensures that no value is present for RecoverToNewSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



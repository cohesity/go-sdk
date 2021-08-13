# KubernetesNamespaceRecoveryNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the parent source to recover the Namespaces. | 

## Methods

### NewKubernetesNamespaceRecoveryNewSourceConfig

`func NewKubernetesNamespaceRecoveryNewSourceConfig(source NullableRecoveryObjectIdentifier, ) *KubernetesNamespaceRecoveryNewSourceConfig`

NewKubernetesNamespaceRecoveryNewSourceConfig instantiates a new KubernetesNamespaceRecoveryNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNamespaceRecoveryNewSourceConfigWithDefaults

`func NewKubernetesNamespaceRecoveryNewSourceConfigWithDefaults() *KubernetesNamespaceRecoveryNewSourceConfig`

NewKubernetesNamespaceRecoveryNewSourceConfigWithDefaults instantiates a new KubernetesNamespaceRecoveryNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *KubernetesNamespaceRecoveryNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KubernetesNamespaceRecoveryNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KubernetesNamespaceRecoveryNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *KubernetesNamespaceRecoveryNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *KubernetesNamespaceRecoveryNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



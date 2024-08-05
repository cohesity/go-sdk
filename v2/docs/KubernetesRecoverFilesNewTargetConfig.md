# KubernetesRecoverFilesNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the path location to recover files to. | 
**TargetNamespace** | Pointer to [**NullableKubernetesRecoverFilesNewTargetConfigTargetNamespace**](KubernetesRecoverFilesNewTargetConfigTargetNamespace.md) |  | [optional] 
**TargetPvc** | [**NullableKubernetesRecoverFilesNewTargetConfigTargetPvc**](KubernetesRecoverFilesNewTargetConfigTargetPvc.md) |  | 
**TargetSource** | Pointer to [**NullableKubernetesRecoverFilesNewTargetConfigTargetSource**](KubernetesRecoverFilesNewTargetConfigTargetSource.md) |  | [optional] 

## Methods

### NewKubernetesRecoverFilesNewTargetConfig

`func NewKubernetesRecoverFilesNewTargetConfig(absolutePath NullableString, targetPvc NullableKubernetesRecoverFilesNewTargetConfigTargetPvc, ) *KubernetesRecoverFilesNewTargetConfig`

NewKubernetesRecoverFilesNewTargetConfig instantiates a new KubernetesRecoverFilesNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesRecoverFilesNewTargetConfigWithDefaults

`func NewKubernetesRecoverFilesNewTargetConfigWithDefaults() *KubernetesRecoverFilesNewTargetConfig`

NewKubernetesRecoverFilesNewTargetConfigWithDefaults instantiates a new KubernetesRecoverFilesNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *KubernetesRecoverFilesNewTargetConfig) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *KubernetesRecoverFilesNewTargetConfig) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *KubernetesRecoverFilesNewTargetConfig) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *KubernetesRecoverFilesNewTargetConfig) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *KubernetesRecoverFilesNewTargetConfig) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetTargetNamespace

`func (o *KubernetesRecoverFilesNewTargetConfig) GetTargetNamespace() KubernetesRecoverFilesNewTargetConfigTargetNamespace`

GetTargetNamespace returns the TargetNamespace field if non-nil, zero value otherwise.

### GetTargetNamespaceOk

`func (o *KubernetesRecoverFilesNewTargetConfig) GetTargetNamespaceOk() (*KubernetesRecoverFilesNewTargetConfigTargetNamespace, bool)`

GetTargetNamespaceOk returns a tuple with the TargetNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetNamespace

`func (o *KubernetesRecoverFilesNewTargetConfig) SetTargetNamespace(v KubernetesRecoverFilesNewTargetConfigTargetNamespace)`

SetTargetNamespace sets TargetNamespace field to given value.

### HasTargetNamespace

`func (o *KubernetesRecoverFilesNewTargetConfig) HasTargetNamespace() bool`

HasTargetNamespace returns a boolean if a field has been set.

### SetTargetNamespaceNil

`func (o *KubernetesRecoverFilesNewTargetConfig) SetTargetNamespaceNil(b bool)`

 SetTargetNamespaceNil sets the value for TargetNamespace to be an explicit nil

### UnsetTargetNamespace
`func (o *KubernetesRecoverFilesNewTargetConfig) UnsetTargetNamespace()`

UnsetTargetNamespace ensures that no value is present for TargetNamespace, not even an explicit nil
### GetTargetPvc

`func (o *KubernetesRecoverFilesNewTargetConfig) GetTargetPvc() KubernetesRecoverFilesNewTargetConfigTargetPvc`

GetTargetPvc returns the TargetPvc field if non-nil, zero value otherwise.

### GetTargetPvcOk

`func (o *KubernetesRecoverFilesNewTargetConfig) GetTargetPvcOk() (*KubernetesRecoverFilesNewTargetConfigTargetPvc, bool)`

GetTargetPvcOk returns a tuple with the TargetPvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPvc

`func (o *KubernetesRecoverFilesNewTargetConfig) SetTargetPvc(v KubernetesRecoverFilesNewTargetConfigTargetPvc)`

SetTargetPvc sets TargetPvc field to given value.


### SetTargetPvcNil

`func (o *KubernetesRecoverFilesNewTargetConfig) SetTargetPvcNil(b bool)`

 SetTargetPvcNil sets the value for TargetPvc to be an explicit nil

### UnsetTargetPvc
`func (o *KubernetesRecoverFilesNewTargetConfig) UnsetTargetPvc()`

UnsetTargetPvc ensures that no value is present for TargetPvc, not even an explicit nil
### GetTargetSource

`func (o *KubernetesRecoverFilesNewTargetConfig) GetTargetSource() KubernetesRecoverFilesNewTargetConfigTargetSource`

GetTargetSource returns the TargetSource field if non-nil, zero value otherwise.

### GetTargetSourceOk

`func (o *KubernetesRecoverFilesNewTargetConfig) GetTargetSourceOk() (*KubernetesRecoverFilesNewTargetConfigTargetSource, bool)`

GetTargetSourceOk returns a tuple with the TargetSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSource

`func (o *KubernetesRecoverFilesNewTargetConfig) SetTargetSource(v KubernetesRecoverFilesNewTargetConfigTargetSource)`

SetTargetSource sets TargetSource field to given value.

### HasTargetSource

`func (o *KubernetesRecoverFilesNewTargetConfig) HasTargetSource() bool`

HasTargetSource returns a boolean if a field has been set.

### SetTargetSourceNil

`func (o *KubernetesRecoverFilesNewTargetConfig) SetTargetSourceNil(b bool)`

 SetTargetSourceNil sets the value for TargetSource to be an explicit nil

### UnsetTargetSource
`func (o *KubernetesRecoverFilesNewTargetConfig) UnsetTargetSource()`

UnsetTargetSource ensures that no value is present for TargetSource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



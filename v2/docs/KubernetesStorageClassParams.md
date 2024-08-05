# KubernetesStorageClassParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageClassMapping** | Pointer to [**[]KubernetesLabel**](KubernetesLabel.md) | Specifies mapping of storage classes | [optional] 
**UseStorageClassMapping** | Pointer to **NullableBool** | Specifies whether or not to use storage class mapping. | [optional] 

## Methods

### NewKubernetesStorageClassParams

`func NewKubernetesStorageClassParams() *KubernetesStorageClassParams`

NewKubernetesStorageClassParams instantiates a new KubernetesStorageClassParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesStorageClassParamsWithDefaults

`func NewKubernetesStorageClassParamsWithDefaults() *KubernetesStorageClassParams`

NewKubernetesStorageClassParamsWithDefaults instantiates a new KubernetesStorageClassParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageClassMapping

`func (o *KubernetesStorageClassParams) GetStorageClassMapping() []KubernetesLabel`

GetStorageClassMapping returns the StorageClassMapping field if non-nil, zero value otherwise.

### GetStorageClassMappingOk

`func (o *KubernetesStorageClassParams) GetStorageClassMappingOk() (*[]KubernetesLabel, bool)`

GetStorageClassMappingOk returns a tuple with the StorageClassMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassMapping

`func (o *KubernetesStorageClassParams) SetStorageClassMapping(v []KubernetesLabel)`

SetStorageClassMapping sets StorageClassMapping field to given value.

### HasStorageClassMapping

`func (o *KubernetesStorageClassParams) HasStorageClassMapping() bool`

HasStorageClassMapping returns a boolean if a field has been set.

### SetStorageClassMappingNil

`func (o *KubernetesStorageClassParams) SetStorageClassMappingNil(b bool)`

 SetStorageClassMappingNil sets the value for StorageClassMapping to be an explicit nil

### UnsetStorageClassMapping
`func (o *KubernetesStorageClassParams) UnsetStorageClassMapping()`

UnsetStorageClassMapping ensures that no value is present for StorageClassMapping, not even an explicit nil
### GetUseStorageClassMapping

`func (o *KubernetesStorageClassParams) GetUseStorageClassMapping() bool`

GetUseStorageClassMapping returns the UseStorageClassMapping field if non-nil, zero value otherwise.

### GetUseStorageClassMappingOk

`func (o *KubernetesStorageClassParams) GetUseStorageClassMappingOk() (*bool, bool)`

GetUseStorageClassMappingOk returns a tuple with the UseStorageClassMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseStorageClassMapping

`func (o *KubernetesStorageClassParams) SetUseStorageClassMapping(v bool)`

SetUseStorageClassMapping sets UseStorageClassMapping field to given value.

### HasUseStorageClassMapping

`func (o *KubernetesStorageClassParams) HasUseStorageClassMapping() bool`

HasUseStorageClassMapping returns a boolean if a field has been set.

### SetUseStorageClassMappingNil

`func (o *KubernetesStorageClassParams) SetUseStorageClassMappingNil(b bool)`

 SetUseStorageClassMappingNil sets the value for UseStorageClassMapping to be an explicit nil

### UnsetUseStorageClassMapping
`func (o *KubernetesStorageClassParams) UnsetUseStorageClassMapping()`

UnsetUseStorageClassMapping ensures that no value is present for UseStorageClassMapping, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# KubernetesNamespaceRecoveryNewSourceConfigSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableInt64** | Specifies the id of the object. | 
**Name** | Pointer to **NullableString** | Specifies the name of the object. | [optional] [readonly] 

## Methods

### NewKubernetesNamespaceRecoveryNewSourceConfigSource

`func NewKubernetesNamespaceRecoveryNewSourceConfigSource(id NullableInt64, ) *KubernetesNamespaceRecoveryNewSourceConfigSource`

NewKubernetesNamespaceRecoveryNewSourceConfigSource instantiates a new KubernetesNamespaceRecoveryNewSourceConfigSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesNamespaceRecoveryNewSourceConfigSourceWithDefaults

`func NewKubernetesNamespaceRecoveryNewSourceConfigSourceWithDefaults() *KubernetesNamespaceRecoveryNewSourceConfigSource`

NewKubernetesNamespaceRecoveryNewSourceConfigSourceWithDefaults instantiates a new KubernetesNamespaceRecoveryNewSourceConfigSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) SetId(v int64)`

SetId sets Id field to given value.


### SetIdNil

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KubernetesNamespaceRecoveryNewSourceConfigSource) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



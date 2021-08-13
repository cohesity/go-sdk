# RecoverKubernetesNamespaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**KubernetesTargetParams** | Pointer to [**NullableKubernetesTargetParamsForRecoverKubernetesNamespace**](KubernetesTargetParamsForRecoverKubernetesNamespace.md) | Specifies the params for recovering to a Kubernetes host. | [optional] 

## Methods

### NewRecoverKubernetesNamespaceParams

`func NewRecoverKubernetesNamespaceParams(targetEnvironment string, ) *RecoverKubernetesNamespaceParams`

NewRecoverKubernetesNamespaceParams instantiates a new RecoverKubernetesNamespaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKubernetesNamespaceParamsWithDefaults

`func NewRecoverKubernetesNamespaceParamsWithDefaults() *RecoverKubernetesNamespaceParams`

NewRecoverKubernetesNamespaceParamsWithDefaults instantiates a new RecoverKubernetesNamespaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetEnvironment

`func (o *RecoverKubernetesNamespaceParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverKubernetesNamespaceParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverKubernetesNamespaceParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetKubernetesTargetParams

`func (o *RecoverKubernetesNamespaceParams) GetKubernetesTargetParams() KubernetesTargetParamsForRecoverKubernetesNamespace`

GetKubernetesTargetParams returns the KubernetesTargetParams field if non-nil, zero value otherwise.

### GetKubernetesTargetParamsOk

`func (o *RecoverKubernetesNamespaceParams) GetKubernetesTargetParamsOk() (*KubernetesTargetParamsForRecoverKubernetesNamespace, bool)`

GetKubernetesTargetParamsOk returns a tuple with the KubernetesTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetParams

`func (o *RecoverKubernetesNamespaceParams) SetKubernetesTargetParams(v KubernetesTargetParamsForRecoverKubernetesNamespace)`

SetKubernetesTargetParams sets KubernetesTargetParams field to given value.

### HasKubernetesTargetParams

`func (o *RecoverKubernetesNamespaceParams) HasKubernetesTargetParams() bool`

HasKubernetesTargetParams returns a boolean if a field has been set.

### SetKubernetesTargetParamsNil

`func (o *RecoverKubernetesNamespaceParams) SetKubernetesTargetParamsNil(b bool)`

 SetKubernetesTargetParamsNil sets the value for KubernetesTargetParams to be an explicit nil

### UnsetKubernetesTargetParams
`func (o *RecoverKubernetesNamespaceParams) UnsetKubernetesTargetParams()`

UnsetKubernetesTargetParams ensures that no value is present for KubernetesTargetParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



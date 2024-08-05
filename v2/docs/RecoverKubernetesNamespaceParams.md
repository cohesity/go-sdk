# RecoverKubernetesNamespaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubernetesTargetParams** | Pointer to [**NullableRecoverKubernetesNamespaceParamsKubernetesTargetParams**](RecoverKubernetesNamespaceParamsKubernetesTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VlanConfig** | Pointer to [**NullableRecoverKubernetesNamespaceParamsVlanConfig**](RecoverKubernetesNamespaceParamsVlanConfig.md) |  | [optional] 

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

### GetKubernetesTargetParams

`func (o *RecoverKubernetesNamespaceParams) GetKubernetesTargetParams() RecoverKubernetesNamespaceParamsKubernetesTargetParams`

GetKubernetesTargetParams returns the KubernetesTargetParams field if non-nil, zero value otherwise.

### GetKubernetesTargetParamsOk

`func (o *RecoverKubernetesNamespaceParams) GetKubernetesTargetParamsOk() (*RecoverKubernetesNamespaceParamsKubernetesTargetParams, bool)`

GetKubernetesTargetParamsOk returns a tuple with the KubernetesTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetParams

`func (o *RecoverKubernetesNamespaceParams) SetKubernetesTargetParams(v RecoverKubernetesNamespaceParamsKubernetesTargetParams)`

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


### GetVlanConfig

`func (o *RecoverKubernetesNamespaceParams) GetVlanConfig() RecoverKubernetesNamespaceParamsVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverKubernetesNamespaceParams) GetVlanConfigOk() (*RecoverKubernetesNamespaceParamsVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverKubernetesNamespaceParams) SetVlanConfig(v RecoverKubernetesNamespaceParamsVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverKubernetesNamespaceParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverKubernetesNamespaceParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverKubernetesNamespaceParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



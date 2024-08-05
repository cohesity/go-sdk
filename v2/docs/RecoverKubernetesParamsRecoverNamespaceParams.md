# RecoverKubernetesParamsRecoverNamespaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubernetesTargetParams** | Pointer to [**NullableRecoverKubernetesNamespaceParamsKubernetesTargetParams**](RecoverKubernetesNamespaceParamsKubernetesTargetParams.md) |  | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VlanConfig** | Pointer to [**NullableRecoverKubernetesNamespaceParamsVlanConfig**](RecoverKubernetesNamespaceParamsVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverKubernetesParamsRecoverNamespaceParams

`func NewRecoverKubernetesParamsRecoverNamespaceParams(targetEnvironment string, ) *RecoverKubernetesParamsRecoverNamespaceParams`

NewRecoverKubernetesParamsRecoverNamespaceParams instantiates a new RecoverKubernetesParamsRecoverNamespaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKubernetesParamsRecoverNamespaceParamsWithDefaults

`func NewRecoverKubernetesParamsRecoverNamespaceParamsWithDefaults() *RecoverKubernetesParamsRecoverNamespaceParams`

NewRecoverKubernetesParamsRecoverNamespaceParamsWithDefaults instantiates a new RecoverKubernetesParamsRecoverNamespaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubernetesTargetParams

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) GetKubernetesTargetParams() RecoverKubernetesNamespaceParamsKubernetesTargetParams`

GetKubernetesTargetParams returns the KubernetesTargetParams field if non-nil, zero value otherwise.

### GetKubernetesTargetParamsOk

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) GetKubernetesTargetParamsOk() (*RecoverKubernetesNamespaceParamsKubernetesTargetParams, bool)`

GetKubernetesTargetParamsOk returns a tuple with the KubernetesTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesTargetParams

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) SetKubernetesTargetParams(v RecoverKubernetesNamespaceParamsKubernetesTargetParams)`

SetKubernetesTargetParams sets KubernetesTargetParams field to given value.

### HasKubernetesTargetParams

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) HasKubernetesTargetParams() bool`

HasKubernetesTargetParams returns a boolean if a field has been set.

### SetKubernetesTargetParamsNil

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) SetKubernetesTargetParamsNil(b bool)`

 SetKubernetesTargetParamsNil sets the value for KubernetesTargetParams to be an explicit nil

### UnsetKubernetesTargetParams
`func (o *RecoverKubernetesParamsRecoverNamespaceParams) UnsetKubernetesTargetParams()`

UnsetKubernetesTargetParams ensures that no value is present for KubernetesTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVlanConfig

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) GetVlanConfig() RecoverKubernetesNamespaceParamsVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) GetVlanConfigOk() (*RecoverKubernetesNamespaceParamsVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) SetVlanConfig(v RecoverKubernetesNamespaceParamsVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverKubernetesParamsRecoverNamespaceParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverKubernetesParamsRecoverNamespaceParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



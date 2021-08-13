# RecoverKubernetesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoveryAction** | **string** | Specifies the type of recover action to be performed. | 
**RecoverNamespaceParams** | Pointer to [**NullableRecoverKubernetesNamespaceParams**](RecoverKubernetesNamespaceParams.md) | Specifies the parameters to recover Kubernetes Namespaces. | [optional] 

## Methods

### NewRecoverKubernetesParams

`func NewRecoverKubernetesParams(recoveryAction string, ) *RecoverKubernetesParams`

NewRecoverKubernetesParams instantiates a new RecoverKubernetesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverKubernetesParamsWithDefaults

`func NewRecoverKubernetesParamsWithDefaults() *RecoverKubernetesParams`

NewRecoverKubernetesParamsWithDefaults instantiates a new RecoverKubernetesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoveryAction

`func (o *RecoverKubernetesParams) GetRecoveryAction() string`

GetRecoveryAction returns the RecoveryAction field if non-nil, zero value otherwise.

### GetRecoveryActionOk

`func (o *RecoverKubernetesParams) GetRecoveryActionOk() (*string, bool)`

GetRecoveryActionOk returns a tuple with the RecoveryAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryAction

`func (o *RecoverKubernetesParams) SetRecoveryAction(v string)`

SetRecoveryAction sets RecoveryAction field to given value.


### GetRecoverNamespaceParams

`func (o *RecoverKubernetesParams) GetRecoverNamespaceParams() RecoverKubernetesNamespaceParams`

GetRecoverNamespaceParams returns the RecoverNamespaceParams field if non-nil, zero value otherwise.

### GetRecoverNamespaceParamsOk

`func (o *RecoverKubernetesParams) GetRecoverNamespaceParamsOk() (*RecoverKubernetesNamespaceParams, bool)`

GetRecoverNamespaceParamsOk returns a tuple with the RecoverNamespaceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverNamespaceParams

`func (o *RecoverKubernetesParams) SetRecoverNamespaceParams(v RecoverKubernetesNamespaceParams)`

SetRecoverNamespaceParams sets RecoverNamespaceParams field to given value.

### HasRecoverNamespaceParams

`func (o *RecoverKubernetesParams) HasRecoverNamespaceParams() bool`

HasRecoverNamespaceParams returns a boolean if a field has been set.

### SetRecoverNamespaceParamsNil

`func (o *RecoverKubernetesParams) SetRecoverNamespaceParamsNil(b bool)`

 SetRecoverNamespaceParamsNil sets the value for RecoverNamespaceParams to be an explicit nil

### UnsetRecoverNamespaceParams
`func (o *RecoverKubernetesParams) UnsetRecoverNamespaceParams()`

UnsetRecoverNamespaceParams ensures that no value is present for RecoverNamespaceParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



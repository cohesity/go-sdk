# RecoverOracleAppParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OracleTargetParams** | Pointer to [**NullableOracleTargetParamsForRecoverOracleApp**](OracleTargetParamsForRecoverOracleApp.md) | Specifies the params for recovering to a oracle host. Provided oracle backup should be recovered to same type of target host. For Example: If you have oracle backup taken from a physical host then that should be recovered to physical host only. | [optional] 
**TargetEnvironment** | **string** | Specifies the environment of the recovery target. The corresponding params below must be filled out. | 
**VlanConfig** | Pointer to [**NullableRecoverKubernetesNamespaceParamsVlanConfig**](RecoverKubernetesNamespaceParamsVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverOracleAppParams

`func NewRecoverOracleAppParams(targetEnvironment string, ) *RecoverOracleAppParams`

NewRecoverOracleAppParams instantiates a new RecoverOracleAppParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverOracleAppParamsWithDefaults

`func NewRecoverOracleAppParamsWithDefaults() *RecoverOracleAppParams`

NewRecoverOracleAppParamsWithDefaults instantiates a new RecoverOracleAppParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOracleTargetParams

`func (o *RecoverOracleAppParams) GetOracleTargetParams() OracleTargetParamsForRecoverOracleApp`

GetOracleTargetParams returns the OracleTargetParams field if non-nil, zero value otherwise.

### GetOracleTargetParamsOk

`func (o *RecoverOracleAppParams) GetOracleTargetParamsOk() (*OracleTargetParamsForRecoverOracleApp, bool)`

GetOracleTargetParamsOk returns a tuple with the OracleTargetParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleTargetParams

`func (o *RecoverOracleAppParams) SetOracleTargetParams(v OracleTargetParamsForRecoverOracleApp)`

SetOracleTargetParams sets OracleTargetParams field to given value.

### HasOracleTargetParams

`func (o *RecoverOracleAppParams) HasOracleTargetParams() bool`

HasOracleTargetParams returns a boolean if a field has been set.

### SetOracleTargetParamsNil

`func (o *RecoverOracleAppParams) SetOracleTargetParamsNil(b bool)`

 SetOracleTargetParamsNil sets the value for OracleTargetParams to be an explicit nil

### UnsetOracleTargetParams
`func (o *RecoverOracleAppParams) UnsetOracleTargetParams()`

UnsetOracleTargetParams ensures that no value is present for OracleTargetParams, not even an explicit nil
### GetTargetEnvironment

`func (o *RecoverOracleAppParams) GetTargetEnvironment() string`

GetTargetEnvironment returns the TargetEnvironment field if non-nil, zero value otherwise.

### GetTargetEnvironmentOk

`func (o *RecoverOracleAppParams) GetTargetEnvironmentOk() (*string, bool)`

GetTargetEnvironmentOk returns a tuple with the TargetEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnvironment

`func (o *RecoverOracleAppParams) SetTargetEnvironment(v string)`

SetTargetEnvironment sets TargetEnvironment field to given value.


### GetVlanConfig

`func (o *RecoverOracleAppParams) GetVlanConfig() RecoverKubernetesNamespaceParamsVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverOracleAppParams) GetVlanConfigOk() (*RecoverKubernetesNamespaceParamsVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverOracleAppParams) SetVlanConfig(v RecoverKubernetesNamespaceParamsVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverOracleAppParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverOracleAppParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverOracleAppParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



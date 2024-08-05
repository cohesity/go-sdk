# KvmTargetParamsForRecoverVmRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**KvmVmRecoveryTargetConfigNewSourceConfig**](KvmVmRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableKvmVmRecoveryTargetConfigOriginalSourceConfig**](KvmVmRecoveryTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewKvmTargetParamsForRecoverVmRecoveryTargetConfig

`func NewKvmTargetParamsForRecoverVmRecoveryTargetConfig(recoverToNewSource bool, ) *KvmTargetParamsForRecoverVmRecoveryTargetConfig`

NewKvmTargetParamsForRecoverVmRecoveryTargetConfig instantiates a new KvmTargetParamsForRecoverVmRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmTargetParamsForRecoverVmRecoveryTargetConfigWithDefaults

`func NewKvmTargetParamsForRecoverVmRecoveryTargetConfigWithDefaults() *KvmTargetParamsForRecoverVmRecoveryTargetConfig`

NewKvmTargetParamsForRecoverVmRecoveryTargetConfigWithDefaults instantiates a new KvmTargetParamsForRecoverVmRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) GetNewSourceConfig() KvmVmRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) GetNewSourceConfigOk() (*KvmVmRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) SetNewSourceConfig(v KvmVmRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### GetOriginalSourceConfig

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) GetOriginalSourceConfig() KvmVmRecoveryTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) GetOriginalSourceConfigOk() (*KvmVmRecoveryTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) SetOriginalSourceConfig(v KvmVmRecoveryTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *KvmTargetParamsForRecoverVmRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



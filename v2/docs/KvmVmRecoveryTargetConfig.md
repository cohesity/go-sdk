# KvmVmRecoveryTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**KvmVmRecoveryTargetConfigNewSourceConfig**](KvmVmRecoveryTargetConfigNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableKvmVmRecoveryTargetConfigOriginalSourceConfig**](KvmVmRecoveryTargetConfigOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 

## Methods

### NewKvmVmRecoveryTargetConfig

`func NewKvmVmRecoveryTargetConfig(recoverToNewSource bool, ) *KvmVmRecoveryTargetConfig`

NewKvmVmRecoveryTargetConfig instantiates a new KvmVmRecoveryTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKvmVmRecoveryTargetConfigWithDefaults

`func NewKvmVmRecoveryTargetConfigWithDefaults() *KvmVmRecoveryTargetConfig`

NewKvmVmRecoveryTargetConfigWithDefaults instantiates a new KvmVmRecoveryTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *KvmVmRecoveryTargetConfig) GetNewSourceConfig() KvmVmRecoveryTargetConfigNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *KvmVmRecoveryTargetConfig) GetNewSourceConfigOk() (*KvmVmRecoveryTargetConfigNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *KvmVmRecoveryTargetConfig) SetNewSourceConfig(v KvmVmRecoveryTargetConfigNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *KvmVmRecoveryTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### GetOriginalSourceConfig

`func (o *KvmVmRecoveryTargetConfig) GetOriginalSourceConfig() KvmVmRecoveryTargetConfigOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *KvmVmRecoveryTargetConfig) GetOriginalSourceConfigOk() (*KvmVmRecoveryTargetConfigOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *KvmVmRecoveryTargetConfig) SetOriginalSourceConfig(v KvmVmRecoveryTargetConfigOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *KvmVmRecoveryTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *KvmVmRecoveryTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *KvmVmRecoveryTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *KvmVmRecoveryTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *KvmVmRecoveryTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *KvmVmRecoveryTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



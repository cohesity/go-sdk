# RecoverAcropolisVmTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecoverToNewSource** | **bool** | Specifies the parameter whether the recovery should be performed to a new or an existing Source Target. | 
**NewSourceConfig** | Pointer to [**NullableRecoverAcropolisVmNewSourceConfig**](RecoverAcropolisVmNewSourceConfig.md) | Specifies the new destination Source configuration parameters where the VMs will be recovered. This is mandatory if recoverToNewSource is set to true. | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverAcropolisVmOriginalSourceConfig**](RecoverAcropolisVmOriginalSourceConfig.md) | Specifies the Source configuration if VM&#39;s are being recovered to Original Source. If not specified, all the configuration parameters will be retained. | [optional] 

## Methods

### NewRecoverAcropolisVmTargetConfig

`func NewRecoverAcropolisVmTargetConfig(recoverToNewSource bool, ) *RecoverAcropolisVmTargetConfig`

NewRecoverAcropolisVmTargetConfig instantiates a new RecoverAcropolisVmTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAcropolisVmTargetConfigWithDefaults

`func NewRecoverAcropolisVmTargetConfigWithDefaults() *RecoverAcropolisVmTargetConfig`

NewRecoverAcropolisVmTargetConfigWithDefaults instantiates a new RecoverAcropolisVmTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecoverToNewSource

`func (o *RecoverAcropolisVmTargetConfig) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverAcropolisVmTargetConfig) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverAcropolisVmTargetConfig) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetNewSourceConfig

`func (o *RecoverAcropolisVmTargetConfig) GetNewSourceConfig() RecoverAcropolisVmNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverAcropolisVmTargetConfig) GetNewSourceConfigOk() (*RecoverAcropolisVmNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverAcropolisVmTargetConfig) SetNewSourceConfig(v RecoverAcropolisVmNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverAcropolisVmTargetConfig) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverAcropolisVmTargetConfig) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverAcropolisVmTargetConfig) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverAcropolisVmTargetConfig) GetOriginalSourceConfig() RecoverAcropolisVmOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverAcropolisVmTargetConfig) GetOriginalSourceConfigOk() (*RecoverAcropolisVmOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverAcropolisVmTargetConfig) SetOriginalSourceConfig(v RecoverAcropolisVmOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverAcropolisVmTargetConfig) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverAcropolisVmTargetConfig) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverAcropolisVmTargetConfig) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



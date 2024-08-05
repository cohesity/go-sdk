# RecoverPureSanVolumeParamsPureTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverPureVolumeTargetParamsNewSourceConfig**](RecoverPureVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverPureVolumeTargetParamsOriginalSourceConfig**](RecoverPureVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies whether to recover to a new source. | 
**UseThinClone** | Pointer to **NullableBool** | Specifies whether to use thin clone to restore storage array snapshots. | [optional] 

## Methods

### NewRecoverPureSanVolumeParamsPureTargetParams

`func NewRecoverPureSanVolumeParamsPureTargetParams(recoverToNewSource bool, ) *RecoverPureSanVolumeParamsPureTargetParams`

NewRecoverPureSanVolumeParamsPureTargetParams instantiates a new RecoverPureSanVolumeParamsPureTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPureSanVolumeParamsPureTargetParamsWithDefaults

`func NewRecoverPureSanVolumeParamsPureTargetParamsWithDefaults() *RecoverPureSanVolumeParamsPureTargetParams`

NewRecoverPureSanVolumeParamsPureTargetParamsWithDefaults instantiates a new RecoverPureSanVolumeParamsPureTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetNewSourceConfig() RecoverPureVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetNewSourceConfigOk() (*RecoverPureVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetNewSourceConfig(v RecoverPureVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverPureSanVolumeParamsPureTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverPureSanVolumeParamsPureTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetOriginalSourceConfig() RecoverPureVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetOriginalSourceConfigOk() (*RecoverPureVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetOriginalSourceConfig(v RecoverPureVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverPureSanVolumeParamsPureTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverPureSanVolumeParamsPureTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetUseThinClone

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetUseThinClone() bool`

GetUseThinClone returns the UseThinClone field if non-nil, zero value otherwise.

### GetUseThinCloneOk

`func (o *RecoverPureSanVolumeParamsPureTargetParams) GetUseThinCloneOk() (*bool, bool)`

GetUseThinCloneOk returns a tuple with the UseThinClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseThinClone

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetUseThinClone(v bool)`

SetUseThinClone sets UseThinClone field to given value.

### HasUseThinClone

`func (o *RecoverPureSanVolumeParamsPureTargetParams) HasUseThinClone() bool`

HasUseThinClone returns a boolean if a field has been set.

### SetUseThinCloneNil

`func (o *RecoverPureSanVolumeParamsPureTargetParams) SetUseThinCloneNil(b bool)`

 SetUseThinCloneNil sets the value for UseThinClone to be an explicit nil

### UnsetUseThinClone
`func (o *RecoverPureSanVolumeParamsPureTargetParams) UnsetUseThinClone()`

UnsetUseThinClone ensures that no value is present for UseThinClone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



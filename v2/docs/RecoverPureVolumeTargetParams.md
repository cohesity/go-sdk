# RecoverPureVolumeTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewSourceConfig** | Pointer to [**NullableRecoverPureVolumeTargetParamsNewSourceConfig**](RecoverPureVolumeTargetParamsNewSourceConfig.md) |  | [optional] 
**OriginalSourceConfig** | Pointer to [**NullableRecoverPureVolumeTargetParamsOriginalSourceConfig**](RecoverPureVolumeTargetParamsOriginalSourceConfig.md) |  | [optional] 
**RecoverToNewSource** | **bool** | Specifies whether to recover to a new source. | 
**UseThinClone** | Pointer to **NullableBool** | Specifies whether to use thin clone to restore storage array snapshots. | [optional] 

## Methods

### NewRecoverPureVolumeTargetParams

`func NewRecoverPureVolumeTargetParams(recoverToNewSource bool, ) *RecoverPureVolumeTargetParams`

NewRecoverPureVolumeTargetParams instantiates a new RecoverPureVolumeTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPureVolumeTargetParamsWithDefaults

`func NewRecoverPureVolumeTargetParamsWithDefaults() *RecoverPureVolumeTargetParams`

NewRecoverPureVolumeTargetParamsWithDefaults instantiates a new RecoverPureVolumeTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewSourceConfig

`func (o *RecoverPureVolumeTargetParams) GetNewSourceConfig() RecoverPureVolumeTargetParamsNewSourceConfig`

GetNewSourceConfig returns the NewSourceConfig field if non-nil, zero value otherwise.

### GetNewSourceConfigOk

`func (o *RecoverPureVolumeTargetParams) GetNewSourceConfigOk() (*RecoverPureVolumeTargetParamsNewSourceConfig, bool)`

GetNewSourceConfigOk returns a tuple with the NewSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSourceConfig

`func (o *RecoverPureVolumeTargetParams) SetNewSourceConfig(v RecoverPureVolumeTargetParamsNewSourceConfig)`

SetNewSourceConfig sets NewSourceConfig field to given value.

### HasNewSourceConfig

`func (o *RecoverPureVolumeTargetParams) HasNewSourceConfig() bool`

HasNewSourceConfig returns a boolean if a field has been set.

### SetNewSourceConfigNil

`func (o *RecoverPureVolumeTargetParams) SetNewSourceConfigNil(b bool)`

 SetNewSourceConfigNil sets the value for NewSourceConfig to be an explicit nil

### UnsetNewSourceConfig
`func (o *RecoverPureVolumeTargetParams) UnsetNewSourceConfig()`

UnsetNewSourceConfig ensures that no value is present for NewSourceConfig, not even an explicit nil
### GetOriginalSourceConfig

`func (o *RecoverPureVolumeTargetParams) GetOriginalSourceConfig() RecoverPureVolumeTargetParamsOriginalSourceConfig`

GetOriginalSourceConfig returns the OriginalSourceConfig field if non-nil, zero value otherwise.

### GetOriginalSourceConfigOk

`func (o *RecoverPureVolumeTargetParams) GetOriginalSourceConfigOk() (*RecoverPureVolumeTargetParamsOriginalSourceConfig, bool)`

GetOriginalSourceConfigOk returns a tuple with the OriginalSourceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSourceConfig

`func (o *RecoverPureVolumeTargetParams) SetOriginalSourceConfig(v RecoverPureVolumeTargetParamsOriginalSourceConfig)`

SetOriginalSourceConfig sets OriginalSourceConfig field to given value.

### HasOriginalSourceConfig

`func (o *RecoverPureVolumeTargetParams) HasOriginalSourceConfig() bool`

HasOriginalSourceConfig returns a boolean if a field has been set.

### SetOriginalSourceConfigNil

`func (o *RecoverPureVolumeTargetParams) SetOriginalSourceConfigNil(b bool)`

 SetOriginalSourceConfigNil sets the value for OriginalSourceConfig to be an explicit nil

### UnsetOriginalSourceConfig
`func (o *RecoverPureVolumeTargetParams) UnsetOriginalSourceConfig()`

UnsetOriginalSourceConfig ensures that no value is present for OriginalSourceConfig, not even an explicit nil
### GetRecoverToNewSource

`func (o *RecoverPureVolumeTargetParams) GetRecoverToNewSource() bool`

GetRecoverToNewSource returns the RecoverToNewSource field if non-nil, zero value otherwise.

### GetRecoverToNewSourceOk

`func (o *RecoverPureVolumeTargetParams) GetRecoverToNewSourceOk() (*bool, bool)`

GetRecoverToNewSourceOk returns a tuple with the RecoverToNewSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToNewSource

`func (o *RecoverPureVolumeTargetParams) SetRecoverToNewSource(v bool)`

SetRecoverToNewSource sets RecoverToNewSource field to given value.


### GetUseThinClone

`func (o *RecoverPureVolumeTargetParams) GetUseThinClone() bool`

GetUseThinClone returns the UseThinClone field if non-nil, zero value otherwise.

### GetUseThinCloneOk

`func (o *RecoverPureVolumeTargetParams) GetUseThinCloneOk() (*bool, bool)`

GetUseThinCloneOk returns a tuple with the UseThinClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseThinClone

`func (o *RecoverPureVolumeTargetParams) SetUseThinClone(v bool)`

SetUseThinClone sets UseThinClone field to given value.

### HasUseThinClone

`func (o *RecoverPureVolumeTargetParams) HasUseThinClone() bool`

HasUseThinClone returns a boolean if a field has been set.

### SetUseThinCloneNil

`func (o *RecoverPureVolumeTargetParams) SetUseThinCloneNil(b bool)`

 SetUseThinCloneNil sets the value for UseThinClone to be an explicit nil

### UnsetUseThinClone
`func (o *RecoverPureVolumeTargetParams) UnsetUseThinClone()`

UnsetUseThinClone ensures that no value is present for UseThinClone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



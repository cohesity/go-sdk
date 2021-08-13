# RecoverPureSanVolumeNewSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**NullableRecoveryObjectIdentifier**](RecoveryObjectIdentifier.md) | Specifies the id of the new target parent source to recover the Pure SAN Volume to. This field must be specified if recoverToNewSource is true. | 
**RenameRecoveredVolumeParams** | Pointer to [**NullableRecoveredOrClonedVmsRenameConfig**](RecoveredOrClonedVmsRenameConfig.md) | Specifies params to rename the recovered SAN volumes. If not specified, the original names of the volumes are preserved. | [optional] 

## Methods

### NewRecoverPureSanVolumeNewSourceConfig

`func NewRecoverPureSanVolumeNewSourceConfig(source NullableRecoveryObjectIdentifier, ) *RecoverPureSanVolumeNewSourceConfig`

NewRecoverPureSanVolumeNewSourceConfig instantiates a new RecoverPureSanVolumeNewSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverPureSanVolumeNewSourceConfigWithDefaults

`func NewRecoverPureSanVolumeNewSourceConfigWithDefaults() *RecoverPureSanVolumeNewSourceConfig`

NewRecoverPureSanVolumeNewSourceConfigWithDefaults instantiates a new RecoverPureSanVolumeNewSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *RecoverPureSanVolumeNewSourceConfig) GetSource() RecoveryObjectIdentifier`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RecoverPureSanVolumeNewSourceConfig) GetSourceOk() (*RecoveryObjectIdentifier, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RecoverPureSanVolumeNewSourceConfig) SetSource(v RecoveryObjectIdentifier)`

SetSource sets Source field to given value.


### SetSourceNil

`func (o *RecoverPureSanVolumeNewSourceConfig) SetSourceNil(b bool)`

 SetSourceNil sets the value for Source to be an explicit nil

### UnsetSource
`func (o *RecoverPureSanVolumeNewSourceConfig) UnsetSource()`

UnsetSource ensures that no value is present for Source, not even an explicit nil
### GetRenameRecoveredVolumeParams

`func (o *RecoverPureSanVolumeNewSourceConfig) GetRenameRecoveredVolumeParams() RecoveredOrClonedVmsRenameConfig`

GetRenameRecoveredVolumeParams returns the RenameRecoveredVolumeParams field if non-nil, zero value otherwise.

### GetRenameRecoveredVolumeParamsOk

`func (o *RecoverPureSanVolumeNewSourceConfig) GetRenameRecoveredVolumeParamsOk() (*RecoveredOrClonedVmsRenameConfig, bool)`

GetRenameRecoveredVolumeParamsOk returns a tuple with the RenameRecoveredVolumeParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameRecoveredVolumeParams

`func (o *RecoverPureSanVolumeNewSourceConfig) SetRenameRecoveredVolumeParams(v RecoveredOrClonedVmsRenameConfig)`

SetRenameRecoveredVolumeParams sets RenameRecoveredVolumeParams field to given value.

### HasRenameRecoveredVolumeParams

`func (o *RecoverPureSanVolumeNewSourceConfig) HasRenameRecoveredVolumeParams() bool`

HasRenameRecoveredVolumeParams returns a boolean if a field has been set.

### SetRenameRecoveredVolumeParamsNil

`func (o *RecoverPureSanVolumeNewSourceConfig) SetRenameRecoveredVolumeParamsNil(b bool)`

 SetRenameRecoveredVolumeParamsNil sets the value for RenameRecoveredVolumeParams to be an explicit nil

### UnsetRenameRecoveredVolumeParams
`func (o *RecoverPureSanVolumeNewSourceConfig) UnsetRenameRecoveredVolumeParams()`

UnsetRenameRecoveredVolumeParams ensures that no value is present for RenameRecoveredVolumeParams, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



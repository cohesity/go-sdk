# RecoverGcpFileAndFolderParamsGcpTargetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other files if one of files or folders failed to recover. Default value is false. | [optional] 
**NewTargetConfig** | Pointer to [**NullableGcpTargetParamsForRecoverFileAndFolderNewTargetConfig**](GcpTargetParamsForRecoverFileAndFolderNewTargetConfig.md) |  | [optional] 
**OriginalTargetConfig** | Pointer to [**NullableGcpTargetParamsForRecoverFileAndFolderOriginalTargetConfig**](GcpTargetParamsForRecoverFileAndFolderOriginalTargetConfig.md) |  | [optional] 
**OverwriteExisting** | Pointer to **NullableBool** | Specifies whether to override the existing files. Default is true. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | Specifies whether to preserve original attributes. Default is true. | [optional] 
**RecoverToOriginalTarget** | **NullableBool** | Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified. | 
**VlanConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverFileAndFolderVlanConfig**](AcropolisTargetParamsForRecoverFileAndFolderVlanConfig.md) |  | [optional] 

## Methods

### NewRecoverGcpFileAndFolderParamsGcpTargetParams

`func NewRecoverGcpFileAndFolderParamsGcpTargetParams(recoverToOriginalTarget NullableBool, ) *RecoverGcpFileAndFolderParamsGcpTargetParams`

NewRecoverGcpFileAndFolderParamsGcpTargetParams instantiates a new RecoverGcpFileAndFolderParamsGcpTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverGcpFileAndFolderParamsGcpTargetParamsWithDefaults

`func NewRecoverGcpFileAndFolderParamsGcpTargetParamsWithDefaults() *RecoverGcpFileAndFolderParamsGcpTargetParams`

NewRecoverGcpFileAndFolderParamsGcpTargetParamsWithDefaults instantiates a new RecoverGcpFileAndFolderParamsGcpTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetNewTargetConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetNewTargetConfig() GcpTargetParamsForRecoverFileAndFolderNewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetNewTargetConfigOk() (*GcpTargetParamsForRecoverFileAndFolderNewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetNewTargetConfig(v GcpTargetParamsForRecoverFileAndFolderNewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetOriginalTargetConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetOriginalTargetConfig() GcpTargetParamsForRecoverFileAndFolderOriginalTargetConfig`

GetOriginalTargetConfig returns the OriginalTargetConfig field if non-nil, zero value otherwise.

### GetOriginalTargetConfigOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetOriginalTargetConfigOk() (*GcpTargetParamsForRecoverFileAndFolderOriginalTargetConfig, bool)`

GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTargetConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetOriginalTargetConfig(v GcpTargetParamsForRecoverFileAndFolderOriginalTargetConfig)`

SetOriginalTargetConfig sets OriginalTargetConfig field to given value.

### HasOriginalTargetConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) HasOriginalTargetConfig() bool`

HasOriginalTargetConfig returns a boolean if a field has been set.

### SetOriginalTargetConfigNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetOriginalTargetConfigNil(b bool)`

 SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil

### UnsetOriginalTargetConfig
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetOriginalTargetConfig()`

UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
### GetOverwriteExisting

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### SetOverwriteExistingNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetOverwriteExistingNil(b bool)`

 SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil

### UnsetOverwriteExisting
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetOverwriteExisting()`

UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
### GetPreserveAttributes

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetRecoverToOriginalTarget

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetRecoverToOriginalTarget() bool`

GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field if non-nil, zero value otherwise.

### GetRecoverToOriginalTargetOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetRecoverToOriginalTargetOk() (*bool, bool)`

GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalTarget

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetRecoverToOriginalTarget(v bool)`

SetRecoverToOriginalTarget sets RecoverToOriginalTarget field to given value.


### SetRecoverToOriginalTargetNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetRecoverToOriginalTargetNil(b bool)`

 SetRecoverToOriginalTargetNil sets the value for RecoverToOriginalTarget to be an explicit nil

### UnsetRecoverToOriginalTarget
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetRecoverToOriginalTarget()`

UnsetRecoverToOriginalTarget ensures that no value is present for RecoverToOriginalTarget, not even an explicit nil
### GetVlanConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetVlanConfig() AcropolisTargetParamsForRecoverFileAndFolderVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverFileAndFolderVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetVlanConfig(v AcropolisTargetParamsForRecoverFileAndFolderVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *RecoverGcpFileAndFolderParamsGcpTargetParams) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



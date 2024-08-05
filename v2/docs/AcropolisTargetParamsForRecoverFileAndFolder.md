# AcropolisTargetParamsForRecoverFileAndFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue recovering other files if one of the objects encounters an error. Default is false. | [optional] 
**NewTargetConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverFileAndFolderNewTargetConfig**](AcropolisTargetParamsForRecoverFileAndFolderNewTargetConfig.md) |  | [optional] 
**OriginalTargetConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverFileAndFolderOriginalTargetConfig**](AcropolisTargetParamsForRecoverFileAndFolderOriginalTargetConfig.md) |  | [optional] 
**OverwriteExisting** | Pointer to **NullableBool** | Specifies whether to overwrite the existing files. Default is true. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | Specifies whether to preserve original file/folder attributes. Default is true. | [optional] 
**RecoverToOriginalTarget** | **NullableBool** | Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified. | 
**VlanConfig** | Pointer to [**NullableAcropolisTargetParamsForRecoverFileAndFolderVlanConfig**](AcropolisTargetParamsForRecoverFileAndFolderVlanConfig.md) |  | [optional] 

## Methods

### NewAcropolisTargetParamsForRecoverFileAndFolder

`func NewAcropolisTargetParamsForRecoverFileAndFolder(recoverToOriginalTarget NullableBool, ) *AcropolisTargetParamsForRecoverFileAndFolder`

NewAcropolisTargetParamsForRecoverFileAndFolder instantiates a new AcropolisTargetParamsForRecoverFileAndFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAcropolisTargetParamsForRecoverFileAndFolderWithDefaults

`func NewAcropolisTargetParamsForRecoverFileAndFolderWithDefaults() *AcropolisTargetParamsForRecoverFileAndFolder`

NewAcropolisTargetParamsForRecoverFileAndFolderWithDefaults instantiates a new AcropolisTargetParamsForRecoverFileAndFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetNewTargetConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetNewTargetConfig() AcropolisTargetParamsForRecoverFileAndFolderNewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetNewTargetConfigOk() (*AcropolisTargetParamsForRecoverFileAndFolderNewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetNewTargetConfig(v AcropolisTargetParamsForRecoverFileAndFolderNewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetOriginalTargetConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetOriginalTargetConfig() AcropolisTargetParamsForRecoverFileAndFolderOriginalTargetConfig`

GetOriginalTargetConfig returns the OriginalTargetConfig field if non-nil, zero value otherwise.

### GetOriginalTargetConfigOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetOriginalTargetConfigOk() (*AcropolisTargetParamsForRecoverFileAndFolderOriginalTargetConfig, bool)`

GetOriginalTargetConfigOk returns a tuple with the OriginalTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalTargetConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetOriginalTargetConfig(v AcropolisTargetParamsForRecoverFileAndFolderOriginalTargetConfig)`

SetOriginalTargetConfig sets OriginalTargetConfig field to given value.

### HasOriginalTargetConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) HasOriginalTargetConfig() bool`

HasOriginalTargetConfig returns a boolean if a field has been set.

### SetOriginalTargetConfigNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetOriginalTargetConfigNil(b bool)`

 SetOriginalTargetConfigNil sets the value for OriginalTargetConfig to be an explicit nil

### UnsetOriginalTargetConfig
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetOriginalTargetConfig()`

UnsetOriginalTargetConfig ensures that no value is present for OriginalTargetConfig, not even an explicit nil
### GetOverwriteExisting

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### SetOverwriteExistingNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetOverwriteExistingNil(b bool)`

 SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil

### UnsetOverwriteExisting
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetOverwriteExisting()`

UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
### GetPreserveAttributes

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetRecoverToOriginalTarget

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalTarget() bool`

GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field if non-nil, zero value otherwise.

### GetRecoverToOriginalTargetOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetRecoverToOriginalTargetOk() (*bool, bool)`

GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalTarget

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetRecoverToOriginalTarget(v bool)`

SetRecoverToOriginalTarget sets RecoverToOriginalTarget field to given value.


### SetRecoverToOriginalTargetNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetRecoverToOriginalTargetNil(b bool)`

 SetRecoverToOriginalTargetNil sets the value for RecoverToOriginalTarget to be an explicit nil

### UnsetRecoverToOriginalTarget
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetRecoverToOriginalTarget()`

UnsetRecoverToOriginalTarget ensures that no value is present for RecoverToOriginalTarget, not even an explicit nil
### GetVlanConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetVlanConfig() AcropolisTargetParamsForRecoverFileAndFolderVlanConfig`

GetVlanConfig returns the VlanConfig field if non-nil, zero value otherwise.

### GetVlanConfigOk

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) GetVlanConfigOk() (*AcropolisTargetParamsForRecoverFileAndFolderVlanConfig, bool)`

GetVlanConfigOk returns a tuple with the VlanConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetVlanConfig(v AcropolisTargetParamsForRecoverFileAndFolderVlanConfig)`

SetVlanConfig sets VlanConfig field to given value.

### HasVlanConfig

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) HasVlanConfig() bool`

HasVlanConfig returns a boolean if a field has been set.

### SetVlanConfigNil

`func (o *AcropolisTargetParamsForRecoverFileAndFolder) SetVlanConfigNil(b bool)`

 SetVlanConfigNil sets the value for VlanConfig to be an explicit nil

### UnsetVlanConfig
`func (o *AcropolisTargetParamsForRecoverFileAndFolder) UnsetVlanConfig()`

UnsetVlanConfig ensures that no value is present for VlanConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



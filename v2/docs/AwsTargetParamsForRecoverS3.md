# AwsTargetParamsForRecoverS3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinueOnError** | Pointer to **NullableBool** | Specifies whether to continue restore on receiving error or not. Default is true. | [optional] 
**NewTargetConfig** | Pointer to [**NullableAwsTargetParamsForRecoverS3NewTargetConfig**](AwsTargetParamsForRecoverS3NewTargetConfig.md) |  | [optional] 
**ObjectPrefix** | Pointer to **NullableString** | Specifies the prefix to be added to all the objects being recovered. | [optional] 
**OverwriteExisting** | Pointer to **NullableBool** | Specifies whether to override the existing objects. Default is false. | [optional] 
**PreserveAttributes** | Pointer to **NullableBool** | Specifies whether to preserve the objects attributes at the time of restore. Default is true. | [optional] 
**RecoverToOriginalTarget** | **NullableBool** | Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified. | 

## Methods

### NewAwsTargetParamsForRecoverS3

`func NewAwsTargetParamsForRecoverS3(recoverToOriginalTarget NullableBool, ) *AwsTargetParamsForRecoverS3`

NewAwsTargetParamsForRecoverS3 instantiates a new AwsTargetParamsForRecoverS3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsTargetParamsForRecoverS3WithDefaults

`func NewAwsTargetParamsForRecoverS3WithDefaults() *AwsTargetParamsForRecoverS3`

NewAwsTargetParamsForRecoverS3WithDefaults instantiates a new AwsTargetParamsForRecoverS3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *AwsTargetParamsForRecoverS3) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *AwsTargetParamsForRecoverS3) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *AwsTargetParamsForRecoverS3) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *AwsTargetParamsForRecoverS3) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *AwsTargetParamsForRecoverS3) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *AwsTargetParamsForRecoverS3) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetNewTargetConfig

`func (o *AwsTargetParamsForRecoverS3) GetNewTargetConfig() AwsTargetParamsForRecoverS3NewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *AwsTargetParamsForRecoverS3) GetNewTargetConfigOk() (*AwsTargetParamsForRecoverS3NewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *AwsTargetParamsForRecoverS3) SetNewTargetConfig(v AwsTargetParamsForRecoverS3NewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *AwsTargetParamsForRecoverS3) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *AwsTargetParamsForRecoverS3) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *AwsTargetParamsForRecoverS3) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetObjectPrefix

`func (o *AwsTargetParamsForRecoverS3) GetObjectPrefix() string`

GetObjectPrefix returns the ObjectPrefix field if non-nil, zero value otherwise.

### GetObjectPrefixOk

`func (o *AwsTargetParamsForRecoverS3) GetObjectPrefixOk() (*string, bool)`

GetObjectPrefixOk returns a tuple with the ObjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPrefix

`func (o *AwsTargetParamsForRecoverS3) SetObjectPrefix(v string)`

SetObjectPrefix sets ObjectPrefix field to given value.

### HasObjectPrefix

`func (o *AwsTargetParamsForRecoverS3) HasObjectPrefix() bool`

HasObjectPrefix returns a boolean if a field has been set.

### SetObjectPrefixNil

`func (o *AwsTargetParamsForRecoverS3) SetObjectPrefixNil(b bool)`

 SetObjectPrefixNil sets the value for ObjectPrefix to be an explicit nil

### UnsetObjectPrefix
`func (o *AwsTargetParamsForRecoverS3) UnsetObjectPrefix()`

UnsetObjectPrefix ensures that no value is present for ObjectPrefix, not even an explicit nil
### GetOverwriteExisting

`func (o *AwsTargetParamsForRecoverS3) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *AwsTargetParamsForRecoverS3) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *AwsTargetParamsForRecoverS3) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *AwsTargetParamsForRecoverS3) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### SetOverwriteExistingNil

`func (o *AwsTargetParamsForRecoverS3) SetOverwriteExistingNil(b bool)`

 SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil

### UnsetOverwriteExisting
`func (o *AwsTargetParamsForRecoverS3) UnsetOverwriteExisting()`

UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
### GetPreserveAttributes

`func (o *AwsTargetParamsForRecoverS3) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *AwsTargetParamsForRecoverS3) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *AwsTargetParamsForRecoverS3) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *AwsTargetParamsForRecoverS3) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *AwsTargetParamsForRecoverS3) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *AwsTargetParamsForRecoverS3) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetRecoverToOriginalTarget

`func (o *AwsTargetParamsForRecoverS3) GetRecoverToOriginalTarget() bool`

GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field if non-nil, zero value otherwise.

### GetRecoverToOriginalTargetOk

`func (o *AwsTargetParamsForRecoverS3) GetRecoverToOriginalTargetOk() (*bool, bool)`

GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalTarget

`func (o *AwsTargetParamsForRecoverS3) SetRecoverToOriginalTarget(v bool)`

SetRecoverToOriginalTarget sets RecoverToOriginalTarget field to given value.


### SetRecoverToOriginalTargetNil

`func (o *AwsTargetParamsForRecoverS3) SetRecoverToOriginalTargetNil(b bool)`

 SetRecoverToOriginalTargetNil sets the value for RecoverToOriginalTarget to be an explicit nil

### UnsetRecoverToOriginalTarget
`func (o *AwsTargetParamsForRecoverS3) UnsetRecoverToOriginalTarget()`

UnsetRecoverToOriginalTarget ensures that no value is present for RecoverToOriginalTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



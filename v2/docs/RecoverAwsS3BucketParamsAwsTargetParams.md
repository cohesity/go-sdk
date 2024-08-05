# RecoverAwsS3BucketParamsAwsTargetParams

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

### NewRecoverAwsS3BucketParamsAwsTargetParams

`func NewRecoverAwsS3BucketParamsAwsTargetParams(recoverToOriginalTarget NullableBool, ) *RecoverAwsS3BucketParamsAwsTargetParams`

NewRecoverAwsS3BucketParamsAwsTargetParams instantiates a new RecoverAwsS3BucketParamsAwsTargetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoverAwsS3BucketParamsAwsTargetParamsWithDefaults

`func NewRecoverAwsS3BucketParamsAwsTargetParamsWithDefaults() *RecoverAwsS3BucketParamsAwsTargetParams`

NewRecoverAwsS3BucketParamsAwsTargetParamsWithDefaults instantiates a new RecoverAwsS3BucketParamsAwsTargetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinueOnError

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetContinueOnError() bool`

GetContinueOnError returns the ContinueOnError field if non-nil, zero value otherwise.

### GetContinueOnErrorOk

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetContinueOnErrorOk() (*bool, bool)`

GetContinueOnErrorOk returns a tuple with the ContinueOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinueOnError

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetContinueOnError(v bool)`

SetContinueOnError sets ContinueOnError field to given value.

### HasContinueOnError

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) HasContinueOnError() bool`

HasContinueOnError returns a boolean if a field has been set.

### SetContinueOnErrorNil

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetContinueOnErrorNil(b bool)`

 SetContinueOnErrorNil sets the value for ContinueOnError to be an explicit nil

### UnsetContinueOnError
`func (o *RecoverAwsS3BucketParamsAwsTargetParams) UnsetContinueOnError()`

UnsetContinueOnError ensures that no value is present for ContinueOnError, not even an explicit nil
### GetNewTargetConfig

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetNewTargetConfig() AwsTargetParamsForRecoverS3NewTargetConfig`

GetNewTargetConfig returns the NewTargetConfig field if non-nil, zero value otherwise.

### GetNewTargetConfigOk

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetNewTargetConfigOk() (*AwsTargetParamsForRecoverS3NewTargetConfig, bool)`

GetNewTargetConfigOk returns a tuple with the NewTargetConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTargetConfig

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetNewTargetConfig(v AwsTargetParamsForRecoverS3NewTargetConfig)`

SetNewTargetConfig sets NewTargetConfig field to given value.

### HasNewTargetConfig

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) HasNewTargetConfig() bool`

HasNewTargetConfig returns a boolean if a field has been set.

### SetNewTargetConfigNil

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetNewTargetConfigNil(b bool)`

 SetNewTargetConfigNil sets the value for NewTargetConfig to be an explicit nil

### UnsetNewTargetConfig
`func (o *RecoverAwsS3BucketParamsAwsTargetParams) UnsetNewTargetConfig()`

UnsetNewTargetConfig ensures that no value is present for NewTargetConfig, not even an explicit nil
### GetObjectPrefix

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetObjectPrefix() string`

GetObjectPrefix returns the ObjectPrefix field if non-nil, zero value otherwise.

### GetObjectPrefixOk

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetObjectPrefixOk() (*string, bool)`

GetObjectPrefixOk returns a tuple with the ObjectPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectPrefix

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetObjectPrefix(v string)`

SetObjectPrefix sets ObjectPrefix field to given value.

### HasObjectPrefix

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) HasObjectPrefix() bool`

HasObjectPrefix returns a boolean if a field has been set.

### SetObjectPrefixNil

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetObjectPrefixNil(b bool)`

 SetObjectPrefixNil sets the value for ObjectPrefix to be an explicit nil

### UnsetObjectPrefix
`func (o *RecoverAwsS3BucketParamsAwsTargetParams) UnsetObjectPrefix()`

UnsetObjectPrefix ensures that no value is present for ObjectPrefix, not even an explicit nil
### GetOverwriteExisting

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetOverwriteExisting() bool`

GetOverwriteExisting returns the OverwriteExisting field if non-nil, zero value otherwise.

### GetOverwriteExistingOk

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetOverwriteExistingOk() (*bool, bool)`

GetOverwriteExistingOk returns a tuple with the OverwriteExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteExisting

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetOverwriteExisting(v bool)`

SetOverwriteExisting sets OverwriteExisting field to given value.

### HasOverwriteExisting

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) HasOverwriteExisting() bool`

HasOverwriteExisting returns a boolean if a field has been set.

### SetOverwriteExistingNil

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetOverwriteExistingNil(b bool)`

 SetOverwriteExistingNil sets the value for OverwriteExisting to be an explicit nil

### UnsetOverwriteExisting
`func (o *RecoverAwsS3BucketParamsAwsTargetParams) UnsetOverwriteExisting()`

UnsetOverwriteExisting ensures that no value is present for OverwriteExisting, not even an explicit nil
### GetPreserveAttributes

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetPreserveAttributes() bool`

GetPreserveAttributes returns the PreserveAttributes field if non-nil, zero value otherwise.

### GetPreserveAttributesOk

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetPreserveAttributesOk() (*bool, bool)`

GetPreserveAttributesOk returns a tuple with the PreserveAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveAttributes

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetPreserveAttributes(v bool)`

SetPreserveAttributes sets PreserveAttributes field to given value.

### HasPreserveAttributes

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) HasPreserveAttributes() bool`

HasPreserveAttributes returns a boolean if a field has been set.

### SetPreserveAttributesNil

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetPreserveAttributesNil(b bool)`

 SetPreserveAttributesNil sets the value for PreserveAttributes to be an explicit nil

### UnsetPreserveAttributes
`func (o *RecoverAwsS3BucketParamsAwsTargetParams) UnsetPreserveAttributes()`

UnsetPreserveAttributes ensures that no value is present for PreserveAttributes, not even an explicit nil
### GetRecoverToOriginalTarget

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetRecoverToOriginalTarget() bool`

GetRecoverToOriginalTarget returns the RecoverToOriginalTarget field if non-nil, zero value otherwise.

### GetRecoverToOriginalTargetOk

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) GetRecoverToOriginalTargetOk() (*bool, bool)`

GetRecoverToOriginalTargetOk returns a tuple with the RecoverToOriginalTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalTarget

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetRecoverToOriginalTarget(v bool)`

SetRecoverToOriginalTarget sets RecoverToOriginalTarget field to given value.


### SetRecoverToOriginalTargetNil

`func (o *RecoverAwsS3BucketParamsAwsTargetParams) SetRecoverToOriginalTargetNil(b bool)`

 SetRecoverToOriginalTargetNil sets the value for RecoverToOriginalTarget to be an explicit nil

### UnsetRecoverToOriginalTarget
`func (o *RecoverAwsS3BucketParamsAwsTargetParams) UnsetRecoverToOriginalTarget()`

UnsetRecoverToOriginalTarget ensures that no value is present for RecoverToOriginalTarget, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



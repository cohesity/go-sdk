# KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternatePath** | Pointer to **NullableString** | Specifies the alternate path location to recover files to. | [optional] 
**RecoverToOriginalPath** | **NullableBool** | Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified. | 

## Methods

### NewKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig

`func NewKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig(recoverToOriginalPath NullableBool, ) *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig`

NewKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig instantiates a new KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults

`func NewKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults() *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig`

NewKubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults instantiates a new KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatePath

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.

### HasAlternatePath

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) HasAlternatePath() bool`

HasAlternatePath returns a boolean if a field has been set.

### SetAlternatePathNil

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
### GetRecoverToOriginalPath

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetRecoverToOriginalPath() bool`

GetRecoverToOriginalPath returns the RecoverToOriginalPath field if non-nil, zero value otherwise.

### GetRecoverToOriginalPathOk

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetRecoverToOriginalPathOk() (*bool, bool)`

GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalPath

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetRecoverToOriginalPath(v bool)`

SetRecoverToOriginalPath sets RecoverToOriginalPath field to given value.


### SetRecoverToOriginalPathNil

`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetRecoverToOriginalPathNil(b bool)`

 SetRecoverToOriginalPathNil sets the value for RecoverToOriginalPath to be an explicit nil

### UnsetRecoverToOriginalPath
`func (o *KubernetesTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetRecoverToOriginalPath()`

UnsetRecoverToOriginalPath ensures that no value is present for RecoverToOriginalPath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternatePath** | Pointer to **NullableString** | Specifies the alternate path location to recover files to. | [optional] 
**RecoverToOriginalPath** | **NullableBool** | Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified. | 
**TargetVmCredentials** | Pointer to [**NullableAcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials**](AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewAzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig

`func NewAzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig(recoverToOriginalPath NullableBool, ) *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig`

NewAzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig instantiates a new AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults

`func NewAzureTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults() *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig`

NewAzureTargetParamsForRecoverFileAndFolderOriginalTargetConfigWithDefaults instantiates a new AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternatePath

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.

### HasAlternatePath

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) HasAlternatePath() bool`

HasAlternatePath returns a boolean if a field has been set.

### SetAlternatePathNil

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil
### GetRecoverToOriginalPath

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetRecoverToOriginalPath() bool`

GetRecoverToOriginalPath returns the RecoverToOriginalPath field if non-nil, zero value otherwise.

### GetRecoverToOriginalPathOk

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetRecoverToOriginalPathOk() (*bool, bool)`

GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalPath

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetRecoverToOriginalPath(v bool)`

SetRecoverToOriginalPath sets RecoverToOriginalPath field to given value.


### SetRecoverToOriginalPathNil

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetRecoverToOriginalPathNil(b bool)`

 SetRecoverToOriginalPathNil sets the value for RecoverToOriginalPath to be an explicit nil

### UnsetRecoverToOriginalPath
`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetRecoverToOriginalPath()`

UnsetRecoverToOriginalPath ensures that no value is present for RecoverToOriginalPath, not even an explicit nil
### GetTargetVmCredentials

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetTargetVmCredentials() AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) GetTargetVmCredentialsOk() (*AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetTargetVmCredentials(v AcropolisRecoverFilesOriginalTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *AzureTargetParamsForRecoverFileAndFolderOriginalTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



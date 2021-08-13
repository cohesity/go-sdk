# AzureRecoverFilesOriginalTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetVmCredentials** | [**NullableCredentials**](Credentials.md) | Specifies the credentials for the target VM. | 
**RecoverToOriginalPath** | **NullableBool** | Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified. | 
**AlternatePath** | Pointer to **NullableString** | Specifies the alternate path location to recover files to. | [optional] 

## Methods

### NewAzureRecoverFilesOriginalTargetConfig

`func NewAzureRecoverFilesOriginalTargetConfig(targetVmCredentials NullableCredentials, recoverToOriginalPath NullableBool, ) *AzureRecoverFilesOriginalTargetConfig`

NewAzureRecoverFilesOriginalTargetConfig instantiates a new AzureRecoverFilesOriginalTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRecoverFilesOriginalTargetConfigWithDefaults

`func NewAzureRecoverFilesOriginalTargetConfigWithDefaults() *AzureRecoverFilesOriginalTargetConfig`

NewAzureRecoverFilesOriginalTargetConfigWithDefaults instantiates a new AzureRecoverFilesOriginalTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetVmCredentials

`func (o *AzureRecoverFilesOriginalTargetConfig) GetTargetVmCredentials() Credentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *AzureRecoverFilesOriginalTargetConfig) GetTargetVmCredentialsOk() (*Credentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *AzureRecoverFilesOriginalTargetConfig) SetTargetVmCredentials(v Credentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.


### SetTargetVmCredentialsNil

`func (o *AzureRecoverFilesOriginalTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *AzureRecoverFilesOriginalTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
### GetRecoverToOriginalPath

`func (o *AzureRecoverFilesOriginalTargetConfig) GetRecoverToOriginalPath() bool`

GetRecoverToOriginalPath returns the RecoverToOriginalPath field if non-nil, zero value otherwise.

### GetRecoverToOriginalPathOk

`func (o *AzureRecoverFilesOriginalTargetConfig) GetRecoverToOriginalPathOk() (*bool, bool)`

GetRecoverToOriginalPathOk returns a tuple with the RecoverToOriginalPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverToOriginalPath

`func (o *AzureRecoverFilesOriginalTargetConfig) SetRecoverToOriginalPath(v bool)`

SetRecoverToOriginalPath sets RecoverToOriginalPath field to given value.


### SetRecoverToOriginalPathNil

`func (o *AzureRecoverFilesOriginalTargetConfig) SetRecoverToOriginalPathNil(b bool)`

 SetRecoverToOriginalPathNil sets the value for RecoverToOriginalPath to be an explicit nil

### UnsetRecoverToOriginalPath
`func (o *AzureRecoverFilesOriginalTargetConfig) UnsetRecoverToOriginalPath()`

UnsetRecoverToOriginalPath ensures that no value is present for RecoverToOriginalPath, not even an explicit nil
### GetAlternatePath

`func (o *AzureRecoverFilesOriginalTargetConfig) GetAlternatePath() string`

GetAlternatePath returns the AlternatePath field if non-nil, zero value otherwise.

### GetAlternatePathOk

`func (o *AzureRecoverFilesOriginalTargetConfig) GetAlternatePathOk() (*string, bool)`

GetAlternatePathOk returns a tuple with the AlternatePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternatePath

`func (o *AzureRecoverFilesOriginalTargetConfig) SetAlternatePath(v string)`

SetAlternatePath sets AlternatePath field to given value.

### HasAlternatePath

`func (o *AzureRecoverFilesOriginalTargetConfig) HasAlternatePath() bool`

HasAlternatePath returns a boolean if a field has been set.

### SetAlternatePathNil

`func (o *AzureRecoverFilesOriginalTargetConfig) SetAlternatePathNil(b bool)`

 SetAlternatePathNil sets the value for AlternatePath to be an explicit nil

### UnsetAlternatePath
`func (o *AzureRecoverFilesOriginalTargetConfig) UnsetAlternatePath()`

UnsetAlternatePath ensures that no value is present for AlternatePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



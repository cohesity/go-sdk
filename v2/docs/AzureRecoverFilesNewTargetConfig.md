# AzureRecoverFilesNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the absolute path location to recover files to. | 
**TargetVm** | [**NullableAcropolisRecoverFilesNewTargetConfigTargetVm**](AcropolisRecoverFilesNewTargetConfigTargetVm.md) |  | 
**TargetVmCredentials** | Pointer to [**NullableAcropolisRecoverFilesNewTargetConfigTargetVmCredentials**](AcropolisRecoverFilesNewTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewAzureRecoverFilesNewTargetConfig

`func NewAzureRecoverFilesNewTargetConfig(absolutePath NullableString, targetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm, ) *AzureRecoverFilesNewTargetConfig`

NewAzureRecoverFilesNewTargetConfig instantiates a new AzureRecoverFilesNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureRecoverFilesNewTargetConfigWithDefaults

`func NewAzureRecoverFilesNewTargetConfigWithDefaults() *AzureRecoverFilesNewTargetConfig`

NewAzureRecoverFilesNewTargetConfigWithDefaults instantiates a new AzureRecoverFilesNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *AzureRecoverFilesNewTargetConfig) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *AzureRecoverFilesNewTargetConfig) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *AzureRecoverFilesNewTargetConfig) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *AzureRecoverFilesNewTargetConfig) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *AzureRecoverFilesNewTargetConfig) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetTargetVm

`func (o *AzureRecoverFilesNewTargetConfig) GetTargetVm() AcropolisRecoverFilesNewTargetConfigTargetVm`

GetTargetVm returns the TargetVm field if non-nil, zero value otherwise.

### GetTargetVmOk

`func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmOk() (*AcropolisRecoverFilesNewTargetConfigTargetVm, bool)`

GetTargetVmOk returns a tuple with the TargetVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVm

`func (o *AzureRecoverFilesNewTargetConfig) SetTargetVm(v AcropolisRecoverFilesNewTargetConfigTargetVm)`

SetTargetVm sets TargetVm field to given value.


### SetTargetVmNil

`func (o *AzureRecoverFilesNewTargetConfig) SetTargetVmNil(b bool)`

 SetTargetVmNil sets the value for TargetVm to be an explicit nil

### UnsetTargetVm
`func (o *AzureRecoverFilesNewTargetConfig) UnsetTargetVm()`

UnsetTargetVm ensures that no value is present for TargetVm, not even an explicit nil
### GetTargetVmCredentials

`func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmCredentials() AcropolisRecoverFilesNewTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *AzureRecoverFilesNewTargetConfig) GetTargetVmCredentialsOk() (*AcropolisRecoverFilesNewTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *AzureRecoverFilesNewTargetConfig) SetTargetVmCredentials(v AcropolisRecoverFilesNewTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *AzureRecoverFilesNewTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *AzureRecoverFilesNewTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *AzureRecoverFilesNewTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



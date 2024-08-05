# AwsRecoverFilesNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the path location to recover files to. | 
**TargetVm** | [**NullableAcropolisRecoverFilesNewTargetConfigTargetVm**](AcropolisRecoverFilesNewTargetConfigTargetVm.md) |  | 
**TargetVmCredentials** | Pointer to [**NullableAcropolisRecoverFilesNewTargetConfigTargetVmCredentials**](AcropolisRecoverFilesNewTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewAwsRecoverFilesNewTargetConfig

`func NewAwsRecoverFilesNewTargetConfig(absolutePath NullableString, targetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm, ) *AwsRecoverFilesNewTargetConfig`

NewAwsRecoverFilesNewTargetConfig instantiates a new AwsRecoverFilesNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsRecoverFilesNewTargetConfigWithDefaults

`func NewAwsRecoverFilesNewTargetConfigWithDefaults() *AwsRecoverFilesNewTargetConfig`

NewAwsRecoverFilesNewTargetConfigWithDefaults instantiates a new AwsRecoverFilesNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *AwsRecoverFilesNewTargetConfig) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *AwsRecoverFilesNewTargetConfig) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *AwsRecoverFilesNewTargetConfig) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *AwsRecoverFilesNewTargetConfig) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *AwsRecoverFilesNewTargetConfig) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetTargetVm

`func (o *AwsRecoverFilesNewTargetConfig) GetTargetVm() AcropolisRecoverFilesNewTargetConfigTargetVm`

GetTargetVm returns the TargetVm field if non-nil, zero value otherwise.

### GetTargetVmOk

`func (o *AwsRecoverFilesNewTargetConfig) GetTargetVmOk() (*AcropolisRecoverFilesNewTargetConfigTargetVm, bool)`

GetTargetVmOk returns a tuple with the TargetVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVm

`func (o *AwsRecoverFilesNewTargetConfig) SetTargetVm(v AcropolisRecoverFilesNewTargetConfigTargetVm)`

SetTargetVm sets TargetVm field to given value.


### SetTargetVmNil

`func (o *AwsRecoverFilesNewTargetConfig) SetTargetVmNil(b bool)`

 SetTargetVmNil sets the value for TargetVm to be an explicit nil

### UnsetTargetVm
`func (o *AwsRecoverFilesNewTargetConfig) UnsetTargetVm()`

UnsetTargetVm ensures that no value is present for TargetVm, not even an explicit nil
### GetTargetVmCredentials

`func (o *AwsRecoverFilesNewTargetConfig) GetTargetVmCredentials() AcropolisRecoverFilesNewTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *AwsRecoverFilesNewTargetConfig) GetTargetVmCredentialsOk() (*AcropolisRecoverFilesNewTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *AwsRecoverFilesNewTargetConfig) SetTargetVmCredentials(v AcropolisRecoverFilesNewTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *AwsRecoverFilesNewTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *AwsRecoverFilesNewTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *AwsRecoverFilesNewTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



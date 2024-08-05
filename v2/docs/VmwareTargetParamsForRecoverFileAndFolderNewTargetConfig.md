# VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsolutePath** | **NullableString** | Specifies the path location to recover files to. | 
**RecoverMethod** | **string** | Specifies the method to recover files and folders. | 
**TargetVm** | [**NullableAcropolisRecoverFilesNewTargetConfigTargetVm**](AcropolisRecoverFilesNewTargetConfigTargetVm.md) |  | 
**TargetVmCredentials** | Pointer to [**NullableVmwareRecoverFilesNewTargetConfigTargetVmCredentials**](VmwareRecoverFilesNewTargetConfigTargetVmCredentials.md) |  | [optional] 

## Methods

### NewVmwareTargetParamsForRecoverFileAndFolderNewTargetConfig

`func NewVmwareTargetParamsForRecoverFileAndFolderNewTargetConfig(absolutePath NullableString, recoverMethod string, targetVm NullableAcropolisRecoverFilesNewTargetConfigTargetVm, ) *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig`

NewVmwareTargetParamsForRecoverFileAndFolderNewTargetConfig instantiates a new VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults

`func NewVmwareTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults() *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig`

NewVmwareTargetParamsForRecoverFileAndFolderNewTargetConfigWithDefaults instantiates a new VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsolutePath

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil
### GetRecoverMethod

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetRecoverMethod() string`

GetRecoverMethod returns the RecoverMethod field if non-nil, zero value otherwise.

### GetRecoverMethodOk

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetRecoverMethodOk() (*string, bool)`

GetRecoverMethodOk returns a tuple with the RecoverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMethod

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetRecoverMethod(v string)`

SetRecoverMethod sets RecoverMethod field to given value.


### GetTargetVm

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVm() AcropolisRecoverFilesNewTargetConfigTargetVm`

GetTargetVm returns the TargetVm field if non-nil, zero value otherwise.

### GetTargetVmOk

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVmOk() (*AcropolisRecoverFilesNewTargetConfigTargetVm, bool)`

GetTargetVmOk returns a tuple with the TargetVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVm

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetVm(v AcropolisRecoverFilesNewTargetConfigTargetVm)`

SetTargetVm sets TargetVm field to given value.


### SetTargetVmNil

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetVmNil(b bool)`

 SetTargetVmNil sets the value for TargetVm to be an explicit nil

### UnsetTargetVm
`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) UnsetTargetVm()`

UnsetTargetVm ensures that no value is present for TargetVm, not even an explicit nil
### GetTargetVmCredentials

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVmCredentials() VmwareRecoverFilesNewTargetConfigTargetVmCredentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) GetTargetVmCredentialsOk() (*VmwareRecoverFilesNewTargetConfigTargetVmCredentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetVmCredentials(v VmwareRecoverFilesNewTargetConfigTargetVmCredentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *VmwareTargetParamsForRecoverFileAndFolderNewTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



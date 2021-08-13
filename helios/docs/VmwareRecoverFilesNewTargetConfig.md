# VmwareRecoverFilesNewTargetConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetVm** | [**NullableRecoverTarget**](RecoverTarget.md) | Specifies the target VM to recover files and folders to. | 
**RecoverMethod** | **string** | Specifies the method to recover files and folders. | 
**TargetVmCredentials** | Pointer to [**NullableCredentials**](Credentials.md) | Specifies the credentials for the target VM. This is mandatory if the recoverMethod is AutoDeploy or UseHypervisorApis. | [optional] 
**AbsolutePath** | **NullableString** | Specifies the path location to recover files to. | 

## Methods

### NewVmwareRecoverFilesNewTargetConfig

`func NewVmwareRecoverFilesNewTargetConfig(targetVm NullableRecoverTarget, recoverMethod string, absolutePath NullableString, ) *VmwareRecoverFilesNewTargetConfig`

NewVmwareRecoverFilesNewTargetConfig instantiates a new VmwareRecoverFilesNewTargetConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVmwareRecoverFilesNewTargetConfigWithDefaults

`func NewVmwareRecoverFilesNewTargetConfigWithDefaults() *VmwareRecoverFilesNewTargetConfig`

NewVmwareRecoverFilesNewTargetConfigWithDefaults instantiates a new VmwareRecoverFilesNewTargetConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetVm

`func (o *VmwareRecoverFilesNewTargetConfig) GetTargetVm() RecoverTarget`

GetTargetVm returns the TargetVm field if non-nil, zero value otherwise.

### GetTargetVmOk

`func (o *VmwareRecoverFilesNewTargetConfig) GetTargetVmOk() (*RecoverTarget, bool)`

GetTargetVmOk returns a tuple with the TargetVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVm

`func (o *VmwareRecoverFilesNewTargetConfig) SetTargetVm(v RecoverTarget)`

SetTargetVm sets TargetVm field to given value.


### SetTargetVmNil

`func (o *VmwareRecoverFilesNewTargetConfig) SetTargetVmNil(b bool)`

 SetTargetVmNil sets the value for TargetVm to be an explicit nil

### UnsetTargetVm
`func (o *VmwareRecoverFilesNewTargetConfig) UnsetTargetVm()`

UnsetTargetVm ensures that no value is present for TargetVm, not even an explicit nil
### GetRecoverMethod

`func (o *VmwareRecoverFilesNewTargetConfig) GetRecoverMethod() string`

GetRecoverMethod returns the RecoverMethod field if non-nil, zero value otherwise.

### GetRecoverMethodOk

`func (o *VmwareRecoverFilesNewTargetConfig) GetRecoverMethodOk() (*string, bool)`

GetRecoverMethodOk returns a tuple with the RecoverMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverMethod

`func (o *VmwareRecoverFilesNewTargetConfig) SetRecoverMethod(v string)`

SetRecoverMethod sets RecoverMethod field to given value.


### GetTargetVmCredentials

`func (o *VmwareRecoverFilesNewTargetConfig) GetTargetVmCredentials() Credentials`

GetTargetVmCredentials returns the TargetVmCredentials field if non-nil, zero value otherwise.

### GetTargetVmCredentialsOk

`func (o *VmwareRecoverFilesNewTargetConfig) GetTargetVmCredentialsOk() (*Credentials, bool)`

GetTargetVmCredentialsOk returns a tuple with the TargetVmCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVmCredentials

`func (o *VmwareRecoverFilesNewTargetConfig) SetTargetVmCredentials(v Credentials)`

SetTargetVmCredentials sets TargetVmCredentials field to given value.

### HasTargetVmCredentials

`func (o *VmwareRecoverFilesNewTargetConfig) HasTargetVmCredentials() bool`

HasTargetVmCredentials returns a boolean if a field has been set.

### SetTargetVmCredentialsNil

`func (o *VmwareRecoverFilesNewTargetConfig) SetTargetVmCredentialsNil(b bool)`

 SetTargetVmCredentialsNil sets the value for TargetVmCredentials to be an explicit nil

### UnsetTargetVmCredentials
`func (o *VmwareRecoverFilesNewTargetConfig) UnsetTargetVmCredentials()`

UnsetTargetVmCredentials ensures that no value is present for TargetVmCredentials, not even an explicit nil
### GetAbsolutePath

`func (o *VmwareRecoverFilesNewTargetConfig) GetAbsolutePath() string`

GetAbsolutePath returns the AbsolutePath field if non-nil, zero value otherwise.

### GetAbsolutePathOk

`func (o *VmwareRecoverFilesNewTargetConfig) GetAbsolutePathOk() (*string, bool)`

GetAbsolutePathOk returns a tuple with the AbsolutePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsolutePath

`func (o *VmwareRecoverFilesNewTargetConfig) SetAbsolutePath(v string)`

SetAbsolutePath sets AbsolutePath field to given value.


### SetAbsolutePathNil

`func (o *VmwareRecoverFilesNewTargetConfig) SetAbsolutePathNil(b bool)`

 SetAbsolutePathNil sets the value for AbsolutePath to be an explicit nil

### UnsetAbsolutePath
`func (o *VmwareRecoverFilesNewTargetConfig) UnsetAbsolutePath()`

UnsetAbsolutePath ensures that no value is present for AbsolutePath, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



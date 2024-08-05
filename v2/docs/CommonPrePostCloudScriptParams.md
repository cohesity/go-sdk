# CommonPrePostCloudScriptParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinuxScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 
**WindowsScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 

## Methods

### NewCommonPrePostCloudScriptParams

`func NewCommonPrePostCloudScriptParams() *CommonPrePostCloudScriptParams`

NewCommonPrePostCloudScriptParams instantiates a new CommonPrePostCloudScriptParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonPrePostCloudScriptParamsWithDefaults

`func NewCommonPrePostCloudScriptParamsWithDefaults() *CommonPrePostCloudScriptParams`

NewCommonPrePostCloudScriptParamsWithDefaults instantiates a new CommonPrePostCloudScriptParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinuxScript

`func (o *CommonPrePostCloudScriptParams) GetLinuxScript() CommonPreBackupScriptParams`

GetLinuxScript returns the LinuxScript field if non-nil, zero value otherwise.

### GetLinuxScriptOk

`func (o *CommonPrePostCloudScriptParams) GetLinuxScriptOk() (*CommonPreBackupScriptParams, bool)`

GetLinuxScriptOk returns a tuple with the LinuxScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxScript

`func (o *CommonPrePostCloudScriptParams) SetLinuxScript(v CommonPreBackupScriptParams)`

SetLinuxScript sets LinuxScript field to given value.

### HasLinuxScript

`func (o *CommonPrePostCloudScriptParams) HasLinuxScript() bool`

HasLinuxScript returns a boolean if a field has been set.

### GetWindowsScript

`func (o *CommonPrePostCloudScriptParams) GetWindowsScript() CommonPreBackupScriptParams`

GetWindowsScript returns the WindowsScript field if non-nil, zero value otherwise.

### GetWindowsScriptOk

`func (o *CommonPrePostCloudScriptParams) GetWindowsScriptOk() (*CommonPreBackupScriptParams, bool)`

GetWindowsScriptOk returns a tuple with the WindowsScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsScript

`func (o *CommonPrePostCloudScriptParams) SetWindowsScript(v CommonPreBackupScriptParams)`

SetWindowsScript sets WindowsScript field to given value.

### HasWindowsScript

`func (o *CommonPrePostCloudScriptParams) HasWindowsScript() bool`

HasWindowsScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



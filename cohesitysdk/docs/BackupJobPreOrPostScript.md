# BackupJobPreOrPostScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupScript** | Pointer to [**ScriptPathAndParams**](ScriptPathAndParams.md) |  | [optional] 
**FullBackupScript** | Pointer to [**ScriptPathAndParams**](ScriptPathAndParams.md) |  | [optional] 
**LogBackupScript** | Pointer to [**ScriptPathAndParams**](ScriptPathAndParams.md) |  | [optional] 
**RemoteHostParams** | Pointer to [**RemoteHostConnectorParams**](RemoteHostConnectorParams.md) |  | [optional] 

## Methods

### NewBackupJobPreOrPostScript

`func NewBackupJobPreOrPostScript() *BackupJobPreOrPostScript`

NewBackupJobPreOrPostScript instantiates a new BackupJobPreOrPostScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupJobPreOrPostScriptWithDefaults

`func NewBackupJobPreOrPostScriptWithDefaults() *BackupJobPreOrPostScript`

NewBackupJobPreOrPostScriptWithDefaults instantiates a new BackupJobPreOrPostScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupScript

`func (o *BackupJobPreOrPostScript) GetBackupScript() ScriptPathAndParams`

GetBackupScript returns the BackupScript field if non-nil, zero value otherwise.

### GetBackupScriptOk

`func (o *BackupJobPreOrPostScript) GetBackupScriptOk() (*ScriptPathAndParams, bool)`

GetBackupScriptOk returns a tuple with the BackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupScript

`func (o *BackupJobPreOrPostScript) SetBackupScript(v ScriptPathAndParams)`

SetBackupScript sets BackupScript field to given value.

### HasBackupScript

`func (o *BackupJobPreOrPostScript) HasBackupScript() bool`

HasBackupScript returns a boolean if a field has been set.

### GetFullBackupScript

`func (o *BackupJobPreOrPostScript) GetFullBackupScript() ScriptPathAndParams`

GetFullBackupScript returns the FullBackupScript field if non-nil, zero value otherwise.

### GetFullBackupScriptOk

`func (o *BackupJobPreOrPostScript) GetFullBackupScriptOk() (*ScriptPathAndParams, bool)`

GetFullBackupScriptOk returns a tuple with the FullBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupScript

`func (o *BackupJobPreOrPostScript) SetFullBackupScript(v ScriptPathAndParams)`

SetFullBackupScript sets FullBackupScript field to given value.

### HasFullBackupScript

`func (o *BackupJobPreOrPostScript) HasFullBackupScript() bool`

HasFullBackupScript returns a boolean if a field has been set.

### GetLogBackupScript

`func (o *BackupJobPreOrPostScript) GetLogBackupScript() ScriptPathAndParams`

GetLogBackupScript returns the LogBackupScript field if non-nil, zero value otherwise.

### GetLogBackupScriptOk

`func (o *BackupJobPreOrPostScript) GetLogBackupScriptOk() (*ScriptPathAndParams, bool)`

GetLogBackupScriptOk returns a tuple with the LogBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupScript

`func (o *BackupJobPreOrPostScript) SetLogBackupScript(v ScriptPathAndParams)`

SetLogBackupScript sets LogBackupScript field to given value.

### HasLogBackupScript

`func (o *BackupJobPreOrPostScript) HasLogBackupScript() bool`

HasLogBackupScript returns a boolean if a field has been set.

### GetRemoteHostParams

`func (o *BackupJobPreOrPostScript) GetRemoteHostParams() RemoteHostConnectorParams`

GetRemoteHostParams returns the RemoteHostParams field if non-nil, zero value otherwise.

### GetRemoteHostParamsOk

`func (o *BackupJobPreOrPostScript) GetRemoteHostParamsOk() (*RemoteHostConnectorParams, bool)`

GetRemoteHostParamsOk returns a tuple with the RemoteHostParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHostParams

`func (o *BackupJobPreOrPostScript) SetRemoteHostParams(v RemoteHostConnectorParams)`

SetRemoteHostParams sets RemoteHostParams field to given value.

### HasRemoteHostParams

`func (o *BackupJobPreOrPostScript) HasRemoteHostParams() bool`

HasRemoteHostParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



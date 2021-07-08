# RemoteJobScript

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullBackupScript** | Pointer to [**NullableRemoteScriptPathAndParams**](RemoteScriptPathAndParams.md) | Specifies the script that should run for the Full (no CBT) backup schedule of a Remote Adapter &#39;kPuppeteer&#39; Job. This field is mandatory if the Policy associated with this Job has a Full (no CBT) backup schedule and this is Remote Adapter &#39;kPuppeteer&#39; Job. | [optional] 
**IncrementalBackupScript** | Pointer to [**NullableRemoteScriptPathAndParams**](RemoteScriptPathAndParams.md) | Specifies the script that should run for the CBT-based backup schedule of a Remote Adapter &#39;kPuppeteer&#39; Job. A CBT-based backup schedule is utilizing Change Block Tracking when capturing Snapshots. This field is mandatory if the Policy associated with this Job has a CBT-based backup schedule and this is Remote Adapter &#39;kPuppeteer&#39; Job. | [optional] 
**LogBackupScript** | Pointer to [**NullableRemoteScriptPathAndParams**](RemoteScriptPathAndParams.md) | Specifies the script that should run for the Log backup schedule of a Remote Adapter &#39;kPuppeteer&#39; Job. This field is mandatory if the Policy associated with this Job has a Log backup schedule and this is Remote Adapter &#39;kPuppeteer&#39; Job. | [optional] 
**RemoteHost** | Pointer to [**NullableRemoteHost**](RemoteHost.md) | Specifies the remote host where the remote scripts are executed. This field must be set for Remote Adapter Jobs. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username that will be used to login to the remote host. For host type &#39;kLinux&#39;, it is expected that user has setup the password-less access. So only username field is required. | [optional] 

## Methods

### NewRemoteJobScript

`func NewRemoteJobScript() *RemoteJobScript`

NewRemoteJobScript instantiates a new RemoteJobScript object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteJobScriptWithDefaults

`func NewRemoteJobScriptWithDefaults() *RemoteJobScript`

NewRemoteJobScriptWithDefaults instantiates a new RemoteJobScript object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullBackupScript

`func (o *RemoteJobScript) GetFullBackupScript() RemoteScriptPathAndParams`

GetFullBackupScript returns the FullBackupScript field if non-nil, zero value otherwise.

### GetFullBackupScriptOk

`func (o *RemoteJobScript) GetFullBackupScriptOk() (*RemoteScriptPathAndParams, bool)`

GetFullBackupScriptOk returns a tuple with the FullBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupScript

`func (o *RemoteJobScript) SetFullBackupScript(v RemoteScriptPathAndParams)`

SetFullBackupScript sets FullBackupScript field to given value.

### HasFullBackupScript

`func (o *RemoteJobScript) HasFullBackupScript() bool`

HasFullBackupScript returns a boolean if a field has been set.

### SetFullBackupScriptNil

`func (o *RemoteJobScript) SetFullBackupScriptNil(b bool)`

 SetFullBackupScriptNil sets the value for FullBackupScript to be an explicit nil

### UnsetFullBackupScript
`func (o *RemoteJobScript) UnsetFullBackupScript()`

UnsetFullBackupScript ensures that no value is present for FullBackupScript, not even an explicit nil
### GetIncrementalBackupScript

`func (o *RemoteJobScript) GetIncrementalBackupScript() RemoteScriptPathAndParams`

GetIncrementalBackupScript returns the IncrementalBackupScript field if non-nil, zero value otherwise.

### GetIncrementalBackupScriptOk

`func (o *RemoteJobScript) GetIncrementalBackupScriptOk() (*RemoteScriptPathAndParams, bool)`

GetIncrementalBackupScriptOk returns a tuple with the IncrementalBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupScript

`func (o *RemoteJobScript) SetIncrementalBackupScript(v RemoteScriptPathAndParams)`

SetIncrementalBackupScript sets IncrementalBackupScript field to given value.

### HasIncrementalBackupScript

`func (o *RemoteJobScript) HasIncrementalBackupScript() bool`

HasIncrementalBackupScript returns a boolean if a field has been set.

### SetIncrementalBackupScriptNil

`func (o *RemoteJobScript) SetIncrementalBackupScriptNil(b bool)`

 SetIncrementalBackupScriptNil sets the value for IncrementalBackupScript to be an explicit nil

### UnsetIncrementalBackupScript
`func (o *RemoteJobScript) UnsetIncrementalBackupScript()`

UnsetIncrementalBackupScript ensures that no value is present for IncrementalBackupScript, not even an explicit nil
### GetLogBackupScript

`func (o *RemoteJobScript) GetLogBackupScript() RemoteScriptPathAndParams`

GetLogBackupScript returns the LogBackupScript field if non-nil, zero value otherwise.

### GetLogBackupScriptOk

`func (o *RemoteJobScript) GetLogBackupScriptOk() (*RemoteScriptPathAndParams, bool)`

GetLogBackupScriptOk returns a tuple with the LogBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupScript

`func (o *RemoteJobScript) SetLogBackupScript(v RemoteScriptPathAndParams)`

SetLogBackupScript sets LogBackupScript field to given value.

### HasLogBackupScript

`func (o *RemoteJobScript) HasLogBackupScript() bool`

HasLogBackupScript returns a boolean if a field has been set.

### SetLogBackupScriptNil

`func (o *RemoteJobScript) SetLogBackupScriptNil(b bool)`

 SetLogBackupScriptNil sets the value for LogBackupScript to be an explicit nil

### UnsetLogBackupScript
`func (o *RemoteJobScript) UnsetLogBackupScript()`

UnsetLogBackupScript ensures that no value is present for LogBackupScript, not even an explicit nil
### GetRemoteHost

`func (o *RemoteJobScript) GetRemoteHost() RemoteHost`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *RemoteJobScript) GetRemoteHostOk() (*RemoteHost, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *RemoteJobScript) SetRemoteHost(v RemoteHost)`

SetRemoteHost sets RemoteHost field to given value.

### HasRemoteHost

`func (o *RemoteJobScript) HasRemoteHost() bool`

HasRemoteHost returns a boolean if a field has been set.

### SetRemoteHostNil

`func (o *RemoteJobScript) SetRemoteHostNil(b bool)`

 SetRemoteHostNil sets the value for RemoteHost to be an explicit nil

### UnsetRemoteHost
`func (o *RemoteJobScript) UnsetRemoteHost()`

UnsetRemoteHost ensures that no value is present for RemoteHost, not even an explicit nil
### GetUsername

`func (o *RemoteJobScript) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RemoteJobScript) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RemoteJobScript) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RemoteJobScript) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RemoteJobScript) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RemoteJobScript) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RemoteAdapterHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostType** | Pointer to **NullableString** | Specifies the Operating system type of the host. | [optional] 
**Hostname** | Pointer to **NullableString** | Specifies the Hostname or IP address of the host where the pre and post script will be run. | [optional] 
**Username** | Pointer to **NullableString** | Specifies the username for the host. | [optional] 
**FullBackupScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 
**IncrementalBackupScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 
**LogBackupScript** | Pointer to [**CommonPreBackupScriptParams**](CommonPreBackupScriptParams.md) |  | [optional] 

## Methods

### NewRemoteAdapterHost

`func NewRemoteAdapterHost() *RemoteAdapterHost`

NewRemoteAdapterHost instantiates a new RemoteAdapterHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteAdapterHostWithDefaults

`func NewRemoteAdapterHostWithDefaults() *RemoteAdapterHost`

NewRemoteAdapterHostWithDefaults instantiates a new RemoteAdapterHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostType

`func (o *RemoteAdapterHost) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *RemoteAdapterHost) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *RemoteAdapterHost) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *RemoteAdapterHost) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### SetHostTypeNil

`func (o *RemoteAdapterHost) SetHostTypeNil(b bool)`

 SetHostTypeNil sets the value for HostType to be an explicit nil

### UnsetHostType
`func (o *RemoteAdapterHost) UnsetHostType()`

UnsetHostType ensures that no value is present for HostType, not even an explicit nil
### GetHostname

`func (o *RemoteAdapterHost) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RemoteAdapterHost) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RemoteAdapterHost) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RemoteAdapterHost) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *RemoteAdapterHost) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *RemoteAdapterHost) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetUsername

`func (o *RemoteAdapterHost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RemoteAdapterHost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RemoteAdapterHost) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RemoteAdapterHost) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RemoteAdapterHost) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RemoteAdapterHost) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetFullBackupScript

`func (o *RemoteAdapterHost) GetFullBackupScript() CommonPreBackupScriptParams`

GetFullBackupScript returns the FullBackupScript field if non-nil, zero value otherwise.

### GetFullBackupScriptOk

`func (o *RemoteAdapterHost) GetFullBackupScriptOk() (*CommonPreBackupScriptParams, bool)`

GetFullBackupScriptOk returns a tuple with the FullBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullBackupScript

`func (o *RemoteAdapterHost) SetFullBackupScript(v CommonPreBackupScriptParams)`

SetFullBackupScript sets FullBackupScript field to given value.

### HasFullBackupScript

`func (o *RemoteAdapterHost) HasFullBackupScript() bool`

HasFullBackupScript returns a boolean if a field has been set.

### GetIncrementalBackupScript

`func (o *RemoteAdapterHost) GetIncrementalBackupScript() CommonPreBackupScriptParams`

GetIncrementalBackupScript returns the IncrementalBackupScript field if non-nil, zero value otherwise.

### GetIncrementalBackupScriptOk

`func (o *RemoteAdapterHost) GetIncrementalBackupScriptOk() (*CommonPreBackupScriptParams, bool)`

GetIncrementalBackupScriptOk returns a tuple with the IncrementalBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrementalBackupScript

`func (o *RemoteAdapterHost) SetIncrementalBackupScript(v CommonPreBackupScriptParams)`

SetIncrementalBackupScript sets IncrementalBackupScript field to given value.

### HasIncrementalBackupScript

`func (o *RemoteAdapterHost) HasIncrementalBackupScript() bool`

HasIncrementalBackupScript returns a boolean if a field has been set.

### GetLogBackupScript

`func (o *RemoteAdapterHost) GetLogBackupScript() CommonPreBackupScriptParams`

GetLogBackupScript returns the LogBackupScript field if non-nil, zero value otherwise.

### GetLogBackupScriptOk

`func (o *RemoteAdapterHost) GetLogBackupScriptOk() (*CommonPreBackupScriptParams, bool)`

GetLogBackupScriptOk returns a tuple with the LogBackupScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogBackupScript

`func (o *RemoteAdapterHost) SetLogBackupScript(v CommonPreBackupScriptParams)`

SetLogBackupScript sets LogBackupScript field to given value.

### HasLogBackupScript

`func (o *RemoteAdapterHost) HasLogBackupScript() bool`

HasLogBackupScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



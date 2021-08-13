# UdaSourceRegistrationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceType** | **string** | Specifies the source type for Universal Data Adapter source. | 
**Hosts** | **[]string** | Specifies the IPs/hostnames for the nodes forming the Universal Data Adapter source cluster. | 
**Credentials** | Pointer to [**UdaSourceRegistrationParamsCredentials**](UdaSourceRegistrationParamsCredentials.md) |  | [optional] 
**ScriptDir** | **string** | Specifies the absolute path of scripts used to interact with the Universal Data Adapter source. | 
**MountView** | Pointer to **NullableBool** | Specifies if SMB/NFS view mounting should be enabled on source. Default value is false. | [optional] 
**ViewParams** | Pointer to [**NullableUdaSourceRegistrationParamsViewParams**](UdaSourceRegistrationParamsViewParams.md) |  | [optional] 
**SourceRegistrationArgs** | Pointer to **NullableString** | Specifies custom arguments to be supplied to the source registration scripts. | [optional] 

## Methods

### NewUdaSourceRegistrationParams

`func NewUdaSourceRegistrationParams(sourceType string, hosts []string, scriptDir string, ) *UdaSourceRegistrationParams`

NewUdaSourceRegistrationParams instantiates a new UdaSourceRegistrationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUdaSourceRegistrationParamsWithDefaults

`func NewUdaSourceRegistrationParamsWithDefaults() *UdaSourceRegistrationParams`

NewUdaSourceRegistrationParamsWithDefaults instantiates a new UdaSourceRegistrationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceType

`func (o *UdaSourceRegistrationParams) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *UdaSourceRegistrationParams) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *UdaSourceRegistrationParams) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetHosts

`func (o *UdaSourceRegistrationParams) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *UdaSourceRegistrationParams) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *UdaSourceRegistrationParams) SetHosts(v []string)`

SetHosts sets Hosts field to given value.


### GetCredentials

`func (o *UdaSourceRegistrationParams) GetCredentials() UdaSourceRegistrationParamsCredentials`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *UdaSourceRegistrationParams) GetCredentialsOk() (*UdaSourceRegistrationParamsCredentials, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *UdaSourceRegistrationParams) SetCredentials(v UdaSourceRegistrationParamsCredentials)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *UdaSourceRegistrationParams) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetScriptDir

`func (o *UdaSourceRegistrationParams) GetScriptDir() string`

GetScriptDir returns the ScriptDir field if non-nil, zero value otherwise.

### GetScriptDirOk

`func (o *UdaSourceRegistrationParams) GetScriptDirOk() (*string, bool)`

GetScriptDirOk returns a tuple with the ScriptDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriptDir

`func (o *UdaSourceRegistrationParams) SetScriptDir(v string)`

SetScriptDir sets ScriptDir field to given value.


### GetMountView

`func (o *UdaSourceRegistrationParams) GetMountView() bool`

GetMountView returns the MountView field if non-nil, zero value otherwise.

### GetMountViewOk

`func (o *UdaSourceRegistrationParams) GetMountViewOk() (*bool, bool)`

GetMountViewOk returns a tuple with the MountView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountView

`func (o *UdaSourceRegistrationParams) SetMountView(v bool)`

SetMountView sets MountView field to given value.

### HasMountView

`func (o *UdaSourceRegistrationParams) HasMountView() bool`

HasMountView returns a boolean if a field has been set.

### SetMountViewNil

`func (o *UdaSourceRegistrationParams) SetMountViewNil(b bool)`

 SetMountViewNil sets the value for MountView to be an explicit nil

### UnsetMountView
`func (o *UdaSourceRegistrationParams) UnsetMountView()`

UnsetMountView ensures that no value is present for MountView, not even an explicit nil
### GetViewParams

`func (o *UdaSourceRegistrationParams) GetViewParams() UdaSourceRegistrationParamsViewParams`

GetViewParams returns the ViewParams field if non-nil, zero value otherwise.

### GetViewParamsOk

`func (o *UdaSourceRegistrationParams) GetViewParamsOk() (*UdaSourceRegistrationParamsViewParams, bool)`

GetViewParamsOk returns a tuple with the ViewParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewParams

`func (o *UdaSourceRegistrationParams) SetViewParams(v UdaSourceRegistrationParamsViewParams)`

SetViewParams sets ViewParams field to given value.

### HasViewParams

`func (o *UdaSourceRegistrationParams) HasViewParams() bool`

HasViewParams returns a boolean if a field has been set.

### SetViewParamsNil

`func (o *UdaSourceRegistrationParams) SetViewParamsNil(b bool)`

 SetViewParamsNil sets the value for ViewParams to be an explicit nil

### UnsetViewParams
`func (o *UdaSourceRegistrationParams) UnsetViewParams()`

UnsetViewParams ensures that no value is present for ViewParams, not even an explicit nil
### GetSourceRegistrationArgs

`func (o *UdaSourceRegistrationParams) GetSourceRegistrationArgs() string`

GetSourceRegistrationArgs returns the SourceRegistrationArgs field if non-nil, zero value otherwise.

### GetSourceRegistrationArgsOk

`func (o *UdaSourceRegistrationParams) GetSourceRegistrationArgsOk() (*string, bool)`

GetSourceRegistrationArgsOk returns a tuple with the SourceRegistrationArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceRegistrationArgs

`func (o *UdaSourceRegistrationParams) SetSourceRegistrationArgs(v string)`

SetSourceRegistrationArgs sets SourceRegistrationArgs field to given value.

### HasSourceRegistrationArgs

`func (o *UdaSourceRegistrationParams) HasSourceRegistrationArgs() bool`

HasSourceRegistrationArgs returns a boolean if a field has been set.

### SetSourceRegistrationArgsNil

`func (o *UdaSourceRegistrationParams) SetSourceRegistrationArgsNil(b bool)`

 SetSourceRegistrationArgsNil sets the value for SourceRegistrationArgs to be an explicit nil

### UnsetSourceRegistrationArgs
`func (o *UdaSourceRegistrationParams) UnsetSourceRegistrationArgs()`

UnsetSourceRegistrationArgs ensures that no value is present for SourceRegistrationArgs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


